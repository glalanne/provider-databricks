package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	authenticationv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// AzureADTokenExchangeAudience is the audience Entra ID expects on a
	// federated token used as a client assertion.
	AzureADTokenExchangeAudience = "api://AzureADTokenExchange"

	// AzureDatabricksScope requests an Entra ID access token for the Azure
	// Databricks first-party application.
	AzureDatabricksScope = "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d/.default"

	// DefaultAzureAuthorityHost is the Entra ID authority for the public cloud.
	DefaultAzureAuthorityHost = "https://login.microsoftonline.com/"

	// serviceAccountTokenExpirationSeconds is the lifetime requested for
	// service account tokens. 600 is the minimum the API server accepts.
	serviceAccountTokenExpirationSeconds = 600

	// tokenRefreshSkew is the minimum remaining lifetime of a cached access
	// token handed out. Tokens are normally refreshed earlier, at half their
	// lifetime, because upjet's async operations keep the token they started
	// with for the whole operation.
	tokenRefreshSkew = 5 * time.Minute

	// maxTokenReuse caps how long a cached access token is reused, whatever its
	// lifetime. Entra ID issues managed identity tokens valid for about 24
	// hours; re-exchanging hourly means revoking the namespace's RBAC or the
	// federated credential takes effect in the provider within an hour.
	maxTokenReuse = time.Hour

	// evictionInterval bounds how often the cache looks for entries to drop.
	evictionInterval = 10 * time.Minute

	// exchangeTimeout bounds a token exchange that the authority accepts but
	// never answers. The reconcile's deadline (three minutes for this
	// provider's controllers) also applies; this is shorter, because a stuck
	// exchange holds its identity's cache lock.
	exchangeTimeout = 30 * time.Second

	errRequestSAToken  = "cannot request a token for service account %s/%s"
	errEmptySAToken    = "token request for service account %s/%s returned an empty token"
	errExchangeRequest = "cannot build the token exchange request"
	errExchangeCall    = "token exchange request failed"
	errExchangeStatus  = "token exchange returned HTTP %d: %s"
	errExchangeDecode  = "cannot decode the token exchange response"
	errExchangeEmpty   = "token exchange response has no access_token"
)

// AccessToken is a bearer token, when it was obtained and when it expires.
type AccessToken struct {
	Value     string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

// needsRefresh reports whether a cached token should be replaced: once it is
// maxTokenReuse old, once half its lifetime has passed, or when less than
// tokenRefreshSkew remains. Handing out only tokens with at least half their
// lifetime left leaves room for long-running operations that keep the token;
// the age cap bounds how long a revoked identity keeps working.
func (t *AccessToken) needsRefresh(now time.Time) bool {
	remaining := t.ExpiresAt.Sub(now)
	return now.Sub(t.IssuedAt) >= maxTokenReuse ||
		remaining < tokenRefreshSkew ||
		remaining < t.ExpiresAt.Sub(t.IssuedAt)/2
}

// TokenSource returns a JWT that identifies a workload, to be presented to an
// Exchanger.
type TokenSource interface {
	// Token returns the JWT.
	Token(ctx context.Context) (string, error)
	// CacheKey identifies the workload the token is issued for. Two sources
	// with the same key produce interchangeable tokens.
	CacheKey() string
}

// Exchanger trades a workload JWT for an access token the downstream API
// accepts.
type Exchanger interface {
	// Audience the workload JWT must carry for this exchange.
	Audience() string
	// Exchange returns an access token for the given JWT.
	Exchange(ctx context.Context, jwt string) (*AccessToken, error)
	// CacheKey identifies the identity the access token is issued for.
	CacheKey() string
}

// ServiceAccountTokenSource requests short-lived tokens for a Kubernetes
// service account through the TokenRequest API. The token's subject is
// system:serviceaccount:<Namespace>:<Name>, which is what binds the resulting
// identity to that namespace.
type ServiceAccountTokenSource struct {
	Client    client.Client
	Namespace string
	Name      string
	// Audiences requested on the token. Must not be empty.
	Audiences []string
}

// Token requests a new service account token.
func (s *ServiceAccountTokenSource) Token(ctx context.Context) (string, error) {
	sa := &corev1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Namespace: s.Namespace, Name: s.Name}}
	expiration := int64(serviceAccountTokenExpirationSeconds)
	tr := &authenticationv1.TokenRequest{
		Spec: authenticationv1.TokenRequestSpec{
			Audiences:         s.Audiences,
			ExpirationSeconds: &expiration,
		},
	}
	if err := s.Client.SubResource("token").Create(ctx, sa, tr); err != nil {
		return "", errors.Wrapf(err, errRequestSAToken, s.Namespace, s.Name)
	}
	if tr.Status.Token == "" {
		return "", errors.Errorf(errEmptySAToken, s.Namespace, s.Name)
	}
	return tr.Status.Token, nil
}

// CacheKey identifies the service account and the audiences requested.
func (s *ServiceAccountTokenSource) CacheKey() string {
	return fmt.Sprintf("serviceaccount:%s/%s|%s", s.Namespace, s.Name, strings.Join(s.Audiences, ","))
}

// AzureADExchanger exchanges a federated JWT for an Entra ID access token
// using the OAuth 2.0 client credentials grant with a JWT client assertion
// (RFC 7523), the same exchange Azure workload identity performs.
type AzureADExchanger struct {
	// HTTPClient used for the exchange. Defaults to a client with a
	// 30-second timeout.
	HTTPClient *http.Client
	// AuthorityHost, e.g. https://login.microsoftonline.com/. Defaults to
	// DefaultAzureAuthorityHost.
	AuthorityHost string
	TenantID      string
	ClientID      string
	// Scope requested. Defaults to AzureDatabricksScope.
	Scope string
	// Now returns the current time. Defaults to time.Now.
	Now func() time.Time
}

type tokenResponse struct {
	AccessToken      string          `json:"access_token"`
	ExpiresIn        json.RawMessage `json:"expires_in"`
	ErrorCode        string          `json:"error"`
	ErrorDescription string          `json:"error_description"`
}

// Exchange performs the client assertion grant.
func (e *AzureADExchanger) Exchange(ctx context.Context, jwt string) (*AccessToken, error) {
	form := url.Values{
		"grant_type":            {"client_credentials"},
		"client_id":             {e.ClientID},
		"scope":                 {e.scope()},
		"client_assertion_type": {"urn:ietf:params:oauth:client-assertion-type:jwt-bearer"},
		"client_assertion":      {jwt},
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, e.tokenEndpoint(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, errors.Wrap(err, errExchangeRequest)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := e.httpClient().Do(req)
	if err != nil {
		return nil, errors.Wrap(err, errExchangeCall)
	}
	defer resp.Body.Close() //nolint:errcheck // nothing useful to do with a close error on a read body
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, errors.Wrap(err, errExchangeDecode)
	}

	var tr tokenResponse
	decodeErr := json.Unmarshal(body, &tr)
	if resp.StatusCode != http.StatusOK {
		msg := strings.TrimSpace(string(body))
		if decodeErr == nil && tr.ErrorCode != "" {
			msg = tr.ErrorCode + ": " + tr.ErrorDescription
		}
		return nil, errors.Errorf(errExchangeStatus, resp.StatusCode, msg)
	}
	if decodeErr != nil {
		return nil, errors.Wrap(decodeErr, errExchangeDecode)
	}
	if tr.AccessToken == "" {
		return nil, errors.New(errExchangeEmpty)
	}
	expiresIn, err := parseExpiresIn(tr.ExpiresIn)
	if err != nil {
		return nil, errors.Wrap(err, errExchangeDecode)
	}
	now := e.now()
	return &AccessToken{Value: tr.AccessToken, IssuedAt: now, ExpiresAt: now.Add(time.Duration(expiresIn) * time.Second)}, nil
}

// Audience is the audience Entra ID requires on a federated client assertion.
func (e *AzureADExchanger) Audience() string {
	return AzureADTokenExchangeAudience
}

// CacheKey identifies the Entra identity and scope the token is issued for.
func (e *AzureADExchanger) CacheKey() string {
	return fmt.Sprintf("azuread:%s|%s|%s|%s", e.authorityHost(), e.TenantID, e.ClientID, e.scope())
}

func (e *AzureADExchanger) tokenEndpoint() string {
	return e.authorityHost() + url.PathEscape(e.TenantID) + "/oauth2/v2.0/token"
}

func (e *AzureADExchanger) authorityHost() string {
	h := e.AuthorityHost
	if h == "" {
		h = DefaultAzureAuthorityHost
	}
	if !strings.HasSuffix(h, "/") {
		h += "/"
	}
	return h
}

func (e *AzureADExchanger) scope() string {
	if e.Scope == "" {
		return AzureDatabricksScope
	}
	return e.Scope
}

// defaultExchangeClient uses the default transport, with a timeout.
var defaultExchangeClient = &http.Client{Timeout: exchangeTimeout}

func (e *AzureADExchanger) httpClient() *http.Client {
	if e.HTTPClient == nil {
		return defaultExchangeClient
	}
	return e.HTTPClient
}

func (e *AzureADExchanger) now() time.Time {
	if e.Now == nil {
		return time.Now()
	}
	return e.Now()
}

// parseExpiresIn accepts expires_in as a JSON number or a string, since
// identity providers differ.
func parseExpiresIn(raw json.RawMessage) (int64, error) {
	if len(raw) == 0 {
		return 0, errors.New("response has no expires_in")
	}
	var n json.Number
	if err := json.Unmarshal(raw, &n); err != nil {
		var s string
		if err := json.Unmarshal(raw, &s); err != nil {
			return 0, errors.Wrap(err, "cannot parse expires_in")
		}
		n = json.Number(s)
	}
	return n.Int64()
}

// TokenCache caches access tokens per workload and identity, so that a token
// is exchanged once per lifetime rather than once per reconcile (upjet builds
// a new Terraform setup on every connect).
type TokenCache struct {
	mu        sync.Mutex
	entries   map[string]*cacheEntry
	lastEvict time.Time
	// Now returns the current time. Defaults to time.Now.
	Now func() time.Time
}

// cacheEntry has its own lock, so a slow exchange for one identity doesn't
// block the others, while concurrent requests for the same identity wait for
// a single exchange.
type cacheEntry struct {
	mu    sync.Mutex
	token *AccessToken
}

// NewTokenCache returns an empty cache.
func NewTokenCache() *TokenCache {
	return &TokenCache{entries: map[string]*cacheEntry{}}
}

// Get returns a cached access token for the source/exchanger pair, or obtains
// a new one when none is cached or the cached one needs refreshing (see
// AccessToken.needsRefresh). Failures are not cached.
func (c *TokenCache) Get(ctx context.Context, src TokenSource, ex Exchanger) (*AccessToken, error) {
	key := src.CacheKey() + "->" + ex.CacheKey()
	now := c.now()

	c.mu.Lock()
	c.evictLocked(now)
	e, ok := c.entries[key]
	if !ok {
		e = &cacheEntry{}
		c.entries[key] = e
	}
	c.mu.Unlock()

	e.mu.Lock()
	defer e.mu.Unlock()
	if e.token != nil && !e.token.needsRefresh(now) {
		return e.token, nil
	}

	jwt, err := src.Token(ctx)
	if err != nil {
		return nil, err
	}
	t, err := ex.Exchange(ctx, jwt)
	if err != nil {
		return nil, err
	}
	e.token = t
	return t, nil
}

// evictLocked drops entries that would be refreshed on their next use anyway
// (or never obtained a token) and that no caller holds, so identities that are
// no longer configured, or old cache keys after a ProviderConfig change, don't
// accumulate. With maxTokenReuse, an unused entry lives at most about an hour
// plus evictionInterval. Runs with c.mu held, at most once per
// evictionInterval.
//
// A caller that looked up an entry but hasn't locked it yet may still use it
// after it is dropped; it then caches into the dropped entry, and the next
// caller exchanges again. That costs one extra exchange, never a wrong token.
func (c *TokenCache) evictLocked(now time.Time) {
	if now.Sub(c.lastEvict) < evictionInterval {
		return
	}
	c.lastEvict = now
	for k, e := range c.entries {
		if !e.mu.TryLock() {
			continue
		}
		if e.token == nil || e.token.needsRefresh(now) {
			delete(c.entries, k)
		}
		e.mu.Unlock()
	}
}

func (c *TokenCache) now() time.Time {
	if c.Now == nil {
		return time.Now()
	}
	return c.Now()
}

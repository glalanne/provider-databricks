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
	"golang.org/x/sync/singleflight"
	authenticationv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// AzureADTokenExchangeAudience is the audience Entra ID requires on a
	// federated client assertion.
	AzureADTokenExchangeAudience = "api://AzureADTokenExchange"

	// AzureDatabricksScope requests an Entra ID access token for the Azure
	// Databricks first-party application.
	AzureDatabricksScope = "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d/.default"

	// DefaultAzureAuthorityHost is the Entra ID authority for the public cloud.
	DefaultAzureAuthorityHost = "https://login.microsoftonline.com/"

	// serviceAccountTokenExpirationSeconds is the lifetime requested for
	// service account tokens. 600 is the minimum the API server accepts.
	serviceAccountTokenExpirationSeconds = 600

	// tokenRefreshSkew is the minimum lifetime left on a cached token handed
	// out. Tokens are normally refreshed at half their lifetime, since upjet's
	// async operations keep their starting token until they finish.
	tokenRefreshSkew = 5 * time.Minute

	// maxTokenReuse caps reuse of a cached token, whatever its lifetime (Entra
	// ID's managed identity tokens last about 24 hours), so revoking the opt-in
	// RBAC or the federated credential takes effect within an hour.
	maxTokenReuse = time.Hour

	// evictionInterval bounds how often the cache looks for entries to drop.
	evictionInterval = 10 * time.Minute

	// exchangeTimeout bounds requesting and exchanging a service account token,
	// even if the authority never answers. It's the only bound, as concurrent
	// callers share the attempt and no caller's deadline cancels it.
	exchangeTimeout = 30 * time.Second

	errTokenPanic      = "obtaining an access token panicked: %v"
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
// maxTokenReuse old, half its lifetime has passed, or less than
// tokenRefreshSkew remains.
func (t *AccessToken) needsRefresh(now time.Time) bool {
	remaining := t.ExpiresAt.Sub(now)
	return now.Sub(t.IssuedAt) >= maxTokenReuse ||
		remaining < tokenRefreshSkew ||
		remaining < t.ExpiresAt.Sub(t.IssuedAt)/2
}

// TokenSource returns a workload's JWT for an Exchanger.
type TokenSource interface {
	// Token returns the JWT.
	Token(ctx context.Context) (string, error)
	// CacheKey identifies the workload: equal keys mean interchangeable tokens.
	CacheKey() string
}

// Exchanger trades a workload JWT for an access token the API accepts.
type Exchanger interface {
	// Audience the workload JWT must carry for this exchange.
	Audience() string
	// Exchange returns an access token for the given JWT.
	Exchange(ctx context.Context, jwt string) (*AccessToken, error)
	// CacheKey identifies the identity the access token is issued for.
	CacheKey() string
}

// ServiceAccountTokenSource requests short-lived service account tokens
// (TokenRequest API). Their subject, system:serviceaccount:<Namespace>:<Name>,
// binds the resulting identity to that namespace.
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

// CacheKey identifies the service account and audiences, quoting each field so
// different sources never share a key.
func (s *ServiceAccountTokenSource) CacheKey() string {
	return fmt.Sprintf("serviceaccount:%q/%q|%q", s.Namespace, s.Name, s.Audiences)
}

// AzureADExchanger exchanges a federated JWT for an Entra ID access token
// using the OAuth 2.0 client credentials grant with a JWT client assertion
// (RFC 7523), as Azure workload identity does.
type AzureADExchanger struct {
	// HTTPClient for the exchange. Defaults to one with a 30-second timeout.
	HTTPClient *http.Client
	// AuthorityHost defaults to DefaultAzureAuthorityHost.
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

// CacheKey identifies the Entra identity and scope, quoting each field so
// different identities never share a key.
func (e *AzureADExchanger) CacheKey() string {
	return fmt.Sprintf("azuread:%q|%q|%q|%q", e.authorityHost(), e.TenantID, e.ClientID, e.scope())
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

// TokenCache caches access tokens per workload and identity, since upjet
// builds a new Terraform setup on every connect.
type TokenCache struct {
	mu        sync.Mutex
	entries   map[string]*AccessToken
	lastEvict time.Time
	// inflight shares one exchange among concurrent requests for an identity,
	// without one identity's slow exchange blocking the others.
	inflight singleflight.Group
	// Now returns the current time. Defaults to time.Now.
	Now func() time.Time
}

// NewTokenCache returns an empty cache.
func NewTokenCache() *TokenCache {
	return &TokenCache{entries: map[string]*AccessToken{}}
}

// Get returns the source/exchanger pair's cached access token, or a new one if
// none is cached or it needs refreshing. Concurrent callers for a pair share
// one attempt; each stops waiting when its ctx is done. Failures aren't cached.
func (c *TokenCache) Get(ctx context.Context, src TokenSource, ex Exchanger) (*AccessToken, error) {
	key := cacheKey(src, ex)
	if t := c.cached(key); t != nil {
		return t, nil
	}
	ch := c.inflight.DoChan(key, func() (any, error) {
		return c.refresh(ctx, key, src, ex)
	})
	select {
	case r := <-ch:
		if r.Err != nil {
			return nil, r.Err
		}
		return r.Val.(*AccessToken), nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// cacheKey combines the source's and the exchanger's keys. Quoting each keeps
// the combination unambiguous, whatever they contain.
func cacheKey(src TokenSource, ex Exchanger) string {
	return fmt.Sprintf("%q->%q", src.CacheKey(), ex.CacheKey())
}

// cached returns the token cached for key if it doesn't need refreshing yet.
func (c *TokenCache) cached(key string) *AccessToken {
	now := c.now()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.evictLocked(now)
	if t := c.entries[key]; t != nil && !t.needsRefresh(now) {
		return t
	}
	return nil
}

// refresh obtains and caches a new token for key. It serves all of key's
// concurrent callers, so it has its own deadline, not the first caller's.
func (c *TokenCache) refresh(ctx context.Context, key string, src TokenSource, ex Exchanger) (t *AccessToken, err error) {
	// singleflight runs this in its own goroutine and re-panics there, where
	// nothing can recover it, so a panic would otherwise crash the provider.
	defer func() {
		if r := recover(); r != nil {
			t, err = nil, errors.Errorf(errTokenPanic, r)
		}
	}()
	// Another refresh may have finished since the caller's lookup.
	if fresh := c.cached(key); fresh != nil {
		return fresh, nil
	}

	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), exchangeTimeout)
	defer cancel()
	jwt, err := src.Token(ctx)
	if err != nil {
		return nil, err
	}
	t, err = ex.Exchange(ctx, jwt)
	if err != nil {
		return nil, err
	}
	c.mu.Lock()
	c.entries[key] = t
	c.mu.Unlock()
	return t, nil
}

// evictLocked drops entries due for refresh, so unused identities and keys
// left by ProviderConfig changes don't accumulate; an unused entry lasts at
// most about maxTokenReuse plus evictionInterval. Runs with c.mu held, at most
// once per evictionInterval.
func (c *TokenCache) evictLocked(now time.Time) {
	if now.Sub(c.lastEvict) < evictionInterval {
		return
	}
	c.lastEvict = now
	for k, t := range c.entries {
		if t.needsRefresh(now) {
			delete(c.entries, k)
		}
	}
}

func (c *TokenCache) now() time.Time {
	if c.Now == nil {
		return time.Now()
	}
	return c.Now()
}

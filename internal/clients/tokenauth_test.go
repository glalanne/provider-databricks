package clients

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	authenticationv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"
)

func TestServiceAccountTokenSource(t *testing.T) {
	sa := &corev1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Namespace: "dp-ns", Name: "dp-canary"}}

	t.Run("requests a short-lived token for the named service account", func(t *testing.T) {
		var got *authenticationv1.TokenRequest
		var gotKey client.ObjectKey
		kube := fake.NewClientBuilder().WithObjects(sa).WithInterceptorFuncs(interceptor.Funcs{
			SubResourceCreate: func(ctx context.Context, c client.Client, sub string, obj client.Object, subObj client.Object, opts ...client.SubResourceCreateOption) error {
				if sub != "token" {
					t.Fatalf("subresource = %q, want token", sub)
				}
				gotKey = client.ObjectKeyFromObject(obj)
				got = subObj.(*authenticationv1.TokenRequest)
				return c.SubResource(sub).Create(ctx, obj, subObj, opts...)
			},
		}).Build()

		src := &ServiceAccountTokenSource{Client: kube, Namespace: "dp-ns", Name: "dp-canary", Audiences: []string{AzureADTokenExchangeAudience}}
		tok, err := src.Token(context.Background())
		if err != nil {
			t.Fatalf("Token() error: %v", err)
		}
		if tok != "fake-token" {
			t.Errorf("Token() = %q, want fake-token", tok)
		}
		if diff := cmp.Diff(client.ObjectKey{Namespace: "dp-ns", Name: "dp-canary"}, gotKey); diff != "" {
			t.Errorf("service account (-want +got):\n%s", diff)
		}
		if diff := cmp.Diff([]string{AzureADTokenExchangeAudience}, got.Spec.Audiences); diff != "" {
			t.Errorf("audiences (-want +got):\n%s", diff)
		}
		if got.Spec.ExpirationSeconds == nil || *got.Spec.ExpirationSeconds != serviceAccountTokenExpirationSeconds {
			t.Errorf("expirationSeconds = %v, want %d", got.Spec.ExpirationSeconds, serviceAccountTokenExpirationSeconds)
		}
	})

	t.Run("fails when the service account doesn't exist", func(t *testing.T) {
		kube := fake.NewClientBuilder().Build()
		src := &ServiceAccountTokenSource{Client: kube, Namespace: "dp-ns", Name: "missing", Audiences: []string{AzureADTokenExchangeAudience}}
		if _, err := src.Token(context.Background()); err == nil {
			t.Fatal("Token() error = nil, want an error")
		}
	})

	t.Run("fails on permission denied", func(t *testing.T) {
		kube := fake.NewClientBuilder().WithObjects(sa).WithInterceptorFuncs(interceptor.Funcs{
			SubResourceCreate: func(context.Context, client.Client, string, client.Object, client.Object, ...client.SubResourceCreateOption) error {
				return errors.New("serviceaccounts \"dp-canary\" is forbidden")
			},
		}).Build()
		src := &ServiceAccountTokenSource{Client: kube, Namespace: "dp-ns", Name: "dp-canary", Audiences: []string{AzureADTokenExchangeAudience}}
		_, err := src.Token(context.Background())
		if err == nil || !strings.Contains(err.Error(), "dp-ns/dp-canary") || !strings.Contains(err.Error(), "forbidden") {
			t.Fatalf("Token() error = %v, want it to name the service account and the cause", err)
		}
	})

	t.Run("cache key names the namespace, account and audiences", func(t *testing.T) {
		a := &ServiceAccountTokenSource{Namespace: "ns-a", Name: "sa", Audiences: []string{"x"}}
		b := &ServiceAccountTokenSource{Namespace: "ns-b", Name: "sa", Audiences: []string{"x"}}
		c := &ServiceAccountTokenSource{Namespace: "ns-a", Name: "sa", Audiences: []string{"y"}}
		if a.CacheKey() == b.CacheKey() || a.CacheKey() == c.CacheKey() {
			t.Errorf("cache keys collide: %q %q %q", a.CacheKey(), b.CacheKey(), c.CacheKey())
		}
	})
}

func TestAzureADExchanger(t *testing.T) {
	now := time.Date(2026, 10, 9, 12, 0, 0, 0, time.UTC)

	t.Run("sends a client assertion grant and returns the access token", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost || r.URL.Path != "/tenant-1/oauth2/v2.0/token" {
				t.Errorf("request = %s %s, want POST /tenant-1/oauth2/v2.0/token", r.Method, r.URL.Path)
			}
			if err := r.ParseForm(); err != nil {
				t.Fatal(err)
			}
			want := map[string]string{
				"grant_type":            "client_credentials",
				"client_id":             "client-1",
				"scope":                 AzureDatabricksScope,
				"client_assertion_type": "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
				"client_assertion":      "the-jwt",
			}
			for k, v := range want {
				if got := r.PostForm.Get(k); got != v {
					t.Errorf("form %s = %q, want %q", k, got, v)
				}
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"token_type":"Bearer","expires_in":3599,"access_token":"entra-token"}`))
		}))
		defer srv.Close()

		ex := &AzureADExchanger{HTTPClient: srv.Client(), AuthorityHost: srv.URL, TenantID: "tenant-1", ClientID: "client-1", Now: func() time.Time { return now }}
		got, err := ex.Exchange(context.Background(), "the-jwt")
		if err != nil {
			t.Fatalf("Exchange() error: %v", err)
		}
		if diff := cmp.Diff(&AccessToken{Value: "entra-token", IssuedAt: now, ExpiresAt: now.Add(3599 * time.Second)}, got); diff != "" {
			t.Errorf("Exchange() (-want +got):\n%s", diff)
		}
	})

	t.Run("accepts expires_in as a string", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"expires_in":"600","access_token":"t"}`))
		}))
		defer srv.Close()
		ex := &AzureADExchanger{HTTPClient: srv.Client(), AuthorityHost: srv.URL, TenantID: "t", ClientID: "c", Now: func() time.Time { return now }}
		got, err := ex.Exchange(context.Background(), "jwt")
		if err != nil || !got.ExpiresAt.Equal(now.Add(600*time.Second)) {
			t.Fatalf("Exchange() = %v, %v; want expiry in 600s", got, err)
		}
	})

	t.Run("reports the Entra error", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error":"invalid_client","error_description":"AADSTS700213: No matching federated identity record found"}`))
		}))
		defer srv.Close()
		ex := &AzureADExchanger{HTTPClient: srv.Client(), AuthorityHost: srv.URL, TenantID: "t", ClientID: "c"}
		_, err := ex.Exchange(context.Background(), "jwt")
		if err == nil || !strings.Contains(err.Error(), "400") || !strings.Contains(err.Error(), "AADSTS700213") {
			t.Fatalf("Exchange() error = %v, want status and Entra message", err)
		}
	})

	t.Run("fails when access_token is missing", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"expires_in":3600}`))
		}))
		defer srv.Close()
		ex := &AzureADExchanger{HTTPClient: srv.Client(), AuthorityHost: srv.URL, TenantID: "t", ClientID: "c"}
		if _, err := ex.Exchange(context.Background(), "jwt"); err == nil {
			t.Fatal("Exchange() error = nil, want an error")
		}
	})

	t.Run("defaults to an HTTP client with a timeout", func(t *testing.T) {
		ex := &AzureADExchanger{TenantID: "t", ClientID: "c"}
		if got := ex.httpClient().Timeout; got != exchangeTimeout {
			t.Errorf("default client timeout = %s, want %s", got, exchangeTimeout)
		}
	})

	t.Run("defaults authority and scope", func(t *testing.T) {
		ex := &AzureADExchanger{TenantID: "t", ClientID: "c"}
		if got, want := ex.tokenEndpoint(), "https://login.microsoftonline.com/t/oauth2/v2.0/token"; got != want {
			t.Errorf("tokenEndpoint() = %q, want %q", got, want)
		}
		if ex.scope() != AzureDatabricksScope {
			t.Errorf("scope() = %q, want %q", ex.scope(), AzureDatabricksScope)
		}
	})
}

type countingSource struct {
	key   string
	calls atomic.Int32
}

func (s *countingSource) Token(context.Context) (string, error) {
	s.calls.Add(1)
	return "jwt-" + s.key, nil
}
func (s *countingSource) CacheKey() string { return s.key }

type countingExchanger struct {
	key      string
	lifetime time.Duration
	now      func() time.Time
	calls    atomic.Int32
	fail     bool
}

func (e *countingExchanger) Exchange(_ context.Context, jwt string) (*AccessToken, error) {
	e.calls.Add(1)
	if e.fail {
		return nil, errors.New("exchange failed")
	}
	now := e.now()
	return &AccessToken{Value: "access-for-" + jwt, IssuedAt: now, ExpiresAt: now.Add(e.lifetime)}, nil
}
func (e *countingExchanger) Audience() string { return "test-audience" }
func (e *countingExchanger) CacheKey() string { return e.key }

func TestTokenCache(t *testing.T) {
	now := time.Date(2026, 10, 9, 12, 0, 0, 0, time.UTC)
	clock := func() time.Time { return now }

	t.Run("reuses a token until half its lifetime has passed", func(t *testing.T) {
		c := NewTokenCache()
		c.Now = clock
		src := &countingSource{key: "sa-a"}
		ex := &countingExchanger{key: "id-a", lifetime: time.Hour, now: clock}

		for range 3 {
			got, err := c.Get(context.Background(), src, ex)
			if err != nil || got.Value != "access-for-jwt-sa-a" {
				t.Fatalf("Get() = %v, %v", got, err)
			}
		}
		if src.calls.Load() != 1 || ex.calls.Load() != 1 {
			t.Errorf("calls: source %d, exchanger %d; want 1 each", src.calls.Load(), ex.calls.Load())
		}

		now = now.Add(29 * time.Minute)
		if _, err := c.Get(context.Background(), src, ex); err != nil {
			t.Fatal(err)
		}
		if ex.calls.Load() != 1 {
			t.Errorf("exchanger calls = %d at 29 of 60 minutes, want 1 (still cached)", ex.calls.Load())
		}

		now = now.Add(2 * time.Minute)
		got, err := c.Get(context.Background(), src, ex)
		if err != nil {
			t.Fatal(err)
		}
		if ex.calls.Load() != 2 {
			t.Errorf("exchanger calls = %d at 31 of 60 minutes, want 2 (refreshed)", ex.calls.Load())
		}
		if remaining := got.ExpiresAt.Sub(now); remaining < 30*time.Minute {
			t.Errorf("handed out a token with %s left, want at least half its lifetime", remaining)
		}
	})

	t.Run("re-exchanges long-lived tokens hourly", func(t *testing.T) {
		c := NewTokenCache()
		c.Now = clock
		src := &countingSource{key: "sa-long"}
		// Entra ID issues managed identity tokens valid for about 24 hours.
		ex := &countingExchanger{key: "id-long", lifetime: 24 * time.Hour, now: clock}

		if _, err := c.Get(context.Background(), src, ex); err != nil {
			t.Fatal(err)
		}
		now = now.Add(59 * time.Minute)
		if _, err := c.Get(context.Background(), src, ex); err != nil {
			t.Fatal(err)
		}
		if ex.calls.Load() != 1 {
			t.Errorf("exchanger calls = %d at 59 minutes, want 1 (still cached)", ex.calls.Load())
		}
		now = now.Add(2 * time.Minute)
		if _, err := c.Get(context.Background(), src, ex); err != nil {
			t.Fatal(err)
		}
		if ex.calls.Load() != 2 {
			t.Errorf("exchanger calls = %d at 61 minutes, want 2 (re-exchanged after maxTokenReuse)", ex.calls.Load())
		}
	})

	t.Run("keeps identities apart", func(t *testing.T) {
		c := NewTokenCache()
		c.Now = clock
		srcA, srcB := &countingSource{key: "sa-a"}, &countingSource{key: "sa-b"}
		ex := &countingExchanger{key: "id", lifetime: time.Hour, now: clock}
		a, _ := c.Get(context.Background(), srcA, ex)
		b, _ := c.Get(context.Background(), srcB, ex)
		if a.Value == b.Value {
			t.Errorf("different service accounts got the same token %q", a.Value)
		}
	})

	t.Run("evicts entries that would be refreshed anyway", func(t *testing.T) {
		c := NewTokenCache()
		c.Now = clock
		ex := &countingExchanger{key: "id", lifetime: 24 * time.Hour, now: clock}
		stale, active := &countingSource{key: "sa-stale"}, &countingSource{key: "sa-active"}
		if _, err := c.Get(context.Background(), stale, ex); err != nil {
			t.Fatal(err)
		}

		// Within maxTokenReuse: nothing is evicted.
		now = now.Add(30 * time.Minute)
		if _, err := c.Get(context.Background(), active, ex); err != nil {
			t.Fatal(err)
		}
		if len(c.entries) != 2 {
			t.Fatalf("entries = %d at 30 minutes, want 2", len(c.entries))
		}

		// The stale entry is past maxTokenReuse; the active one isn't.
		now = now.Add(31 * time.Minute)
		if _, err := c.Get(context.Background(), active, ex); err != nil {
			t.Fatal(err)
		}
		if _, ok := c.entries["sa-stale->id"]; ok {
			t.Error("stale entry kept, want evicted")
		}
		if _, ok := c.entries["sa-active->id"]; !ok {
			t.Error("active entry evicted, want kept")
		}
		if ex.calls.Load() != 2 {
			t.Errorf("exchanger calls = %d, want 2 (the active entry is still cached)", ex.calls.Load())
		}
	})

	t.Run("doesn't evict an entry in use", func(t *testing.T) {
		c := NewTokenCache()
		c.Now = clock
		ex := &countingExchanger{key: "id", lifetime: time.Hour, now: clock}
		src := &countingSource{key: "sa"}
		if _, err := c.Get(context.Background(), src, ex); err != nil {
			t.Fatal(err)
		}
		now = now.Add(2 * time.Hour)
		e := c.entries["sa->id"]
		e.mu.Lock()
		c.mu.Lock()
		c.evictLocked(now)
		c.mu.Unlock()
		e.mu.Unlock()
		if _, ok := c.entries["sa->id"]; !ok {
			t.Error("locked entry evicted, want kept")
		}
	})

	t.Run("refuses the cached token when the hourly re-exchange fails", func(t *testing.T) {
		c := NewTokenCache()
		c.Now = clock
		src := &countingSource{key: "sa-revoked"}
		// Entra ID issues managed identity tokens valid for about 24 hours.
		ex := &countingExchanger{key: "id-revoked", lifetime: 24 * time.Hour, now: clock}
		if _, err := c.Get(context.Background(), src, ex); err != nil {
			t.Fatal(err)
		}

		// A use shortly before maxTokenReuse runs the eviction sweep, so the
		// next Get finds the stale entry rather than an evicted one.
		now = now.Add(maxTokenReuse - 5*time.Minute)
		if _, err := c.Get(context.Background(), src, ex); err != nil {
			t.Fatal(err)
		}

		// The identity is revoked (federated credential, RBAC or identity
		// removed): the cached token is still valid at Entra ID, but past
		// maxTokenReuse it must not be handed out again.
		ex.fail = true
		now = now.Add(5 * time.Minute)
		if _, ok := c.entries["sa-revoked->id-revoked"]; !ok {
			t.Fatal("entry evicted before the re-exchange; the test wouldn't exercise the cached token")
		}
		got, err := c.Get(context.Background(), src, ex)
		if err == nil {
			t.Fatalf("Get() = %q, nil; want an error, not the cached token", got.Value)
		}
		if got != nil {
			t.Errorf("Get() returned a token alongside the error: %q", got.Value)
		}
	})

	t.Run("doesn't cache failures", func(t *testing.T) {
		c := NewTokenCache()
		c.Now = clock
		src := &countingSource{key: "sa"}
		ex := &countingExchanger{key: "id", lifetime: time.Hour, now: clock, fail: true}
		for range 2 {
			if _, err := c.Get(context.Background(), src, ex); err == nil {
				t.Fatal("Get() error = nil, want an error")
			}
		}
		if ex.calls.Load() != 2 {
			t.Errorf("exchanger calls = %d, want 2 (failures retried)", ex.calls.Load())
		}
	})
}

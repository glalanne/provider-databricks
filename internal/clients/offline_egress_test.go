package clients

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"github.com/crossplane/upjet/v2/pkg/diffserver"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	dbclient "github.com/databricks/databricks-sdk-go/client"
	dbconfig "github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/iam"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	namespacedv1beta1 "github.com/glalanne/provider-databricks/apis/namespaced/v1beta1"
)

const egressChildMode = "DATABRICKS_OFFLINE_EGRESS_TEST_CHILD"

// Each test runs in a fresh process: net/http caches proxy environment settings
// on first use. The proxy records CONNECT as well as plaintext requests and never
// forwards them. Unlike replacing DefaultTransport, it also observes SDK-owned
// transports that use ProxyFromEnvironment.
func TestOfflineEgressBoundary(t *testing.T) {
	if mode := os.Getenv(egressChildMode); mode != "" {
		runOfflineEgressChild(t, mode)
		return
	}
	for _, tc := range []struct {
		mode       string
		wantEgress bool
	}{
		{mode: "offline"},
		{mode: "unguarded-discovery", wantEgress: true},
		{mode: "unguarded-api", wantEgress: true},
		{mode: "unguarded-auth", wantEgress: true},
		{mode: "unguarded-diff", wantEgress: true},
		{mode: "default-transport", wantEgress: true},
		{mode: "separate-transport", wantEgress: true},
	} {
		t.Run(tc.mode, func(t *testing.T) {
			var mu sync.Mutex
			var requests []string
			proxy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				mu.Lock()
				requests = append(requests, r.Method+" "+r.Host)
				mu.Unlock()
				// A non-retryable response keeps positive controls fast.
				http.Error(w, "blocked by offline egress test", http.StatusForbidden)
			}))
			defer proxy.Close()
			binary, err := os.Executable()
			if err != nil {
				t.Fatal(err)
			}
			ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
			defer cancel()
			cmd := exec.CommandContext(ctx, binary, "-test.run=^TestOfflineEgressBoundary$", "-test.count=1", "-test.timeout=40s")
			for _, env := range os.Environ() {
				key, _, _ := strings.Cut(env, "=")
				upper := strings.ToUpper(key)
				if strings.HasPrefix(upper, "DATABRICKS_") || strings.HasPrefix(upper, "ARM_") ||
					strings.HasPrefix(upper, "AZURE_") || strings.HasPrefix(upper, "GOOGLE_") ||
					strings.HasPrefix(upper, "MSI_") || strings.HasPrefix(upper, "IDENTITY_") {
					continue
				}
				switch upper {
				case egressChildMode, "HTTP_PROXY", "HTTPS_PROXY", "ALL_PROXY", "NO_PROXY":
					continue
				}
				cmd.Env = append(cmd.Env, env)
			}
			cmd.Env = append(cmd.Env, egressChildMode+"="+tc.mode,
				"HTTP_PROXY="+proxy.URL, "HTTPS_PROXY="+proxy.URL, "ALL_PROXY="+proxy.URL,
				"NO_PROXY=", "http_proxy="+proxy.URL, "https_proxy="+proxy.URL, "all_proxy="+proxy.URL, "no_proxy=")
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Errorf("egress test subprocess failed: %v\n%s", err, output)
			}
			mu.Lock()
			seen := append([]string(nil), requests...)
			mu.Unlock()
			if tc.wantEgress && len(seen) == 0 {
				t.Fatal("positive control did not reach the proxy; the egress tripwire is ineffective")
			}
			if !tc.wantEgress && len(seen) != 0 {
				t.Fatalf("offline client leaked HTTP requests: %v\n%s", seen, output)
			}
		})
	}
}

func runOfflineEgressChild(t *testing.T, mode string) {
	t.Helper()
	t.Setenv("DATABRICKS_CONFIG_FILE", filepath.Join(t.TempDir(), "missing-databrickscfg"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	const host = "https://workspace.cloud.databricks.com"
	noDiscovery := func(context.Context, string) (*dbconfig.HostMetadata, error) { return nil, nil }
	switch mode {
	case "offline":
		exerciseOfflineEgressPaths(t)
	case "unguarded-discovery":
		// Deliberately retain normal provider Configure and SDK discovery.
		ps := terraform.Setup{Configuration: map[string]any{
			keyHost: host, keyAuthType: authTypePAT, keyAuthToken: offlineToken,
		}}
		if err := configureNoForkDBXClient(ctx, &ps, *offlineTestProvider(t)); err != nil {
			t.Fatal(err)
		}
	case "unguarded-api", "unguarded-auth":
		cfg := &dbconfig.Config{
			Host: host, Loaders: []dbconfig.Loader{offlineConfigLoader{}},
			HostMetadataResolver: noDiscovery,
		}
		if mode == "unguarded-auth" {
			cfg.AuthType, cfg.ClientID, cfg.ClientSecret = "oauth-m2m", "offline-client", "offline-secret"
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, host, nil)
			if err != nil {
				t.Fatal(err)
			}
			if err := cfg.Authenticate(req); err == nil {
				t.Fatal("expected proxy to reject OAuth authentication")
			}
			return
		}
		cfg.Token, cfg.Credentials = offlineToken, dbconfig.PatCredentials{}
		c, err := dbclient.New(cfg)
		if err != nil {
			t.Fatal(err)
		}
		if err := c.Do(ctx, http.MethodGet, "/api/2.0/test", nil, nil, nil, nil); err == nil {
			t.Fatal("expected proxy to reject unguarded API request")
		}
	case "unguarded-diff":
		p := offlineTestProvider(t)
		ps := offlineTestSetup(t, p, namespacedv1beta1.ProviderConfigSpec{Host: ref(host)})
		c := offlineTestMeta(t, ps)
		cfg, err := c.Config.NewWithWorkspaceHost(host)
		if err != nil {
			t.Fatal(err)
		}
		// Remove only the API guard, keeping offline credentials and loaders.
		cfg.HTTPTransport = nil
		c.DatabricksClient, err = dbclient.New(cfg)
		if err != nil {
			t.Fatal(err)
		}
		_, err = p.ResourcesMap["databricks_cluster_policy"].Diff(ctx, offlineAPIDiffState(),
			tfsdk.NewResourceConfigRaw(map[string]any{"name": "new-name", "definition": `{"a":1}`}), ps.Meta)
		if err == nil {
			t.Fatal("expected proxy to reject the CustomizeDiff workspace lookup")
		}
	case "default-transport", "separate-transport":
		c := &http.Client{Timeout: 5 * time.Second}
		if mode == "separate-transport" {
			transport := &http.Transport{Proxy: http.ProxyFromEnvironment}
			defer transport.CloseIdleConnections()
			c.Transport = transport
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, host, nil)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := c.Do(req)
		if err == nil {
			resp.Body.Close()
			t.Fatal("expected proxy to reject HTTPS CONNECT")
		}
	default:
		t.Fatalf("unknown egress test mode %q", mode)
	}
}

func exerciseOfflineEgressPaths(t *testing.T) {
	t.Helper()
	// ProxyFromEnvironment bypasses loopback destinations. Watch one directly
	// too, so a ProviderConfig or metadata endpoint cannot hide local traffic.
	var localRequests atomic.Int64
	sink := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		localRequests.Add(1)
		http.Error(w, "blocked by local egress test", http.StatusForbidden)
	}))
	t.Cleanup(func() {
		sink.Close()
		if n := localRequests.Load(); n != 0 {
			t.Errorf("offline client leaked %d requests to a loopback endpoint", n)
		}
	})
	// Verify the loopback tripwire independently before arming it.
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, sink.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := sink.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if n := localRequests.Swap(0); n != 1 {
		t.Fatalf("loopback tripwire positive control recorded %d requests, want 1", n)
	}
	dir := t.TempDir()
	profile := filepath.Join(dir, "databrickscfg")
	if err := os.WriteFile(profile, []byte("[DEFAULT]\nhost = https://profile.invalid\nclient_id = profile-client\nclient_secret = profile-secret\n"), 0600); err != nil {
		t.Fatal(err)
	}
	for key, value := range map[string]string{
		"DATABRICKS_HOST":                   "https://environment.invalid",
		"DATABRICKS_TOKEN":                  "environment-token",
		"DATABRICKS_CLIENT_ID":              "environment-client",
		"DATABRICKS_CLIENT_SECRET":          "environment-secret",
		"DATABRICKS_ACCOUNT_ID":             "environment-account",
		"DATABRICKS_CONFIG_FILE":            profile,
		"DATABRICKS_CONFIG_PROFILE":         "DEFAULT",
		"DATABRICKS_AZURE_RESOURCE_ID":      "/subscriptions/offline/resourceGroups/offline/providers/Microsoft.Databricks/workspaces/offline",
		"DATABRICKS_GOOGLE_SERVICE_ACCOUNT": "offline@example.invalid",
		"DATABRICKS_METADATA_SERVICE_URL":   sink.URL + "/metadata",
		"ARM_CLIENT_ID":                     "azure-client",
		"ARM_CLIENT_SECRET":                 "azure-secret",
		"ARM_TENANT_ID":                     "azure-tenant",
		"ARM_USE_MSI":                       "true",
		"ARM_METADATA_HOSTNAME":             "metadata.invalid",
		"MSI_ENDPOINT":                      sink.URL + "/msi/token",
		"IDENTITY_ENDPOINT":                 sink.URL + "/identity/token",
		"AZURE_FEDERATED_TOKEN_FILE":        filepath.Join(dir, "missing-token"),
		"GOOGLE_APPLICATION_CREDENTIALS":    filepath.Join(dir, "missing-google.json"),
	} {
		t.Setenv(key, value)
	}
	p := offlineTestProvider(t)
	for _, source := range []xpv2.CredentialsSource{
		xpv2.CredentialsSourceNone, xpv2.CredentialsSourceSecret, xpv2.CredentialsSourceFilesystem,
		credentialsSourceUserAssignedManagedIdentity, credentialsSourceSystemAssignedManagedIdentity,
		credentialsSourceOIDCTokenFile, credentialsSourceUpbound, namespacedv1beta1.CredentialsSourceOIDCTokenRequest,
	} {
		for _, authType := range []string{"pat", "oauth-m2m", "azure-client-secret", "azure-msi", "azure-cli", "databricks-cli", "google-credentials", "google-id", "metadata-service"} {
			t.Run(string(source)+"/"+authType, func(t *testing.T) {
				t.Setenv("DATABRICKS_AUTH_TYPE", authType)
				ps := offlineTestSetup(t, p, namespacedv1beta1.ProviderConfigSpec{
					Credentials: namespacedv1beta1.ProviderCredentials{Source: source},
				})
				cfg := offlineTestMeta(t, ps).Config
				req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, cfg.Host, nil)
				if err != nil {
					t.Fatal(err)
				}
				if err := cfg.Authenticate(req); err != nil {
					t.Fatal(err)
				}
				if req.Header.Get("Authorization") != "Bearer "+offlineToken {
					t.Fatal("offline authentication did not use the placeholder PAT")
				}
			})
		}
	}
	for _, host := range []string{
		offlineHost, sink.URL, "http://plaintext.invalid", "https://workspace.cloud.databricks.com",
		"https://workspace.azuredatabricks.net", "https://workspace.gcp.databricks.com",
		"https://accounts.cloud.databricks.com", "https://accounts.azuredatabricks.net",
		"https://accounts.gcp.databricks.com",
	} {
		t.Run(host, func(t *testing.T) {
			ps := offlineTestSetup(t, p, namespacedv1beta1.ProviderConfigSpec{Host: ref(host), AccountID: ref("offline-account")})
			c := offlineTestMeta(t, ps)
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			blocked := func(err error) {
				t.Helper()
				if !errors.Is(err, diffserver.ErrDiffComputationNotSupported) {
					t.Fatalf("expected unsupported diff, got %v", err)
				}
			}
			for _, method := range []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete} {
				blocked(c.Do(ctx, method, "/api/2.0/test", nil, nil, nil, nil))
			}
			blocked(c.Do(ctx, http.MethodGet, "https://absolute-target.invalid/api/2.0/test", nil, nil, nil, nil))
			if c.Config.IsAccountClient() {
				a, err := c.AccountClient()
				if err != nil {
					t.Fatal(err)
				}
				_, err = a.Groups.ListAll(ctx, iam.ListAccountGroupsRequest{})
				blocked(err)
			} else {
				w, err := c.WorkspaceClient()
				if err != nil {
					t.Fatal(err)
				}
				_, err = w.CurrentUser.Me(ctx, iam.MeRequest{})
				blocked(err)
			}
			for _, target := range []string{"https://derived.invalid", "http://plaintext-derived.invalid", sink.URL} {
				derived, err := c.ClientForHost(ctx, target)
				if err != nil {
					t.Fatal(err)
				}
				blocked(derived.Get(ctx, "/api/2.0/test", nil, nil))
			}
			// Exercise the real Terraform CustomizeDiff path with usable meta.
			if c.Config.IsAccountClient() {
				return
			}
			_, err := p.ResourcesMap["databricks_cluster_policy"].Diff(ctx,
				&tfsdk.InstanceState{ID: "policy-id", Attributes: map[string]string{"id": "policy-id", "name": "old-name", "definition": `{"a":1}`}},
				tfsdk.NewResourceConfigRaw(map[string]any{"name": "new-name", "definition": `{"a":1}`}), ps.Meta)
			if err != nil {
				t.Fatal(err)
			}
			_, err = p.ResourcesMap["databricks_cluster_policy"].Diff(ctx, offlineAPIDiffState(),
				tfsdk.NewResourceConfigRaw(map[string]any{"name": "new-name", "definition": `{"a":1}`}), ps.Meta)
			if err == nil {
				t.Fatal("expected API-dependent CustomizeDiff to fail offline")
			}
		})
	}
}

func offlineAPIDiffState() *tfsdk.InstanceState {
	return &tfsdk.InstanceState{ID: "policy-id", Attributes: map[string]string{
		"id": "policy-id", "name": "old-name", "definition": `{"a":1}`,
		"provider_config.#": "1", "provider_config.0.workspace_id": "123",
	}}
}

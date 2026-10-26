package clients

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"github.com/crossplane/upjet/v2/pkg/diffserver"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	dbconfig "github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/xpprovider"
	"github.com/google/go-cmp/cmp"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"

	clusterv1beta1 "github.com/glalanne/provider-databricks/apis/cluster/v1beta1"
	clusterworkspacev1beta1 "github.com/glalanne/provider-databricks/apis/cluster/workspace/v1beta1"
	namespacedv1beta1 "github.com/glalanne/provider-databricks/apis/namespaced/v1beta1"
	workspacev1beta1 "github.com/glalanne/provider-databricks/apis/namespaced/workspace/v1beta1"
)

type offlineRoundTripFunc func(*http.Request) (*http.Response, error)

func (f offlineRoundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func forbidOfflineEgress(t *testing.T) {
	t.Helper()
	var calls atomic.Int64
	original := http.DefaultTransport
	http.DefaultTransport = offlineRoundTripFunc(func(*http.Request) (*http.Response, error) {
		calls.Add(1)
		return nil, errors.New("unexpected network access in offline test")
	})
	t.Cleanup(func() {
		http.DefaultTransport = original
		if n := calls.Load(); n != 0 {
			t.Errorf("offline operation attempted %d network requests", n)
		}
	})
}

func offlineTestProvider(t *testing.T) *schema.Provider {
	t.Helper()
	_, p, err := xpprovider.GetProvider(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	return p
}

func offlineTestSetupBuilder(t *testing.T, p *schema.Provider) terraform.SetupFn {
	t.Helper()
	fw, _, err := xpprovider.GetProvider(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	return OfflineTerraformSetupBuilder(fw, p)
}

func offlineTestSetup(t *testing.T, p *schema.Provider, spec namespacedv1beta1.ProviderConfigSpec) terraform.Setup {
	t.Helper()
	s := testScheme(t)
	pc := &namespacedv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: "default", Namespace: "test"},
		Spec:       spec,
	}
	kube := fake.NewClientBuilder().WithScheme(s).WithObjects(pc).WithInterceptorFuncs(interceptor.Funcs{
		Get: func(ctx context.Context, c client.WithWatch, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
			switch obj.(type) {
			case *namespacedv1beta1.ProviderConfig, *namespacedv1beta1.ProviderConfigUsage:
			default:
				t.Errorf("offline setup read credential object %T", obj)
				return errors.New("credential reads are forbidden")
			}
			return c.Get(ctx, key, obj, opts...)
		},
		SubResourceCreate: func(context.Context, client.Client, string, client.Object, client.Object, ...client.SubResourceCreateOption) error {
			t.Error("offline setup requested a service account token")
			return errors.New("token requests are forbidden")
		},
	}).Build()
	mg := &workspacev1beta1.Notebook{ObjectMeta: metav1.ObjectMeta{Name: "notebook", Namespace: "test", UID: "test-uid"}}
	mg.SetProviderConfigReference(&xpv2.ProviderConfigReference{Kind: "ProviderConfig", Name: "default"})
	ps, err := offlineTestSetupBuilder(t, p)(context.Background(), kube, mg)
	if err != nil {
		t.Fatal(err)
	}
	return ps
}

func offlineTestMeta(t *testing.T, ps terraform.Setup) *common.DatabricksClient {
	t.Helper()
	meta, ok := ps.Meta.(*common.DatabricksClient)
	if !ok || meta == nil || meta.DatabricksClient == nil {
		t.Fatalf("expected initialized *common.DatabricksClient, got %T", ps.Meta)
	}
	return meta
}

func TestOfflineAuthentication(t *testing.T) {
	forbidOfflineEgress(t)
	// These must neither supply real credentials nor cause conflicting-auth
	// validation, profile reads, CLI execution, or metadata discovery.
	for key, value := range map[string]string{
		"DATABRICKS_HOST":                "https://environment.invalid",
		"DATABRICKS_TOKEN":               "environment-secret",
		"DATABRICKS_AUTH_TYPE":           "azure-cli",
		"DATABRICKS_CLIENT_ID":           "environment-client",
		"DATABRICKS_CLIENT_SECRET":       "environment-client-secret",
		"DATABRICKS_ACCOUNT_ID":          "environment-account",
		"DATABRICKS_CONFIG_FILE":         "/nonexistent/offline-test/databrickscfg",
		"DATABRICKS_CONFIG_PROFILE":      "nonexistent",
		"ARM_USE_MSI":                    "true",
		"AZURE_FEDERATED_TOKEN_FILE":     "/nonexistent/offline-test/token",
		"GOOGLE_APPLICATION_CREDENTIALS": "/nonexistent/offline-test/google.json",
	} {
		t.Setenv(key, value)
	}
	p := offlineTestProvider(t)
	originalConfigure := p.ConfigureContextFunc
	for _, source := range []xpv2.CredentialsSource{
		xpv2.CredentialsSourceNone, xpv2.CredentialsSourceSecret, xpv2.CredentialsSourceFilesystem,
		credentialsSourceUserAssignedManagedIdentity, credentialsSourceSystemAssignedManagedIdentity,
		credentialsSourceOIDCTokenFile, credentialsSourceUpbound, namespacedv1beta1.CredentialsSourceOIDCTokenRequest,
	} {
		t.Run(string(source), func(t *testing.T) {
			ps := offlineTestSetup(t, p, namespacedv1beta1.ProviderConfigSpec{
				Credentials: namespacedv1beta1.ProviderCredentials{
					Source:            source,
					ServiceAccountRef: &namespacedv1beta1.ServiceAccountReference{Name: "missing"},
				},
			})
			want := map[string]any{keyHost: offlineHost, keyAuthType: authTypePAT, keyAuthToken: offlineToken}
			if diff := cmp.Diff(want, map[string]any(ps.Configuration)); diff != "" {
				t.Errorf("configuration (-want +got):\n%s", diff)
			}
			cfg := offlineTestMeta(t, ps).Config
			if cfg.Host != offlineHost || cfg.AccountID != "" || cfg.AzureClientID != "" || cfg.ClientSecret != "" {
				t.Fatalf("environment leaked into offline config: host=%q account=%q", cfg.Host, cfg.AccountID)
			}
			for i := 0; i < 2; i++ {
				req, err := http.NewRequest(http.MethodGet, offlineHost, nil)
				if err != nil {
					t.Fatal(err)
				}
				if err := cfg.Authenticate(req); err != nil {
					t.Fatal(err)
				}
				if got := req.Header.Get("Authorization"); got != "Bearer "+offlineToken {
					t.Errorf("unexpected authorization header %q", got)
				}
			}
		})
	}
	if p.Meta() != nil || p.ConfigureContextFunc == nil || originalConfigure == nil {
		t.Error("offline setup mutated the shared provider")
	}
	// A separately constructed SDK config must retain normal PAT authentication.
	normal := &dbconfig.Config{Host: testHost, Token: "normal-token",
		Loaders:              []dbconfig.Loader{offlineConfigLoader{}},
		HostMetadataResolver: func(context.Context, string) (*dbconfig.HostMetadata, error) { return nil, nil },
	}
	req, _ := http.NewRequest(http.MethodGet, testHost, nil)
	if err := normal.Authenticate(req); err != nil {
		t.Fatal(err)
	}
	if req.Header.Get("Authorization") != "Bearer normal-token" {
		t.Error("offline setup changed authentication for other clients")
	}
}

func TestOfflineClientBlocksRequests(t *testing.T) {
	forbidOfflineEgress(t)
	p := offlineTestProvider(t)
	for _, host := range []string{testHost, "https://workspace.cloud.databricks.com", "https://workspace.gcp.databricks.com", "https://accounts.azuredatabricks.net"} {
		t.Run(host, func(t *testing.T) {
			ps := offlineTestSetup(t, p, namespacedv1beta1.ProviderConfigSpec{Host: ref(host), AccountID: ref("offline-account")})
			meta := offlineTestMeta(t, ps)
			var configured fwprovider.ConfigureResponse
			ps.FrameworkProvider.Configure(t.Context(), fwprovider.ConfigureRequest{}, &configured)
			if configured.Diagnostics.HasError() || configured.ResourceData != meta || configured.DataSourceData != meta {
				t.Fatal("Framework resources did not receive the isolated guarded client")
			}
			if meta.Config.Host != host || meta.Config.AccountID != "offline-account" {
				t.Fatal("ProviderConfig host/account ID were not preserved")
			}
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			check := func(err error) {
				t.Helper()
				if !errors.Is(err, diffserver.ErrDiffComputationNotSupported) {
					t.Fatalf("expected unsupported diff error, got %v", err)
				}
			}
			check(meta.Get(ctx, "/api/2.0/test", nil, nil))
			if meta.Config.IsAccountClient() {
				a, err := meta.AccountClient()
				if err != nil {
					t.Fatal(err)
				}
				_, err = a.Groups.ListAll(ctx, iam.ListAccountGroupsRequest{})
				check(err)
			} else {
				w, err := meta.WorkspaceClient()
				if err != nil {
					t.Fatal(err)
				}
				_, err = w.CurrentUser.Me(ctx, iam.MeRequest{})
				check(err)
			}
			derived, err := meta.ClientForHost(ctx, "https://derived.cloud.databricks.com")
			if err != nil {
				t.Fatal(err)
			}
			check(derived.Get(ctx, "/api/2.0/test", nil, nil))
		})
	}
}

func TestOfflineTransportError(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "https://user:password@host.invalid/api/2.0/test?sensitive-query=secret", strings.NewReader("sensitive-body"))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := (offlineTransport{}).RoundTrip(req)
	if resp != nil || !errors.Is(err, diffserver.ErrDiffComputationNotSupported) {
		t.Fatalf("expected no response and an unsupported diff error, got %v, %v", resp, err)
	}
	for _, sensitive := range []string{"password", "host.invalid", "sensitive-query", "secret", "sensitive-body"} {
		if strings.Contains(err.Error(), sensitive) {
			t.Errorf("transport error contains %q", sensitive)
		}
	}
	if !strings.Contains(err.Error(), "POST /api/2.0/test") {
		t.Errorf("transport error omitted method/path: %v", err)
	}
}

func TestOfflineTerraformDiff(t *testing.T) {
	forbidOfflineEgress(t)
	p := offlineTestProvider(t)
	ps := offlineTestSetup(t, p, namespacedv1beta1.ProviderConfigSpec{})
	r := p.ResourcesMap["databricks_cluster_policy"]
	state := &tfsdk.InstanceState{ID: "policy-id", Attributes: map[string]string{
		"id": "policy-id", "name": "old-name", "definition": `{"a":1}`,
	}}
	for _, tc := range []struct {
		name       string
		definition string
		wantChange bool
	}{
		{name: "old-name", definition: " \n" + `{"a":1}` + "\n "},
		{name: "new-name", definition: `{"a":1}`, wantChange: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			diff, err := r.Diff(context.Background(), state, tfsdk.NewResourceConfigRaw(map[string]any{
				"name": tc.name, "definition": tc.definition,
			}), ps.Meta)
			if err != nil {
				t.Fatal(err)
			}
			if tc.wantChange {
				if diff == nil || diff.Attributes["name"] == nil {
					t.Fatal("expected name diff")
				}
				if got := diff.Attributes["name"]; got.Old != "old-name" || got.New != "new-name" || got.RequiresNew {
					t.Fatalf("unexpected name diff: %+v", got)
				}
			} else if diff != nil && diff.Attributes["name"] != nil {
				t.Fatal("unchanged name should not produce a diff")
			}
			if diff != nil && diff.Attributes["definition"] != nil {
				got := diff.Attributes["definition"]
				if got.Old != got.New || got.NewComputed || got.NewRemoved || got.RequiresNew {
					t.Fatalf("surrounding whitespace should not produce a definition diff: %+v", got)
				}
			}
		})
	}
}

func TestOfflineProviderConfigResolution(t *testing.T) {
	forbidOfflineEgress(t)
	s := testScheme(t)
	for _, add := range []func() error{
		func() error { return clusterv1beta1.SchemeBuilder.AddToScheme(s) },
		func() error { return clusterworkspacev1beta1.AddToScheme(s) },
	} {
		if err := add(); err != nil {
			t.Fatal(err)
		}
	}
	p := offlineTestProvider(t)
	t.Run("legacy", func(t *testing.T) {
		pc := &clusterv1beta1.ProviderConfig{ObjectMeta: metav1.ObjectMeta{Name: "default"},
			Spec: clusterv1beta1.ProviderConfigSpec{Host: ref(testHost)}}
		mg := &clusterworkspacev1beta1.Notebook{ObjectMeta: metav1.ObjectMeta{Name: "legacy", UID: "legacy-uid"}}
		mg.SetProviderConfigReference(&xpv2.Reference{Name: "default"})
		kube := fake.NewClientBuilder().WithScheme(s).WithObjects(pc).Build()
		ps, err := offlineTestSetupBuilder(t, p)(context.Background(), kube, mg)
		if err != nil {
			t.Fatal(err)
		}
		if offlineTestMeta(t, ps).Config.Host != testHost {
			t.Fatal("legacy host not preserved")
		}
	})
	for _, kind := range []string{"ProviderConfig", "ClusterProviderConfig"} {
		t.Run(kind, func(t *testing.T) {
			mg := &workspacev1beta1.Notebook{ObjectMeta: metav1.ObjectMeta{Name: "modern", Namespace: "test", UID: "modern-uid"}}
			mg.SetProviderConfigReference(&xpv2.ProviderConfigReference{Kind: kind, Name: "default"})
			var pc client.Object = &namespacedv1beta1.ProviderConfig{ObjectMeta: metav1.ObjectMeta{Name: "default", Namespace: "test"}}
			if kind == "ClusterProviderConfig" {
				pc = &namespacedv1beta1.ClusterProviderConfig{ObjectMeta: metav1.ObjectMeta{Name: "default"}}
			}
			kube := fake.NewClientBuilder().WithScheme(s).WithObjects(pc).Build()
			if _, err := offlineTestSetupBuilder(t, p)(context.Background(), kube, mg); err != nil {
				t.Fatal(err)
			}
		})
	}
	for _, tc := range []struct {
		name string
		ref  *xpv2.ProviderConfigReference
		want string
	}{
		{name: "no reference", want: errNoProviderConfig},
		{name: "missing config", ref: &xpv2.ProviderConfigReference{Kind: "ProviderConfig", Name: "missing"}, want: errGetProviderConfig},
		{name: "invalid kind", ref: &xpv2.ProviderConfigReference{Kind: "Invalid", Name: "default"}, want: "is invalid"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			mg := &workspacev1beta1.Notebook{}
			mg.SetProviderConfigReference(tc.ref)
			kube := fake.NewClientBuilder().WithScheme(s).Build()
			ps, err := offlineTestSetupBuilder(t, p)(context.Background(), kube, mg)
			if err == nil || !strings.Contains(err.Error(), tc.want) || ps.Meta != nil {
				t.Fatalf("expected empty setup and %q error, got %+v, %v", tc.want, ps, err)
			}
		})
	}
}

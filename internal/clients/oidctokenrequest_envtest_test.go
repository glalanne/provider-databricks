//go:build envtest

// Integration tests against a real kube-apiserver (envtest). Run with:
//
//	KUBEBUILDER_ASSETS=$(setup-envtest use 1.36.x -p path) go test -tags envtest ./internal/clients/ -run Envtest
package clients

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
	"time"

	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	clusterv1beta1 "github.com/glalanne/provider-databricks/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/glalanne/provider-databricks/apis/namespaced/v1beta1"
)

func startEnvtest(t *testing.T) (*envtest.Environment, *runtime.Scheme, client.Client) {
	t.Helper()
	env := &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join("..", "..", "package", "crds", "databricks.crossplane.io_providerconfigs.yaml"),
			filepath.Join("..", "..", "package", "crds", "databricks.m.crossplane.io_providerconfigs.yaml"),
			filepath.Join("..", "..", "package", "crds", "databricks.m.crossplane.io_clusterproviderconfigs.yaml"),
		},
		ErrorIfCRDPathMissing: true,
	}
	cfg, err := env.Start()
	if err != nil {
		t.Fatalf("starting envtest: %v", err)
	}
	t.Cleanup(func() { _ = env.Stop() })

	s := runtime.NewScheme()
	for _, add := range []func(*runtime.Scheme) error{clientgoscheme.AddToScheme, namespacedv1beta1.SchemeBuilder.AddToScheme, clusterv1beta1.SchemeBuilder.AddToScheme} {
		if err := add(s); err != nil {
			t.Fatal(err)
		}
	}
	admin, err := client.New(cfg, client.Options{Scheme: s})
	if err != nil {
		t.Fatal(err)
	}
	return env, s, admin
}

func mustCreate(ctx context.Context, t *testing.T, c client.Client, objs ...client.Object) {
	t.Helper()
	for _, o := range objs {
		if err := c.Create(ctx, o); err != nil {
			t.Fatalf("create %T %s: %v", o, o.GetName(), err)
		}
	}
}

// TestEnvtestCELRules proves the CRD validation rules compile and are enforced
// by a real API server, and that existing configs are unaffected.
func TestEnvtestCELRules(t *testing.T) {
	_, _, admin := startEnvtest(t)
	ctx := context.Background()
	mustCreate(ctx, t, admin, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "dp-a"}})

	nsPC := func(name string, mutate func(*namespacedv1beta1.ProviderConfigSpec)) *namespacedv1beta1.ProviderConfig {
		spec := namespacedv1beta1.ProviderConfigSpec{
			Credentials: namespacedv1beta1.ProviderCredentials{
				Source:            namespacedv1beta1.CredentialsSourceOIDCTokenRequest,
				ServiceAccountRef: &namespacedv1beta1.ServiceAccountReference{Name: "dp-a"},
			},
			Host: ref(testHost),
		}
		if mutate != nil {
			mutate(&spec)
		}
		return &namespacedv1beta1.ProviderConfig{ObjectMeta: metav1.ObjectMeta{Namespace: "dp-a", Name: name}, Spec: spec}
	}
	legacyPC := func(name string, mutate func(*clusterv1beta1.ProviderConfigSpec)) *clusterv1beta1.ProviderConfig {
		spec := clusterv1beta1.ProviderConfigSpec{
			Credentials: clusterv1beta1.ProviderCredentials{
				Source:            clusterv1beta1.CredentialsSourceOIDCTokenRequest,
				ServiceAccountRef: &clusterv1beta1.ServiceAccountReference{Name: "dp-a", Namespace: ref("dp-a")},
			},
			Host: ref(testHost),
		}
		if mutate != nil {
			mutate(&spec)
		}
		return &clusterv1beta1.ProviderConfig{ObjectMeta: metav1.ObjectMeta{Name: name}, Spec: spec}
	}
	withAuthority := func(host namespacedv1beta1.AzureAuthorityHost) func(*namespacedv1beta1.ProviderConfigSpec) {
		return func(s *namespacedv1beta1.ProviderConfigSpec) {
			s.Credentials.TokenExchange = &namespacedv1beta1.TokenExchange{AzureAD: &namespacedv1beta1.AzureADTokenExchange{AuthorityHost: &host}}
		}
	}
	withHost := func(host string) func(*namespacedv1beta1.ProviderConfigSpec) {
		return func(s *namespacedv1beta1.ProviderConfigSpec) { s.Host = &host }
	}
	withAzureAD := func(clientID, tenantID string) func(*namespacedv1beta1.ProviderConfigSpec) {
		return func(s *namespacedv1beta1.ProviderConfigSpec) {
			s.Credentials.TokenExchange = &namespacedv1beta1.TokenExchange{AzureAD: &namespacedv1beta1.AzureADTokenExchange{ClientID: &clientID, TenantID: &tenantID}}
		}
	}
	withSA := func(name string) func(*namespacedv1beta1.ProviderConfigSpec) {
		return func(s *namespacedv1beta1.ProviderConfigSpec) { s.Credentials.ServiceAccountRef.Name = name }
	}

	cases := map[string]struct {
		obj     client.Object
		wantErr string // empty = must be accepted
	}{
		"namespaced: valid OIDCTokenRequest": {obj: nsPC("ok", nil)},
		"namespaced: valid with tokenExchange and audiences": {obj: nsPC("ok-full", func(s *namespacedv1beta1.ProviderConfigSpec) {
			typ := namespacedv1beta1.TokenExchangeAzureAD
			s.Credentials.TokenExchange = &namespacedv1beta1.TokenExchange{Type: &typ, AzureAD: &namespacedv1beta1.AzureADTokenExchange{
				ClientID:      ref(testClientID),
				TenantID:      ref(testTenantID),
				AuthorityHost: ref(namespacedv1beta1.AzureAuthorityHostPublic),
			}}
			s.Credentials.ServiceAccountRef.Audiences = []string{"databricks://dp-a/ok-full"}
		})},
		"namespaced: US Government authority host": {obj: nsPC("authority-us", withAuthority(namespacedv1beta1.AzureAuthorityHostUSGovernment))},
		"namespaced: China authority host":         {obj: nsPC("authority-cn", withAuthority(namespacedv1beta1.AzureAuthorityHostChina))},
		"namespaced: China partner authority host": {obj: nsPC("authority-cn-partner", withAuthority(namespacedv1beta1.AzureAuthorityHostChinaPartner))},
		"namespaced: authority host over plain HTTP": {
			obj: nsPC("authority-http", withAuthority("http://login.microsoftonline.com/")), wantErr: "authorityHost: Unsupported value",
		},
		"namespaced: authority host that isn't an Entra authority": {
			obj: nsPC("authority-other", withAuthority("https://login.example/")), wantErr: "authorityHost: Unsupported value",
		},
		"namespaced: authority host without the trailing slash": {
			obj: nsPC("authority-no-slash", withAuthority("https://login.microsoftonline.com")), wantErr: "authorityHost: Unsupported value",
		},
		"namespaced: US Government workspace host": {obj: nsPC("host-us", withHost("https://adb-1234567890123456.7.databricks.azure.us"))},
		"namespaced: China workspace host":         {obj: nsPC("host-cn", withHost("https://adb-1234567890123456.7.databricks.azure.cn"))},
		"namespaced: host with a trailing slash":   {obj: nsPC("host-slash", withHost(testHost+"/"))},
		"namespaced: host over plain HTTP": {
			obj: nsPC("host-http", withHost("http://adb-1234567890123456.7.azuredatabricks.net")), wantErr: "host must be an https URL on an Azure Databricks domain",
		},
		"namespaced: host that isn't Azure Databricks": {
			obj: nsPC("host-other", withHost("https://databricks.example")), wantErr: "host must be an https URL on an Azure Databricks domain",
		},
		"namespaced: Databricks on AWS host": {
			obj: nsPC("host-aws", withHost("https://dbc-12345678-9abc.cloud.databricks.com")), wantErr: "host must be an https URL on an Azure Databricks domain",
		},
		"namespaced: host without a scheme": {
			obj: nsPC("host-no-scheme", withHost("adb-1234567890123456.7.azuredatabricks.net")), wantErr: "host must be an https URL on an Azure Databricks domain",
		},
		"namespaced: host with a port": {
			obj: nsPC("host-port", withHost(testHost+":443")), wantErr: "host must be an https URL on an Azure Databricks domain",
		},
		"namespaced: host with a path": {
			obj: nsPC("host-path", withHost(testHost+"/api")), wantErr: "host must be an https URL on an Azure Databricks domain",
		},
		"namespaced: host with an Azure Databricks domain inside a longer hostname": {
			obj: nsPC("host-longer", withHost(testHost+".example")), wantErr: "host must be an https URL on an Azure Databricks domain",
		},
		"namespaced: host that isn't Azure Databricks, with tokenExchange.type AzureAD": {obj: nsPC("host-other-typed", func(s *namespacedv1beta1.ProviderConfigSpec) {
			typ := namespacedv1beta1.TokenExchangeAzureAD
			s.Credentials.TokenExchange = &namespacedv1beta1.TokenExchange{Type: &typ}
			s.Host = ref("https://databricks.example")
		}), wantErr: "host must be an https URL on an Azure Databricks domain"},
		"namespaced: host that isn't Azure Databricks, with AzureAD settings and no type": {obj: nsPC("host-other-untyped", func(s *namespacedv1beta1.ProviderConfigSpec) {
			withAzureAD(testClientID, testTenantID)(s)
			s.Host = ref("https://databricks.example")
		}), wantErr: "host must be an https URL on an Azure Databricks domain"},
		"namespaced: Secret config's host isn't checked": {obj: &namespacedv1beta1.ProviderConfig{
			ObjectMeta: metav1.ObjectMeta{Namespace: "dp-a", Name: "secret-host"},
			Spec: namespacedv1beta1.ProviderConfigSpec{
				Credentials: namespacedv1beta1.ProviderCredentials{
					Source: xpv2.CredentialsSourceSecret,
					CommonCredentialSelectors: xpv2.CommonCredentialSelectors{SecretRef: &xpv2.SecretKeySelector{
						SecretReference: xpv2.SecretReference{Namespace: "dp-a", Name: "creds"}, Key: "credential",
					}},
				},
				Host: ref("https://databricks.example"),
			},
		}},
		"namespaced: client and tenant IDs in upper case": {
			obj: nsPC("ids-upper", withAzureAD("AAAAAAAA-BBBB-CCCC-DDDD-EEEEEEEEEEEE", "FFFFFFFF-AAAA-BBBB-CCCC-DDDDDDDDDDDD")),
		},
		"namespaced: client ID that isn't a GUID": {
			obj: nsPC("client-not-guid", withAzureAD("databricks-sp", testTenantID)), wantErr: "spec.credentials.tokenExchange.azureAD.clientID in body should match",
		},
		"namespaced: tenant ID that isn't a GUID": {
			obj: nsPC("tenant-not-guid", withAzureAD(testClientID, "contoso.onmicrosoft.com")), wantErr: "spec.credentials.tokenExchange.azureAD.tenantID in body should match",
		},
		"namespaced: account ID that isn't a GUID": {obj: nsPC("account-not-guid", func(s *namespacedv1beta1.ProviderConfigSpec) {
			s.Host = ref("https://accounts.azuredatabricks.net")
			s.AccountID = ref("my-account")
		}), wantErr: "spec.accountID in body should match"},
		"namespaced: service account name with a dot": {obj: nsPC("sa-dotted", withSA("databricks.sa"))},
		"namespaced: service account name that isn't a DNS subdomain": {
			obj: nsPC("sa-not-subdomain", withSA("Databricks_SA")), wantErr: "spec.credentials.serviceAccountRef.name in body should match",
		},
		"namespaced: service account name that's too long": {
			obj: nsPC("sa-too-long", withSA(strings.Repeat("a", 254))), wantErr: "spec.credentials.serviceAccountRef.name: Too long",
		},
		"namespaced: valid account-level": {obj: nsPC("ok-account", func(s *namespacedv1beta1.ProviderConfigSpec) {
			s.Host = ref("https://accounts.azuredatabricks.net")
			s.AccountID = ref("00000000-0000-0000-0000-000000000001")
		})},
		"namespaced: missing host":               {obj: nsPC("no-host", func(s *namespacedv1beta1.ProviderConfigSpec) { s.Host = nil }), wantErr: "host is required"},
		"namespaced: missing serviceAccountRef":  {obj: nsPC("no-ref", func(s *namespacedv1beta1.ProviderConfigSpec) { s.Credentials.ServiceAccountRef = nil }), wantErr: "serviceAccountRef is required"},
		"namespaced: empty service account name": {obj: nsPC("empty-name", func(s *namespacedv1beta1.ProviderConfigSpec) { s.Credentials.ServiceAccountRef.Name = "" }), wantErr: "spec.credentials.serviceAccountRef.name"},
		"namespaced: unknown exchange type": {obj: nsPC("bad-type", func(s *namespacedv1beta1.ProviderConfigSpec) {
			typ := namespacedv1beta1.TokenExchangeType("Other")
			s.Credentials.TokenExchange = &namespacedv1beta1.TokenExchange{Type: &typ}
		}), wantErr: "spec.credentials.tokenExchange.type"},
		"namespaced: duplicate audiences": {obj: nsPC("dup-aud", func(s *namespacedv1beta1.ProviderConfigSpec) {
			s.Credentials.ServiceAccountRef.Audiences = []string{"x", "x"}
		}), wantErr: "Duplicate value"},
		"namespaced: existing Secret config unaffected": {obj: &namespacedv1beta1.ProviderConfig{
			ObjectMeta: metav1.ObjectMeta{Namespace: "dp-a", Name: "secret"},
			Spec: namespacedv1beta1.ProviderConfigSpec{Credentials: namespacedv1beta1.ProviderCredentials{
				Source: xpv2.CredentialsSourceSecret,
				CommonCredentialSelectors: xpv2.CommonCredentialSelectors{SecretRef: &xpv2.SecretKeySelector{
					SecretReference: xpv2.SecretReference{Namespace: "dp-a", Name: "creds"}, Key: "credential",
				}},
			}},
		}},
		"cluster-scoped: valid ClusterProviderConfig": {obj: &namespacedv1beta1.ClusterProviderConfig{
			ObjectMeta: metav1.ObjectMeta{Name: "cluster-ok"},
			Spec:       nsPC("unused", func(s *namespacedv1beta1.ProviderConfigSpec) { s.Credentials.ServiceAccountRef.Namespace = ref("dp-a") }).Spec,
		}},
		"cluster-scoped: missing namespace": {obj: &namespacedv1beta1.ClusterProviderConfig{
			ObjectMeta: metav1.ObjectMeta{Name: "cluster-no-ns"},
			Spec:       nsPC("unused", nil).Spec,
		}, wantErr: "serviceAccountRef with a namespace is required"},
		"cluster-scoped: missing serviceAccountRef": {obj: &namespacedv1beta1.ClusterProviderConfig{
			ObjectMeta: metav1.ObjectMeta{Name: "cluster-no-ref"},
			Spec:       nsPC("unused", func(s *namespacedv1beta1.ProviderConfigSpec) { s.Credentials.ServiceAccountRef = nil }).Spec,
		}, wantErr: "serviceAccountRef with a namespace is required"},
		"cluster-scoped: authority host that isn't an Entra authority": {obj: &namespacedv1beta1.ClusterProviderConfig{
			ObjectMeta: metav1.ObjectMeta{Name: "cluster-authority-other"},
			Spec: nsPC("unused", func(s *namespacedv1beta1.ProviderConfigSpec) {
				s.Credentials.ServiceAccountRef.Namespace = ref("dp-a")
				withAuthority("https://login.example/")(s)
			}).Spec,
		}, wantErr: "authorityHost: Unsupported value"},
		"cluster-scoped: host that isn't Azure Databricks": {obj: &namespacedv1beta1.ClusterProviderConfig{
			ObjectMeta: metav1.ObjectMeta{Name: "cluster-host-other"},
			Spec: nsPC("unused", func(s *namespacedv1beta1.ProviderConfigSpec) {
				s.Credentials.ServiceAccountRef.Namespace = ref("dp-a")
				withHost("https://databricks.example")(s)
			}).Spec,
		}, wantErr: "host must be an https URL on an Azure Databricks domain"},
		"cluster-scoped: service account namespace that isn't a DNS label": {obj: &namespacedv1beta1.ClusterProviderConfig{
			ObjectMeta: metav1.ObjectMeta{Name: "cluster-ns-not-label"},
			Spec: nsPC("unused", func(s *namespacedv1beta1.ProviderConfigSpec) {
				s.Credentials.ServiceAccountRef.Namespace = ref("team.a")
			}).Spec,
		}, wantErr: "spec.credentials.serviceAccountRef.namespace in body should match"},
		"cluster-scoped: service account namespace that's too long": {obj: &namespacedv1beta1.ClusterProviderConfig{
			ObjectMeta: metav1.ObjectMeta{Name: "cluster-ns-too-long"},
			Spec: nsPC("unused", func(s *namespacedv1beta1.ProviderConfigSpec) {
				s.Credentials.ServiceAccountRef.Namespace = ref(strings.Repeat("a", 64))
			}).Spec,
		}, wantErr: "spec.credentials.serviceAccountRef.namespace: Too long"},
		"cluster-scoped: existing Secret config unaffected": {obj: &namespacedv1beta1.ClusterProviderConfig{
			ObjectMeta: metav1.ObjectMeta{Name: "cluster-secret"},
			Spec: namespacedv1beta1.ProviderConfigSpec{Credentials: namespacedv1beta1.ProviderCredentials{
				Source: xpv2.CredentialsSourceSecret,
				CommonCredentialSelectors: xpv2.CommonCredentialSelectors{SecretRef: &xpv2.SecretKeySelector{
					SecretReference: xpv2.SecretReference{Namespace: "crossplane-system", Name: "creds"}, Key: "credentials",
				}},
			}},
		}},
		"legacy: valid":             {obj: legacyPC("legacy-ok", nil)},
		"legacy: missing namespace": {obj: legacyPC("legacy-no-ns", func(s *clusterv1beta1.ProviderConfigSpec) { s.Credentials.ServiceAccountRef.Namespace = nil }), wantErr: "serviceAccountRef with a namespace is required"},
		"legacy: missing host":      {obj: legacyPC("legacy-no-host", func(s *clusterv1beta1.ProviderConfigSpec) { s.Host = nil }), wantErr: "host is required"},
		"legacy: authority host that isn't an Entra authority": {obj: legacyPC("legacy-authority-other", func(s *clusterv1beta1.ProviderConfigSpec) {
			s.Credentials.TokenExchange = &clusterv1beta1.TokenExchange{AzureAD: &clusterv1beta1.AzureADTokenExchange{
				AuthorityHost: ref(clusterv1beta1.AzureAuthorityHost("https://login.example/")),
			}}
		}), wantErr: "authorityHost: Unsupported value"},
		"legacy: host that isn't Azure Databricks": {obj: legacyPC("legacy-host-other", func(s *clusterv1beta1.ProviderConfigSpec) {
			s.Host = ref("https://databricks.example")
		}), wantErr: "host must be an https URL on an Azure Databricks domain"},
		"legacy: client ID that isn't a GUID": {obj: legacyPC("legacy-client-not-guid", func(s *clusterv1beta1.ProviderConfigSpec) {
			s.Credentials.TokenExchange = &clusterv1beta1.TokenExchange{AzureAD: &clusterv1beta1.AzureADTokenExchange{ClientID: ref("databricks-sp")}}
		}), wantErr: "spec.credentials.tokenExchange.azureAD.clientID in body should match"},
		"legacy: tenant ID that isn't a GUID": {obj: legacyPC("legacy-tenant-not-guid", func(s *clusterv1beta1.ProviderConfigSpec) {
			s.Credentials.TokenExchange = &clusterv1beta1.TokenExchange{AzureAD: &clusterv1beta1.AzureADTokenExchange{TenantID: ref("contoso.onmicrosoft.com")}}
		}), wantErr: "spec.credentials.tokenExchange.azureAD.tenantID in body should match"},
		"legacy: account ID that isn't a GUID": {obj: legacyPC("legacy-account-not-guid", func(s *clusterv1beta1.ProviderConfigSpec) {
			s.Host = ref("https://accounts.azuredatabricks.net")
			s.AccountID = ref("my-account")
		}), wantErr: "spec.accountID in body should match"},
		"legacy: service account name that isn't a DNS subdomain": {obj: legacyPC("legacy-sa-not-subdomain", func(s *clusterv1beta1.ProviderConfigSpec) {
			s.Credentials.ServiceAccountRef.Name = "Databricks_SA"
		}), wantErr: "spec.credentials.serviceAccountRef.name in body should match"},
		"legacy: service account name that's too long": {obj: legacyPC("legacy-sa-too-long", func(s *clusterv1beta1.ProviderConfigSpec) {
			s.Credentials.ServiceAccountRef.Name = strings.Repeat("a", 254)
		}), wantErr: "spec.credentials.serviceAccountRef.name: Too long"},
		"legacy: service account namespace that isn't a DNS label": {obj: legacyPC("legacy-ns-not-label", func(s *clusterv1beta1.ProviderConfigSpec) {
			s.Credentials.ServiceAccountRef.Namespace = ref("team.a")
		}), wantErr: "spec.credentials.serviceAccountRef.namespace in body should match"},
		"legacy: service account namespace that's too long": {obj: legacyPC("legacy-ns-too-long", func(s *clusterv1beta1.ProviderConfigSpec) {
			s.Credentials.ServiceAccountRef.Namespace = ref(strings.Repeat("a", 64))
		}), wantErr: "spec.credentials.serviceAccountRef.namespace: Too long"},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := admin.Create(ctx, tc.obj)
			switch {
			case tc.wantErr == "" && err != nil:
				t.Fatalf("rejected, want accepted: %v", err)
			case tc.wantErr != "" && err == nil:
				t.Fatalf("accepted, want rejection containing %q", tc.wantErr)
			case tc.wantErr != "" && !strings.Contains(err.Error(), tc.wantErr):
				t.Fatalf("rejected with %v, want it to contain %q", err, tc.wantErr)
			}
		})
	}
}

func jwtClaims(t *testing.T, token string) map[string]any {
	t.Helper()
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		t.Fatalf("token has %d parts, want 3", len(parts))
	}
	raw, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		t.Fatal(err)
	}
	var claims map[string]any
	if err := json.Unmarshal(raw, &claims); err != nil {
		t.Fatal(err)
	}
	return claims
}

// TestEnvtestOptInRBAC proves the permissions model against a real API
// server: a provider identity with only a namespace's opt-in Role can mint
// tokens for that one service account and read it, can't for any other, and
// reads service accounts without cluster-wide list/watch.
func TestEnvtestOptInRBAC(t *testing.T) {
	env, s, admin := startEnvtest(t)
	ctx := context.Background()

	mustCreate(ctx, t, admin,
		&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "dp-a"}},
		&corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "dp-b"}},
		annotatedSA("dp-a", "dp-a"),
		annotatedSA("dp-a", "other"),
		annotatedSA("dp-b", "dp-b"),
		// The namespace's opt-in: the provider may mint tokens for, and read,
		// this one service account. Nothing cluster-wide.
		&rbacv1.Role{
			ObjectMeta: metav1.ObjectMeta{Namespace: "dp-a", Name: "provider-databricks-tokenrequest"},
			Rules: []rbacv1.PolicyRule{
				{APIGroups: []string{""}, Resources: []string{"serviceaccounts/token"}, ResourceNames: []string{"dp-a"}, Verbs: []string{"create"}},
				{APIGroups: []string{""}, Resources: []string{"serviceaccounts"}, ResourceNames: []string{"dp-a"}, Verbs: []string{"get"}},
			},
		},
		&rbacv1.RoleBinding{
			ObjectMeta: metav1.ObjectMeta{Namespace: "dp-a", Name: "provider-databricks-tokenrequest"},
			RoleRef:    rbacv1.RoleRef{APIGroup: "rbac.authorization.k8s.io", Kind: "Role", Name: "provider-databricks-tokenrequest"},
			Subjects:   []rbacv1.Subject{{Kind: "User", Name: "provider"}},
		},
	)

	user, err := env.ControlPlane.AddUser(envtest.User{Name: "provider"}, nil)
	if err != nil {
		t.Fatal(err)
	}
	providerCfg := user.Config()

	// The provider's real client: cached, with ServiceAccounts excluded, as
	// configured in cmd/provider/main.go.
	newClient := func(cfg *rest.Config, disableSA bool) client.Client {
		t.Helper()
		c, err := cache.New(cfg, cache.Options{Scheme: s})
		if err != nil {
			t.Fatal(err)
		}
		cctx, cancel := context.WithCancel(ctx)
		t.Cleanup(cancel)
		go func() { _ = c.Start(cctx) }()
		if !c.WaitForCacheSync(cctx) {
			t.Fatal("cache did not start")
		}
		opts := client.Options{Scheme: s, Cache: &client.CacheOptions{Reader: c}}
		if disableSA {
			opts.Cache.DisableFor = []client.Object{&corev1.ServiceAccount{}}
		}
		cl, err := client.New(cfg, opts)
		if err != nil {
			t.Fatal(err)
		}
		return cl
	}
	newProviderClient := func(disableSA bool) client.Client { return newClient(providerCfg, disableSA) }
	provider := newProviderClient(true)

	t.Run("mints a token for the opted-in service account", func(t *testing.T) {
		src := &ServiceAccountTokenSource{Client: provider, Namespace: "dp-a", Name: "dp-a", Audiences: withAudience(AzureADTokenExchangeAudience, []string{"databricks://dp-a/default"})}
		tok, err := src.Token(ctx)
		if err != nil {
			t.Fatalf("Token() error: %v", err)
		}
		claims := jwtClaims(t, tok)
		if got, want := claims["sub"], "system:serviceaccount:dp-a:dp-a"; got != want {
			t.Errorf("sub = %v, want %s", got, want)
		}
		aud, _ := claims["aud"].([]any)
		if len(aud) != 2 || aud[0] != AzureADTokenExchangeAudience || aud[1] != "databricks://dp-a/default" {
			t.Errorf("aud = %v, want [%s databricks://dp-a/default]", claims["aud"], AzureADTokenExchangeAudience)
		}
		exp, _ := claims["exp"].(float64)
		iat, _ := claims["iat"].(float64)
		if life := time.Duration(exp-iat) * time.Second; life != serviceAccountTokenExpirationSeconds*time.Second {
			t.Errorf("token lifetime = %s, want %ds", life, serviceAccountTokenExpirationSeconds)
		}
	})

	for name, tc := range map[string]struct{ ns, sa string }{
		"another service account in the same namespace":       {"dp-a", "other"},
		"a service account in a namespace that didn't opt in": {"dp-b", "dp-b"},
	} {
		t.Run("refuses "+name, func(t *testing.T) {
			src := &ServiceAccountTokenSource{Client: provider, Namespace: tc.ns, Name: tc.sa, Audiences: []string{AzureADTokenExchangeAudience}}
			_, err := src.Token(ctx)
			if err == nil || !strings.Contains(err.Error(), "forbidden") {
				t.Fatalf("Token() error = %v, want forbidden", err)
			}
		})
	}

	t.Run("reads the service account's annotations without list/watch", func(t *testing.T) {
		getCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		ex, err := azureADExchanger(getCtx, provider, &namespacedv1beta1.ProviderConfigSpec{}, "dp-a", "dp-a")
		if err != nil {
			t.Fatalf("azureADExchanger() error: %v", err)
		}
		if ex.ClientID != testClientID || ex.TenantID != testTenantID {
			t.Errorf("got client %q tenant %q, want %q %q", ex.ClientID, ex.TenantID, testClientID, testTenantID)
		}
	})

	t.Run("control: a cached ServiceAccount read would need list/watch", func(t *testing.T) {
		cached := newProviderClient(false)
		getCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		err := cached.Get(getCtx, client.ObjectKey{Namespace: "dp-a", Name: "dp-a"}, &corev1.ServiceAccount{})
		if err == nil {
			t.Fatal("cached Get succeeded; the cache exclusion would be unnecessary")
		}
		// The cached read starts an informer that must list/watch
		// ServiceAccounts cluster-wide; without that permission it never
		// syncs, so the read times out rather than returning the object.
		if !strings.Contains(err.Error(), "Informer to sync") {
			t.Fatalf("cached Get failed for an unrelated reason: %v", err)
		}
		t.Logf("cached Get without list/watch fails as expected: %v", err)
	})

	t.Run("control: the same cached read works with cluster-wide list/watch", func(t *testing.T) {
		mustCreate(ctx, t, admin,
			&rbacv1.ClusterRole{
				ObjectMeta: metav1.ObjectMeta{Name: "sa-lister"},
				Rules:      []rbacv1.PolicyRule{{APIGroups: []string{""}, Resources: []string{"serviceaccounts"}, Verbs: []string{"get", "list", "watch"}}},
			},
			&rbacv1.ClusterRoleBinding{
				ObjectMeta: metav1.ObjectMeta{Name: "sa-lister"},
				RoleRef:    rbacv1.RoleRef{APIGroup: "rbac.authorization.k8s.io", Kind: "ClusterRole", Name: "sa-lister"},
				Subjects:   []rbacv1.Subject{{Kind: "User", Name: "lister"}},
			},
		)
		lister, err := env.ControlPlane.AddUser(envtest.User{Name: "lister"}, nil)
		if err != nil {
			t.Fatal(err)
		}
		cached := newClient(lister.Config(), false)
		getCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := cached.Get(getCtx, client.ObjectKey{Namespace: "dp-a", Name: "dp-a"}, &corev1.ServiceAccount{}); err != nil {
			t.Fatalf("cached Get with list/watch failed: %v", err)
		}
	})
}

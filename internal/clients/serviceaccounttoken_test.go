package clients

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/google/go-cmp/cmp"
	authenticationv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"

	namespacedv1beta1 "github.com/glalanne/provider-databricks/apis/namespaced/v1beta1"
	workspacev1beta1 "github.com/glalanne/provider-databricks/apis/namespaced/workspace/v1beta1"
)

const (
	testHost     = "https://adb-1234567890123456.7.azuredatabricks.net"
	testClientID = "11111111-1111-1111-1111-111111111111"
	testTenantID = "22222222-2222-2222-2222-222222222222"
)

func testScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	s := runtime.NewScheme()
	for _, add := range []func(*runtime.Scheme) error{clientgoscheme.AddToScheme, namespacedv1beta1.SchemeBuilder.AddToScheme, workspacev1beta1.AddToScheme} {
		if err := add(s); err != nil {
			t.Fatal(err)
		}
	}
	return s
}

func annotatedSA(ns, name string) *corev1.ServiceAccount {
	return &corev1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{
		Namespace: ns, Name: name,
		Annotations: map[string]string{annotationAzureClientID: testClientID, annotationAzureTenantID: testTenantID},
	}}
}

// fakeEntra accepts only the given client assertion for the given client ID,
// like a federated identity credential bound to one service account.
func fakeEntra(t *testing.T, wantClientID, wantAssertion string) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}
		if r.URL.Path != "/"+testTenantID+"/oauth2/v2.0/token" || r.PostForm.Get("client_id") != wantClientID || r.PostForm.Get("client_assertion") != wantAssertion {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":"invalid_client","error_description":"AADSTS700213: No matching federated identity record found"}`))
			return
		}
		_, _ = w.Write([]byte(`{"token_type":"Bearer","expires_in":3599,"access_token":"entra-access-token"}`))
	}))
	t.Cleanup(srv.Close)
	return srv
}

// tokenFor makes the fake API server issue a token that names the service
// account it was requested for, as a real one would (the token's subject).
func tokenFor(kubeObjs ...client.Object) *fake.ClientBuilder {
	return fake.NewClientBuilder().WithObjects(kubeObjs...).WithInterceptorFuncs(interceptor.Funcs{
		SubResourceCreate: func(ctx context.Context, c client.Client, sub string, obj client.Object, subObj client.Object, opts ...client.SubResourceCreateOption) error {
			if err := c.SubResource(sub).Create(ctx, obj, subObj, opts...); err != nil {
				return err
			}
			subObj.(*authenticationv1.TokenRequest).Status.Token = "sa-token:" + obj.GetNamespace() + "/" + obj.GetName()
			return nil
		},
	})
}

func saSpec(ns, name, authority string) *namespacedv1beta1.ProviderConfigSpec {
	return &namespacedv1beta1.ProviderConfigSpec{
		Credentials: namespacedv1beta1.ProviderCredentials{
			Source:            namespacedv1beta1.CredentialsSourceServiceAccount,
			ServiceAccountRef: &namespacedv1beta1.ServiceAccountReference{Name: name, Namespace: ref(ns)},
			TokenExchange: &namespacedv1beta1.TokenExchange{
				AzureAD: &namespacedv1beta1.AzureADTokenExchange{AuthorityHost: ref(authority)},
			},
		},
		Host: ref(testHost),
	}
}

func TestServiceAccountTokenAuth(t *testing.T) {
	t.Run("configures host and the exchanged token from the service account's annotations", func(t *testing.T) {
		entra := fakeEntra(t, testClientID, "sa-token:dp-a/dp-a")
		kube := tokenFor(annotatedSA("dp-a", "dp-a")).Build()

		ps := &terraform.Setup{Configuration: map[string]any{}}
		if err := serviceAccountTokenAuth(context.Background(), kube, saSpec("dp-a", "dp-a", entra.URL), ps, NewTokenCache()); err != nil {
			t.Fatalf("serviceAccountTokenAuth() error: %v", err)
		}
		want := map[string]any{keyHost: testHost, keyAuthType: authTypePAT, keyAuthToken: "entra-access-token"}
		if diff := cmp.Diff(want, map[string]any(ps.Configuration)); diff != "" {
			t.Errorf("configuration (-want +got):\n%s", diff)
		}
	})

	t.Run("configures account_id for an accounts console host", func(t *testing.T) {
		entra := fakeEntra(t, testClientID, "sa-token:dp-a/dp-a")
		kube := tokenFor(annotatedSA("dp-a", "dp-a")).Build()

		spec := saSpec("dp-a", "dp-a", entra.URL)
		spec.Host = ref("https://accounts.azuredatabricks.net")
		spec.AccountID = ref("00000000-0000-0000-0000-000000000001")
		ps := &terraform.Setup{Configuration: map[string]any{}}
		if err := serviceAccountTokenAuth(context.Background(), kube, spec, ps, NewTokenCache()); err != nil {
			t.Fatalf("serviceAccountTokenAuth() error: %v", err)
		}
		want := map[string]any{
			keyHost:      "https://accounts.azuredatabricks.net",
			keyAccountID: "00000000-0000-0000-0000-000000000001",
			keyAuthType:  authTypePAT,
			keyAuthToken: "entra-access-token",
		}
		if diff := cmp.Diff(want, map[string]any(ps.Configuration)); diff != "" {
			t.Errorf("configuration (-want +got):\n%s", diff)
		}
	})

	t.Run("requests the exchanger's audience first, then the extras, without duplicates", func(t *testing.T) {
		entra := fakeEntra(t, testClientID, "fake-token")
		var got []string
		kube := fake.NewClientBuilder().WithObjects(annotatedSA("dp-a", "dp-a")).WithInterceptorFuncs(interceptor.Funcs{
			SubResourceCreate: func(ctx context.Context, c client.Client, sub string, obj client.Object, subObj client.Object, opts ...client.SubResourceCreateOption) error {
				got = subObj.(*authenticationv1.TokenRequest).Spec.Audiences
				return c.SubResource(sub).Create(ctx, obj, subObj, opts...)
			},
		}).Build()
		spec := saSpec("dp-a", "dp-a", entra.URL)
		spec.Credentials.ServiceAccountRef.Audiences = []string{AzureADTokenExchangeAudience, "databricks://dp-a/default"}
		ps := &terraform.Setup{Configuration: map[string]any{}}
		if err := serviceAccountTokenAuth(context.Background(), kube, spec, ps, NewTokenCache()); err != nil {
			t.Fatalf("serviceAccountTokenAuth() error: %v", err)
		}
		if diff := cmp.Diff([]string{AzureADTokenExchangeAudience, "databricks://dp-a/default"}, got); diff != "" {
			t.Errorf("audiences (-want +got):\n%s", diff)
		}
	})

	t.Run("clientID and tenantID in tokenExchange.azureAD override the annotations", func(t *testing.T) {
		const specClientID = "33333333-3333-3333-3333-333333333333"
		entra := fakeEntra(t, specClientID, "sa-token:dp-a/dp-a")
		bare := &corev1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Namespace: "dp-a", Name: "dp-a"}}
		kube := tokenFor(bare).Build()

		spec := saSpec("dp-a", "dp-a", entra.URL)
		azureAD := spec.Credentials.TokenExchange.AzureAD
		azureAD.ClientID, azureAD.TenantID = ref(specClientID), ref(testTenantID)
		ps := &terraform.Setup{Configuration: map[string]any{}}
		if err := serviceAccountTokenAuth(context.Background(), kube, spec, ps, NewTokenCache()); err != nil {
			t.Fatalf("serviceAccountTokenAuth() error: %v", err)
		}
	})

	// The top-level clientID and tenantID belong to the Azure-derived sources;
	// this source reads its IDs only from tokenExchange.azureAD.
	t.Run("top-level clientID and tenantID are ignored", func(t *testing.T) {
		entra := fakeEntra(t, testClientID, "sa-token:dp-a/dp-a")
		kube := tokenFor(annotatedSA("dp-a", "dp-a")).Build()

		spec := saSpec("dp-a", "dp-a", entra.URL)
		spec.ClientID, spec.TenantID = ref("33333333-3333-3333-3333-333333333333"), ref("44444444-4444-4444-4444-444444444444")
		ps := &terraform.Setup{Configuration: map[string]any{}}
		if err := serviceAccountTokenAuth(context.Background(), kube, spec, ps, NewTokenCache()); err != nil {
			t.Fatalf("serviceAccountTokenAuth() error: %v", err)
		}
	})

	// The confused-deputy case at the identity provider: knowing another data
	// product's client ID isn't enough, because the token names the wrong
	// service account and that identity only trusts its own.
	t.Run("another data product's client ID is refused for this namespace's token", func(t *testing.T) {
		entra := fakeEntra(t, testClientID, "sa-token:dp-a/dp-a") // dp-a's identity trusts only dp-a's SA
		kube := tokenFor(annotatedSA("dp-b", "dp-b")).Build()

		spec := saSpec("dp-b", "dp-b", entra.URL)
		azureAD := spec.Credentials.TokenExchange.AzureAD
		azureAD.ClientID, azureAD.TenantID = ref(testClientID), ref(testTenantID) // dp-b names dp-a's identity
		ps := &terraform.Setup{Configuration: map[string]any{}}
		err := serviceAccountTokenAuth(context.Background(), kube, spec, ps, NewTokenCache())
		if err == nil || !strings.Contains(err.Error(), "AADSTS700213") || !strings.Contains(err.Error(), "dp-b/dp-b") {
			t.Fatalf("serviceAccountTokenAuth() error = %v, want an Entra refusal naming dp-b/dp-b", err)
		}
		if len(ps.Configuration) != 0 {
			t.Errorf("configuration = %v, want nothing set on failure", ps.Configuration)
		}
	})

	errorCases := map[string]struct {
		spec func(authority string) *namespacedv1beta1.ProviderConfigSpec
		objs []client.Object
		want string
	}{
		"no serviceAccountRef": {
			spec: func(a string) *namespacedv1beta1.ProviderConfigSpec {
				s := saSpec("dp-a", "dp-a", a)
				s.Credentials.ServiceAccountRef = nil
				return s
			},
			want: errNoServiceAccountRef,
		},
		"no namespace (cluster-scoped config)": {
			spec: func(a string) *namespacedv1beta1.ProviderConfigSpec {
				s := saSpec("dp-a", "dp-a", a)
				s.Credentials.ServiceAccountRef.Namespace = nil
				return s
			},
			want: errNoServiceAccountNamespace,
		},
		"no host": {
			spec: func(a string) *namespacedv1beta1.ProviderConfigSpec {
				s := saSpec("dp-a", "dp-a", a)
				s.Host = nil
				return s
			},
			want: errNoHost,
		},
		"unsupported exchange type": {
			spec: func(a string) *namespacedv1beta1.ProviderConfigSpec {
				s := saSpec("dp-a", "dp-a", a)
				other := namespacedv1beta1.TokenExchangeType("Other")
				s.Credentials.TokenExchange.Type = &other
				return s
			},
			want: `unsupported credentials.tokenExchange.type "Other"`,
		},
		"service account without client ID annotation": {
			spec: func(a string) *namespacedv1beta1.ProviderConfigSpec { return saSpec("dp-a", "dp-a", a) },
			objs: []client.Object{&corev1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Namespace: "dp-a", Name: "dp-a"}}},
			want: "no client ID",
		},
		"service account missing": {
			spec: func(a string) *namespacedv1beta1.ProviderConfigSpec { return saSpec("dp-a", "dp-a", a) },
			want: "cannot get service account dp-a/dp-a",
		},
	}
	for name, tc := range errorCases {
		t.Run(name, func(t *testing.T) {
			entra := fakeEntra(t, testClientID, "sa-token:dp-a/dp-a")
			kube := tokenFor(tc.objs...).Build()
			ps := &terraform.Setup{Configuration: map[string]any{}}
			err := serviceAccountTokenAuth(context.Background(), kube, tc.spec(entra.URL), ps, NewTokenCache())
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("serviceAccountTokenAuth() error = %v, want it to contain %q", err, tc.want)
			}
		})
	}
}

func TestEnrichLocalServiceAccountRef(t *testing.T) {
	mg := &workspacev1beta1.Directory{ObjectMeta: metav1.ObjectMeta{Namespace: "dp-a", Name: "dir"}}
	pc := func(ns *string) *namespacedv1beta1.ProviderConfig {
		return &namespacedv1beta1.ProviderConfig{Spec: namespacedv1beta1.ProviderConfigSpec{Credentials: namespacedv1beta1.ProviderCredentials{
			Source:            namespacedv1beta1.CredentialsSourceServiceAccount,
			ServiceAccountRef: &namespacedv1beta1.ServiceAccountReference{Name: "dp-a", Namespace: ns},
		}}}
	}

	t.Run("defaults to the ProviderConfig's own namespace", func(t *testing.T) {
		p := pc(nil)
		if err := enrichLocalServiceAccountRef(p, mg); err != nil {
			t.Fatal(err)
		}
		if got := deref(p.Spec.Credentials.ServiceAccountRef.Namespace); got != "dp-a" {
			t.Errorf("namespace = %q, want dp-a", got)
		}
	})
	t.Run("accepts its own namespace spelled out", func(t *testing.T) {
		if err := enrichLocalServiceAccountRef(pc(ref("dp-a")), mg); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("rejects another namespace", func(t *testing.T) {
		err := enrichLocalServiceAccountRef(pc(ref("dp-b")), mg)
		if err == nil || !strings.Contains(err.Error(), `"dp-b"`) {
			t.Fatalf("error = %v, want a rejection naming dp-b", err)
		}
	})
	t.Run("leaves configs without serviceAccountRef alone", func(t *testing.T) {
		p := &namespacedv1beta1.ProviderConfig{}
		if err := enrichLocalServiceAccountRef(p, mg); err != nil || p.Spec.Credentials.ServiceAccountRef != nil {
			t.Fatalf("got %v, %v; want no change", p.Spec.Credentials.ServiceAccountRef, err)
		}
	})
}

// The confused-deputy case at the provider: a namespaced ProviderConfig can't
// point at a service account in another namespace, so it can never obtain
// another namespace's token in the first place.
func TestResolveProviderConfigRejectsOtherNamespaceServiceAccount(t *testing.T) {
	pc := &namespacedv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Namespace: "dp-b", Name: "default"},
		Spec: namespacedv1beta1.ProviderConfigSpec{
			Credentials: namespacedv1beta1.ProviderCredentials{
				Source:            namespacedv1beta1.CredentialsSourceServiceAccount,
				ServiceAccountRef: &namespacedv1beta1.ServiceAccountReference{Name: "dp-a", Namespace: ref("dp-a")},
			},
			Host: ref(testHost),
		},
	}
	mg := &workspacev1beta1.Directory{ObjectMeta: metav1.ObjectMeta{Namespace: "dp-b", Name: "dir"}}
	mg.SetProviderConfigReference(&xpv2.ProviderConfigReference{Kind: "ProviderConfig", Name: "default"})

	kube := fake.NewClientBuilder().WithScheme(testScheme(t)).WithObjects(pc).Build()
	_, err := resolveProviderConfig(context.Background(), kube, mg)
	t.Logf("resolveProviderConfig() error: %v", err)
	if err == nil || !strings.Contains(err.Error(), "can only use service accounts in its own namespace") {
		t.Fatalf("resolveProviderConfig() error = %v, want a rejection naming dp-a", err)
	}
}

// The existing Secret path must be untouched by this change: the same secret
// produces exactly the same Terraform configuration.
func TestDefaultAuthSecretUnchanged(t *testing.T) {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Namespace: "dp-a", Name: "creds"},
		Data: map[string][]byte{"credential": []byte(`{
			"auth_type": "azure-client-secret",
			"azure_client_id": "app-id",
			"azure_client_secret": "app-secret",
			"azure_tenant_id": "tenant-id",
			"host": "` + testHost + `"
		}`)},
	}
	spec := &namespacedv1beta1.ProviderConfigSpec{Credentials: namespacedv1beta1.ProviderCredentials{
		Source: xpv2.CredentialsSourceSecret,
		CommonCredentialSelectors: xpv2.CommonCredentialSelectors{SecretRef: &xpv2.SecretKeySelector{
			SecretReference: xpv2.SecretReference{Namespace: "dp-a", Name: "creds"}, Key: "credential",
		}},
	}}
	kube := fake.NewClientBuilder().WithObjects(secret).Build()
	ps := &terraform.Setup{Configuration: map[string]any{}}
	if err := defaultAuth(context.Background(), spec, ps, kube); err != nil {
		t.Fatalf("defaultAuth() error: %v", err)
	}
	want := map[string]any{
		keyAuthType:          "azure-client-secret",
		keyAzureClientID:     "app-id",
		keyAzureClientSecret: "app-secret",
		keyAzureTenantID:     "tenant-id",
		keyHost:              testHost,
	}
	if diff := cmp.Diff(want, map[string]any(ps.Configuration)); diff != "" {
		t.Errorf("configuration (-want +got):\n%s", diff)
	}
}

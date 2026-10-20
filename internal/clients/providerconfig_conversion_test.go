package clients

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	clusterv1beta1 "github.com/glalanne/provider-databricks/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/glalanne/provider-databricks/apis/namespaced/v1beta1"
)

// The legacy cluster-scoped ProviderConfig is converted to the namespaced spec
// by a JSON round-trip, which silently drops any field the two trees don't
// share. Every field added for service account token auth must survive it.
func TestLegacyToModernProviderConfigSpecCarriesTokenAuthFields(t *testing.T) {
	azureAD := clusterv1beta1.TokenExchangeAzureAD
	legacy := &clusterv1beta1.ProviderConfig{Spec: clusterv1beta1.ProviderConfigSpec{
		Credentials: clusterv1beta1.ProviderCredentials{
			Source: clusterv1beta1.CredentialsSourceOIDCTokenRequest,
			ServiceAccountRef: &clusterv1beta1.ServiceAccountReference{
				Name:      "dp-canary",
				Namespace: ref("dp-ns"),
				Audiences: []string{"extra"},
			},
			TokenExchange: &clusterv1beta1.TokenExchange{
				Type: &azureAD,
				AzureAD: &clusterv1beta1.AzureADTokenExchange{
					ClientID:      ref("11111111-1111-1111-1111-111111111111"),
					TenantID:      ref("22222222-2222-2222-2222-222222222222"),
					AuthorityHost: ref("https://login.example/"),
				},
			},
		},
		Host:      ref("https://adb-1.0.azuredatabricks.net"),
		AccountID: ref("00000000-0000-0000-0000-000000000001"),
	}}

	got, err := legacyToModernProviderConfigSpec(legacy)
	if err != nil {
		t.Fatalf("legacyToModernProviderConfigSpec() error: %v", err)
	}

	modernAzureAD := namespacedv1beta1.TokenExchangeAzureAD
	want := &namespacedv1beta1.ProviderConfigSpec{
		Credentials: namespacedv1beta1.ProviderCredentials{
			Source: namespacedv1beta1.CredentialsSourceOIDCTokenRequest,
			ServiceAccountRef: &namespacedv1beta1.ServiceAccountReference{
				Name:      "dp-canary",
				Namespace: ref("dp-ns"),
				Audiences: []string{"extra"},
			},
			TokenExchange: &namespacedv1beta1.TokenExchange{
				Type: &modernAzureAD,
				AzureAD: &namespacedv1beta1.AzureADTokenExchange{
					ClientID:      ref("11111111-1111-1111-1111-111111111111"),
					TenantID:      ref("22222222-2222-2222-2222-222222222222"),
					AuthorityHost: ref("https://login.example/"),
				},
			},
		},
		Host:      ref("https://adb-1.0.azuredatabricks.net"),
		AccountID: ref("00000000-0000-0000-0000-000000000001"),
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("legacyToModernProviderConfigSpec() (-want +got):\n%s", diff)
	}
}

func ref[T any](v T) *T { return &v }

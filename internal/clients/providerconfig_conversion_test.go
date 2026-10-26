package clients

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/util/yaml"

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
					AuthorityHost: ref(clusterv1beta1.AzureAuthorityHostUSGovernment),
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
					AuthorityHost: ref(namespacedv1beta1.AzureAuthorityHostUSGovernment),
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

// TestLegacyAndModernTokenExchangeSchemasMatch keeps the two API trees in
// step. The round-trip above drops what one tree has and the other lacks, so
// a tokenExchange type, settings member or rule added to only one of them
// would be lost without an error.
func TestLegacyAndModernTokenExchangeSchemasMatch(t *testing.T) {
	type specSchema struct {
		Properties struct {
			Credentials struct {
				Properties struct {
					TokenExchange map[string]any `json:"tokenExchange"`
				} `json:"properties"`
			} `json:"credentials"`
		} `json:"properties"`
		Rules []map[string]any `json:"x-kubernetes-validations"`
	}
	load := func(file string) specSchema {
		t.Helper()
		raw, err := os.ReadFile(filepath.Join("..", "..", "package", "crds", file))
		if err != nil {
			t.Fatal(err)
		}
		var crd struct {
			Spec struct {
				Versions []struct {
					Schema struct {
						OpenAPIV3Schema struct {
							Properties struct {
								Spec specSchema `json:"spec"`
							} `json:"properties"`
						} `json:"openAPIV3Schema"`
					} `json:"schema"`
				} `json:"versions"`
			} `json:"spec"`
		}
		if err := yaml.Unmarshal(raw, &crd); err != nil {
			t.Fatal(err)
		}
		if len(crd.Spec.Versions) != 1 {
			t.Fatalf("%s has %d versions, want 1", file, len(crd.Spec.Versions))
		}
		s := crd.Spec.Versions[0].Schema.OpenAPIV3Schema.Properties.Spec
		if s.Properties.Credentials.Properties.TokenExchange == nil || len(s.Rules) == 0 {
			t.Fatalf("%s has no credentials.tokenExchange schema or spec rules", file)
		}
		return s
	}

	modern := load("databricks.m.crossplane.io_providerconfigs.yaml")
	legacy := load("databricks.crossplane.io_providerconfigs.yaml")
	if diff := cmp.Diff(modern.Properties.Credentials.Properties.TokenExchange, legacy.Properties.Credentials.Properties.TokenExchange); diff != "" {
		t.Errorf("credentials.tokenExchange differs between the namespaced and legacy CRDs (-namespaced +legacy):\n%s", diff)
	}
	if diff := cmp.Diff(modern.Rules, legacy.Rules); diff != "" {
		t.Errorf("spec rules differ between the namespaced and legacy CRDs (-namespaced +legacy):\n%s", diff)
	}
}

func ref[T any](v T) *T { return &v }

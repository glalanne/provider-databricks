package clients

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/terraform"
	namespacedv1beta1 "github.com/glalanne/provider-databricks/apis/namespaced/v1beta1"
	"github.com/google/go-cmp/cmp"
)

func Test_oidcAuth_tokenFilePath(t *testing.T) {
	subID, tenantID, clientID := "sub", "tenant", "client"
	explicitPath := "/explicit/path/azure-identity-token"
	envPath := "/var/run/secrets/azure/wi/token/azure-identity-token"

	cases := map[string]struct {
		oidcTokenFilePath *string
		envValue          string
		envSet            bool
		want              string
	}{
		"explicit_path_wins_over_env": {
			oidcTokenFilePath: &explicitPath,
			envValue:          envPath,
			envSet:            true,
			want:              explicitPath,
		},
		"env_used_when_no_explicit_path": {
			envValue: envPath,
			envSet:   true,
			want:     envPath,
		},
		"falls_back_to_default_when_env_unset": {
			want: defaultOidcTokenFilePath,
		},
		"falls_back_to_default_when_env_empty": {
			envValue: "",
			envSet:   true,
			want:     defaultOidcTokenFilePath,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if tc.envSet {
				t.Setenv(envAzureFederatedTokenFile, tc.envValue)
			}
			pcSpec := &namespacedv1beta1.ProviderConfigSpec{
				SubscriptionID:    &subID,
				TenantID:          &tenantID,
				ClientID:          &clientID,
				OidcTokenFilePath: tc.oidcTokenFilePath,
			}
			ps := &terraform.Setup{Configuration: terraform.ProviderConfiguration{}}
			if err := oidcAuth(pcSpec, ps); err != nil {
				t.Fatalf("oidcAuth() returned unexpected error: %v", err)
			}
			got, _ := ps.Configuration[keyOidcTokenFilePath].(string)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("oidc_token_file_path mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_msiAuth(t *testing.T) {
	subID, tenantID, clientID, endpoint, environment := "sub", "tenant", "client", "endpoint", "environment"
	pcSpec := &namespacedv1beta1.ProviderConfigSpec{
		SubscriptionID: &subID,
		TenantID:       &tenantID,
		ClientID:       &clientID,
		MSIEndpoint:    &endpoint,
		Environment:    &environment,
	}
	ps := &terraform.Setup{Configuration: terraform.ProviderConfiguration{}}

	if err := msiAuth(pcSpec, ps); err != nil {
		t.Fatalf("msiAuth() returned unexpected error: %v", err)
	}

	want := terraform.ProviderConfiguration{
		keySubscriptionID: subID,
		keyTenantID:       tenantID,
		keyClientID:       clientID,
		keyUseMSI:         "true",
		keyMSIEndpoint:    endpoint,
		keyEnvironment:    environment,
	}
	if diff := cmp.Diff(want, ps.Configuration); diff != "" {
		t.Errorf("MSI configuration mismatch (-want +got):\n%s", diff)
	}
}

func Test_upboundAuth(t *testing.T) {
	subID, tenantID, clientID, environment := "sub", "tenant", "client", "environment"
	pcSpec := &namespacedv1beta1.ProviderConfigSpec{
		SubscriptionID: &subID,
		TenantID:       &tenantID,
		ClientID:       &clientID,
		Environment:    &environment,
	}
	ps := &terraform.Setup{Configuration: terraform.ProviderConfiguration{}}

	if err := upboundAuth(pcSpec, ps); err != nil {
		t.Fatalf("upboundAuth() returned unexpected error: %v", err)
	}

	want := terraform.ProviderConfiguration{
		keyOidcTokenFilePath: upboundProviderIdentityTokenFile,
		keySubscriptionID:    subID,
		keyTenantID:          tenantID,
		keyClientID:          clientID,
		keyUseOIDC:           "true",
		keyEnvironment:       environment,
	}
	if diff := cmp.Diff(want, ps.Configuration); diff != "" {
		t.Errorf("Upbound configuration mismatch (-want +got):\n%s", diff)
	}
}

func Test_sourceAuth_requiredFields(t *testing.T) {
	fields := []struct {
		name string
		make func(*namespacedv1beta1.ProviderConfigSpec, *terraform.Setup) error
		want string
	}{
		{name: "MSI subscription ID", make: msiAuth, want: errSubscriptionIDNotSet},
		{name: "OIDC subscription ID", make: oidcAuth, want: errSubscriptionIDNotSet},
		{name: "Upbound subscription ID", make: upboundAuth, want: errSubscriptionIDNotSet},
	}
	for _, tc := range fields {
		t.Run(tc.name, func(t *testing.T) {
			pcSpec := &namespacedv1beta1.ProviderConfigSpec{}
			ps := &terraform.Setup{Configuration: terraform.ProviderConfiguration{}}
			if err := tc.make(pcSpec, ps); err == nil || err.Error() != tc.want {
				t.Fatalf("expected error %q, got %v", tc.want, err)
			}
		})
	}
}

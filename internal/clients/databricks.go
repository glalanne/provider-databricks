/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/glalanne/provider-databricks/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal databricks credentials as JSON"

	keyHost                     = "host"
	keyAzureWorkspaceResourceID = "azure_workspace_resource_id"
	keyAzureUseMsi              = "azure_use_msi"
	keyAzureClientId            = "azure_client_id"
	keyAzureClientSecret        = "azure_client_secret"
	keyAzureTenantId            = "azure_tenant_id"
	keyClientId      			= "client_id"
	keyClientSecret 			= "client_secret"
	keyAccountId	 			= "account_id"
	keyAuthType                 = "auth_type"
	keyAuthToken                = "token"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn { // nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{}
		if v, ok := creds[keyHost]; ok {
			ps.Configuration[keyHost] = v
		}
		if v, ok := creds[keyAzureWorkspaceResourceID]; ok {
			ps.Configuration[keyAzureWorkspaceResourceID] = v
		}
		if v, ok := creds[keyAzureUseMsi]; ok {
			ps.Configuration[keyAzureUseMsi] = v
		}
		if v, ok := creds[keyAzureClientId]; ok {
			ps.Configuration[keyAzureClientId] = v
		}
		if v, ok := creds[keyAzureClientSecret]; ok {
			ps.Configuration[keyAzureClientSecret] = v
		}
		if v, ok := creds[keyAzureTenantId]; ok {
			ps.Configuration[keyAzureTenantId] = v
		}
		if v, ok := creds[keyAuthType]; ok {
			ps.Configuration[keyAuthType] = v
		}
		if v, ok := creds[keyAuthToken]; ok {
			ps.Configuration[keyAuthToken] = v
		}
		if v, ok := creds[keyClientId]; ok {
			ps.Configuration[keyClientId] = v
		}
		if v, ok := creds[keyClientSecret]; ok {
			ps.Configuration[keyClientSecret] = v
		}
		if v, ok := creds[keyAccountId]; ok {
			ps.Configuration[keyAccountId] = v
		}
		return ps, nil
	}
}

/*
Copyright 2026 The Crossplane Authors.
*/

package clients

import (
	"context"
	"net/http"

	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/diffserver"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	dbclient "github.com/databricks/databricks-sdk-go/client"
	dbconfig "github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"sigs.k8s.io/controller-runtime/pkg/client"

	namespacedv1beta1 "github.com/glalanne/provider-databricks/apis/namespaced/v1beta1"
)

const (
	offlineHost  = "https://offline.databricks.invalid"
	offlineToken = "offline-diff-placeholder"
)

// OfflineTerraformSetupBuilder configures SDKv2 and Framework provider meta for the diff
// server without reading credentials or contacting Databricks. ProviderConfig
// resolution still uses the supplied Kubernetes client; host and account ID
// must be set in its spec if they affect the diff. Secrets are never read.
// Unlike normal reconciliation, API-dependent diffs fail as unsupported.
func OfflineTerraformSetupBuilder(fwProvider fwprovider.Provider, tfProvider *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, kube client.Client, mg xpresource.Managed) (terraform.Setup, error) {
		pcSpec, err := resolveProviderConfig(ctx, kube, mg)
		if err != nil {
			return terraform.Setup{}, err
		}

		ps := terraform.Setup{Configuration: offlineConfiguration(pcSpec)}
		p := *tfProvider
		// The upstream Configure callback has no per-client SDK customization
		// hook. Replace it on this copy so the egress guard is installed before
		// SDK config resolution (which otherwise performs host discovery).
		p.ConfigureFunc = nil
		p.ConfigureProvider = nil
		p.ConfigureContextFunc = configureOfflineDatabricksClient
		if err := configureNoForkDBXClient(ctx, &ps, p); err != nil {
			return terraform.Setup{}, errors.Wrap(err, "failed to configure the offline Databricks client")
		}
		ps.FrameworkProvider = offlineFrameworkProvider{Provider: fwProvider, meta: ps.Meta}
		return ps, nil
	}
}

// Override only Configure: resource factories and schemas remain upstream's,
// while every resource receives the same per-request, egress-guarded client.
type offlineFrameworkProvider struct {
	fwprovider.Provider
	meta any
}

func (p offlineFrameworkProvider) Configure(_ context.Context, _ fwprovider.ConfigureRequest, resp *fwprovider.ConfigureResponse) {
	resp.ResourceData = p.meta
	resp.DataSourceData = p.meta
}

func offlineConfiguration(pcSpec *namespacedv1beta1.ProviderConfigSpec) map[string]any {
	cfg := map[string]any{
		keyHost:      offlineHost,
		keyAuthType:  authTypePAT,
		keyAuthToken: offlineToken,
	}
	if pcSpec.Host != nil && *pcSpec.Host != "" {
		cfg[keyHost] = *pcSpec.Host
	}
	if pcSpec.AccountID != nil && *pcSpec.AccountID != "" {
		cfg[keyAccountID] = *pcSpec.AccountID
	}
	return cfg
}

func configureOfflineDatabricksClient(_ context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	cfg := &dbconfig.Config{
		Host:          d.Get(keyHost).(string),
		AccountID:     d.Get(keyAccountID).(string),
		Token:         offlineToken,
		AuthType:      authTypePAT,
		Credentials:   dbconfig.PatCredentials{},
		HTTPTransport: offlineTransport{},
		// An empty loader list enables the SDK's environment/profile defaults.
		Loaders: []dbconfig.Loader{offlineConfigLoader{}},
		HostMetadataResolver: func(context.Context, string) (*dbconfig.HostMetadata, error) {
			// Offline diffs use the SDK's URL-based cloud/host classification.
			return nil, nil
		},
	}
	c, err := dbclient.New(cfg)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	meta := &common.DatabricksClient{DatabricksClient: c}
	meta.WithCommandExecutor(func(ctx context.Context, c *common.DatabricksClient) common.CommandExecutor {
		return commands.NewCommandsAPI(ctx, c)
	})
	return meta, nil
}

type offlineConfigLoader struct{}

func (offlineConfigLoader) Name() string                     { return "offline" }
func (offlineConfigLoader) Configure(*dbconfig.Config) error { return nil }

type offlineTransport struct{}

func (offlineTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Omit host, query and body, which may contain credentials or resource data.
	return nil, diffserver.NewDiffComputationNotSupportedError(errors.Errorf(
		"offline diff server blocked an outbound Databricks request (%s %s): computing this diff requires an API call",
		req.Method, req.URL.Path))
}

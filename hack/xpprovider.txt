package xpprovider

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetProvider(_ context.Context) (fwprovider.Provider, *schema.Provider, error) {
	return pluginfw.GetDatabricksProviderPluginFramework(),
		sdkv2.DatabricksProvider(), nil
}

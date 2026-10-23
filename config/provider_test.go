package config

import (
	"context"
	"testing"

	upjetconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/xpprovider"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestProviderLibrarySDKCompatibility(t *testing.T) {
	t.Setenv("USE_SDK_V2_RESOURCES", "")
	for _, scope := range []struct {
		name string
		get  func(context.Context, fwprovider.Provider, *schema.Provider, bool) (*upjetconfig.Provider, error)
	}{
		{name: "cluster", get: GetProvider},
		{name: "namespaced", get: GetProviderNamespaced},
		{name: "cluster/legacy", get: GetProviderV1Alpha1Legacy},
		{name: "namespaced/legacy", get: GetProviderNamespacedV1Alpha1Legacy},
	} {
		for _, generation := range []bool{false, true} {
			mode := "runtime"
			if generation {
				mode = "generation"
			}
			t.Run(scope.name+"/"+mode, func(t *testing.T) {
				ctx := t.Context()
				fw, sdk, err := xpprovider.GetProvider(ctx)
				if err != nil {
					t.Fatal(err)
				}
				delete(sdk.ResourcesMap, "databricks_library")
				p, err := scope.get(ctx, fw, sdk, generation)
				if err != nil {
					t.Fatal(err)
				}
				r := p.Resources["databricks_library"]
				if r == nil || !r.ShouldUseTerraformPluginSDKClient() || r.ShouldUseTerraformPluginFrameworkClient() {
					t.Fatal("Library must continue using the SDKv2 reconciler")
				}
				if len(r.TFListConversionPaths()) != 0 || len(r.CRDListConversionPaths()) != 0 {
					t.Fatal("Library must retain list-shaped blocks in both API versions")
				}
				sdkLibrary := sdk.ResourcesMap["databricks_library"]
				if sdkLibrary == nil || sdkLibrary.CreateContext == nil || sdkLibrary.ReadContext == nil || sdkLibrary.DeleteContext == nil {
					t.Fatal("Library must have the SDKv2 schema and CRUD implementation")
				}
				if !generation && r.TerraformResource != sdkLibrary {
					t.Fatal("runtime Library must use the restored SDKv2 resource")
				}
			})
		}
	}
}

func TestProviderPreservesSDKLibrary(t *testing.T) {
	ctx := t.Context()
	fw, sdk, err := xpprovider.GetProvider(ctx)
	if err != nil {
		t.Fatal(err)
	}
	library := clusters.ResourceLibrary().ToResource()
	sdk.ResourcesMap["databricks_library"] = library
	for range 2 {
		p, err := GetProvider(ctx, fw, sdk, false)
		if err != nil {
			t.Fatal(err)
		}
		if p.Resources["databricks_library"].TerraformResource != library {
			t.Fatal("an existing SDKv2 Library resource must not be replaced")
		}
	}
}

package cluster

import (
	"context"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestConfigureResourcePreservesReconstructedID(t *testing.T) {
	const externalName = "0827-181218-mkm88eql"

	r := &config.Resource{
		ExternalName:      config.IdentifierFromProvider,
		TerraformResource: clusters.ResourceCluster().ToResource(),
	}
	configureResource(r)

	id, err := r.ExternalName.GetIDFn(context.Background(), externalName, nil, nil)
	if err != nil {
		t.Fatalf("GetIDFn() error = %v", err)
	}
	stateValue, err := schema.JSONMapToStateValue(map[string]any{
		"cluster_name": "test-cluster",
		"id":           id,
		"library": []any{
			map[string]any{
				"cran":  []any{map[string]any{}},
				"maven": []any{map[string]any{}},
				"pypi": []any{
					map[string]any{
						"package": "cleo==2.0.0",
						"repo":    "",
					},
				},
			},
		},
	}, r.TerraformResource.CoreConfigSchema())
	if err != nil {
		t.Fatalf("JSONMapToStateValue() error = %v", err)
	}
	state, err := r.TerraformResource.ShimInstanceStateFromValue(stateValue)
	if err != nil {
		t.Fatalf("ShimInstanceStateFromValue() error = %v", err)
	}
	if state.ID != externalName {
		t.Fatalf("reconstructed state ID = %q, want %q", state.ID, externalName)
	}
}

package pipeline

import (
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_pipeline", func(r *config.Resource) {
		r.ShortGroup = "compute"

		r.ExternalName.OmittedFields = []string{
			"last_modified",
		}

		// Move to computed so that it is not required in the spec.
		if s, ok := r.TerraformResource.Schema["state"]; ok {
			s.Optional = false
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["latest_updates"]; ok {
			s.Optional = false
			s.Computed = true
		}

		r.TerraformCustomDiff = func(
			diff *terraform.InstanceDiff,
			state *terraform.InstanceState,
			cfg *terraform.ResourceConfig,
		) (*terraform.InstanceDiff, error) {
			if diff == nil {
				return diff, nil
			}

			for k := range diff.Attributes {
				if strings.HasPrefix(k, "latest_updates.") {
					delete(diff.Attributes, k)
				}
			}

			return diff, nil
		}
	})
}

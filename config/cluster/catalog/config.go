package catalog

import (
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_catalog", func(r *config.Resource) {
		r.ShortGroup = "unity"
		r.SchemaElementOptions.SetAddToObservation("provisioning_info")

		r.TerraformCustomDiff = func(
			diff *terraform.InstanceDiff,
			state *terraform.InstanceState,
			cfg *terraform.ResourceConfig,
		) (*terraform.InstanceDiff, error) {
			if diff == nil {
				return diff, nil
			}

			delete(diff.Attributes, "provisioning_info.#")

			for k := range diff.Attributes {
				if strings.HasPrefix(k, "provisioning_info.") {
					delete(diff.Attributes, k)
				}
			}

			return diff, nil
		}
	})
}

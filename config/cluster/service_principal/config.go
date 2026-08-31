package service_principal

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_service_principal", func(r *config.Resource) {
		r.ShortGroup = "security"

		// The Databricks SCIM API cannot update displayName of an existing
		// service principal, so drop it from the diff after creation.
		// https://community.databricks.com/t5/administration-architecture/how-to-change-the-display-name-for-a-service-principal/td-p/139696
		r.TerraformCustomDiff = func(
			diff *terraform.InstanceDiff,
			state *terraform.InstanceState,
			cfg *terraform.ResourceConfig,
		) (*terraform.InstanceDiff, error) {
			if diff == nil || diff.Destroy || state == nil || state.ID == "" {
				return diff, nil
			}

			delete(diff.Attributes, "display_name")

			return diff, nil
		}
	})
}

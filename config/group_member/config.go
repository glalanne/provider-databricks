package group_member

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_group_member", func(r *config.Resource) {
		r.ShortGroup = "security"
		r.References["group_id"] = config.Reference{
			TerraformName: "databricks_group",
		}
	})
}

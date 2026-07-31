package environments_workspace_base_environment

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_environments_workspace_base_environment", func(r *config.Resource) {
		r.ShortGroup = "envs"
	})
}

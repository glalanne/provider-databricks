package workspace_file

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_workspace_file", func(r *config.Resource) {
		r.ShortGroup = "workspace"
	})
}

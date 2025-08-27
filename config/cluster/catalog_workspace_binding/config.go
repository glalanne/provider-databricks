package catalog_workspace_binding

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_catalog_workspace_binding", func(r *config.Resource) {
		r.ShortGroup = "unity"
	})
}

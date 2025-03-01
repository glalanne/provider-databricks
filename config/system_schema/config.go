package system_schema

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_system_schema", func(r *config.Resource) {
		r.ShortGroup = "unity"
	})
}

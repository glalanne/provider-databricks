package online_table

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_online_table", func(r *config.Resource) {
		r.ShortGroup = "unity"
	})
}

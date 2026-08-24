package data_quality_refresh

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_data_quality_refresh", func(r *config.Resource) {
		r.ShortGroup = "unity"
	})
}

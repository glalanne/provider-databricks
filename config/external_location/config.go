package external_location

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_external_location", func(r *config.Resource) {
		r.ShortGroup = "unity"
	})
}

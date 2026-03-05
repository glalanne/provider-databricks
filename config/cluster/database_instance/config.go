package database_instance

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_database_instance", func(r *config.Resource) {
		r.ShortGroup = "databases"
	})
}

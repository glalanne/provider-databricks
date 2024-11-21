package sql_visualization

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_sql_visualization", func(r *config.Resource) {
		r.ShortGroup = "sql"
	})
}

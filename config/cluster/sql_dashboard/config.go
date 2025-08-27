package sql_dashboard

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_sql_dashboard", func(r *config.Resource) {
		r.ShortGroup = "sql"
	})
}

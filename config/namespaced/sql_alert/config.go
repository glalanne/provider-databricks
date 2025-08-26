package sql_alert

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_sql_alert", func(r *config.Resource) {
		r.ShortGroup = "sql"
	})
}

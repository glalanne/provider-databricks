package mws_log_delivery

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_mws_log_delivery", func(r *config.Resource) {
		r.ShortGroup = "log"
	})
}

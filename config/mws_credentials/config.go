package mws_credentials

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_mws_credentials", func(r *config.Resource) {
		r.ShortGroup = "deployment"
	})
}

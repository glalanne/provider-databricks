package mws_private_access_settings

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_mws_private_access_settings", func(r *config.Resource) {
		r.ShortGroup = "databricks"
	})
}
package mws_network_connectivity_config

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_mws_network_connectivity_config", func(r *config.Resource) {
		r.ShortGroup = "deployment"
	})
}

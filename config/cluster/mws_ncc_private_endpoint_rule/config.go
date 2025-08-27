package mws_ncc_private_endpoint_rule

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_mws_ncc_private_endpoint_rule", func(r *config.Resource) {
		r.ShortGroup = "deployment"
	})
}

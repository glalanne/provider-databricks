package mws_permission_assignment

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_mws_permission_assignment", func(r *config.Resource) {
		r.ShortGroup = "security"
	})
}

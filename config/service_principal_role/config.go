package service_principal_role

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_service_principal_role", func(r *config.Resource) {
		r.ShortGroup = "security"
	})
}

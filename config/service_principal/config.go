package service_principal

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_service_principal", func(r *config.Resource) {
		r.ShortGroup = "security"
	})
}

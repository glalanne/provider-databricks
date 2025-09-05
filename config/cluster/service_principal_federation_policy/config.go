package service_principal_federation_policy

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
// TODO: waiting on protov6 support with Upjet https://github.com/crossplane/upjet/issues/372
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_service_principal_federation_policy", func(r *config.Resource) {
		r.ShortGroup = "oauth"
	})
}

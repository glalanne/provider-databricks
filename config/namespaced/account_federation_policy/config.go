package account_federation_policy

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_account_federation_policy", func(r *config.Resource) {
		r.ShortGroup = "oauth"
	})
}

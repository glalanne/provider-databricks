package disaster_recovery_stable_url

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_disaster_recovery_stable_url", func(r *config.Resource) {
		r.ShortGroup = "dr"
	})
}

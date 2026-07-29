package disaster_recovery_failover_group

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_disaster_recovery_failover_group", func(r *config.Resource) {
		r.ShortGroup = "dr"
	})
}

package access_control_rule_set

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_access_control_rule_set", func(r *config.Resource) {
		r.ShortGroup = "security"
	})
}

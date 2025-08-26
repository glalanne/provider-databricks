package compliance_security_profile_workspace_setting

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_compliance_security_profile_workspace_setting", func(r *config.Resource) {
		r.ShortGroup = "settings"
	})
}

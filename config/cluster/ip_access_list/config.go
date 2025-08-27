package ip_access_list

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_ip_access_list", func(r *config.Resource) {
		r.ShortGroup = "security"
	})
}

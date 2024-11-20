package aws_s3_mount

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_aws_s3_mount", func(r *config.Resource) {
		r.ShortGroup = "databricks"
	})
}
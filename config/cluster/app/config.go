package app

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
// Does not support nested_types for now
// https://github.com/crossplane/upjet/issues/372
// https://github.com/crossplane/upjet/pull/448
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_app", func(r *config.Resource) {
		r.ShortGroup = "apps"
	})
}

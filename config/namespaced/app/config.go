package app

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_app", func(r *config.Resource) {
		r.ShortGroup = "apps"
	})
}

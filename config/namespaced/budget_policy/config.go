package budget_policy

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/glalanne/provider-databricks/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_budget_policy", func(r *config.Resource) {
		r.ShortGroup = "billing"
		r.TerraformPluginFrameworkIsStateEmptyFn = common.BudgetPolicyIsStateEmpty
	})
}

package common

import (
	"fmt"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// IdentifierArgumentExternalName is an external-name configuration for
// plugin framework resources whose only identifier is the computed attribute
// field and which have no "id" attribute. The external name is the value of
// field. Because upjet only reaches the Terraform prior state through
// SetIdentifierArgumentFn when there is no "id", it also writes the external
// name back to field; without that, Read runs with no identifier and the state
// comes back empty.
func IdentifierArgumentExternalName(field string) config.ExternalName {
	en := config.NewExternalNameFrom(
		config.IdentifierFromProvider,
		config.WithSetIdentifierArgumentsFn(func(_ config.SetIdentifierArgumentsFn, base map[string]any, externalName string) {
			if externalName != "" {
				base[field] = externalName
			}
		}),
		config.WithGetExternalNameFn(func(_ config.GetExternalNameFn, tfState map[string]any) (string, error) {
			value, ok := tfState[field]
			if !ok || value == nil || fmt.Sprint(value) == "" {
				return "", fmt.Errorf("cannot find %q in tfstate", field)
			}
			return fmt.Sprint(value), nil
		}),
	)
	// field is optional+computed, so keep it out of the config sent to
	// Terraform; it is still set on the prior state via SetIdentifierArgumentFn.
	en.TFPluginFrameworkOptions.ComputedIdentifierAttributes = []string{field}
	return en
}

var BudgetPolicyExternalName = IdentifierArgumentExternalName("policy_id")

package common

import (
	"context"
	"fmt"

	"github.com/crossplane/upjet/v2/pkg/config"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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

// BudgetPolicyIsStateEmpty prevents a non-null Read result without an identifier
// from being treated as an existing policy and passed to the external-name
// extractor. Unknown or malformed state must not authorize resource creation.
func BudgetPolicyIsStateEmpty(_ context.Context, state tftypes.Value, _ rschema.Schema) (bool, error) {
	if !state.IsKnown() {
		return false, fmt.Errorf("cannot check BudgetPolicy existence: Terraform state is unknown")
	}
	if state.IsNull() {
		return true, nil
	}
	if _, ok := state.Type().(tftypes.Object); !ok {
		return false, fmt.Errorf("cannot check BudgetPolicy existence: expected object state, got %s", state.Type())
	}
	var attributes map[string]tftypes.Value
	if err := state.As(&attributes); err != nil {
		return false, fmt.Errorf("cannot decode BudgetPolicy state: %w", err)
	}
	id, ok := attributes["policy_id"]
	if !ok {
		return true, nil
	}
	if !id.IsKnown() {
		return false, fmt.Errorf("cannot check BudgetPolicy existence: policy_id is unknown")
	}
	if !id.Type().Is(tftypes.String) {
		return false, fmt.Errorf("cannot check BudgetPolicy existence: expected string policy_id, got %s", id.Type())
	}
	if id.IsNull() {
		return true, nil
	}
	var policyID string
	if err := id.As(&policyID); err != nil {
		return false, fmt.Errorf("cannot decode BudgetPolicy policy_id: %w", err)
	}
	return policyID == "", nil
}

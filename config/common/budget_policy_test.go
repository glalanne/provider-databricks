package common_test

import (
	"strings"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	cluster "github.com/glalanne/provider-databricks/config/cluster/budget_policy"
	"github.com/glalanne/provider-databricks/config/common"
	namespaced "github.com/glalanne/provider-databricks/config/namespaced/budget_policy"
)

func TestBudgetPolicyStateExistenceBothScopes(t *testing.T) {
	stringStateType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{
		"policy_id": tftypes.String, "policy_name": tftypes.String,
	}}
	stateWithID := func(id any) tftypes.Value {
		return tftypes.NewValue(stringStateType, map[string]tftypes.Value{
			"policy_id":   tftypes.NewValue(tftypes.String, id),
			"policy_name": tftypes.NewValue(tftypes.String, "existing-name"),
		})
	}
	cases := []struct {
		name      string
		state     tftypes.Value
		wantEmpty bool
		wantError string
	}{
		{name: "null state", state: tftypes.NewValue(stringStateType, nil), wantEmpty: true},
		{name: "empty object", state: tftypes.NewValue(tftypes.Object{AttributeTypes: map[string]tftypes.Type{}}, map[string]tftypes.Value{}), wantEmpty: true},
		{name: "missing id with other fields", state: tftypes.NewValue(tftypes.Object{AttributeTypes: map[string]tftypes.Type{"policy_name": tftypes.String}},
			map[string]tftypes.Value{"policy_name": tftypes.NewValue(tftypes.String, "existing-name")}), wantEmpty: true},
		{name: "null id", state: stateWithID(nil), wantEmpty: true},
		{name: "empty id", state: stateWithID(""), wantEmpty: true},
		{name: "existing policy", state: stateWithID("policy-123")},
		{name: "unknown state", state: tftypes.NewValue(stringStateType, tftypes.UnknownValue), wantError: "state is unknown"},
		{name: "unknown id", state: stateWithID(tftypes.UnknownValue), wantError: "policy_id is unknown"},
		{name: "non-object state", state: tftypes.NewValue(tftypes.String, "invalid"), wantError: "expected object state"},
		{name: "non-string id", state: tftypes.NewValue(tftypes.Object{AttributeTypes: map[string]tftypes.Type{"policy_id": tftypes.Number}},
			map[string]tftypes.Value{"policy_id": tftypes.NewValue(tftypes.Number, 123)}), wantError: "expected string policy_id"},
	}
	for scope, configure := range map[string]func(*config.Provider){"cluster": cluster.Configure, "namespaced": namespaced.Configure} {
		t.Run(scope, func(t *testing.T) {
			p := config.NewProvider([]byte(`{"format_version":"0.2","provider_schemas":{"databricks":{"resource_schemas":{}}}}`),
				"databricks", "github.com/glalanne/provider-databricks", []byte("{}"))
			r := &config.Resource{ExternalName: common.BudgetPolicyExternalName}
			unrelated := &config.Resource{}
			p.Resources["databricks_budget_policy"] = r
			p.Resources["databricks_other"] = unrelated
			configure(p)
			p.ConfigureResources()
			if r.ShortGroup != "billing" || r.TerraformPluginFrameworkIsStateEmptyFn == nil {
				t.Fatal("BudgetPolicy configurator did not register its billing group and state check")
			}
			if unrelated.TerraformPluginFrameworkIsStateEmptyFn != nil {
				t.Fatal("state check must be resource-scoped")
			}
			for _, tc := range cases {
				t.Run(tc.name, func(t *testing.T) {
					empty, err := r.TerraformPluginFrameworkIsStateEmptyFn(t.Context(), tc.state, rschema.Schema{})
					if tc.wantError != "" {
						if err == nil || !strings.Contains(err.Error(), tc.wantError) || empty {
							t.Fatalf("expected non-empty and %q error, got empty=%v, err=%v", tc.wantError, empty, err)
						}
						return
					}
					if err != nil || empty != tc.wantEmpty {
						t.Fatalf("empty=%v, err=%v; want empty=%v", empty, err, tc.wantEmpty)
					}
					if !empty {
						id, err := r.ExternalName.GetExternalNameFn(map[string]any{"policy_id": "policy-123"})
						if err != nil || id != "policy-123" {
							t.Fatalf("existing policy external name=%q, err=%v", id, err)
						}
					}
				})
			}
		})
	}
}

package common

import (
	"reflect"
	"strings"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/databricks/terraform-provider-databricks/xpprovider"
)

func TestEmptyObjectCleaner(t *testing.T) {
	tests := map[string]struct {
		mode   config.Mode
		params map[string]any
		want   map[string]any
	}{
		"clears selected empty objects from Terraform": {
			mode: config.FromTerraform,
			params: map[string]any{
				"schedule": map[string]any{},
				"trigger":  map[string]any{},
			},
			want: map[string]any{
				"schedule": nil,
				"trigger":  nil,
			},
		},
		"preserves populated and unselected objects": {
			mode: config.FromTerraform,
			params: map[string]any{
				"schedule": map[string]any{"timezone_id": "UTC"},
				"other":    map[string]any{},
			},
			want: map[string]any{
				"schedule": map[string]any{"timezone_id": "UTC"},
				"other":    map[string]any{},
			},
		},
		"does not alter parameters sent to Terraform": {
			mode: config.ToTerraform,
			params: map[string]any{
				"schedule": map[string]any{},
			},
			want: map[string]any{
				"schedule": map[string]any{},
			},
		},
	}

	cleaner := EmptyObjectCleaner{"schedule", "trigger"}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := cleaner.Convert(tt.params, nil, tt.mode)
			if err != nil {
				t.Fatalf("Convert: %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert: got %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestEmptyObjectCleanerAfterSingletonConversion(t *testing.T) {
	params := map[string]any{}
	cleaner := JobEmptyObjectCleaner(configuredJobResource(t))
	for _, field := range cleaner {
		if strings.Contains(field, "[*]") || strings.Contains(field, ".") {
			continue
		}
		// Upjet's singleton-list conversion produces an empty map for an
		// empty list; TypeMap fields such as tags have the same shape.
		params[field] = map[string]any{}
	}
	got, err := cleaner.Convert(params, nil, config.FromTerraform)
	if err != nil {
		t.Fatalf("Convert: %v", err)
	}
	for field := range params {
		if got[field] != nil {
			t.Errorf("%s: got %#v, want nil", field, got[field])
		}
	}
}

func TestJobEmptyObjectCleanerClearsNestedObject(t *testing.T) {
	params := map[string]any{
		"task": []any{map[string]any{
			"task_key":            "main",
			"email_notifications": map[string]any{},
		}},
	}
	cleaned, err := JobEmptyObjectCleaner(configuredJobResource(t)).Convert(params, nil, config.FromTerraform)
	if err != nil {
		t.Fatalf("Convert: %v", err)
	}
	tasks := cleaned["task"].([]any)
	nested := tasks[0].(map[string]any)["email_notifications"]
	if nested != nil {
		t.Fatalf("email_notifications: got %#v, want nil", nested)
	}
}

func configuredJobResource(t *testing.T) *config.Resource {
	t.Helper()
	_, sdkProvider, err := xpprovider.GetProvider(t.Context())
	if err != nil {
		t.Fatalf("GetProvider: %s", err)
	}
	job, ok := sdkProvider.ResourcesMap["databricks_job"]
	if !ok {
		t.Fatal("databricks_job is not in the provider schema")
	}
	r := config.DefaultResource("databricks_job", job, nil, nil)
	if err := config.TraverseSchemas("databricks_job", r, &config.SingletonListEmbedder{}); err != nil {
		t.Fatalf("TraverseSchemas: %v", err)
	}
	return r
}

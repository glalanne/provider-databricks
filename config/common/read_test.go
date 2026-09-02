package common

import (
	"context"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func testResource(read schema.ReadContextFunc) *config.Resource {
	return &config.Resource{
		TerraformResource: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"schedule": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"quartz_cron_expression": {
								Type:     schema.TypeString,
								Optional: true,
							},
						},
					},
				},
				"tags": {
					Type:     schema.TypeMap,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			},
			ReadContext: read,
		},
	}
}

func staleData(t *testing.T, r *config.Resource) *schema.ResourceData {
	t.Helper()
	return schema.TestResourceDataRaw(t, r.TerraformResource.Schema, map[string]any{
		"name": "job",
		"schedule": []any{map[string]any{
			"quartz_cron_expression": "0 0 12 * * ?",
		}},
		"tags": map[string]any{"env": "dev"},
	})
}

func TestClearFieldsBeforeReadClearsStaleValues(t *testing.T) {
	r := testResource(func(_ context.Context, _ *schema.ResourceData, _ any) diag.Diagnostics {
		// The Databricks API omits a removed schedule, so the upstream Read
		// does not touch the field.
		return nil
	})
	ClearFieldsBeforeRead(r, "schedule", "tags")

	d := staleData(t, r)
	if diags := r.TerraformResource.ReadContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if got := d.Get("schedule").([]any); len(got) != 0 {
		t.Errorf("expected schedule to be cleared, got %v", got)
	}
	if got := d.Get("tags").(map[string]any); len(got) != 0 {
		t.Errorf("expected tags to be cleared, got %v", got)
	}
	if got := d.Get("name").(string); got != "job" {
		t.Errorf("expected name to be preserved, got %q", got)
	}
}

func TestClearFieldsBeforeReadKeepsValuesReturnedByRead(t *testing.T) {
	r := testResource(func(_ context.Context, d *schema.ResourceData, _ any) diag.Diagnostics {
		if err := d.Set("schedule", []any{map[string]any{
			"quartz_cron_expression": "0 0 6 * * ?",
		}}); err != nil {
			return diag.FromErr(err)
		}
		return nil
	})
	ClearFieldsBeforeRead(r, "schedule")

	d := staleData(t, r)
	if diags := r.TerraformResource.ReadContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	got := d.Get("schedule").([]any)
	if len(got) != 1 {
		t.Fatalf("expected schedule to be populated, got %v", got)
	}
	if cron := got[0].(map[string]any)["quartz_cron_expression"]; cron != "0 0 6 * * ?" {
		t.Errorf("expected the schedule returned by Read, got %v", cron)
	}
}

func TestClearFieldsBeforeReadWrapsOnce(t *testing.T) {
	reads := 0
	r := testResource(func(_ context.Context, _ *schema.ResourceData, _ any) diag.Diagnostics {
		reads++
		return nil
	})

	// The cluster-scoped and namespaced configurations share the same
	// Terraform resource pointer.
	ClearFieldsBeforeRead(r, "schedule")
	ClearFieldsBeforeRead(r, "tags")

	cleanersMu.Lock()
	c, ok := cleaners[r.TerraformResource]
	cleanersMu.Unlock()
	if !ok {
		t.Fatal("expected the resource to be registered")
	}
	if len(c.names()) != 2 {
		t.Errorf("expected both fields to be registered on a single wrapper, got %v", c.names())
	}

	d := staleData(t, r)
	if diags := r.TerraformResource.ReadContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if reads != 1 {
		t.Errorf("expected the original Read to run once, ran %d times", reads)
	}
}

func TestClearFieldsBeforeReadIgnoresUnknownAndPrimitiveFields(t *testing.T) {
	r := testResource(func(_ context.Context, _ *schema.ResourceData, _ any) diag.Diagnostics {
		return nil
	})
	ClearFieldsBeforeRead(r, "name", "does_not_exist")

	cleanersMu.Lock()
	c := cleaners[r.TerraformResource]
	cleanersMu.Unlock()
	if len(c.names()) != 0 {
		t.Errorf("expected no fields to be registered, got %v", c.names())
	}
}

func TestClearFieldsBeforeReadDuringGeneration(t *testing.T) {
	// The Terraform provider is not wired in while generating the APIs.
	r := testResource(nil)
	ClearFieldsBeforeRead(r, "schedule")

	if r.TerraformResource.ReadContext != nil {
		t.Error("expected ReadContext to stay nil")
	}
	ClearFieldsBeforeRead(nil, "schedule")
	ClearFieldsBeforeRead(&config.Resource{}, "schedule")
	ClearStaleBlocksBeforeRead(nil)
	ClearStaleBlocksBeforeRead(&config.Resource{})
}

func TestClearStaleBlocksBeforeReadSelectsFields(t *testing.T) {
	block := func(s map[string]*schema.Schema) *schema.Schema {
		return &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     &schema.Resource{Schema: s},
		}
	}
	str := &schema.Schema{Type: schema.TypeString, Optional: true}
	secret := &schema.Schema{Type: schema.TypeString, Optional: true, Sensitive: true}

	r := &config.Resource{TerraformResource: &schema.Resource{
		ReadContext: func(_ context.Context, _ *schema.ResourceData, _ any) diag.Diagnostics {
			return nil
		},
		Schema: map[string]*schema.Schema{
			"name":     str,
			"schedule": block(map[string]*schema.Schema{"cron": str}),
			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"task": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Resource{Schema: map[string]*schema.Schema{"key": str}},
			},
			"state":           {Type: schema.TypeList, Computed: true, Elem: &schema.Resource{Schema: map[string]*schema.Schema{"phase": str}}},
			"token":           block(map[string]*schema.Schema{"value": secret}),
			"nested_token":    block(map[string]*schema.Schema{"inner": block(map[string]*schema.Schema{"value": secret})}),
			"provider_config": block(map[string]*schema.Schema{"host": str}),
		},
	}}

	ClearStaleBlocksBeforeRead(r, "provider_config")

	cleanersMu.Lock()
	c := cleaners[r.TerraformResource]
	cleanersMu.Unlock()

	got := map[string]bool{}
	for _, n := range c.names() {
		got[n] = true
	}
	want := []string{"schedule", "tags"}
	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, c.names())
	}
	for _, w := range want {
		if !got[w] {
			t.Errorf("expected %q to be cleared, got %v", w, c.names())
		}
	}
}

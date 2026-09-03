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
				"keys": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"task": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"key": {
								Type:     schema.TypeString,
								Optional: true,
							},
							"email_notifications": {
								Type:     schema.TypeList,
								Optional: true,
								MaxItems: 1,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"on_failure": {
											Type:     schema.TypeList,
											Optional: true,
											Elem:     &schema.Schema{Type: schema.TypeString},
										},
									},
								},
							},
						},
					},
				},
			},
			ReadContext: read,
		},
	}
}

func staleData(t *testing.T, r *config.Resource) *schema.ResourceData {
	t.Helper()
	d := schema.TestResourceDataRaw(t, r.TerraformResource.Schema, map[string]any{
		"name": "job",
		"schedule": []any{map[string]any{
			"quartz_cron_expression": "0 0 12 * * ?",
		}},
		"tags": map[string]any{"env": "dev"},
		"keys": []any{"ssh-rsa AAA"},
		"task": []any{map[string]any{
			"key": "main",
			"email_notifications": []any{map[string]any{
				"on_failure": []any{"oncall@example.com"},
			}},
		}},
	})
	d.SetId("1234")
	return d
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
		if got := d.Get("name").(string); got != "job" {
			return diag.Errorf("scalar read input was cleared: name=%q", got)
		}
		if !d.IsNewResource() {
			return diag.Errorf("expected refresh data to be marked as new")
		}
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

func TestClearFieldsBeforeReadPreservesDefaultSuppression(t *testing.T) {
	r := testResource(func(_ context.Context, d *schema.ResourceData, _ any) diag.Diagnostics {
		if err := d.Set("name", "server default"); err != nil {
			return diag.FromErr(err)
		}
		return nil
	})
	ClearFieldsBeforeRead(r, "schedule")

	d := schema.TestResourceDataRaw(t, r.TerraformResource.Schema, map[string]any{})
	d.SetId("1234")
	if diags := r.TerraformResource.ReadContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if got := d.Get("name").(string); got != "" {
		t.Errorf("expected unconfigured field to remain suppressed, got %q", got)
	}
}

func TestClearFieldsBeforeReadRefreshesConfiguredScalar(t *testing.T) {
	r := testResource(func(_ context.Context, d *schema.ResourceData, _ any) diag.Diagnostics {
		if err := d.Set("name", "remote name"); err != nil {
			return diag.FromErr(err)
		}
		return nil
	})
	ClearFieldsBeforeRead(r, "schedule")

	d := schema.TestResourceDataRaw(t, r.TerraformResource.Schema, map[string]any{"name": "configured name"})
	d.SetId("1234")
	if diags := r.TerraformResource.ReadContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if got := d.Get("name").(string); got != "remote name" {
		t.Errorf("expected configured scalar to refresh, got %q", got)
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

func TestClearFieldsBeforeReadClearsSets(t *testing.T) {
	r := testResource(func(_ context.Context, _ *schema.ResourceData, _ any) diag.Diagnostics {
		return nil
	})
	ClearFieldsBeforeRead(r, "keys")

	d := staleData(t, r)
	if diags := r.TerraformResource.ReadContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if got := d.Get("keys").(*schema.Set); got.Len() != 0 {
		t.Errorf("expected keys to be cleared, got %v", got.List())
	}
}

func TestClearFieldsBeforeReadReplacesNestedValues(t *testing.T) {
	// task is not registered for clearing: setting the parent must still drop
	// the nested email_notifications the API no longer returns.
	r := testResource(func(_ context.Context, d *schema.ResourceData, _ any) diag.Diagnostics {
		if err := d.Set("task", []any{map[string]any{"key": "main"}}); err != nil {
			return diag.FromErr(err)
		}
		return nil
	})
	ClearFieldsBeforeRead(r, "schedule")

	d := staleData(t, r)
	if diags := r.TerraformResource.ReadContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	task := d.Get("task").([]any)
	if len(task) != 1 {
		t.Fatalf("expected one task, got %v", task)
	}
	if got := task[0].(map[string]any)["email_notifications"].([]any); len(got) != 0 {
		t.Errorf("expected the nested block to be replaced, got %v", got)
	}
}

func TestClearFieldsBeforeReadKeepsID(t *testing.T) {
	r := testResource(func(_ context.Context, _ *schema.ResourceData, _ any) diag.Diagnostics {
		return nil
	})
	ClearFieldsBeforeRead(r, "schedule", "tags", "keys", "task")

	d := staleData(t, r)
	if diags := r.TerraformResource.ReadContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if d.Id() != "1234" {
		t.Errorf("expected the ID to be preserved, got %q", d.Id())
	}
}

func TestClearFieldsBeforeReadPropagatesDiagnostics(t *testing.T) {
	r := testResource(func(_ context.Context, _ *schema.ResourceData, _ any) diag.Diagnostics {
		return diag.Errorf("boom")
	})
	ClearFieldsBeforeRead(r, "schedule")

	diags := r.TerraformResource.ReadContext(context.Background(), staleData(t, r), nil)
	if !diags.HasError() {
		t.Fatal("expected the upstream error to be propagated")
	}
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
			"retries":  {Type: schema.TypeInt, Optional: true},
			"enabled":  {Type: schema.TypeBool, Optional: true},
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
			"parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ssh_public_keys": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"secret_values": {
				Type:      schema.TypeList,
				Optional:  true,
				Sensitive: true,
				Elem:      &schema.Schema{Type: schema.TypeString},
			},
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
	want := []string{"schedule", "tags", "parameters", "ssh_public_keys"}
	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, c.names())
	}
	for _, w := range want {
		if !got[w] {
			t.Errorf("expected %q to be cleared, got %v", w, c.names())
		}
	}
}

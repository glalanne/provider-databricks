package common

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/databricks/terraform-provider-databricks/xpprovider"
)

// TestUseAuthoritativeReadOnJobSchema pins the selection against the real
// databricks_job schema, so an upstream schema change that flips a flag is
// caught here rather than in a cluster.
func TestUseAuthoritativeReadOnJobSchema(t *testing.T) {
	_, sdkProvider, err := xpprovider.GetProvider(t.Context())
	if err != nil {
		t.Fatalf("GetProvider: %s", err)
	}
	job, ok := sdkProvider.ResourcesMap["databricks_job"]
	if !ok {
		t.Fatal("databricks_job is not in the provider schema")
	}

	r := &config.Resource{TerraformResource: job}
	UseAuthoritativeRead(r, "provider_config", "always_running", "control_run_state")

	cleanersMu.Lock()
	c := cleaners[job]
	cleanersMu.Unlock()
	if c == nil {
		t.Fatal("expected the job resource to be registered")
	}

	selected := map[string]bool{}
	for _, n := range c.names() {
		selected[n] = true
	}

	for _, name := range []string{
		"name", "description", "timeout_seconds", "max_retries", "max_concurrent_runs", "edit_mode",
		"schedule", "continuous", "trigger", "queue", "health", "deployment", "run_job_task",
		"git_source", "email_notifications", "webhook_notifications", "notification_settings",
		"library", "parameter", "environment", "tags",
	} {
		if !selected[name] {
			t.Errorf("expected %q to be cleared before the Read", name)
		}
	}

	for _, name := range []string{
		// Client-side routing, never returned by the Jobs API.
		"provider_config",
		// Carry a nested sensitive value, such as the docker image credentials,
		// that the API does not return either.
		"task", "job_cluster", "new_cluster",
		// Computed fields and write-only lifecycle controls.
		"id", "format", "url", "always_running", "control_run_state",
	} {
		if selected[name] {
			t.Errorf("did not expect %q to be cleared before the Read", name)
		}
	}
}

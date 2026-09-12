package common

import (
	"strings"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/databricks/terraform-provider-databricks/xpprovider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestClearStaleBlocksBeforeReadOnJobSchema pins the selection against the real
// databricks_job schema, so an upstream schema change that flips a flag is
// caught here rather than in a cluster.
func TestClearStaleBlocksBeforeReadOnJobSchema(t *testing.T) {
	r := configuredJobResource(t)
	job := r.TerraformResource
	ClearStaleBlocksBeforeRead(r, "provider_config")

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
		// Scalars the Read recomputes or never returns.
		"id", "name", "description", "format", "url", "timeout_seconds", "max_retries",
		"max_concurrent_runs", "edit_mode", "always_running", "control_run_state",
	} {
		if selected[name] {
			t.Errorf("did not expect %q to be cleared before the Read", name)
		}
	}

	wantEmptyObjectCleanup := map[string]bool{}
	for name := range selected {
		s := job.Schema[name]
		_, nestedBlock := s.Elem.(*schema.Resource)
		if s.Type == schema.TypeMap || (s.Type == schema.TypeList && s.MaxItems == 1 && nestedBlock) {
			wantEmptyObjectCleanup[name] = true
		}
	}
	gotEmptyObjectCleanup := map[string]bool{}
	cleaner := JobEmptyObjectCleaner(r)
	for _, name := range cleaner {
		if !strings.Contains(name, ".") && !strings.Contains(name, "[*]") {
			gotEmptyObjectCleanup[name] = true
		}
	}
	if len(gotEmptyObjectCleanup) != len(wantEmptyObjectCleanup) {
		t.Fatalf("empty object cleanup fields: got %v, want %v", gotEmptyObjectCleanup, wantEmptyObjectCleanup)
	}
	for name := range wantEmptyObjectCleanup {
		if !gotEmptyObjectCleanup[name] {
			t.Errorf("expected %q to be cleaned after Terraform conversion", name)
		}
	}

	cleanedPaths := map[string]bool{}
	for _, path := range cleaner {
		cleanedPaths[path] = true
	}
	for _, path := range r.TFListConversionPaths() {
		if isStaleProneBlock(schemaAtPath(job.Schema, path)) && !cleanedPaths[path] {
			t.Errorf("expected nested singleton path %q to be cleaned", path)
		}
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

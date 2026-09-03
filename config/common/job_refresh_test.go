package common

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	dbclient "github.com/databricks/databricks-sdk-go/client"
	dbconfig "github.com/databricks/databricks-sdk-go/config"
	sdkjobs "github.com/databricks/databricks-sdk-go/service/jobs"
	tfcommon "github.com/databricks/terraform-provider-databricks/common"
	tfjobs "github.com/databricks/terraform-provider-databricks/jobs"
	"github.com/databricks/terraform-provider-databricks/xpprovider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestClearStaleBlocksBeforeReadRefreshesJobSchedule(t *testing.T) {
	tests := map[string]struct {
		schedule  *tfjobs.CronSchedule
		wantCount int
		wantCron  string
	}{
		"omitted schedule clears prior state": {},
		"returned schedule populates state": {
			schedule: &tfjobs.CronSchedule{
				QuartzCronExpression: "0 0 6 * * ?",
				TimezoneID:           "UTC",
				PauseStatus:          "UNPAUSED",
			},
			wantCount: 1,
			wantCron:  "0 0 6 * * ?",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, sdkProvider, err := xpprovider.GetProvider(t.Context())
			if err != nil {
				t.Fatalf("GetProvider: %v", err)
			}
			job := sdkProvider.ResourcesMap["databricks_job"]
			ClearStaleBlocksBeforeRead(&config.Resource{TerraformResource: job}, "provider_config")

			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				if req.URL.Path == "/.well-known/databricks-config" {
					rw.WriteHeader(http.StatusNotFound)
					return
				}
				if req.Method != http.MethodGet || req.URL.RequestURI() != "/api/2.0/jobs/get?job_id=789" {
					t.Errorf("unexpected request: %s %s", req.Method, req.URL.RequestURI())
					rw.WriteHeader(http.StatusNotFound)
					return
				}
				if err := json.NewEncoder(rw).Encode(tfjobs.Job{
					JobID: 789,
					Settings: &tfjobs.JobSettings{
						ExistingClusterID: "abc",
						Name:              "job",
						Schedule:          tt.schedule,
					},
				}); err != nil {
					t.Errorf("encode response: %v", err)
				}
			}))
			defer server.Close()

			client, err := dbclient.New(&dbconfig.Config{Host: server.URL, Token: "test"})
			if err != nil {
				t.Fatalf("new Databricks client: %v", err)
			}
			meta := &tfcommon.DatabricksClient{DatabricksClient: client}
			meta.SetCachedWorkspaceID(12345)
			refreshJobSchedule(t, context.Background(), job, meta, tt.wantCount, tt.wantCron)
		})
	}
}

func TestClearStaleBlocksBeforeReadRefreshesMultiTaskJobSchedule(t *testing.T) {
	_, sdkProvider, err := xpprovider.GetProvider(t.Context())
	if err != nil {
		t.Fatalf("GetProvider: %v", err)
	}
	job := sdkProvider.ResourcesMap["databricks_job"]
	ClearStaleBlocksBeforeRead(&config.Resource{TerraformResource: job}, "provider_config")

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/.well-known/databricks-config" {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		if req.Method != http.MethodGet || req.URL.RequestURI() != "/api/2.2/jobs/get?job_id=789" {
			t.Errorf("unexpected request: %s %s", req.Method, req.URL.RequestURI())
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(rw).Encode(sdkjobs.Job{
			JobId: 789,
			Settings: &sdkjobs.JobSettings{
				Name:   "job",
				Format: sdkjobs.FormatMultiTask,
				Schedule: &sdkjobs.CronSchedule{
					QuartzCronExpression: "0 0 6 * * ?",
					TimezoneId:           "UTC",
					PauseStatus:          sdkjobs.PauseStatusUnpaused,
				},
			},
		}); err != nil {
			t.Errorf("encode response: %v", err)
		}
	}))
	defer server.Close()

	client, err := dbclient.New(&dbconfig.Config{Host: server.URL, Token: "test"})
	if err != nil {
		t.Fatalf("new Databricks client: %v", err)
	}
	meta := &tfcommon.DatabricksClient{DatabricksClient: client}
	meta.SetCachedWorkspaceID(12345)

	d := schema.TestResourceDataRaw(t, job.Schema, map[string]any{
		"format": "MULTI_TASK",
		"schedule": []any{map[string]any{
			"quartz_cron_expression": "0 0 12 * * ?",
			"timezone_id":            "UTC",
			"pause_status":           "UNPAUSED",
		}},
	})
	d.SetId("789")
	state, diags := job.RefreshWithoutUpgrade(context.Background(), d.State(), meta)
	if diags.HasError() {
		t.Fatalf("RefreshWithoutUpgrade: %v", diags)
	}
	refreshed, err := schema.InternalMap(job.Schema).Data(state, nil)
	if err != nil {
		t.Fatalf("state to ResourceData: %v", err)
	}
	if got := refreshed.Get("schedule.0.quartz_cron_expression").(string); got != "0 0 6 * * ?" {
		t.Errorf("schedule cron: got %q, want %q", got, "0 0 6 * * ?")
	}
}

func refreshJobSchedule(t *testing.T, ctx context.Context, job *schema.Resource, client any, wantCount int, wantCron string) {
	t.Helper()
	d := schema.TestResourceDataRaw(t, job.Schema, map[string]any{
		"existing_cluster_id": "abc",
		"schedule": []any{map[string]any{
			"quartz_cron_expression": "0 0 12 * * ?",
			"timezone_id":            "UTC",
			"pause_status":           "UNPAUSED",
		}},
	})
	d.SetId("789")

	state, diags := job.RefreshWithoutUpgrade(ctx, d.State(), client)
	if diags.HasError() {
		t.Fatalf("RefreshWithoutUpgrade: %v", diags)
	}
	if state == nil {
		t.Fatal("expected refreshed state")
	}

	refreshed, err := schema.InternalMap(job.Schema).Data(state, nil)
	if err != nil {
		t.Fatalf("state to ResourceData: %v", err)
	}
	if got := refreshed.Get("schedule.#").(int); got != wantCount {
		t.Fatalf("schedule count: got %d, want %d", got, wantCount)
	}
	if wantCount > 0 {
		if got := refreshed.Get("schedule.0.quartz_cron_expression").(string); got != wantCron {
			t.Errorf("schedule cron: got %q, want %q", got, wantCron)
		}
	}
}

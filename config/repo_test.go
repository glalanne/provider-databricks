package config

import (
	"context"
	"slices"
	"testing"

	upjetconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/resource"
	dbcommon "github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/repos"
	"github.com/databricks/terraform-provider-databricks/xpprovider"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	clusteralpha "github.com/glalanne/provider-databricks/apis/cluster/workspace/v1alpha1"
	clusterbeta "github.com/glalanne/provider-databricks/apis/cluster/workspace/v1beta1"
	namespacedalpha "github.com/glalanne/provider-databricks/apis/namespaced/workspace/v1alpha1"
	namespacedbeta "github.com/glalanne/provider-databricks/apis/namespaced/workspace/v1beta1"
)

func TestRepoSchemaAndResponseConversion(t *testing.T) {
	for _, scope := range []struct {
		name string
		get  func(context.Context, fwprovider.Provider, *schema.Provider, bool) (*upjetconfig.Provider, error)
	}{
		{name: "cluster", get: GetProvider},
		{name: "namespaced", get: GetProviderNamespaced},
	} {
		for _, generation := range []bool{false, true} {
			mode := "runtime"
			if generation {
				mode = "generation"
			}
			t.Run(scope.name+"/"+mode, func(t *testing.T) {
				ctx := context.Background()
				fw, sdk, err := xpprovider.GetProvider(ctx)
				if err != nil {
					t.Fatal(err)
				}
				p, err := scope.get(ctx, fw, sdk, generation)
				if err != nil {
					t.Fatal(err)
				}
				r := p.Resources["databricks_repo"]
				field := r.TerraformResource.Schema["commit_hash"]
				if !field.Optional || !field.Computed {
					t.Fatal("commit_hash must remain Optional+Computed")
				}
				if !slices.Contains(r.LateInitializer.IgnoredFields, "commit_hash") {
					t.Fatal("commit_hash must be excluded from late initialization")
				}
				if generation {
					return
				}
				d := r.TerraformResource.Data(&terraform.InstanceState{ID: "123"})
				err = dbcommon.StructToData(repos.ReposInformation{
					ID: 123, Url: "https://github.com/example/repo.git",
					HeadCommitID: "observed-hash",
				}, r.TerraformResource.Schema, d)
				if err != nil {
					t.Fatalf("Repo response conversion failed: %v", err)
				}
				if got := d.Get("commit_hash"); got != "observed-hash" {
					t.Fatalf("commit_hash = %v, want observed-hash", got)
				}
			})
		}
	}
}

func TestRepoCommitHashNotLateInitialized(t *testing.T) {
	for _, tc := range []struct {
		name string
		repo resource.Terraformed
	}{
		{name: "cluster/v1alpha1", repo: &clusteralpha.Repo{}},
		{name: "cluster/v1beta1", repo: &clusterbeta.Repo{}},
		{name: "namespaced/v1alpha1", repo: &namespacedalpha.Repo{}},
		{name: "namespaced/v1beta1", repo: &namespacedbeta.Repo{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.repo.LateInitialize([]byte(`{"commit_hash":"observed-hash","path":"/Repos/example/repo"}`))
			if err != nil {
				t.Fatal(err)
			}
			params, err := tc.repo.GetParameters()
			if err != nil {
				t.Fatal(err)
			}
			if _, ok := params["commit_hash"]; ok {
				t.Fatal("observed commit_hash was copied into spec")
			}
			if params["path"] != "/Repos/example/repo" {
				t.Fatal("unrelated fields must still late-initialize")
			}
			if err := tc.repo.SetObservation(map[string]any{"commit_hash": "observed-hash"}); err != nil {
				t.Fatal(err)
			}
			observed, err := tc.repo.GetObservation()
			if err != nil {
				t.Fatal(err)
			}
			if observed["commit_hash"] != "observed-hash" {
				t.Fatal("commit_hash must remain available in status")
			}
		})
	}
}

func TestRepoOmittedCommitHashDoesNotCauseDiff(t *testing.T) {
	r := repos.ResourceRepo().ToResource()
	state := &terraform.InstanceState{
		ID: "123",
		Attributes: map[string]string{
			"id": "123", "url": "https://github.com/example/repo.git",
			"commit_hash": "observed-hash", "tag": "",
		},
	}
	for _, configured := range []bool{false, true} {
		params := map[string]any{"url": state.Attributes["url"], "tag": "v0.1.0"}
		if configured {
			params["commit_hash"] = "stale-hash"
		}
		diff, err := schema.InternalMap(r.Schema).Diff(
			context.Background(), state, terraform.NewResourceConfigRaw(params), nil, nil, false,
		)
		if err != nil {
			t.Fatal(err)
		}
		if diff == nil || diff.Attributes["tag"] == nil {
			t.Fatal("legitimate tag changes must still produce a diff")
		}
		hashDiff := diff.Attributes["commit_hash"]
		if configured {
			if hashDiff == nil || hashDiff.Old != "observed-hash" || hashDiff.New != "stale-hash" {
				t.Fatalf("expected stale configured hash diff, got %#v", hashDiff)
			}
		} else if hashDiff != nil {
			t.Fatalf("omitted Optional+Computed hash caused a diff: %#v", hashDiff)
		}
	}
}

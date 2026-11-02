package config

import (
	"context"
	"encoding/json"
	"reflect"
	"slices"
	"strings"
	"testing"

	upjetconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/resource"
	sdkcatalog "github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/catalog"
	dbcommon "github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/xpprovider"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	clusteralpha "github.com/glalanne/provider-databricks/apis/cluster/unity/v1alpha1"
	clusterbeta "github.com/glalanne/provider-databricks/apis/cluster/unity/v1beta1"
	namespacedalpha "github.com/glalanne/provider-databricks/apis/namespaced/unity/v1alpha1"
	namespacedbeta "github.com/glalanne/provider-databricks/apis/namespaced/unity/v1beta1"
)

var credentialOutputFields = []string{
	"created_at", "created_by", "full_name", "metastore_id",
	"updated_at", "updated_by", "used_for_managed_storage",
	"aws_iam_role.external_id", "aws_iam_role.unity_catalog_iam_arn",
	"azure_managed_identity.credential_id",
	"databricks_gcp_service_account.credential_id",
	"databricks_gcp_service_account.email",
	"databricks_gcp_service_account.private_key_id",
}

func TestCredentialSchemaAndResponseConversion(t *testing.T) {
	for _, scope := range []struct {
		name string
		get  func(context.Context, fwprovider.Provider, *schema.Provider, bool) (*upjetconfig.Provider, error)
	}{
		{"cluster", GetProvider},
		{"namespaced", GetProviderNamespaced},
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
				r := p.Resources["databricks_credential"]
				if !reflect.DeepEqual(r.LateInitializer.IgnoredFields, credentialOutputFields) {
					t.Fatalf("unexpected late-init exclusions: %v", r.LateInitializer.IgnoredFields)
				}
				for _, path := range credentialOutputFields {
					m := r.TerraformResource.Schema
					parts := strings.Split(path, ".")
					for _, part := range parts[:len(parts)-1] {
						m = m[part].Elem.(*schema.Resource).Schema
					}
					field := m[parts[len(parts)-1]]
					if !field.Optional || !field.Computed {
						t.Fatalf("%s must remain Optional+Computed", path)
					}
				}
				for _, input := range []string{"owner", "isolation_mode", "provider_config"} {
					if slices.Contains(r.LateInitializer.IgnoredFields, input) {
						t.Fatalf("input %s must still late-initialize", input)
					}
				}
				if generation {
					return
				}
				d := r.TerraformResource.Data(&terraform.InstanceState{ID: "credential"})
				err = dbcommon.StructToData(sdkcatalog.CredentialInfo{
					Name: "credential", CreatedAt: 1000, UpdatedAt: 2000,
					AwsIamRole: &sdkcatalog.AwsIamRole{RoleArn: "arn:role", ExternalId: "external"},
					AzureManagedIdentity: &sdkcatalog.AzureManagedIdentity{
						AccessConnectorId: "/connector", CredentialId: "azure-id",
					},
					DatabricksGcpServiceAccount: &sdkcatalog.DatabricksGcpServiceAccount{
						CredentialId: "gcp-id", Email: "service@example.com", PrivateKeyId: "key-id",
					},
				}, r.TerraformResource.Schema, d)
				if err != nil {
					t.Fatalf("credential response conversion failed: %v", err)
				}
				if d.Get("updated_at") != 2000 {
					t.Fatalf("updated_at = %v, want 2000", d.Get("updated_at"))
				}
			})
		}
	}
}

func TestCredentialOutputLateInitialization(t *testing.T) {
	for _, version := range []struct {
		name string
		new  func() resource.Terraformed
		list bool
	}{
		{"cluster/v1alpha1", func() resource.Terraformed { return &clusteralpha.Credential{} }, true},
		{"cluster/v1beta1", func() resource.Terraformed { return &clusterbeta.Credential{} }, false},
		{"namespaced/v1alpha1", func() resource.Terraformed { return &namespacedalpha.Credential{} }, true},
		{"namespaced/v1beta1", func() resource.Terraformed { return &namespacedbeta.Credential{} }, false},
	} {
		for _, cloud := range []struct {
			block   string
			inputs  map[string]any
			outputs map[string]any
		}{
			{"aws_iam_role", map[string]any{"role_arn": "arn:role"},
				map[string]any{"external_id": "external", "unity_catalog_iam_arn": "arn:uc"}},
			{"azure_managed_identity", map[string]any{"access_connector_id": "/connector", "managed_identity_id": "/identity"},
				map[string]any{"credential_id": "azure-id"}},
			{"databricks_gcp_service_account", map[string]any{},
				map[string]any{"credential_id": "gcp-id", "email": "service@example.com", "private_key_id": "key-id"}},
		} {
			for _, existingBlock := range []bool{false, true} {
				name := "absent-block"
				if existingBlock {
					name = "configured-block"
				}
				t.Run(version.name+"/"+cloud.block+"/"+name, func(t *testing.T) {
					obj := version.new()
					block := func(m map[string]any) any {
						if version.list {
							return []any{m}
						}
						return m
					}
					if existingBlock {
						if err := obj.SetParameters(map[string]any{cloud.block: block(cloud.inputs)}); err != nil {
							t.Fatal(err)
						}
					}
					nested := map[string]any{}
					for k, v := range cloud.inputs {
						nested[k] = v
					}
					for k, v := range cloud.outputs {
						nested[k] = v
					}
					state := map[string]any{
						"created_at": 1000, "created_by": "creator", "updated_at": 2000, "updated_by": "editor",
						"full_name": "credential", "metastore_id": "metastore", "used_for_managed_storage": true,
						"owner": "owner", "isolation_mode": "ISOLATION_MODE_OPEN", cloud.block: block(nested),
					}
					data, err := json.Marshal(state)
					if err != nil {
						t.Fatal(err)
					}
					if _, err := obj.LateInitialize(data); err != nil {
						t.Fatal(err)
					}
					params, err := obj.GetParameters()
					if err != nil {
						t.Fatal(err)
					}
					for _, path := range credentialOutputFields {
						if !strings.Contains(path, ".") {
							if _, ok := params[path]; ok {
								t.Fatalf("%s was late-initialized", path)
							}
						}
					}
					got := map[string]any{}
					if raw, ok := params[cloud.block]; ok {
						if version.list {
							got = raw.([]any)[0].(map[string]any)
						} else {
							got = raw.(map[string]any)
						}
					}
					for key := range cloud.outputs {
						if _, ok := got[key]; ok {
							t.Fatalf("%s.%s was late-initialized", cloud.block, key)
						}
					}
					for key, want := range cloud.inputs {
						if got[key] != want {
							t.Fatalf("input %s = %v, want %v", key, got[key], want)
						}
					}
					if params["owner"] != "owner" || params["isolation_mode"] != "ISOLATION_MODE_OPEN" {
						t.Fatal("unrelated inputs must still late-initialize")
					}
					if err := obj.SetObservation(state); err != nil {
						t.Fatal(err)
					}
					obs, err := obj.GetObservation()
					if err != nil {
						t.Fatal(err)
					}
					normalized := map[string]any{}
					if err := json.Unmarshal(data, &normalized); err != nil {
						t.Fatal(err)
					}
					if !reflect.DeepEqual(obs, normalized) {
						t.Fatalf("observation did not round-trip: %v", obs)
					}
					if err := obj.SetParameters(state); err != nil {
						t.Fatal(err)
					}
					if _, err := obj.LateInitialize(data); err != nil {
						t.Fatal(err)
					}
					params, err = obj.GetParameters()
					if err != nil {
						t.Fatal(err)
					}
					if !reflect.DeepEqual(params, normalized) {
						t.Fatal("explicitly supplied output values must not be removed")
					}
				})
			}
		}
	}
}

func TestCredentialOmittedMetadataDoesNotCauseDiff(t *testing.T) {
	r := catalog.ResourceCredential().ToResource()
	state := &terraform.InstanceState{
		ID: "credential",
		Attributes: map[string]string{
			"id": "credential", "name": "credential", "purpose": "SERVICE",
			"updated_at": "2000", "updated_by": "editor",
			"aws_iam_role.#": "1", "aws_iam_role.0.role_arn": "arn:role",
		},
	}
	for _, configured := range []bool{false, true} {
		params := map[string]any{
			"name": "credential", "purpose": "SERVICE", "comment": "new comment",
			"aws_iam_role": []any{map[string]any{"role_arn": "arn:role"}},
		}
		if configured {
			params["updated_at"] = 1000
			params["updated_by"] = "old-editor"
		}
		diff, err := schema.InternalMap(r.Schema).Diff(
			context.Background(), state, terraform.NewResourceConfigRaw(params), nil, nil, false,
		)
		if err != nil {
			t.Fatal(err)
		}
		if diff == nil || diff.Attributes["comment"] == nil {
			t.Fatal("legitimate comment change must still produce a diff")
		}
		for _, field := range []string{"updated_at", "updated_by"} {
			if (diff.Attributes[field] != nil) != configured {
				t.Fatalf("configured=%v: unexpected %s diff: %#v", configured, field, diff.Attributes[field])
			}
		}
	}
}

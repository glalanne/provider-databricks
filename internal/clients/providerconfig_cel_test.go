package clients

import (
	"os"
	"path/filepath"
	"testing"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	structuralschema "k8s.io/apiextensions-apiserver/pkg/apiserver/schema"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/schema/cel"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/schema/cel/model"
	"k8s.io/apimachinery/pkg/util/version"
	"k8s.io/apimachinery/pkg/util/yaml"
	celconfig "k8s.io/apiserver/pkg/apis/cel"
	"k8s.io/apiserver/pkg/cel/environment"
)

// TestProviderConfigRulesCompileOnKubernetes131 compiles the ProviderConfig
// CRDs' CEL rules at CEL compatibility version 1.30, which a Kubernetes 1.31
// API server applies to new CRDs. Newer API servers, envtest's included, also
// accept a CEL reserved word such as namespace as a field name; 1.31 and
// kubectl-validate v0.0.4 accept only the escaped form, __namespace__.
func TestProviderConfigRulesCompileOnKubernetes131(t *testing.T) {
	envSet := environment.MustBaseEnvSet(version.MajorMinor(1, 30))
	files, err := filepath.Glob(filepath.Join("..", "..", "package", "crds", "databricks.*crossplane.io_*.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	rules := 0
	for _, file := range files {
		raw, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		var crd apiextensionsv1.CustomResourceDefinition
		if err := yaml.Unmarshal(raw, &crd); err != nil {
			t.Fatal(err)
		}
		for _, v := range crd.Spec.Versions {
			var props apiextensions.JSONSchemaProps
			if err := apiextensionsv1.Convert_v1_JSONSchemaProps_To_apiextensions_JSONSchemaProps(v.Schema.OpenAPIV3Schema, &props, nil); err != nil {
				t.Fatal(err)
			}
			s, err := structuralschema.NewStructural(&props)
			if err != nil {
				t.Fatal(err)
			}
			rules += compileRules(t, envSet, filepath.Base(file)+" "+v.Name, s, true)
		}
	}
	if rules == 0 {
		t.Fatal("found no CEL rules in the ProviderConfig CRDs")
	}
}

// compileRules compiles the rules on s and every schema below it, and returns
// how many it compiled.
func compileRules(t *testing.T, envSet *environment.EnvSet, path string, s *structuralschema.Structural, isRoot bool) int {
	t.Helper()
	n := len(s.XValidations)
	if n > 0 {
		results, err := cel.Compile(s, model.SchemaDeclType(s, isRoot), celconfig.PerCallLimit, envSet, cel.NewExpressionsEnvLoader())
		if err != nil {
			t.Fatalf("%s: %v", path, err)
		}
		for _, r := range results {
			if r.Error != nil {
				t.Errorf("%s: %v", path, r.Error)
			}
			if r.MessageExpressionError != nil {
				t.Errorf("%s: messageExpression: %v", path, r.MessageExpressionError)
			}
		}
	}
	for name, p := range s.Properties {
		n += compileRules(t, envSet, path+"."+name, &p, false)
	}
	if s.Items != nil {
		n += compileRules(t, envSet, path+"[*]", s.Items, false)
	}
	if s.AdditionalProperties != nil && s.AdditionalProperties.Structural != nil {
		n += compileRules(t, envSet, path+"[*]", s.AdditionalProperties.Structural, false)
	}
	return n
}

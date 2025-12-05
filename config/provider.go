/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	"context"
	_ "embed"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"
	uname "github.com/crossplane/upjet/v2/pkg/types/name"
	tfjson "github.com/hashicorp/terraform-json"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/glalanne/provider-databricks/config/cluster"
	"github.com/glalanne/provider-databricks/config/namespaced"

	"github.com/pkg/errors"
)

const (
	resourcePrefix = "databricks"
	modulePath     = "github.com/glalanne/provider-databricks"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

func getProviderSchema(s string) (*tfschema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &tfschema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(_ context.Context, fwProvider fwprovider.Provider, sdkProvider *tfschema.Provider, generationProvider bool) (*config.Provider, error) {

	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	pc := config.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		// ujconfig.WithShortName("databricks"),
		config.WithRootGroup("databricks.crossplane.io"),
		config.WithIncludeList(CLIReconciledResourceList()),
		config.WithTerraformPluginSDKIncludeList(TerraformPluginSDKResourceList()),
		config.WithTerraformPluginFrameworkIncludeList(TerraformPluginFrameworkResourceList()),
		config.WithDefaultResourceOptions(ResourceConfigurator()),
		config.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector(modulePath)}),
		config.WithFeaturesPackage("internal/features"),
		config.WithTerraformProvider(sdkProvider),
		config.WithTerraformPluginFrameworkProvider(fwProvider),
	)

	// Rename resources to make it more pleasing to the eye
	for _, r := range pc.Resources {

		parts := strings.Split(r.Name, "_")
		if len(parts) > 1 {
			r.ShortGroup = resourcePrefix
			r.Kind = uname.NewFromSnake(strings.Join(parts[1:], "_")).Camel
		}

	}

	// add custom config functions
	for _, configure := range cluster.ProviderConfiguration {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// GetProviderNamespaced returns provider configuration for namespace-scoped resources
func GetProviderNamespaced(_ context.Context, fwProvider fwprovider.Provider, sdkProvider *tfschema.Provider, generationProvider bool) (*config.Provider, error) {
	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	pc := config.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		// config.WithShortName("databricks"),
		config.WithRootGroup("databricks.m.crossplane.io"),
		config.WithIncludeList(CLIReconciledResourceList()),
		config.WithTerraformPluginSDKIncludeList(TerraformPluginSDKResourceList()),
		config.WithTerraformPluginFrameworkIncludeList(TerraformPluginFrameworkResourceList()),
		config.WithDefaultResourceOptions(ResourceConfigurator()),
		config.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector(modulePath)}),
		config.WithFeaturesPackage("internal/features"),
		config.WithTerraformProvider(sdkProvider),
		config.WithTerraformPluginFrameworkProvider(fwProvider),
	)

	registerTerraformConversions(pc)

	// Rename resources to make it more pleasing to the eye
	for _, r := range pc.Resources {

		parts := strings.Split(r.Name, "_")
		if len(parts) > 1 {
			r.ShortGroup = resourcePrefix
			r.Kind = uname.NewFromSnake(strings.Join(parts[1:], "_")).Camel
		}
	}

	// add custom config functions
	for _, configure := range namespaced.ProviderConfiguration {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

func registerTerraformConversions(pc *config.Provider) {
	for name, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}

		r.TerraformConversions = []config.TerraformConversion{
			config.NewTFSingletonConversion(),
		}
		pc.Resources[name] = r
	}
}

// CLIReconciledResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the TF CLI based architecture.
func CLIReconciledResourceList() []string {
	l := make([]string, len(CLIReconciledExternalNameConfigs))
	i := 0
	for name := range CLIReconciledExternalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

// TerraformPluginSDKResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the no-fork architecture.
func TerraformPluginSDKResourceList() []string {
	l := make([]string, len(TerraformPluginSDKExternalNameConfigs))
	i := 0
	for name := range TerraformPluginSDKExternalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

func TerraformPluginFrameworkResourceList() []string {
	l := make([]string, len(TerraformPluginFrameworkExternalNameConfigs))
	i := 0
	for name := range TerraformPluginFrameworkExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

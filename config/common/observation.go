package common

import (
	"sort"
	"strings"

	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

// EmptyObjectCleaner converts selected empty observed objects to nil so
// generated observation setters clear values that disappeared remotely.
type EmptyObjectCleaner []string

// EmptyObjectCleanerForResource returns every singleton block and map path,
// including nested paths, whose empty value can merge with stale observation.
func EmptyObjectCleanerForResource(r *config.Resource) EmptyObjectCleaner {
	if r == nil || r.TerraformResource == nil {
		return nil
	}

	paths := map[string]struct{}{}
	for _, path := range r.TFListConversionPaths() {
		if fieldSchema := schemaAtPath(r.TerraformResource.Schema, path); isStaleProneBlock(fieldSchema) {
			paths[embeddedObjectPath(r.TerraformResource.Schema, path)] = struct{}{}
		}
	}
	addMapPaths(r.TerraformResource.Schema, "", paths)

	cleaner := make(EmptyObjectCleaner, 0, len(paths))
	for path := range paths {
		cleaner = append(cleaner, path)
	}
	sort.Strings(cleaner)
	return cleaner
}

// Convert implements config.TerraformConversion.
func (c EmptyObjectCleaner) Convert(params map[string]any, _ *config.Resource, mode config.Mode) (map[string]any, error) {
	if mode != config.FromTerraform {
		return params, nil
	}

	paved := fieldpath.Pave(params)
	for _, path := range c {
		expanded, err := paved.ExpandWildcards(path)
		if fieldpath.IsNotFound(err) {
			continue
		}
		if err != nil {
			return nil, errors.Wrapf(err, "cannot expand empty object path %q", path)
		}
		for _, concretePath := range expanded {
			value, err := paved.GetValue(concretePath)
			if err != nil {
				return nil, errors.Wrapf(err, "cannot read empty object path %q", concretePath)
			}
			if object, ok := value.(map[string]any); ok && len(object) == 0 {
				if err := paved.SetValue(concretePath, nil); err != nil {
					return nil, errors.Wrapf(err, "cannot clear empty object path %q", concretePath)
				}
			}
		}
	}
	return params, nil
}

func addMapPaths(fields map[string]*schema.Schema, prefix string, paths map[string]struct{}) {
	for name, fieldSchema := range fields {
		path := strings.TrimPrefix(prefix+"."+name, ".")
		if fieldSchema.Type == schema.TypeMap && isStaleProneBlock(fieldSchema) {
			paths[path] = struct{}{}
		}

		nested, ok := fieldSchema.Elem.(*schema.Resource)
		if !ok {
			continue
		}
		if (fieldSchema.Type == schema.TypeList || fieldSchema.Type == schema.TypeSet) && fieldSchema.MaxItems != 1 {
			path += "[*]"
		}
		addMapPaths(nested.Schema, path, paths)
	}
}

func schemaAtPath(fields map[string]*schema.Schema, path string) *schema.Schema {
	segments := strings.Split(path, ".")
	for i, segment := range segments {
		name := strings.TrimSuffix(strings.TrimSuffix(segment, "[*]"), "[0]")
		fieldSchema := fields[name]
		if fieldSchema == nil || i == len(segments)-1 {
			return fieldSchema
		}
		nested, ok := fieldSchema.Elem.(*schema.Resource)
		if !ok {
			return nil
		}
		fields = nested.Schema
	}
	return nil
}

func embeddedObjectPath(fields map[string]*schema.Schema, path string) string {
	segments := strings.Split(path, ".")
	for i, segment := range segments {
		name := strings.TrimSuffix(strings.TrimSuffix(segment, "[*]"), "[0]")
		fieldSchema := fields[name]
		if fieldSchema == nil {
			break
		}
		if fieldSchema.MaxItems == 1 {
			segments[i] = name
		}
		nested, ok := fieldSchema.Elem.(*schema.Resource)
		if !ok {
			break
		}
		fields = nested.Schema
	}
	return strings.Join(segments, ".")
}

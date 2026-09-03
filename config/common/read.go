// Package common holds configuration helpers shared by the cluster-scoped and
// namespaced resource configurations.
package common

import (
	"context"
	"sync"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type readCleaner struct {
	mu     sync.Mutex
	fields map[string]struct{}
}

func (c *readCleaner) add(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.fields[name] = struct{}{}
}

func (c *readCleaner) names() []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	names := make([]string, 0, len(c.fields))
	for n := range c.fields {
		names = append(names, n)
	}
	return names
}

var (
	cleanersMu sync.Mutex
	// The cluster-scoped and namespaced configurations share the same
	// *schema.Resource pointers, so the wrapper must only be installed once.
	cleaners = map[*schema.Resource]*readCleaner{}
)

// ClearStaleBlocksBeforeRead resets every optional nested block and map of the
// resource to its empty value before the resource's Terraform Read runs.
//
// The Databricks API omits blocks that are no longer configured instead of
// returning them empty, and the upstream Read only sets what it receives. A
// removed block therefore survives in the Terraform state, surfaces in
// status.atProvider and is copied back into spec.forProvider by late
// initialization. Clearing first lets the upstream Read repopulate only the
// blocks that actually exist remotely.
//
// StructToData normally ignores optional values returned for fields that are
// empty in ResourceData. Marking this refresh data as new bypasses that gate so
// the Read can repopulate a cleared block. Only collection fields are cleared;
// scalar fields used by provider Read implementations remain intact.
//
// Blocks the API never echoes back are skipped automatically when they are
// computed or carry a sensitive field. Pass the names of any remaining
// write-only blocks in excluded so they keep their state.
func ClearStaleBlocksBeforeRead(r *config.Resource, excluded ...string) {
	if r == nil || r.TerraformResource == nil {
		return
	}
	skip := make(map[string]struct{}, len(excluded))
	for _, e := range excluded {
		skip[e] = struct{}{}
	}

	var fields []string
	for name, s := range r.TerraformResource.Schema {
		if _, ok := skip[name]; ok {
			continue
		}
		if isStaleProneBlock(s) {
			fields = append(fields, name)
		}
	}
	ClearFieldsBeforeRead(r, fields...)
}

// isStaleProneBlock reports whether a field holds a collection that the Read
// leaves untouched once the API stops returning it.
func isStaleProneBlock(s *schema.Schema) bool {
	if !s.Optional || s.Required || s.Computed || s.Sensitive || emptyValue(s) == nil {
		return false
	}
	elem, ok := s.Elem.(*schema.Resource)
	if !ok {
		return s.Type == schema.TypeList || s.Type == schema.TypeSet || s.Type == schema.TypeMap
	}
	// A block holding a value the API never returns must keep its state.
	return !containsSensitive(elem, map[*schema.Resource]struct{}{})
}

func containsSensitive(r *schema.Resource, seen map[*schema.Resource]struct{}) bool {
	if r == nil {
		return false
	}
	if _, ok := seen[r]; ok {
		return false
	}
	seen[r] = struct{}{}
	for _, s := range r.Schema {
		if s.Sensitive {
			return true
		}
		if elem, ok := s.Elem.(*schema.Resource); ok && containsSensitive(elem, seen) {
			return true
		}
	}
	return false
}

// ClearFieldsBeforeRead resets the given top-level list, set or map fields to
// their empty value before the resource's Terraform Read runs. Prefer
// ClearStaleBlocksBeforeRead unless a single field has to be targeted.
func ClearFieldsBeforeRead(r *config.Resource, fields ...string) {
	if r == nil || r.TerraformResource == nil || r.TerraformResource.ReadContext == nil {
		// No Terraform provider is wired in during API generation.
		return
	}

	cleanersMu.Lock()
	defer cleanersMu.Unlock()

	tr := r.TerraformResource
	c, wrapped := cleaners[tr]
	if !wrapped {
		c = &readCleaner{fields: map[string]struct{}{}}
		cleaners[tr] = c
		wrapReadCleaner(tr, c)
	}

	for _, f := range fields {
		s, ok := tr.Schema[f]
		if !ok || emptyValue(s) == nil {
			continue
		}
		c.add(f)
	}
}

func wrapReadCleaner(tr *schema.Resource, cleaner *readCleaner) {
	read := tr.ReadContext
	tr.ReadContext = func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
		names := cleaner.names()
		selected := make(map[string]struct{}, len(names))
		for _, name := range names {
			selected[name] = struct{}{}
			if err := d.Set(name, emptyValue(tr.Schema[name])); err != nil {
				return diag.FromErr(err)
			}
		}
		if len(selected) == 0 {
			return read(ctx, d, meta)
		}

		unconfigured := unconfiguredFields(d, tr.Schema, selected)
		// StructToData uses this marker only to decide whether values returned
		// by the API may populate fields absent from ResourceData.
		d.MarkNewResource()
		diags := read(ctx, d, meta)
		for name := range unconfigured {
			if err := d.Set(name, nil); err != nil {
				return append(diags, diag.FromErr(err)...)
			}
		}
		return diags
	}
}

// unconfiguredFields captures fields StructToData would normally suppress on
// an existing resource. Restoring them after Read confines MarkNewResource to
// the selected fields.
func unconfiguredFields(d *schema.ResourceData, schemas map[string]*schema.Schema, selected map[string]struct{}) map[string]struct{} {
	fields := map[string]struct{}{}
	for name, s := range schemas {
		if _, ok := selected[name]; ok || s.Computed || s.Type == schema.TypeBool {
			continue
		}
		if _, configured := d.GetOk(name); !configured {
			fields[name] = struct{}{}
		}
	}
	return fields
}

func emptyValue(s *schema.Schema) any {
	if s == nil {
		return nil
	}
	switch s.Type {
	case schema.TypeList, schema.TypeSet:
		return []any{}
	case schema.TypeMap:
		return map[string]any{}
	case schema.TypeInvalid, schema.TypeBool, schema.TypeInt, schema.TypeFloat, schema.TypeString:
		return nil
	}
	return nil
}

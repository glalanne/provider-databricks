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

// UseAuthoritativeRead makes the API response authoritative for every
// optional field that is safe to clear before the resource's Terraform Read.
//
// The Databricks API omits fields that are no longer configured instead of
// returning them empty, and the upstream Read only sets what it receives. A
// removed field therefore survives in the Terraform state, surfaces in
// status.atProvider and is copied back into spec.forProvider by late
// initialization. Clearing first lets the upstream Read repopulate only the
// fields that actually exist remotely. Marking the temporary ResourceData as
// new makes StructToData populate fields returned by the API even when they
// are absent from config.
//
// Fields the API never echoes back are skipped automatically when they are
// computed or carry a sensitive field. Pass the names of any remaining
// write-only fields in excluded so they keep their state.
func UseAuthoritativeRead(r *config.Resource, excluded ...string) {
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
		if isStaleProneField(s) {
			fields = append(fields, name)
		}
	}
	ClearFieldsBeforeRead(r, fields...)
}

// isStaleProneField reports whether a field can be cleared before Read and
// reconstructed from the API response.
func isStaleProneField(s *schema.Schema) bool {
	if !s.Optional || s.Required || s.Computed || s.Sensitive || emptyValue(s) == nil {
		return false
	}
	elem, ok := s.Elem.(*schema.Resource)
	if !ok {
		return true
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

// ClearFieldsBeforeRead resets the given top-level fields to
// their empty value if the resource's Terraform Read does not populate them.
// Prefer UseAuthoritativeRead unless a single field has to be targeted.
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

		read := tr.ReadContext
		tr.ReadContext = func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
			names := c.names()
			for _, name := range names {
				if err := d.Set(name, emptyValue(tr.Schema[name])); err != nil {
					return diag.FromErr(err)
				}
			}
			d.MarkNewResource()
			return read(ctx, d, meta)
		}
	}

	for _, f := range fields {
		s, ok := tr.Schema[f]
		if !ok || emptyValue(s) == nil {
			continue
		}
		c.add(f)
	}
}

func emptyValue(s *schema.Schema) any {
	if s == nil {
		return nil
	}
	switch s.Type {
	case schema.TypeBool:
		return false
	case schema.TypeInt:
		return 0
	case schema.TypeFloat:
		return float64(0)
	case schema.TypeString:
		return ""
	case schema.TypeList, schema.TypeSet:
		return []any{}
	case schema.TypeMap:
		return map[string]any{}
	case schema.TypeInvalid:
		return nil
	}
	return nil
}

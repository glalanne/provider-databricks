package common

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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

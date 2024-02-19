// Copyright 2024 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package typeset

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Normalizer takes a value and converts it to a stable representation.
type Normalizer interface {
	Normalize(v any) any
}

// ElemField defines a set element field. ElemField must always implement
// Normalizer to ensure that each set element hash is stable.
type ElemField interface {
	Normalizer

	// Schema returns the schema for the element field.
	Schema() *schema.Schema
}

// NotNormalized returns an ElemField for a element field schema which already
// encodes stable values and thus does need to implement Normalizer.
//
// Note that any field which implements a DiffSuppressFunc will be rejected if
// it is declared as NotNormalized.
func NotNormalized(s *schema.Schema) ElemField {
	return notNormalized{s}
}

type notNormalized struct {
	schema *schema.Schema
}

func (f notNormalized) Schema() *schema.Schema {
	return f.schema
}

func (f notNormalized) Normalize(v any) any {
	return v
}

// Set defines the parameters of a TypeSet field in a Terraform schema. Set must
// be used to build any TypeSet schemas which contains complex objects to ensure
// that diff-ing is implemented correctly (TL;DR TypeSet has tons of bugs and
// requires very specific configuration).
type Set struct {
	Required bool
	MinItems int

	// ElemFields defines the element field schema, where the keys are the name
	// of the element fields and the values are the schemas for said field.
	ElemFields map[string]ElemField
}

// Schema returns the Terraform schema of the Set.
func (s Set) Schema() *schema.Schema {
	elemResource := s.elemResource()
	return &schema.Schema{
		Type:     schema.TypeSet,
		Required: s.Required,
		Optional: !s.Required,
		MinItems: s.MinItems,
		Elem:     elemResource,
		Set:      s.setFunc(elemResource),
	}
}

func (s Set) elemResource() *schema.Resource {
	elemSchema := make(map[string]*schema.Schema)
	for name, f := range s.ElemFields {
		if err := validateField(f); err != nil {
			panic(fmt.Errorf("invalid field %q: %v", name, err))
		}
		elemSchema[name] = f.Schema()
	}
	return &schema.Resource{
		Schema: elemSchema,
	}
}

func (s Set) setFunc(resource *schema.Resource) schema.SchemaSetFunc {
	hasher := schema.HashResource(resource)

	return func(elem any) int {
		elemMap := elem.(map[string]any)
		normalized := make(map[string]any)

		for k, v := range elemMap {
			normalized[k] = s.normalize(k, v)
		}

		return hasher(normalized)
	}
}

func (s Set) normalize(k string, v any) any {
	f, ok := s.ElemFields[k]
	if !ok {
		return v
	}
	return f.Normalize(v)
}

func validateField(f ElemField) error {
	s := f.Schema()
	switch f.(type) {
	case notNormalized:
		if s.DiffSuppressFunc != nil {
			return errors.New("cannot DiffSuppressFunc without implementing Normalize")
		}
	default:
		if s.DiffSuppressFunc == nil {
			return errors.New("must define DiffSuppressFunc when implementing Normalize")
		}
	}
	return nil
}

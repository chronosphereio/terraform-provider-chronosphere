// Copyright 2023 Chronosphere Inc.
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

package chronosphere

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/go-cty/cty"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema/intschematag"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/shared/pkg/container/set"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
)

// indexPath is used in path when indexing into a slice / map.
const indexPath = "[]"

// dummyRef is used in dry run validation to populate fields that reference other entities
// that don't exist yet.
var dummyRef = tfid.Slug("dummy_value")

// setUnknownReferences sets entity reference fields (those having the type tfid.ID) that
// a) do not currently have a value and
// b) have a non-nil value in the raw config
// This occurs when two related objects are being created at the same time, so that the referred to object does not
// have a slug at the time plan is executed.
// Populating this dummy value allows entities to pass validations that their reference field is populated.
// A database level validation will still trigger, but these are errors ignored by the TF provider.
func setUnknownReferences(v any, rawCfg cty.Value, skip []string) {
	rv := reflect.ValueOf(v)

	setUnknownReferencesParams{
		rawCfg: rawCfg,
		skip:   set.New(skip...),
	}.set(rv, nil /* path */)
}

type setUnknownReferencesParams struct {
	rawCfg cty.Value
	skip   set.Set[string]
}

func (p setUnknownReferencesParams) set(v reflect.Value, path []string) {
	if id, ok := v.Interface().(tfid.ID); ok {
		p.setID(v, id, path)
		return
	}

	switch v.Type().Kind() {
	case reflect.Bool, reflect.String, reflect.Float64, reflect.Int64:
		// no-op
	case reflect.Pointer:
		if v.IsNil() {
			// Create a dummy empty value, and recurse into it to verify
			// there's no nested unsupported tfid fields.
			// v is (*T)(nil), we need to create new(T), hence .Elem().
			v = reflect.New(v.Type().Elem())
		}
		p.set(v.Elem(), path)
	case reflect.Slice:
		// Create a dummy slice with an empty value, and recurse into it to verify
		// there's no nested unsupported tfid fields.
		if v.Len() == 0 {
			v = reflect.MakeSlice(v.Type(), 1, 1)
		}
		newPath := append(path, indexPath)
		for i := 0; i < v.Len(); i++ {
			p.set(v.Index(i), newPath)
		}
	case reflect.Map:
		if v.Len() == 0 {
			// Create a new map element, and recurse into that to verify
			// there's no nested unsupported tfid fields.
			emptyElem := reflect.New(v.Type().Elem())
			newPath := append(path, indexPath)
			p.set(emptyElem, newPath)
			return
		}

		newPath := append(path, indexPath)
		for _, k := range v.MapKeys() {
			p.set(v.MapIndex(k), newPath)
		}
	case reflect.Struct:
		for i := 0; i < v.Type().NumField(); i++ {
			tag := intschematag.Unmarshal(v.Type().Field(i))
			newPath := append(path, tag.TFName)
			p.set(v.Field(i), newPath)
		}
	default:
		panic(fmt.Errorf("unsupported type %v", v.Type()))
	}
}

func (p setUnknownReferencesParams) setID(v reflect.Value, id tfid.ID, path []string) {
	idPath := strings.Join(path, ".")
	if p.skip.Has(idPath) {
		// Explicitly allow-listed to be skipped, expected for nested tfid fields, see below.
		return
	}

	if len(path) > 1 {
		// We don't support nested references for simplicity, as they're only used in a single
		// resource, and add complexity:
		// * Nested structs are actually represented as slices in the raw config.
		// * Sets are represented as slices but with different order in the intschema slice
		//    compared to the raw config.
		// * Slice and map lookups need to use .Index vs .GetAttr.
		panic(fmt.Sprintf("setUnknownReferences found unsupported tfid in a slice/map at %v", idPath))
	}

	// If the id value is empty, but the underlying config is not null, then the reference is to
	// an object that may not have been created yet, but will likely be created as part of the
	// same apply operation, so we set the field to a dummy value.
	hasConfig := !p.rawCfg.GetAttr(path[0]).IsNull()
	if id == (tfid.ID{}) && hasConfig {
		v.Set(reflect.ValueOf(dummyRef))
	}
}

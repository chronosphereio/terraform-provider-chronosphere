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

package hclmarshal

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema/intschematag"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"go.uber.org/multierr"
)

var typeTFID = reflect.TypeOf(tfid.ID{})

// MarshalIntSchema marshals an intschema struct using reflection.
func MarshalIntSchema(v any, b *Block) error {
	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Pointer || rv.Elem().Kind() != reflect.Struct {
		panic(errors.New("MarshalIntSchema must be called with pointer to an intschema struct"))
	}

	return marshalStruct(rv.Elem(), b)
}

func marshalStruct(rv reflect.Value, b *Block) error {
	var errs error

	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		t := intschematag.Unmarshal(rt.Field(i))
		if t.TFName == intschematag.InternalFieldName {
			continue
		}

		errs = multierr.Append(errs, marshalField(t, rv.Field(i), b))
	}

	return errs
}

func marshalField(t intschematag.Tag, rv reflect.Value, b *Block) error {
	rt := rv.Type()

	if rt == typeTFID {
		b.AddRef(t.TFName, rv.Interface().(tfid.ID))
		return nil
	}

	if t.File {
		marshalFile(t, rv.Interface(), b)
		return nil
	}

	if isPrimitive(rv.Type()) {
		marshalValue(t, rv.Interface(), b)
		return nil
	}

	switch rt.Kind() {
	case reflect.Map:
		if rv.Len() == 0 {
			// Optional empty containers should be skipped, and required containers should not be empty.
			return nil
		}

		if isPrimitive(rt.Elem()) {
			marshalValue(t, rv.Interface(), b)
			return nil
		}

		panic(fmt.Errorf("map %q has an unsupported type", t.TFName))
	case reflect.Slice:
		if rv.Len() == 0 {
			// Optional empty containers should be skipped, and required containers should not be empty.
			return nil
		}

		if isPrimitive(rt.Elem()) {
			marshalValue(t, rv.Interface(), b)
			return nil
		}

		if rt.Elem() == typeTFID {
			b.AddRefs(t.TFName, rv.Interface().([]tfid.ID))
			return nil
		}

		// Slice of complex objects uses a list of sequential blocks.
		var errs error
		for i := 0; i < rv.Len(); i++ {
			b.formatNewLines(true /* nextFieldComplex */)
			nested := b.AddBlock(t.TFName)
			errs = multierr.Append(errs, marshalStruct(rv.Index(i), nested))
		}
		return errs
	case reflect.Struct:
		b.formatNewLines(true /* nextFieldComplex */)
		nested := b.AddBlock(t.TFName)
		return marshalStruct(rv, nested)
	case reflect.Pointer:
		if !t.Optional {
			// We should never get here, since only optional fields use pointers.
			panic(fmt.Errorf("field %q is required, but is using a pointer", t.TFName))
		}

		if rv.IsZero() {
			// No need to marshal anything
			return nil
		}

		return marshalField(t, rv.Elem(), b)
	default:
		panic(fmt.Errorf("field %q has unsupported type %v", t.TFName, rv.Type().Name()))
	}
}

func marshalFile(t intschematag.Tag, v any, b *Block) {
	fileName, ok := v.(string)
	if !ok {
		panic(fmt.Errorf("file field %q is not a string", t.TFName))
	}

	if fileName == "" {
		return
	}

	b.AddFuncCall(t.TFName, "file", fileName)
}

func marshalValue(t intschematag.Tag, v any, b *Block) {
	// If !Optional, then the computed field is considered read-only and only settable
	// by the provider, so never marshal a value.
	if !t.Optional && t.Computed {
		return
	}

	if t.Optional {
		// Optionals without a default are only marshalled if they
		// are not the 0 value for the type, which AddOptional handles.
		if t.Default == "" {
			b.AddOptional(t.TFName, v)
			return
		}

		// Optional fields with a default are only marshalled if they
		// are not the configured default value for the type.
		if fmt.Sprint(v) == t.Default {
			return
		}
	}

	// required fields are always marshalled, as are non-default optionals.
	b.Add(t.TFName, v)
}

func isPrimitive(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Bool, reflect.Float64, reflect.Int64, reflect.String:
		return true
	default:
		return false
	}
}

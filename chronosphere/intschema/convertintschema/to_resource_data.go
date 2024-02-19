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

package convertintschema

import (
	"fmt"
	"reflect"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema/intschematag"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema/overridecreate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ToResourceData unloads obj into d, where obj is a generated intschema struct.
func ToResourceData(obj any, d *schema.ResourceData) diag.Diagnostics {
	toSet := unloadStruct(reflect.ValueOf(obj).Elem(), func(tag intschematag.Tag, field reflect.Value) bool {
		if tag.Optional {
			_, ok := d.GetOk(tag.TFName)
			// For optional fields, only call Set if the local value is not
			// empty or the server value is not empty.
			return ok || !isEmpty(field)
		}
		return true
	})
	var errs diag.Diagnostics
	for k, v := range toSet {
		if err := d.Set(k, v); err != nil {
			errs = append(errs, diag.FromErr(err)...)
		}
	}
	return errs
}

func unloadStruct(
	v reflect.Value,
	shouldSet func(tag intschematag.Tag, field reflect.Value) bool,
) map[string]any {
	out := make(map[string]any)
	for i := 0; i < v.NumField(); i++ {
		fieldType := v.Type().Field(i)
		field := v.Field(i)
		tag := intschematag.Unmarshal(v.Type().Field(i))

		if tag.IgnoreResourceData() ||
			// NB: Ignore server value for override_create fields. Server will always return an empty value because
			// override_create is a Terraform concept and will only be present on the manually-defined TF resource.
			// We don't want to overwrite the value defined in the resource to empty.
			tag.TFName == overridecreate.Field {
			if !isEmpty(field) {
				panic(fmt.Sprintf(
					"cannot set field %s when calling ToResourceData",
					fieldType.Name))
			}
			continue
		}

		if shouldSet != nil && !shouldSet(tag, field) {
			continue
		}

		data := unloadValue(field)
		if tag.ListEncodedObject {
			if data != nil {
				data = []any{data}
			}
		}
		out[tag.TFName] = data
	}
	return out
}

func unloadSlice(v reflect.Value) any {
	out := make([]any, 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		out = append(out, unloadValue(v.Index(i)))
	}
	return out
}

func unloadMap(v reflect.Value) any {
	out := make(map[string]any)
	for _, k := range v.MapKeys() {
		out[k.Interface().(string)] = v.MapIndex(k).Interface()
	}
	return out
}

func unloadValue(v reflect.Value) any {
	switch v.Type().Kind() {
	case reflect.Pointer:
		if v.IsNil() {
			return nil
		}
		return unloadValue(v.Elem())
	case reflect.Struct:
		if id, ok := v.Interface().(tfid.ID); ok {
			return id.Slug()
		}
		return unloadStruct(v, nil /* set all */)
	case reflect.Slice:
		return unloadSlice(v)
	case reflect.Map:
		return unloadMap(v)
	case reflect.Bool, reflect.String, reflect.Float64, reflect.Int64:
		return v.Interface()
	default:
		panic(fmt.Sprintf("unhandled go type: %s", v.Type().Kind()))
	}
}

func isEmpty(v reflect.Value) bool {
	if v.IsZero() {
		return true
	}
	switch v.Kind() {
	case reflect.Map, reflect.Slice, reflect.Array:
		return v.Len() == 0
	}
	return false
}

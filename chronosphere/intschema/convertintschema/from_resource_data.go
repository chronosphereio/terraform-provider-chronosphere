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

package convertintschema

import (
	"fmt"
	"reflect"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema/intschematag"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceGetter is a subset of the read-only interface to read resource data
// from schema.ResourceData and schema.ResourceDiff.
type ResourceGetter interface {
	Get(key string) any
	Id() string
}

// FromResourceData unloads d into outObj, where outObj is a generated intschema
// struct and objSchema is its Terraform schema.
func FromResourceData(
	objSchema map[string]*schema.Schema, d ResourceGetter, outObj any,
) error {
	outVal := reflect.ValueOf(outObj).Elem()
	outVal.FieldByName(intschematag.StateIDField).SetString(d.Id())
	return loadObject(objSchema, dataAsMap(objSchema, d), outVal)
}

func dataAsMap(objSchema map[string]*schema.Schema, d ResourceGetter) map[string]any {
	m := make(map[string]any)
	for k := range objSchema {
		m[k] = d.Get(k)
	}
	return m
}

func loadObject(
	objSchema map[string]*schema.Schema, data map[string]any, outVal reflect.Value,
) error {
	for i := 0; i < outVal.NumField(); i++ {
		tag := intschematag.Unmarshal(outVal.Type().Field(i))
		if tag.IgnoreResourceData() {
			continue
		}
		s, ok := objSchema[tag.TFName]
		if !ok {
			return fmt.Errorf("no field schema for %q", tag.TFName)
		}
		f := outVal.Field(i)
		if err := loadSchema(s, data[tag.TFName], f); err != nil {
			return fmt.Errorf("load %q schema: %s", tag.TFName, err)
		}
	}
	return nil
}

func loadSchema(s *schema.Schema, data any, outVal reflect.Value) error {
	switch s.Type {
	case schema.TypeBool:
		outVal.SetBool(data.(bool))
	case schema.TypeString:
		s := data.(string)
		if _, ok := outVal.Interface().(tfid.ID); ok {
			outVal.Set(reflect.ValueOf(tfid.Slug(s)))
		} else {
			outVal.SetString(s)
		}
	case schema.TypeFloat:
		outVal.SetFloat(data.(float64))
	case schema.TypeInt:
		outVal.SetInt(int64(data.(int)))
	case schema.TypeSet:
		l := data.(*schema.Set).List()
		if err := loadSlice(s, outVal, l); err != nil {
			return err
		}
	case schema.TypeList:
		l := data.([]any)
		if tfschema.IsListEncodedObject(s) {
			if err := loadListEncodedObject(s, outVal, l); err != nil {
				return err
			}
		} else {
			if err := loadSlice(s, outVal, l); err != nil {
				return err
			}
		}
	case schema.TypeMap:
		if err := loadMap(s, outVal, data.(map[string]any)); err != nil {
			return err
		}
	default:
		panic(fmt.Sprintf("unhandled terraform type: %s", s.Type))
	}
	return nil
}

func loadSlice(s *schema.Schema, field reflect.Value, data []any) error {
	if len(data) == 0 {
		return nil
	}
	slice := reflect.MakeSlice(field.Type(), 0, len(data))
	for i := 0; i < len(data); i++ {
		v := reflect.New(field.Type().Elem())
		if err := loadElem(s, v, data[i]); err != nil {
			return err
		}
		slice = reflect.Append(slice, v.Elem())
	}
	field.Set(slice)
	return nil
}

func loadMap(s *schema.Schema, field reflect.Value, data map[string]any) error {
	if len(data) == 0 {
		return nil
	}
	m := reflect.MakeMap(field.Type())
	for k := range data {
		v := reflect.New(field.Type().Elem())
		if err := loadElem(s, v, data[k]); err != nil {
			return err
		}
		m.SetMapIndex(reflect.ValueOf(k), v.Elem())
	}
	field.Set(m)
	return nil
}

func loadListEncodedObject(s *schema.Schema, v reflect.Value, data []any) error {
	if len(data) == 0 {
		return nil
	}
	var ptr reflect.Value
	if v.Kind() == reflect.Pointer {
		// If the field is a pointer to a struct it will be nil initially, thus
		// we must manually initialize the underlying value.
		v.Set(reflect.New(v.Type().Elem()))
		ptr = v
	} else {
		ptr = v.Addr()
	}
	return loadElem(s, ptr, data[0])
}

func loadElem(s *schema.Schema, v reflect.Value, data any) error {
	if data == nil {
		return nil
	}
	switch t := s.Elem.(type) {
	case *schema.Resource:
		if err := loadObject(t.Schema, data.(map[string]any), v.Elem()); err != nil {
			return err
		}
	case *schema.Schema:
		if err := loadSchema(t, data, v.Elem()); err != nil {
			return err
		}
	default:
		panic(fmt.Sprintf("unhandled terraform elem type: %T", t))
	}
	return nil
}

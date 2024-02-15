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

package intschematag

import (
	"fmt"
	"reflect"
	"strings"
)

// Internal field names.
const (
	HCLIDField   = "HCLID"
	StateIDField = "StateID"
)

// InternalFieldName is used for fields not based on the TF schema.
const InternalFieldName = "-"

const (
	tagName               = "intschema"
	optionalFlag          = "optional"
	computed              = "computed"
	listEncodedObjectFlag = "list_encoded_object"
	fileFieldType         = "file"
	defaultFlag           = "default"
)

// Tag contains metadata for a intschema struct tag.
type Tag struct {
	TFName            string
	Optional          bool
	Computed          bool
	ListEncodedObject bool
	Default           string
	File              bool
}

// IgnoreResourceData returns true for fields that should not be
// used when setting to/from ResourceData. E.g., fields only relevant
// to HCL marshalling like HCLID.
func (t Tag) IgnoreResourceData() bool {
	return t.TFName == InternalFieldName || t.File
}

// InternalFieldTag marks a field which is not part of the resource schema.
func InternalFieldTag() Tag {
	return Tag{TFName: InternalFieldName}
}

// Marshal marshals t into a struct tag string.
func (t Tag) Marshal() string {
	vs := []string{
		t.TFName,
	}
	if t.Optional {
		vs = append(vs, optionalFlag)
	}
	if t.File {
		vs = append(vs, fileFieldType)
	}
	if t.Computed {
		vs = append(vs, computed)
	}
	if t.ListEncodedObject {
		vs = append(vs, listEncodedObjectFlag)
	}
	if t.Default != "" {
		if strings.ContainsAny(t.Default, `'", `) {
			panic(fmt.Sprintf("field %q has default %q with unsupported characters", t.TFName, t.Default))
		}
		vs = append(vs, fmt.Sprintf("%s:%s", defaultFlag, t.Default))
	}
	return fmt.Sprintf("%s:%q", tagName, strings.Join(vs, ","))
}

// Unmarshal unmarshals an intschema Tag for the specified struct field.
func Unmarshal(f reflect.StructField) Tag {
	var t Tag
	vs := strings.Split(f.Tag.Get(tagName), ",")
	t.TFName = vs[0]
	for _, v := range vs[1:] {
		if v == optionalFlag {
			t.Optional = true
		}
		if v == fileFieldType {
			t.File = true
		}
		if v == computed {
			t.Computed = true
		}
		if v == listEncodedObjectFlag {
			t.ListEncodedObject = true
		}
		if key, val, ok := KeyVal(v); ok {
			switch key {
			case defaultFlag:
				t.Default = val
			}
		}
	}
	return t
}

// KeyVal parses a string in the format "<key>:<value>" form.
func KeyVal(s string) (key, val string, ok bool) {
	return strings.Cut(s, ":")
}

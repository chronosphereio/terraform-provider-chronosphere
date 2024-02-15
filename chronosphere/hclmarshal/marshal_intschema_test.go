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

package hclmarshal_test

import (
	"strings"
	"testing"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/hclmarshal"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/hclmarshal/hcltest"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/stretchr/testify/assert"
)

func TestMarshalIntSchema_Success(t *testing.T) {
	tests := []struct {
		msg  string
		r    intschema.TestResource
		want string
	}{
		{
			msg: "minimal",
			r: intschema.TestResource{
				HCLID:     "foo",
				Notifiers: []tfid.ID{tfid.Slug("some-notifier")},
			},
			want: `resource "chronosphere_test_resource" "foo" {
  dashboard_json = ""
  some_bool      = false
  some_float     = 0
  some_int       = 0

  some_object {
    inner_bool   = false
    inner_string = ""
  }

  some_string                = ""
  notifiers                  = ["some-notifier"]
  optional_bool_with_default = false
}`,
		},
		{
			msg: "all fields set",
			r: intschema.TestResource{
				HCLID: "bar",
				CollectionId: tfid.Ref{
					Type: "chronosphere_collection",
					ID:   "col1",
				}.AsID(),
				Notifiers: []tfid.ID{
					{},
					tfid.Slug("some-slug"),
					tfid.Ref{Type: "chronosphere_test_notifier", ID: "some_notifier"}.AsID(),
				},
				SomeBool:  true,
				SomeFloat: 1.1,
				SomeInt:   1,
				SomeObject: intschema.TestResourceSomeObject{
					InnerBool:   true,
					InnerString: "obj-inner",
				},
				SomeObjectSet: []intschema.TestResourceSomeObjectSet{
					{
						InnerBool:   true,
						InnerString: "obj-set-inner",
					},
				},
				SomeString:     "str",
				SomeStringList: []string{"str-list-1", "str-list-2"},
				SomeStringMap: map[string]string{
					"k": "v",
				},
				OptionalObject: &intschema.TestResourceOptionalObject{
					InnerStringList: []string{"inner-list-entry"},
				},
				OptionalStringList: []string{
					"opt-str-list-1", "opt-str-list-2",
				},
				OptionalBoolWithDefault: true,
				ComputedAndNotOptional:  "this-should-not-be-in-the-output",
				ComputedAndOptional:     "output-value",
				HCLFileDashboardJson:    "/tmp/dashboard.json",
			},
			want: `resource "chronosphere_test_resource" "bar" {
  collection_id  = chronosphere_collection.col1.id
  dashboard_json = file("/tmp/dashboard.json")
  some_bool      = true
  some_float     = 1.1
  some_int       = 1

  some_object {
    inner_bool   = true
    inner_string = "obj-inner"
  }

  some_object_set {
    inner_bool   = true
    inner_string = "obj-set-inner"
  }

  some_string      = "str"
  some_string_list = ["str-list-1", "str-list-2"]

  some_string_map = {
    k = "v"
  }

  computed_and_optional = "output-value"
  notifiers             = ["some-slug", chronosphere_test_notifier.some_notifier.id]

  optional_object {
    inner_string_list = ["inner-list-entry"]
  }

  optional_string_list = ["opt-str-list-1", "opt-str-list-2"]
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			got := hcltest.MustMarshalString(t, &tt.r)
			assert.Equal(t, tt.want, strings.TrimSpace(got))
		})
	}
}

func TestMarshalIntSchema_Data(t *testing.T) {
	s := &intschema.DataBucket{
		HCLID: "f",
		Slug:  "foo",
	}
	want := `
data "chronosphere_bucket" "f" {
  slug = "foo"
}
`
	assert.Equal(t, hcltest.MustMarshalString(t, s), want)
	assert.Equal(t, tfid.Ref{
		Datasource: true,
		Type:       "chronosphere_bucket",
		ID:         "f",
	}.AsID(), s.Ref())
}

func TestMarshalIntSchema_Panic(t *testing.T) {
	tests := []struct {
		msg       string
		s         any
		wantPanic string
	}{
		{
			msg:       "must pass pointer to struct",
			s:         struct{}{},
			wantPanic: "MarshalIntSchema must be called with pointer to an intschema struct",
		},
		{
			msg: "required object uses pointer",
			s: &struct {
				ObjectPtr *intschema.TestResourceSomeObject `intschema:"object_ptr"`
			}{},
			wantPanic: `field "object_ptr" is required, but is using a pointer`,
		},
		{
			msg: "unknown type",
			s: &struct {
				F32 float32 `intschema:"f32"`
			}{},
			wantPanic: `field "f32" has unsupported type float32`,
		},
		{
			msg: "map has unsupported value type",
			s: &struct {
				Objects map[string]intschema.TestResourceSomeObject `intschema:"objects"`
			}{
				Objects: map[string]intschema.TestResourceSomeObject{
					"foo": {},
				},
			},
			wantPanic: `map "objects" has an unsupported type`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			m := hclmarshal.New()
			b := m.AddResource("marshal_test", "foo")
			assert.PanicsWithError(t, tt.wantPanic, func() {
				_ = hclmarshal.MarshalIntSchema(tt.s, b)
			})
		})
	}
}

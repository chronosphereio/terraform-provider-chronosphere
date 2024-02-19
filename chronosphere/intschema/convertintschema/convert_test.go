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

package convertintschema_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema"
)

type dict = map[string]any

func TestConvert(t *testing.T) {
	tests := []struct {
		name  string
		input dict
		want  intschema.TestResource
	}{
		{
			name: "fully populated",
			input: dict{
				"collection_id":    "some-collection-slug",
				"notifiers":        []any{"notifier-1", "notifier-2"},
				"some_string":      "hello",
				"some_bool":        true,
				"some_float":       5.5,
				"some_int":         5,
				"some_string_list": []any{"foo", "bar", "baz"},
				"some_object_set": []any{
					dict{
						"inner_bool":   true,
						"inner_string": "x",
					},
					dict{
						"inner_bool":   false,
						"inner_string": "y",
					},
				},
				"some_string_map": dict{
					"dog": "woof",
					"cat": "meow",
				},
				"some_object": []any{
					dict{
						"inner_bool":   true,
						"inner_string": "z",
					},
				},
				"optional_object": []any{
					dict{
						"inner_string_list": []any{"foo"},
					},
				},
				"optional_bool_with_default": false,
			},
			want: intschema.TestResource{
				CollectionId: tfid.Slug("some-collection-slug"),
				// TF uses a deterministic but not alphabetical order for sets.
				Notifiers:      []tfid.ID{tfid.Slug("notifier-2"), tfid.Slug("notifier-1")},
				SomeString:     "hello",
				SomeBool:       true,
				SomeFloat:      5.5,
				SomeInt:        5,
				SomeStringList: []string{"foo", "bar", "baz"},
				SomeObjectSet: []intschema.TestResourceSomeObjectSet{
					{
						InnerBool:   true,
						InnerString: "x",
					},
					{
						InnerBool:   false,
						InnerString: "y",
					},
				},
				SomeStringMap: map[string]string{
					"dog": "woof",
					"cat": "meow",
				},
				SomeObject: intschema.TestResourceSomeObject{
					InnerBool:   true,
					InnerString: "z",
				},
				OptionalObject: &intschema.TestResourceOptionalObject{
					InnerStringList: []string{"foo"},
				},
				OptionalBoolWithDefault: false,
			},
		},
		{
			name:  "empty",
			input: dict{},
			want: intschema.TestResource{
				OptionalBoolWithDefault: true, // unset uses the Default.
			},
		},
		{
			name: "explicitly set default values",
			input: dict{
				"optional_bool_with_default": true,
			},
			want: intschema.TestResource{
				OptionalBoolWithDefault: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputData := schema.TestResourceDataRaw(t, tfschema.TestResource, tt.input)
			inputData.SetId("some-state-id")

			var r1 intschema.TestResource
			require.NoError(t, r1.FromResourceData(inputData))

			// Only available on the FromResourceData; is rejected by
			// ToResourceData, so we clear the field on the way back.
			require.Equal(t, "some-state-id", r1.StateID)
			r1.StateID = ""

			require.Equal(t, tt.want, r1)

			outputData := schema.TestResourceDataRaw(t, tfschema.TestResource, nil)
			require.Nil(t, r1.ToResourceData(outputData))

			var r2 intschema.TestResource
			require.NoError(t, r2.FromResourceData(outputData))
			require.Equal(t, tt.want, r2)
		})
	}
}

func TestOptionalFieldServerChangesAreDetectedLocally(t *testing.T) {
	tests := []struct {
		local  []string
		server []string
		want   []string
	}{
		{local: nil, server: nil, want: nil},

		{local: []string{}, server: []string{}, want: nil},

		{local: []string{"foo"}, server: []string{"bar"}, want: []string{"bar"}},

		{local: nil, server: []string{"foo"}, want: []string{"foo"}},
		{local: []string{"foo"}, server: nil, want: nil},

		{local: []string{}, server: []string{"foo"}, want: []string{"foo"}},
		{local: []string{"foo"}, server: []string{}, want: nil},

		// This is what the special optional field code actually does: physical
		// changes to container types are ignored if the two values aren't
		// semantically different (nil and empty are treated as equal).
		{local: nil, server: []string{}, want: nil},
		{local: []string{}, server: nil, want: nil},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%#v vs %#v", tt.local, tt.server), func(t *testing.T) {
			// Server has some value.
			server := intschema.TestResource{
				OptionalStringList: tt.server,
			}

			// Local data has some value.
			localData := schema.TestResourceDataRaw(t, tfschema.TestResource, nil)
			require.NoError(t, localData.Set("optional_string_list", tt.local))

			// Write the server values to local data.
			require.Nil(t, server.ToResourceData(localData))

			// This actually doesn't check whether Set was or
			// wasn't called, since GetOk will still return ok=false if Set was
			// called with nil.
			v, ok := localData.GetOk("optional_string_list")
			if tt.want != nil {
				require.True(t, ok)
				var vStr []string
				for _, x := range v.([]any) {
					vStr = append(vStr, x.(string))
				}
				require.Equal(t, tt.want, vStr)
			} else {
				require.False(t, ok)
			}

			// For completeness, lift the local data back into a struct and
			// ensure we still have the expected value.
			updatedLocal := intschema.TestResource{}
			require.NoError(t, updatedLocal.FromResourceData(localData))
			require.Equal(t, tt.want, updatedLocal.OptionalStringList)
		})
	}
}

func TestCannotSetInternalFields(t *testing.T) {
	tests := []struct {
		field    string
		resource *intschema.TestResource
	}{
		{
			field:    "HCLID",
			resource: &intschema.TestResource{HCLID: "something"},
		},
		{
			field:    "StateID",
			resource: &intschema.TestResource{StateID: "something"},
		},
		{
			field:    "HCLFileDashboardJson",
			resource: &intschema.TestResource{HCLFileDashboardJson: "/tmp/file"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.field, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, tfschema.TestResource, nil)
			wantMsg := fmt.Sprintf("cannot set field %s when calling ToResourceData", tt.field)
			require.PanicsWithValue(t, wantMsg, func() {
				tt.resource.ToResourceData(d)
			})
		})
	}
}

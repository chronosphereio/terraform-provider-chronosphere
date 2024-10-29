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

package chronosphere

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/stretchr/testify/assert"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema/intschematest"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/shared/pkg/container/set"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
)

func TestSetUnknown(t *testing.T) {
	type noTFID struct {
		Name string `intschema:"name"`
	}
	type singleTFID struct {
		Name   string  `intschema:"name"`
		RefVal tfid.ID `intschema:"ref_val"`
	}
	type multipleTFID struct {
		Name    string  `intschema:"name"`
		RefVal  tfid.ID `intschema:"ref_val"`
		RefVal2 tfid.ID `intschema:"ref_val2"`
	}
	type nestedTFID struct {
		Name   string     `intschema:"name"`
		Nested singleTFID `intschema:"nested"`
		RefVal tfid.ID    `intschema:"ref_val"`
	}
	type sliceOfRefs struct {
		Name    string    `intschema:"name"`
		RefVals []tfid.ID `intschema:"ref_vals"`
		RefVal  tfid.ID   `intschema:"ref_val"`
	}
	type mapWithRefs struct {
		Name    string             `intschema:"name"`
		MapRefs map[string]tfid.ID `intschema:"map_vals"`
		RefVal  tfid.ID            `intschema:"ref_val"`
	}
	type mapWithName struct {
		Name string `intschema:"name"`
	}
	type mapWithNestedName struct {
		Nested *mapWithName `intschema:"nested,list_encoded_object"`
	}

	testCases := []struct {
		name           string
		input          any
		rawConfig      cty.Value
		skipIDs        set.Set[string]
		dryRunDefaults map[string]any
		want           any
		wantPanicMsg   string
	}{
		{
			name:  "no tfid values",
			input: &noTFID{Name: "michael"},
			want:  &noTFID{Name: "michael"},
		},
		{
			name: "single empty top level tfid value",
			input: &singleTFID{
				Name: "michael",
			},
			rawConfig: cty.ObjectVal(map[string]cty.Value{
				"ref_val": cty.StringVal("ref_slug"),
			}),
			want: &singleTFID{
				Name:   "michael",
				RefVal: dummyRef,
			},
		},
		{
			name: "single empty top level tfid value with nil raw config is not changed",
			input: &singleTFID{
				Name: "michael",
			},
			rawConfig: cty.ObjectVal(map[string]cty.Value{
				"ref_val": cty.NilVal,
			}),
			want: &singleTFID{
				Name: "michael",
			},
		},
		{
			name: "populated tfid value is not changed",
			input: &singleTFID{
				Name:   "michael",
				RefVal: tfid.Slug("populated"),
			},
			rawConfig: cty.ObjectVal(map[string]cty.Value{
				"ref_val": cty.StringVal("ignored"),
			}),
			want: &singleTFID{
				Name:   "michael",
				RefVal: tfid.Slug("populated"),
			},
		},
		{
			name: "multiple empty top level tfid values",
			input: &multipleTFID{
				Name: "michael",
			},
			rawConfig: cty.ObjectVal(map[string]cty.Value{
				"ref_val":  cty.StringVal("ref_slug"),
				"ref_val2": cty.StringVal("ref_slug"),
			}),
			want: &multipleTFID{
				Name:    "michael",
				RefVal:  dummyRef,
				RefVal2: dummyRef,
			},
		},
		{
			name: "nested tfid unsupported",
			input: &nestedTFID{
				Name: "michael",
				Nested: singleTFID{
					Name: "michael",
				},
			},
			wantPanicMsg: "setUnknown found unsupported tfid in a slice/map at nested.ref_val",
		},
		{
			name: "nested tfid skipped",
			input: &nestedTFID{
				Name: "michael",
				Nested: singleTFID{
					Name: "michael",
				},
			},
			rawConfig: cty.ObjectVal(map[string]cty.Value{
				"ref_val": cty.StringVal("ref_slug"),
			}),
			skipIDs: set.New("nested.ref_val"),
			want: &nestedTFID{
				Name:   "michael",
				RefVal: dummyRef,
				Nested: singleTFID{
					Name: "michael",
				},
			},
		},
		{
			name: "top level nil slice of tfids without skip panics",
			input: &sliceOfRefs{
				Name: "michael",
			},
			wantPanicMsg: "setUnknown found unsupported tfid in a slice/map at ref_vals.[]",
		},
		// NOTE: the below behavior does not match the behaviour on unknown notifier references,
		// which uses an empty list. See the note in resource_notification_policy on dry-run validation.
		{
			name: "top level slice of empty tfids without skip panics",
			input: &sliceOfRefs{
				Name:    "michael",
				RefVals: []tfid.ID{tfid.Slug(""), tfid.Slug("")},
			},
			wantPanicMsg: "setUnknown found unsupported tfid in a slice/map at ref_vals.[]",
		},
		// NOTE: the below behavior does not match the behaviour on unknown notifier references,
		// which uses an empty list. See the note in resource_notification_policy on dry-run validation.
		{
			name: "top level slice of tfids with skip",
			input: &sliceOfRefs{
				Name:    "michael",
				RefVals: []tfid.ID{tfid.Slug(""), tfid.Slug("")},
			},
			rawConfig: cty.ObjectVal(map[string]cty.Value{
				"ref_val": cty.StringVal("ref_slug"),
			}),
			skipIDs: set.New("ref_vals.[]"),
			want: &sliceOfRefs{
				Name:    "michael",
				RefVals: []tfid.ID{tfid.Slug(""), tfid.Slug("")},
				RefVal:  dummyRef,
			},
		},
		{
			name: "top level map with tfids without skip panics",
			input: &mapWithRefs{
				Name: "michael",
				MapRefs: map[string]tfid.ID{
					"one": tfid.Slug(""),
				},
			},
			wantPanicMsg: "setUnknown found unsupported tfid in a slice/map at map_vals.[]",
		},
		{
			name: "top level nil map with tfids without skip panics",
			input: &mapWithRefs{
				Name: "michael",
			},
			wantPanicMsg: "setUnknown found unsupported tfid in a slice/map at map_vals.[]",
		},
		{
			name: "top level empty map with tfids without skip panics",
			input: &mapWithRefs{
				Name:    "michael",
				MapRefs: map[string]tfid.ID{},
			},
			wantPanicMsg: "setUnknown found unsupported tfid in a slice/map at map_vals.[]",
		},
		{
			name: "top level map with tfids with skip",
			input: &mapWithRefs{
				Name: "michael",
				MapRefs: map[string]tfid.ID{
					"one": tfid.Slug(""),
				},
			},
			rawConfig: cty.ObjectVal(map[string]cty.Value{
				"ref_val": cty.StringVal("ref_slug"),
			}),
			skipIDs: set.New("map_vals.[]"),
			want: &mapWithRefs{
				Name: "michael",
				MapRefs: map[string]tfid.ID{
					"one": tfid.Slug(""),
				},
				RefVal: dummyRef,
			},
		},
		{
			name: "dry run default at top level",
			input: &mapWithName{
				Name: "",
			},
			rawConfig: cty.ObjectVal(map[string]cty.Value{
				"name": cty.UnknownVal(cty.String),
			}),
			dryRunDefaults: map[string]any{"name": "dry-run-default"},
			want: &mapWithName{
				Name: "dry-run-default",
			},
		},
		{
			name: "dry run default in nested struct",
			input: &mapWithNestedName{
				Nested: &mapWithName{
					Name: "",
				},
			},
			rawConfig: cty.ObjectVal(map[string]cty.Value{
				"nested": cty.ListVal([]cty.Value{
					cty.ObjectVal(map[string]cty.Value{
						"name": cty.UnknownVal(cty.String),
					}),
				}),
			}),
			dryRunDefaults: map[string]any{"nested.[0].name": "dry-run-default"},
			want: &mapWithNestedName{
				Nested: &mapWithName{
					Name: "dry-run-default",
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			p := setUnknownParams{
				rawConfig:      tt.rawConfig,
				skipIDs:        tt.skipIDs,
				dryRunDefaults: tt.dryRunDefaults,
			}
			if tt.wantPanicMsg != "" {
				assert.PanicsWithValue(t, tt.wantPanicMsg,
					func() {
						setUnknown(tt.input, p)
					},
				)
				return
			}

			r := intschematest.Clone(tt.input)
			setUnknown(r, p)
			assert.Equal(t, tt.want, r)
		})
	}
}

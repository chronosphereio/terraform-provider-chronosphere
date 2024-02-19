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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
)

func TestSetUnknownReferences(t *testing.T) {
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

	testCases := []struct {
		name         string
		input        any
		skip         []string
		rawConfig    cty.Value
		want         any
		wantPanicMsg string
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
			wantPanicMsg: "setUnknownReferences found unsupported tfid in a slice/map at nested.ref_val",
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
			skip: []string{"nested.ref_val"},
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
			wantPanicMsg: "setUnknownReferences found unsupported tfid in a slice/map at ref_vals.[]",
		},
		// NOTE: the below behavior does not match the behaviour on unknown notifier references,
		// which uses an empty list. See the note in resource_notification_policy on dry-run validation.
		{
			name: "top level slice of empty tfids without skip panics",
			input: &sliceOfRefs{
				Name:    "michael",
				RefVals: []tfid.ID{tfid.Slug(""), tfid.Slug("")},
			},
			wantPanicMsg: "setUnknownReferences found unsupported tfid in a slice/map at ref_vals.[]",
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
			skip: []string{"ref_vals.[]"},
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
			wantPanicMsg: "setUnknownReferences found unsupported tfid in a slice/map at map_vals.[]",
		},
		{
			name: "top level nil map with tfids without skip panics",
			input: &mapWithRefs{
				Name: "michael",
			},
			wantPanicMsg: "setUnknownReferences found unsupported tfid in a slice/map at map_vals.[]",
		},
		{
			name: "top level empty map with tfids without skip panics",
			input: &mapWithRefs{
				Name:    "michael",
				MapRefs: map[string]tfid.ID{},
			},
			wantPanicMsg: "setUnknownReferences found unsupported tfid in a slice/map at map_vals.[]",
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
			skip: []string{"map_vals.[]"},
			want: &mapWithRefs{
				Name: "michael",
				MapRefs: map[string]tfid.ID{
					"one": tfid.Slug(""),
				},
				RefVal: dummyRef,
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanicMsg != "" {
				assert.PanicsWithValue(t, tt.wantPanicMsg,
					func() {
						setUnknownReferences(tt.input, tt.rawConfig, nil)
					},
				)
				return
			}

			r := intschematest.Clone(tt.input)
			setUnknownReferences(r, tt.rawConfig, tt.skip)
			assert.Equal(t, tt.want, r)
		})
	}
}

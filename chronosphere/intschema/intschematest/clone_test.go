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

package intschematest

import (
	"testing"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/intschema"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {
	tests := []struct {
		msg string
		v   any
	}{
		{
			msg: "collection with team_id slug",
			v: &intschema.Collection{
				HCLID:  "col1",
				Name:   "collection name",
				Slug:   "col-1",
				TeamId: tfid.Slug("t1"),
			},
		},
		{
			msg: "collection with team_id ref",
			v: &intschema.Collection{
				HCLID:       "col1",
				Name:        "collection name",
				Slug:        "col-1",
				TeamId:      tfid.Ref{Type: "chronosphere_team", ID: "team1"}.AsID(),
				Description: "this is a collection",
			},
		},
		{
			msg: "empty TestResource",
			v:   &intschema.TestResource{},
		},
		{
			msg: "TestResource with non-nil empty containers",
			v: &intschema.TestResource{
				SomeObjectSet:  make([]intschema.TestResourceSomeObjectSet, 0),
				SomeStringMap:  make(map[string]string),
				SomeStringList: make([]string, 0),
			},
		},
		{
			msg: "full TestResource",
			v: intschema.TestResource{
				HCLID: "bar",
				CollectionId: tfid.Ref{
					Type: "chronosphere_collection",
					ID:   "col1",
				}.AsID(),
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
					InnerStringList: []string{"foo"},
				},
				OptionalStringList: []string{
					"opt-str-list-1", "opt-str-list-2",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			assert.Equal(t, tt.v, Clone(tt.v))
		})
	}
}

func TestCloneUnsupported(t *testing.T) {
	assert.Panics(t, func() {
		Clone(struct{ F func() }{})
	})
}

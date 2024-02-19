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

package tfschema

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// TestResource is used exclusively for unit testing.
var TestResource = map[string]*schema.Schema{
	"some_string": {
		Type: schema.TypeString,
	},
	"some_bool": {
		Type: schema.TypeBool,
	},
	"some_float": {
		Type: schema.TypeFloat,
	},
	"some_int": {
		Type: schema.TypeInt,
	},
	"some_string_list": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"some_object_set": {
		Type: schema.TypeSet,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"inner_string": {
					Type: schema.TypeString,
				},
				"inner_bool": {
					Type: schema.TypeBool,
				},
			},
		},
	},
	"some_string_map": {
		Type: schema.TypeMap,
		Elem: &schema.Schema{Type: schema.TypeString},
	},
	"some_object": {
		Type:     schema.TypeList,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"inner_string": {
					Type: schema.TypeString,
				},
				"inner_bool": {
					Type: schema.TypeBool,
				},
			},
		},
	},
	"optional_object": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"inner_string_list": {
					Type: schema.TypeList,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	},
	"optional_string_list": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional: true,
	},
	"collection_id": {
		Type: schema.TypeString,
	},
	// Intentionally match the exact notifiers schema since it's the only
	// list of TF IDs, and we want to ensure it works correctly.
	"notifiers": NotificationRouteSchema.Elem.(*schema.Resource).Schema["notifiers"],

	// Matches send_resolved in real resources.
	"optional_bool_with_default": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  true,
	},
	"computed_and_not_optional": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"computed_and_optional": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
	"dashboard_json": {
		Type: schema.TypeString,
	},
}

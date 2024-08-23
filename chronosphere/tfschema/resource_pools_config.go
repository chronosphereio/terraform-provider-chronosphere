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

const maxResourcePools = 128

var ResourcePoolsConfig = map[string]*schema.Schema{
	"default_pool": {
		Type: schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allocation": ResourcePoolDefaultPoolAllocationSchema,
				"priorities": ResourcePoolPrioritiesSchema,
			},
			SchemaVersion: 1,
		},
		Required: true,
		MaxItems: 1,
	},
	"pools": {
		Type:          schema.TypeList,
		Elem:          ResourcePoolElemSchema,
		Optional:      true,
		ConflictsWith: []string{"pool"},
		Deprecated:    "Use pool instead of pools",
		MaxItems:      maxResourcePools,
	},
	"pool": {
		Type:          schema.TypeList,
		Elem:          ResourcePoolElemSchema,
		ConflictsWith: []string{"pools"},
		Optional:      true,
		MaxItems:      maxResourcePools,
	},
}

var ResourcePoolDefaultPoolAllocationSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"percent_of_license": {
				Type:     schema.TypeFloat,
				Required: false,
			},
		},
	},
	MaxItems:   1,
	Required:   false,
	Deprecated: "default_pool allocation is not required as it can be derived",
}

var ResourcePoolAllocationSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"percent_of_license": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"fixed_values": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePoolAllocationFixedValueSchema,
				MinItems: 1,
			},
		},
	},
	MaxItems: 1,
	Required: true,
}

var ResourcePoolAllocationFixedValueSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"license": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	},
	MaxItems: 1,
	Required: true,
}

var ResourcePoolElemSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"allocation": ResourcePoolAllocationSchema,
		"match_rule": {
			Type:       schema.TypeString,
			Optional:   true,
			Deprecated: "use match_rules",
		},
		"match_rules": {
			Type:     schema.TypeList,
			Elem:     &schema.Schema{Type: schema.TypeString},
			MinItems: 1,
			Optional: true,
		},
		"priorities": ResourcePoolPrioritiesSchema,
	},
}

var ResourcePoolPrioritiesSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"high_priority_match_rules": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				MinItems: 1,
				Optional: true,
			},
			"low_priority_match_rules": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				MinItems: 1,
				Optional: true,
			},
		},
	},
	MaxItems: 1,
	Optional: true,
}

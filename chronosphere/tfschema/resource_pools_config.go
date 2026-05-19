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
		Type:        schema.TypeList,
		Description: "Catch-all pool that receives metrics not matched by any other pool. Also receives any license allocation left unassigned by the other pools.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allocation":          ResourcePoolAllocationSchema,
				"priorities":          ResourcePoolPrioritiesSchema,
				"priority_thresholds": ResourcePoolAllocationThresholdsSchema,
			},
			SchemaVersion: 1,
		},
		Optional: true,
		MaxItems: 1,
	},
	"pools": {
		Type:          schema.TypeList,
		Elem:          ResourcePoolElemSchema,
		Optional:      true,
		ConflictsWith: []string{"pool"},
		Deprecated:    "Use pool instead of pools",
		Description:   "Deprecated: use `pool` instead. Set of named pools that partition the license.",
		MaxItems:      maxResourcePools,
	},
	"pool": {
		Type:          schema.TypeList,
		Elem:          ResourcePoolElemSchema,
		ConflictsWith: []string{"pools"},
		Optional:      true,
		Description:   "Named pools that partition each license across teams or workloads. Pools are matched in declaration order via their `match_rules`.",
		MaxItems:      maxResourcePools,
	},
}

var ResourcePoolAllocationSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "License allocation for the pool. Can be expressed as a percentage of the license (`percent_of_license`) or as per-license fixed values (`fixed_value`).",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"percent_of_license": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "Percent of each license to allocate to this pool, between 0 and 100. Across non-default pools, the sum must not exceed 100; the default pool receives the remainder.",
			},
			"fixed_value": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        ResourcePoolAllocationFixedValueSchema,
				MinItems:    1,
				Description: "Per-license fixed allocations that override `percent_of_license` for the named licenses. When any pool sets a fixed value for a license, every pool must also set one for that license.",
			},
			"priority_thresholds": ResourcePoolAllocationThresholdsSchema,
		},
	},
	MaxItems: 1,
	Optional: true,
}

var ResourcePoolAllocationFixedValueSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"license": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "License this fixed-value allocation applies to (e.g. `PERSISTED_WRITES`).",
		},
		"value": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "Fixed amount of the license to allocate, in the license's native unit.",
		},
	},
}

var ResourcePoolElemSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Unique name of the pool.",
		},
		"allocation": ResourcePoolAllocationSchema,
		"match_rule": {
			Type:        schema.TypeString,
			Optional:    true,
			Deprecated:  "use match_rules",
			Description: "Deprecated: use `match_rules` instead. Single matcher selecting metrics that belong to this pool.",
		},
		"match_rules": {
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			MinItems:    1,
			Optional:    true,
			Description: "Matchers selecting metrics that map to this pool. A metric matching any rule is assigned to the pool.",
		},
		"priorities": ResourcePoolPrioritiesSchema,
	},
}

var ResourcePoolPrioritiesSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Optional high/low priority sub-classifications within the pool. Low-priority metrics are dropped first; high-priority metrics are dropped last when limits are hit.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"high_priority_match_rules": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				MinItems:    1,
				Optional:    true,
				Description: "Matchers selecting metrics within the pool that should be treated as high priority and dropped last.",
			},
			"low_priority_match_rules": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				MinItems:    1,
				Optional:    true,
				Description: "Matchers selecting metrics within the pool that should be treated as low priority and dropped first.",
			},
		},
	},
	MaxItems: 1,
	Optional: true,
}

var ResourcePoolAllocationThresholdsSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Per-license drop thresholds for `PERSISTED_CARDINALITY_STANDARD` and `PERSISTED_CARDINALITY_HISTOGRAM` only. Defines strict upper bounds beyond which new consumption is dropped, optionally segmented by priority class.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"license": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "License the thresholds apply to.",
			},
			"all_priorities":           ResourcePoolAllocationThresholdSchema,
			"default_and_low_priority": ResourcePoolAllocationThresholdSchema,
			"low_priority":             ResourcePoolAllocationThresholdSchema,
		},
	},
	MinItems: 1,
	Optional: true,
}

var ResourcePoolAllocationThresholdSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Threshold value, expressed as either a percent of the pool's allocation or as a fixed value in license units.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"percent_of_pool_allocation": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "Threshold as a percent of the pool's allocation. `100` equals the full allocation; values above 100 allow the pool to exceed its baseline allocation.",
			},
			"fixed_value": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Threshold expressed as a fixed value of the license, in the license's native unit.",
			},
		},
	},
	MaxItems: 1,
	Optional: true,
}

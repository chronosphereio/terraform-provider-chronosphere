// Copyright 2026 Chronosphere Inc.
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

var QueryResourcePools = map[string]*schema.Schema{
	"pool": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Pools that group automated query sources. Query sources map to the first pool that matches, in declaration order.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Unique name of the pool.",
				},
				"source_types": {
					Type:        schema.TypeList,
					Required:    true,
					MinItems:    1,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Description: "Source types that map to this pool: `MONITOR`, `RECORDING_RULE`, `SLO`, or `SERVICE_ACCOUNT`. A query source maps to the first pool whose `source_types` matches its source type and whose `source_names` (if set) matches its source name.",
				},
				"source_names": {
					Type:        schema.TypeList,
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Description: "Narrows the pool to specific named sources of the pool's source types: service account IDs for `SERVICE_ACCOUNT`, monitor slugs for `MONITOR`, recording rule slugs for `RECORDING_RULE`, and SLO slugs for `SLO`.",
				},
				"data_read_limit": QueryResourcePoolDataReadLimitSchema,
			},
		},
	},
	"default_pool": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Catch-all pool that receives query sources not matched by any other pool.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"data_read_limit": QueryResourcePoolDataReadLimitSchema,
			},
		},
	},
}

var QueryResourcePoolDataReadLimitSchema = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	MaxItems:    1,
	Description: "Data read limit enforced on the pool. Omit to leave the pool unlimited.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_datapoints_read_per_second": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Maximum datapoints read per second for the pool, aggregated across all sources in the pool. A value of `0` blocks all queries from the pool.",
			},
			"sustain_secs": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "How long the pool must remain above `max_datapoints_read_per_second` before its queries are rejected. Must be a whole multiple of 60 seconds, between 60 and 300; `0` (or omitted) applies the 60-second default. Must be omitted when `max_datapoints_read_per_second` is `0`, which rejects queries immediately.",
			},
		},
	},
}

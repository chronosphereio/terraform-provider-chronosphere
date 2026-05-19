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

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/aggregationfilter"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var MappingRule = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the mapping rule. Can be changed after creation.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the mapping rule. Generated from `name` if omitted. Immutable after creation.",
	},
	"bucket_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "ID of the bucket the mapping rule belongs to.",
	},
	"filter": Filter{
		KVDelimiter: aggregationfilter.MappingRuleDelimiter,
		Description: "Space-delimited list of `label=value_glob` matchers that select the metrics this rule applies to. A metric must match every filter to be considered.",
	}.Schema(),
	"aggregations": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Aggregation type applied to matching metrics. Cannot be set if `drop` is `true`.",
		Elem: Enum{
			Value: enum.AggregationType.ToStrings(),
		}.Schema(),
	},
	// Storage policies to apply to the mapped metrics.
	"storage_policy": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Storage policy controlling resolution and retention of mapped metrics. Deprecated: use `interval` instead.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"resolution": Duration{
					Required:    true,
					Description: "Resolution at which mapped data points are stored.",
				}.Schema(),
				"retention": Duration{
					Required:    true,
					Description: "Retention duration for mapped data points.",
				}.Schema(),
			},
		},
		Deprecated: "use `interval` instead",
	},
	"drop": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "If `true`, drops the matching metrics instead of aggregating them. Cannot be set together with `aggregations`. Defaults to `false`.",
	},
	// Whether or not to drop the timestamp.
	"drop_timestamp": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Deprecated:  "drop timestamp is no longer supported",
		Description: "Deprecated: no longer supported.",
	},
	"interval": {
		Type:     schema.TypeString,
		Optional: true,
		// When no interval is specified, a server-side default is used.
		Computed:      true,
		ConflictsWith: []string{"storage_policy"},
		Description:   "Interval between aggregated data points produced by this mapping rule. Defaults to a server-side value when unset. Conflicts with `storage_policy`.",
	},
	"mode": Enum{
		Value:       enum.MappingModeType.ToStrings(),
		Optional:    true,
		Description: "Mapping rule mode controlling whether it is active or in a preview state.",
	}.Schema(),
}

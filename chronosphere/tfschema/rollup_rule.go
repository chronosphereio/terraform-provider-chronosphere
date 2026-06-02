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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/aggregationfilter"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
)

var RollupRule = map[string]*schema.Schema{
	"bucket_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "ID of the bucket the rollup rule belongs to.",
	},
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the rollup rule. Can be changed after creation.",
	},
	"slug": {
		Type: schema.TypeString,
		// This is intentionally different from other resources due to unique issues with Rollup Rules
		Required:    true,
		ForceNew:    true,
		Description: "Stable identifier for the rollup rule. Immutable after creation. Unlike most resources, the slug is required and is not auto-generated from `name`.",
	},
	"filter": Filter{
		KVDelimiter: aggregationfilter.RollupRuleDelimiter,
		Description: "Space-delimited list of `label:value_glob` matchers that select the input series. Supports glob patterns and special filters like `__name__`, `__metric_type__`, and `__metric_source__`.",
	}.Schema(),
	"new_metric": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Name of the output metric produced by the rollup. Supports the `{{.MetricName}}` template variable to reference the source metric name. Optional for Graphite rollup rules.",
	},
	"metric_type": Enum{
		Value:       enum.MetricType.ToStrings(),
		Required:    true,
		Description: "Type of the source metric being rolled up (e.g. `gauge`, `counter`, `histogram`).",
	}.Schema(),
	"aggregation": Enum{
		Value:       enum.AggregationType.ToStrings(),
		Optional:    true,
		Description: "Aggregation function applied across grouped series (e.g. `sum`, `min`, `max`, `last`).",
	}.Schema(),
	"storage_policies": {
		Type:        schema.TypeList,
		Optional:    true,
		MinItems:    1,
		MaxItems:    1,
		Description: "Storage policy controlling resolution and retention of rolled-up metrics. Deprecated: use `interval` instead.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"resolution": Duration{
					Required:    true,
					Description: "Resolution at which rolled-up data points are stored.",
				}.Schema(),
				"retention": Duration{
					Required:    true,
					Description: "Retention duration for rolled-up data points.",
				}.Schema(),
			},
		},
		Deprecated: "use `interval` instead",
	},
	"interval": {
		Type:     schema.TypeString,
		Optional: true,
		// When no interval is specified, a server-side default is used.
		Computed:      true,
		ConflictsWith: []string{"storage_policies"},
		Description:   "Interval between aggregated data points produced by the rollup. Defaults to a server-side value when unset. Conflicts with `storage_policies`.",
	},
	"group_by": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Labels to preserve when aggregating; all other labels are dropped. Mutually exclusive with `exclude_by`.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"exclude_by": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Labels to drop when aggregating; all other labels are preserved. Mutually exclusive with `group_by`.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"metric_type_tag": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "Whether to add a `__rollup_type__` label to the output metric identifying the rollup type. Defaults to `false`.",
	},
	"drop_raw": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "If `true`, automatically generates a drop rule that removes the raw input metrics matching this rollup. Defaults to `false`.",
	},
	"permissive": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Deprecated:  "permissive is no longer supported",
		Description: "Deprecated: no longer supported.",
	},
	"mode": Enum{
		Value:       enum.RollupModeType.ToStrings(),
		Optional:    true,
		Description: "Rollup mode controlling whether the rule is active or in a preview state.",
	}.Schema(),
	"skip_on_conflict": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		Description: "If `true`, this rule is skipped when another rollup rule already produces a metric with the same output name. Defaults to `false`.",
	},
	"graphite_label_policy": {
		Type:        schema.TypeList,
		MaxItems:    1,
		Optional:    true,
		Description: "Graphite-specific label policy applied to positional labels (`__gX__`) on the output metric.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"replace": {
					Type:        schema.TypeList,
					Optional:    true,
					Description: "List of positional Graphite label replacements applied to the output metric.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Required:    true,
								Description: "Positional Graphite label to replace (e.g. `__g1__`).",
							},
							"new_value": {
								Type:        schema.TypeString,
								Required:    true,
								Description: "Replacement value for the named positional label.",
							},
						},
					},
				},
			},
		},
	},
}

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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
)

var TraceMetricsRule = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the trace metrics rule.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the trace metrics rule. Generated from `name` if omitted. Immutable after creation.",
	},
	"metric_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Base name of the generated Prometheus metrics emitted by this rule.",
	},
	"metric_labels": {
		Type:        schema.TypeMap,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Static key/value labels added to every metric series emitted by the rule.",
	},
	"histogram_buckets_seconds": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeFloat,
		},
		Optional:    true,
		Description: "Histogram bucket upper bounds in seconds for the generated span-duration histogram metric.",
	},
	"trace_filter": TraceSearchFilterSchema,
	"scope_filter": TraceScopeFilterSchema,
	"group_by": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Span attributes to project into metric labels. Each entry maps a key on the matched span to a label on the resulting metric series.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key": {
					Type:        schema.TypeList,
					MinItems:    1,
					MaxItems:    1,
					Required:    true,
					Description: "Span attribute to group by.",
					Elem:        traceMetricsRuleGroupByKeySchema,
				},
				"label": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Name of the resulting metric label.",
				},
			},
		},
	},
}

var traceMetricsRuleGroupByKeySchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"type": Enum{
			Value:       enum.TraceMetricsRuleGroupByType.ToStrings(),
			Required:    true,
			Description: "Category of span attribute to group by (for example a well-known field such as `SERVICE` or `OPERATION`, or a generic span `TAG`).",
		}.Schema(),
		"named_key": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the span tag when `type` requires one (for example `TAG`). Ignored for fixed-key types.",
		},
	},
}

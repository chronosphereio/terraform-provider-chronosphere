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
		Type:     schema.TypeString,
		Required: true,
	},
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"metric_name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"metric_labels": {
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
	"histogram_buckets_seconds": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeFloat,
		},
		Optional: true,
	},
	"trace_filter": TraceSearchFilterSchema,
	"group_by": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key": {
					Type:     schema.TypeList,
					MinItems: 1,
					MaxItems: 1,
					Required: true,
					Elem:     traceMetricsRuleGroupByKeySchema,
				},
				"label": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	},
}

var traceMetricsRuleGroupByKeySchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"type": Enum{
			Value:    enum.TraceMetricsRuleGroupByType.ToStrings(),
			Required: true,
		}.Schema(),
		"named_key": {
			Type:     schema.TypeString,
			Optional: true,
		},
	},
}

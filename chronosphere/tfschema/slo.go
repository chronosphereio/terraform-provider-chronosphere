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

var Slo = map[string]*schema.Schema{
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
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"collection_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"notification_policy_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"sli": {
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		MinItems: 1,
		Elem: &schema.Resource{
			Schema: SLI,
		},
	},
	"definition": {
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		MinItems: 1,
		Elem: &schema.Resource{
			Schema: SloDefinition,
		},
	},
	"signal_grouping": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"annotations": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
}

var SloDefinition = map[string]*schema.Schema{
	"objective": {
		Type:     schema.TypeFloat,
		Required: true,
	},
	"reporting_windows": {
		Type:     schema.TypeSet,
		Required: true,
		MinItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"duration": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	},
	"low_volume": {
		Type:     schema.TypeBool,
		Optional: true,
	},
}

var SLI = map[string]*schema.Schema{
	"lens_template_indicator": {
		Type:         schema.TypeString,
		Optional:     true,
		RequiredWith: []string{"endpoint_availability", "endpoint_latency"},
	},
	"endpoint_label": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"custom_indicator": {
		Type:         schema.TypeList,
		Optional:     true,
		MaxItems:     1,
		ExactlyOneOf: []string{"custom_indicator", "endpoint_availability", "endpoint_latency"},
		Elem: &schema.Resource{
			Schema: SloCustomIndicator,
		},
	},
	"endpoint_availability": {
		Type:         schema.TypeList,
		Optional:     true,
		MaxItems:     1,
		ExactlyOneOf: []string{"custom_indicator", "endpoint_availability", "endpoint_latency"},
		Elem: &schema.Resource{
			Schema: SloEndpointAvailability,
		},
	},
	"endpoint_latency": {
		Type:         schema.TypeList,
		Optional:     true,
		MaxItems:     1,
		ExactlyOneOf: []string{"custom_indicator", "endpoint_availability", "endpoint_latency"},
		Elem: &schema.Resource{
			Schema: SloEndpointLatency,
		},
	},
}

var SloEndpointAvailability = map[string]*schema.Schema{
	"endpoints_monitored": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"success_codes": {
		Type:         schema.TypeSet,
		Optional:     true,
		ExactlyOneOf: []string{"success_codes", "error_codes"},
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"error_codes": {
		Type:         schema.TypeSet,
		Optional:     true,
		ExactlyOneOf: []string{"success_codes", "error_codes"},
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"additional_promql_filters": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Required: true,
				},
				"label_name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"value": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	},
}

var SloEndpointLatency = map[string]*schema.Schema{
	"endpoints_monitored": {
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"latency_bucket": {
		Type:     schema.TypeString,
		Required: true,
	},
	"additional_promql_filters": {
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Required: true,
				},
				"label_name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"value": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	},
}

var SloCustomIndicator = map[string]*schema.Schema{
	"good_query_template": {
		Type:         schema.TypeString,
		Optional:     true,
		ExactlyOneOf: []string{"good_query_template", "bad_query_template"},
	},
	"bad_query_template": {
		Type:         schema.TypeString,
		Optional:     true,
		ExactlyOneOf: []string{"good_query_template", "bad_query_template"},
	},
	"total_query_template": {
		Type:     schema.TypeString,
		Required: true,
	},
}

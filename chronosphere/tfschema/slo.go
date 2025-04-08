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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	sliTypes                      = []string{"sli.0.custom_indicator", "sli.0.endpoint_availability", "sli.0.endpoint_latency"}
	customIndicatorQueryTemplates = []string{"sli.0.custom_indicator.0.good_query_template", "sli.0.custom_indicator.0.bad_query_template"}
)

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
	"signal_grouping": SignalGrouping,
	"annotations": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"labels": {
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
	"burn_rate_alerting_config": {
		Type:     schema.TypeList,
		Computed: true,
		Optional: true,
		Elem: &schema.Resource{
			Schema: BurnRateDefinition,
		},
	},
}

var SLI = map[string]*schema.Schema{
	"custom_indicator": {
		Type:         schema.TypeList,
		Optional:     true,
		MaxItems:     1,
		ExactlyOneOf: sliTypes,
		Elem: &schema.Resource{
			Schema: SloCustomIndicator,
		},
	},
	"custom_dimension_labels": {
		Type:     schema.TypeList,
		Optional: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
	"additional_promql_filters": SLOAdditionalPromQLFilters,
}

var SloCustomIndicator = map[string]*schema.Schema{
	"good_query_template": {
		Type:         schema.TypeString,
		Optional:     true,
		ExactlyOneOf: customIndicatorQueryTemplates,
	},
	"bad_query_template": {
		Type:         schema.TypeString,
		Optional:     true,
		ExactlyOneOf: customIndicatorQueryTemplates,
	},
	"total_query_template": {
		Type:     schema.TypeString,
		Required: true,
	},
}

var SLOAdditionalPromQLFilters = &schema.Schema{
	Type:     schema.TypeSet,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": Enum{
				Value:    enum.PromQLMatcherType.ToStrings(),
				Required: true,
			}.Schema(),
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	},
}

var BurnRateDefinition = map[string]*schema.Schema{
	"window": {
		Type:     schema.TypeString,
		Required: true,
	},
	"budget": {
		Type:     schema.TypeFloat,
		Required: true,
	},
	"severity": {
		Type:     schema.TypeString,
		Required: true,
	},
	"labels": {
		Type:     schema.TypeMap,
		Optional: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
}

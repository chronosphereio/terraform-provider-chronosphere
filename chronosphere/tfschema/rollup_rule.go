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
		Type:     schema.TypeString,
		Optional: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"slug": {
		Type: schema.TypeString,
		// This is intentionally different from other resources due to unique issues with Rollup Rules
		Required: true,
		ForceNew: true,
	},
	"filter": Filter{
		KVDelimiter: aggregationfilter.RollupRuleDelimiter,
	}.Schema(),
	"new_metric": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"metric_type": Enum{
		Value:    enum.MetricType.ToStrings(),
		Required: true,
	}.Schema(),
	"aggregation": Enum{
		Value:    enum.AggregationType.ToStrings(),
		Optional: true,
	}.Schema(),
	"storage_policies": {
		Type:     schema.TypeList,
		Optional: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"resolution": Duration{
					Required: true,
				}.Schema(),
				"retention": Duration{
					Required: true,
				}.Schema(),
			},
		},
		// When no policies are specified, the server-side will set the defaults.
		Computed:   true,
		Deprecated: "use `interval` instead",
	},
	"interval": {
		Type:     schema.TypeString,
		Optional: true,
		// When no interval is specified, a server-side default is used.
		Computed:      true,
		ConflictsWith: []string{"storage_policies"},
	},
	"group_by": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"exclude_by": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"metric_type_tag": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
	"drop_raw": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
	"permissive": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
	"mode": Enum{
		Value:    enum.RollupModeType.ToStrings(),
		Optional: true,
	}.Schema(),
	"graphite_label_policy": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"replace": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name": {
								Type:     schema.TypeString,
								Required: true,
							},
							"new_value": {
								Type:     schema.TypeString,
								Required: true,
							},
						},
					},
				},
			},
		},
	},
}

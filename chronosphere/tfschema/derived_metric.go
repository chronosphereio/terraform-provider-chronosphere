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

var DerivedMetric = map[string]*schema.Schema{
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the derived metric. Generated from `name` if omitted. Immutable after creation.",
	},
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the derived metric. Can be changed after creation.",
	},
	"metric_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Name of the derived metric as referenced in queries. Must be unique across the system.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Free-form description of the derived metric.",
	},
	"queries": {
		Type:        schema.TypeList,
		Required:    true,
		Description: "Ordered list of selector/query pairs. When the derived metric is used, the first entry whose `selector` matches the usage's labels supplies the PromQL `query`.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"selector": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: "Label matchers that must be present on the derived metric usage for this query to be selected. If omitted, the query matches any usage.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"labels": {
								Type:        schema.TypeMap,
								Optional:    true,
								Description: "Labels that must match (key/value) on the derived metric usage for the selector to apply.",
								Elem:        &schema.Schema{Type: schema.TypeString},
							},
						},
					},
				},
				"query": {
					Type:        schema.TypeList,
					Required:    true,
					MinItems:    1,
					MaxItems:    1,
					Description: "PromQL query executed when this selector matches.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"expr": {
								Type:        schema.TypeString,
								Required:    true,
								Description: "PromQL expression for the derived metric. References declared variables using `$name` syntax (e.g. `cpu_usage{$service}`).",
							},
							"variables": {
								Type:        schema.TypeList,
								Optional:    true,
								Description: "Variables that can be substituted into `expr` at query time as label selectors.",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"name": {
											Type:        schema.TypeString,
											Required:    true,
											Description: "Variable name as referenced in `expr` (e.g. `service` for `$service`).",
										},
										"default_selector": {
											Type:        schema.TypeString,
											Required:    true,
											Description: "PromQL label selector used when no override is supplied by the derived metric usage.",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

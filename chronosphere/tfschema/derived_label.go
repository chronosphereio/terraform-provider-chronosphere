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

var DerivedLabel = map[string]*schema.Schema{
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the derived label. Generated from `name` if omitted. Immutable after creation.",
	},
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the derived label. Can be changed after creation.",
	},
	"label_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Name of the label exposed on derived series. Must be unique across the system.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Free-form description of the derived label.",
	},
	"existing_label_policy": Enum{
		Value:       enum.LabelPolicy.ToStrings(),
		Optional:    true,
		Description: "Policy controlling behavior when the target label already exists on the source series (e.g. keep, replace).",
	}.Schema(),
	"metric_label": {
		Type:        schema.TypeList,
		MaxItems:    1,
		Optional:    true,
		Description: "Derives a label for metrics, either by constructing a new value from filters or by mapping an existing label. Mutually exclusive with `span_tag`.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"constructed_label": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: "Constructs the derived label value from a list of value definitions, each gated by a filter on existing labels.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"value_definitions": {
								Type:        schema.TypeList,
								Required:    true,
								Description: "Ordered list of value definitions. The first definition whose filters match produces the derived label value.",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"value": {
											Type:        schema.TypeString,
											Required:    true,
											Description: "Value assigned to the derived label when this definition's filters match.",
										},
										"filters": {
											Type:        schema.TypeList,
											Required:    true,
											Description: "Label filters that must all match for this value definition to apply.",
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Type:        schema.TypeString,
														Required:    true,
														Description: "Name of the label to match.",
													},
													"value_glob": {
														Type:        schema.TypeString,
														Required:    true,
														Description: "Glob pattern matched against the label value.",
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
				"mapping_label": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: "Derives the label value by mapping from an existing source label, optionally translating its values.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name_mappings": {
								Type:        schema.TypeList,
								Optional:    true,
								Description: "Ordered list of name mappings. The first mapping whose filters match supplies the derived label from its `source_label`.",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"source_label": {
											Type:        schema.TypeString,
											Required:    true,
											Description: "Source label on the ingested time series to copy into the derived label.",
										},
										"filters": {
											Type:        schema.TypeList,
											Required:    true,
											Description: "Label filters that must all match for this name mapping to apply.",
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Type:        schema.TypeString,
														Required:    true,
														Description: "Name of the label to match.",
													},
													"value_glob": {
														Type:        schema.TypeString,
														Required:    true,
														Description: "Glob pattern matched against the label value.",
													},
												},
											},
										},
										"value_mappings": ValueMappingsSchema,
									},
								},
							},
							"value_mappings": ValueMappingsSchema,
						},
					},
				},
			},
		},
	},
	"span_tag": {
		Type:        schema.TypeList,
		MaxItems:    1,
		Optional:    true,
		Description: "Derives a label for trace spans by mapping from an existing span tag. Mutually exclusive with `metric_label`.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name_mappings": {
					Type:        schema.TypeList,
					Optional:    true,
					Description: "Ordered list of name mappings. The first mapping that matches supplies the derived label from its `source_tag`.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"source_tag": {
								Type:        schema.TypeString,
								Required:    true,
								Description: "Source span tag name to copy into the derived label.",
							},
						},
					},
				},
			},
		},
	},
}

var ValueMappingsSchema = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	Description: "Translations from source label values to a normalized target value. Each entry maps a set of source globs to a single target.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"target_value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Value to assign on the derived label when any `source_value_globs` matches.",
			},
			"source_value_globs": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Glob patterns matched against the source label value. A match maps the value to `target_value`.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	},
}

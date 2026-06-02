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
)

var GcpMetricsIntegration = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the GCP metrics integration.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Stable identifier for the integration. Generated from `name` if omitted. Immutable after creation.",
	},
	"service_account": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Google Cloud service account that Chronosphere impersonates to read metrics.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"client_email": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Email address of the Google Cloud service account to impersonate for authentication.",
				},
			},
		},
		MinItems: 1,
		MaxItems: 1,
	},
	"metric_groups": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Groups of Google Cloud metrics to ingest. Each group targets a specific project and set of metric prefixes.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"project_id": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Google Cloud project ID to read metrics from. The configured service account must have access.",
				},
				"prefixes": {
					Type:        schema.TypeList,
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Description: "List of Google Cloud metric prefixes to ingest (e.g. `compute.googleapis.com/`).",
				},
				"filters": {
					Type:        schema.TypeList,
					Optional:    true,
					Description: "Label filters applied to metrics in this group. All filters must match for a metric to be ingested.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Optional:    true,
								Description: "Label name to filter on.",
							},
							"value_glob": {
								Type:        schema.TypeString,
								Optional:    true,
								Description: "Value pattern using glob syntax (e.g. `prod-*`). An exact match is applied when no glob characters are present.",
							},
							"context": {
								Type:        schema.TypeString,
								Optional:    true,
								Description: "Label context, e.g. resource vs. metric label. See the Chronosphere GCP integration documentation for accepted values.",
							},
						},
					},
				},
				"rollup_rules": {
					Type:        schema.TypeList,
					Optional:    true,
					Description: "Server-side aggregation rules applied to metrics in this group before they are stored.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"metric_name": {
								Type:        schema.TypeString,
								Optional:    true,
								Description: "Fully-qualified Google Cloud metric name the rollup rule targets (e.g. `cloudsql.googleapis.com/database/uptime`).",
							},
							"aggregation": {
								Type:        schema.TypeString,
								Optional:    true,
								Description: "Aggregation function applied across the dropped labels (e.g. sum, max).",
							},
							"label_policy": {
								Type:        schema.TypeList,
								Optional:    true,
								MaxItems:    1,
								Description: "Specifies which labels to preserve during aggregation. Labels not listed are dropped.",
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"keep": {
											Type:        schema.TypeList,
											Optional:    true,
											Description: "Labels to retain after aggregation.",
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"name": {
														Type:        schema.TypeString,
														Optional:    true,
														Description: "Name of the label to keep.",
													},
													"context": {
														Type:        schema.TypeString,
														Optional:    true,
														Description: "Label context, e.g. resource vs. metric label.",
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
			},
		},
	},
}

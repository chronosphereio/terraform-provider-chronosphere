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

var AzureMetricsIntegration = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the Azure metrics integration.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the integration. Generated from `name` if omitted. Immutable after creation.",
	},
	"principal": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Azure managed identity principal used to authenticate with Azure Monitor.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"tenant_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "ID of the Azure tenant that hosts the managed identity principal.",
				},
				"client_id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "OAuth2 client ID of the managed identity principal.",
				},
			},
		},
		MinItems: 1,
		MaxItems: 1,
	},
	"scrape_config": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Scope of Azure subscriptions, locations, and resource types from which to ingest metrics.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"subscription_ids": {
					Type:        schema.TypeList,
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Description: "Azure subscription IDs to target. Leave empty to scrape from all subscriptions accessible to the principal.",
				},
				"locations": {
					Type:        schema.TypeList,
					Optional:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Description: "Azure locations (regions) to ingest from, applied across all subscriptions. Leave empty for all locations.",
				},
				"resource_type": {
					Type:        schema.TypeList,
					Optional:    true,
					Description: "Azure resource types to scrape metrics from. Each entry can constrain the set of metric names to a subset.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"name": {
								Type:        schema.TypeString,
								Optional:    true,
								Description: "Azure resource type identifier (e.g. `Microsoft.Compute/virtualMachines`).",
							},
							"metric_names": {
								Type:        schema.TypeList,
								Optional:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
								Description: "Metric names to ingest for this resource type. Leave empty for all metrics.",
							},
						},
					},
				},
			},
		},
		MinItems: 1,
		MaxItems: 1,
	},
	"count_metrics_enabled": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "If true, enables Azure count metrics for the configured resources.",
	},
	"usage_metrics_enabled": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "If true, enables collection of Azure usage metrics under this principal (Microsoft.Compute, Microsoft.Network, Microsoft.Storage).",
	},
	"propagate_tags": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "If true, propagates Azure resource, group, and subscription tags as metric labels.",
	},
}

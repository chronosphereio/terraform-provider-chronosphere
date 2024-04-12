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

var GcpMetricsIntegration = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Optional: false,
	},
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"service_account": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"client_email": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
		MinItems: 1,
		MaxItems: 1,
	},
	"metric_groups": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"project_id": {
					Type:     schema.TypeString,
					Required: true,
				},
				"prefixes": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	},
}

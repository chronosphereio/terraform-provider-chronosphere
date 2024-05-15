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

var OtelMetricsIngestion = map[string]*schema.Schema{
	"resource_attributes": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"flatten_mode": Enum{
					Value:    enum.ResourceAttributesFlattenMode.ToStrings(),
					Optional: true,
				}.Schema(),
				"filter_mode": Enum{
					Value:    enum.ResourceAttributesFilterMode.ToStrings(),
					Optional: true,
				}.Schema(),
				"exclude_keys": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"generate_target_info": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
	},
}

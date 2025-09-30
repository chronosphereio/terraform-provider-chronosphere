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

var DropRule = map[string]*schema.Schema{
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"active": {
		Type:       schema.TypeBool,
		Optional:   true,
		Default:    false,
		Deprecated: "use `mode` instead",
	},
	"mode": Enum{
		Value:    enum.DropRuleModeType.ToStrings(),
		Optional: true,
		Required: false,
	}.Schema(),
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"query": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		MinItems: 1,
		Required: true,
	},
	"conditional_drop": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"drop_nan_value": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"activated_drop_duration": Duration{
		Optional: true,
	}.Schema(),
	"rate_limit_threshold": {
		Type:             schema.TypeFloat,
		Optional:         true,
		ValidateDiagFunc: float64RangeValidator(0, 100.0),
	},
	"value_based_drop": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"target_drop_value": {
					Type:     schema.TypeFloat,
					Required: true,
				},
			},
		},
	},
}

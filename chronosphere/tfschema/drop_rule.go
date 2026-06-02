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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the drop rule. Generated from `name` if omitted. Immutable after creation.",
	},
	"active": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Deprecated:  "use `mode` instead",
		Description: "Whether the drop rule is active. Deprecated: use `mode` instead.",
	},
	"mode": Enum{
		Value:       enum.DropRuleModeType.ToStrings(),
		Optional:    true,
		Default:     "ENABLED",
		Description: "Drop rule mode controlling whether it is enabled, disabled, or in a preview state. Defaults to `ENABLED`.",
	}.Schema(),
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the drop rule. Can be changed after creation.",
	},
	"query": {
		Type:        schema.TypeList,
		Elem:        &schema.Schema{Type: schema.TypeString},
		MinItems:    1,
		Required:    true,
		Description: "List of label filter queries that select which metrics to drop. A metric is dropped if it matches all filters in any one query.",
	},
	"conditional_drop": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "If `true`, the drop only activates when the configured `rate_limit_threshold` is exceeded.",
	},
	"drop_nan_value": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "If `true`, drops data points whose value is NaN, including any associated staleness markers.",
	},
	"activated_drop_duration": Duration{
		Optional:    true,
		Description: "Once a conditional drop activates, how long it stays activated before re-checking against `rate_limit_threshold`.",
	}.Schema(),
	"rate_limit_threshold": {
		Type:             schema.TypeFloat,
		Optional:         true,
		ValidateDiagFunc: float64RangeValidator(0, 100.0),
		Description:      "Percentage of the licensed metrics limit (0-100) at which a conditional drop activates.",
	},
	"value_based_drop": {
		Type:        schema.TypeList,
		MaxItems:    1,
		Optional:    true,
		Description: "Configuration for dropping data points whose value matches a target.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"target_drop_value": {
					Type:        schema.TypeFloat,
					Required:    true,
					Description: "Data point value at which matching points are dropped.",
				},
			},
		},
	},
}

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

var MatcherListSchema = &schema.Schema{
	Type:        schema.TypeList,
	Required:    true,
	MinItems:    1,
	Description: "List of label matchers used to select a subset of series.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Label name to match.",
			},
			"type": Enum{
				Value:       enum.MatcherType.ToStrings(),
				Required:    true,
				Description: "Match operator: one of `=`, `!=`, `=~` (regex), `!~` (regex negation).",
			}.Schema(),
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Label value (or regex pattern, for regex matchers) to match against.",
			},
		},
	},
}

// SignalGrouping is used by both Monitor and SLO
var SignalGrouping = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	MaxItems:    1,
	Description: "Controls how individual time series are grouped into signals for alerting purposes.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"label_names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "Labels to group by. Series sharing the same values for these labels produce one signal. Defaults to no grouping (one signal per series).",
			},
			"signal_per_series": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, treat each individual series as its own signal. Mutually exclusive with `label_names`.",
			},
		},
	},
}

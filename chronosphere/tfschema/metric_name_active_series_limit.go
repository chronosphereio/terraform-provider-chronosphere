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

var MetricNameActiveSeriesLimit = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the metric name active series limit.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the metric name active series limit. Generated from `name` if omitted. Immutable after creation.",
	},
	"metric_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "The exact `__name__` this limit applies to. Only one limit may exist per metric name.",
	},
	"max_active_series": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "The total active time series allowed for this metric name over the active time series rolling window. Must be greater than zero.",
	},
}

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

var MetricNameDefaultActiveSeriesLimit = map[string]*schema.Schema{
	"default_max_active_standard_series": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "The default active standard (non-histogram) time series allowed for each metric name over the active time series rolling window. If zero or unset, the Chronosphere-managed default limit is used. A `chronosphere_metric_name_active_series_limit` overrides this default for its metric name. Can't exceed the Chronosphere-managed default limit.",
	},
	"default_max_active_histogram_series": {
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "The default active histogram time series allowed for each metric name over the active time series rolling window. If zero or unset, the Chronosphere-managed default limit is used. A `chronosphere_metric_name_active_series_limit` overrides this default for its metric name. Can't exceed the Chronosphere-managed default limit.",
	},
}

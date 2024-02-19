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

package enum

import (
	configv1 "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// MetricType is an enum.
var MetricType = newEnum("MetricType", []value[
	string,
	configv1.RollupRuleMetricType,
]{
	{
		legacy:    "UNKNOWN_MT",
		isDefault: true,
	},
	{
		legacy: "COUNTER",
		v1:     configv1.RollupRuleMetricTypeCOUNTER,
		alias:  "COUNTER",
	},
	{
		legacy: "GAUGE",
		v1:     configv1.RollupRuleMetricTypeGAUGE,
		alias:  "GAUGE",
	},
	{
		legacy: "DELTA",
		v1:     configv1.RollupRuleMetricTypeDELTA,
		alias:  "DELTA",
	},
	{
		legacy: "DISTRIBUTION",
		v1:     configv1.RollupRuleMetricTypeDISTRIBUTION,
		alias:  "DISTRIBUTION",
	},
	{
		v1:    configv1.RollupRuleMetricTypeCUMULATIVEEXPONENTIALHISTOGRAM,
		alias: "CUMULATIVE_EXPONENTIAL_HISTOGRAM",
	},
	{
		v1:    configv1.RollupRuleMetricTypeMEASUREMENT,
		alias: "MEASUREMENT",
	},
})

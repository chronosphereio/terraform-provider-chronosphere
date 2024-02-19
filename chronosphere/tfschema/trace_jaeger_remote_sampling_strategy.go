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

var TraceJaegerRemoteSamplingStrategy = map[string]*schema.Schema{
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"service_name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"applied_strategy": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: appliedStrategySchema,
		},
	},
}

var appliedStrategySchema = map[string]*schema.Schema{
	"per_operation_strategies": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: perOperationStrategiesSchema,
		},
	},
	"probabilistic_strategy": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: probabilisticStrategySchema,
		},
	},
	"rate_limiting_strategy": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: rateLimitingStrategySchema,
		},
	},
}

var probabilisticStrategySchema = map[string]*schema.Schema{
	"sampling_rate": {
		Type:             schema.TypeFloat,
		Required:         true,
		ValidateDiagFunc: float64RangeValidator(0, 1.0),
	},
}

var rateLimitingStrategySchema = map[string]*schema.Schema{
	"max_traces_per_second": {
		Type:     schema.TypeInt,
		Required: true,
	},
}

var perOperationStrategySchema = map[string]*schema.Schema{
	"operation": {
		Type:     schema.TypeString,
		Required: true,
	},
	"probabilistic_strategy": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: probabilisticStrategySchema,
		},
	},
}

var perOperationStrategiesSchema = map[string]*schema.Schema{
	"default_sampling_rate": {
		Type:             schema.TypeFloat,
		Required:         true,
		ValidateDiagFunc: float64RangeValidator(0, 1.0),
	},
	"default_lower_bound_traces_per_second": {
		Type:     schema.TypeFloat,
		Optional: true,
	},
	"default_upper_bound_traces_per_second": {
		Type:     schema.TypeFloat,
		Optional: true,
	},
	"per_operation_strategies": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: perOperationStrategySchema,
		},
	},
}

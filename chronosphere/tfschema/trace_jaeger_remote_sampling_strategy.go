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
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the Jaeger remote sampling strategy. Generated from `name` if omitted. Immutable after creation.",
	},
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the Jaeger remote sampling strategy.",
	},
	"service_name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Value of `service.name` the strategy applies to. Jaeger SDK clients reporting under this service receive this sampling configuration.",
	},
	"applied_strategy": {
		Type:        schema.TypeList,
		Required:    true,
		MinItems:    1,
		MaxItems:    1,
		Description: "Sampling strategy returned to the Jaeger client. Exactly one of `probabilistic_strategy`, `rate_limiting_strategy`, or `per_operation_strategies` must be set.",
		Elem: &schema.Resource{
			Schema: appliedStrategySchema,
		},
	},
}

var appliedStrategySchema = map[string]*schema.Schema{
	"per_operation_strategies": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Per-operation sampling configuration with a service-wide default and optional per-operation overrides.",
		Elem: &schema.Resource{
			Schema: perOperationStrategiesSchema,
		},
	},
	"probabilistic_strategy": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Probabilistic sampling: each trace is sampled with a fixed probability.",
		Elem: &schema.Resource{
			Schema: probabilisticStrategySchema,
		},
	},
	"rate_limiting_strategy": {
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Rate-limiting sampling: cap the number of sampled traces per second using a leaky bucket.",
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
		Description:      "Probability in the range `[0.0, 1.0]` that any given trace is sampled. `0` samples no traces, `1` samples every trace.",
	},
}

var rateLimitingStrategySchema = map[string]*schema.Schema{
	"max_traces_per_second": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Maximum number of traces to sample per second for the service.",
	},
}

var perOperationStrategySchema = map[string]*schema.Schema{
	"operation": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Span operation (span name) this override applies to.",
	},
	"probabilistic_strategy": {
		Type:        schema.TypeList,
		Required:    true,
		MinItems:    1,
		MaxItems:    1,
		Description: "Probabilistic sampling configuration applied to spans whose operation matches.",
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
		Description:      "Service-wide sampling probability in the range `[0.0, 1.0]` applied when no per-operation override matches.",
	},
	"default_lower_bound_traces_per_second": {
		Type:        schema.TypeFloat,
		Optional:    true,
		Description: "Minimum number of traces per second sampled for any operation in the service, even when the probabilistic rate would yield fewer.",
	},
	"default_upper_bound_traces_per_second": {
		Type:        schema.TypeFloat,
		Optional:    true,
		Description: "Maximum number of traces per second sampled for any operation in the service, regardless of matching per-operation strategy.",
	},
	"per_operation_strategies": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Operation-specific overrides that take precedence over `default_sampling_rate`.",
		Elem: &schema.Resource{
			Schema: perOperationStrategySchema,
		},
	},
}

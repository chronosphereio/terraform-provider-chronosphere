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

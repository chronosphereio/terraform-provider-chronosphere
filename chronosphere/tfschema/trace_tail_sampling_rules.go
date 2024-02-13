package tfschema

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var TraceTailSamplingRules = map[string]*schema.Schema{
	"default_sample_rate": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"enabled": {
					Type:     schema.TypeBool,
					Optional: true,
				},
				"sample_rate": SampleRateSchema,
			},
		},
	},
	"rules": rulesSchema,
}

var rulesSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true, // this can be set to required: true once all callers are upgraded
			},
			"system_name": {
				Type:     schema.TypeString,
				Optional: true, // this can be set to required: true once all callers are upgraded
			},
			"sample_rate": SampleRateSchema,
			"filter": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trace": tailSamplingTraceFilterSchema,
						"span":  tailSamplingSpanFilterSchema,
					},
				},
			},
		},
	},
}

var SampleRateSchema = &schema.Schema{
	Type:             schema.TypeFloat,
	Required:         true,
	ValidateDiagFunc: float64RangeValidator(0, 1.0),
}

var tailSamplingTraceFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	MaxItems: 1,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": TailSamplingDurationFilterSchema,
			"error":    TailSamplingBoolFilterSchema,
		},
	},
}

var tailSamplingSpanFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_type": Enum{
				Value:    enum.SpanFilterMatchType.ToStrings(),
				Optional: true,
			}.Schema(),
			"service":          TailSamplingStringFilterSchema,
			"operation":        TailSamplingStringFilterSchema,
			"parent_service":   TailSamplingStringFilterSchema,
			"parent_operation": TailSamplingStringFilterSchema,
			"duration":         TailSamplingDurationFilterSchema,
			"error":            TailSamplingBoolFilterSchema,
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value":         TailSamplingStringFilterSchema,
						"numeric_value": TailSamplingNumericFilterSchema,
					},
				},
			},
			"span_count": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	},
}

var TailSamplingStringFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	MaxItems: 1,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"match": Enum{
				Value:    enum.StringFilterMatchType.ToStrings(),
				Optional: true,
			}.Schema(),
		},
	},
}

var TailSamplingNumericFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	MaxItems: 1,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeFloat,
				Required: true,
			},
			"comparison": Enum{
				Value:    enum.NumericFilterComparisonType.ToStrings(),
				Required: true,
			}.Schema(),
		},
	},
}

var TailSamplingDurationFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	MaxItems: 1,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"min_secs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_secs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	},
}

var TailSamplingBoolFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	MaxItems: 1,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	},
}

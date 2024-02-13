package tfschema

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var DropRule = map[string]*schema.Schema{
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"active": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  false,
	},
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

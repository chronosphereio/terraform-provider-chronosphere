package tfschema

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
)

var LogControlConfig = map[string]*schema.Schema{
	"rules": {
		Type:     schema.TypeList,
		Elem:     logControlRuleResource,
		Optional: true,
	},
}

var logControlRuleResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"mode": Enum{
			Value:    enum.LogControlRuleMode.ToStrings(),
			Optional: true,
		}.Schema(),
		"filter": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "LogQL query to select logs. Only matching logs will have control action applied.",
		},
		"type": Enum{
			Value:    enum.LogControlRuleType.ToStrings(),
			Required: true,
		}.Schema(),
		"sample": {
			Type:     schema.TypeList,
			Elem:     logControlRuleSampleResource,
			Optional: true,
			MaxItems: 1,
		},
		"drop_field": {
			Type:     schema.TypeList,
			Elem:     logControlRuleDropFieldResource,
			Optional: true,
			MaxItems: 1,
		},
	},
}

var logControlRuleSampleResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"rate": {
			Type:        schema.TypeFloat,
			Required:    true,
			Description: "Percentage of matching logs to keep. Must be in the range [0, 1].",
		},
	},
}

var logControlRuleDropFieldResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"field_regex": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Regular expression to match the field(s) to drop.",
		},
		"parent_path": {
			Type:     schema.TypeList,
			Elem:     logFieldPathResource,
			Optional: true,
			MaxItems: 1,
		},
	},
}

var logFieldPathResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"selector": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "LogQL Selector to indicate field path. Use 'parent[child]' syntax to indicate nesting.",
		},
	},
}

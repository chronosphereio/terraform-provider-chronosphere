package tfschema

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
)

const maxLogParsers = 10

var LogIngestConfig = map[string]*schema.Schema{
	"plaintext_parser": {
		Type:     schema.TypeList,
		Elem:     PlaintextParserSchema,
		Optional: true,
		MaxItems: maxLogParsers,
	},
	"field_parser": {
		Type:     schema.TypeList,
		Elem:     LogFieldParserSchema,
		Optional: true,
		MaxItems: maxLogParsers,
	},
}

var PlaintextParserSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"mode": Enum{
			Value:    enum.PlaintextParserMode.ToStrings(),
			Optional: true,
		}.Schema(),
		"parser": LogParserSchema,
		"drop_original": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	},
}

var LogFieldParserSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"mode": Enum{
			Value:    enum.LogFieldParserMode.ToStrings(),
			Optional: true,
		}.Schema(),
		"source":      LogFieldPathSchema,
		"destination": LogFieldPathSchema,
		"parser":      LogParserSchema,
	},
}

var LogFieldPathSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"selector": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "LogQL Selector to indicate field path. Use 'parent[child]' syntax to indicate nesting.",
			},
		},
	},
	Required: true,
	MaxItems: 1,
}

var LogParserSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"parser_type": Enum{
				Value:    enum.LogParserType.ToStrings(),
				Required: true,
			}.Schema(),
			"json_parser":      JSONLogParserSchema,
			"regex_parser":     RegexLogParserSchema,
			"key_value_parser": KeyValueLogParserSchema,
		},
	},
	Required: true,
	MaxItems: 1,
}

var JSONLogParserSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{},
	},
	Optional: true,
	MaxItems: 1,
}

var RegexLogParserSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"regex": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Re2 regex parser pattern to apply. Named capturing groups become named fields in the extracted log.",
			},
		},
	},
	Optional: true,
	MaxItems: 1,
}

var KeyValueLogParserSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pair_separator": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "String used to split each pair into its key and value.",
			},
			"delimiter": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "String used to split the input into key-value pairs.",
			},
			"trim_set": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "All leading and trailing characters contained in the trim set will be removed from keys and values.",
			},
		},
	},
	Optional: true,
	MaxItems: 1,
}

package tfschema

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
)

var LogIngestConfig = map[string]*schema.Schema{
	"plaintext_parser": {
		Type:     schema.TypeList,
		Elem:     plaintextParserResource,
		Optional: true,
	},
	"field_parser": {
		Type:     schema.TypeList,
		Elem:     logFieldParserResource,
		Optional: true,
	},
	"field_normalization": {
		Type:     schema.TypeList,
		Elem:     fieldNormalizationResource,
		Optional: true,
		MaxItems: 1,
	},
}

var plaintextParserResource = &schema.Resource{
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
		"keep_original": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	},
}

var logFieldParserResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"mode": Enum{
			Value:    enum.LogFieldParserMode.ToStrings(),
			Optional: true,
		}.Schema(),
		"source": {
			Type:     schema.TypeList,
			Elem:     LogFieldPathResource,
			Required: true,
			MaxItems: 1,
		},
		"destination": {
			Type:     schema.TypeList,
			Elem:     LogFieldPathResource,
			Optional: true,
			MaxItems: 1,
		},
		"parser": LogParserSchema,
	},
}

var LogFieldPathResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"selector": {
			Type:     schema.TypeString,
			Required: true,
		},
	},
}

var LogParserSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"parser_type": Enum{
				Value:    enum.LogParserType.ToStrings(),
				Required: true,
			}.Schema(),
			"regex_parser":     RegexLogParserSchema,
			"key_value_parser": KeyValueLogParserSchema,
		},
	},
	Required: true,
	MaxItems: 1,
}

var RegexLogParserSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"regex": {
				Type:     schema.TypeString,
				Required: true,
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
				Type:     schema.TypeString,
				Required: true,
			},
			"delimiter": {
				Type:     schema.TypeString,
				Required: true,
			},
			"trim_set": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	},
	Optional: true,
	MaxItems: 1,
}

var fieldNormalizationResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"custom_field_normalization": {
			Type:     schema.TypeList,
			Elem:     NamedStringNormalizationResource,
			Optional: true,
		},
		"message":  LogIngestConfigStringNormalizationSchema,
		"service":  LogIngestConfigStringNormalizationSchema,
		"severity": LogIngestConfigStringNormalizationSchema,
		"timestamp": {
			Type:     schema.TypeList,
			Elem:     timestampNormalizationResource,
			Optional: true,
			MaxItems: 1,
		},
	},
}

var timestampNormalizationResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"source": {
			Type:     schema.TypeList,
			Elem:     LogFieldPathResource,
			Optional: true,
		},
	},
}

var LogIngestConfigStringNormalizationSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"default_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sanitize_patterns": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"source": {
				Type:     schema.TypeList,
				Elem:     LogFieldPathResource,
				Optional: true,
			},
			"value_map": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	},
}

var NamedStringNormalizationResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"normalization": LogIngestConfigStringNormalizationSchema,
		"target": {
			Type:     schema.TypeString,
			Optional: true,
		},
	},
}

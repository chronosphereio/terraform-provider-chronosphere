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
			Elem:     LogFieldSelectorResource,
			Required: true,
			MaxItems: 1,
		},
		"destination": {
			Type:     schema.TypeList,
			Elem:     LogFieldSelectorResource,
			Optional: true,
			MaxItems: 1,
		},
		"parser": LogParserSchema,
	},
}

// LogFieldSelectorResource is shared across log configs
var LogFieldSelectorResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"selector": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "LogQL Selector to indicate field path. Use 'parent[child]' syntax to indicate nesting.",
		},
	},
}

// Shared schemas for field normalization - these are used across multiple resources
// to ensure the generator recognizes them as shared types
var StringNormalizationSchema = &schema.Schema{
	Type:     schema.TypeList,
	Elem:     StringNormalizationResource,
	Optional: true,
	MaxItems: 1,
}

var NamedStringNormalizationSchema = &schema.Schema{
	Type:     schema.TypeList,
	Elem:     NamedStringNormalizationResource,
	Optional: true,
	MaxItems: 1,
}

// NamedStringNormalizationListSchema is for lists of named string normalizations (no MaxItems)
var NamedStringNormalizationListSchema = &schema.Schema{
	Type:     schema.TypeList,
	Elem:     NamedStringNormalizationResource,
	Optional: true,
}

var TimestampNormalizationSchema = &schema.Schema{
	Type:     schema.TypeList,
	Elem:     TimestampNormalizationResource,
	Optional: true,
	MaxItems: 1,
}

var LogFieldSelectorSchema = &schema.Schema{
	Type:     schema.TypeList,
	Elem:     LogFieldSelectorResource,
	Optional: true,
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

var fieldNormalizationResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"custom_field_normalization": NamedStringNormalizationListSchema,
		"message":                    StringNormalizationSchema,
		"primary_key":                NamedStringNormalizationSchema,
		"severity":                   StringNormalizationSchema,
		"timestamp":                  TimestampNormalizationSchema,
	},
}

// StringNormalizationResource is shared across log configs
var StringNormalizationResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"default_value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Default value to use when no source fields contain values.",
		},
		"sanitize_patterns": {
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Optional:    true,
			Description: "Optional regex patterns to extract and sanitize values. Each pattern must have exactly one capturing group that will be used as the result.",
		},
		"source": {
			Type:        schema.TypeList,
			Elem:        LogFieldSelectorResource,
			Optional:    true,
			Description: "List of field paths to check for values, in priority order. The first non-empty value found will be used.",
		},
		"value_map": {
			Type:        schema.TypeMap,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Optional:    true,
			Description: "Optional mapping to normalize values. For example: {\"warn\": \"WARNING\", \"err\": \"ERROR\"} to standardize severity levels.",
		},
	},
}

// NamedStringNormalizationResource is shared across log configs
var NamedStringNormalizationResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"normalization": {
			Type:     schema.TypeList,
			Elem:     StringNormalizationResource,
			Optional: true,
			MaxItems: 1,
		},
		"target": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the target field where the normalized value will be stored.",
		},
	},
}

// TimestampNormalizationResource is shared across log configs
var TimestampNormalizationResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"source": {
			Type:        schema.TypeList,
			Elem:        LogFieldSelectorResource,
			Optional:    true,
			Description: "List of field paths to check for timestamp values, in priority order. Common fields include \"timestamp\", \"@timestamp\", \"time\", \"datetime\".",
		},
	},
}

package tfschema

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
)

var LogIngestConfig = map[string]*schema.Schema{
	"plaintext_parser": {
		Type:        schema.TypeList,
		Elem:        plaintextParserResource,
		Optional:    true,
		Description: "Parsers applied to plaintext logs as they enter the ingest pipeline. The first parser that matches a log is used.",
	},
	"field_parser": {
		Type:        schema.TypeList,
		Elem:        logFieldParserResource,
		Optional:    true,
		Description: "Parsers applied to specific fields within structured logs (or to fields produced by a plaintext parser).",
	},
	"field_normalization": {
		Type:        schema.TypeList,
		Elem:        fieldNormalizationResource,
		Optional:    true,
		MaxItems:    1,
		Description: "Field normalization rules that map and standardize well-known fields (timestamp, severity, message, service) across log formats. Runs after parsing.",
	},
}

var plaintextParserResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the parser. Must be unique within the configuration.",
		},
		"mode": Enum{
			Value:       enum.PlaintextParserMode.ToStrings(),
			Optional:    true,
			Description: "Mode that controls how the parser matches incoming plaintext logs.",
		}.Schema(),
		"parser": LogParserSchema,
		"keep_original": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "If `true`, the original log is retained after parsing and stored under the `plaintext_log` key. Defaults to `false`.",
		},
	},
}

var logFieldParserResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"mode": Enum{
			Value:       enum.LogFieldParserMode.ToStrings(),
			Optional:    true,
			Description: "Mode that controls when the field parser runs on incoming logs.",
		}.Schema(),
		"source": {
			Type:        schema.TypeList,
			Elem:        LogFieldPathResource,
			Required:    true,
			MaxItems:    1,
			Description: "Path of the field to parse.",
		},
		"destination": {
			Type:        schema.TypeList,
			Elem:        LogFieldPathResource,
			Optional:    true,
			MaxItems:    1,
			Description: "Path to write the parsed output to. If omitted, parsed fields are written at the root.",
		},
		"parser": LogParserSchema,
	},
}

var LogFieldPathResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"selector": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Field path selector. Use `parent[child]` syntax to indicate nesting.",
		},
	},
}

var LogParserSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"parser_type": Enum{
				Value:       enum.LogParserType.ToStrings(),
				Required:    true,
				Description: "Type of parser to apply. Determines which of `regex_parser`, `key_value_parser`, or `grok_parser` must be set.",
			}.Schema(),
			"regex_parser":     RegexLogParserSchema,
			"key_value_parser": KeyValueLogParserSchema,
			"grok_parser":      GrokLogParserSchema,
		},
	},
	Required:    true,
	MaxItems:    1,
	Description: "Parser configuration. Exactly one of `regex_parser`, `key_value_parser`, or `grok_parser` must be set, matching `parser_type`.",
}

var RegexLogParserSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"regex": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "RE2 regular expression pattern. Named capturing groups become named fields in the extracted log.",
			},
		},
	},
	Optional:    true,
	MaxItems:    1,
	Description: "Regex parser configuration. Only set when `parser_type` is `REGEX`.",
}

var KeyValueLogParserSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pair_separator": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "String used to split each pair into a key and value.",
			},
			"delimiter": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "String used to split the input into individual key/value pairs.",
			},
			"trim_set": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unicode code points to trim from the beginning and end of each key and value.",
			},
		},
	},
	Optional:    true,
	MaxItems:    1,
	Description: "Key/value parser configuration. Only set when `parser_type` is `KEY_VALUE`. Duplicate keys keep the first occurrence.",
}

var GrokLogParserSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pattern": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Grok pattern to apply. Named capture groups become named fields in the extracted log.",
			},
		},
	},
	Optional:    true,
	MaxItems:    1,
	Description: "Grok parser configuration. Only set when `parser_type` is `GROK`.",
}

var fieldNormalizationResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"custom_field_normalization": {
			Type:        schema.TypeList,
			Elem:        NamedStringNormalizationResource,
			Optional:    true,
			Description: "Normalization rules for additional custom fields. These fields are not indexed; use them for things like environment, region, or user ID.",
		},
		"message":  LogIngestConfigStringNormalizationSchema,
		"service":  LogIngestConfigStringNormalizationSchema,
		"severity": LogIngestConfigStringNormalizationSchema,
		"timestamp": {
			Type:        schema.TypeList,
			Elem:        timestampNormalizationResource,
			Optional:    true,
			MaxItems:    1,
			Description: "Normalization rule for the well-known `timestamp` field.",
		},
	},
}

var timestampNormalizationResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"source": {
			Type:        schema.TypeList,
			Elem:        LogFieldPathResource,
			Optional:    true,
			Description: "Ordered list of field paths to check for timestamp values. The first non-empty value found is used. Common sources include `timestamp`, `@timestamp`, `time`, and `datetime`.",
		},
	},
}

var LogIngestConfigStringNormalizationSchema = &schema.Schema{
	Type:        schema.TypeList,
	Optional:    true,
	MaxItems:    1,
	Description: "Rule that extracts and transforms a string value from a log field, with optional regex sanitization, default value, and value mapping.",
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"default_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Value to use when no source field contains a value.",
			},
			"sanitize_patterns": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Regex patterns used to extract and sanitize the value. Each pattern must have exactly one capturing group, whose contents are used as the result.",
			},
			"source": {
				Type:        schema.TypeList,
				Elem:        LogFieldPathResource,
				Optional:    true,
				Description: "Ordered list of field paths to check for values. The first non-empty value found is used.",
			},
			"value_map": {
				Type:        schema.TypeMap,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Optional mapping that normalizes raw values to canonical ones (for example, `warn` to `WARNING`).",
			},
		},
	},
}

var NamedStringNormalizationResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"normalization": LogIngestConfigStringNormalizationSchema,
		"target": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the target field where the normalized value is stored.",
		},
	},
}

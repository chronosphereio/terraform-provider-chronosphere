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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
)

var LogControlConfig = map[string]*schema.Schema{
	"rules": {
		Type:        schema.TypeList,
		Elem:        logControlRuleResource,
		Optional:    true,
		Description: "Ordered list of log control rules applied to the log ingest pipeline. Rules are evaluated in order.",
	},
}

var logControlRuleResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "User-defined name for the control rule.",
		},
		"mode": Enum{
			Value:       enum.LogControlRuleMode.ToStrings(),
			Optional:    true,
			Description: "Execution mode for the rule (for example, `ENABLED` or `DISABLED`).",
		}.Schema(),
		"filter": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Log query filter that selects matching logs. The control action applies only to logs that match.",
		},
		"type": Enum{
			Value:       enum.LogControlRuleType.ToStrings(),
			Optional:    true,
			Description: "Type of control action this rule performs. Exactly one of the matching action blocks (`sample`, `drop_field`, `emit_metrics`, `replace_field`, `parse_field`) must be configured.",
		}.Schema(),
		"sample": {
			Type:        schema.TypeList,
			Elem:        logControlRuleSampleResource,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for the `SAMPLE_LOGS` action, which keeps a fraction of matching logs.",
		},
		"drop_field": {
			Type:        schema.TypeList,
			Elem:        logControlRuleDropFieldResource,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for the `DROP_FIELD` action, which removes fields from matching logs.",
		},
		"emit_metrics": {
			Type:        schema.TypeList,
			Elem:        logControlRuleEmitMetricsResource,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for the `EMIT_METRICS` action, which derives Prometheus metrics from matching logs.",
		},
		"replace_field": {
			Type:        schema.TypeList,
			Elem:        logControlRuleReplaceFieldResource,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for the `REPLACE_FIELD` action, which rewrites field values in matching logs.",
		},
		"parse_field": {
			Type:        schema.TypeList,
			Elem:        logControlRuleParseFieldResource,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for the `PARSE_FIELD` action, which parses a field with a regex, key/value, or grok parser and writes the result to another field.",
		},
	},
}

var logControlRuleSampleResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"rate": {
			Type:        schema.TypeFloat,
			Optional:    true,
			Description: "Fraction of matching logs to keep, in the range `[0, 1]` (for example, `0.25` keeps 25%).",
		},
	},
}

var logControlRuleDropFieldResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"field_regex": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Regular expression that selects which fields to drop.",
		},
		"parent_path": LogFieldPathSchema,
	},
}

var LogFieldPathSchema = &schema.Schema{
	Type: schema.TypeList,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"selector": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Field path selector. Use `parent[child]` syntax to indicate nesting.",
			},
		},
	},
	Optional:    true,
	MaxItems:    1,
	Description: "Path to a field within a log record.",
}

var logControlRuleEmitMetricsResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"counter": {
			Type:        schema.TypeList,
			Elem:        logControlRuleEmitMetricsCounterResource,
			Optional:    true,
			MaxItems:    1,
			Description: "Emit a counter metric. Exactly one of `counter`, `gauge`, or `histogram` must be set.",
		},
		"drop_log": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "If `true`, drops the entire log after emitting the metric.",
		},
		"gauge": {
			Type:        schema.TypeList,
			Elem:        logControlRuleEmitMetricsGaugeResource,
			Optional:    true,
			MaxItems:    1,
			Description: "Emit a gauge metric. Exactly one of `counter`, `gauge`, or `histogram` must be set.",
		},
		"histogram": {
			Type:        schema.TypeList,
			Elem:        logControlRuleEmitMetricsHistogramResource,
			Optional:    true,
			MaxItems:    1,
			Description: "Emit a histogram metric. Exactly one of `counter`, `gauge`, or `histogram` must be set.",
		},
		"labels": {
			Type:        schema.TypeList,
			Elem:        logControlRuleEmitMetricsLabelResource,
			Optional:    true,
			Description: "Labels to attach to the generated metric, specified as key/value pairs mapping a Prometheus label name to a log field path.",
		},
		"mode": Enum{
			Value:       enum.EmitMetricsMetricMode.ToStrings(),
			Optional:    true,
			Description: "Metric emission mode that controls how the metric is generated from matching logs.",
		}.Schema(),
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the generated metric. Must conform to Prometheus naming conventions and be unique within the tenant.",
		},
	},
}

var logControlRuleEmitMetricsCounterResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"value": LogFieldPathSchema,
	},
}

var logControlRuleEmitMetricsGaugeResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"aggregation_type": Enum{
			Value:       enum.GaugeAggregationType.ToStrings(),
			Optional:    true,
			Description: "How multiple values are aggregated into the emitted gauge (for example, `LAST`, `MIN`, `MAX`).",
		}.Schema(),
		"value": LogFieldPathSchema,
	},
}

var logControlRuleEmitMetricsHistogramResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"value": LogFieldPathSchema,
	},
}

var logControlRuleEmitMetricsLabelResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"key": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Prometheus label name to set on the emitted metric.",
		},
		"value": LogFieldPathSchema,
	},
}

var logControlRuleReplaceFieldResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"field": LogFieldPathSchema,
		"mapped_value": {
			Type:        schema.TypeList,
			Elem:        logControlRuleReplaceFieldMappedValueResource,
			Optional:    true,
			MaxItems:    1,
			Description: "Replace field values using a key/value lookup table. Exactly one of `mapped_value` or `static_value` must be set.",
		},
		"replace_all": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "If `true`, replaces all matches. If `false`, replaces only the first match.",
		},
		"replace_mode": Enum{
			Value:       enum.ReplaceFieldReplaceMode.ToStrings(),
			Optional:    true,
			Description: "Mode that controls how the replacement is applied to matched content.",
		}.Schema(),
		"replace_regex": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Regular expression that selects which part of the field value to replace.",
		},
		"static_value": {
			Type:        schema.TypeList,
			Elem:        logControlRuleReplaceFieldStaticValueResource,
			Optional:    true,
			MaxItems:    1,
			Description: "Replace matched content with a static string. Exactly one of `mapped_value` or `static_value` must be set.",
		},
	},
}

var logControlRuleReplaceFieldMappedValueResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"default_value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Value to substitute when no matching key is found, when `use_default` is `true`.",
		},
		"pairs": {
			Type:        schema.TypeList,
			Elem:        logControlRuleReplaceFieldMappedValuePairResource,
			Optional:    true,
			Description: "List of key/value pairs that map matched content to replacement values.",
		},
		"use_default": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "If `true`, falls back to `default_value` when no key matches. If `false`, leaves the value unchanged on a miss.",
		},
	},
}

var logControlRuleReplaceFieldMappedValuePairResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"key": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Matched content to look up.",
		},
		"value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Replacement value to substitute when the key matches.",
		},
	},
}

var logControlRuleReplaceFieldStaticValueResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Replacement value. If empty, the matched content is removed.",
		},
	},
}

var logControlRuleParseFieldResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"source":      LogFieldPathSchema,
		"destination": LogFieldPathSchema,
		"parser":      LogParserSchema,
	},
}

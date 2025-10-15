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
		Type:     schema.TypeList,
		Elem:     logControlRuleResource,
		Optional: true,
	},
}

var logControlRuleResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"mode": Enum{
			Value:    enum.LogControlRuleMode.ToStrings(),
			Optional: true,
		}.Schema(),
		"filter": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": Enum{
			Value:    enum.LogControlRuleType.ToStrings(),
			Optional: true,
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
		"emit_metrics": {
			Type:     schema.TypeList,
			Elem:     logControlRuleEmitMetricsResource,
			Optional: true,
			MaxItems: 1,
		},
		"replace_field": {
			Type:     schema.TypeList,
			Elem:     logControlRuleReplaceFieldResource,
			Optional: true,
			MaxItems: 1,
		},
	},
}

var logControlRuleSampleResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"rate": {
			Type:     schema.TypeFloat,
			Optional: true,
		},
	},
}

var logControlRuleDropFieldResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"field_regex": {
			Type:     schema.TypeString,
			Optional: true,
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
			Type:     schema.TypeString,
			Optional: true,
		},
	},
}

var LogFieldPathSchema = &schema.Schema{
	Type:     schema.TypeList,
	Elem:     logFieldPathResource,
	Optional: true,
	MaxItems: 1,
}

var logControlRuleEmitMetricsResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"counter": {
			Type:     schema.TypeList,
			Elem:     logControlRuleEmitMetricsCounterResource,
			Optional: true,
			MaxItems: 1,
		},
		"drop_log": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"gauge": {
			Type:     schema.TypeList,
			Elem:     logControlRuleEmitMetricsGaugeResource,
			Optional: true,
			MaxItems: 1,
		},
		"histogram": {
			Type:     schema.TypeList,
			Elem:     logControlRuleEmitMetricsHistogramResource,
			Optional: true,
			MaxItems: 1,
		},
		"labels": {
			Type:     schema.TypeList,
			Elem:     logControlRuleEmitMetricsLabelResource,
			Optional: true,
		},
		"mode": Enum{
			Value:    enum.EmitMetricsMetricMode.ToStrings(),
			Optional: true,
		}.Schema(),
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	},
}

var logControlRuleEmitMetricsCounterResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"value": {
			Type:     schema.TypeList,
			Elem:     logFieldPathResource,
			Optional: true,
			MaxItems: 1,
		},
	},
}

var logControlRuleEmitMetricsGaugeResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"aggregation_type": Enum{
			Value:    enum.GaugeAggregationType.ToStrings(),
			Optional: true,
		}.Schema(),
		"value": {
			Type:     schema.TypeList,
			Elem:     logFieldPathResource,
			Optional: true,
			MaxItems: 1,
		},
	},
}

var logControlRuleEmitMetricsHistogramResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"value": {
			Type:     schema.TypeList,
			Elem:     logFieldPathResource,
			Optional: true,
			MaxItems: 1,
		},
	},
}

var logControlRuleEmitMetricsLabelResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"key": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"value": {
			Type:     schema.TypeList,
			Elem:     logFieldPathResource,
			Optional: true,
			MaxItems: 1,
		},
	},
}

var logControlRuleReplaceFieldResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"field": {
			Type:     schema.TypeList,
			Elem:     logFieldPathResource,
			Optional: true,
			MaxItems: 1,
		},
		"mapped_value": {
			Type:     schema.TypeList,
			Elem:     logControlRuleReplaceFieldMappedValueResource,
			Optional: true,
			MaxItems: 1,
		},
		"replace_all": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"replace_mode": Enum{
			Value:    enum.ReplaceFieldReplaceMode.ToStrings(),
			Optional: true,
		}.Schema(),
		"replace_regex": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"static_value": {
			Type:     schema.TypeList,
			Elem:     logControlRuleReplaceFieldStaticValueResource,
			Optional: true,
			MaxItems: 1,
		},
	},
}

var logControlRuleReplaceFieldMappedValueResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"default_value": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"pairs": {
			Type:     schema.TypeList,
			Elem:     logControlRuleReplaceFieldMappedValuePairResource,
			Optional: true,
		},
		"use_default": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	},
}

var logControlRuleReplaceFieldMappedValuePairResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"key": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"value": {
			Type:     schema.TypeString,
			Optional: true,
		},
	},
}

var logControlRuleReplaceFieldStaticValueResource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"value": {
			Type:     schema.TypeString,
			Optional: true,
		},
	},
}

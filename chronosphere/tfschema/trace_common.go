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
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/prettyenum"
)

var TraceSearchFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	MinItems: 1,
	MaxItems: 1,
	Required: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"trace": TraceFilterSchema,
			"span":  TraceSpanFilterListSchema,
		},
	},
}

var TraceFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	MinItems: 0,
	MaxItems: 1,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": TraceDurationFilterSchema,
			"error":    TraceBoolFilterSchema,
		},
	},
}

var TraceSpanFilterListSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          prettyenum.SpanFilterMatchTypeInclude,
				ValidateDiagFunc: validateSpanFilterMatchType,
				DiffSuppressFunc: diffSuppressSpanFilterMatchType,
			},
			"service":          TraceStringFilterSchema,
			"operation":        TraceStringFilterSchema,
			"parent_service":   TraceStringFilterSchema,
			"parent_operation": TraceStringFilterSchema,
			"duration":         TraceDurationFilterSchema,
			"error":            TraceBoolFilterSchema,
			// Note: this is the preferable form (singular) of this list field.
			"tag": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     traceTagFilterSchema,
			},
			// This field is still defined for error messaging to existing users, but is deprecated. Prefer tag.
			"tags": &schema.Schema{
				Type:       schema.TypeList,
				Optional:   true,
				Elem:       traceTagFilterSchema,
				Deprecated: "use tag instead",
			},
			"span_count": TraceSpanCountFilterSchema,
		},
	},
}

var TraceSpanCountFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MinItems: 0,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"min": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	},
}

var TraceBoolFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MinItems: 0,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"value": {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	},
}

var TraceDurationFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	MaxItems: 1,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"min_secs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  0.0,
			},
			"min_seconds": { // This field is still defined for error messaging to existing users, but is deprecated. Prefer min_secs.
				Type:       schema.TypeFloat,
				Optional:   true,
				Default:    0.0,
				Deprecated: "use min_secs instead",
			},
			"max_secs": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"max_seconds": { // This field is still defined for error messaging to existing users, but is deprecated. Prefer max_secs.
				Type:       schema.TypeFloat,
				Optional:   true,
				Default:    0.0,
				Deprecated: "use max_secs instead",
			},
		},
	},
}

var TraceStringFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MinItems: 0,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"match": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          prettyenum.StringFilterMatchTypeExact,
				ValidateDiagFunc: validateStringFilterMatchType,
				DiffSuppressFunc: diffSuppressStringFilterMatchType,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	},
}

var TraceNumericFilterSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MinItems: 0,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"comparison": Enum{
				Value:    enum.NumericFilterComparisonType.ToStrings(),
				Required: true,
			}.Schema(),
			"value": {
				Type:     schema.TypeFloat,
				Required: true,
			},
		},
	},
}

var traceTagFilterSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"key": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"value":         TraceStringFilterSchema,
		"numeric_value": TraceNumericFilterSchema,
	},
}

func validateStringFilterMatchType(i interface{}, _ cty.Path) diag.Diagnostics {
	rawType, ok := i.(string)
	if !ok {
		return diag.Errorf("expected match to be a string, got %T", i)
	}

	if err := prettyenum.ValidateStringFilterMatchType(rawType); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// diffSuppressStringFilterMatchType sanitizes and then diffs two string filter match type payloads.
func diffSuppressStringFilterMatchType(_, old, new string, _ *schema.ResourceData) bool {
	if old == new {
		return true
	}
	mtOld, err := prettyenum.NewStringFilterMatchType(old)
	if err != nil {
		return false
	}
	mtNew, err := prettyenum.NewStringFilterMatchType(new)
	if err != nil {
		return false
	}
	return mtOld.Model() == mtNew.Model()
}

func validateSpanFilterMatchType(i interface{}, _ cty.Path) diag.Diagnostics {
	rawType, ok := i.(string)
	if !ok {
		return diag.Errorf("expected match_type to be a string, got %T", i)
	}

	if err := prettyenum.ValidateSpanFilterMatchType(rawType); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// diffSuppressSpanFilterMatchType sanitizes and then diffs two span filter match type payloads.
func diffSuppressSpanFilterMatchType(_, old, new string, _ *schema.ResourceData) bool {
	if old == new {
		return true
	}
	mtOld, err := prettyenum.NewSpanFilterMatchType(old)
	if err != nil {
		return false
	}
	mtNew, err := prettyenum.NewSpanFilterMatchType(new)
	if err != nil {
		return false
	}
	return mtOld.Model() == mtNew.Model()
}

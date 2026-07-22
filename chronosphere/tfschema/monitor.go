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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/localid"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema/typeset"
)

var Monitor = map[string]*schema.Schema{
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the monitor. Generated from `name` if omitted. Immutable after creation.",
	},
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the monitor.",
	},
	"bucket_id": {
		Type:         schema.TypeString,
		Optional:     true,
		AtLeastOneOf: []string{"bucket_id", "collection_id"},
		Description:  "ID of the bucket the monitor belongs to. Exactly one of `bucket_id` or `collection_id` must be set.",
	},
	"collection_id": {
		Type:         schema.TypeString,
		Optional:     true,
		AtLeastOneOf: []string{"bucket_id", "collection_id"},
		Description:  "ID of the collection the monitor belongs to. Exactly one of `bucket_id` or `collection_id` must be set.",
	},
	"notification_policy_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "ID of the notification policy that routes signals from this monitor. If omitted, the parent collection's default policy applies. Must reference a named policy (anonymous policies are rejected).",
		ValidateDiagFunc: func(policyID any, _ cty.Path) diag.Diagnostics {
			if localid.IsLocalID(policyID.(string)) {
				return diag.Errorf("cannot directly reference unnamed notification policy, use a notification policy with name set")
			}
			return nil
		},
	},
	"labels": {
		Type:        schema.TypeMap,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Key/value labels attached to every signal emitted by the monitor. Used for routing and filtering.",
	},
	"annotations": {
		Type:        schema.TypeMap,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Free-form key/value pairs attached to every signal, intended for human consumption such as runbook URLs and descriptions.",
	},
	"query": {
		Type:        schema.TypeList,
		Required:    true,
		MinItems:    1,
		MaxItems:    1,
		Description: "Query that produces the time series evaluated by the monitor. Exactly one of `prometheus_expr`, `graphite_expr`, or `logging_expr` must be set.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"prometheus_expr": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "PromQL expression evaluated by the monitor.",
				},
				"graphite_expr": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Graphite expression evaluated by the monitor.",
				},
				"logging_expr": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Log query expression evaluated by the monitor.",
				},
			},
		},
	},
	"signal_grouping": SignalGrouping,
	"series_conditions": {
		Type:        schema.TypeList,
		Required:    true,
		MinItems:    1,
		MaxItems:    1,
		Description: "Conditions that determine when a series fires a signal.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"condition": MonitorSeriesConditionSchema,
				"override": {
					Type:        schema.TypeList,
					Optional:    true,
					Description: "Per-series overrides that apply different conditions to series matching a set of label matchers.",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"label_matcher": MatcherListSchema,
							"condition":     MonitorSeriesConditionSchema,
						},
					},
				},
			},
		},
	},
	"interval": Duration{
		Optional:    true,
		Description: "Evaluation interval (e.g. `30s`, `1m`). Defaults to the system default if unset.",
	}.Schema(),
	"notification_template": {
		Type:        schema.TypeList,
		MaxItems:    1,
		Optional:    true,
		Description: "Templated title/description rendered into outbound notifications. Supports Go templating with access to signal labels and annotations.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"title": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Title template for the notification.",
				},
				"description": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Body/description template for the notification.",
				},
			},
		},
	},
	"schedule": {
		Type:        schema.TypeList,
		MaxItems:    1,
		Optional:    true,
		Description: "Optional schedule restricting when the monitor evaluates and fires.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"timezone": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "IANA timezone name (e.g. `America/New_York`) used to interpret `range` values.",
				},
				"range": typeset.Set{
					Description: "Time-of-day ranges during which the monitor is active. The monitor is inactive outside these ranges.",
					ElemFields: map[string]typeset.ElemField{
						"day": CaseInsensitiveString{
							Required:    true,
							Description: "Day of week, e.g. `monday`. Case-insensitive.",
						},
						"start": typeset.NotNormalized(&schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Start time of day, 24-hour `HH:MM` format.",
						}),
						"end": typeset.NotNormalized(&schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "End time of day, 24-hour `HH:MM` format.",
						}),
					},
				}.Schema(),
			},
		},
	},
}

var MonitorSeriesConditionSchema = typeset.Set{
	Required:    true,
	MinItems:    1,
	Description: "One or more severity/threshold conditions. Multiple conditions enable multi-severity monitors (e.g. warn at one threshold, page at a higher one).",
	ElemFields: map[string]typeset.ElemField{
		// Note, severity is case-sensitive.
		"severity": typeset.NotNormalized(&schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Severity assigned when this condition matches (e.g. `warn`, `critical`). Case-sensitive.",
		}),
		"op": Enum{
			Value:       enum.ConditionOp.ToStrings(),
			Required:    true,
			Description: "Comparison operator between the query value and `value` (e.g. `gt`, `lt`, `eq`).",
		},
		"value": typeset.NotNormalized(&schema.Schema{
			Type:        schema.TypeFloat,
			Optional:    true,
			Default:     0,
			Description: "Threshold compared against the query result using `op`. Defaults to 0.",
		}),
		"sustain": Duration{
			Optional:    true,
			Description: "Duration the condition must hold continuously before a signal fires.",
		},
		"resolve_sustain": Duration{
			Optional:    true,
			Description: "Duration the condition must remain false continuously before an active signal resolves.",
		},
		"resolve_value": typeset.NotNormalized(&schema.Schema{
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Optional separate threshold used for resolution, enabling hysteresis (e.g. fire at >90, resolve at <80).",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"value": {
						Type:        schema.TypeFloat,
						Required:    true,
						Description: "Resolution threshold value.",
					},
					"enabled": {
						Type:        schema.TypeBool,
						Required:    true,
						Description: "Whether the resolve-value threshold is active.",
					},
				},
			},
		}),
		"resolve_sustain_for_no_data": resolveSustainForNoData{},
	},
}.Schema()

// resolveSustainForNoData is the ElemField for the resolve_sustain_for_no_data
// block. It implements Normalize (rather than typeset.NotNormalized) because
// the nested sustain duration must hash stably inside the condition set.
type resolveSustainForNoData struct{}

func (resolveSustainForNoData) Schema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "Controls how a firing condition resolves once its series stops returning data, independently of `resolve_sustain` (which governs recovery while data is still present).",
		DiffSuppressFunc: func(_, old, new string, _ *schema.ResourceData) bool {
			return Duration{}.Normalize(old) == Duration{}.Normalize(new)
		},
		DiffSuppressOnRefresh: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"enabled": {
					Type:        schema.TypeBool,
					Required:    true,
					Description: "When false (default), missing data is treated the same as passing data and `resolve_sustain` governs recovery. When true, a firing signal keeps firing while its data is missing and auto-resolves once data has been missing for `sustain`.",
				},
				"sustain": Duration{
					Optional:    true,
					Description: "How long a firing signal keeps firing after its data goes missing before auto-resolving. Only applies when `enabled` is true. Omitted or `0` resolves immediately on missing data. Must be less than 3 days.",
				}.Schema(),
			},
		},
	}
}

func (resolveSustainForNoData) Normalize(v any) any {
	list, ok := v.([]any)
	if !ok {
		return v
	}
	out := make([]any, 0, len(list))
	for _, e := range list {
		elem, ok := e.(map[string]any)
		if !ok {
			out = append(out, e)
			continue
		}
		normalized := make(map[string]any, len(elem))
		for k, ev := range elem {
			if s, isStr := ev.(string); isStr && k == "sustain" {
				ev = Duration{}.Normalize(s)
			}
			normalized[k] = ev
		}
		out = append(out, normalized)
	}
	return out
}

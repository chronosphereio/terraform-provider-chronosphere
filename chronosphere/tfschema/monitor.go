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
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"bucket_id": {
		Type:         schema.TypeString,
		Optional:     true,
		AtLeastOneOf: []string{"bucket_id", "collection_id"},
	},
	"collection_id": {
		Type:         schema.TypeString,
		Optional:     true,
		AtLeastOneOf: []string{"bucket_id", "collection_id"},
	},
	"notification_policy_id": {
		Type:     schema.TypeString,
		Optional: true,
		ValidateDiagFunc: func(policyID any, _ cty.Path) diag.Diagnostics {
			if localid.IsLocalID(policyID.(string)) {
				return diag.Errorf("cannot directly reference unnamed notification policy, use a notification policy with name set")
			}
			return nil
		},
	},
	"labels": {
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
	"annotations": {
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
	"query": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"prometheus_expr": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"graphite_expr": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	},
	"signal_grouping": {
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"label_names": {
					Type: schema.TypeList,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional: true,
				},
				"signal_per_series": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
	},
	"series_conditions": {
		Type:     schema.TypeList,
		Required: true,
		MinItems: 1,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"condition": MonitorSeriesConditionSchema,
				"override": {
					Type:     schema.TypeList,
					Optional: true,
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
		Optional: true,
	}.Schema(),
	"schedule": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"timezone": {
					Type:     schema.TypeString,
					Required: true,
				},
				"range": typeset.Set{
					ElemFields: map[string]typeset.ElemField{
						"day": CaseInsensitiveString{
							Required: true,
						},
						"start": typeset.NotNormalized(&schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						}),
						"end": typeset.NotNormalized(&schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						}),
					},
				}.Schema(),
			},
		},
	},
}

var MonitorSeriesConditionSchema = typeset.Set{
	Required: true,
	MinItems: 1,
	ElemFields: map[string]typeset.ElemField{
		// Note, severity is case-sensitive.
		"severity": typeset.NotNormalized(&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		}),
		"op": Enum{
			Value:    enum.ConditionOp.ToStrings(),
			Required: true,
		},
		"value": typeset.NotNormalized(&schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  0,
		}),
		"sustain": Duration{
			Optional: true,
		},
		"resolve_sustain": Duration{
			Optional: true,
		},
	},
}.Schema()

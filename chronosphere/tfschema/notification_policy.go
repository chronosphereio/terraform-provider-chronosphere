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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfschema/typeset"
)

// When a notification policy is independent, we always set the
// notification_policy_data to this static sentinel value such that the bucket
// resource can detect when it's referencing an independent policy.
const IndependentNotificationPolicyData = "__independent"

var NotificationPolicy = map[string]*schema.Schema{
	// NB: slug can only be set if name is set, but cannot be set if name is unset, e.g. inline.
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the notification policy. Can only be set when `name` is set. Generated from `name` if omitted. Immutable after creation.",
	},
	// NB: there is custom name ForceNew behavior in the DiffSuppressFunc.
	"name": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Display name of the notification policy. If unset, the policy is treated as anonymous/inline and cannot be referenced by ID. Changing between set/unset forces resource replacement.",
	},
	"team_id": {
		Type:         schema.TypeString,
		Optional:     true,
		RequiredWith: []string{"name"},
		Description:  "ID of the team that owns this notification policy. Required when `name` is set (anonymous policies cannot be owned).",
	},
	"route": NotificationRouteSchema,
	"override": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Ordered overrides that route alerts matching specific label matchers to different destinations. The first matching override is applied; non-matching alerts fall through to the default `route`.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"alert_label_matcher": MatcherListSchema,
				"route":               NotificationRouteSchema,
			},
		},
	},
	"notification_policy_data": {
		Type:             schema.TypeString,
		Optional:         true,
		Computed:         true,
		DiffSuppressFunc: JSONNotificationPolicyDiffSuppress,
		ValidateFunc:     ValidateNotificationPolicyData,
		Description:      "Computed/optional JSON serialization of the policy. Primarily used to attach inline policy data to other resources (e.g. buckets).",
	},
	// This field is for internal use only. We use it to force new resources when the name
	// of a notification policy changes (from inline to independent or vice versa)
	"is_independent": {
		Type:        schema.TypeBool,
		Computed:    true,
		ForceNew:    true,
		Description: "Read-only internal marker tracking whether the policy is independent (named) or inline. Used to force replacement when transitioning between the two.",
	},
}

var NotificationRouteSchema = typeset.Set{
	Description: "Per-severity routing rules. Each entry maps a severity (e.g. `warn`, `critical`) to a set of notifiers, destinations, grouping, and repeat behavior.",
	ElemFields: map[string]typeset.ElemField{
		// Note, severity is case-sensitive.
		"severity": typeset.NotNormalized(&schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Severity this route applies to (e.g. `warn`, `critical`). Case-sensitive.",
		}),
		"notifiers": typeset.NotNormalized(&schema.Schema{
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Slugs of notifier resources that receive alerts at this severity. Cannot be combined with `destination`.",
		}),
		"repeat_interval": Duration{
			Optional:    true,
			Description: "How often to resend unresolved alerts at this severity (e.g. `4h`).",
		},
		"group_by": typeset.NotNormalized(&schema.Schema{
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "Optional grouping configuration controlling how alerts are batched before delivery.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"label_names": {
						Type:        schema.TypeList,
						Optional:    true,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Description: "Label names to group alerts by. Alerts with identical values for these labels are bundled into a single notification.",
					},
				},
			},
		}),
		"destination": typeset.NotNormalized(&schema.Schema{
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Inline notification destinations defined directly on the route. Each block sets at most one of `slack`, `pagerduty`, `webhook`, `ops_genie`, `victor_ops`, or `email`. Cannot be combined with `notifiers`.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"disable_resolves": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "If true, do not send notifications when alerts resolve. Defaults to false.",
					},
					"slack": {
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Description: "Slack delivery configuration for this destination.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"external_connection_slug": {
									Type:        schema.TypeString,
									Required:    true,
									Description: "Slug of the Slack external connection holding the integration credentials.",
								},
								"channels": {
									Type:        schema.TypeList,
									Optional:    true,
									Elem:        &schema.Schema{Type: schema.TypeString},
									Description: "Slack channels to send notifications to.",
								},
							},
						},
					},
					"pagerduty": {
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Description: "PagerDuty delivery configuration for this destination.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"external_connection_slug": {
									Type:        schema.TypeString,
									Required:    true,
									Description: "Slug of the PagerDuty external connection holding the integration credentials.",
								},
							},
						},
					},
					"webhook": {
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Description: "Generic webhook delivery configuration for this destination.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"external_connection_slug": {
									Type:        schema.TypeString,
									Required:    true,
									Description: "Slug of the webhook external connection holding the endpoint URL and auth settings.",
								},
								"query_parameter": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "Additional query parameters appended to the webhook URL when delivering this notification.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"key": {
												Type:        schema.TypeString,
												Required:    true,
												Description: "Query parameter name.",
											},
											"value": {
												Type:        schema.TypeString,
												Required:    true,
												Description: "Query parameter value.",
											},
										},
									},
								},
							},
						},
					},
					"ops_genie": {
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Description: "OpsGenie delivery configuration for this destination.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"external_connection_slug": {
									Type:        schema.TypeString,
									Required:    true,
									Description: "Slug of the OpsGenie external connection holding the integration credentials.",
								},
							},
						},
					},
					"victor_ops": {
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Description: "VictorOps (Splunk On-Call) delivery configuration for this destination.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"external_connection_slug": {
									Type:        schema.TypeString,
									Required:    true,
									Description: "Slug of the VictorOps external connection holding the integration credentials.",
								},
								"routing_keys": {
									Type:        schema.TypeList,
									Required:    true,
									Elem:        &schema.Schema{Type: schema.TypeString},
									Description: "VictorOps routing keys identifying the destination escalation policies.",
								},
							},
						},
					},
					"email": {
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Description: "Email delivery configuration for this destination.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"addresses": {
									Type:        schema.TypeList,
									Required:    true,
									Elem:        &schema.Schema{Type: schema.TypeString},
									Description: "Email addresses to deliver notifications to.",
								},
							},
						},
					},
				},
			},
		}),
	},
}.Schema()

func deprecated(s *schema.Schema, msg string) *schema.Schema {
	// Shallow copy to preserve the Elem address so we can still share intschema
	// structs.
	c := *s
	c.Deprecated = msg
	return &c
}

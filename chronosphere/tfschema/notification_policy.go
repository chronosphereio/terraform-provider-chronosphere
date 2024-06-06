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
	// NB: slug may be set if bucket_id or team_id is set, but cannot be set if bucket_id and team_id are unset.
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"name": {
		Type:     schema.TypeString,
		Optional: true,
	},
	// NB: there is custom team_id ForceNew behavior in the DiffSuppressFunc.
	"team_id": {
		Type:         schema.TypeString,
		Optional:     true,
		RequiredWith: []string{"name"},
	},
	"route": NotificationRouteSchema,
	"override": {
		Type:     schema.TypeList,
		Optional: true,
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
	},
	// This field is for internal use only. We use it to force new resources when the ownership
	// of a notification policy changes (from unowned to independent or vice versa)
	"is_independent": {
		Type:     schema.TypeBool,
		Computed: true,
		ForceNew: true,
	},
}

var NotificationRouteSchema = typeset.Set{
	ElemFields: map[string]typeset.ElemField{
		// Note, severity is case-sensitive.
		"severity": typeset.NotNormalized(&schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		}),
		"notifiers": typeset.NotNormalized(&schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		}),
		"repeat_interval": Duration{
			Optional: true,
		},
		"group_by": typeset.NotNormalized(&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
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

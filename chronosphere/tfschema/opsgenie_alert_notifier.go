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

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var OpsgenieAlertNotifier = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the notifier.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the notifier. Generated from `name` if omitted. Immutable after creation.",
	},
	"send_resolved": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Whether to send a follow-up notification when an alert is resolved. Defaults to true.",
	},
	"api_key": {
		Type:        schema.TypeString,
		Required:    true,
		Sensitive:   true,
		Description: "Opsgenie API key used to authenticate requests. Treat as a secret.",
	},
	"api_url": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Opsgenie API URL to send requests to (e.g. `https://api.opsgenie.com/`).",
	},
	"message": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Alert text shown in Opsgenie. Supports Go templating.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Detailed description of the alert. Supports Go templating.",
	},
	"source": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Backlink to the sender of the notification. Supports Go templating.",
	},
	"details": {
		Type:        schema.TypeMap,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Arbitrary key/value pairs attached to the alert as additional context. Values support Go templating.",
	},
	"responder": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Responders that Opsgenie will notify for the alert. See https://docs.opsgenie.com/docs/alert-api for accepted shapes.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Opsgenie identifier of the responder. Use instead of `name` or `username`.",
				},
				"name": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Name of the responder team, schedule, or escalation policy.",
				},
				"username": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Username of a user responder.",
				},

				"type": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "Responder type. One of `team`, `user`, `escalation`, or `schedule`.",
				},
			},
		},
	},
	"tags": {
		Type: schema.TypeSet,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		Description: "Tags attached to the Opsgenie alert.",
	},
	"note": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Additional note appended to the alert. Supports Go templating.",
	},
	"priority": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Priority level of the alert. One of `P1`, `P2`, `P3`, `P4`, or `P5`.",
	},

	"basic_auth_username": {
		Type:          schema.TypeString,
		Optional:      true,
		RequiredWith:  []string{"basic_auth_password"},
		ConflictsWith: []string{"bearer_token"},
		Description:   "Username for HTTP basic auth when calling the Opsgenie API. Mutually exclusive with `bearer_token`.",
	},
	"basic_auth_password": {
		Type:         schema.TypeString,
		Optional:     true,
		RequiredWith: []string{"basic_auth_password"},
		Sensitive:    true,
		Description:  "Password for HTTP basic auth when calling the Opsgenie API. Treat as a secret.",
	},
	"bearer_token": {
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"basic_auth_username"},
		Description:   "Bearer token sent in the `Authorization` header when calling the Opsgenie API. Treat as a secret. Mutually exclusive with basic auth.",
	},
	"proxy_url": {
		Type:        schema.TypeString,
		Optional:    true,
		Deprecated:  "custom proxy URLs are not supported",
		Description: "Deprecated and ignored. Custom proxy URLs are not supported.",
	},
	"tls_insecure_skip_verify": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "If true, skip TLS certificate verification when calling the Opsgenie API. Disable only in trusted environments.",
	},
}

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

var PagerdutyAlertNotifier = map[string]*schema.Schema{
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
	"class": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Class of the event reported to PagerDuty (e.g. `cpu`, `database`). Supports Go templating.",
	},
	"client": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Name of the monitoring client identified in the notification.",
	},
	"client_url": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Backlink to the sender of the notification, shown in PagerDuty.",
	},
	"component": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Part or component of the affected system that is broken. Supports Go templating.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Summary of the incident. Supports Go templating.",
	},
	"details": {
		Type:        schema.TypeMap,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Arbitrary key/value pairs attached to the incident as additional context. Values support Go templating.",
	},
	"group": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Logical grouping of services the incident belongs to. Supports Go templating.",
	},
	"basic_auth_username": {
		Type:          schema.TypeString,
		Optional:      true,
		RequiredWith:  []string{"basic_auth_password"},
		ConflictsWith: []string{"bearer_token"},
		Description:   "Username for HTTP basic auth when calling the PagerDuty API. Mutually exclusive with `bearer_token`.",
	},
	"basic_auth_password": {
		Type:         schema.TypeString,
		Optional:     true,
		RequiredWith: []string{"basic_auth_password"},
		Sensitive:    true,
		Description:  "Password for HTTP basic auth when calling the PagerDuty API. Treat as a secret.",
	},
	"bearer_token": {
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"basic_auth_username"},
		Description:   "Bearer token sent in the `Authorization` header when calling the PagerDuty API. Treat as a secret. Mutually exclusive with basic auth.",
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
		Description: "If true, skip TLS certificate verification when calling the PagerDuty API. Disable only in trusted environments.",
	},
	"image": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Images attached to the PagerDuty incident. See https://developer.pagerduty.com/docs/events-api-v2/trigger-events/.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"alt": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Alternate text shown when the image cannot be rendered.",
				},
				"href": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Optional URL the image links to when clicked.",
				},
				"src": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "URL of the image to attach.",
				},
			},
		},
	},
	"link": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Hyperlinks attached to the PagerDuty incident. See https://developer.pagerduty.com/docs/events-api-v2/trigger-events/.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"href": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "URL the link points to.",
				},
				"text": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Display text for the link.",
				},
			},
		},
	},
	"routing_key": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "PagerDuty integration key when using the `Events API v2` integration type. Treat as a secret. Mutually exclusive with `service_key`.",
	},
	"service_key": {
		Type:        schema.TypeString,
		Optional:    true,
		Sensitive:   true,
		Description: "PagerDuty integration key when using the `Prometheus` integration type. Treat as a secret. Mutually exclusive with `routing_key`.",
	},
	"severity": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Severity of the incident. One of `critical`, `error`, `warning`, or `info`.",
	},
	"url": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "PagerDuty API URL to send events to (e.g. `https://events.pagerduty.com/v2/enqueue`).",
	},
}

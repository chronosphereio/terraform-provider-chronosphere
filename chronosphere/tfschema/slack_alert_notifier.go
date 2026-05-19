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

var SlackAlertNotifier = map[string]*schema.Schema{
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
	"action": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Interactive buttons appended to the Slack message. See https://api.slack.com/reference/messaging/attachments#action_fields.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Identifier sent back to Slack when the button is clicked.",
				},
				"style": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Visual style of the button: `default`, `primary`, or `danger`.",
				},
				"text": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Label shown on the button.",
				},
				"type": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Action type. Typically `button`.",
				},
				"url": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Link the button navigates to when clicked.",
				},
				"value": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Opaque value sent back to Slack alongside `name` when the button is clicked.",
				},
				"action_confirm_text": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Body text of the confirmation dialog shown before the action runs.",
				},
				"action_confirm_tile": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Title of the confirmation dialog shown before the action runs.",
				},
				"action_confirm_ok_text": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Label for the confirm button in the confirmation dialog.",
				},
				"action_confirm_dismiss_text": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Label for the cancel button in the confirmation dialog.",
				},
			},
		},
	},
	"api_url": {
		Type:        schema.TypeString,
		Required:    true,
		Sensitive:   true,
		Description: "Slack incoming webhook URL that receives the notifications. Treat this as a secret.",
	},
	"send_resolved": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true,
		Description: "Whether to send a follow-up notification when an alert is resolved. Defaults to true.",
	},
	"callback_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Slack callback ID used to identify the source of interactive actions.",
	},
	"channel": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Slack channel to post notifications to (e.g. `#alerts`).",
	},
	"color": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Color of the attachment border. Hex code or one of `good`, `warning`, `danger`. Supports Go templating.",
	},
	"fallback": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Plain-text fallback shown in notifications and clients that don't render attachments. Supports Go templating.",
	},
	"fields": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Structured field/value pairs rendered as a table in the attachment.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"title": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Bold heading shown above the value.",
				},
				"value": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Value text. Supports Go templating.",
				},
				"short": {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: "If true, the field is short enough to be shown side-by-side with the next field.",
				},
			},
		},
	},
	"footer": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Footer text shown at the bottom of the attachment. Supports Go templating.",
	},
	"basic_auth_username": {
		Type:          schema.TypeString,
		Optional:      true,
		RequiredWith:  []string{"basic_auth_password"},
		ConflictsWith: []string{"bearer_token"},
		Description:   "Username for HTTP basic auth when calling the webhook. Mutually exclusive with `bearer_token`.",
	},
	"basic_auth_password": {
		Type:         schema.TypeString,
		Optional:     true,
		RequiredWith: []string{"basic_auth_password"},
		Sensitive:    true,
		Description:  "Password for HTTP basic auth when calling the webhook.",
	},
	"bearer_token": {
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"basic_auth_username"},
		Description:   "Bearer token sent in the `Authorization` header when calling the webhook. Mutually exclusive with basic auth.",
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
		Description: "If true, skip TLS certificate verification when calling the webhook. Disable only in trusted environments.",
	},
	"icon_emoji": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Slack emoji to use as the bot avatar (e.g. `:fire:`). Mutually exclusive with `icon_url` at Slack.",
	},
	"icon_url": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "URL of an image to use as the bot avatar.",
	},
	"image_url": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "URL of an image attached to the message.",
	},
	"link_names": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "If true, find and link channel names and usernames in the message text.",
	},
	"mrkdwn_in": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Attachment fields in which Slack should parse `mrkdwn` formatting. Common values: `pretext`, `text`, `fields`.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"pretext": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Text shown above the attachment. Supports Go templating.",
	},
	"short_fields": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "If true, render all `fields` with `short: true` regardless of per-field setting.",
	},
	"text": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Body text of the attachment. Supports Go templating.",
	},
	"thumb_url": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "URL of a small thumbnail image shown to the right of the attachment.",
	},
	"title": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Attachment title, rendered as a header. Supports Go templating.",
	},
	"title_link": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "URL the title links to when clicked.",
	},
	"username": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Display name of the bot posting the message.",
	},
}

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

var SlackExternalConnection = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the external connection.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the connection. Generated from `name` if omitted. Immutable after creation.",
	},
	"api_url": {
		Type:     schema.TypeString,
		Optional: true,
		Description: "Base URL of the Slack Web API that messages are posted to, with `/chat.postMessage` appended " +
			"unless the URL already ends with it. Defaults to `https://slack.com/api`; set it only for another " +
			"Slack host, such as `https://slack-gov.com/api`, or a proxy. This is not an incoming webhook URL: " +
			"to post through a webhook, use `chronosphere_slack_alert_notifier`.",
	},
	"token": {
		Type:      schema.TypeString,
		Required:  true,
		Sensitive: true,
		Description: "Slack bot token, such as `xoxb-...`, used to post messages with `chat.postMessage`. The bot " +
			"needs the `chat:write` scope and must be a member of each channel it posts to, except a public " +
			"channel when it also has `chat:write.public`. The API never returns the token, so a token changed " +
			"outside Terraform is not detected, and an imported connection plans one update to set it.",
	},
}

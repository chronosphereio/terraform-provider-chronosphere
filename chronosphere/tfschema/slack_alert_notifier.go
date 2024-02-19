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
		Type:     schema.TypeString,
		Required: true,
	},
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"action": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"style": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"text": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"type": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"url": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"value": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"action_confirm_text": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"action_confirm_tile": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"action_confirm_ok_text": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"action_confirm_dismiss_text": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	},
	"api_url": {
		Type:      schema.TypeString,
		Required:  true,
		Sensitive: true,
	},
	"send_resolved": {
		Type:     schema.TypeBool,
		Optional: true,
		Default:  true,
	},
	"callback_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"channel": {
		Type:     schema.TypeString,
		Required: true,
	},
	"color": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"fallback": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"fields": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"title": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"value": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"short": {
					Type:     schema.TypeBool,
					Optional: true,
				},
			},
		},
	},
	"footer": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"basic_auth_username": {
		Type:          schema.TypeString,
		Optional:      true,
		RequiredWith:  []string{"basic_auth_password"},
		ConflictsWith: []string{"bearer_token"},
	},
	"basic_auth_password": {
		Type:         schema.TypeString,
		Optional:     true,
		RequiredWith: []string{"basic_auth_password"},
		Sensitive:    true,
	},
	"bearer_token": {
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"basic_auth_username"},
	},
	"proxy_url": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"tls_insecure_skip_verify": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"icon_emoji": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"icon_url": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"image_url": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"link_names": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"mrkdwn_in": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"pretext": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"short_fields": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"text": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"thumb_url": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"title": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"title_link": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"username": {
		Type:     schema.TypeString,
		Optional: true,
	},
}

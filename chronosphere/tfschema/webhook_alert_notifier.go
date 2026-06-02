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

var WebhookAlertNotifier = map[string]*schema.Schema{
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
	"url": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Webhook URL that receives the alert payload via HTTP POST.",
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
		Description:  "Password for HTTP basic auth when calling the webhook. Treat as a secret.",
	},
	"bearer_token": {
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"basic_auth_username"},
		Description:   "Bearer token sent in the `Authorization` header when calling the webhook. Treat as a secret. Mutually exclusive with basic auth.",
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
}

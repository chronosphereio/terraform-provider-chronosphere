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

var WebhookExternalConnection = map[string]*schema.Schema{
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
	"url": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Destination URL that receives webhook POST requests for this connection.",
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
		RequiredWith: []string{"basic_auth_username"},
		Sensitive:    true,
		Description:  "Password for HTTP basic auth when calling the webhook. Treat as a secret.",
	},
	"bearer_token": {
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"basic_auth_username"},
		Description:   "Bearer token sent in the `Authorization` header when calling the webhook. Mutually exclusive with basic auth. Treat as a secret.",
	},
	"tls_insecure_skip_verify": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "If true, skip TLS certificate verification when calling the webhook. Disable only in trusted environments.",
	},
}

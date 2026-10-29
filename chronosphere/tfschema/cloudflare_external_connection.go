// Copyright 2026 Chronosphere Inc.
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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var CloudflareExternalConnection = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the external connection.",
	},
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
		Description: "Stable identifier for the connection. Generated from `name` if omitted. Immutable after creation, " +
			"so changing it replaces the connection. If a `chronosphere_cloud_integration` uses the connection, the " +
			"replacement needs `create_before_destroy` and a new `name` in the same apply, because a connection in " +
			"use cannot be deleted and names are unique.",
	},
	"api_token_wo":         cloudflareAPITokenSchema(),
	"api_token_wo_version": SecretVersionSchema("api_token"),
}

func cloudflareAPITokenSchema() *schema.Schema {
	s := WriteOnlySecret{
		Name:        "api_token",
		Description: "Cloudflare API token used to read analytics.",
		Required:    true,
	}.Schema()
	s.ValidateFunc = validation.StringIsNotEmpty
	return s
}

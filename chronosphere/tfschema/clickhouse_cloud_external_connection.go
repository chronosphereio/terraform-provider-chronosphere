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

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

var ClickhouseCloudExternalConnection = map[string]*schema.Schema{
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
	"key_id": {
		Type:     schema.TypeString,
		Required: true,
		Description: "ID of a ClickHouse Cloud API key, sent as the HTTP Basic username. At most 256 bytes, with no " +
			"colon, whitespace or control characters.",
	},
	"key_secret": {
		Type:      schema.TypeString,
		Required:  true,
		Sensitive: true,
		Description: "Secret of the same ClickHouse Cloud API key, sent as the HTTP Basic password, with no whitespace " +
			"or control characters. The API never returns the secret, so a secret changed outside Terraform is not detected, and " +
			"an imported connection plans one update to set it.",
	},
}

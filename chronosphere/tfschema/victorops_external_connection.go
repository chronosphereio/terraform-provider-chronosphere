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

var VictoropsExternalConnection = map[string]*schema.Schema{
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
	"api_key": {
		Type:        schema.TypeString,
		Optional:    true,
		Sensitive:   true,
		Description: "VictorOps (Splunk On-Call) REST integration API key used to authenticate alert delivery. Treat as a secret.",
	},
	"api_url": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "VictorOps REST endpoint URL that receives alert payloads. Override to target a custom endpoint.",
	},
}

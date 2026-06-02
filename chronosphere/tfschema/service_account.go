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

import (
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/enum"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var serviceAccountPermOneOfFields = []string{"restriction", "unrestricted"}

var ServiceAccount = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "Display name of the service account. Immutable after creation.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the service account. Generated from `name` if omitted. Immutable after creation.",
	},
	"email": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Read-only: synthetic email address assigned to the service account by the server.",
	},
	"token": {
		Type:        schema.TypeString,
		Computed:    true,
		Sensitive:   true,
		Description: "Read-only: API token generated for the service account. Returned only at creation time; store it securely. If lost, the service account must be recreated.",
	},
	"unrestricted": {
		Type:         schema.TypeBool,
		Optional:     true,
		ForceNew:     true,
		ExactlyOneOf: serviceAccountPermOneOfFields,
		Description:  "If true, grants the service account access to all Chronosphere APIs within the access controls defined by team membership. Exactly one of `unrestricted` or `restriction` must be set.",
	},
	"restriction": {
		Type:        schema.TypeList,
		Optional:    true,
		ForceNew:    true,
		MaxItems:    1,
		Description: "Restricts the service account to a specific permission and optional metric label scope. Exactly one of `unrestricted` or `restriction` must be set.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"permission": Enum{
					Value:       enum.Permission.ToStrings(),
					Required:    true,
					Description: "Permission level granted by this restriction (e.g. metric read/write).",
				}.Schema(),
				"labels": {
					Type:        schema.TypeMap,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Optional:    true,
					Description: "Optional label matchers further scoping the restriction to metrics whose labels match these key/value pairs.",
				},
			},
		},
		ExactlyOneOf: serviceAccountPermOneOfFields,
	},
}

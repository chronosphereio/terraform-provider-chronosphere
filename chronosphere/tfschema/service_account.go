// Copyright 2023 Chronosphere Inc.
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
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"email": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"token": {
		Type:      schema.TypeString,
		Computed:  true,
		Sensitive: true,
	},
	"unrestricted": {
		Type:         schema.TypeBool,
		Optional:     true,
		ForceNew:     true,
		ExactlyOneOf: serviceAccountPermOneOfFields,
	},
	"restriction": {
		Type:     schema.TypeList,
		Optional: true,
		ForceNew: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"permission": Enum{
					Value:    enum.Permission.ToStrings(),
					Required: true,
				}.Schema(),
				"labels": {
					Type:     schema.TypeMap,
					Elem:     &schema.Schema{Type: schema.TypeString},
					Optional: true,
				},
			},
		},
		ExactlyOneOf: serviceAccountPermOneOfFields,
	},
}

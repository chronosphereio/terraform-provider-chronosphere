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

var CommandCenterGroup = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the command center group.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the command center group. Generated from `name` if omitted. Immutable after creation.",
	},
	"group_slo_reference": {
		Type:          schema.TypeList,
		Optional:      true,
		MaxItems:      1,
		ConflictsWith: []string{"primary_slo_reference"},
		Deprecated:    "use `primary_slo_reference` instead",
		Description:   "Deprecated: use `primary_slo_reference` instead. Reference to the primary SLO tracked by this group.",
		Elem:          CommandCenterSLOReferenceElemSchema,
	},
	"primary_slo_reference": {
		Type:          schema.TypeList,
		Optional:      true,
		MaxItems:      1,
		ConflictsWith: []string{"group_slo_reference"},
		Description:   "Reference to the primary SLO tracked by this group.",
		Elem:          CommandCenterSLOReferenceElemSchema,
	},
	"related_slo_references": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Related SLOs tracked by this group, secondary to the primary one.",
		Elem:        CommandCenterSLOReferenceElemSchema,
	},
}

var CommandCenterSLOReferenceElemSchema = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"slug": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Slug of the referenced SLO.",
		},
	},
}

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

var Bucket = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the bucket. Can be changed after creation.",
	},
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the bucket. Generated from `name` if omitted. Immutable after creation.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Free-form description of the bucket.",
	},
	"labels": {
		Type:        schema.TypeMap,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Key/value labels attached to the bucket for organization and filtering.",
	},
	// notification_policy_slug is an internal field used to track the slug of inline policies
	// set via notification_policy_data.
	// Users who want to specify a default policy should use notification_policy_id.
	"notification_policy_slug": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Internal field tracking the slug of an inline notification policy defined via `notification_policy_data`. Use `notification_policy_id` to reference a named policy.",
	},
	"notification_policy_id": {
		Type:          schema.TypeString,
		ConflictsWith: []string{"notification_policy_data"},
		Optional:      true,
		Description:   "ID of the default notification policy applied to monitors in this bucket that do not explicitly reference one. Conflicts with `notification_policy_data`.",
	},
	"team_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "ID of the team that owns this bucket.",
	},
	"notification_policy_data": {
		Type:             schema.TypeString,
		Optional:         true,
		DiffSuppressFunc: JSONNotificationPolicyDiffSuppress,
		ValidateFunc:     ValidateNotificationPolicyData,
		Description:      "Inline notification policy serialized as JSON. Conflicts with `notification_policy_id`. For reusability, reference a named policy instead.",
	},
}

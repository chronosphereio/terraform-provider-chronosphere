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
		Type:     schema.TypeString,
		Required: true,
	},
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"labels": {
		Type:     schema.TypeMap,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
	// notification_policy_slug is an internal field used to track the slug of inline policies
	// set via notification_policy_data.
	// Users who want to specify a default policy should use notification_policy_id.
	"notification_policy_slug": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"notification_policy_id": {
		Type:          schema.TypeString,
		ConflictsWith: []string{"notification_policy_data"},
		Optional:      true,
	},
	"team_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"notification_policy_data": {
		Type:             schema.TypeString,
		Optional:         true,
		DiffSuppressFunc: JSONNotificationPolicyDiffSuppress,
		ValidateFunc:     ValidateNotificationPolicyData,
	},
}

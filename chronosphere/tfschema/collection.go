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

var Collection = map[string]*schema.Schema{
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the collection. Generated from `name` if omitted. Immutable after creation.",
	},
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the collection. Can be changed after creation.",
	},
	"team_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "ID of the team that owns this collection.",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Free-form description of the collection.",
	},
	"notification_policy_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "ID of the default notification policy applied to monitors in this collection that do not explicitly reference one. Monitors that set their own `notification_policy_id` are not overridden.",
	},
}

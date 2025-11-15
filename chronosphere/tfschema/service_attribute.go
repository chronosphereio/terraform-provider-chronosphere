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

var ServiceAttribute = map[string]*schema.Schema{
	"service_slug": {
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "The slug of the service to associate attributes with",
	},
	"name": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Name of the service",
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Human-readable description of the service",
	},
	"team_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Team that owns the service",
	},
	"notification_policy_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Notification policy ID for alerts",
	},
}

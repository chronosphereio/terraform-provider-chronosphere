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

var PagerdutyExternalConnection = map[string]*schema.Schema{
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
	"pagerduty_events_version": {
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"pagerduty_rest_api_key"},
		Description:   "PagerDuty Events API version used to deliver alerts: `PAGERDUTY_EVENTS_VERSION_V1` (legacy) or `PAGERDUTY_EVENTS_VERSION_V2` (default, recommended). Mutually exclusive with `pagerduty_rest_api_key`.",
	},
	"pagerduty_api_key": {
		Type:          schema.TypeString,
		Optional:      true,
		Sensitive:     true,
		ConflictsWith: []string{"pagerduty_rest_api_key"},
		Description:   "PagerDuty Events API integration key used to authenticate alert delivery. Called the routing key in Events v2 and the service key in Events v1. Treat as a secret. Mutually exclusive with `pagerduty_rest_api_key`.",
	},
	"pagerduty_rest_api_key": {
		Type:          schema.TypeString,
		Optional:      true,
		Sensitive:     true,
		ConflictsWith: []string{"pagerduty_api_key", "pagerduty_events_version"},
		Description:   "PagerDuty REST API token used to authenticate incident note polling. Treat as a secret. Mutually exclusive with `pagerduty_api_key` and `pagerduty_events_version`.",
	},
}

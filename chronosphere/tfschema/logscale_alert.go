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

var LogscaleAlert = map[string]*schema.Schema{
	"slug": {
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "Stable identifier for the LogScale alert. Generated from `name` if omitted. Immutable after creation.",
	},
	"repository": {
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "Name of the LogScale repository the alert belongs to. Immutable after creation.",
	},
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Display name of the LogScale alert.",
	},
	"alert_type": Enum{
		Value:       enum.LogscaleAlertType.ToStrings(),
		Required:    true,
		Description: "Type of LogScale alert. `STANDARD` runs the query on a schedule over a time window; `FILTER` evaluates the query against each incoming event.",
	}.Schema(),
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Human-readable description of the alert.",
	},
	"disabled": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "If `true`, the alert will not evaluate or trigger actions.",
	},
	"query": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "LogScale query that the alert evaluates. Example: `level = ERROR | severity > 3 | count(as=numErrors) | numErrors > 500`.",
	},
	"time_window": durationWithDescription(Duration{
		Optional: true,
	}, "Lookback window for the alert query. Required for `STANDARD` alerts, ignored for `FILTER` alerts."),
	"throttle_duration": durationWithDescription(Duration{
		Optional: true,
	}, "Minimum interval between consecutive triggers of the alert. Required for `STANDARD` alerts, optional for `FILTER` alerts."),
	"throttle_field": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Optional field whose value is used to scope throttling, so the alert is throttled per distinct value of this field rather than globally.",
	},
	"action_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		MinItems:    0,
		Description: "Slugs of LogScale actions to invoke when the alert triggers. The alert does not fire if this list is empty.",
	},
	"tags": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional:    true,
		MinItems:    0,
		Description: "Tags attached to the alert for organization and filtering.",
	},
	"run_as_user": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Email of the user that the alert runs on behalf of",
	},
}

func durationWithDescription(d Duration, description string) *schema.Schema {
	s := d.Schema()
	s.Description = description
	return s
}

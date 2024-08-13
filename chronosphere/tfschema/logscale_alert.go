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
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"repository": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
	},
	"alert_type": Enum{
		Value:    enum.LogscaleAlertType.ToStrings(),
		Required: true,
	}.Schema(),
	"description": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"disabled": {
		Type:     schema.TypeBool,
		Optional: true,
	},
	"query": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"time_window": durationWithDescription(Duration{
		Optional: true,
	}, "Required for STANDARD type alerts, ignored for FILTER type alerts"),
	"throttle_duration": durationWithDescription(Duration{
		Optional: true,
	}, "Required for STANDARD type alerts, optional for FILTER type alerts"),
	"throttle_field": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"action_ids": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional: true,
		MinItems: 0,
	},
	"tags": {
		Type: schema.TypeList,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
		Optional: true,
		MinItems: 0,
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

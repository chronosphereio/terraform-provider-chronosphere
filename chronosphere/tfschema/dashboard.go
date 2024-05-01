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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/dashboard"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var Dashboard = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"slug": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"collection_id": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"dashboard_json": {
		Type:                  schema.TypeString,
		Required:              true,
		DiffSuppressFunc:      dashboardJSONDiffSuppress,
		DiffSuppressOnRefresh: true,
	},
}

// dashboardJSONDiffSuppress sanitizes and then diffs two dashboard JSON payloads.
func dashboardJSONDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	sanitizedOld, err := dashboard.SanitizedDashboardJSON(old)
	if err != nil {
		return false
	}

	sanitizedNew, err := dashboard.SanitizedDashboardJSON(new)
	if err != nil {
		return false
	}

	return sanitizedOld == sanitizedNew
}

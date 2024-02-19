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
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/grafana"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var GrafanaDashboard = map[string]*schema.Schema{
	"bucket_id": {
		Type:         schema.TypeString,
		Optional:     true,
		AtLeastOneOf: []string{"bucket_id", "collection_id"},
	},
	"collection_id": {
		Type:         schema.TypeString,
		Optional:     true,
		AtLeastOneOf: []string{"bucket_id", "collection_id"},
	},
	"dashboard_json": {
		Type:             schema.TypeString,
		Required:         true,
		DiffSuppressFunc: grafanaDashboardJSONDiffSuppress,
	},
}

// grafanaDashboardJSONDiffSuppress sanitizes and then diffs two dashboard JSON payloads.
func grafanaDashboardJSONDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	sanitizedOld, err := grafana.SanitizedDashboardJSON(old, grafana.WithUID(""))
	if err != nil {
		return false
	}

	sanitizedNew, err := grafana.SanitizedDashboardJSON(new, grafana.WithUID(""))
	if err != nil {
		return false
	}

	return sanitizedOld == sanitizedNew
}

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
	"encoding/json"
	"fmt"

	xjson "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/x/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var ClassicDashboard = map[string]*schema.Schema{
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
		DiffSuppressFunc: classicDashboardJSONDiffSuppress,
	},
}

// classicDashboardJSONDiffSuppress sanitizes and then diffs two dashboard JSON payloads.
func classicDashboardJSONDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	sanitizedOld, err := SanitizedDashboardJSON(old, WithDashboardUID(""))
	if err != nil {
		return false
	}

	sanitizedNew, err := SanitizedDashboardJSON(new, WithDashboardUID(""))
	if err != nil {
		return false
	}

	return sanitizedOld == sanitizedNew
}

// sanitizedDashboardFields is a set of Grafana dashboard JSON fields that can be sanitized
// when comparing dashboards.
var sanitizedDashboardFields = map[string]bool{
	"id":      true,
	"version": true,
}

// SanitizeOpt is an option passed to SanitizedDashboardJSON.
type SanitizeOpt func(map[string]any) error

// SanitizedDashboardJSON sanitizes a dashboard JSON payload,
// clearing fields irrelevant to reading or upserting dashboards.
func SanitizedDashboardJSON(data string, opts ...SanitizeOpt) (string, error) {
	var dash map[string]any
	if err := xjson.Unmarshal([]byte(data), &dash); err != nil {
		return "", fmt.Errorf("invalid dashboard JSON: %w", err)
	}

	for _, opt := range opts {
		if err := opt(dash); err != nil {
			return "", err
		}
	}

	for field := range sanitizedDashboardFields {
		delete(dash, field)
	}

	sanitizedJSON, err := json.Marshal(dash)
	if err != nil {
		return "", fmt.Errorf("invalid dashboard JSON: %w", err)
	}

	return string(sanitizedJSON), nil
}

// WithDashboardUID sets or clears the dashboard `uid` field.
// If uid is given, the `uid` field is set to it in the JSON payload.
// Otherwise, the payload omits the `uid` field.
func WithDashboardUID(uid string) SanitizeOpt {
	return func(dash map[string]any) error {
		if uid == "" {
			delete(dash, "uid")
		} else {
			dash["uid"] = uid
		}
		return nil
	}
}

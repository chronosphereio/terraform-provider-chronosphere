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

package grafana

import (
	"encoding/json"
	"fmt"

	xjson "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/x/json"
)

// sanitizedDashboardFields is a set of Grafana dashboard JSON fields that can be sanitized
// when comparing dashboards.
var sanitizedDashboardFields = map[string]bool{
	"id":      true,
	"version": true,
}

// SanitizeOpt is an option passed to SanitizedDashboardJSON.
type SanitizeOpt func(map[string]any) error

// SanitizedDashboardJSON sanitizes a Grafana dashboard JSON payload,
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

// WithUID sets or clears the dashboard `uid` field.
// If uid is given, the `uid` field is set to it in the JSON payload.
// Otherwise, the payload omits the `uid` field.
func WithUID(uid string) SanitizeOpt {
	return func(dash map[string]any) error {
		if uid == "" {
			delete(dash, "uid")
		} else {
			dash["uid"] = uid
		}
		return nil
	}
}

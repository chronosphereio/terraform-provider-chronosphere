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

package dashboard

import (
	"encoding/json"
	"errors"
	"fmt"

	xjson "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/x/json"
)

// SanitizedDashboardJSON sanitizes a Dashboard JSON payload.
func SanitizedDashboardJSON(data string) (string, error) {
	var dash map[string]any
	if err := xjson.Unmarshal([]byte(data), &dash); err != nil {
		return "", fmt.Errorf("invalid dashboard JSON: %w", err)
	}

	rawMetadata, ok := dash["metadata"]
	if ok {
		metadata, ok := rawMetadata.(map[string]any)
		if !ok {
			return "", errors.New("invalid dashboard metadata")
		}
		delete(metadata, "created_at")
		delete(metadata, "updated_at")
		delete(metadata, "version")

		dash["metadata"] = metadata
	}

	ignoreNullAndEmptyFields(dash)

	sanitizedJSON, err := json.Marshal(dash)
	if err != nil {
		return "", fmt.Errorf("invalid dashboard JSON: %w", err)
	}

	return string(sanitizedJSON), nil
}

// Ignore fields with null or empty values so that null and empty are treated as equivalent in sanitized
// payloads.
func ignoreNullAndEmptyFields(jsonObj map[string]any) {
	for k, v := range jsonObj {
		if v == nil || v == "" {
			delete(jsonObj, k)
		}

		if nestedObj, ok := v.(map[string]any); ok {
			ignoreNullAndEmptyFields(nestedObj)
		}
	}
}

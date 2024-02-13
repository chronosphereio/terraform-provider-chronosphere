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

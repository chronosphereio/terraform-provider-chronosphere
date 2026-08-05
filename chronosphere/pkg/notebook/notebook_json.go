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

// Package notebook provides helpers for working with the raw JSON payload of
// a Chronosphere notebook.
package notebook

import (
	"encoding/json"
	"fmt"

	xjson "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/x/json"
)

// CanonicalNotebookJSON returns a canonical encoding of a notebook JSON
// payload: object keys sorted and insignificant whitespace removed. Two
// payloads that differ only in formatting produce the same output, which lets
// callers compare them for semantic equality.
//
// Unlike dashboards, the config API stores and returns notebook_json verbatim
// and populates no fields inside it, so no fields need stripping here. If that
// changes, this is where the sanitizing would go.
func CanonicalNotebookJSON(data string) (string, error) {
	// Decoded into any rather than map[string]any so a payload whose top level
	// is not an object still round-trips instead of erroring.
	var notebook any
	if err := xjson.Unmarshal([]byte(data), &notebook); err != nil {
		return "", fmt.Errorf("invalid notebook JSON: %w", err)
	}

	// json.Marshal sorts map keys, which is what makes the output canonical.
	canonical, err := json.Marshal(notebook)
	if err != nil {
		return "", fmt.Errorf("invalid notebook JSON: %w", err)
	}

	return string(canonical), nil
}

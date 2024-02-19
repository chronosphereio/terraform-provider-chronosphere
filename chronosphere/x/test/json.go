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

package xtest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

// JSONMarshalIndentedString returns an indentend JSON string.
func JSONMarshalIndentedString(t *testing.T, str string) string {
	var unmarshalled any
	err := json.Unmarshal([]byte(str), &unmarshalled)
	require.NoError(t, err)
	pretty, err := json.MarshalIndent(unmarshalled, "", "  ")
	require.NoError(t, err)
	return string(pretty)
}

// JSONMarshalIndentedValue returns an indentend JSON string.
func JSONMarshalIndentedValue(t *testing.T, value any) string {
	pretty, err := json.MarshalIndent(value, "", "  ")
	require.NoError(t, err)
	return string(pretty)
}

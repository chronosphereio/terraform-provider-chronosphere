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

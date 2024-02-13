package xtest

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var replacementPatterns = map[string]*regexp.Regexp{
	// UUID is any version UUID
	`$uuid`: regexp.MustCompile(`(` +
		`([a-f0-9]{32,36})|` +
		`([a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})|` +
		`([A-F0-9]{8}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{12})` +
		`)`),

	// Datetime is an ISO8601 date time string with optional decimal part
	`$datetime`: regexp.MustCompile(`(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])(\.[0-9]+)?(Z)?`),

	// Time is a RFC 1123 date time string.
	`$time`: regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}.[0-9]{1,20}.[-,\+][0-9]{4} [A-Z]{3} [a-z]{1,5}=\+[0-9].[0-9]{1,20}`),
}

// RegexpEscapedPatternWithTypes returns a regexp pattern with specific
// type aliases replaced with their regexp patterns and all other characters
// escaped.
// The supported types are:
//   - string: $uuid
//     pattern: "?[a-zA-Z0-9-]{32,36}"?
func RegexpEscapedPatternWithTypes(t *testing.T, input string) *regexp.Regexp {
	replacerInput := make([]string, 0, len(replacementPatterns)*2)
	for k, v := range replacementPatterns {
		replacerInput = append(replacerInput, fmt.Sprintf("\\%s", k))
		replacerInput = append(replacerInput, v.String())
	}
	replacer := strings.NewReplacer(replacerInput...)

	escapedWithTypesReplaced := replacer.Replace(regexp.QuoteMeta(input))
	re, err := regexp.Compile(escapedWithTypesReplaced)
	require.NoError(t, err)

	return re
}

package json

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// This delegates to json.Unmarshal, returning more friendly error messages that include line numbers.
func Unmarshal(data []byte, v any) error {
	err := json.Unmarshal(data, v)
	if err == nil {
		return nil
	}

	if _, ok := err.(*json.InvalidUnmarshalError); ok {
		return err
	}

	if errStr := err.Error(); errStr == "EOF" || errStr == "unexpected EOF" || errStr == "unexpected end of JSON input" {
		return &UnmarshalError{error: "unexpected end of input"}
	}

	dataStr := string(data)
	offsetLine := func(offset int64) int {
		if int64(len(dataStr)) < offset {
			return 0
		}

		start := strings.LastIndex(dataStr[:offset], "\n") + 1
		return strings.Count(dataStr[:start], "\n") + 1
	}

	if syntaxErr, ok := err.(*json.SyntaxError); ok {
		return &UnmarshalError{error: trimErrorPrefix(syntaxErr.Error()), line: offsetLine(syntaxErr.Offset)}
	}
	if typeErr, ok := err.(*json.UnmarshalTypeError); ok {
		return &UnmarshalError{error: trimErrorPrefix(typeErr.Error()), line: offsetLine(typeErr.Offset)}
	}

	if errStr := err.Error(); errStr == "EOF" || errStr == "unexpected EOF" {
		return &UnmarshalError{error: "unexpected end of input"}
	}

	return &UnmarshalError{error: trimErrorPrefix(err.Error())}
}

// UnmarshalError is an error unmarshaling JSON data.
type UnmarshalError struct {
	error string
	line  int // Optional: is 0 if no relevant line
}

// Error returns the error string.
func (d *UnmarshalError) Error() string {
	if d.line > 0 {
		return fmt.Sprintf("%s (line %d)", d.error, d.line)
	}

	return d.error
}

func trimErrorPrefix(jsonErr string) string {
	if strings.HasPrefix(jsonErr, "json: ") {
		return jsonErr[6:]
	}

	return jsonErr
}

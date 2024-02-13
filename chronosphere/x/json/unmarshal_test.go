package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalIntoMap(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output map[string]any
		err    string
	}{
		{
			input: ``,
			err:   "unexpected end of input",
		},
		{
			input:  `{}`,
			output: map[string]any{},
		},
		{
			input: `[]`,
			err:   "cannot unmarshal array into Go value of type map[string]interface {}",
		},
		{
			input:  `{"hello":"world","asdf":123}`,
			output: map[string]any{"hello": "world", "asdf": 123.0},
		},
		{
			input: `foo`,
			err:   "invalid character 'o' in literal false (expecting 'a') (line 1)",
		},
		{
			input: `{`,
			err:   "unexpected end of input",
		},
		{
			input: `{"hello":"world}`,
			err:   "unexpected end of input",
		},
		{
			input: `{
"foo":"bar",
}`,
			err: "invalid character '}' looking for beginning of object key string (line 3)",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.input, func(t *testing.T) {
			var out map[string]any
			err := Unmarshal([]byte(tt.input), &out)

			if tt.err != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.output, out)
			}
		})
	}
}

func TestUnmarshalIntoStruct(t *testing.T) {
	type cat struct {
		Name string `json:"name"`
		Age  uint   `json:"age"`
	}

	testCases := []struct {
		name   string
		input  string
		output cat
		err    string
	}{
		{
			input: ``,
			err:   "unexpected end of input",
		},
		{
			input: `{}`,
		},
		{
			input: `[]`,
			err:   "cannot unmarshal array into Go value of type json.cat (line 1)",
		},
		{
			input:  `{"name":"ume","age":8}`,
			output: cat{Name: "ume", Age: 8},
		},
		{
			input: `{"name":"ume","age":"oops"}`,
			err:   "cannot unmarshal string into Go struct field cat.age of type uint (line 1)",
		},
		{
			input: `foo`,
			err:   "invalid character 'o' in literal false (expecting 'a') (line 1)",
		},
		{
			input: `{`,
			err:   "unexpected end of input",
		},
		{
			input: `{"hello":"world}`,
			err:   "unexpected end of input",
		},
		{
			input: `{
"foo":"bar",
}`,
			err: "invalid character '}' looking for beginning of object key string (line 3)",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.input, func(t *testing.T) {
			var out cat
			err := Unmarshal([]byte(tt.input), &out)

			if tt.err != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.output, out)
			}
		})
	}
}

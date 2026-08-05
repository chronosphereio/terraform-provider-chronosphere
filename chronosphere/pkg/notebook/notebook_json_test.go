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

package notebook

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanonicalNotebookJSON(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "sorts object keys",
			in:   `{"b": 1, "a": 2}`,
			want: `{"a":2,"b":1}`,
		},
		{
			name: "strips insignificant whitespace",
			in: `{
				"name" : "n" ,
				"cells": [ ]
			}`,
			want: `{"cells":[],"name":"n"}`,
		},
		{
			name: "sorts nested object keys but preserves array order",
			in:   `{"cells":[{"z":1,"a":2},{"b":3}]}`,
			want: `{"cells":[{"a":2,"z":1},{"b":3}]}`,
		},
		{
			name: "non-object top level round-trips",
			in:   `[2, 1]`,
			want: `[2,1]`,
		},
		{
			name: "scalar top level round-trips",
			in:   `  "just a string"  `,
			want: `"just a string"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CanonicalNotebookJSON(tt.in)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCanonicalNotebookJSONEquivalence(t *testing.T) {
	// The point of canonicalization: two payloads that a human would call the
	// same must produce byte-identical output.
	compact := `{"cells":[{"kind":"markdown","text":"hi"}],"version":1}`
	pretty := `{
	  "version": 1,
	  "cells": [
	    {
	      "text": "hi",
	      "kind": "markdown"
	    }
	  ]
	}`

	gotCompact, err := CanonicalNotebookJSON(compact)
	require.NoError(t, err)
	gotPretty, err := CanonicalNotebookJSON(pretty)
	require.NoError(t, err)

	assert.Equal(t, gotCompact, gotPretty)
}

func TestCanonicalNotebookJSONInvalid(t *testing.T) {
	for _, in := range []string{"", "{", `{"a": }`, "not json"} {
		_, err := CanonicalNotebookJSON(in)
		assert.ErrorContains(t, err, "invalid notebook JSON", "input %q", in)
	}
}

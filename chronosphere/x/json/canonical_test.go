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

package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanonical(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr string
	}{
		{
			name:  "sorts keys and drops whitespace",
			input: "{\n  \"b\": [1, 2],\n  \"a\": {\"d\": true, \"c\": null}\n}",
			want:  `{"a":{"c":null,"d":true},"b":[1,2]}`,
		},
		{
			name:  "keeps large integer literals",
			input: `{"id": 9007199254740993}`,
			want:  `{"id":9007199254740993}`,
		},
		{
			name:  "keeps array order",
			input: `[2, 1]`,
			want:  `[2,1]`,
		},
		{
			name:  "scalar",
			input: ` "x" `,
			want:  `"x"`,
		},
		{
			name:    "empty input",
			input:   "",
			wantErr: "EOF",
		},
		{
			name:    "truncated",
			input:   `{"a":`,
			wantErr: "EOF",
		},
		{
			name:    "trailing garbage",
			input:   `{"a":1}}`,
			wantErr: "trailing data after JSON value",
		},
		{
			name:    "two values",
			input:   `{} {}`,
			wantErr: "trailing data after JSON value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Canonical(tt.input)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

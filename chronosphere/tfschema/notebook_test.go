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

package tfschema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotebookJSONDiffSuppress(t *testing.T) {
	tests := []struct {
		name string
		old  string
		new  string
		want bool
	}{
		{
			name: "identical",
			old:  `{"type":"doc","content":[]}`,
			new:  `{"type":"doc","content":[]}`,
			want: true,
		},
		{
			name: "reordered keys and whitespace",
			old:  `{"type":"doc","content":[{"type":"paragraph","attrs":{"a":1,"b":2}}]}`,
			new:  "{\n  \"content\": [ {\"attrs\": {\"b\": 2, \"a\": 1}, \"type\": \"paragraph\"} ],\n  \"type\": \"doc\"\n}",
			want: true,
		},
		{
			name: "large integer survives",
			old:  `{"id":9007199254740993}`,
			new:  `{"id": 9007199254740993}`,
			want: true,
		},
		{
			name: "different large integers are not equal",
			old:  `{"id":9007199254740993}`,
			new:  `{"id":9007199254740992}`,
			want: false,
		},
		{
			name: "different content",
			old:  `{"type":"doc","content":[]}`,
			new:  `{"type":"doc","content":[{"type":"paragraph"}]}`,
			want: false,
		},
		{
			name: "array order matters",
			old:  `{"content":[1,2]}`,
			new:  `{"content":[2,1]}`,
			want: false,
		},
		{
			name: "trailing garbage is not suppressed against valid state",
			old:  `{"a":1}`,
			new:  `{"a":1}}`,
			want: false,
		},
		{
			name: "invalid old JSON",
			old:  `{`,
			new:  `{}`,
			want: false,
		},
		{
			name: "invalid new JSON",
			old:  `{}`,
			new:  `not json`,
			want: false,
		},
		{
			name: "empty strings",
			old:  "",
			new:  "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, notebookJSONDiffSuppress("notebook_json", tt.old, tt.new, nil))
		})
	}
}

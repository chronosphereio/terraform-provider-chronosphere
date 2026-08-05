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
			old:  `{"a":1}`,
			new:  `{"a":1}`,
			want: true,
		},
		{
			name: "differs only in key order",
			old:  `{"a":1,"b":2}`,
			new:  `{"b":2,"a":1}`,
			want: true,
		},
		{
			name: "differs only in whitespace",
			old:  `{"a":1}`,
			new:  "{\n  \"a\": 1\n}",
			want: true,
		},
		{
			name: "different value",
			old:  `{"a":1}`,
			new:  `{"a":2}`,
			want: false,
		},
		{
			name: "different array order is a real change",
			old:  `{"cells":[1,2]}`,
			new:  `{"cells":[2,1]}`,
			want: false,
		},
		{
			// An unparseable value must never be suppressed: Terraform should
			// show the diff and let the API reject it with a real error.
			name: "invalid new",
			old:  `{"a":1}`,
			new:  `{"a":`,
			want: false,
		},
		{
			name: "invalid old",
			old:  `{"a":`,
			new:  `{"a":1}`,
			want: false,
		},
		{
			name: "both empty",
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

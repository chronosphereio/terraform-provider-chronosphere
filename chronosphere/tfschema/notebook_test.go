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

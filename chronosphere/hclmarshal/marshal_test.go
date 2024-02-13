package hclmarshal

import (
	"bytes"
	"testing"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalerResourceBlock(t *testing.T) {
	tests := []struct {
		msg     string
		marshal func(b *Block)
		wantErr string
		want    string
	}{
		{
			msg: "required empty values",
			marshal: func(b *Block) {
				b.Add("bool", false)
				b.Add("float", 0.0)
				b.Add("int", 0)
				b.Add("str", "")
				b.Add("list", []string{})
				b.Add("list_null", []string(nil))
				b.Add("map", map[string]int{})
				b.Add("map_null", map[string]int(nil))
			},
			want: `
  bool      = false
  float     = 0
  int       = 0
  str       = ""
  list      = []
  list_null = null

  map = {}

  map_null = null
`,
		},
		{
			msg: "optional empty values",
			marshal: func(b *Block) {
				b.AddOptional("bool", false)
				b.AddOptional("float", 0.0)
				b.AddOptional("int", 0)
				b.AddOptional("str", "")
				b.AddOptional("list", []string{})
				b.AddOptional("list_null", []string(nil))
				b.AddOptional("map", map[string]int{})
				b.AddOptional("map_null", map[string]int(nil))
			},
			want: `
`,
		},
		{
			msg: "required non-empty values",
			marshal: func(b *Block) {
				b.Add("bool", true)
				b.Add("float", 1.3)
				b.Add("int", 1)
				b.Add("str", "s")
				b.Add("list", []string{"a", "b", "c"})
				b.Add("map", map[string]int{"a": 1, "b": 2})
			},
			want: `
  bool  = true
  float = 1.3
  int   = 1
  str   = "s"
  list  = ["a", "b", "c"]

  map = {
    a = 1
    b = 2
  }
`,
		},
		{
			msg: "optional non-empty values",
			marshal: func(b *Block) {
				b.AddOptional("bool", true)
				b.AddOptional("float", 1.3)
				b.AddOptional("int", 1)
				b.AddOptional("str", "s")
				b.AddOptional("list", []string{"a", "b", "c"})
				b.AddOptional("map", map[string]int{"a": 1, "b": 2})
			},
			want: `
  bool  = true
  float = 1.3
  int   = 1
  str   = "s"
  list  = ["a", "b", "c"]

  map = {
    a = 1
    b = 2
  }
`,
		},
		{
			msg: "add refs",
			marshal: func(b *Block) {
				b.AddRef("ref_str", tfid.Slug("foobar"))
				b.AddRef("ref_path", tfid.Ref{
					Type:  "chronosphere_test_resource",
					ID:    "baz",
					Field: "data",
				}.AsID())
				b.AddRef("ref_null", tfid.ID{})
				b.AddTFRef("tf_ref", tfid.Ref{
					Type:       "chronosphere_test_resource",
					ID:         "qux",
					Datasource: true,
				})
				b.AddRefs("refs_null", nil /* ids */)
				b.AddRefs("refs_empty", []tfid.ID{{}, {}})
				b.AddRefs("refs_mixed", []tfid.ID{
					{},
					tfid.Slug("slug-ref"),
					{},
					tfid.Ref{
						Type: "chronosphere_test_resource",
						ID:   "qux",
					}.AsID(),
				})
			},
			want: `
  ref_str    = "foobar"
  ref_path   = chronosphere_test_resource.baz.data
  tf_ref     = data.chronosphere_test_resource.qux.id
  refs_mixed = ["slug-ref", chronosphere_test_resource.qux.id]
`,
		},
		{
			msg: "nested blocks",
			marshal: func(b *Block) {
				b.Add("attr", 1)

				b.AddNewLine()
				nested := b.AddBlock("nested")
				nested.Add("attr", 2)

				nested2 := nested.AddBlock("nested2")
				nested2.Add("attr", 3)

				b.Add("attr_after_nest", true)
			},
			want: `
  attr = 1

  nested {
    attr = 2
    nested2 {
      attr = 3
    }
  }
  attr_after_nest = true
`,
		},
		{
			msg: "multi-line string",
			marshal: func(b *Block) {
				multiline := `{
  "title": "title",
  "nest": {
    "foo": "bar"
  }
}
`
				b.Add("multi_str", multiline)
				nested := b.AddBlock("nested")
				nested.Add("another_multi", multiline)
			},
			want: `
  multi_str = <<-EOF
  {
    "title": "title",
    "nest": {
      "foo": "bar"
    }
  }
  EOF
  nested {
    another_multi = <<-EOF
    {
      "title": "title",
      "nest": {
        "foo": "bar"
      }
    }
    EOF
  }
`,
		},
		{
			msg: "multi-line no trailing newline",
			marshal: func(b *Block) {
				b.Add("multi_line_str", "foo\nbar")
			},
			want: `
  multi_line_str = chomp(<<-EOF
  foo
  bar
  EOF
  )
`,
		},
		{
			msg: "multi-line with EOF",
			marshal: func(b *Block) {
				b.Add("use_heredoc", "foo\nbar\n")
				b.Add("use_normal_str", "foo\nEOF\nbar\n")
			},
			want: `
  use_heredoc = <<-EOF
  foo
  bar
  EOF
  use_normal_str = "foo\nEOF\nbar\n"
`,
		},
		{
			msg: "unserializable type",
			marshal: func(b *Block) {
				b.Add("chan", make(chan string))
				b.AddOptional("chan_opt", make(chan string))
			},
			wantErr: "no cty.Type for chan",
		},
		{
			msg: "no additional whitespace at beginning or end of block",
			marshal: func(b *Block) {
				b.AddOptional("map", map[string]int{"a": 1, "b": 2})
			},
			want: `
  map = {
    a = 1
    b = 2
  }
`,
		},
		{
			msg: "attribute with function call",
			marshal: func(b *Block) {
				b.AddFuncCall("f1", "noArgs")
				b.AddFuncCall("f2", "singleArg", "str")
				b.AddFuncCall("f3", "multiArgs", 1, true, "str")
			},
			want: `
  f1 = noArgs()
  f2 = singleArg("str")
  f3 = multiArgs(1, true, "str")
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			m := New()
			b := m.AddResource("chronosphere_test_resource", "foobar")
			tt.marshal(b)

			buf := &bytes.Buffer{}
			err := m.MarshalTo(buf)
			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
				return
			}

			require.NoError(t, err)
			want := "\n" + `resource "chronosphere_test_resource" "foobar" {` + tt.want + "}\n"
			assert.Equal(t, want, buf.String())
		})
	}
}

func TestMarshalerMultipleBlocks(t *testing.T) {
	m := New()

	b1 := m.AddData("chronosphere_test_resource", "foobar")
	b1.Add("attr", 1)

	b2 := m.AddResource("chronosphere_test_resource", "foobar2")
	b2.Add("attr", 2)

	want := `
data "chronosphere_test_resource" "foobar" {
  attr = 1
}

resource "chronosphere_test_resource" "foobar2" {
  attr = 2
}
`

	buf := &bytes.Buffer{}
	require.NoError(t, m.MarshalTo(buf))
	assert.Equal(t, want, buf.String())
}

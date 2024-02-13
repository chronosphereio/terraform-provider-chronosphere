package intschematag

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundtrip(t *testing.T) {
	tests := []struct {
		msg                string
		tag                Tag
		structField        any
		ignoreResourceData bool
	}{
		{
			msg: "empty tag",
			tag: Tag{},
			structField: struct {
				Field int `intschema:""`
			}{},
		},
		{
			msg: "only name",
			tag: Tag{
				TFName: "foo",
			},
			structField: struct {
				Field int `intschema:"foo"`
			}{},
		},
		{
			msg: "optional",
			tag: Tag{
				TFName:   "foo",
				Optional: true,
			},
			structField: struct {
				Field int `intschema:"foo,optional"`
			}{},
		},
		{
			msg: "computed",
			tag: Tag{
				TFName:   "foo",
				Computed: true,
			},
			structField: struct {
				Field int `intschema:"foo,computed"`
			}{},
		},
		{
			msg: "file",
			tag: Tag{
				TFName: "foo",
				File:   true,
			},
			structField: struct {
				Field string `intschema:"foo,file"`
			}{},
			ignoreResourceData: true,
		},
		{
			msg: "internal field",
			tag: Tag{
				TFName: InternalFieldName,
			},
			structField: struct {
				Field string `intschema:"-"`
			}{},
			ignoreResourceData: true,
		},
		{
			msg: "all fields set",
			tag: Tag{
				TFName:            "foo",
				Optional:          true,
				ListEncodedObject: true,
			},
			structField: struct {
				Field int `intschema:"foo,optional,list_encoded_object"`
			}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			structFieldType := reflect.ValueOf(tt.structField).Type().Field(0)

			assert.Equal(t, string(structFieldType.Tag), tt.tag.Marshal(), "Marshal")
			assert.Equal(t, tt.tag, Unmarshal(structFieldType), "Unmarshal")
			assert.Equal(t, tt.ignoreResourceData, tt.tag.IgnoreResourceData(), "IgnoreResourceData")
		})
	}
}

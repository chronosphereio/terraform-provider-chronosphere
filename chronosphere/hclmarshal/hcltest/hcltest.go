package hcltest

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

// Marshaler is a type that can be marshalled to HCL, and is implementd by intschema generated types.
type Marshaler interface {
	MarshalHCL(w io.Writer) error
}

// MarshalString marshals the given marshalable type to a string.
func MarshalString(m Marshaler) (string, error) {
	buf := &bytes.Buffer{}
	err := m.MarshalHCL(buf)
	return buf.String(), err
}

// MustMarshalString is the same as MarshalString, but requires success.
func MustMarshalString(t testing.TB, m Marshaler) string {
	s, err := MarshalString(m)
	require.NoError(t, err, "MarshalString failed")
	return s
}

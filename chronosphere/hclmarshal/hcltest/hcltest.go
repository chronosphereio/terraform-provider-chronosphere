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

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

package tfid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSlug(t *testing.T) {
	id := Slug("foo")
	require.Equal(t, TypeSlug, id.Type())
	require.Equal(t, "foo", id.Slug())
	require.PanicsWithValue(t, "ID is not a local ref", func() { id.LocalRef() })
}

func TestLocalRef(t *testing.T) {
	r := Ref{
		Datasource: true,
		Type:       "chronosphere_bucket",
		ID:         "foo",
		Field:      "id",
	}
	id := LocalRef(r)
	assert.Equal(t, id, r.AsID())
	require.Equal(t, TypeLocalRef, id.Type())
	require.Equal(t, r, id.LocalRef())
	require.PanicsWithValue(t, "ID is not a slug", func() { id.Slug() })
}

func TestEmpty(t *testing.T) {
	id := ID{}
	require.Equal(t, TypeEmpty, id.Type())
	require.Equal(t, "", id.Slug())
	require.Equal(t, Ref{}, id.LocalRef())

	// Empty values are always considered empty.
	require.Equal(t, TypeEmpty, Slug("").Type())
	require.Equal(t, TypeEmpty, LocalRef(Ref{}).Type())
}

func TestSafeID(t *testing.T) {
	tests := []struct {
		slug string
		id   string
	}{
		{"foo1bar", "foo1bar"},
		{"foo-bar", "foo_bar"},
		{"-foo-bar", "_foo_bar"},
		{"1foo2", "_1foo2"},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			assert.Equal(t, tt.id, SafeID(tt.slug))
		})
	}
}

// Copyright 2023 Chronosphere Inc.
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

package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	s := New[string]()
	s.Add("a")
	assert.ElementsMatch(t, []string{"a"}, s.Elements())

	s.Add("a")
	assert.ElementsMatch(t, []string{"a"}, s.Elements())

	s.Add("b")
	assert.ElementsMatch(t, []string{"a", "b"}, s.Elements())
}

func TestAddAll(t *testing.T) {
	t.Parallel()

	s := New[string]()
	s.AddAll("a", "b", "a")
	assert.ElementsMatch(t, []string{"a", "b"}, s.Elements())

	s.AddAll("c", "b")
	assert.ElementsMatch(t, []string{"a", "b", "c"}, s.Elements())

	s.AddAll()
	assert.ElementsMatch(t, []string{"a", "b", "c"}, s.Elements())
}

func TestHas(t *testing.T) {
	t.Parallel()

	intSet := New(1, 2)
	intSet.Add(3)

	assert.True(t, intSet.Has(1))
	assert.True(t, intSet.Has(2))
	assert.True(t, intSet.Has(3))

	assert.False(t, intSet.Has(0))
	assert.False(t, intSet.Has(4))
}

func TestRemove(t *testing.T) {
	t.Parallel()

	s := New(1, 2, 3)
	s.Remove(2)
	assert.ElementsMatch(t, []int{1, 3}, s.Elements())

	s.Remove(2)
	assert.ElementsMatch(t, []int{1, 3}, s.Elements())

	s.Remove(1)
	assert.ElementsMatch(t, []int{3}, s.Elements())

	s.Remove(12345)
	assert.ElementsMatch(t, []int{3}, s.Elements())
}

func TestRemoveAll(t *testing.T) {
	t.Parallel()

	s := New(1, 2, 3, 4)
	s.RemoveAll(2, 4, 6)
	assert.ElementsMatch(t, []int{1, 3}, s.Elements())

	s.RemoveAll(2, 4, 6)
	assert.ElementsMatch(t, []int{1, 3}, s.Elements())

	s.RemoveAll()
	assert.ElementsMatch(t, []int{1, 3}, s.Elements())

	s.RemoveAll(1)
	assert.ElementsMatch(t, []int{3}, s.Elements())
}

func TestClear(t *testing.T) {
	t.Parallel()

	s := New(1, 2, 3)
	s.Clear()
	assert.Equal(t, 0, s.Len())

	s.AddAll(4, 5)
	s.Clear()
	assert.Equal(t, 0, s.Len())
}

func TestElements(t *testing.T) {
	t.Parallel()

	assert.ElementsMatch(t, []int{}, New[int]().Elements())
	assert.ElementsMatch(t, []int{1}, New(1).Elements())
	assert.ElementsMatch(t, []int{1, 2, 3}, New(1, 2, 3).Elements())
}

func TestIterate(t *testing.T) {
	stringSet := New("a", "b", "c")

	t.Run("full iteration", func(t *testing.T) {
		var visited []string
		stringSet.Iterate(func(value string) bool {
			visited = append(visited, value)
			return true
		})
		assert.ElementsMatch(t, []string{"a", "b", "c"}, visited)
	})

	t.Run("partial iteration", func(t *testing.T) {
		count := 0
		var visited []string

		// Iterate over 2 values
		stringSet.Iterate(func(value string) bool {
			visited = append(visited, value)
			count++
			return count < 2
		})
		assert.Len(t, visited, 2)

		unexpectedVisited := Difference[string](New(visited...), stringSet)
		assert.Len(t, unexpectedVisited.Elements(), 0)
	})
}

func TestLen(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, New[string]().Len())
	assert.Equal(t, 1, New("a").Len())
	assert.Equal(t, 2, New("a", "b").Len())
}

func TestDiff(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    MapSet[string]
		b    MapSet[string]
		want MapSet[string]
	}{
		{
			name: "some same some different",
			a:    New("x", "y", "z"),
			b:    New("w", "x", "y"),
			want: New("w", "z"),
		},
		{
			name: "all same",
			a:    New("x", "y", "z"),
			b:    New("x", "y", "z"),
			want: New[string](),
		},
		{
			name: "all different",
			a:    New("t", "u", "v"),
			b:    New("x", "y", "z"),
			want: New("t", "u", "v", "x", "y", "z"),
		},
		{
			name: "one set empty",
			a:    New[string](),
			b:    New("x", "y", "z"),
			want: New("x", "y", "z"),
		},
		{
			name: "both empty",
			a:    New[string](),
			b:    New[string](),
			want: New[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.a.Diff(tt.b))
			assert.Equal(t, tt.want, tt.b.Diff(tt.a))
		})
	}
}

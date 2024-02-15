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

func TestEqual(t *testing.T) {
	t.Parallel()

	assert.True(t, Equal[int](New[int](), New[int]()))
	assert.True(t, Equal[int](New(1), New(1)))
	assert.True(t, Equal[int](New(1, 2), New(1, 2)))

	assert.False(t, Equal[int](New(1), New[int]()))
	assert.False(t, Equal[int](New(1, 2), New[int]()))

	assert.False(t, Equal[int](New[int](), New(1)))
	assert.False(t, Equal[int](New[int](), New(1, 2)))
}

func TestSortedElements(t *testing.T) {
	t.Parallel()

	assert.Nil(t, SortedElements[int](New[int]()))
	assert.Equal(t, []int{1}, SortedElements[int](New(1)))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, SortedElements[int](New(1, 3, 5, 2, 4)))
}

func TestCopyInto(t *testing.T) {
	t.Parallel()

	{
		dest := New[int]()
		src := New(1, 2, 3)
		CopyInto[int](dest, src)
		assert.ElementsMatch(t, []int{1, 2, 3}, dest.Elements())
		assert.ElementsMatch(t, []int{1, 2, 3}, src.Elements())
	}

	{
		dest := New(1, 2, 3)
		src := New(3, 4, 5)
		CopyInto[int](dest, src)
		assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, dest.Elements())
		assert.ElementsMatch(t, []int{3, 4, 5}, src.Elements())
	}
}

func TestDifference(t *testing.T) {
	t.Parallel()

	assert.Empty(t, Difference[int](New[int](), New[int]()).Elements())
	assert.ElementsMatch(t, []int{1, 2, 3}, Difference[int](New(1, 2, 3), New[int]()).Elements())
	assert.ElementsMatch(t, []int{1, 3, 5}, Difference[int](New(1, 2, 3, 4, 5), New(2, 4, 6)).Elements())
	assert.Empty(t, Difference[int](New[int](), New(1, 2, 3)).Elements())
}

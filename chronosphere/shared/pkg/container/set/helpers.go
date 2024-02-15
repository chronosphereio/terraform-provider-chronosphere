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
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Equal returns true if two sets have the same elements.
func Equal[T comparable](a, b Set[T]) bool {
	if a.Len() != b.Len() {
		return false
	}

	equal := true
	a.Iterate(func(elem T) bool {
		if !b.Has(elem) {
			equal = false
		}
		return equal
	})

	return equal
}

// SortedElements returns the set's entries as an ordered slice.
func SortedElements[T constraints.Ordered](s Set[T]) []T {
	elems := s.Elements()
	if len(elems) == 0 {
		return elems
	}

	slices.Sort(elems)

	return elems
}

// CopyInto copies all the elements from src into dest.
func CopyInto[T comparable](dest, src Set[T]) {
	src.Iterate(func(elem T) bool {
		dest.Add(elem)
		return true
	})
}

// Difference returns a new set representing the difference between sets a and b,
// that is, elements which only exist in set a and not set b.
// The first argument is a function which returns a new empty set into which the resulting elements are added.
func Difference[T comparable](a, b Set[T]) Set[T] {
	diffSet := New[T]()

	CopyInto[T](diffSet, a)

	b.Iterate(func(elem T) bool {
		diffSet.Remove(elem)
		return true
	})

	return diffSet
}

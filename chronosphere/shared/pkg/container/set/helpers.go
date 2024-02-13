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

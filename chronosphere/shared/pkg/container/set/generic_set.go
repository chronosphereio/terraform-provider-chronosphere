package set

// Set is a mutable unordered set of comparable elements.
// This is a minimal interface for a set type.
type Set[T comparable] interface {
	// Add puts an element into the set.
	Add(elem T)

	// AddAll puts many elements into the set.
	AddAll(elems ...T)

	// Remove removes an element from the set.
	Remove(elem T)

	// RemoveAll removes all the given elements from the set.
	RemoveAll(elems ...T)

	// Clear removes all elements from the set.
	Clear()

	// Has returns true if the set contains the given element.
	Has(elem T) bool

	// Elements returns all elements in the set in a non-deterministic order.
	// For empty sets, this may return either nil or an empty slice.
	Elements() []T

	// Iterate invokes the callback for every element in the set.
	// Values are iterated in a non-deterministic order.
	// If the callback returns false then iteration will stop.
	Iterate(callback func(elem T) bool)

	// Len returns the number of elements in the set.
	Len() int
}

// New creates a new set with the given elements.
func New[T comparable](elems ...T) MapSet[T] {
	if len(elems) == 0 {
		return MapSet[T]{elems: make(map[T]struct{})}
	}

	newSet := MapSet[T]{elems: make(map[T]struct{}, len(elems))}
	newSet.AddAll(elems...)

	return newSet
}

// MapSet is an unordered Set backed by the builtin map type.
type MapSet[T comparable] struct {
	elems map[T]struct{}
}

func (m MapSet[T]) Add(elem T) {
	m.elems[elem] = struct{}{}
}

func (m MapSet[T]) AddAll(elems ...T) {
	for _, v := range elems {
		m.elems[v] = struct{}{}
	}
}

func (m MapSet[T]) Remove(elem T) {
	delete(m.elems, elem)
}

func (m MapSet[T]) RemoveAll(elems ...T) {
	for _, v := range elems {
		delete(m.elems, v)
	}
}

func (m MapSet[T]) Clear() {
	for k := range m.elems {
		delete(m.elems, k)
	}
}

func (m MapSet[T]) Has(value T) bool {
	_, ok := m.elems[value]
	return ok
}

func (m MapSet[T]) Elements() []T {
	if len(m.elems) == 0 {
		return nil
	}

	elems := make([]T, 0, len(m.elems))
	for v := range m.elems {
		elems = append(elems, v)
	}

	return elems
}

func (m MapSet[T]) Iterate(callback func(value T) bool) {
	for v := range m.elems {
		if !callback(v) {
			return
		}
	}
}

func (m MapSet[T]) Len() int {
	return len(m.elems)
}

func (m MapSet[T]) Diff(o MapSet[T]) MapSet[T] {
	diff := New[T]()
	for e := range m.elems {
		if !o.Has(e) {
			diff.Add(e)
		}
	}
	for e := range o.elems {
		if !m.Has(e) {
			diff.Add(e)
		}
	}
	return diff
}

var _ Set[string] = MapSet[string]{}

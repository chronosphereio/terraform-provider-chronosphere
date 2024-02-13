// Package set contains set implementations.
package set

import (
	"sort"
	stdstrings "strings"
	"sync"
)

// SyncStrings is a set of strings where all operations are protected by a mutex.
// Prefer using the generic Set and New instead of SyncStrings and NewSyncStrings.
type SyncStrings interface {
	Has(n string) bool
	Add(elts ...string)
	Del(elts ...string)
	Clone() SyncStrings
	Minus(other SyncStrings) SyncStrings
	Entries() []string
	SortedEntries() []string
	Empty() bool
	Len() int
	Join(sep string) string
	Equal(other SyncStrings) bool
}

// syncStrings is a goroutine-safe implementation of SyncStrings interface.
type syncStrings struct {
	sync.RWMutex
	data map[string]struct{}
}

// NewSyncStrings returns a new string set pre-populated with elements.
func NewSyncStrings(elts ...string) SyncStrings {
	s := &syncStrings{
		data: map[string]struct{}{},
	}
	s.Add(elts...)
	return s
}

// Has checks whether the set has the given string.
func (s *syncStrings) Has(n string) bool {
	s.RLock()
	defer s.RUnlock()

	_, ok := s.data[n]
	return ok
}

// Add adds new entries to the set.
func (s *syncStrings) Add(elts ...string) {
	s.Lock()
	defer s.Unlock()

	for _, elt := range elts {
		s.data[elt] = struct{}{}
	}
}

// Del removes entries from the set.
func (s *syncStrings) Del(elts ...string) {
	s.Lock()
	defer s.Unlock()

	for _, elt := range elts {
		delete(s.data, elt)
	}
}

// Clone returns a copy of the set.
func (s *syncStrings) Clone() SyncStrings {
	s.RLock()
	defer s.RUnlock()

	clone := make(map[string]struct{}, len(s.data))
	for v := range s.data {
		clone[v] = struct{}{}
	}

	return &syncStrings{
		data: clone,
	}
}

// Minus returns a new set without the contents of the other set.
func (s *syncStrings) Minus(other SyncStrings) SyncStrings {
	res := s.Clone()
	res.Del(other.Entries()...)
	return res
}

// Entries returns the entries in the set in a non-deterministic order.
func (s *syncStrings) Entries() []string {
	s.RLock()
	defer s.RUnlock()

	entries := make([]string, 0, len(s.data))
	for elt := range s.data {
		entries = append(entries, elt)
	}
	return entries
}

// SortedEntries returns the entries in the set
// in sorted order.
func (s *syncStrings) SortedEntries() []string {
	entries := s.Entries()
	sort.Strings(entries)
	return entries
}

// Empty returns true if the set is empty.
func (s *syncStrings) Empty() bool {
	s.RLock()
	defer s.RUnlock()

	return len(s.data) == 0
}

// Len returns the length of the set.
func (s *syncStrings) Len() int {
	s.RLock()
	defer s.RUnlock()

	return len(s.data)
}

// Join concatenates the elements of `s` to create a single string. The separator
// string `sep` is placed between elements in the resulting string.
func (s *syncStrings) Join(sep string) string {
	s.RLock()
	defer s.RUnlock()

	return stdstrings.Join(s.SortedEntries(), sep)
}

// Equal returns true if s contains the same elements as other.
func (s *syncStrings) Equal(other SyncStrings) bool {
	if s.Len() != other.Len() {
		return false
	}

	for _, entry := range s.Entries() {
		if !other.Has(entry) {
			return false
		}
	}

	return true
}

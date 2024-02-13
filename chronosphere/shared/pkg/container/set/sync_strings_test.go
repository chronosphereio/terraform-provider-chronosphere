package set

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrings(t *testing.T) {
	t.Parallel()
	s := NewSyncStrings("foo", "zed", "bar")
	assert.True(t, s.Has("foo"))
	assert.False(t, s.Has("blerp"))
	s.Add("blerp")
	assert.True(t, s.Has("blerp"))
	s.Del("foo", "zed", "bar")
	assert.False(t, s.Has("foo"))
	assert.True(t, s.Has("blerp"))
}

func TestStringsClone(t *testing.T) {
	t.Parallel()
	set := NewSyncStrings("foo", "bar")

	clone := set.Clone()
	assert.Equal(t, set.SortedEntries(), clone.SortedEntries())

	set.Add("baz")
	assert.NotEqual(t, set.SortedEntries(), clone.SortedEntries())
}

func TestStringsMinus(t *testing.T) {
	t.Parallel()

	set := NewSyncStrings("foo", "bar", "baz")

	assert.Equal(t, []string{"foo"}, set.Minus(NewSyncStrings("baz", "bar", "asdf")).SortedEntries())
	assert.Equal(t, []string{}, set.Minus(set).SortedEntries())
	assert.Equal(t, set.SortedEntries(), set.Minus(NewSyncStrings("asdf")).SortedEntries())
	assert.Equal(t, set.SortedEntries(), set.Minus(NewSyncStrings()).SortedEntries())
}

func TestStringsEmpty(t *testing.T) {
	t.Parallel()
	assert.True(t, NewSyncStrings().Empty())
	assert.False(t, NewSyncStrings("foo", "bar").Empty())
}

func TestStringsLen(t *testing.T) {
	t.Parallel()
	assert.Equal(t, 0, NewSyncStrings().Len())
	assert.Equal(t, 2, NewSyncStrings("foo", "bar").Len())
}

func TestStringsSynchronisation(t *testing.T) {
	t.Parallel()

	set := NewSyncStrings()
	wait := sync.WaitGroup{}
	wait.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			set.Add(fmt.Sprintf("item-%d", i))
		}
		wait.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			set.Len()
		}
		wait.Done()
	}()

	wait.Wait()
	assert.Equal(t, 10, set.Len())
}

func TestStringsJoin(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		elts []string
		sep  string
		want string
	}{
		{
			name: "no separator",
			elts: []string{"a", "b", "c"},
			sep:  "",
			want: "abc",
		},
		{
			name: "comma separator",
			elts: []string{"a", "b", "c"},
			sep:  ",",
			want: "a,b,c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			s := NewSyncStrings(tt.elts...)
			actual := s.Join(tt.sep)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func TestStringsEqual(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		left, right SyncStrings
		want        bool
	}{
		{
			name:  "equal",
			left:  NewSyncStrings("a", "b"),
			right: NewSyncStrings("a", "b"),
			want:  true,
		},
		{
			name:  "unequal with same length",
			left:  NewSyncStrings("a", "b"),
			right: NewSyncStrings("a", "c"),
			want:  false,
		},
		{
			name:  "unequal with different lengths",
			left:  NewSyncStrings("a", "b"),
			right: NewSyncStrings("a", "b", "c"),
			want:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := tt.left.Equal(tt.right)
			assert.Equal(t, tt.want, actual)
		})
	}
}

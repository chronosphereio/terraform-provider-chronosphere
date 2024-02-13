package sliceutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStructA struct {
	A string
}

type testStructB struct {
	B string
}

func TestMap(t *testing.T) {
	tests := []struct {
		name string
		in   []testStructA
		out  []testStructB
	}{
		{
			name: "nil",
			in:   nil,
			out:  nil,
		},
		{
			name: "convert",
			in:   []testStructA{{A: "a"}, {A: "b"}, {A: "c"}},
			out:  []testStructB{{B: "a"}, {B: "b"}, {B: "c"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := Map(test.in, func(a testStructA) testStructB {
				return testStructB{B: a.A}
			})
			assert.Equal(t, test.out, b)
		})
	}
}

func TestMapErr(t *testing.T) {
	tests := []struct {
		name     string
		in       []testStructA
		out      []testStructB
		hasError bool
	}{
		{
			name:     "nil",
			in:       nil,
			out:      nil,
			hasError: false,
		},
		{
			name:     "convert",
			in:       []testStructA{{A: "a"}, {A: "b"}, {A: "c"}},
			out:      []testStructB{{B: "a"}, {B: "b"}, {B: "c"}},
			hasError: false,
		},
		{
			name:     "error",
			in:       []testStructA{{A: "a"}, {A: "b"}, {A: "c"}},
			out:      []testStructB{{B: "a"}, {B: "b"}, {B: "c"}},
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b, err := MapErr(test.in, func(a testStructA) (testStructB, error) {
				if test.hasError {
					return testStructB{}, assert.AnError
				}
				return testStructB{B: a.A}, nil
			})
			if test.hasError {
				assert.Error(t, assert.AnError)
			} else {
				assert.Equal(t, test.out, b)
				assert.NoError(t, err)
			}
		})
	}
}

package mockmatcher

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type anyEq struct {
	t       testing.TB
	want    interface{}
	message string
}

// AnyEq returns a matcher that matches any object, but then ensures that the
// object is equal to a passed-in value using assert.Equal.
// This is different from `gomock.Eq`, as this matcher will always match, but then use
// `assert.Equal` to print a diff of the difference, while `Eq` will be considered a
// missing call, and not do a diff of the mismatched arguments.
func AnyEq(t testing.TB, want interface{}, message string) gomock.Matcher {
	return &anyEq{t, want, message}
}

func (m anyEq) Matches(x interface{}) bool {
	assert.Equal(m.t, m.want, x, m.message)
	return true
}

func (m anyEq) String() string {
	return fmt.Sprintf("AnyEq(%v) %v", m.want, m.message)
}

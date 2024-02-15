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

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

package localid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLocalID(t *testing.T) {
	assert.True(t, IsLocalID(NewLocalID()))
	assert.True(t, IsLocalID(NewImportedID()))
	assert.True(t, IsLocalID("1234567")) // Older format, no longer used but possibly present in state files

	assert.False(t, IsLocalID(""))
	assert.False(t, IsLocalID("foo"))
	assert.False(t, IsLocalID("foo-123"))
	assert.False(t, IsLocalID("foo-bar"))
}

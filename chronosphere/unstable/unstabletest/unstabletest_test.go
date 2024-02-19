// Copyright 2024 Chronosphere Inc.
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

package unstabletest

import (
	"os"
	"testing"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/unstable"
	"github.com/stretchr/testify/require"
)

func TestWithEnabled(t *testing.T) {
	require.False(t, unstable.Enabled())

	WithEnabled(t)

	require.True(t, unstable.Enabled())
	require.Equal(t, os.Getenv(unstable.Env), "1")
}

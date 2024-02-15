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

package tfschema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFloat64RangeValidator(t *testing.T) {
	tests := []struct {
		f       float64
		hi      float64
		lo      float64
		wantErr string
	}{
		{
			f:  1,
			hi: 1.0,
			lo: 0,
		},
		{
			f:  0.5,
			hi: 1.0,
			lo: 0,
		},
		{
			f:  0,
			hi: 1.0,
			lo: 0,
		},
		{
			f:       2.0,
			hi:      1.0,
			lo:      0,
			wantErr: "value must be within range [0.000000, 1.000000]",
		},
		{
			f:       -1,
			hi:      1.0,
			lo:      0,
			wantErr: "value must be within range [0.000000, 1.000000]",
		},
	}

	for _, tt := range tests {
		validator := float64RangeValidator(tt.lo, tt.hi)
		err := DiagError(validator(tt.f, nil))
		if tt.wantErr != "" {
			require.Error(t, err) // ensure error severity is set
			assert.Contains(t, err.Error(), tt.wantErr)
		} else {
			assert.Nil(t, err)
		}
	}
}

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

package tfschema

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseDuration(t *testing.T) {
	for _, test := range []struct {
		time     string
		expected time.Duration
		wantErr  bool
	}{
		{
			time:     "30s",
			expected: 30 * time.Second,
		},
		{
			time:     "3d",
			expected: 72 * time.Hour,
		},
		{
			time:     "1m0s",
			expected: 1 * time.Minute,
		},
		{
			time:    "1m0",
			wantErr: true,
		},
	} {
		t.Run(test.time, func(t *testing.T) {
			duration, err := ParseDuration(test.time)
			if test.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.expected, duration)
		})
	}
}

func TestDurationValidate(t *testing.T) {
	tests := []struct {
		name       string
		input      any
		wantErrors []string
	}{
		{
			name:  "valid duration",
			input: "10s",
		},
		{
			name:  "duration with days",
			input: "3d", // stdlib doesn't support days.
		},
		{
			name:  "invalid type",
			input: 10,
			wantErrors: []string{
				"expected type to be string",
			},
		},
		{
			name:  "invalid duration",
			input: "1minute",
			wantErrors: []string{
				`"1minute" is not a valid duration: time: unknown unit "minute" in duration "1minute"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagErrs := Duration{}.validate(tt.input, nil)
			require.Len(t, diagErrs, len(tt.wantErrors))
			for i := range diagErrs {
				assert.Equal(t, tt.wantErrors[i], diagErrs[i].Summary)
				assert.Equal(t, diag.Error, diagErrs[i].Severity)
			}
		})
	}
}

func TestDurationNormalize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "no change",
			input: "30s",
			want:  "30s",
		},
		{
			name:  "normalize supports prom units",
			input: "25h",
			want:  "1d1h",
		},
		{
			name:  "different format",
			input: "90s",
			want:  "1m30s",
		},
		{
			name:  "empty is same as 0s",
			input: "",
			want:  "0s",
		},
		{
			name:  "0 of any unit is 0s",
			input: "0d",
			want:  "0s",
		},
		{
			name:  "invalid",
			input: "invalid",
			want:  "invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, Duration{}.Normalize(tt.input))
		})
	}
}

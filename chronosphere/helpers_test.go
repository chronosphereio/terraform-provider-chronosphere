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

package chronosphere

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDurationToSecs(t *testing.T) {
	tests := []struct {
		name    string
		give    string
		want    int32
		wantErr string
	}{
		{
			name: "ok",
			give: "5s",
			want: 5,
		},
		{
			name:    "err if millis",
			give:    "5s500ms",
			wantErr: `invalid duration "5s500ms": must use seconds granularity`,
		},
		{
			name:    "err if under 1s",
			give:    "50ms",
			wantErr: `invalid duration "50ms": must use seconds granularity`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := durationToSecs(tt.give)
			if tt.wantErr != "" {
				require.EqualError(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

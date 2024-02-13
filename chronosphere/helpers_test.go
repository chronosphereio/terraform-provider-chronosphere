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

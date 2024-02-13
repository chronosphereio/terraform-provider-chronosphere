package prettyenum

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

func TestSpanFilterMatchType(t *testing.T) {
	testCases := []struct {
		raw    string
		expErr error
		model  models.SpanFilterSpanFilterMatchType
	}{
		{
			raw:   "include",
			model: models.SpanFilterSpanFilterMatchTypeINCLUDE,
		},
		{
			raw:   "INCLUDE",
			model: models.SpanFilterSpanFilterMatchTypeINCLUDE,
		},
		{
			raw:   "exclude",
			model: models.SpanFilterSpanFilterMatchTypeEXCLUDE,
		},
		{
			raw:   "EXCLUDE",
			model: models.SpanFilterSpanFilterMatchTypeEXCLUDE,
		},
		{
			raw:    "bad",
			expErr: errors.New("invalid match_type: bad"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.raw, func(t *testing.T) {
			at, err := NewSpanFilterMatchType(testCase.raw)
			if testCase.expErr != nil {
				require.Equal(t, testCase.expErr, err)
				return
			}
			require.Equal(t, testCase.model, at.Model())
		})
	}
}

// TestSpanFilterMatchTypeFromModel ensure that all models.TraceSearchFilterSpanFilterMatchType are mapped in spanFilterMatchTypeFromModel
func TestSpanFilterMatchTypeFromModel(t *testing.T) {
	for _, orig := range modelFromSpanFilterMatchType {
		s := SpanFilterMatchTypeFromModel(orig)
		got := s.Model()
		require.Equal(t, orig, got)
	}
}

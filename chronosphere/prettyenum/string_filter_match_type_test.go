package prettyenum

import (
	"errors"
	"testing"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/stretchr/testify/require"
)

func TestStringFilterMatchType(t *testing.T) {
	testCases := []struct {
		raw    string
		expErr error
		model  models.StringFilterStringFilterMatchType
	}{
		{
			raw:   "exact",
			model: models.StringFilterStringFilterMatchTypeEXACT,
		},
		{
			raw:   "EXACT",
			model: models.StringFilterStringFilterMatchTypeEXACT,
		},
		{
			raw:   "regex",
			model: models.StringFilterStringFilterMatchTypeREGEX,
		},
		{
			raw:   "REGEX",
			model: models.StringFilterStringFilterMatchTypeREGEX,
		},
		{
			raw:   "exact_negation",
			model: models.StringFilterStringFilterMatchTypeEXACTNEGATION,
		},
		{
			raw:   "EXACT_NEGATION",
			model: models.StringFilterStringFilterMatchTypeEXACTNEGATION,
		},
		{
			raw:   "regex_negation",
			model: models.StringFilterStringFilterMatchTypeREGEXNEGATION,
		},
		{
			raw:   "REGEX_NEGATION",
			model: models.StringFilterStringFilterMatchTypeREGEXNEGATION,
		},
		{
			raw:    "bad",
			expErr: errors.New("invalid match: bad"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.raw, func(t *testing.T) {
			at, err := NewStringFilterMatchType(testCase.raw)
			if testCase.expErr != nil {
				require.Equal(t, testCase.expErr, err)
				return
			}
			require.Equal(t, testCase.model, at.Model())
		})
	}
}

// TestStringFilterMatchTypeFromModel ensure that all models.TraceSearchFilterStringFilterMatchType are mapped in stringFilterMatchTypeFromModel
func TestStringFilterMatchTypeFromModel(t *testing.T) {
	for _, orig := range modelFromStringFilterMatchType {
		s := StringFilterMatchTypeFromModel(orig)
		got := s.Model()
		require.Equal(t, orig, got)
	}
}

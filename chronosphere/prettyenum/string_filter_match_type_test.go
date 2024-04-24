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
			raw:   "in",
			model: models.StringFilterStringFilterMatchTypeIN,
		},
		{
			raw:   "IN",
			model: models.StringFilterStringFilterMatchTypeIN,
		},
		{
			raw:   "not_in",
			model: models.StringFilterStringFilterMatchTypeNOTIN,
		},
		{
			raw:   "NOT_IN",
			model: models.StringFilterStringFilterMatchTypeNOTIN,
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

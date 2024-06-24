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

	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

func TestDatasetType(t *testing.T) {
	testCases := []struct {
		raw    string
		expErr error
		model  models.DatasetDatasetType
	}{
		{
			raw:   "traces",
			model: models.DatasetDatasetTypeTRACES,
		},
		{
			raw:   "TRACES",
			model: models.DatasetDatasetTypeTRACES,
		},
		{
			raw:   "logs",
			model: models.DatasetDatasetTypeLOGS,
		},
		{
			raw:   "LOGS",
			model: models.DatasetDatasetTypeLOGS,
		},
		{
			raw:    "bad",
			expErr: errors.New("invalid dataset type: bad"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.raw, func(t *testing.T) {
			at, err := NewDatasetDatasetType(testCase.raw)
			if testCase.expErr != nil {
				require.Equal(t, testCase.expErr, err)
				return
			}
			require.Equal(t, testCase.model, at.Model())
		})
	}
}

func TestDatasetTypeFromModel(t *testing.T) {
	testCases := []struct {
		expErr   error
		model    models.DatasetDatasetType
		expected DatasetType
	}{
		{
			model:    models.DatasetDatasetTypeTRACES,
			expected: DatasetDatasetTypeTracesModel,
		},
		{
			model:    models.DatasetDatasetTypeLOGS,
			expected: DatasetDatasetTypeLogsModel,
		},
	}

	for _, testCase := range testCases {
		t.Run(string(testCase.model), func(t *testing.T) {
			require.Equal(t, testCase.expected, DatasetTypeFromModel(testCase.model))
		})
	}
}

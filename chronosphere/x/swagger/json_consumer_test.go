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

package swagger

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

func TestJSONConsumerErrors(t *testing.T) {
	tests := []struct {
		name         string
		reader       io.Reader
		expectedData *models.APIError
		output       string
	}{
		{
			name:   "vanilla runtime error",
			reader: io.NopCloser(strings.NewReader(`{"code":14, "message": "lizard"}`)),
			expectedData: &models.APIError{
				Code:    14,
				Message: "lizard",
			},
		},
		{
			name:   "requestID'd runtime error",
			reader: &RequestIDBody{ReadCloser: io.NopCloser(strings.NewReader(`{"code":14, "message": "lizard"}`)), requestID: "abc-123"},
			expectedData: &models.APIError{
				Code:    14,
				Message: "lizard (request_id=abc-123)",
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			consumer := JSONConsumer()
			data := &models.APIError{}
			require.NoError(t, consumer.Consume(tt.reader, data))
			assert.Equal(t, tt.expectedData, data)
		})
	}
}

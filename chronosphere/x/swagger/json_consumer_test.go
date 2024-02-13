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

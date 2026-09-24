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

package clienterror

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	configv1models "github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// fakeAPIError mimics a go-swagger generated response error: it carries an HTTP
// status code (via Code()) and a structured payload with an optional
// application-level code (via GetPayload()), matching the interfaces the real
// generated clients implement.
type fakeAPIError struct {
	httpCode int
	payload  *configv1models.APIError
}

func (e *fakeAPIError) Code() int                            { return e.httpCode }
func (e *fakeAPIError) GetPayload() *configv1models.APIError { return e.payload }
func (e *fakeAPIError) Error() string                        { return e.payload.Message }

func TestIsClientError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil",
			err:  nil,
			want: false,
		},
		{
			// Regression: "monitor configuration is too large" is a generic
			// HTTP 400 BAD_REQUEST without the entity-validation application
			// code (400008). This used to be silently ignored during dry-run,
			// letting an invalid plan pass.
			name: "generic 400 monitor too large",
			err: &fakeAPIError{
				httpCode: 400,
				payload: &configv1models.APIError{
					Message: "CreateMonitor failed: monitor configuration is too large",
				},
			},
			want: true,
		},
		{
			name: "entity validation failed (400008)",
			err: &fakeAPIError{
				httpCode: 400,
				payload: &configv1models.APIError{
					Code:    entityValidationFailedCode,
					Message: "entity validation failed",
				},
			},
			want: true,
		},
		{
			name: "not found (404)",
			err: &fakeAPIError{
				httpCode: 404,
				payload:  &configv1models.APIError{Message: "not found"},
			},
			want: true,
		},
		{
			// Transient server errors must not fail a plan.
			name: "internal server error (500)",
			err: &fakeAPIError{
				httpCode: 500,
				payload:  &configv1models.APIError{Message: "boom"},
			},
			want: false,
		},
		{
			// Non-HTTP errors (e.g. network/transport failures) have no status
			// code and are not treated as client errors.
			name: "plain error",
			err:  errors.New("dial tcp: connection refused"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsClientError(tt.err))
		})
	}
}

// TestDryRunErrorClassification documents the root cause of the dry-run bug:
// IsEntityValidationFailed only matches application code 400008, so a generic
// 400 BAD_REQUEST like "monitor configuration is too large" was ignored. The
// broader IsClientError check catches it while still matching 400008.
func TestDryRunErrorClassification(t *testing.T) {
	tooLarge := &fakeAPIError{
		httpCode: 400,
		payload: &configv1models.APIError{
			Message: "CreateMonitor failed: monitor configuration is too large",
		},
	}

	// Old behavior: not caught -> plan incorrectly passed.
	assert.False(t, IsEntityValidationFailed(tooLarge))
	// New behavior: caught -> plan fails as expected.
	assert.True(t, IsClientError(tooLarge))

	entityValidation := &fakeAPIError{
		httpCode: 400,
		payload: &configv1models.APIError{
			Code:    entityValidationFailedCode,
			Message: "entity validation failed",
		},
	}
	// The previously-handled case is still handled.
	assert.True(t, IsEntityValidationFailed(entityValidation))
	assert.True(t, IsClientError(entityValidation))
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BaselineBehaviorOptionsFastSampleOptions baseline behavior options fast sample options
//
// swagger:model BaselineBehaviorOptionsFastSampleOptions
type BaselineBehaviorOptionsFastSampleOptions struct {

	// Duration in seconds under which traces are sampled
	// according to the given sample rate.
	MaxDurationSeconds float64 `json:"max_duration_seconds,omitempty"`

	// Sample rate for traces under the given duration.
	SampleRate float64 `json:"sample_rate,omitempty"`
}

// Validate validates this baseline behavior options fast sample options
func (m *BaselineBehaviorOptionsFastSampleOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this baseline behavior options fast sample options based on context it is used
func (m *BaselineBehaviorOptionsFastSampleOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaselineBehaviorOptionsFastSampleOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaselineBehaviorOptionsFastSampleOptions) UnmarshalBinary(b []byte) error {
	var res BaselineBehaviorOptionsFastSampleOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

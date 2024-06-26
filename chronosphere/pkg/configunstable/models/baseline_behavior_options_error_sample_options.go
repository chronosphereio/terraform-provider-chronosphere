// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BaselineBehaviorOptionsErrorSampleOptions baseline behavior options error sample options
//
// swagger:model BaselineBehaviorOptionsErrorSampleOptions
type BaselineBehaviorOptionsErrorSampleOptions struct {

	// Sample rate for traces with errors.
	SampleRate float64 `json:"sample_rate,omitempty"`
}

// Validate validates this baseline behavior options error sample options
func (m *BaselineBehaviorOptionsErrorSampleOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this baseline behavior options error sample options based on context it is used
func (m *BaselineBehaviorOptionsErrorSampleOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaselineBehaviorOptionsErrorSampleOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaselineBehaviorOptionsErrorSampleOptions) UnmarshalBinary(b []byte) error {
	var res BaselineBehaviorOptionsErrorSampleOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

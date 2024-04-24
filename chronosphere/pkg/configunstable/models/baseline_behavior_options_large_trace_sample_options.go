// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BaselineBehaviorOptionsLargeTraceSampleOptions baseline behavior options large trace sample options
//
// swagger:model BaselineBehaviorOptionsLargeTraceSampleOptions
type BaselineBehaviorOptionsLargeTraceSampleOptions struct {

	// sample rate
	SampleRate float64 `json:"sample_rate,omitempty"`

	// For N = number of spans in the trace, if N >= span_count_threshold, the trace is sampled according to the
	// given sample rate.
	SpanCountThreshold int64 `json:"span_count_threshold,omitempty"`
}

// Validate validates this baseline behavior options large trace sample options
func (m *BaselineBehaviorOptionsLargeTraceSampleOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this baseline behavior options large trace sample options based on context it is used
func (m *BaselineBehaviorOptionsLargeTraceSampleOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaselineBehaviorOptionsLargeTraceSampleOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaselineBehaviorOptionsLargeTraceSampleOptions) UnmarshalBinary(b []byte) error {
	var res BaselineBehaviorOptionsLargeTraceSampleOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

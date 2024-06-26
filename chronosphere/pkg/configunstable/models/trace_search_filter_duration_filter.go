// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TraceSearchFilterDurationFilter trace search filter duration filter
//
// swagger:model TraceSearchFilterDurationFilter
type TraceSearchFilterDurationFilter struct {

	// Maximum duration, in seconds, required for a span or trace to match.
	MaxSecs float64 `json:"max_secs,omitempty"`

	// Minimum duration, in seconds, required for a span or trace to match.
	MinSecs float64 `json:"min_secs,omitempty"`
}

// Validate validates this trace search filter duration filter
func (m *TraceSearchFilterDurationFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this trace search filter duration filter based on context it is used
func (m *TraceSearchFilterDurationFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TraceSearchFilterDurationFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraceSearchFilterDurationFilter) UnmarshalBinary(b []byte) error {
	var res TraceSearchFilterDurationFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SavedTraceSearchRelativeTimeFilter saved trace search relative time filter
//
// swagger:model SavedTraceSearchRelativeTimeFilter
type SavedTraceSearchRelativeTimeFilter struct {

	// The duration, in seconds, from now to the end of the search interval.
	EndRelativeOffsetSecs int32 `json:"end_relative_offset_secs,omitempty"`

	// The duration, in seconds, from now to the beginning of the search interval.
	StartRelativeOffsetSecs int32 `json:"start_relative_offset_secs,omitempty"`
}

// Validate validates this saved trace search relative time filter
func (m *SavedTraceSearchRelativeTimeFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this saved trace search relative time filter based on context it is used
func (m *SavedTraceSearchRelativeTimeFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SavedTraceSearchRelativeTimeFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SavedTraceSearchRelativeTimeFilter) UnmarshalBinary(b []byte) error {
	var res SavedTraceSearchRelativeTimeFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

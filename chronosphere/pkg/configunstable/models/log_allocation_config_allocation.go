// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogAllocationConfigAllocation Configuration for allocating resources to a dataset.
//
// swagger:model LogAllocationConfigAllocation
type LogAllocationConfigAllocation struct {

	// percent of license
	PercentOfLicense float64 `json:"percent_of_license,omitempty"`
}

// Validate validates this log allocation config allocation
func (m *LogAllocationConfigAllocation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this log allocation config allocation based on context it is used
func (m *LogAllocationConfigAllocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogAllocationConfigAllocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogAllocationConfigAllocation) UnmarshalBinary(b []byte) error {
	var res LogAllocationConfigAllocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
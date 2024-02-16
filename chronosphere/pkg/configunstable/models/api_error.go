// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIError api error
//
// swagger:model apiError
type APIError struct {

	// An optional private error code whose values are undefined.
	Code int32 `json:"code,omitempty"`

	// An error message describing what went wrong.
	Message string `json:"message,omitempty"`
}

// Validate validates this api error
func (m *APIError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api error based on context it is used
func (m *APIError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIError) UnmarshalBinary(b []byte) error {
	var res APIError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
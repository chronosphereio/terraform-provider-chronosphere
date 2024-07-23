// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Configv1LogSearchFilter configv1 log search filter
//
// swagger:model configv1LogSearchFilter
type Configv1LogSearchFilter struct {

	// Matches logs which match this query.
	// Query may only include top level operations (no nested clauses).
	// Only one type of operator AND/OR is allowed.
	Query string `json:"query,omitempty"`
}

// Validate validates this configv1 log search filter
func (m *Configv1LogSearchFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this configv1 log search filter based on context it is used
func (m *Configv1LogSearchFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Configv1LogSearchFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1LogSearchFilter) UnmarshalBinary(b []byte) error {
	var res Configv1LogSearchFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
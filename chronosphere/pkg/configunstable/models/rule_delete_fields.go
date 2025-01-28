// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RuleDeleteFields DeleteFields is the configuration for a delete fields rule.
//
// swagger:model RuleDeleteFields
type RuleDeleteFields struct {

	// field regex
	FieldRegex string `json:"field_regex,omitempty"`
}

// Validate validates this rule delete fields
func (m *RuleDeleteFields) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rule delete fields based on context it is used
func (m *RuleDeleteFields) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RuleDeleteFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleDeleteFields) UnmarshalBinary(b []byte) error {
	var res RuleDeleteFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

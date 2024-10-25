// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RuleDrop Drop is the configuration for a drop filter
//
// swagger:model RuleDrop
type RuleDrop struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// label
	Label *ConfigunstableLabel `json:"label,omitempty"`
}

// Validate validates this rule drop
func (m *RuleDrop) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuleDrop) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if m.Label != nil {
		if err := m.Label.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("label")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("label")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rule drop based on the context it is used
func (m *RuleDrop) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuleDrop) contextValidateLabel(ctx context.Context, formats strfmt.Registry) error {

	if m.Label != nil {
		if err := m.Label.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("label")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("label")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuleDrop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleDrop) UnmarshalBinary(b []byte) error {
	var res RuleDrop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

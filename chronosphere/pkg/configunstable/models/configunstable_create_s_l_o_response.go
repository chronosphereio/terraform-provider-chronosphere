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

// ConfigunstableCreateSLOResponse configunstable create s l o response
//
// swagger:model configunstableCreateSLOResponse
type ConfigunstableCreateSLOResponse struct {

	// slo
	Slo *ConfigunstableSLO `json:"slo,omitempty"`
}

// Validate validates this configunstable create s l o response
func (m *ConfigunstableCreateSLOResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSlo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateSLOResponse) validateSlo(formats strfmt.Registry) error {
	if swag.IsZero(m.Slo) { // not required
		return nil
	}

	if m.Slo != nil {
		if err := m.Slo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable create s l o response based on the context it is used
func (m *ConfigunstableCreateSLOResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSlo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateSLOResponse) contextValidateSlo(ctx context.Context, formats strfmt.Registry) error {

	if m.Slo != nil {
		if err := m.Slo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableCreateSLOResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableCreateSLOResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableCreateSLOResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

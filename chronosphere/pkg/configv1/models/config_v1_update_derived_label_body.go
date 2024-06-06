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

// ConfigV1UpdateDerivedLabelBody config v1 update derived label body
//
// swagger:model ConfigV1UpdateDerivedLabelBody
type ConfigV1UpdateDerivedLabelBody struct {

	// If true, the DerivedLabel will be created if it does not already exist, identified by slug. If false, an error will be returned if the DerivedLabel does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// derived label
	DerivedLabel *Configv1DerivedLabel `json:"derived_label,omitempty"`

	// If true, the DerivedLabel will not be created nor updated, and no response DerivedLabel will be returned. The response will return an error if the given DerivedLabel is invalid.
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this config v1 update derived label body
func (m *ConfigV1UpdateDerivedLabelBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDerivedLabel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateDerivedLabelBody) validateDerivedLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.DerivedLabel) { // not required
		return nil
	}

	if m.DerivedLabel != nil {
		if err := m.DerivedLabel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("derived_label")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("derived_label")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config v1 update derived label body based on the context it is used
func (m *ConfigV1UpdateDerivedLabelBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDerivedLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateDerivedLabelBody) contextValidateDerivedLabel(ctx context.Context, formats strfmt.Registry) error {

	if m.DerivedLabel != nil {
		if err := m.DerivedLabel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("derived_label")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("derived_label")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigV1UpdateDerivedLabelBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigV1UpdateDerivedLabelBody) UnmarshalBinary(b []byte) error {
	var res ConfigV1UpdateDerivedLabelBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

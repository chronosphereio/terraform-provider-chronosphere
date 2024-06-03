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

// ConfigV1UpdateNotifierBody config v1 update notifier body
//
// swagger:model ConfigV1UpdateNotifierBody
type ConfigV1UpdateNotifierBody struct {

	// If true, the Notifier will be created if it does not already exist, identified by slug. If false, an error will be returned if the Notifier does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the Notifier will not be created nor updated, and no response Notifier will be returned. The response will return an error if the given Notifier is invalid.
	DryRun bool `json:"dry_run,omitempty"`

	// notifier
	Notifier *Configv1Notifier `json:"notifier,omitempty"`
}

// Validate validates this config v1 update notifier body
func (m *ConfigV1UpdateNotifierBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotifier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateNotifierBody) validateNotifier(formats strfmt.Registry) error {
	if swag.IsZero(m.Notifier) { // not required
		return nil
	}

	if m.Notifier != nil {
		if err := m.Notifier.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifier")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config v1 update notifier body based on the context it is used
func (m *ConfigV1UpdateNotifierBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNotifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateNotifierBody) contextValidateNotifier(ctx context.Context, formats strfmt.Registry) error {

	if m.Notifier != nil {
		if err := m.Notifier.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifier")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifier")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigV1UpdateNotifierBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigV1UpdateNotifierBody) UnmarshalBinary(b []byte) error {
	var res ConfigV1UpdateNotifierBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

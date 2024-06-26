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

// ConfigunstableCreateNoopEntityRequest configunstable create noop entity request
//
// swagger:model configunstableCreateNoopEntityRequest
type ConfigunstableCreateNoopEntityRequest struct {

	// noop entity
	NoopEntity *ConfigunstableNoopEntity `json:"noop_entity,omitempty"`
}

// Validate validates this configunstable create noop entity request
func (m *ConfigunstableCreateNoopEntityRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNoopEntity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateNoopEntityRequest) validateNoopEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.NoopEntity) { // not required
		return nil
	}

	if m.NoopEntity != nil {
		if err := m.NoopEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("noop_entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("noop_entity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable create noop entity request based on the context it is used
func (m *ConfigunstableCreateNoopEntityRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNoopEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateNoopEntityRequest) contextValidateNoopEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.NoopEntity != nil {
		if err := m.NoopEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("noop_entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("noop_entity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableCreateNoopEntityRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableCreateNoopEntityRequest) UnmarshalBinary(b []byte) error {
	var res ConfigunstableCreateNoopEntityRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

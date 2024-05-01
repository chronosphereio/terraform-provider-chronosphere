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

// ConfigunstableCreateLogAllocationConfigResponse configunstable create log allocation config response
//
// swagger:model configunstableCreateLogAllocationConfigResponse
type ConfigunstableCreateLogAllocationConfigResponse struct {

	// log allocation config
	LogAllocationConfig *ConfigunstableLogAllocationConfig `json:"log_allocation_config,omitempty"`
}

// Validate validates this configunstable create log allocation config response
func (m *ConfigunstableCreateLogAllocationConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogAllocationConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateLogAllocationConfigResponse) validateLogAllocationConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.LogAllocationConfig) { // not required
		return nil
	}

	if m.LogAllocationConfig != nil {
		if err := m.LogAllocationConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_allocation_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_allocation_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable create log allocation config response based on the context it is used
func (m *ConfigunstableCreateLogAllocationConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogAllocationConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateLogAllocationConfigResponse) contextValidateLogAllocationConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.LogAllocationConfig != nil {
		if err := m.LogAllocationConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_allocation_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_allocation_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableCreateLogAllocationConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableCreateLogAllocationConfigResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableCreateLogAllocationConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

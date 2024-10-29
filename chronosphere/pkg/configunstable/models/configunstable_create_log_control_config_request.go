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

// ConfigunstableCreateLogControlConfigRequest configunstable create log control config request
//
// swagger:model configunstableCreateLogControlConfigRequest
type ConfigunstableCreateLogControlConfigRequest struct {

	// log control config
	LogControlConfig *ConfigunstableLogControlConfig `json:"log_control_config,omitempty"`
}

// Validate validates this configunstable create log control config request
func (m *ConfigunstableCreateLogControlConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogControlConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateLogControlConfigRequest) validateLogControlConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.LogControlConfig) { // not required
		return nil
	}

	if m.LogControlConfig != nil {
		if err := m.LogControlConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_control_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_control_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable create log control config request based on the context it is used
func (m *ConfigunstableCreateLogControlConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogControlConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateLogControlConfigRequest) contextValidateLogControlConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.LogControlConfig != nil {
		if err := m.LogControlConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_control_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_control_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableCreateLogControlConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableCreateLogControlConfigRequest) UnmarshalBinary(b []byte) error {
	var res ConfigunstableCreateLogControlConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

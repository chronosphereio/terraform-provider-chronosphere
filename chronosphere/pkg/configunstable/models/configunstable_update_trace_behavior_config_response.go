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

// ConfigunstableUpdateTraceBehaviorConfigResponse configunstable update trace behavior config response
//
// swagger:model configunstableUpdateTraceBehaviorConfigResponse
type ConfigunstableUpdateTraceBehaviorConfigResponse struct {

	// trace behavior config
	TraceBehaviorConfig *ConfigunstableTraceBehaviorConfig `json:"trace_behavior_config,omitempty"`
}

// Validate validates this configunstable update trace behavior config response
func (m *ConfigunstableUpdateTraceBehaviorConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraceBehaviorConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableUpdateTraceBehaviorConfigResponse) validateTraceBehaviorConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceBehaviorConfig) { // not required
		return nil
	}

	if m.TraceBehaviorConfig != nil {
		if err := m.TraceBehaviorConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_behavior_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_behavior_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable update trace behavior config response based on the context it is used
func (m *ConfigunstableUpdateTraceBehaviorConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTraceBehaviorConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableUpdateTraceBehaviorConfigResponse) contextValidateTraceBehaviorConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.TraceBehaviorConfig != nil {
		if err := m.TraceBehaviorConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_behavior_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_behavior_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableUpdateTraceBehaviorConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableUpdateTraceBehaviorConfigResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableUpdateTraceBehaviorConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

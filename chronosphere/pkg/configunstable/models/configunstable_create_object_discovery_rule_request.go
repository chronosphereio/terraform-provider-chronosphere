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

// ConfigunstableCreateObjectDiscoveryRuleRequest configunstable create object discovery rule request
//
// swagger:model configunstableCreateObjectDiscoveryRuleRequest
type ConfigunstableCreateObjectDiscoveryRuleRequest struct {

	// object discovery rule
	ObjectDiscoveryRule *ConfigunstableObjectDiscoveryRule `json:"object_discovery_rule,omitempty"`
}

// Validate validates this configunstable create object discovery rule request
func (m *ConfigunstableCreateObjectDiscoveryRuleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectDiscoveryRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateObjectDiscoveryRuleRequest) validateObjectDiscoveryRule(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectDiscoveryRule) { // not required
		return nil
	}

	if m.ObjectDiscoveryRule != nil {
		if err := m.ObjectDiscoveryRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_discovery_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object_discovery_rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable create object discovery rule request based on the context it is used
func (m *ConfigunstableCreateObjectDiscoveryRuleRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjectDiscoveryRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateObjectDiscoveryRuleRequest) contextValidateObjectDiscoveryRule(ctx context.Context, formats strfmt.Registry) error {

	if m.ObjectDiscoveryRule != nil {
		if err := m.ObjectDiscoveryRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object_discovery_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object_discovery_rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableCreateObjectDiscoveryRuleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableCreateObjectDiscoveryRuleRequest) UnmarshalBinary(b []byte) error {
	var res ConfigunstableCreateObjectDiscoveryRuleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// Configv1CreateMappingRuleRequest configv1 create mapping rule request
//
// swagger:model configv1CreateMappingRuleRequest
type Configv1CreateMappingRuleRequest struct {

	// If true, the MappingRule will not be created, and no response MappingRule will be returned. The response will return an error if the given MappingRule is invalid.
	DryRun bool `json:"dry_run,omitempty"`

	// mapping rule
	MappingRule *Configv1MappingRule `json:"mapping_rule,omitempty"`
}

// Validate validates this configv1 create mapping rule request
func (m *Configv1CreateMappingRuleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMappingRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateMappingRuleRequest) validateMappingRule(formats strfmt.Registry) error {
	if swag.IsZero(m.MappingRule) { // not required
		return nil
	}

	if m.MappingRule != nil {
		if err := m.MappingRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mapping_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mapping_rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 create mapping rule request based on the context it is used
func (m *Configv1CreateMappingRuleRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMappingRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateMappingRuleRequest) contextValidateMappingRule(ctx context.Context, formats strfmt.Registry) error {

	if m.MappingRule != nil {
		if err := m.MappingRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mapping_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mapping_rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1CreateMappingRuleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1CreateMappingRuleRequest) UnmarshalBinary(b []byte) error {
	var res Configv1CreateMappingRuleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

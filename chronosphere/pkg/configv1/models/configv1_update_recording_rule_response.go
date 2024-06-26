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

// Configv1UpdateRecordingRuleResponse configv1 update recording rule response
//
// swagger:model configv1UpdateRecordingRuleResponse
type Configv1UpdateRecordingRuleResponse struct {

	// recording rule
	RecordingRule *Configv1RecordingRule `json:"recording_rule,omitempty"`
}

// Validate validates this configv1 update recording rule response
func (m *Configv1UpdateRecordingRuleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecordingRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1UpdateRecordingRuleResponse) validateRecordingRule(formats strfmt.Registry) error {
	if swag.IsZero(m.RecordingRule) { // not required
		return nil
	}

	if m.RecordingRule != nil {
		if err := m.RecordingRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recording_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recording_rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 update recording rule response based on the context it is used
func (m *Configv1UpdateRecordingRuleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecordingRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1UpdateRecordingRuleResponse) contextValidateRecordingRule(ctx context.Context, formats strfmt.Registry) error {

	if m.RecordingRule != nil {
		if err := m.RecordingRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recording_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recording_rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1UpdateRecordingRuleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1UpdateRecordingRuleResponse) UnmarshalBinary(b []byte) error {
	var res Configv1UpdateRecordingRuleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

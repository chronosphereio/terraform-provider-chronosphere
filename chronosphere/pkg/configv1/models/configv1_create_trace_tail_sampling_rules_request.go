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

// Configv1CreateTraceTailSamplingRulesRequest configv1 create trace tail sampling rules request
//
// swagger:model configv1CreateTraceTailSamplingRulesRequest
type Configv1CreateTraceTailSamplingRulesRequest struct {

	// If true, the TraceTailSamplingRules isn't created, and no response TraceTailSamplingRules will be returned. The response will return an error if the given TraceTailSamplingRules is invalid.
	DryRun bool `json:"dry_run,omitempty"`

	// trace tail sampling rules
	TraceTailSamplingRules *Configv1TraceTailSamplingRules `json:"trace_tail_sampling_rules,omitempty"`
}

// Validate validates this configv1 create trace tail sampling rules request
func (m *Configv1CreateTraceTailSamplingRulesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraceTailSamplingRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateTraceTailSamplingRulesRequest) validateTraceTailSamplingRules(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceTailSamplingRules) { // not required
		return nil
	}

	if m.TraceTailSamplingRules != nil {
		if err := m.TraceTailSamplingRules.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_tail_sampling_rules")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_tail_sampling_rules")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 create trace tail sampling rules request based on the context it is used
func (m *Configv1CreateTraceTailSamplingRulesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTraceTailSamplingRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateTraceTailSamplingRulesRequest) contextValidateTraceTailSamplingRules(ctx context.Context, formats strfmt.Registry) error {

	if m.TraceTailSamplingRules != nil {
		if err := m.TraceTailSamplingRules.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_tail_sampling_rules")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_tail_sampling_rules")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1CreateTraceTailSamplingRulesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1CreateTraceTailSamplingRulesRequest) UnmarshalBinary(b []byte) error {
	var res Configv1CreateTraceTailSamplingRulesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// ConfigunstableCreateTraceBehaviorRequest configunstable create trace behavior request
//
// swagger:model configunstableCreateTraceBehaviorRequest
type ConfigunstableCreateTraceBehaviorRequest struct {

	// If true, the TraceBehavior will not be created, and no response TraceBehavior will be returned. The response will return an error if the given TraceBehavior is invalid.
	DryRun bool `json:"dry_run,omitempty"`

	// trace behavior
	TraceBehavior *ConfigunstableTraceBehavior `json:"trace_behavior,omitempty"`
}

// Validate validates this configunstable create trace behavior request
func (m *ConfigunstableCreateTraceBehaviorRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraceBehavior(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateTraceBehaviorRequest) validateTraceBehavior(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceBehavior) { // not required
		return nil
	}

	if m.TraceBehavior != nil {
		if err := m.TraceBehavior.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_behavior")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_behavior")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable create trace behavior request based on the context it is used
func (m *ConfigunstableCreateTraceBehaviorRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTraceBehavior(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateTraceBehaviorRequest) contextValidateTraceBehavior(ctx context.Context, formats strfmt.Registry) error {

	if m.TraceBehavior != nil {
		if err := m.TraceBehavior.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_behavior")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_behavior")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableCreateTraceBehaviorRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableCreateTraceBehaviorRequest) UnmarshalBinary(b []byte) error {
	var res ConfigunstableCreateTraceBehaviorRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

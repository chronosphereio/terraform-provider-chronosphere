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

// ConfigunstableReadBudgetControlResponse configunstable read budget control response
//
// swagger:model configunstableReadBudgetControlResponse
type ConfigunstableReadBudgetControlResponse struct {

	// budget control
	BudgetControl *ConfigunstableBudgetControl `json:"budget_control,omitempty"`
}

// Validate validates this configunstable read budget control response
func (m *ConfigunstableReadBudgetControlResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBudgetControl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableReadBudgetControlResponse) validateBudgetControl(formats strfmt.Registry) error {
	if swag.IsZero(m.BudgetControl) { // not required
		return nil
	}

	if m.BudgetControl != nil {
		if err := m.BudgetControl.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("budget_control")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("budget_control")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable read budget control response based on the context it is used
func (m *ConfigunstableReadBudgetControlResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBudgetControl(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableReadBudgetControlResponse) contextValidateBudgetControl(ctx context.Context, formats strfmt.Registry) error {

	if m.BudgetControl != nil {
		if err := m.BudgetControl.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("budget_control")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("budget_control")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableReadBudgetControlResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableReadBudgetControlResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableReadBudgetControlResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

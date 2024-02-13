// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Configv1ListDropRulesResponse configv1 list drop rules response
//
// swagger:model configv1ListDropRulesResponse
type Configv1ListDropRulesResponse struct {

	// drop rules
	DropRules []*Configv1DropRule `json:"drop_rules"`

	// page
	Page *Configv1PageResult `json:"page,omitempty"`
}

// Validate validates this configv1 list drop rules response
func (m *Configv1ListDropRulesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDropRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1ListDropRulesResponse) validateDropRules(formats strfmt.Registry) error {
	if swag.IsZero(m.DropRules) { // not required
		return nil
	}

	for i := 0; i < len(m.DropRules); i++ {
		if swag.IsZero(m.DropRules[i]) { // not required
			continue
		}

		if m.DropRules[i] != nil {
			if err := m.DropRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("drop_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("drop_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Configv1ListDropRulesResponse) validatePage(formats strfmt.Registry) error {
	if swag.IsZero(m.Page) { // not required
		return nil
	}

	if m.Page != nil {
		if err := m.Page.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 list drop rules response based on the context it is used
func (m *Configv1ListDropRulesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDropRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1ListDropRulesResponse) contextValidateDropRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DropRules); i++ {

		if m.DropRules[i] != nil {
			if err := m.DropRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("drop_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("drop_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Configv1ListDropRulesResponse) contextValidatePage(ctx context.Context, formats strfmt.Registry) error {

	if m.Page != nil {
		if err := m.Page.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1ListDropRulesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1ListDropRulesResponse) UnmarshalBinary(b []byte) error {
	var res Configv1ListDropRulesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// ConfigunstableListTraceBehaviorsResponse configunstable list trace behaviors response
//
// swagger:model configunstableListTraceBehaviorsResponse
type ConfigunstableListTraceBehaviorsResponse struct {

	// page
	Page *Configv1PageResult `json:"page,omitempty"`

	// trace behaviors
	TraceBehaviors []*ConfigunstableTraceBehavior `json:"trace_behaviors"`
}

// Validate validates this configunstable list trace behaviors response
func (m *ConfigunstableListTraceBehaviorsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraceBehaviors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableListTraceBehaviorsResponse) validatePage(formats strfmt.Registry) error {
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

func (m *ConfigunstableListTraceBehaviorsResponse) validateTraceBehaviors(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceBehaviors) { // not required
		return nil
	}

	for i := 0; i < len(m.TraceBehaviors); i++ {
		if swag.IsZero(m.TraceBehaviors[i]) { // not required
			continue
		}

		if m.TraceBehaviors[i] != nil {
			if err := m.TraceBehaviors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("trace_behaviors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("trace_behaviors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this configunstable list trace behaviors response based on the context it is used
func (m *ConfigunstableListTraceBehaviorsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTraceBehaviors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableListTraceBehaviorsResponse) contextValidatePage(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConfigunstableListTraceBehaviorsResponse) contextValidateTraceBehaviors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TraceBehaviors); i++ {

		if m.TraceBehaviors[i] != nil {
			if err := m.TraceBehaviors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("trace_behaviors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("trace_behaviors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableListTraceBehaviorsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableListTraceBehaviorsResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableListTraceBehaviorsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

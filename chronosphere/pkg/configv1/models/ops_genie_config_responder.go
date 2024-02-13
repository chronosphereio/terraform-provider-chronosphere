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

// OpsGenieConfigResponder ops genie config responder
//
// swagger:model OpsGenieConfigResponder
type OpsGenieConfigResponder struct {

	// ID of the responder. Cannot be set if name or username are set.
	ID string `json:"id,omitempty"`

	// Name of the responder. Cannot be set if id or username are set.
	Name string `json:"name,omitempty"`

	// responder type
	ResponderType ResponderResponderType `json:"responder_type,omitempty"`

	// Username of the responder. Cannot be set if id or name are set.
	Username string `json:"username,omitempty"`
}

// Validate validates this ops genie config responder
func (m *OpsGenieConfigResponder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResponderType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpsGenieConfigResponder) validateResponderType(formats strfmt.Registry) error {
	if swag.IsZero(m.ResponderType) { // not required
		return nil
	}

	if err := m.ResponderType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("responder_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("responder_type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this ops genie config responder based on the context it is used
func (m *OpsGenieConfigResponder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResponderType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpsGenieConfigResponder) contextValidateResponderType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ResponderType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("responder_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("responder_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpsGenieConfigResponder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpsGenieConfigResponder) UnmarshalBinary(b []byte) error {
	var res OpsGenieConfigResponder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

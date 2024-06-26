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

// ConfigunstableUpdateLinkTemplateResponse configunstable update link template response
//
// swagger:model configunstableUpdateLinkTemplateResponse
type ConfigunstableUpdateLinkTemplateResponse struct {

	// link template
	LinkTemplate *ConfigunstableLinkTemplate `json:"link_template,omitempty"`
}

// Validate validates this configunstable update link template response
func (m *ConfigunstableUpdateLinkTemplateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinkTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableUpdateLinkTemplateResponse) validateLinkTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.LinkTemplate) { // not required
		return nil
	}

	if m.LinkTemplate != nil {
		if err := m.LinkTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link_template")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable update link template response based on the context it is used
func (m *ConfigunstableUpdateLinkTemplateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinkTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableUpdateLinkTemplateResponse) contextValidateLinkTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.LinkTemplate != nil {
		if err := m.LinkTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link_template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableUpdateLinkTemplateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableUpdateLinkTemplateResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableUpdateLinkTemplateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

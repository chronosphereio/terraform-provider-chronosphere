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

// ConfigunstableCreateLinkTemplateResponse configunstable create link template response
//
// swagger:model configunstableCreateLinkTemplateResponse
type ConfigunstableCreateLinkTemplateResponse struct {

	// link template
	LinkTemplate *ConfigunstableLinkTemplate `json:"link_template,omitempty"`
}

// Validate validates this configunstable create link template response
func (m *ConfigunstableCreateLinkTemplateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinkTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateLinkTemplateResponse) validateLinkTemplate(formats strfmt.Registry) error {
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

// ContextValidate validate this configunstable create link template response based on the context it is used
func (m *ConfigunstableCreateLinkTemplateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinkTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateLinkTemplateResponse) contextValidateLinkTemplate(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ConfigunstableCreateLinkTemplateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableCreateLinkTemplateResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableCreateLinkTemplateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

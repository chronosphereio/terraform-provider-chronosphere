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

// ConfigunstableListLinkTemplatesResponse configunstable list link templates response
//
// swagger:model configunstableListLinkTemplatesResponse
type ConfigunstableListLinkTemplatesResponse struct {

	// link templates
	LinkTemplates []*ConfigunstableLinkTemplate `json:"link_templates"`

	// page
	Page *Configv1PageResult `json:"page,omitempty"`
}

// Validate validates this configunstable list link templates response
func (m *ConfigunstableListLinkTemplatesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinkTemplates(formats); err != nil {
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

func (m *ConfigunstableListLinkTemplatesResponse) validateLinkTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.LinkTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.LinkTemplates); i++ {
		if swag.IsZero(m.LinkTemplates[i]) { // not required
			continue
		}

		if m.LinkTemplates[i] != nil {
			if err := m.LinkTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("link_templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("link_templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigunstableListLinkTemplatesResponse) validatePage(formats strfmt.Registry) error {
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

// ContextValidate validate this configunstable list link templates response based on the context it is used
func (m *ConfigunstableListLinkTemplatesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinkTemplates(ctx, formats); err != nil {
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

func (m *ConfigunstableListLinkTemplatesResponse) contextValidateLinkTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LinkTemplates); i++ {

		if m.LinkTemplates[i] != nil {
			if err := m.LinkTemplates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("link_templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("link_templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigunstableListLinkTemplatesResponse) contextValidatePage(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ConfigunstableListLinkTemplatesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableListLinkTemplatesResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableListLinkTemplatesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

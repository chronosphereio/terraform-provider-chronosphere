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

// NotifierWebhookConfig notifier webhook config
//
// swagger:model NotifierWebhookConfig
type NotifierWebhookConfig struct {

	// http config
	HTTPConfig *NotifierHTTPConfig `json:"http_config,omitempty"`

	// Required webhook URL (will be called as a POST request).
	URL string `json:"url,omitempty"`
}

// Validate validates this notifier webhook config
func (m *NotifierWebhookConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTPConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotifierWebhookConfig) validateHTTPConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTPConfig) { // not required
		return nil
	}

	if m.HTTPConfig != nil {
		if err := m.HTTPConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this notifier webhook config based on the context it is used
func (m *NotifierWebhookConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHTTPConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotifierWebhookConfig) contextValidateHTTPConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTPConfig != nil {
		if err := m.HTTPConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotifierWebhookConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotifierWebhookConfig) UnmarshalBinary(b []byte) error {
	var res NotifierWebhookConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

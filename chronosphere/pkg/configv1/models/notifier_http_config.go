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

// NotifierHTTPConfig notifier HTTP config
//
// swagger:model NotifierHTTPConfig
type NotifierHTTPConfig struct {

	// basic auth
	BasicAuth *HTTPConfigBasicAuth `json:"basic_auth,omitempty"`

	// Bearer token authentication. Cannot be set if basic_auth is set.
	BearerToken string `json:"bearer_token,omitempty"`

	// Optional proxy URL.
	ProxyURL string `json:"proxy_url,omitempty"`

	// tls config
	TLSConfig *HTTPConfigTLSConfig `json:"tls_config,omitempty"`
}

// Validate validates this notifier HTTP config
func (m *NotifierHTTPConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBasicAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLSConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotifierHTTPConfig) validateBasicAuth(formats strfmt.Registry) error {
	if swag.IsZero(m.BasicAuth) { // not required
		return nil
	}

	if m.BasicAuth != nil {
		if err := m.BasicAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basic_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basic_auth")
			}
			return err
		}
	}

	return nil
}

func (m *NotifierHTTPConfig) validateTLSConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.TLSConfig) { // not required
		return nil
	}

	if m.TLSConfig != nil {
		if err := m.TLSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this notifier HTTP config based on the context it is used
func (m *NotifierHTTPConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBasicAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTLSConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotifierHTTPConfig) contextValidateBasicAuth(ctx context.Context, formats strfmt.Registry) error {

	if m.BasicAuth != nil {
		if err := m.BasicAuth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basic_auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basic_auth")
			}
			return err
		}
	}

	return nil
}

func (m *NotifierHTTPConfig) contextValidateTLSConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.TLSConfig != nil {
		if err := m.TLSConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotifierHTTPConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotifierHTTPConfig) UnmarshalBinary(b []byte) error {
	var res NotifierHTTPConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

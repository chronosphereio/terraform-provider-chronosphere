// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotifierEmailConfig notifier email config
//
// swagger:model NotifierEmailConfig
type NotifierEmailConfig struct {

	// Optional HTML body of the email.
	HTML string `json:"html,omitempty"`

	// Optional text body of the email.
	Text string `json:"text,omitempty"`

	// Required email address to send notifications to.
	To string `json:"to,omitempty"`
}

// Validate validates this notifier email config
func (m *NotifierEmailConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this notifier email config based on context it is used
func (m *NotifierEmailConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotifierEmailConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotifierEmailConfig) UnmarshalBinary(b []byte) error {
	var res NotifierEmailConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

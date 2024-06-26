// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogScaleActionSlackPostMessageAction log scale action slack post message action
//
// swagger:model LogScaleActionSlackPostMessageAction
type LogScaleActionSlackPostMessageAction struct {

	// api token
	APIToken string `json:"api_token,omitempty"`

	// channels
	Channels []string `json:"channels"`

	// fields
	Fields map[string]string `json:"fields,omitempty"`

	// use proxy
	UseProxy bool `json:"use_proxy,omitempty"`
}

// Validate validates this log scale action slack post message action
func (m *LogScaleActionSlackPostMessageAction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this log scale action slack post message action based on context it is used
func (m *LogScaleActionSlackPostMessageAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogScaleActionSlackPostMessageAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogScaleActionSlackPostMessageAction) UnmarshalBinary(b []byte) error {
	var res LogScaleActionSlackPostMessageAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

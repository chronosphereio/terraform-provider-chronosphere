// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotificationPolicyRoutesGroupBy notification policy routes group by
//
// swagger:model NotificationPolicyRoutesGroupBy
type NotificationPolicyRoutesGroupBy struct {

	// Set of label names used to group alerts.
	// For example, if label_names is ["service", "code"] then all alerts including labels {service="foo",code="404"}
	// will be grouped together.
	LabelNames []string `json:"label_names"`
}

// Validate validates this notification policy routes group by
func (m *NotificationPolicyRoutesGroupBy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this notification policy routes group by based on context it is used
func (m *NotificationPolicyRoutesGroupBy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotificationPolicyRoutesGroupBy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationPolicyRoutesGroupBy) UnmarshalBinary(b []byte) error {
	var res NotificationPolicyRoutesGroupBy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

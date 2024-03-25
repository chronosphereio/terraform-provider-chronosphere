// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GcpMetricsIntegrationServiceAccount gcp metrics integration service account
//
// swagger:model GcpMetricsIntegrationServiceAccount
type GcpMetricsIntegrationServiceAccount struct {

	// Email address of the service account to impersonate.
	ClientEmail string `json:"client_email,omitempty"`
}

// Validate validates this gcp metrics integration service account
func (m *GcpMetricsIntegrationServiceAccount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gcp metrics integration service account based on context it is used
func (m *GcpMetricsIntegrationServiceAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GcpMetricsIntegrationServiceAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpMetricsIntegrationServiceAccount) UnmarshalBinary(b []byte) error {
	var res GcpMetricsIntegrationServiceAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

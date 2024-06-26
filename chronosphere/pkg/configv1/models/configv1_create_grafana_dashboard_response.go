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

// Configv1CreateGrafanaDashboardResponse configv1 create grafana dashboard response
//
// swagger:model configv1CreateGrafanaDashboardResponse
type Configv1CreateGrafanaDashboardResponse struct {

	// grafana dashboard
	GrafanaDashboard *Configv1GrafanaDashboard `json:"grafana_dashboard,omitempty"`
}

// Validate validates this configv1 create grafana dashboard response
func (m *Configv1CreateGrafanaDashboardResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGrafanaDashboard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateGrafanaDashboardResponse) validateGrafanaDashboard(formats strfmt.Registry) error {
	if swag.IsZero(m.GrafanaDashboard) { // not required
		return nil
	}

	if m.GrafanaDashboard != nil {
		if err := m.GrafanaDashboard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grafana_dashboard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grafana_dashboard")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 create grafana dashboard response based on the context it is used
func (m *Configv1CreateGrafanaDashboardResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGrafanaDashboard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1CreateGrafanaDashboardResponse) contextValidateGrafanaDashboard(ctx context.Context, formats strfmt.Registry) error {

	if m.GrafanaDashboard != nil {
		if err := m.GrafanaDashboard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grafana_dashboard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grafana_dashboard")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1CreateGrafanaDashboardResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1CreateGrafanaDashboardResponse) UnmarshalBinary(b []byte) error {
	var res Configv1CreateGrafanaDashboardResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

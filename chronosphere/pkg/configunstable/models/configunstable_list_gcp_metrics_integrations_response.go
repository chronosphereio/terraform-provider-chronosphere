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

// ConfigunstableListGcpMetricsIntegrationsResponse configunstable list gcp metrics integrations response
//
// swagger:model configunstableListGcpMetricsIntegrationsResponse
type ConfigunstableListGcpMetricsIntegrationsResponse struct {

	// gcp metrics integrations
	GcpMetricsIntegrations []*ConfigunstableGcpMetricsIntegration `json:"gcp_metrics_integrations"`

	// page
	Page *Configv1PageResult `json:"page,omitempty"`
}

// Validate validates this configunstable list gcp metrics integrations response
func (m *ConfigunstableListGcpMetricsIntegrationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGcpMetricsIntegrations(formats); err != nil {
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

func (m *ConfigunstableListGcpMetricsIntegrationsResponse) validateGcpMetricsIntegrations(formats strfmt.Registry) error {
	if swag.IsZero(m.GcpMetricsIntegrations) { // not required
		return nil
	}

	for i := 0; i < len(m.GcpMetricsIntegrations); i++ {
		if swag.IsZero(m.GcpMetricsIntegrations[i]) { // not required
			continue
		}

		if m.GcpMetricsIntegrations[i] != nil {
			if err := m.GcpMetricsIntegrations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gcp_metrics_integrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gcp_metrics_integrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigunstableListGcpMetricsIntegrationsResponse) validatePage(formats strfmt.Registry) error {
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

// ContextValidate validate this configunstable list gcp metrics integrations response based on the context it is used
func (m *ConfigunstableListGcpMetricsIntegrationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGcpMetricsIntegrations(ctx, formats); err != nil {
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

func (m *ConfigunstableListGcpMetricsIntegrationsResponse) contextValidateGcpMetricsIntegrations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GcpMetricsIntegrations); i++ {

		if m.GcpMetricsIntegrations[i] != nil {
			if err := m.GcpMetricsIntegrations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gcp_metrics_integrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gcp_metrics_integrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigunstableListGcpMetricsIntegrationsResponse) contextValidatePage(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ConfigunstableListGcpMetricsIntegrationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableListGcpMetricsIntegrationsResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableListGcpMetricsIntegrationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

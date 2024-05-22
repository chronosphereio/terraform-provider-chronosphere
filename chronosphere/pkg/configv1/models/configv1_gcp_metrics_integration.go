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
	"github.com/go-openapi/validate"
)

// Configv1GcpMetricsIntegration configv1 gcp metrics integration
//
// swagger:model configv1GcpMetricsIntegration
type Configv1GcpMetricsIntegration struct {

	// Timestamp of when the GcpMetricsIntegration was created. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Metric groups to be ingested for this integration.
	MetricGroups []*GcpMetricsIntegrationMetricGroup `json:"metric_groups"`

	// Required name of the GcpMetricsIntegration. May be modified after the GcpMetricsIntegration is created.
	Name string `json:"name,omitempty"`

	// service account
	ServiceAccount *Configv1GcpMetricsIntegrationServiceAccount `json:"service_account,omitempty"`

	// Unique identifier of the GcpMetricsIntegration. If slug is not provided, one will be generated based of the name field. Cannot be modified after the GcpMetricsIntegration is created.
	Slug string `json:"slug,omitempty"`

	// Timestamp of when the GcpMetricsIntegration was last updated. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this configv1 gcp metrics integration
func (m *Configv1GcpMetricsIntegration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1GcpMetricsIntegration) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1GcpMetricsIntegration) validateMetricGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.MetricGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.MetricGroups); i++ {
		if swag.IsZero(m.MetricGroups[i]) { // not required
			continue
		}

		if m.MetricGroups[i] != nil {
			if err := m.MetricGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metric_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metric_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Configv1GcpMetricsIntegration) validateServiceAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceAccount) { // not required
		return nil
	}

	if m.ServiceAccount != nil {
		if err := m.ServiceAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_account")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1GcpMetricsIntegration) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this configv1 gcp metrics integration based on the context it is used
func (m *Configv1GcpMetricsIntegration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetricGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1GcpMetricsIntegration) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1GcpMetricsIntegration) contextValidateMetricGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MetricGroups); i++ {

		if m.MetricGroups[i] != nil {
			if err := m.MetricGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metric_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metric_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Configv1GcpMetricsIntegration) contextValidateServiceAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceAccount != nil {
		if err := m.ServiceAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_account")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1GcpMetricsIntegration) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1GcpMetricsIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1GcpMetricsIntegration) UnmarshalBinary(b []byte) error {
	var res Configv1GcpMetricsIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
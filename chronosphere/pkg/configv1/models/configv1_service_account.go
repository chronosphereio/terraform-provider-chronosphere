// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Configv1ServiceAccount configv1 service account
//
// swagger:model configv1ServiceAccount
type Configv1ServiceAccount struct {

	// Timestamp of when the ServiceAccount was created. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// email is the generated email address of the service account. Cannot be set
	// by clients.
	// Read Only: true
	Email string `json:"email,omitempty"`

	// metrics restriction
	MetricsRestriction *ServiceAccountMetricsRestriction `json:"metrics_restriction,omitempty"`

	// Required name of the ServiceAccount. May be modified after the ServiceAccount is created.
	Name string `json:"name,omitempty"`

	// Unique identifier of the ServiceAccount. If slug is not provided, one will be generated based of the name field. Cannot be modified after the ServiceAccount is created.
	Slug string `json:"slug,omitempty"`

	// token is the generated API token of the service account. Cannot be set by
	// clients.
	//
	// token is only set once by the server in the CreateServiceAccount response.
	// ReadServiceAccount will always return an empty token. Therefore, when
	// creating a service account, clients are responsible for securely storing
	// the response token on their end, as they will not be able to read it
	// again.
	// Read Only: true
	Token string `json:"token,omitempty"`

	// If set, grants the service account access to all Chronosphere APIs
	// (including resource configuration and metric data) within the access
	// controls defined by the service account's team membership.
	//
	// Exactly one of unrestricted or metrics_restriction must be set.
	Unrestricted bool `json:"unrestricted,omitempty"`

	// Timestamp of when the ServiceAccount was last updated. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this configv1 service account
func (m *Configv1ServiceAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricsRestriction(formats); err != nil {
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

func (m *Configv1ServiceAccount) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1ServiceAccount) validateMetricsRestriction(formats strfmt.Registry) error {
	if swag.IsZero(m.MetricsRestriction) { // not required
		return nil
	}

	if m.MetricsRestriction != nil {
		if err := m.MetricsRestriction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics_restriction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metrics_restriction")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1ServiceAccount) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this configv1 service account based on the context it is used
func (m *Configv1ServiceAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetricsRestriction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToken(ctx, formats); err != nil {
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

func (m *Configv1ServiceAccount) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1ServiceAccount) contextValidateEmail(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "email", "body", string(m.Email)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1ServiceAccount) contextValidateMetricsRestriction(ctx context.Context, formats strfmt.Registry) error {

	if m.MetricsRestriction != nil {
		if err := m.MetricsRestriction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metrics_restriction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metrics_restriction")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1ServiceAccount) contextValidateToken(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "token", "body", string(m.Token)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1ServiceAccount) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1ServiceAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1ServiceAccount) UnmarshalBinary(b []byte) error {
	var res Configv1ServiceAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

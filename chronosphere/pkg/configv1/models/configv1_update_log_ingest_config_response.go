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

// Configv1UpdateLogIngestConfigResponse configv1 update log ingest config response
//
// swagger:model configv1UpdateLogIngestConfigResponse
type Configv1UpdateLogIngestConfigResponse struct {

	// log ingest config
	LogIngestConfig *Configv1LogIngestConfig `json:"log_ingest_config,omitempty"`
}

// Validate validates this configv1 update log ingest config response
func (m *Configv1UpdateLogIngestConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogIngestConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1UpdateLogIngestConfigResponse) validateLogIngestConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.LogIngestConfig) { // not required
		return nil
	}

	if m.LogIngestConfig != nil {
		if err := m.LogIngestConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_ingest_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_ingest_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configv1 update log ingest config response based on the context it is used
func (m *Configv1UpdateLogIngestConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogIngestConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1UpdateLogIngestConfigResponse) contextValidateLogIngestConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.LogIngestConfig != nil {
		if err := m.LogIngestConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("log_ingest_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("log_ingest_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1UpdateLogIngestConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1UpdateLogIngestConfigResponse) UnmarshalBinary(b []byte) error {
	var res Configv1UpdateLogIngestConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// ChronoConfigSeverityMapping chrono config severity mapping
//
// swagger:model ChronoConfigSeverityMapping
type ChronoConfigSeverityMapping struct {

	// chronosphere severity
	ChronosphereSeverity ChronoConfigSeverityMappingSeverity `json:"chronosphere_severity,omitempty"`

	// severity
	Severity string `json:"severity,omitempty"`
}

// Validate validates this chrono config severity mapping
func (m *ChronoConfigSeverityMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChronosphereSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChronoConfigSeverityMapping) validateChronosphereSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.ChronosphereSeverity) { // not required
		return nil
	}

	if err := m.ChronosphereSeverity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("chronosphere_severity")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("chronosphere_severity")
		}
		return err
	}

	return nil
}

// ContextValidate validate this chrono config severity mapping based on the context it is used
func (m *ChronoConfigSeverityMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChronosphereSeverity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChronoConfigSeverityMapping) contextValidateChronosphereSeverity(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ChronosphereSeverity.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("chronosphere_severity")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("chronosphere_severity")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChronoConfigSeverityMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChronoConfigSeverityMapping) UnmarshalBinary(b []byte) error {
	var res ChronoConfigSeverityMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

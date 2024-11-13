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

// SLODefinition s l o definition
//
// swagger:model SLODefinition
type SLODefinition struct {

	// Configured whether this SLO is for a low volume event (< 1/s). This will
	// adjust the SLI queries to account for the low volume nature of the event.
	LowVolume bool `json:"low_volume,omitempty"`

	// The SLO objective
	Objective float64 `json:"objective,omitempty"`

	// The reporting windows for this SLO. The SLO is considered breached if the
	// error budget is depleted in any of these windows.
	ReportingWindows []*DefinitionTimeWindow `json:"reporting_windows"`
}

// Validate validates this s l o definition
func (m *SLODefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReportingWindows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SLODefinition) validateReportingWindows(formats strfmt.Registry) error {
	if swag.IsZero(m.ReportingWindows) { // not required
		return nil
	}

	for i := 0; i < len(m.ReportingWindows); i++ {
		if swag.IsZero(m.ReportingWindows[i]) { // not required
			continue
		}

		if m.ReportingWindows[i] != nil {
			if err := m.ReportingWindows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reporting_windows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reporting_windows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this s l o definition based on the context it is used
func (m *SLODefinition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReportingWindows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SLODefinition) contextValidateReportingWindows(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReportingWindows); i++ {

		if m.ReportingWindows[i] != nil {
			if err := m.ReportingWindows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reporting_windows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reporting_windows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SLODefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SLODefinition) UnmarshalBinary(b []byte) error {
	var res SLODefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

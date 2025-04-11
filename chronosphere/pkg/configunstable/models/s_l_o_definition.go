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

	// Provides the burn rate alert configuration for the SLO. If not provided the
	// default burn rates will be used. The configuration is only valid if the
	// enable_burn_rate_alerting flag is set to true.
	BurnRateAlertingConfig []*DefinitionBurnRateDefinition `json:"burn_rate_alerting_config"`

	// If true enables burn rate alerting.
	EnableBurnRateAlerting bool `json:"enable_burn_rate_alerting,omitempty"`

	// The SLO objective
	Objective float64 `json:"objective,omitempty"`

	// This is deprecated.
	ReportingWindows []*DefinitionTimeWindow `json:"reporting_windows"`

	// time window
	TimeWindow *DefinitionTimeWindow `json:"time_window,omitempty"`
}

// Validate validates this s l o definition
func (m *SLODefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBurnRateAlertingConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportingWindows(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeWindow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SLODefinition) validateBurnRateAlertingConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.BurnRateAlertingConfig) { // not required
		return nil
	}

	for i := 0; i < len(m.BurnRateAlertingConfig); i++ {
		if swag.IsZero(m.BurnRateAlertingConfig[i]) { // not required
			continue
		}

		if m.BurnRateAlertingConfig[i] != nil {
			if err := m.BurnRateAlertingConfig[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("burn_rate_alerting_config" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("burn_rate_alerting_config" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

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

func (m *SLODefinition) validateTimeWindow(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeWindow) { // not required
		return nil
	}

	if m.TimeWindow != nil {
		if err := m.TimeWindow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time_window")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("time_window")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s l o definition based on the context it is used
func (m *SLODefinition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBurnRateAlertingConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReportingWindows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeWindow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SLODefinition) contextValidateBurnRateAlertingConfig(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BurnRateAlertingConfig); i++ {

		if m.BurnRateAlertingConfig[i] != nil {
			if err := m.BurnRateAlertingConfig[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("burn_rate_alerting_config" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("burn_rate_alerting_config" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

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

func (m *SLODefinition) contextValidateTimeWindow(ctx context.Context, formats strfmt.Registry) error {

	if m.TimeWindow != nil {
		if err := m.TimeWindow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time_window")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("time_window")
			}
			return err
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

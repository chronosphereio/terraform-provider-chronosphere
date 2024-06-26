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

// MonitorSchedule monitor schedule
//
// swagger:model MonitorSchedule
type MonitorSchedule struct {

	// The timezone of the time ranges.
	Timezone string `json:"timezone,omitempty"`

	// weekly schedule
	WeeklySchedule *ScheduleWeeklySchedule `json:"weekly_schedule,omitempty"`
}

// Validate validates this monitor schedule
func (m *MonitorSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWeeklySchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MonitorSchedule) validateWeeklySchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.WeeklySchedule) { // not required
		return nil
	}

	if m.WeeklySchedule != nil {
		if err := m.WeeklySchedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weekly_schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weekly_schedule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this monitor schedule based on the context it is used
func (m *MonitorSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWeeklySchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MonitorSchedule) contextValidateWeeklySchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.WeeklySchedule != nil {
		if err := m.WeeklySchedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weekly_schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weekly_schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MonitorSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MonitorSchedule) UnmarshalBinary(b []byte) error {
	var res MonitorSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

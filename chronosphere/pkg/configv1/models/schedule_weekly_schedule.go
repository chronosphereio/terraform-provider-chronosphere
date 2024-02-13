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

// ScheduleWeeklySchedule schedule weekly schedule
//
// swagger:model ScheduleWeeklySchedule
type ScheduleWeeklySchedule struct {

	// friday
	Friday *ScheduleScheduleDay `json:"friday,omitempty"`

	// monday
	Monday *ScheduleScheduleDay `json:"monday,omitempty"`

	// saturday
	Saturday *ScheduleScheduleDay `json:"saturday,omitempty"`

	// sunday
	Sunday *ScheduleScheduleDay `json:"sunday,omitempty"`

	// thursday
	Thursday *ScheduleScheduleDay `json:"thursday,omitempty"`

	// tuesday
	Tuesday *ScheduleScheduleDay `json:"tuesday,omitempty"`

	// wednesday
	Wednesday *ScheduleScheduleDay `json:"wednesday,omitempty"`
}

// Validate validates this schedule weekly schedule
func (m *ScheduleWeeklySchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFriday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSaturday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSunday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThursday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTuesday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWednesday(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleWeeklySchedule) validateFriday(formats strfmt.Registry) error {
	if swag.IsZero(m.Friday) { // not required
		return nil
	}

	if m.Friday != nil {
		if err := m.Friday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("friday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("friday")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleWeeklySchedule) validateMonday(formats strfmt.Registry) error {
	if swag.IsZero(m.Monday) { // not required
		return nil
	}

	if m.Monday != nil {
		if err := m.Monday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monday")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleWeeklySchedule) validateSaturday(formats strfmt.Registry) error {
	if swag.IsZero(m.Saturday) { // not required
		return nil
	}

	if m.Saturday != nil {
		if err := m.Saturday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("saturday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("saturday")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleWeeklySchedule) validateSunday(formats strfmt.Registry) error {
	if swag.IsZero(m.Sunday) { // not required
		return nil
	}

	if m.Sunday != nil {
		if err := m.Sunday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sunday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sunday")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleWeeklySchedule) validateThursday(formats strfmt.Registry) error {
	if swag.IsZero(m.Thursday) { // not required
		return nil
	}

	if m.Thursday != nil {
		if err := m.Thursday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thursday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("thursday")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleWeeklySchedule) validateTuesday(formats strfmt.Registry) error {
	if swag.IsZero(m.Tuesday) { // not required
		return nil
	}

	if m.Tuesday != nil {
		if err := m.Tuesday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tuesday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tuesday")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleWeeklySchedule) validateWednesday(formats strfmt.Registry) error {
	if swag.IsZero(m.Wednesday) { // not required
		return nil
	}

	if m.Wednesday != nil {
		if err := m.Wednesday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wednesday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wednesday")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this schedule weekly schedule based on the context it is used
func (m *ScheduleWeeklySchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFriday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSaturday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSunday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThursday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTuesday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWednesday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleWeeklySchedule) contextValidateFriday(ctx context.Context, formats strfmt.Registry) error {

	if m.Friday != nil {
		if err := m.Friday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("friday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("friday")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleWeeklySchedule) contextValidateMonday(ctx context.Context, formats strfmt.Registry) error {

	if m.Monday != nil {
		if err := m.Monday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monday")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleWeeklySchedule) contextValidateSaturday(ctx context.Context, formats strfmt.Registry) error {

	if m.Saturday != nil {
		if err := m.Saturday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("saturday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("saturday")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleWeeklySchedule) contextValidateSunday(ctx context.Context, formats strfmt.Registry) error {

	if m.Sunday != nil {
		if err := m.Sunday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sunday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sunday")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleWeeklySchedule) contextValidateThursday(ctx context.Context, formats strfmt.Registry) error {

	if m.Thursday != nil {
		if err := m.Thursday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thursday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("thursday")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleWeeklySchedule) contextValidateTuesday(ctx context.Context, formats strfmt.Registry) error {

	if m.Tuesday != nil {
		if err := m.Tuesday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tuesday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tuesday")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleWeeklySchedule) contextValidateWednesday(ctx context.Context, formats strfmt.Registry) error {

	if m.Wednesday != nil {
		if err := m.Wednesday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wednesday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wednesday")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduleWeeklySchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleWeeklySchedule) UnmarshalBinary(b []byte) error {
	var res ScheduleWeeklySchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

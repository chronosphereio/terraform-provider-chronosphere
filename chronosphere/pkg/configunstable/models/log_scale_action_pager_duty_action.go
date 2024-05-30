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

// LogScaleActionPagerDutyAction log scale action pager duty action
//
// swagger:model LogScaleActionPagerDutyAction
type LogScaleActionPagerDutyAction struct {

	// routing key
	RoutingKey string `json:"routing_key,omitempty"`

	// severity
	Severity LogScaleActionPagerDutyActionSeverity `json:"severity,omitempty"`

	// use proxy
	UseProxy bool `json:"use_proxy,omitempty"`
}

// Validate validates this log scale action pager duty action
func (m *LogScaleActionPagerDutyAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogScaleActionPagerDutyAction) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	if err := m.Severity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("severity")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("severity")
		}
		return err
	}

	return nil
}

// ContextValidate validate this log scale action pager duty action based on the context it is used
func (m *LogScaleActionPagerDutyAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSeverity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogScaleActionPagerDutyAction) contextValidateSeverity(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Severity.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("severity")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("severity")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogScaleActionPagerDutyAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogScaleActionPagerDutyAction) UnmarshalBinary(b []byte) error {
	var res LogScaleActionPagerDutyAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

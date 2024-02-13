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

// TraceMetricsRuleGroupBy GroupBy contains fields required to group the resultant metrics of a TraceMetricsRule by a specific key.
//
// swagger:model TraceMetricsRuleGroupBy
type TraceMetricsRuleGroupBy struct {

	// key
	Key *GroupByGroupByKey `json:"key,omitempty"`

	// The label to use in the resultant metrics.
	Label string `json:"label,omitempty"`
}

// Validate validates this trace metrics rule group by
func (m *TraceMetricsRuleGroupBy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceMetricsRuleGroupBy) validateKey(formats strfmt.Registry) error {
	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if m.Key != nil {
		if err := m.Key.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this trace metrics rule group by based on the context it is used
func (m *TraceMetricsRuleGroupBy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceMetricsRuleGroupBy) contextValidateKey(ctx context.Context, formats strfmt.Registry) error {

	if m.Key != nil {
		if err := m.Key.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TraceMetricsRuleGroupBy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraceMetricsRuleGroupBy) UnmarshalBinary(b []byte) error {
	var res TraceMetricsRuleGroupBy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

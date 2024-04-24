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

// TraceSearchFilterStringFilter trace search filter string filter
//
// swagger:model TraceSearchFilterStringFilter
type TraceSearchFilterStringFilter struct {

	// Values the filter tests against when using IN or NOT_IN match type.
	InValues []string `json:"in_values"`

	// match
	Match StringFilterStringFilterMatchType `json:"match,omitempty"`

	// The value the filter compares to the target trace or span field.
	Value string `json:"value,omitempty"`
}

// Validate validates this trace search filter string filter
func (m *TraceSearchFilterStringFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceSearchFilterStringFilter) validateMatch(formats strfmt.Registry) error {
	if swag.IsZero(m.Match) { // not required
		return nil
	}

	if err := m.Match.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("match")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("match")
		}
		return err
	}

	return nil
}

// ContextValidate validate this trace search filter string filter based on the context it is used
func (m *TraceSearchFilterStringFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMatch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceSearchFilterStringFilter) contextValidateMatch(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Match.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("match")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("match")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TraceSearchFilterStringFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraceSearchFilterStringFilter) UnmarshalBinary(b []byte) error {
	var res TraceSearchFilterStringFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

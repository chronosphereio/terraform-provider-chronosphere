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

// SavedTraceSearchTimeFilter saved trace search time filter
//
// swagger:model SavedTraceSearchTimeFilter
type SavedTraceSearchTimeFilter struct {

	// between
	Between *SavedTraceSearchBetweenTimeFilter `json:"between,omitempty"`

	// close to
	CloseTo *SavedTraceSearchCloseToTimeFilter `json:"close_to,omitempty"`

	// relative
	Relative *SavedTraceSearchRelativeTimeFilter `json:"relative,omitempty"`
}

// Validate validates this saved trace search time filter
func (m *SavedTraceSearchTimeFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBetween(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloseTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelative(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SavedTraceSearchTimeFilter) validateBetween(formats strfmt.Registry) error {
	if swag.IsZero(m.Between) { // not required
		return nil
	}

	if m.Between != nil {
		if err := m.Between.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("between")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("between")
			}
			return err
		}
	}

	return nil
}

func (m *SavedTraceSearchTimeFilter) validateCloseTo(formats strfmt.Registry) error {
	if swag.IsZero(m.CloseTo) { // not required
		return nil
	}

	if m.CloseTo != nil {
		if err := m.CloseTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("close_to")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("close_to")
			}
			return err
		}
	}

	return nil
}

func (m *SavedTraceSearchTimeFilter) validateRelative(formats strfmt.Registry) error {
	if swag.IsZero(m.Relative) { // not required
		return nil
	}

	if m.Relative != nil {
		if err := m.Relative.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relative")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("relative")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this saved trace search time filter based on the context it is used
func (m *SavedTraceSearchTimeFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBetween(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCloseTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelative(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SavedTraceSearchTimeFilter) contextValidateBetween(ctx context.Context, formats strfmt.Registry) error {

	if m.Between != nil {
		if err := m.Between.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("between")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("between")
			}
			return err
		}
	}

	return nil
}

func (m *SavedTraceSearchTimeFilter) contextValidateCloseTo(ctx context.Context, formats strfmt.Registry) error {

	if m.CloseTo != nil {
		if err := m.CloseTo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("close_to")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("close_to")
			}
			return err
		}
	}

	return nil
}

func (m *SavedTraceSearchTimeFilter) contextValidateRelative(ctx context.Context, formats strfmt.Registry) error {

	if m.Relative != nil {
		if err := m.Relative.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relative")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("relative")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SavedTraceSearchTimeFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SavedTraceSearchTimeFilter) UnmarshalBinary(b []byte) error {
	var res SavedTraceSearchTimeFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
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

// TraceBehaviorBaselineBehaviorOptions trace behavior baseline behavior options
//
// swagger:model TraceBehaviorBaselineBehaviorOptions
type TraceBehaviorBaselineBehaviorOptions struct {

	// Sample rate for head sampling. This applies to all root spans that are enrolled in head sampling,
	// but do not have a specific rule defined for their service.
	BaseHeadSampleRate float64 `json:"base_head_sample_rate,omitempty"`

	// Sample rate for fully assembled traces that do not apply to the fast, slow, or error sampling options.
	BaseTailSampleRate float64 `json:"base_tail_sample_rate,omitempty"`

	// error sample options
	ErrorSampleOptions *BaselineBehaviorOptionsErrorSampleOptions `json:"error_sample_options,omitempty"`

	// fast sample options
	FastSampleOptions *BaselineBehaviorOptionsFastSampleOptions `json:"fast_sample_options,omitempty"`

	// slow sample options
	SlowSampleOptions *BaselineBehaviorOptionsSlowSampleOptions `json:"slow_sample_options,omitempty"`
}

// Validate validates this trace behavior baseline behavior options
func (m *TraceBehaviorBaselineBehaviorOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrorSampleOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFastSampleOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlowSampleOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceBehaviorBaselineBehaviorOptions) validateErrorSampleOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrorSampleOptions) { // not required
		return nil
	}

	if m.ErrorSampleOptions != nil {
		if err := m.ErrorSampleOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error_sample_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error_sample_options")
			}
			return err
		}
	}

	return nil
}

func (m *TraceBehaviorBaselineBehaviorOptions) validateFastSampleOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.FastSampleOptions) { // not required
		return nil
	}

	if m.FastSampleOptions != nil {
		if err := m.FastSampleOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fast_sample_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fast_sample_options")
			}
			return err
		}
	}

	return nil
}

func (m *TraceBehaviorBaselineBehaviorOptions) validateSlowSampleOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.SlowSampleOptions) { // not required
		return nil
	}

	if m.SlowSampleOptions != nil {
		if err := m.SlowSampleOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slow_sample_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slow_sample_options")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this trace behavior baseline behavior options based on the context it is used
func (m *TraceBehaviorBaselineBehaviorOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrorSampleOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFastSampleOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlowSampleOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceBehaviorBaselineBehaviorOptions) contextValidateErrorSampleOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.ErrorSampleOptions != nil {
		if err := m.ErrorSampleOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error_sample_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error_sample_options")
			}
			return err
		}
	}

	return nil
}

func (m *TraceBehaviorBaselineBehaviorOptions) contextValidateFastSampleOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.FastSampleOptions != nil {
		if err := m.FastSampleOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fast_sample_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fast_sample_options")
			}
			return err
		}
	}

	return nil
}

func (m *TraceBehaviorBaselineBehaviorOptions) contextValidateSlowSampleOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.SlowSampleOptions != nil {
		if err := m.SlowSampleOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slow_sample_options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slow_sample_options")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TraceBehaviorBaselineBehaviorOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraceBehaviorBaselineBehaviorOptions) UnmarshalBinary(b []byte) error {
	var res TraceBehaviorBaselineBehaviorOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

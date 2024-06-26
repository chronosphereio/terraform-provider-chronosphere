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

// ConfigunstableUpdateTraceJaegerRemoteSamplingStrategyResponse configunstable update trace jaeger remote sampling strategy response
//
// swagger:model configunstableUpdateTraceJaegerRemoteSamplingStrategyResponse
type ConfigunstableUpdateTraceJaegerRemoteSamplingStrategyResponse struct {

	// trace jaeger remote sampling strategy
	TraceJaegerRemoteSamplingStrategy *ConfigunstableTraceJaegerRemoteSamplingStrategy `json:"trace_jaeger_remote_sampling_strategy,omitempty"`
}

// Validate validates this configunstable update trace jaeger remote sampling strategy response
func (m *ConfigunstableUpdateTraceJaegerRemoteSamplingStrategyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraceJaegerRemoteSamplingStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableUpdateTraceJaegerRemoteSamplingStrategyResponse) validateTraceJaegerRemoteSamplingStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceJaegerRemoteSamplingStrategy) { // not required
		return nil
	}

	if m.TraceJaegerRemoteSamplingStrategy != nil {
		if err := m.TraceJaegerRemoteSamplingStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_jaeger_remote_sampling_strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_jaeger_remote_sampling_strategy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable update trace jaeger remote sampling strategy response based on the context it is used
func (m *ConfigunstableUpdateTraceJaegerRemoteSamplingStrategyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTraceJaegerRemoteSamplingStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableUpdateTraceJaegerRemoteSamplingStrategyResponse) contextValidateTraceJaegerRemoteSamplingStrategy(ctx context.Context, formats strfmt.Registry) error {

	if m.TraceJaegerRemoteSamplingStrategy != nil {
		if err := m.TraceJaegerRemoteSamplingStrategy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_jaeger_remote_sampling_strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_jaeger_remote_sampling_strategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableUpdateTraceJaegerRemoteSamplingStrategyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableUpdateTraceJaegerRemoteSamplingStrategyResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableUpdateTraceJaegerRemoteSamplingStrategyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

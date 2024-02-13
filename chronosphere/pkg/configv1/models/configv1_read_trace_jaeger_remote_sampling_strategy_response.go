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

// Configv1ReadTraceJaegerRemoteSamplingStrategyResponse configv1 read trace jaeger remote sampling strategy response
//
// swagger:model configv1ReadTraceJaegerRemoteSamplingStrategyResponse
type Configv1ReadTraceJaegerRemoteSamplingStrategyResponse struct {

	// trace jaeger remote sampling strategy
	TraceJaegerRemoteSamplingStrategy *Configv1TraceJaegerRemoteSamplingStrategy `json:"trace_jaeger_remote_sampling_strategy,omitempty"`
}

// Validate validates this configv1 read trace jaeger remote sampling strategy response
func (m *Configv1ReadTraceJaegerRemoteSamplingStrategyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraceJaegerRemoteSamplingStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1ReadTraceJaegerRemoteSamplingStrategyResponse) validateTraceJaegerRemoteSamplingStrategy(formats strfmt.Registry) error {
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

// ContextValidate validate this configv1 read trace jaeger remote sampling strategy response based on the context it is used
func (m *Configv1ReadTraceJaegerRemoteSamplingStrategyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTraceJaegerRemoteSamplingStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1ReadTraceJaegerRemoteSamplingStrategyResponse) contextValidateTraceJaegerRemoteSamplingStrategy(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Configv1ReadTraceJaegerRemoteSamplingStrategyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1ReadTraceJaegerRemoteSamplingStrategyResponse) UnmarshalBinary(b []byte) error {
	var res Configv1ReadTraceJaegerRemoteSamplingStrategyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

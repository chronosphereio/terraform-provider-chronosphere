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

// Configv1ListTraceJaegerRemoteSamplingStrategiesResponse configv1 list trace jaeger remote sampling strategies response
//
// swagger:model configv1ListTraceJaegerRemoteSamplingStrategiesResponse
type Configv1ListTraceJaegerRemoteSamplingStrategiesResponse struct {

	// page
	Page *Configv1PageResult `json:"page,omitempty"`

	// trace jaeger remote sampling strategies
	TraceJaegerRemoteSamplingStrategies []*Configv1TraceJaegerRemoteSamplingStrategy `json:"trace_jaeger_remote_sampling_strategies"`
}

// Validate validates this configv1 list trace jaeger remote sampling strategies response
func (m *Configv1ListTraceJaegerRemoteSamplingStrategiesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraceJaegerRemoteSamplingStrategies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1ListTraceJaegerRemoteSamplingStrategiesResponse) validatePage(formats strfmt.Registry) error {
	if swag.IsZero(m.Page) { // not required
		return nil
	}

	if m.Page != nil {
		if err := m.Page.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1ListTraceJaegerRemoteSamplingStrategiesResponse) validateTraceJaegerRemoteSamplingStrategies(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceJaegerRemoteSamplingStrategies) { // not required
		return nil
	}

	for i := 0; i < len(m.TraceJaegerRemoteSamplingStrategies); i++ {
		if swag.IsZero(m.TraceJaegerRemoteSamplingStrategies[i]) { // not required
			continue
		}

		if m.TraceJaegerRemoteSamplingStrategies[i] != nil {
			if err := m.TraceJaegerRemoteSamplingStrategies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("trace_jaeger_remote_sampling_strategies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("trace_jaeger_remote_sampling_strategies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this configv1 list trace jaeger remote sampling strategies response based on the context it is used
func (m *Configv1ListTraceJaegerRemoteSamplingStrategiesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTraceJaegerRemoteSamplingStrategies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Configv1ListTraceJaegerRemoteSamplingStrategiesResponse) contextValidatePage(ctx context.Context, formats strfmt.Registry) error {

	if m.Page != nil {
		if err := m.Page.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1ListTraceJaegerRemoteSamplingStrategiesResponse) contextValidateTraceJaegerRemoteSamplingStrategies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TraceJaegerRemoteSamplingStrategies); i++ {

		if m.TraceJaegerRemoteSamplingStrategies[i] != nil {
			if err := m.TraceJaegerRemoteSamplingStrategies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("trace_jaeger_remote_sampling_strategies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("trace_jaeger_remote_sampling_strategies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1ListTraceJaegerRemoteSamplingStrategiesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1ListTraceJaegerRemoteSamplingStrategiesResponse) UnmarshalBinary(b []byte) error {
	var res Configv1ListTraceJaegerRemoteSamplingStrategiesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
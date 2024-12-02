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

// SLIEndpointAvailabilityConfig Configuration for an endpoint availability SLI.
//
// swagger:model SLIEndpointAvailabilityConfig
type SLIEndpointAvailabilityConfig struct {

	// These are added to _every_ query and are intended to be used for things
	// like `cluster!~"dev"`
	AdditionalPromqlFilters []*ConfigunstablePromQLMatcher `json:"additional_promql_filters"`

	// The endpoints that are monitored by this SLI.
	EndpointsMonitored []string `json:"endpoints_monitored"`

	// A list of result codes that indicate an unsuccessful event. Either this
	// or success_codes must be set.
	ErrorCodes []string `json:"error_codes"`
}

// Validate validates this s l i endpoint availability config
func (m *SLIEndpointAvailabilityConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalPromqlFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SLIEndpointAvailabilityConfig) validateAdditionalPromqlFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalPromqlFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalPromqlFilters); i++ {
		if swag.IsZero(m.AdditionalPromqlFilters[i]) { // not required
			continue
		}

		if m.AdditionalPromqlFilters[i] != nil {
			if err := m.AdditionalPromqlFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additional_promql_filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additional_promql_filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this s l i endpoint availability config based on the context it is used
func (m *SLIEndpointAvailabilityConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalPromqlFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SLIEndpointAvailabilityConfig) contextValidateAdditionalPromqlFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalPromqlFilters); i++ {

		if m.AdditionalPromqlFilters[i] != nil {
			if err := m.AdditionalPromqlFilters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additional_promql_filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additional_promql_filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SLIEndpointAvailabilityConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SLIEndpointAvailabilityConfig) UnmarshalBinary(b []byte) error {
	var res SLIEndpointAvailabilityConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

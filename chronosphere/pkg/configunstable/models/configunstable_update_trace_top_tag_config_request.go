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

// ConfigunstableUpdateTraceTopTagConfigRequest configunstable update trace top tag config request
//
// swagger:model configunstableUpdateTraceTopTagConfigRequest
type ConfigunstableUpdateTraceTopTagConfigRequest struct {

	// If true, the TraceTopTagConfig will be created if it does not already exist. If false, an error will be returned if the TraceTopTagConfig does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the TraceTopTagConfig will not be created nor updated, and no response TraceTopTagConfig will be returned. The response will return an error if the given TraceTopTagConfig is invalid.
	DryRun bool `json:"dry_run,omitempty"`

	// trace top tag config
	TraceTopTagConfig *ConfigunstableTraceTopTagConfig `json:"trace_top_tag_config,omitempty"`
}

// Validate validates this configunstable update trace top tag config request
func (m *ConfigunstableUpdateTraceTopTagConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraceTopTagConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableUpdateTraceTopTagConfigRequest) validateTraceTopTagConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.TraceTopTagConfig) { // not required
		return nil
	}

	if m.TraceTopTagConfig != nil {
		if err := m.TraceTopTagConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_top_tag_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_top_tag_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable update trace top tag config request based on the context it is used
func (m *ConfigunstableUpdateTraceTopTagConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTraceTopTagConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableUpdateTraceTopTagConfigRequest) contextValidateTraceTopTagConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.TraceTopTagConfig != nil {
		if err := m.TraceTopTagConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trace_top_tag_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trace_top_tag_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableUpdateTraceTopTagConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableUpdateTraceTopTagConfigRequest) UnmarshalBinary(b []byte) error {
	var res ConfigunstableUpdateTraceTopTagConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

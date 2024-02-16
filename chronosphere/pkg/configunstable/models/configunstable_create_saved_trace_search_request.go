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

// ConfigunstableCreateSavedTraceSearchRequest configunstable create saved trace search request
//
// swagger:model configunstableCreateSavedTraceSearchRequest
type ConfigunstableCreateSavedTraceSearchRequest struct {

	// saved trace search
	SavedTraceSearch *ConfigunstableSavedTraceSearch `json:"saved_trace_search,omitempty"`
}

// Validate validates this configunstable create saved trace search request
func (m *ConfigunstableCreateSavedTraceSearchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSavedTraceSearch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateSavedTraceSearchRequest) validateSavedTraceSearch(formats strfmt.Registry) error {
	if swag.IsZero(m.SavedTraceSearch) { // not required
		return nil
	}

	if m.SavedTraceSearch != nil {
		if err := m.SavedTraceSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("saved_trace_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("saved_trace_search")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable create saved trace search request based on the context it is used
func (m *ConfigunstableCreateSavedTraceSearchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSavedTraceSearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableCreateSavedTraceSearchRequest) contextValidateSavedTraceSearch(ctx context.Context, formats strfmt.Registry) error {

	if m.SavedTraceSearch != nil {
		if err := m.SavedTraceSearch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("saved_trace_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("saved_trace_search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableCreateSavedTraceSearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableCreateSavedTraceSearchRequest) UnmarshalBinary(b []byte) error {
	var res ConfigunstableCreateSavedTraceSearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
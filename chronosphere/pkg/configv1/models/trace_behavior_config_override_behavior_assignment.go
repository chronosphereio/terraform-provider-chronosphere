// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TraceBehaviorConfigOverrideBehaviorAssignment trace behavior config override behavior assignment
//
// swagger:model TraceBehaviorConfigOverrideBehaviorAssignment
type TraceBehaviorConfigOverrideBehaviorAssignment struct {

	// The slug reference of a TraceBehavior
	BehaviorSlug string `json:"behavior_slug,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The author or creator of the entry.
	CreatedBy string `json:"created_by,omitempty"`

	// The slug reference of a TraceDataset
	DatasetSlug string `json:"dataset_slug,omitempty"`

	// A description of the entry.
	Description string `json:"description,omitempty"`

	// The ending time of the override.
	// Format: date-time
	EndTime strfmt.DateTime `json:"end_time,omitempty"`

	// The starting time of the override.
	// Format: date-time
	StartTime strfmt.DateTime `json:"start_time,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this trace behavior config override behavior assignment
func (m *TraceBehaviorConfigOverrideBehaviorAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceBehaviorConfigOverrideBehaviorAssignment) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TraceBehaviorConfigOverrideBehaviorAssignment) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TraceBehaviorConfigOverrideBehaviorAssignment) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TraceBehaviorConfigOverrideBehaviorAssignment) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this trace behavior config override behavior assignment based on context it is used
func (m *TraceBehaviorConfigOverrideBehaviorAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TraceBehaviorConfigOverrideBehaviorAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraceBehaviorConfigOverrideBehaviorAssignment) UnmarshalBinary(b []byte) error {
	var res TraceBehaviorConfigOverrideBehaviorAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

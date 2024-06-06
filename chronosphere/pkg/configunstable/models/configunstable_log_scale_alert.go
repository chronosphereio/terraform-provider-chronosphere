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

// ConfigunstableLogScaleAlert configunstable log scale alert
//
// swagger:model configunstableLogScaleAlert
type ConfigunstableLogScaleAlert struct {

	// Timestamp of when the LogScaleAlert was created. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// Optional. When value is empty this alert will not trigger anything.
	LogScaleActionSlugs []string `json:"log_scale_action_slugs"`

	// log scale query
	LogScaleQuery string `json:"log_scale_query,omitempty"`

	// Required name of the LogScaleAlert. May be modified after the LogScaleAlert is created.
	Name string `json:"name,omitempty"`

	// repository
	Repository string `json:"repository,omitempty"`

	// Unique identifier of the LogScaleAlert. If slug is not provided, one will be generated based of the name field. Cannot be modified after the LogScaleAlert is created.
	Slug string `json:"slug,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// The alert is triggered at most once per throttle period.
	ThrottleSecs int32 `json:"throttle_secs,omitempty"`

	// An alert uses a sliding time window for its search.
	// If this is set to 86400 seconds (24 hours), only the events from the last 24 hours will be considered when the alert query is run.
	TimeWindowSecs int32 `json:"time_window_secs,omitempty"`

	// Timestamp of when the LogScaleAlert was last updated. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this configunstable log scale alert
func (m *ConfigunstableLogScaleAlert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
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

func (m *ConfigunstableLogScaleAlert) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableLogScaleAlert) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this configunstable log scale alert based on the context it is used
func (m *ConfigunstableLogScaleAlert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableLogScaleAlert) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableLogScaleAlert) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableLogScaleAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableLogScaleAlert) UnmarshalBinary(b []byte) error {
	var res ConfigunstableLogScaleAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

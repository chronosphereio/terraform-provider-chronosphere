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

// Configv1RecordingRule configv1 recording rule
//
// swagger:model configv1RecordingRule
type Configv1RecordingRule struct {

	// Optional slug of the bucket the RecordingRule belongs to.
	BucketSlug string `json:"bucket_slug,omitempty"`

	// Timestamp of when the RecordingRule was created. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Optional execution_group in which this rule is to be evaluated.
	// At least one of bucket_slug and execution_group must be set. If both are set, then they are expected to match.
	ExecutionGroup string `json:"execution_group,omitempty"`

	// Optional interval for evaluating the recording rule.
	IntervalSecs int32 `json:"interval_secs,omitempty"`

	// label policy
	LabelPolicy *Configv1RecordingRuleLabelPolicy `json:"label_policy,omitempty"`

	// The name of the time series to use for output, which must be a valid
	// metric name.
	MetricName string `json:"metric_name,omitempty"`

	// Required. Name of the RecordingRule. You can modify this value after the RecordingRule is created.
	Name string `json:"name,omitempty"`

	// The PromQL expression to evaluate at the time of each evaluation cycle.
	// The result is recorded as a new time series with its metric name
	// defined by the metric_name (or name) field.
	PrometheusExpr string `json:"prometheus_expr,omitempty"`

	// Unique identifier of the RecordingRule. If a `slug` isn't provided, one will be generated based of the `name` field. You can't modify this field after the RecordingRule is created.
	Slug string `json:"slug,omitempty"`

	// Timestamp of when the RecordingRule was last updated. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this configv1 recording rule
func (m *Configv1RecordingRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelPolicy(formats); err != nil {
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

func (m *Configv1RecordingRule) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Configv1RecordingRule) validateLabelPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelPolicy) { // not required
		return nil
	}

	if m.LabelPolicy != nil {
		if err := m.LabelPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("label_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("label_policy")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1RecordingRule) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this configv1 recording rule based on the context it is used
func (m *Configv1RecordingRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelPolicy(ctx, formats); err != nil {
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

func (m *Configv1RecordingRule) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Configv1RecordingRule) contextValidateLabelPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelPolicy != nil {
		if err := m.LabelPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("label_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("label_policy")
			}
			return err
		}
	}

	return nil
}

func (m *Configv1RecordingRule) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Configv1RecordingRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Configv1RecordingRule) UnmarshalBinary(b []byte) error {
	var res Configv1RecordingRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

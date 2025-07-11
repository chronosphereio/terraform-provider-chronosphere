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

// ConfigV1UpdateNotificationPolicyBody config v1 update notification policy body
//
// swagger:model ConfigV1UpdateNotificationPolicyBody
type ConfigV1UpdateNotificationPolicyBody struct {

	// If true, the NotificationPolicy will be created if it does not already exist, identified by slug. If false, an error will be returned if the NotificationPolicy does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the NotificationPolicy isn't created or updated, and no response NotificationPolicy will be returned. The response will return an error if the given NotificationPolicy is invalid.
	DryRun bool `json:"dry_run,omitempty"`

	// notification policy
	NotificationPolicy *Configv1NotificationPolicy `json:"notification_policy,omitempty"`
}

// Validate validates this config v1 update notification policy body
func (m *ConfigV1UpdateNotificationPolicyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotificationPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateNotificationPolicyBody) validateNotificationPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.NotificationPolicy) { // not required
		return nil
	}

	if m.NotificationPolicy != nil {
		if err := m.NotificationPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notification_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notification_policy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config v1 update notification policy body based on the context it is used
func (m *ConfigV1UpdateNotificationPolicyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNotificationPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigV1UpdateNotificationPolicyBody) contextValidateNotificationPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.NotificationPolicy != nil {
		if err := m.NotificationPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notification_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notification_policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigV1UpdateNotificationPolicyBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigV1UpdateNotificationPolicyBody) UnmarshalBinary(b []byte) error {
	var res ConfigV1UpdateNotificationPolicyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

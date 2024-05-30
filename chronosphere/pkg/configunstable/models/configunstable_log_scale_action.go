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

// ConfigunstableLogScaleAction configunstable log scale action
//
// swagger:model configunstableLogScaleAction
type ConfigunstableLogScaleAction struct {

	// action type
	ActionType LogScaleActionActionType `json:"action_type,omitempty"`

	// Timestamp of when the LogScaleAction was created. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// email action
	EmailAction *LogScaleActionEmailAction `json:"email_action,omitempty"`

	// humio action
	HumioAction *LogScaleActionHumioRepoAction `json:"humio_action,omitempty"`

	// Required name of the LogScaleAction. May be modified after the LogScaleAction is created.
	Name string `json:"name,omitempty"`

	// ops genie action
	OpsGenieAction *LogScaleActionOpsGenieAction `json:"ops_genie_action,omitempty"`

	// pager duty action
	PagerDutyAction *LogScaleActionPagerDutyAction `json:"pager_duty_action,omitempty"`

	// repository
	Repository string `json:"repository,omitempty"`

	// slack action
	SlackAction *LogScaleActionSlackAction `json:"slack_action,omitempty"`

	// slack post message action
	SlackPostMessageAction *LogScaleActionSlackPostMessageAction `json:"slack_post_message_action,omitempty"`

	// Unique identifier of the LogScaleAction. If slug is not provided, one will be generated based of the name field. Cannot be modified after the LogScaleAction is created.
	Slug string `json:"slug,omitempty"`

	// Timestamp of when the LogScaleAction was last updated. Cannot be set by clients.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// upload file action
	UploadFileAction *LogScaleActionUploadFileAction `json:"upload_file_action,omitempty"`

	// victor ops action
	VictorOpsAction *LogScaleActionVictorOpsAction `json:"victor_ops_action,omitempty"`

	// webhook action
	WebhookAction *LogScaleActionWebhookAction `json:"webhook_action,omitempty"`
}

// Validate validates this configunstable log scale action
func (m *ConfigunstableLogScaleAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHumioAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpsGenieAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagerDutyAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlackAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlackPostMessageAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadFileAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVictorOpsAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableLogScaleAction) validateActionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionType) { // not required
		return nil
	}

	if err := m.ActionType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("action_type")
		}
		return err
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) validateEmailAction(formats strfmt.Registry) error {
	if swag.IsZero(m.EmailAction) { // not required
		return nil
	}

	if m.EmailAction != nil {
		if err := m.EmailAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("email_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("email_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) validateHumioAction(formats strfmt.Registry) error {
	if swag.IsZero(m.HumioAction) { // not required
		return nil
	}

	if m.HumioAction != nil {
		if err := m.HumioAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("humio_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("humio_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) validateOpsGenieAction(formats strfmt.Registry) error {
	if swag.IsZero(m.OpsGenieAction) { // not required
		return nil
	}

	if m.OpsGenieAction != nil {
		if err := m.OpsGenieAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ops_genie_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ops_genie_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) validatePagerDutyAction(formats strfmt.Registry) error {
	if swag.IsZero(m.PagerDutyAction) { // not required
		return nil
	}

	if m.PagerDutyAction != nil {
		if err := m.PagerDutyAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pager_duty_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pager_duty_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) validateSlackAction(formats strfmt.Registry) error {
	if swag.IsZero(m.SlackAction) { // not required
		return nil
	}

	if m.SlackAction != nil {
		if err := m.SlackAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slack_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slack_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) validateSlackPostMessageAction(formats strfmt.Registry) error {
	if swag.IsZero(m.SlackPostMessageAction) { // not required
		return nil
	}

	if m.SlackPostMessageAction != nil {
		if err := m.SlackPostMessageAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slack_post_message_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slack_post_message_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) validateUploadFileAction(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadFileAction) { // not required
		return nil
	}

	if m.UploadFileAction != nil {
		if err := m.UploadFileAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upload_file_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upload_file_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) validateVictorOpsAction(formats strfmt.Registry) error {
	if swag.IsZero(m.VictorOpsAction) { // not required
		return nil
	}

	if m.VictorOpsAction != nil {
		if err := m.VictorOpsAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("victor_ops_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("victor_ops_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) validateWebhookAction(formats strfmt.Registry) error {
	if swag.IsZero(m.WebhookAction) { // not required
		return nil
	}

	if m.WebhookAction != nil {
		if err := m.WebhookAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_action")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable log scale action based on the context it is used
func (m *ConfigunstableLogScaleAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmailAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHumioAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpsGenieAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagerDutyAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlackAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlackPostMessageAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUploadFileAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVictorOpsAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebhookAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableLogScaleAction) contextValidateActionType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ActionType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("action_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("action_type")
		}
		return err
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) contextValidateEmailAction(ctx context.Context, formats strfmt.Registry) error {

	if m.EmailAction != nil {
		if err := m.EmailAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("email_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("email_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) contextValidateHumioAction(ctx context.Context, formats strfmt.Registry) error {

	if m.HumioAction != nil {
		if err := m.HumioAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("humio_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("humio_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) contextValidateOpsGenieAction(ctx context.Context, formats strfmt.Registry) error {

	if m.OpsGenieAction != nil {
		if err := m.OpsGenieAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ops_genie_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ops_genie_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) contextValidatePagerDutyAction(ctx context.Context, formats strfmt.Registry) error {

	if m.PagerDutyAction != nil {
		if err := m.PagerDutyAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pager_duty_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pager_duty_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) contextValidateSlackAction(ctx context.Context, formats strfmt.Registry) error {

	if m.SlackAction != nil {
		if err := m.SlackAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slack_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slack_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) contextValidateSlackPostMessageAction(ctx context.Context, formats strfmt.Registry) error {

	if m.SlackPostMessageAction != nil {
		if err := m.SlackPostMessageAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slack_post_message_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slack_post_message_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) contextValidateUploadFileAction(ctx context.Context, formats strfmt.Registry) error {

	if m.UploadFileAction != nil {
		if err := m.UploadFileAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upload_file_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("upload_file_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) contextValidateVictorOpsAction(ctx context.Context, formats strfmt.Registry) error {

	if m.VictorOpsAction != nil {
		if err := m.VictorOpsAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("victor_ops_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("victor_ops_action")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigunstableLogScaleAction) contextValidateWebhookAction(ctx context.Context, formats strfmt.Registry) error {

	if m.WebhookAction != nil {
		if err := m.WebhookAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_action")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableLogScaleAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableLogScaleAction) UnmarshalBinary(b []byte) error {
	var res ConfigunstableLogScaleAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// TraceSearchFilterSpanFilter trace search filter span filter
//
// swagger:model TraceSearchFilterSpanFilter
type TraceSearchFilterSpanFilter struct {

	// duration
	Duration *TraceSearchFilterDurationFilter `json:"duration,omitempty"`

	// error
	Error *TraceSearchFilterBoolFilter `json:"error,omitempty"`

	// match type
	MatchType SpanFilterSpanFilterMatchType `json:"match_type,omitempty"`

	// operation
	Operation *TraceSearchFilterStringFilter `json:"operation,omitempty"`

	// parent operation
	ParentOperation *TraceSearchFilterStringFilter `json:"parent_operation,omitempty"`

	// parent service
	ParentService *TraceSearchFilterStringFilter `json:"parent_service,omitempty"`

	// service
	Service *TraceSearchFilterStringFilter `json:"service,omitempty"`

	// span count
	SpanCount *TraceSearchFilterCountFilter `json:"span_count,omitempty"`

	// Matches the tags of the candidate.
	Tags []*TraceSearchFilterTagFilter `json:"tags"`
}

// Validate validates this trace search filter span filter
func (m *TraceSearchFilterSpanFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpanCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceSearchFilterSpanFilter) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	if m.Duration != nil {
		if err := m.Duration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("duration")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) validateMatchType(formats strfmt.Registry) error {
	if swag.IsZero(m.MatchType) { // not required
		return nil
	}

	if err := m.MatchType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("match_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("match_type")
		}
		return err
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) validateOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) validateParentOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentOperation) { // not required
		return nil
	}

	if m.ParentOperation != nil {
		if err := m.ParentOperation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent_operation")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) validateParentService(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentService) { // not required
		return nil
	}

	if m.ParentService != nil {
		if err := m.ParentService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent_service")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) validateSpanCount(formats strfmt.Registry) error {
	if swag.IsZero(m.SpanCount) { // not required
		return nil
	}

	if m.SpanCount != nil {
		if err := m.SpanCount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("span_count")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("span_count")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this trace search filter span filter based on the context it is used
func (m *TraceSearchFilterSpanFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMatchType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpanCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TraceSearchFilterSpanFilter) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if m.Duration != nil {
		if err := m.Duration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("duration")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {
		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) contextValidateMatchType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MatchType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("match_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("match_type")
		}
		return err
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) contextValidateOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.Operation != nil {
		if err := m.Operation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) contextValidateParentOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentOperation != nil {
		if err := m.ParentOperation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent_operation")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) contextValidateParentService(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentService != nil {
		if err := m.ParentService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent_service")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

	if m.Service != nil {
		if err := m.Service.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) contextValidateSpanCount(ctx context.Context, formats strfmt.Registry) error {

	if m.SpanCount != nil {
		if err := m.SpanCount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("span_count")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("span_count")
			}
			return err
		}
	}

	return nil
}

func (m *TraceSearchFilterSpanFilter) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TraceSearchFilterSpanFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TraceSearchFilterSpanFilter) UnmarshalBinary(b []byte) error {
	var res TraceSearchFilterSpanFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

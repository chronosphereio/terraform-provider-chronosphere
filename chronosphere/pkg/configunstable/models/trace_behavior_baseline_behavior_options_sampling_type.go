// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// TraceBehaviorBaselineBehaviorOptionsSamplingType  - LOW_VALUE: Match indicates a low value trace. With multiple low value matches sample at the lowest rate.
//   - HIGH_VALUE: Match indicates a high value trace. With multiple high value matches sample at the highest rate.
//
// swagger:model TraceBehaviorBaselineBehaviorOptionsSamplingType
type TraceBehaviorBaselineBehaviorOptionsSamplingType string

func NewTraceBehaviorBaselineBehaviorOptionsSamplingType(value TraceBehaviorBaselineBehaviorOptionsSamplingType) *TraceBehaviorBaselineBehaviorOptionsSamplingType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated TraceBehaviorBaselineBehaviorOptionsSamplingType.
func (m TraceBehaviorBaselineBehaviorOptionsSamplingType) Pointer() *TraceBehaviorBaselineBehaviorOptionsSamplingType {
	return &m
}

const (

	// TraceBehaviorBaselineBehaviorOptionsSamplingTypeLOWVALUE captures enum value "LOW_VALUE"
	TraceBehaviorBaselineBehaviorOptionsSamplingTypeLOWVALUE TraceBehaviorBaselineBehaviorOptionsSamplingType = "LOW_VALUE"

	// TraceBehaviorBaselineBehaviorOptionsSamplingTypeHIGHVALUE captures enum value "HIGH_VALUE"
	TraceBehaviorBaselineBehaviorOptionsSamplingTypeHIGHVALUE TraceBehaviorBaselineBehaviorOptionsSamplingType = "HIGH_VALUE"
)

// for schema
var traceBehaviorBaselineBehaviorOptionsSamplingTypeEnum []interface{}

func init() {
	var res []TraceBehaviorBaselineBehaviorOptionsSamplingType
	if err := json.Unmarshal([]byte(`["LOW_VALUE","HIGH_VALUE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		traceBehaviorBaselineBehaviorOptionsSamplingTypeEnum = append(traceBehaviorBaselineBehaviorOptionsSamplingTypeEnum, v)
	}
}

func (m TraceBehaviorBaselineBehaviorOptionsSamplingType) validateTraceBehaviorBaselineBehaviorOptionsSamplingTypeEnum(path, location string, value TraceBehaviorBaselineBehaviorOptionsSamplingType) error {
	if err := validate.EnumCase(path, location, value, traceBehaviorBaselineBehaviorOptionsSamplingTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this trace behavior baseline behavior options sampling type
func (m TraceBehaviorBaselineBehaviorOptionsSamplingType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTraceBehaviorBaselineBehaviorOptionsSamplingTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this trace behavior baseline behavior options sampling type based on context it is used
func (m TraceBehaviorBaselineBehaviorOptionsSamplingType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

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

// Configv1MutingRuleLabelMatcherMatcherType configv1 muting rule label matcher matcher type
//
// swagger:model configv1MutingRuleLabelMatcherMatcherType
type Configv1MutingRuleLabelMatcherMatcherType string

func NewConfigv1MutingRuleLabelMatcherMatcherType(value Configv1MutingRuleLabelMatcherMatcherType) *Configv1MutingRuleLabelMatcherMatcherType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Configv1MutingRuleLabelMatcherMatcherType.
func (m Configv1MutingRuleLabelMatcherMatcherType) Pointer() *Configv1MutingRuleLabelMatcherMatcherType {
	return &m
}

const (

	// Configv1MutingRuleLabelMatcherMatcherTypeEXACT captures enum value "EXACT"
	Configv1MutingRuleLabelMatcherMatcherTypeEXACT Configv1MutingRuleLabelMatcherMatcherType = "EXACT"

	// Configv1MutingRuleLabelMatcherMatcherTypeREGEX captures enum value "REGEX"
	Configv1MutingRuleLabelMatcherMatcherTypeREGEX Configv1MutingRuleLabelMatcherMatcherType = "REGEX"

	// Configv1MutingRuleLabelMatcherMatcherTypeNOTEXACT captures enum value "NOT_EXACT"
	Configv1MutingRuleLabelMatcherMatcherTypeNOTEXACT Configv1MutingRuleLabelMatcherMatcherType = "NOT_EXACT"

	// Configv1MutingRuleLabelMatcherMatcherTypeNOTREGEXP captures enum value "NOT_REGEXP"
	Configv1MutingRuleLabelMatcherMatcherTypeNOTREGEXP Configv1MutingRuleLabelMatcherMatcherType = "NOT_REGEXP"
)

// for schema
var configv1MutingRuleLabelMatcherMatcherTypeEnum []interface{}

func init() {
	var res []Configv1MutingRuleLabelMatcherMatcherType
	if err := json.Unmarshal([]byte(`["EXACT","REGEX","NOT_EXACT","NOT_REGEXP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configv1MutingRuleLabelMatcherMatcherTypeEnum = append(configv1MutingRuleLabelMatcherMatcherTypeEnum, v)
	}
}

func (m Configv1MutingRuleLabelMatcherMatcherType) validateConfigv1MutingRuleLabelMatcherMatcherTypeEnum(path, location string, value Configv1MutingRuleLabelMatcherMatcherType) error {
	if err := validate.EnumCase(path, location, value, configv1MutingRuleLabelMatcherMatcherTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this configv1 muting rule label matcher matcher type
func (m Configv1MutingRuleLabelMatcherMatcherType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigv1MutingRuleLabelMatcherMatcherTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this configv1 muting rule label matcher matcher type based on context it is used
func (m Configv1MutingRuleLabelMatcherMatcherType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
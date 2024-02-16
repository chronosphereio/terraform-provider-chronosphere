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

// Configv1CollectionReferenceType Type values must match entitiespb.Collection.CollectionType.
//
// swagger:model configv1CollectionReferenceType
type Configv1CollectionReferenceType string

func NewConfigv1CollectionReferenceType(value Configv1CollectionReferenceType) *Configv1CollectionReferenceType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Configv1CollectionReferenceType.
func (m Configv1CollectionReferenceType) Pointer() *Configv1CollectionReferenceType {
	return &m
}

const (

	// Configv1CollectionReferenceTypeSIMPLE captures enum value "SIMPLE"
	Configv1CollectionReferenceTypeSIMPLE Configv1CollectionReferenceType = "SIMPLE"

	// Configv1CollectionReferenceTypeSERVICE captures enum value "SERVICE"
	Configv1CollectionReferenceTypeSERVICE Configv1CollectionReferenceType = "SERVICE"
)

// for schema
var configv1CollectionReferenceTypeEnum []interface{}

func init() {
	var res []Configv1CollectionReferenceType
	if err := json.Unmarshal([]byte(`["SIMPLE","SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configv1CollectionReferenceTypeEnum = append(configv1CollectionReferenceTypeEnum, v)
	}
}

func (m Configv1CollectionReferenceType) validateConfigv1CollectionReferenceTypeEnum(path, location string, value Configv1CollectionReferenceType) error {
	if err := validate.EnumCase(path, location, value, configv1CollectionReferenceTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this configv1 collection reference type
func (m Configv1CollectionReferenceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigv1CollectionReferenceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this configv1 collection reference type based on context it is used
func (m Configv1CollectionReferenceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
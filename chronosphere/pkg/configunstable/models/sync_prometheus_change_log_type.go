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

// SyncPrometheusChangeLogType sync prometheus change log type
//
// swagger:model SyncPrometheusChangeLogType
type SyncPrometheusChangeLogType string

func NewSyncPrometheusChangeLogType(value SyncPrometheusChangeLogType) *SyncPrometheusChangeLogType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SyncPrometheusChangeLogType.
func (m SyncPrometheusChangeLogType) Pointer() *SyncPrometheusChangeLogType {
	return &m
}

const (

	// SyncPrometheusChangeLogTypeCREATEDNOTIFIER captures enum value "CREATED_NOTIFIER"
	SyncPrometheusChangeLogTypeCREATEDNOTIFIER SyncPrometheusChangeLogType = "CREATED_NOTIFIER"

	// SyncPrometheusChangeLogTypeUPDATEDNOTIFIER captures enum value "UPDATED_NOTIFIER"
	SyncPrometheusChangeLogTypeUPDATEDNOTIFIER SyncPrometheusChangeLogType = "UPDATED_NOTIFIER"

	// SyncPrometheusChangeLogTypeDELETEDNOTIFIER captures enum value "DELETED_NOTIFIER"
	SyncPrometheusChangeLogTypeDELETEDNOTIFIER SyncPrometheusChangeLogType = "DELETED_NOTIFIER"

	// SyncPrometheusChangeLogTypeCREATEDBUCKET captures enum value "CREATED_BUCKET"
	SyncPrometheusChangeLogTypeCREATEDBUCKET SyncPrometheusChangeLogType = "CREATED_BUCKET"

	// SyncPrometheusChangeLogTypeUPDATEDBUCKET captures enum value "UPDATED_BUCKET"
	SyncPrometheusChangeLogTypeUPDATEDBUCKET SyncPrometheusChangeLogType = "UPDATED_BUCKET"

	// SyncPrometheusChangeLogTypeDELETEDBUCKET captures enum value "DELETED_BUCKET"
	SyncPrometheusChangeLogTypeDELETEDBUCKET SyncPrometheusChangeLogType = "DELETED_BUCKET"

	// SyncPrometheusChangeLogTypeCREATEDCOLLECTION captures enum value "CREATED_COLLECTION"
	SyncPrometheusChangeLogTypeCREATEDCOLLECTION SyncPrometheusChangeLogType = "CREATED_COLLECTION"

	// SyncPrometheusChangeLogTypeUPDATEDCOLLECTION captures enum value "UPDATED_COLLECTION"
	SyncPrometheusChangeLogTypeUPDATEDCOLLECTION SyncPrometheusChangeLogType = "UPDATED_COLLECTION"

	// SyncPrometheusChangeLogTypeDELETEDCOLLECTION captures enum value "DELETED_COLLECTION"
	SyncPrometheusChangeLogTypeDELETEDCOLLECTION SyncPrometheusChangeLogType = "DELETED_COLLECTION"

	// SyncPrometheusChangeLogTypeCREATEDNOTIFICATIONPOLICY captures enum value "CREATED_NOTIFICATION_POLICY"
	SyncPrometheusChangeLogTypeCREATEDNOTIFICATIONPOLICY SyncPrometheusChangeLogType = "CREATED_NOTIFICATION_POLICY"

	// SyncPrometheusChangeLogTypeUPDATEDNOTIFICATIONPOLICY captures enum value "UPDATED_NOTIFICATION_POLICY"
	SyncPrometheusChangeLogTypeUPDATEDNOTIFICATIONPOLICY SyncPrometheusChangeLogType = "UPDATED_NOTIFICATION_POLICY"

	// SyncPrometheusChangeLogTypeDELETEDNOTIFICATIONPOLICY captures enum value "DELETED_NOTIFICATION_POLICY"
	SyncPrometheusChangeLogTypeDELETEDNOTIFICATIONPOLICY SyncPrometheusChangeLogType = "DELETED_NOTIFICATION_POLICY"

	// SyncPrometheusChangeLogTypeCREATEDMONITOR captures enum value "CREATED_MONITOR"
	SyncPrometheusChangeLogTypeCREATEDMONITOR SyncPrometheusChangeLogType = "CREATED_MONITOR"

	// SyncPrometheusChangeLogTypeUPDATEDMONITOR captures enum value "UPDATED_MONITOR"
	SyncPrometheusChangeLogTypeUPDATEDMONITOR SyncPrometheusChangeLogType = "UPDATED_MONITOR"

	// SyncPrometheusChangeLogTypeDELETEDMONITOR captures enum value "DELETED_MONITOR"
	SyncPrometheusChangeLogTypeDELETEDMONITOR SyncPrometheusChangeLogType = "DELETED_MONITOR"

	// SyncPrometheusChangeLogTypeCREATEDRECORDINGRULE captures enum value "CREATED_RECORDING_RULE"
	SyncPrometheusChangeLogTypeCREATEDRECORDINGRULE SyncPrometheusChangeLogType = "CREATED_RECORDING_RULE"

	// SyncPrometheusChangeLogTypeUPDATEDRECORDINGRULE captures enum value "UPDATED_RECORDING_RULE"
	SyncPrometheusChangeLogTypeUPDATEDRECORDINGRULE SyncPrometheusChangeLogType = "UPDATED_RECORDING_RULE"

	// SyncPrometheusChangeLogTypeDELETEDRECORDINGRULE captures enum value "DELETED_RECORDING_RULE"
	SyncPrometheusChangeLogTypeDELETEDRECORDINGRULE SyncPrometheusChangeLogType = "DELETED_RECORDING_RULE"
)

// for schema
var syncPrometheusChangeLogTypeEnum []interface{}

func init() {
	var res []SyncPrometheusChangeLogType
	if err := json.Unmarshal([]byte(`["CREATED_NOTIFIER","UPDATED_NOTIFIER","DELETED_NOTIFIER","CREATED_BUCKET","UPDATED_BUCKET","DELETED_BUCKET","CREATED_COLLECTION","UPDATED_COLLECTION","DELETED_COLLECTION","CREATED_NOTIFICATION_POLICY","UPDATED_NOTIFICATION_POLICY","DELETED_NOTIFICATION_POLICY","CREATED_MONITOR","UPDATED_MONITOR","DELETED_MONITOR","CREATED_RECORDING_RULE","UPDATED_RECORDING_RULE","DELETED_RECORDING_RULE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		syncPrometheusChangeLogTypeEnum = append(syncPrometheusChangeLogTypeEnum, v)
	}
}

func (m SyncPrometheusChangeLogType) validateSyncPrometheusChangeLogTypeEnum(path, location string, value SyncPrometheusChangeLogType) error {
	if err := validate.EnumCase(path, location, value, syncPrometheusChangeLogTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this sync prometheus change log type
func (m SyncPrometheusChangeLogType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSyncPrometheusChangeLogTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this sync prometheus change log type based on context it is used
func (m SyncPrometheusChangeLogType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

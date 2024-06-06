// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LogScaleActionHumioRepoAction log scale action humio repo action
//
// swagger:model LogScaleActionHumioRepoAction
type LogScaleActionHumioRepoAction struct {

	// ingest token
	IngestToken string `json:"ingest_token,omitempty"`
}

// Validate validates this log scale action humio repo action
func (m *LogScaleActionHumioRepoAction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this log scale action humio repo action based on context it is used
func (m *LogScaleActionHumioRepoAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogScaleActionHumioRepoAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogScaleActionHumioRepoAction) UnmarshalBinary(b []byte) error {
	var res LogScaleActionHumioRepoAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
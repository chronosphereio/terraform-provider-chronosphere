// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmitMetricshistogramMetric emit metricshistogram metric
//
// swagger:model EmitMetricshistogramMetric
type EmitMetricshistogramMetric struct {

	// Buckets specify the buckets to use for the histogram.
	Buckets []float64 `json:"buckets"`

	// Value field is used to specify what field holds the numerical value.
	ValueField string `json:"value_field,omitempty"`
}

// Validate validates this emit metricshistogram metric
func (m *EmitMetricshistogramMetric) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this emit metricshistogram metric based on context it is used
func (m *EmitMetricshistogramMetric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmitMetricshistogramMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmitMetricshistogramMetric) UnmarshalBinary(b []byte) error {
	var res EmitMetricshistogramMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package otel_metrics_ingestion

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// NewUpdateOtelMetricsIngestionParams creates a new UpdateOtelMetricsIngestionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOtelMetricsIngestionParams() *UpdateOtelMetricsIngestionParams {
	return &UpdateOtelMetricsIngestionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOtelMetricsIngestionParamsWithTimeout creates a new UpdateOtelMetricsIngestionParams object
// with the ability to set a timeout on a request.
func NewUpdateOtelMetricsIngestionParamsWithTimeout(timeout time.Duration) *UpdateOtelMetricsIngestionParams {
	return &UpdateOtelMetricsIngestionParams{
		timeout: timeout,
	}
}

// NewUpdateOtelMetricsIngestionParamsWithContext creates a new UpdateOtelMetricsIngestionParams object
// with the ability to set a context for a request.
func NewUpdateOtelMetricsIngestionParamsWithContext(ctx context.Context) *UpdateOtelMetricsIngestionParams {
	return &UpdateOtelMetricsIngestionParams{
		Context: ctx,
	}
}

// NewUpdateOtelMetricsIngestionParamsWithHTTPClient creates a new UpdateOtelMetricsIngestionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOtelMetricsIngestionParamsWithHTTPClient(client *http.Client) *UpdateOtelMetricsIngestionParams {
	return &UpdateOtelMetricsIngestionParams{
		HTTPClient: client,
	}
}

/*
UpdateOtelMetricsIngestionParams contains all the parameters to send to the API endpoint

	for the update otel metrics ingestion operation.

	Typically these are written to a http.Request.
*/
type UpdateOtelMetricsIngestionParams struct {

	// Body.
	Body *models.Configv1UpdateOtelMetricsIngestionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update otel metrics ingestion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOtelMetricsIngestionParams) WithDefaults() *UpdateOtelMetricsIngestionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update otel metrics ingestion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOtelMetricsIngestionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update otel metrics ingestion params
func (o *UpdateOtelMetricsIngestionParams) WithTimeout(timeout time.Duration) *UpdateOtelMetricsIngestionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update otel metrics ingestion params
func (o *UpdateOtelMetricsIngestionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update otel metrics ingestion params
func (o *UpdateOtelMetricsIngestionParams) WithContext(ctx context.Context) *UpdateOtelMetricsIngestionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update otel metrics ingestion params
func (o *UpdateOtelMetricsIngestionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update otel metrics ingestion params
func (o *UpdateOtelMetricsIngestionParams) WithHTTPClient(client *http.Client) *UpdateOtelMetricsIngestionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update otel metrics ingestion params
func (o *UpdateOtelMetricsIngestionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update otel metrics ingestion params
func (o *UpdateOtelMetricsIngestionParams) WithBody(body *models.Configv1UpdateOtelMetricsIngestionRequest) *UpdateOtelMetricsIngestionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update otel metrics ingestion params
func (o *UpdateOtelMetricsIngestionParams) SetBody(body *models.Configv1UpdateOtelMetricsIngestionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOtelMetricsIngestionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
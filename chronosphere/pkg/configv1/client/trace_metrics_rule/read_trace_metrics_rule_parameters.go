// Code generated by go-swagger; DO NOT EDIT.

package trace_metrics_rule

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
)

// NewReadTraceMetricsRuleParams creates a new ReadTraceMetricsRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadTraceMetricsRuleParams() *ReadTraceMetricsRuleParams {
	return &ReadTraceMetricsRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadTraceMetricsRuleParamsWithTimeout creates a new ReadTraceMetricsRuleParams object
// with the ability to set a timeout on a request.
func NewReadTraceMetricsRuleParamsWithTimeout(timeout time.Duration) *ReadTraceMetricsRuleParams {
	return &ReadTraceMetricsRuleParams{
		timeout: timeout,
	}
}

// NewReadTraceMetricsRuleParamsWithContext creates a new ReadTraceMetricsRuleParams object
// with the ability to set a context for a request.
func NewReadTraceMetricsRuleParamsWithContext(ctx context.Context) *ReadTraceMetricsRuleParams {
	return &ReadTraceMetricsRuleParams{
		Context: ctx,
	}
}

// NewReadTraceMetricsRuleParamsWithHTTPClient creates a new ReadTraceMetricsRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadTraceMetricsRuleParamsWithHTTPClient(client *http.Client) *ReadTraceMetricsRuleParams {
	return &ReadTraceMetricsRuleParams{
		HTTPClient: client,
	}
}

/*
ReadTraceMetricsRuleParams contains all the parameters to send to the API endpoint

	for the read trace metrics rule operation.

	Typically these are written to a http.Request.
*/
type ReadTraceMetricsRuleParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read trace metrics rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadTraceMetricsRuleParams) WithDefaults() *ReadTraceMetricsRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read trace metrics rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadTraceMetricsRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read trace metrics rule params
func (o *ReadTraceMetricsRuleParams) WithTimeout(timeout time.Duration) *ReadTraceMetricsRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read trace metrics rule params
func (o *ReadTraceMetricsRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read trace metrics rule params
func (o *ReadTraceMetricsRuleParams) WithContext(ctx context.Context) *ReadTraceMetricsRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read trace metrics rule params
func (o *ReadTraceMetricsRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read trace metrics rule params
func (o *ReadTraceMetricsRuleParams) WithHTTPClient(client *http.Client) *ReadTraceMetricsRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read trace metrics rule params
func (o *ReadTraceMetricsRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the read trace metrics rule params
func (o *ReadTraceMetricsRuleParams) WithSlug(slug string) *ReadTraceMetricsRuleParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the read trace metrics rule params
func (o *ReadTraceMetricsRuleParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ReadTraceMetricsRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param slug
	if err := r.SetPathParam("slug", o.Slug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

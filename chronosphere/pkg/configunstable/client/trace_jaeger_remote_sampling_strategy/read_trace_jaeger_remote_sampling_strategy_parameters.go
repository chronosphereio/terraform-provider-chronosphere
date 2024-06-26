// Code generated by go-swagger; DO NOT EDIT.

package trace_jaeger_remote_sampling_strategy

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

// NewReadTraceJaegerRemoteSamplingStrategyParams creates a new ReadTraceJaegerRemoteSamplingStrategyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadTraceJaegerRemoteSamplingStrategyParams() *ReadTraceJaegerRemoteSamplingStrategyParams {
	return &ReadTraceJaegerRemoteSamplingStrategyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadTraceJaegerRemoteSamplingStrategyParamsWithTimeout creates a new ReadTraceJaegerRemoteSamplingStrategyParams object
// with the ability to set a timeout on a request.
func NewReadTraceJaegerRemoteSamplingStrategyParamsWithTimeout(timeout time.Duration) *ReadTraceJaegerRemoteSamplingStrategyParams {
	return &ReadTraceJaegerRemoteSamplingStrategyParams{
		timeout: timeout,
	}
}

// NewReadTraceJaegerRemoteSamplingStrategyParamsWithContext creates a new ReadTraceJaegerRemoteSamplingStrategyParams object
// with the ability to set a context for a request.
func NewReadTraceJaegerRemoteSamplingStrategyParamsWithContext(ctx context.Context) *ReadTraceJaegerRemoteSamplingStrategyParams {
	return &ReadTraceJaegerRemoteSamplingStrategyParams{
		Context: ctx,
	}
}

// NewReadTraceJaegerRemoteSamplingStrategyParamsWithHTTPClient creates a new ReadTraceJaegerRemoteSamplingStrategyParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadTraceJaegerRemoteSamplingStrategyParamsWithHTTPClient(client *http.Client) *ReadTraceJaegerRemoteSamplingStrategyParams {
	return &ReadTraceJaegerRemoteSamplingStrategyParams{
		HTTPClient: client,
	}
}

/*
ReadTraceJaegerRemoteSamplingStrategyParams contains all the parameters to send to the API endpoint

	for the read trace jaeger remote sampling strategy operation.

	Typically these are written to a http.Request.
*/
type ReadTraceJaegerRemoteSamplingStrategyParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read trace jaeger remote sampling strategy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadTraceJaegerRemoteSamplingStrategyParams) WithDefaults() *ReadTraceJaegerRemoteSamplingStrategyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read trace jaeger remote sampling strategy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadTraceJaegerRemoteSamplingStrategyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read trace jaeger remote sampling strategy params
func (o *ReadTraceJaegerRemoteSamplingStrategyParams) WithTimeout(timeout time.Duration) *ReadTraceJaegerRemoteSamplingStrategyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read trace jaeger remote sampling strategy params
func (o *ReadTraceJaegerRemoteSamplingStrategyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read trace jaeger remote sampling strategy params
func (o *ReadTraceJaegerRemoteSamplingStrategyParams) WithContext(ctx context.Context) *ReadTraceJaegerRemoteSamplingStrategyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read trace jaeger remote sampling strategy params
func (o *ReadTraceJaegerRemoteSamplingStrategyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read trace jaeger remote sampling strategy params
func (o *ReadTraceJaegerRemoteSamplingStrategyParams) WithHTTPClient(client *http.Client) *ReadTraceJaegerRemoteSamplingStrategyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read trace jaeger remote sampling strategy params
func (o *ReadTraceJaegerRemoteSamplingStrategyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the read trace jaeger remote sampling strategy params
func (o *ReadTraceJaegerRemoteSamplingStrategyParams) WithSlug(slug string) *ReadTraceJaegerRemoteSamplingStrategyParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the read trace jaeger remote sampling strategy params
func (o *ReadTraceJaegerRemoteSamplingStrategyParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ReadTraceJaegerRemoteSamplingStrategyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

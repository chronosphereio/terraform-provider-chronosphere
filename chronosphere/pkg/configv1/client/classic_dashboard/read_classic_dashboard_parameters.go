// Code generated by go-swagger; DO NOT EDIT.

package classic_dashboard

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

// NewReadClassicDashboardParams creates a new ReadClassicDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadClassicDashboardParams() *ReadClassicDashboardParams {
	return &ReadClassicDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadClassicDashboardParamsWithTimeout creates a new ReadClassicDashboardParams object
// with the ability to set a timeout on a request.
func NewReadClassicDashboardParamsWithTimeout(timeout time.Duration) *ReadClassicDashboardParams {
	return &ReadClassicDashboardParams{
		timeout: timeout,
	}
}

// NewReadClassicDashboardParamsWithContext creates a new ReadClassicDashboardParams object
// with the ability to set a context for a request.
func NewReadClassicDashboardParamsWithContext(ctx context.Context) *ReadClassicDashboardParams {
	return &ReadClassicDashboardParams{
		Context: ctx,
	}
}

// NewReadClassicDashboardParamsWithHTTPClient creates a new ReadClassicDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadClassicDashboardParamsWithHTTPClient(client *http.Client) *ReadClassicDashboardParams {
	return &ReadClassicDashboardParams{
		HTTPClient: client,
	}
}

/*
ReadClassicDashboardParams contains all the parameters to send to the API endpoint

	for the read classic dashboard operation.

	Typically these are written to a http.Request.
*/
type ReadClassicDashboardParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read classic dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadClassicDashboardParams) WithDefaults() *ReadClassicDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read classic dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadClassicDashboardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read classic dashboard params
func (o *ReadClassicDashboardParams) WithTimeout(timeout time.Duration) *ReadClassicDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read classic dashboard params
func (o *ReadClassicDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read classic dashboard params
func (o *ReadClassicDashboardParams) WithContext(ctx context.Context) *ReadClassicDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read classic dashboard params
func (o *ReadClassicDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read classic dashboard params
func (o *ReadClassicDashboardParams) WithHTTPClient(client *http.Client) *ReadClassicDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read classic dashboard params
func (o *ReadClassicDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the read classic dashboard params
func (o *ReadClassicDashboardParams) WithSlug(slug string) *ReadClassicDashboardParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the read classic dashboard params
func (o *ReadClassicDashboardParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ReadClassicDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

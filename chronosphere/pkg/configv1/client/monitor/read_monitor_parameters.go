// Code generated by go-swagger; DO NOT EDIT.

package monitor

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

// NewReadMonitorParams creates a new ReadMonitorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadMonitorParams() *ReadMonitorParams {
	return &ReadMonitorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadMonitorParamsWithTimeout creates a new ReadMonitorParams object
// with the ability to set a timeout on a request.
func NewReadMonitorParamsWithTimeout(timeout time.Duration) *ReadMonitorParams {
	return &ReadMonitorParams{
		timeout: timeout,
	}
}

// NewReadMonitorParamsWithContext creates a new ReadMonitorParams object
// with the ability to set a context for a request.
func NewReadMonitorParamsWithContext(ctx context.Context) *ReadMonitorParams {
	return &ReadMonitorParams{
		Context: ctx,
	}
}

// NewReadMonitorParamsWithHTTPClient creates a new ReadMonitorParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadMonitorParamsWithHTTPClient(client *http.Client) *ReadMonitorParams {
	return &ReadMonitorParams{
		HTTPClient: client,
	}
}

/*
ReadMonitorParams contains all the parameters to send to the API endpoint

	for the read monitor operation.

	Typically these are written to a http.Request.
*/
type ReadMonitorParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read monitor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadMonitorParams) WithDefaults() *ReadMonitorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read monitor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadMonitorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read monitor params
func (o *ReadMonitorParams) WithTimeout(timeout time.Duration) *ReadMonitorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read monitor params
func (o *ReadMonitorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read monitor params
func (o *ReadMonitorParams) WithContext(ctx context.Context) *ReadMonitorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read monitor params
func (o *ReadMonitorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read monitor params
func (o *ReadMonitorParams) WithHTTPClient(client *http.Client) *ReadMonitorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read monitor params
func (o *ReadMonitorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the read monitor params
func (o *ReadMonitorParams) WithSlug(slug string) *ReadMonitorParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the read monitor params
func (o *ReadMonitorParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ReadMonitorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package bucket

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

// NewReadBucketParams creates a new ReadBucketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadBucketParams() *ReadBucketParams {
	return &ReadBucketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadBucketParamsWithTimeout creates a new ReadBucketParams object
// with the ability to set a timeout on a request.
func NewReadBucketParamsWithTimeout(timeout time.Duration) *ReadBucketParams {
	return &ReadBucketParams{
		timeout: timeout,
	}
}

// NewReadBucketParamsWithContext creates a new ReadBucketParams object
// with the ability to set a context for a request.
func NewReadBucketParamsWithContext(ctx context.Context) *ReadBucketParams {
	return &ReadBucketParams{
		Context: ctx,
	}
}

// NewReadBucketParamsWithHTTPClient creates a new ReadBucketParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadBucketParamsWithHTTPClient(client *http.Client) *ReadBucketParams {
	return &ReadBucketParams{
		HTTPClient: client,
	}
}

/*
ReadBucketParams contains all the parameters to send to the API endpoint

	for the read bucket operation.

	Typically these are written to a http.Request.
*/
type ReadBucketParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadBucketParams) WithDefaults() *ReadBucketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadBucketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read bucket params
func (o *ReadBucketParams) WithTimeout(timeout time.Duration) *ReadBucketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read bucket params
func (o *ReadBucketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read bucket params
func (o *ReadBucketParams) WithContext(ctx context.Context) *ReadBucketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read bucket params
func (o *ReadBucketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read bucket params
func (o *ReadBucketParams) WithHTTPClient(client *http.Client) *ReadBucketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read bucket params
func (o *ReadBucketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the read bucket params
func (o *ReadBucketParams) WithSlug(slug string) *ReadBucketParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the read bucket params
func (o *ReadBucketParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ReadBucketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

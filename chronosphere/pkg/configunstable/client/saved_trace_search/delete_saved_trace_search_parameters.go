// Code generated by go-swagger; DO NOT EDIT.

package saved_trace_search

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

// NewDeleteSavedTraceSearchParams creates a new DeleteSavedTraceSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSavedTraceSearchParams() *DeleteSavedTraceSearchParams {
	return &DeleteSavedTraceSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSavedTraceSearchParamsWithTimeout creates a new DeleteSavedTraceSearchParams object
// with the ability to set a timeout on a request.
func NewDeleteSavedTraceSearchParamsWithTimeout(timeout time.Duration) *DeleteSavedTraceSearchParams {
	return &DeleteSavedTraceSearchParams{
		timeout: timeout,
	}
}

// NewDeleteSavedTraceSearchParamsWithContext creates a new DeleteSavedTraceSearchParams object
// with the ability to set a context for a request.
func NewDeleteSavedTraceSearchParamsWithContext(ctx context.Context) *DeleteSavedTraceSearchParams {
	return &DeleteSavedTraceSearchParams{
		Context: ctx,
	}
}

// NewDeleteSavedTraceSearchParamsWithHTTPClient creates a new DeleteSavedTraceSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSavedTraceSearchParamsWithHTTPClient(client *http.Client) *DeleteSavedTraceSearchParams {
	return &DeleteSavedTraceSearchParams{
		HTTPClient: client,
	}
}

/*
DeleteSavedTraceSearchParams contains all the parameters to send to the API endpoint

	for the delete saved trace search operation.

	Typically these are written to a http.Request.
*/
type DeleteSavedTraceSearchParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete saved trace search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSavedTraceSearchParams) WithDefaults() *DeleteSavedTraceSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete saved trace search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSavedTraceSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete saved trace search params
func (o *DeleteSavedTraceSearchParams) WithTimeout(timeout time.Duration) *DeleteSavedTraceSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete saved trace search params
func (o *DeleteSavedTraceSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete saved trace search params
func (o *DeleteSavedTraceSearchParams) WithContext(ctx context.Context) *DeleteSavedTraceSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete saved trace search params
func (o *DeleteSavedTraceSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete saved trace search params
func (o *DeleteSavedTraceSearchParams) WithHTTPClient(client *http.Client) *DeleteSavedTraceSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete saved trace search params
func (o *DeleteSavedTraceSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the delete saved trace search params
func (o *DeleteSavedTraceSearchParams) WithSlug(slug string) *DeleteSavedTraceSearchParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the delete saved trace search params
func (o *DeleteSavedTraceSearchParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSavedTraceSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
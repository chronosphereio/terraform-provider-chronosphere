// Code generated by go-swagger; DO NOT EDIT.

package log_scale_action

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

// NewDeleteLogScaleActionParams creates a new DeleteLogScaleActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLogScaleActionParams() *DeleteLogScaleActionParams {
	return &DeleteLogScaleActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLogScaleActionParamsWithTimeout creates a new DeleteLogScaleActionParams object
// with the ability to set a timeout on a request.
func NewDeleteLogScaleActionParamsWithTimeout(timeout time.Duration) *DeleteLogScaleActionParams {
	return &DeleteLogScaleActionParams{
		timeout: timeout,
	}
}

// NewDeleteLogScaleActionParamsWithContext creates a new DeleteLogScaleActionParams object
// with the ability to set a context for a request.
func NewDeleteLogScaleActionParamsWithContext(ctx context.Context) *DeleteLogScaleActionParams {
	return &DeleteLogScaleActionParams{
		Context: ctx,
	}
}

// NewDeleteLogScaleActionParamsWithHTTPClient creates a new DeleteLogScaleActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLogScaleActionParamsWithHTTPClient(client *http.Client) *DeleteLogScaleActionParams {
	return &DeleteLogScaleActionParams{
		HTTPClient: client,
	}
}

/*
DeleteLogScaleActionParams contains all the parameters to send to the API endpoint

	for the delete log scale action operation.

	Typically these are written to a http.Request.
*/
type DeleteLogScaleActionParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete log scale action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLogScaleActionParams) WithDefaults() *DeleteLogScaleActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete log scale action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLogScaleActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete log scale action params
func (o *DeleteLogScaleActionParams) WithTimeout(timeout time.Duration) *DeleteLogScaleActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete log scale action params
func (o *DeleteLogScaleActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete log scale action params
func (o *DeleteLogScaleActionParams) WithContext(ctx context.Context) *DeleteLogScaleActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete log scale action params
func (o *DeleteLogScaleActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete log scale action params
func (o *DeleteLogScaleActionParams) WithHTTPClient(client *http.Client) *DeleteLogScaleActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete log scale action params
func (o *DeleteLogScaleActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the delete log scale action params
func (o *DeleteLogScaleActionParams) WithSlug(slug string) *DeleteLogScaleActionParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the delete log scale action params
func (o *DeleteLogScaleActionParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLogScaleActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
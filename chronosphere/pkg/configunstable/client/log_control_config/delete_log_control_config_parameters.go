// Code generated by go-swagger; DO NOT EDIT.

package log_control_config

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

// NewDeleteLogControlConfigParams creates a new DeleteLogControlConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLogControlConfigParams() *DeleteLogControlConfigParams {
	return &DeleteLogControlConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLogControlConfigParamsWithTimeout creates a new DeleteLogControlConfigParams object
// with the ability to set a timeout on a request.
func NewDeleteLogControlConfigParamsWithTimeout(timeout time.Duration) *DeleteLogControlConfigParams {
	return &DeleteLogControlConfigParams{
		timeout: timeout,
	}
}

// NewDeleteLogControlConfigParamsWithContext creates a new DeleteLogControlConfigParams object
// with the ability to set a context for a request.
func NewDeleteLogControlConfigParamsWithContext(ctx context.Context) *DeleteLogControlConfigParams {
	return &DeleteLogControlConfigParams{
		Context: ctx,
	}
}

// NewDeleteLogControlConfigParamsWithHTTPClient creates a new DeleteLogControlConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLogControlConfigParamsWithHTTPClient(client *http.Client) *DeleteLogControlConfigParams {
	return &DeleteLogControlConfigParams{
		HTTPClient: client,
	}
}

/*
DeleteLogControlConfigParams contains all the parameters to send to the API endpoint

	for the delete log control config operation.

	Typically these are written to a http.Request.
*/
type DeleteLogControlConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete log control config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLogControlConfigParams) WithDefaults() *DeleteLogControlConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete log control config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLogControlConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete log control config params
func (o *DeleteLogControlConfigParams) WithTimeout(timeout time.Duration) *DeleteLogControlConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete log control config params
func (o *DeleteLogControlConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete log control config params
func (o *DeleteLogControlConfigParams) WithContext(ctx context.Context) *DeleteLogControlConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete log control config params
func (o *DeleteLogControlConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete log control config params
func (o *DeleteLogControlConfigParams) WithHTTPClient(client *http.Client) *DeleteLogControlConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete log control config params
func (o *DeleteLogControlConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLogControlConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package muting_rule

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

// NewReadMutingRuleParams creates a new ReadMutingRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadMutingRuleParams() *ReadMutingRuleParams {
	return &ReadMutingRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadMutingRuleParamsWithTimeout creates a new ReadMutingRuleParams object
// with the ability to set a timeout on a request.
func NewReadMutingRuleParamsWithTimeout(timeout time.Duration) *ReadMutingRuleParams {
	return &ReadMutingRuleParams{
		timeout: timeout,
	}
}

// NewReadMutingRuleParamsWithContext creates a new ReadMutingRuleParams object
// with the ability to set a context for a request.
func NewReadMutingRuleParamsWithContext(ctx context.Context) *ReadMutingRuleParams {
	return &ReadMutingRuleParams{
		Context: ctx,
	}
}

// NewReadMutingRuleParamsWithHTTPClient creates a new ReadMutingRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadMutingRuleParamsWithHTTPClient(client *http.Client) *ReadMutingRuleParams {
	return &ReadMutingRuleParams{
		HTTPClient: client,
	}
}

/*
ReadMutingRuleParams contains all the parameters to send to the API endpoint

	for the read muting rule operation.

	Typically these are written to a http.Request.
*/
type ReadMutingRuleParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read muting rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadMutingRuleParams) WithDefaults() *ReadMutingRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read muting rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadMutingRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read muting rule params
func (o *ReadMutingRuleParams) WithTimeout(timeout time.Duration) *ReadMutingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read muting rule params
func (o *ReadMutingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read muting rule params
func (o *ReadMutingRuleParams) WithContext(ctx context.Context) *ReadMutingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read muting rule params
func (o *ReadMutingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read muting rule params
func (o *ReadMutingRuleParams) WithHTTPClient(client *http.Client) *ReadMutingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read muting rule params
func (o *ReadMutingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the read muting rule params
func (o *ReadMutingRuleParams) WithSlug(slug string) *ReadMutingRuleParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the read muting rule params
func (o *ReadMutingRuleParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ReadMutingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

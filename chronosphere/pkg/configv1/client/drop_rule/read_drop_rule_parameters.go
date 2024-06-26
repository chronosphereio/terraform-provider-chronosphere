// Code generated by go-swagger; DO NOT EDIT.

package drop_rule

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

// NewReadDropRuleParams creates a new ReadDropRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadDropRuleParams() *ReadDropRuleParams {
	return &ReadDropRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadDropRuleParamsWithTimeout creates a new ReadDropRuleParams object
// with the ability to set a timeout on a request.
func NewReadDropRuleParamsWithTimeout(timeout time.Duration) *ReadDropRuleParams {
	return &ReadDropRuleParams{
		timeout: timeout,
	}
}

// NewReadDropRuleParamsWithContext creates a new ReadDropRuleParams object
// with the ability to set a context for a request.
func NewReadDropRuleParamsWithContext(ctx context.Context) *ReadDropRuleParams {
	return &ReadDropRuleParams{
		Context: ctx,
	}
}

// NewReadDropRuleParamsWithHTTPClient creates a new ReadDropRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadDropRuleParamsWithHTTPClient(client *http.Client) *ReadDropRuleParams {
	return &ReadDropRuleParams{
		HTTPClient: client,
	}
}

/*
ReadDropRuleParams contains all the parameters to send to the API endpoint

	for the read drop rule operation.

	Typically these are written to a http.Request.
*/
type ReadDropRuleParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read drop rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadDropRuleParams) WithDefaults() *ReadDropRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read drop rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadDropRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read drop rule params
func (o *ReadDropRuleParams) WithTimeout(timeout time.Duration) *ReadDropRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read drop rule params
func (o *ReadDropRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read drop rule params
func (o *ReadDropRuleParams) WithContext(ctx context.Context) *ReadDropRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read drop rule params
func (o *ReadDropRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read drop rule params
func (o *ReadDropRuleParams) WithHTTPClient(client *http.Client) *ReadDropRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read drop rule params
func (o *ReadDropRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the read drop rule params
func (o *ReadDropRuleParams) WithSlug(slug string) *ReadDropRuleParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the read drop rule params
func (o *ReadDropRuleParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ReadDropRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

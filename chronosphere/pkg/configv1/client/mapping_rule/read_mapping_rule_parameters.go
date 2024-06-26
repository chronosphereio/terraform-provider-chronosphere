// Code generated by go-swagger; DO NOT EDIT.

package mapping_rule

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

// NewReadMappingRuleParams creates a new ReadMappingRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadMappingRuleParams() *ReadMappingRuleParams {
	return &ReadMappingRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadMappingRuleParamsWithTimeout creates a new ReadMappingRuleParams object
// with the ability to set a timeout on a request.
func NewReadMappingRuleParamsWithTimeout(timeout time.Duration) *ReadMappingRuleParams {
	return &ReadMappingRuleParams{
		timeout: timeout,
	}
}

// NewReadMappingRuleParamsWithContext creates a new ReadMappingRuleParams object
// with the ability to set a context for a request.
func NewReadMappingRuleParamsWithContext(ctx context.Context) *ReadMappingRuleParams {
	return &ReadMappingRuleParams{
		Context: ctx,
	}
}

// NewReadMappingRuleParamsWithHTTPClient creates a new ReadMappingRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadMappingRuleParamsWithHTTPClient(client *http.Client) *ReadMappingRuleParams {
	return &ReadMappingRuleParams{
		HTTPClient: client,
	}
}

/*
ReadMappingRuleParams contains all the parameters to send to the API endpoint

	for the read mapping rule operation.

	Typically these are written to a http.Request.
*/
type ReadMappingRuleParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read mapping rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadMappingRuleParams) WithDefaults() *ReadMappingRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read mapping rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadMappingRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read mapping rule params
func (o *ReadMappingRuleParams) WithTimeout(timeout time.Duration) *ReadMappingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read mapping rule params
func (o *ReadMappingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read mapping rule params
func (o *ReadMappingRuleParams) WithContext(ctx context.Context) *ReadMappingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read mapping rule params
func (o *ReadMappingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read mapping rule params
func (o *ReadMappingRuleParams) WithHTTPClient(client *http.Client) *ReadMappingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read mapping rule params
func (o *ReadMappingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the read mapping rule params
func (o *ReadMappingRuleParams) WithSlug(slug string) *ReadMappingRuleParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the read mapping rule params
func (o *ReadMappingRuleParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ReadMappingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

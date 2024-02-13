// Code generated by go-swagger; DO NOT EDIT.

package rollup_rule

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

// NewUpdateRollupRuleParams creates a new UpdateRollupRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRollupRuleParams() *UpdateRollupRuleParams {
	return &UpdateRollupRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRollupRuleParamsWithTimeout creates a new UpdateRollupRuleParams object
// with the ability to set a timeout on a request.
func NewUpdateRollupRuleParamsWithTimeout(timeout time.Duration) *UpdateRollupRuleParams {
	return &UpdateRollupRuleParams{
		timeout: timeout,
	}
}

// NewUpdateRollupRuleParamsWithContext creates a new UpdateRollupRuleParams object
// with the ability to set a context for a request.
func NewUpdateRollupRuleParamsWithContext(ctx context.Context) *UpdateRollupRuleParams {
	return &UpdateRollupRuleParams{
		Context: ctx,
	}
}

// NewUpdateRollupRuleParamsWithHTTPClient creates a new UpdateRollupRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRollupRuleParamsWithHTTPClient(client *http.Client) *UpdateRollupRuleParams {
	return &UpdateRollupRuleParams{
		HTTPClient: client,
	}
}

/*
UpdateRollupRuleParams contains all the parameters to send to the API endpoint

	for the update rollup rule operation.

	Typically these are written to a http.Request.
*/
type UpdateRollupRuleParams struct {

	// Body.
	Body UpdateRollupRuleBody

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update rollup rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRollupRuleParams) WithDefaults() *UpdateRollupRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update rollup rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRollupRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update rollup rule params
func (o *UpdateRollupRuleParams) WithTimeout(timeout time.Duration) *UpdateRollupRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update rollup rule params
func (o *UpdateRollupRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update rollup rule params
func (o *UpdateRollupRuleParams) WithContext(ctx context.Context) *UpdateRollupRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update rollup rule params
func (o *UpdateRollupRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update rollup rule params
func (o *UpdateRollupRuleParams) WithHTTPClient(client *http.Client) *UpdateRollupRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update rollup rule params
func (o *UpdateRollupRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update rollup rule params
func (o *UpdateRollupRuleParams) WithBody(body UpdateRollupRuleBody) *UpdateRollupRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update rollup rule params
func (o *UpdateRollupRuleParams) SetBody(body UpdateRollupRuleBody) {
	o.Body = body
}

// WithSlug adds the slug to the update rollup rule params
func (o *UpdateRollupRuleParams) WithSlug(slug string) *UpdateRollupRuleParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the update rollup rule params
func (o *UpdateRollupRuleParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRollupRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param slug
	if err := r.SetPathParam("slug", o.Slug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

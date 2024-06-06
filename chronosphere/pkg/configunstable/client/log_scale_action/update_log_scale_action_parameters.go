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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

// NewUpdateLogScaleActionParams creates a new UpdateLogScaleActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLogScaleActionParams() *UpdateLogScaleActionParams {
	return &UpdateLogScaleActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLogScaleActionParamsWithTimeout creates a new UpdateLogScaleActionParams object
// with the ability to set a timeout on a request.
func NewUpdateLogScaleActionParamsWithTimeout(timeout time.Duration) *UpdateLogScaleActionParams {
	return &UpdateLogScaleActionParams{
		timeout: timeout,
	}
}

// NewUpdateLogScaleActionParamsWithContext creates a new UpdateLogScaleActionParams object
// with the ability to set a context for a request.
func NewUpdateLogScaleActionParamsWithContext(ctx context.Context) *UpdateLogScaleActionParams {
	return &UpdateLogScaleActionParams{
		Context: ctx,
	}
}

// NewUpdateLogScaleActionParamsWithHTTPClient creates a new UpdateLogScaleActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLogScaleActionParamsWithHTTPClient(client *http.Client) *UpdateLogScaleActionParams {
	return &UpdateLogScaleActionParams{
		HTTPClient: client,
	}
}

/*
UpdateLogScaleActionParams contains all the parameters to send to the API endpoint

	for the update log scale action operation.

	Typically these are written to a http.Request.
*/
type UpdateLogScaleActionParams struct {

	// Body.
	Body *models.ConfigUnstableUpdateLogScaleActionBody

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update log scale action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLogScaleActionParams) WithDefaults() *UpdateLogScaleActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update log scale action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLogScaleActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update log scale action params
func (o *UpdateLogScaleActionParams) WithTimeout(timeout time.Duration) *UpdateLogScaleActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update log scale action params
func (o *UpdateLogScaleActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update log scale action params
func (o *UpdateLogScaleActionParams) WithContext(ctx context.Context) *UpdateLogScaleActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update log scale action params
func (o *UpdateLogScaleActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update log scale action params
func (o *UpdateLogScaleActionParams) WithHTTPClient(client *http.Client) *UpdateLogScaleActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update log scale action params
func (o *UpdateLogScaleActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update log scale action params
func (o *UpdateLogScaleActionParams) WithBody(body *models.ConfigUnstableUpdateLogScaleActionBody) *UpdateLogScaleActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update log scale action params
func (o *UpdateLogScaleActionParams) SetBody(body *models.ConfigUnstableUpdateLogScaleActionBody) {
	o.Body = body
}

// WithSlug adds the slug to the update log scale action params
func (o *UpdateLogScaleActionParams) WithSlug(slug string) *UpdateLogScaleActionParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the update log scale action params
func (o *UpdateLogScaleActionParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLogScaleActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

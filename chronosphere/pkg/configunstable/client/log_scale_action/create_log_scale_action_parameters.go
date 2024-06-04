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

// NewCreateLogScaleActionParams creates a new CreateLogScaleActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateLogScaleActionParams() *CreateLogScaleActionParams {
	return &CreateLogScaleActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLogScaleActionParamsWithTimeout creates a new CreateLogScaleActionParams object
// with the ability to set a timeout on a request.
func NewCreateLogScaleActionParamsWithTimeout(timeout time.Duration) *CreateLogScaleActionParams {
	return &CreateLogScaleActionParams{
		timeout: timeout,
	}
}

// NewCreateLogScaleActionParamsWithContext creates a new CreateLogScaleActionParams object
// with the ability to set a context for a request.
func NewCreateLogScaleActionParamsWithContext(ctx context.Context) *CreateLogScaleActionParams {
	return &CreateLogScaleActionParams{
		Context: ctx,
	}
}

// NewCreateLogScaleActionParamsWithHTTPClient creates a new CreateLogScaleActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateLogScaleActionParamsWithHTTPClient(client *http.Client) *CreateLogScaleActionParams {
	return &CreateLogScaleActionParams{
		HTTPClient: client,
	}
}

/*
CreateLogScaleActionParams contains all the parameters to send to the API endpoint

	for the create log scale action operation.

	Typically these are written to a http.Request.
*/
type CreateLogScaleActionParams struct {

	// Body.
	Body *models.ConfigunstableCreateLogScaleActionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create log scale action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateLogScaleActionParams) WithDefaults() *CreateLogScaleActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create log scale action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateLogScaleActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create log scale action params
func (o *CreateLogScaleActionParams) WithTimeout(timeout time.Duration) *CreateLogScaleActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create log scale action params
func (o *CreateLogScaleActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create log scale action params
func (o *CreateLogScaleActionParams) WithContext(ctx context.Context) *CreateLogScaleActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create log scale action params
func (o *CreateLogScaleActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create log scale action params
func (o *CreateLogScaleActionParams) WithHTTPClient(client *http.Client) *CreateLogScaleActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create log scale action params
func (o *CreateLogScaleActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create log scale action params
func (o *CreateLogScaleActionParams) WithBody(body *models.ConfigunstableCreateLogScaleActionRequest) *CreateLogScaleActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create log scale action params
func (o *CreateLogScaleActionParams) SetBody(body *models.ConfigunstableCreateLogScaleActionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLogScaleActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

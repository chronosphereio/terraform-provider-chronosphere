// Code generated by go-swagger; DO NOT EDIT.

package trace_tail_sampling_rules

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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// NewUpdateTraceTailSamplingRulesParams creates a new UpdateTraceTailSamplingRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateTraceTailSamplingRulesParams() *UpdateTraceTailSamplingRulesParams {
	return &UpdateTraceTailSamplingRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTraceTailSamplingRulesParamsWithTimeout creates a new UpdateTraceTailSamplingRulesParams object
// with the ability to set a timeout on a request.
func NewUpdateTraceTailSamplingRulesParamsWithTimeout(timeout time.Duration) *UpdateTraceTailSamplingRulesParams {
	return &UpdateTraceTailSamplingRulesParams{
		timeout: timeout,
	}
}

// NewUpdateTraceTailSamplingRulesParamsWithContext creates a new UpdateTraceTailSamplingRulesParams object
// with the ability to set a context for a request.
func NewUpdateTraceTailSamplingRulesParamsWithContext(ctx context.Context) *UpdateTraceTailSamplingRulesParams {
	return &UpdateTraceTailSamplingRulesParams{
		Context: ctx,
	}
}

// NewUpdateTraceTailSamplingRulesParamsWithHTTPClient creates a new UpdateTraceTailSamplingRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateTraceTailSamplingRulesParamsWithHTTPClient(client *http.Client) *UpdateTraceTailSamplingRulesParams {
	return &UpdateTraceTailSamplingRulesParams{
		HTTPClient: client,
	}
}

/*
UpdateTraceTailSamplingRulesParams contains all the parameters to send to the API endpoint

	for the update trace tail sampling rules operation.

	Typically these are written to a http.Request.
*/
type UpdateTraceTailSamplingRulesParams struct {

	// Body.
	Body *models.Configv1UpdateTraceTailSamplingRulesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update trace tail sampling rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTraceTailSamplingRulesParams) WithDefaults() *UpdateTraceTailSamplingRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update trace tail sampling rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTraceTailSamplingRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update trace tail sampling rules params
func (o *UpdateTraceTailSamplingRulesParams) WithTimeout(timeout time.Duration) *UpdateTraceTailSamplingRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update trace tail sampling rules params
func (o *UpdateTraceTailSamplingRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update trace tail sampling rules params
func (o *UpdateTraceTailSamplingRulesParams) WithContext(ctx context.Context) *UpdateTraceTailSamplingRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update trace tail sampling rules params
func (o *UpdateTraceTailSamplingRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update trace tail sampling rules params
func (o *UpdateTraceTailSamplingRulesParams) WithHTTPClient(client *http.Client) *UpdateTraceTailSamplingRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update trace tail sampling rules params
func (o *UpdateTraceTailSamplingRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update trace tail sampling rules params
func (o *UpdateTraceTailSamplingRulesParams) WithBody(body *models.Configv1UpdateTraceTailSamplingRulesRequest) *UpdateTraceTailSamplingRulesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update trace tail sampling rules params
func (o *UpdateTraceTailSamplingRulesParams) SetBody(body *models.Configv1UpdateTraceTailSamplingRulesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTraceTailSamplingRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

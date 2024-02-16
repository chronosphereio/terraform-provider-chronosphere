// Code generated by go-swagger; DO NOT EDIT.

package derived_metric

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

// NewCreateDerivedMetricParams creates a new CreateDerivedMetricParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDerivedMetricParams() *CreateDerivedMetricParams {
	return &CreateDerivedMetricParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDerivedMetricParamsWithTimeout creates a new CreateDerivedMetricParams object
// with the ability to set a timeout on a request.
func NewCreateDerivedMetricParamsWithTimeout(timeout time.Duration) *CreateDerivedMetricParams {
	return &CreateDerivedMetricParams{
		timeout: timeout,
	}
}

// NewCreateDerivedMetricParamsWithContext creates a new CreateDerivedMetricParams object
// with the ability to set a context for a request.
func NewCreateDerivedMetricParamsWithContext(ctx context.Context) *CreateDerivedMetricParams {
	return &CreateDerivedMetricParams{
		Context: ctx,
	}
}

// NewCreateDerivedMetricParamsWithHTTPClient creates a new CreateDerivedMetricParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDerivedMetricParamsWithHTTPClient(client *http.Client) *CreateDerivedMetricParams {
	return &CreateDerivedMetricParams{
		HTTPClient: client,
	}
}

/*
CreateDerivedMetricParams contains all the parameters to send to the API endpoint

	for the create derived metric operation.

	Typically these are written to a http.Request.
*/
type CreateDerivedMetricParams struct {

	// Body.
	Body *models.Configv1CreateDerivedMetricRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create derived metric params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDerivedMetricParams) WithDefaults() *CreateDerivedMetricParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create derived metric params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDerivedMetricParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create derived metric params
func (o *CreateDerivedMetricParams) WithTimeout(timeout time.Duration) *CreateDerivedMetricParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create derived metric params
func (o *CreateDerivedMetricParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create derived metric params
func (o *CreateDerivedMetricParams) WithContext(ctx context.Context) *CreateDerivedMetricParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create derived metric params
func (o *CreateDerivedMetricParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create derived metric params
func (o *CreateDerivedMetricParams) WithHTTPClient(client *http.Client) *CreateDerivedMetricParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create derived metric params
func (o *CreateDerivedMetricParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create derived metric params
func (o *CreateDerivedMetricParams) WithBody(body *models.Configv1CreateDerivedMetricRequest) *CreateDerivedMetricParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create derived metric params
func (o *CreateDerivedMetricParams) SetBody(body *models.Configv1CreateDerivedMetricRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDerivedMetricParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
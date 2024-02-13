// Code generated by go-swagger; DO NOT EDIT.

package service

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

// NewCreateServiceParams creates a new CreateServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateServiceParams() *CreateServiceParams {
	return &CreateServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateServiceParamsWithTimeout creates a new CreateServiceParams object
// with the ability to set a timeout on a request.
func NewCreateServiceParamsWithTimeout(timeout time.Duration) *CreateServiceParams {
	return &CreateServiceParams{
		timeout: timeout,
	}
}

// NewCreateServiceParamsWithContext creates a new CreateServiceParams object
// with the ability to set a context for a request.
func NewCreateServiceParamsWithContext(ctx context.Context) *CreateServiceParams {
	return &CreateServiceParams{
		Context: ctx,
	}
}

// NewCreateServiceParamsWithHTTPClient creates a new CreateServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateServiceParamsWithHTTPClient(client *http.Client) *CreateServiceParams {
	return &CreateServiceParams{
		HTTPClient: client,
	}
}

/*
CreateServiceParams contains all the parameters to send to the API endpoint

	for the create service operation.

	Typically these are written to a http.Request.
*/
type CreateServiceParams struct {

	// Body.
	Body *models.ConfigunstableCreateServiceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServiceParams) WithDefaults() *CreateServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create service params
func (o *CreateServiceParams) WithTimeout(timeout time.Duration) *CreateServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create service params
func (o *CreateServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create service params
func (o *CreateServiceParams) WithContext(ctx context.Context) *CreateServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create service params
func (o *CreateServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create service params
func (o *CreateServiceParams) WithHTTPClient(client *http.Client) *CreateServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create service params
func (o *CreateServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create service params
func (o *CreateServiceParams) WithBody(body *models.ConfigunstableCreateServiceRequest) *CreateServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create service params
func (o *CreateServiceParams) SetBody(body *models.ConfigunstableCreateServiceRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

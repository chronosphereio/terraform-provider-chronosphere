// Code generated by go-swagger; DO NOT EDIT.

package resource_pools

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

// NewCreateResourcePoolsParams creates a new CreateResourcePoolsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateResourcePoolsParams() *CreateResourcePoolsParams {
	return &CreateResourcePoolsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateResourcePoolsParamsWithTimeout creates a new CreateResourcePoolsParams object
// with the ability to set a timeout on a request.
func NewCreateResourcePoolsParamsWithTimeout(timeout time.Duration) *CreateResourcePoolsParams {
	return &CreateResourcePoolsParams{
		timeout: timeout,
	}
}

// NewCreateResourcePoolsParamsWithContext creates a new CreateResourcePoolsParams object
// with the ability to set a context for a request.
func NewCreateResourcePoolsParamsWithContext(ctx context.Context) *CreateResourcePoolsParams {
	return &CreateResourcePoolsParams{
		Context: ctx,
	}
}

// NewCreateResourcePoolsParamsWithHTTPClient creates a new CreateResourcePoolsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateResourcePoolsParamsWithHTTPClient(client *http.Client) *CreateResourcePoolsParams {
	return &CreateResourcePoolsParams{
		HTTPClient: client,
	}
}

/*
CreateResourcePoolsParams contains all the parameters to send to the API endpoint

	for the create resource pools operation.

	Typically these are written to a http.Request.
*/
type CreateResourcePoolsParams struct {

	// Body.
	Body *models.Configv1CreateResourcePoolsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create resource pools params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateResourcePoolsParams) WithDefaults() *CreateResourcePoolsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create resource pools params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateResourcePoolsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create resource pools params
func (o *CreateResourcePoolsParams) WithTimeout(timeout time.Duration) *CreateResourcePoolsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create resource pools params
func (o *CreateResourcePoolsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create resource pools params
func (o *CreateResourcePoolsParams) WithContext(ctx context.Context) *CreateResourcePoolsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create resource pools params
func (o *CreateResourcePoolsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create resource pools params
func (o *CreateResourcePoolsParams) WithHTTPClient(client *http.Client) *CreateResourcePoolsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create resource pools params
func (o *CreateResourcePoolsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create resource pools params
func (o *CreateResourcePoolsParams) WithBody(body *models.Configv1CreateResourcePoolsRequest) *CreateResourcePoolsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create resource pools params
func (o *CreateResourcePoolsParams) SetBody(body *models.Configv1CreateResourcePoolsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateResourcePoolsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
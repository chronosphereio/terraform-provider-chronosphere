// Code generated by go-swagger; DO NOT EDIT.

package grafana_dashboard

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

// NewDeleteGrafanaDashboardParams creates a new DeleteGrafanaDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGrafanaDashboardParams() *DeleteGrafanaDashboardParams {
	return &DeleteGrafanaDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGrafanaDashboardParamsWithTimeout creates a new DeleteGrafanaDashboardParams object
// with the ability to set a timeout on a request.
func NewDeleteGrafanaDashboardParamsWithTimeout(timeout time.Duration) *DeleteGrafanaDashboardParams {
	return &DeleteGrafanaDashboardParams{
		timeout: timeout,
	}
}

// NewDeleteGrafanaDashboardParamsWithContext creates a new DeleteGrafanaDashboardParams object
// with the ability to set a context for a request.
func NewDeleteGrafanaDashboardParamsWithContext(ctx context.Context) *DeleteGrafanaDashboardParams {
	return &DeleteGrafanaDashboardParams{
		Context: ctx,
	}
}

// NewDeleteGrafanaDashboardParamsWithHTTPClient creates a new DeleteGrafanaDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGrafanaDashboardParamsWithHTTPClient(client *http.Client) *DeleteGrafanaDashboardParams {
	return &DeleteGrafanaDashboardParams{
		HTTPClient: client,
	}
}

/*
DeleteGrafanaDashboardParams contains all the parameters to send to the API endpoint

	for the delete grafana dashboard operation.

	Typically these are written to a http.Request.
*/
type DeleteGrafanaDashboardParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete grafana dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGrafanaDashboardParams) WithDefaults() *DeleteGrafanaDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete grafana dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGrafanaDashboardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete grafana dashboard params
func (o *DeleteGrafanaDashboardParams) WithTimeout(timeout time.Duration) *DeleteGrafanaDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete grafana dashboard params
func (o *DeleteGrafanaDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete grafana dashboard params
func (o *DeleteGrafanaDashboardParams) WithContext(ctx context.Context) *DeleteGrafanaDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete grafana dashboard params
func (o *DeleteGrafanaDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete grafana dashboard params
func (o *DeleteGrafanaDashboardParams) WithHTTPClient(client *http.Client) *DeleteGrafanaDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete grafana dashboard params
func (o *DeleteGrafanaDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the delete grafana dashboard params
func (o *DeleteGrafanaDashboardParams) WithSlug(slug string) *DeleteGrafanaDashboardParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the delete grafana dashboard params
func (o *DeleteGrafanaDashboardParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGrafanaDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

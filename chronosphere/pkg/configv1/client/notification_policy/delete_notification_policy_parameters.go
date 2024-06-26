// Code generated by go-swagger; DO NOT EDIT.

package notification_policy

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

// NewDeleteNotificationPolicyParams creates a new DeleteNotificationPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteNotificationPolicyParams() *DeleteNotificationPolicyParams {
	return &DeleteNotificationPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNotificationPolicyParamsWithTimeout creates a new DeleteNotificationPolicyParams object
// with the ability to set a timeout on a request.
func NewDeleteNotificationPolicyParamsWithTimeout(timeout time.Duration) *DeleteNotificationPolicyParams {
	return &DeleteNotificationPolicyParams{
		timeout: timeout,
	}
}

// NewDeleteNotificationPolicyParamsWithContext creates a new DeleteNotificationPolicyParams object
// with the ability to set a context for a request.
func NewDeleteNotificationPolicyParamsWithContext(ctx context.Context) *DeleteNotificationPolicyParams {
	return &DeleteNotificationPolicyParams{
		Context: ctx,
	}
}

// NewDeleteNotificationPolicyParamsWithHTTPClient creates a new DeleteNotificationPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteNotificationPolicyParamsWithHTTPClient(client *http.Client) *DeleteNotificationPolicyParams {
	return &DeleteNotificationPolicyParams{
		HTTPClient: client,
	}
}

/*
DeleteNotificationPolicyParams contains all the parameters to send to the API endpoint

	for the delete notification policy operation.

	Typically these are written to a http.Request.
*/
type DeleteNotificationPolicyParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete notification policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNotificationPolicyParams) WithDefaults() *DeleteNotificationPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete notification policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteNotificationPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete notification policy params
func (o *DeleteNotificationPolicyParams) WithTimeout(timeout time.Duration) *DeleteNotificationPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete notification policy params
func (o *DeleteNotificationPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete notification policy params
func (o *DeleteNotificationPolicyParams) WithContext(ctx context.Context) *DeleteNotificationPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete notification policy params
func (o *DeleteNotificationPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete notification policy params
func (o *DeleteNotificationPolicyParams) WithHTTPClient(client *http.Client) *DeleteNotificationPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete notification policy params
func (o *DeleteNotificationPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the delete notification policy params
func (o *DeleteNotificationPolicyParams) WithSlug(slug string) *DeleteNotificationPolicyParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the delete notification policy params
func (o *DeleteNotificationPolicyParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNotificationPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

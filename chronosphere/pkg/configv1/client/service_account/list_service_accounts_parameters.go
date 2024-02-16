// Code generated by go-swagger; DO NOT EDIT.

package service_account

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
	"github.com/go-openapi/swag"
)

// NewListServiceAccountsParams creates a new ListServiceAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListServiceAccountsParams() *ListServiceAccountsParams {
	return &ListServiceAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListServiceAccountsParamsWithTimeout creates a new ListServiceAccountsParams object
// with the ability to set a timeout on a request.
func NewListServiceAccountsParamsWithTimeout(timeout time.Duration) *ListServiceAccountsParams {
	return &ListServiceAccountsParams{
		timeout: timeout,
	}
}

// NewListServiceAccountsParamsWithContext creates a new ListServiceAccountsParams object
// with the ability to set a context for a request.
func NewListServiceAccountsParamsWithContext(ctx context.Context) *ListServiceAccountsParams {
	return &ListServiceAccountsParams{
		Context: ctx,
	}
}

// NewListServiceAccountsParamsWithHTTPClient creates a new ListServiceAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListServiceAccountsParamsWithHTTPClient(client *http.Client) *ListServiceAccountsParams {
	return &ListServiceAccountsParams{
		HTTPClient: client,
	}
}

/*
ListServiceAccountsParams contains all the parameters to send to the API endpoint

	for the list service accounts operation.

	Typically these are written to a http.Request.
*/
type ListServiceAccountsParams struct {

	/* Names.

	   Filters results by name, where any ServiceAccount with a matching name in the given list (and matches all other filters) is returned.
	*/
	Names []string

	/* PageMaxSize.

	     Page size preference (i.e. how many items are returned in the next
	page). If zero, the server will use a default. Regardless of what size
	is given, clients must never assume how many items will be returned.

	     Format: int64
	*/
	PageMaxSize *int64

	/* PageToken.

	     Opaque page token identifying which page to request. An empty token
	identifies the first page.
	*/
	PageToken *string

	/* Slugs.

	   Filters results by slug, where any ServiceAccount with a matching slug in the given list (and matches all other filters) is returned.
	*/
	Slugs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list service accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServiceAccountsParams) WithDefaults() *ListServiceAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list service accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServiceAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list service accounts params
func (o *ListServiceAccountsParams) WithTimeout(timeout time.Duration) *ListServiceAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list service accounts params
func (o *ListServiceAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list service accounts params
func (o *ListServiceAccountsParams) WithContext(ctx context.Context) *ListServiceAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list service accounts params
func (o *ListServiceAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list service accounts params
func (o *ListServiceAccountsParams) WithHTTPClient(client *http.Client) *ListServiceAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list service accounts params
func (o *ListServiceAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNames adds the names to the list service accounts params
func (o *ListServiceAccountsParams) WithNames(names []string) *ListServiceAccountsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the list service accounts params
func (o *ListServiceAccountsParams) SetNames(names []string) {
	o.Names = names
}

// WithPageMaxSize adds the pageMaxSize to the list service accounts params
func (o *ListServiceAccountsParams) WithPageMaxSize(pageMaxSize *int64) *ListServiceAccountsParams {
	o.SetPageMaxSize(pageMaxSize)
	return o
}

// SetPageMaxSize adds the pageMaxSize to the list service accounts params
func (o *ListServiceAccountsParams) SetPageMaxSize(pageMaxSize *int64) {
	o.PageMaxSize = pageMaxSize
}

// WithPageToken adds the pageToken to the list service accounts params
func (o *ListServiceAccountsParams) WithPageToken(pageToken *string) *ListServiceAccountsParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list service accounts params
func (o *ListServiceAccountsParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithSlugs adds the slugs to the list service accounts params
func (o *ListServiceAccountsParams) WithSlugs(slugs []string) *ListServiceAccountsParams {
	o.SetSlugs(slugs)
	return o
}

// SetSlugs adds the slugs to the list service accounts params
func (o *ListServiceAccountsParams) SetSlugs(slugs []string) {
	o.Slugs = slugs
}

// WriteToRequest writes these params to a swagger request
func (o *ListServiceAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Names != nil {

		// binding items for names
		joinedNames := o.bindParamNames(reg)

		// query array param names
		if err := r.SetQueryParam("names", joinedNames...); err != nil {
			return err
		}
	}

	if o.PageMaxSize != nil {

		// query param page.max_size
		var qrPageMaxSize int64

		if o.PageMaxSize != nil {
			qrPageMaxSize = *o.PageMaxSize
		}
		qPageMaxSize := swag.FormatInt64(qrPageMaxSize)
		if qPageMaxSize != "" {

			if err := r.SetQueryParam("page.max_size", qPageMaxSize); err != nil {
				return err
			}
		}
	}

	if o.PageToken != nil {

		// query param page.token
		var qrPageToken string

		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {

			if err := r.SetQueryParam("page.token", qPageToken); err != nil {
				return err
			}
		}
	}

	if o.Slugs != nil {

		// binding items for slugs
		joinedSlugs := o.bindParamSlugs(reg)

		// query array param slugs
		if err := r.SetQueryParam("slugs", joinedSlugs...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListServiceAccounts binds the parameter names
func (o *ListServiceAccountsParams) bindParamNames(formats strfmt.Registry) []string {
	namesIR := o.Names

	var namesIC []string
	for _, namesIIR := range namesIR { // explode []string

		namesIIV := namesIIR // string as string
		namesIC = append(namesIC, namesIIV)
	}

	// items.CollectionFormat: "multi"
	namesIS := swag.JoinByFormat(namesIC, "multi")

	return namesIS
}

// bindParamListServiceAccounts binds the parameter slugs
func (o *ListServiceAccountsParams) bindParamSlugs(formats strfmt.Registry) []string {
	slugsIR := o.Slugs

	var slugsIC []string
	for _, slugsIIR := range slugsIR { // explode []string

		slugsIIV := slugsIIR // string as string
		slugsIC = append(slugsIC, slugsIIV)
	}

	// items.CollectionFormat: "multi"
	slugsIS := swag.JoinByFormat(slugsIC, "multi")

	return slugsIS
}
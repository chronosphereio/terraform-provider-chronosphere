// Code generated by go-swagger; DO NOT EDIT.

package s_l_o

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

// NewListSLOsParams creates a new ListSLOsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSLOsParams() *ListSLOsParams {
	return &ListSLOsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSLOsParamsWithTimeout creates a new ListSLOsParams object
// with the ability to set a timeout on a request.
func NewListSLOsParamsWithTimeout(timeout time.Duration) *ListSLOsParams {
	return &ListSLOsParams{
		timeout: timeout,
	}
}

// NewListSLOsParamsWithContext creates a new ListSLOsParams object
// with the ability to set a context for a request.
func NewListSLOsParamsWithContext(ctx context.Context) *ListSLOsParams {
	return &ListSLOsParams{
		Context: ctx,
	}
}

// NewListSLOsParamsWithHTTPClient creates a new ListSLOsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSLOsParamsWithHTTPClient(client *http.Client) *ListSLOsParams {
	return &ListSLOsParams{
		HTTPClient: client,
	}
}

/*
ListSLOsParams contains all the parameters to send to the API endpoint

	for the list s l os operation.

	Typically these are written to a http.Request.
*/
type ListSLOsParams struct {

	// CollectionSlugs.
	CollectionSlugs []string

	/* Names.

	   Filters results by name, where any SLO with a matching name in the given list (and matches all other filters) is returned.
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

	// ServiceSlugs.
	ServiceSlugs []string

	/* Slugs.

	   Filters results by slug, where any SLO with a matching slug in the given list (and matches all other filters) is returned.
	*/
	Slugs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list s l os params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSLOsParams) WithDefaults() *ListSLOsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list s l os params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSLOsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list s l os params
func (o *ListSLOsParams) WithTimeout(timeout time.Duration) *ListSLOsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list s l os params
func (o *ListSLOsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list s l os params
func (o *ListSLOsParams) WithContext(ctx context.Context) *ListSLOsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list s l os params
func (o *ListSLOsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list s l os params
func (o *ListSLOsParams) WithHTTPClient(client *http.Client) *ListSLOsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list s l os params
func (o *ListSLOsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectionSlugs adds the collectionSlugs to the list s l os params
func (o *ListSLOsParams) WithCollectionSlugs(collectionSlugs []string) *ListSLOsParams {
	o.SetCollectionSlugs(collectionSlugs)
	return o
}

// SetCollectionSlugs adds the collectionSlugs to the list s l os params
func (o *ListSLOsParams) SetCollectionSlugs(collectionSlugs []string) {
	o.CollectionSlugs = collectionSlugs
}

// WithNames adds the names to the list s l os params
func (o *ListSLOsParams) WithNames(names []string) *ListSLOsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the list s l os params
func (o *ListSLOsParams) SetNames(names []string) {
	o.Names = names
}

// WithPageMaxSize adds the pageMaxSize to the list s l os params
func (o *ListSLOsParams) WithPageMaxSize(pageMaxSize *int64) *ListSLOsParams {
	o.SetPageMaxSize(pageMaxSize)
	return o
}

// SetPageMaxSize adds the pageMaxSize to the list s l os params
func (o *ListSLOsParams) SetPageMaxSize(pageMaxSize *int64) {
	o.PageMaxSize = pageMaxSize
}

// WithPageToken adds the pageToken to the list s l os params
func (o *ListSLOsParams) WithPageToken(pageToken *string) *ListSLOsParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list s l os params
func (o *ListSLOsParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithServiceSlugs adds the serviceSlugs to the list s l os params
func (o *ListSLOsParams) WithServiceSlugs(serviceSlugs []string) *ListSLOsParams {
	o.SetServiceSlugs(serviceSlugs)
	return o
}

// SetServiceSlugs adds the serviceSlugs to the list s l os params
func (o *ListSLOsParams) SetServiceSlugs(serviceSlugs []string) {
	o.ServiceSlugs = serviceSlugs
}

// WithSlugs adds the slugs to the list s l os params
func (o *ListSLOsParams) WithSlugs(slugs []string) *ListSLOsParams {
	o.SetSlugs(slugs)
	return o
}

// SetSlugs adds the slugs to the list s l os params
func (o *ListSLOsParams) SetSlugs(slugs []string) {
	o.Slugs = slugs
}

// WriteToRequest writes these params to a swagger request
func (o *ListSLOsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CollectionSlugs != nil {

		// binding items for collection_slugs
		joinedCollectionSlugs := o.bindParamCollectionSlugs(reg)

		// query array param collection_slugs
		if err := r.SetQueryParam("collection_slugs", joinedCollectionSlugs...); err != nil {
			return err
		}
	}

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

	if o.ServiceSlugs != nil {

		// binding items for service_slugs
		joinedServiceSlugs := o.bindParamServiceSlugs(reg)

		// query array param service_slugs
		if err := r.SetQueryParam("service_slugs", joinedServiceSlugs...); err != nil {
			return err
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

// bindParamListSLOs binds the parameter collection_slugs
func (o *ListSLOsParams) bindParamCollectionSlugs(formats strfmt.Registry) []string {
	collectionSlugsIR := o.CollectionSlugs

	var collectionSlugsIC []string
	for _, collectionSlugsIIR := range collectionSlugsIR { // explode []string

		collectionSlugsIIV := collectionSlugsIIR // string as string
		collectionSlugsIC = append(collectionSlugsIC, collectionSlugsIIV)
	}

	// items.CollectionFormat: "multi"
	collectionSlugsIS := swag.JoinByFormat(collectionSlugsIC, "multi")

	return collectionSlugsIS
}

// bindParamListSLOs binds the parameter names
func (o *ListSLOsParams) bindParamNames(formats strfmt.Registry) []string {
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

// bindParamListSLOs binds the parameter service_slugs
func (o *ListSLOsParams) bindParamServiceSlugs(formats strfmt.Registry) []string {
	serviceSlugsIR := o.ServiceSlugs

	var serviceSlugsIC []string
	for _, serviceSlugsIIR := range serviceSlugsIR { // explode []string

		serviceSlugsIIV := serviceSlugsIIR // string as string
		serviceSlugsIC = append(serviceSlugsIC, serviceSlugsIIV)
	}

	// items.CollectionFormat: "multi"
	serviceSlugsIS := swag.JoinByFormat(serviceSlugsIC, "multi")

	return serviceSlugsIS
}

// bindParamListSLOs binds the parameter slugs
func (o *ListSLOsParams) bindParamSlugs(formats strfmt.Registry) []string {
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
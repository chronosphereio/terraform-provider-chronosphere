// Code generated by go-swagger; DO NOT EDIT.

package mapping_rule

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

// NewListMappingRulesParams creates a new ListMappingRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListMappingRulesParams() *ListMappingRulesParams {
	return &ListMappingRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListMappingRulesParamsWithTimeout creates a new ListMappingRulesParams object
// with the ability to set a timeout on a request.
func NewListMappingRulesParamsWithTimeout(timeout time.Duration) *ListMappingRulesParams {
	return &ListMappingRulesParams{
		timeout: timeout,
	}
}

// NewListMappingRulesParamsWithContext creates a new ListMappingRulesParams object
// with the ability to set a context for a request.
func NewListMappingRulesParamsWithContext(ctx context.Context) *ListMappingRulesParams {
	return &ListMappingRulesParams{
		Context: ctx,
	}
}

// NewListMappingRulesParamsWithHTTPClient creates a new ListMappingRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListMappingRulesParamsWithHTTPClient(client *http.Client) *ListMappingRulesParams {
	return &ListMappingRulesParams{
		HTTPClient: client,
	}
}

/*
ListMappingRulesParams contains all the parameters to send to the API endpoint

	for the list mapping rules operation.

	Typically these are written to a http.Request.
*/
type ListMappingRulesParams struct {

	/* BucketSlugs.

	   Filters results by bucket_slug, where any MappingRule with a matching bucket_slug in the given list (and matches all other filters) is returned.
	*/
	BucketSlugs []string

	/* Names.

	   Filters results by name, where any MappingRule with a matching name in the given list (and matches all other filters) is returned.
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

	   Filters results by slug, where any MappingRule with a matching slug in the given list (and matches all other filters) is returned.
	*/
	Slugs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list mapping rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMappingRulesParams) WithDefaults() *ListMappingRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list mapping rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMappingRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list mapping rules params
func (o *ListMappingRulesParams) WithTimeout(timeout time.Duration) *ListMappingRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list mapping rules params
func (o *ListMappingRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list mapping rules params
func (o *ListMappingRulesParams) WithContext(ctx context.Context) *ListMappingRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list mapping rules params
func (o *ListMappingRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list mapping rules params
func (o *ListMappingRulesParams) WithHTTPClient(client *http.Client) *ListMappingRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list mapping rules params
func (o *ListMappingRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketSlugs adds the bucketSlugs to the list mapping rules params
func (o *ListMappingRulesParams) WithBucketSlugs(bucketSlugs []string) *ListMappingRulesParams {
	o.SetBucketSlugs(bucketSlugs)
	return o
}

// SetBucketSlugs adds the bucketSlugs to the list mapping rules params
func (o *ListMappingRulesParams) SetBucketSlugs(bucketSlugs []string) {
	o.BucketSlugs = bucketSlugs
}

// WithNames adds the names to the list mapping rules params
func (o *ListMappingRulesParams) WithNames(names []string) *ListMappingRulesParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the list mapping rules params
func (o *ListMappingRulesParams) SetNames(names []string) {
	o.Names = names
}

// WithPageMaxSize adds the pageMaxSize to the list mapping rules params
func (o *ListMappingRulesParams) WithPageMaxSize(pageMaxSize *int64) *ListMappingRulesParams {
	o.SetPageMaxSize(pageMaxSize)
	return o
}

// SetPageMaxSize adds the pageMaxSize to the list mapping rules params
func (o *ListMappingRulesParams) SetPageMaxSize(pageMaxSize *int64) {
	o.PageMaxSize = pageMaxSize
}

// WithPageToken adds the pageToken to the list mapping rules params
func (o *ListMappingRulesParams) WithPageToken(pageToken *string) *ListMappingRulesParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list mapping rules params
func (o *ListMappingRulesParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithSlugs adds the slugs to the list mapping rules params
func (o *ListMappingRulesParams) WithSlugs(slugs []string) *ListMappingRulesParams {
	o.SetSlugs(slugs)
	return o
}

// SetSlugs adds the slugs to the list mapping rules params
func (o *ListMappingRulesParams) SetSlugs(slugs []string) {
	o.Slugs = slugs
}

// WriteToRequest writes these params to a swagger request
func (o *ListMappingRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BucketSlugs != nil {

		// binding items for bucket_slugs
		joinedBucketSlugs := o.bindParamBucketSlugs(reg)

		// query array param bucket_slugs
		if err := r.SetQueryParam("bucket_slugs", joinedBucketSlugs...); err != nil {
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

// bindParamListMappingRules binds the parameter bucket_slugs
func (o *ListMappingRulesParams) bindParamBucketSlugs(formats strfmt.Registry) []string {
	bucketSlugsIR := o.BucketSlugs

	var bucketSlugsIC []string
	for _, bucketSlugsIIR := range bucketSlugsIR { // explode []string

		bucketSlugsIIV := bucketSlugsIIR // string as string
		bucketSlugsIC = append(bucketSlugsIC, bucketSlugsIIV)
	}

	// items.CollectionFormat: "multi"
	bucketSlugsIS := swag.JoinByFormat(bucketSlugsIC, "multi")

	return bucketSlugsIS
}

// bindParamListMappingRules binds the parameter names
func (o *ListMappingRulesParams) bindParamNames(formats strfmt.Registry) []string {
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

// bindParamListMappingRules binds the parameter slugs
func (o *ListMappingRulesParams) bindParamSlugs(formats strfmt.Registry) []string {
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

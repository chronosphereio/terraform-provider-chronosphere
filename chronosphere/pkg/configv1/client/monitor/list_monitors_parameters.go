// Code generated by go-swagger; DO NOT EDIT.

package monitor

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

// NewListMonitorsParams creates a new ListMonitorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListMonitorsParams() *ListMonitorsParams {
	return &ListMonitorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListMonitorsParamsWithTimeout creates a new ListMonitorsParams object
// with the ability to set a timeout on a request.
func NewListMonitorsParamsWithTimeout(timeout time.Duration) *ListMonitorsParams {
	return &ListMonitorsParams{
		timeout: timeout,
	}
}

// NewListMonitorsParamsWithContext creates a new ListMonitorsParams object
// with the ability to set a context for a request.
func NewListMonitorsParamsWithContext(ctx context.Context) *ListMonitorsParams {
	return &ListMonitorsParams{
		Context: ctx,
	}
}

// NewListMonitorsParamsWithHTTPClient creates a new ListMonitorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListMonitorsParamsWithHTTPClient(client *http.Client) *ListMonitorsParams {
	return &ListMonitorsParams{
		HTTPClient: client,
	}
}

/*
ListMonitorsParams contains all the parameters to send to the API endpoint

	for the list monitors operation.

	Typically these are written to a http.Request.
*/
type ListMonitorsParams struct {

	/* BucketSlugs.

	   Filters results by bucket_slug, where any Monitor with a matching bucket_slug in the given list (and matches all other filters) is returned.
	*/
	BucketSlugs []string

	/* CollectionSlugs.

	   Filters results by collection_slug, where any Monitor with a matching collection_slug in the given list (and matches all other filters) is returned.
	*/
	CollectionSlugs []string

	/* Names.

	   Filters results by name, where any Monitor with a matching name in the given list (and matches all other filters) is returned.
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

	   Filters results by slug, where any Monitor with a matching slug in the given list (and matches all other filters) is returned.
	*/
	Slugs []string

	/* TeamSlugs.

	   Filter returned monitors by the teams that own the collections that they belong to.
	*/
	TeamSlugs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list monitors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMonitorsParams) WithDefaults() *ListMonitorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list monitors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMonitorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list monitors params
func (o *ListMonitorsParams) WithTimeout(timeout time.Duration) *ListMonitorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list monitors params
func (o *ListMonitorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list monitors params
func (o *ListMonitorsParams) WithContext(ctx context.Context) *ListMonitorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list monitors params
func (o *ListMonitorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list monitors params
func (o *ListMonitorsParams) WithHTTPClient(client *http.Client) *ListMonitorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list monitors params
func (o *ListMonitorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketSlugs adds the bucketSlugs to the list monitors params
func (o *ListMonitorsParams) WithBucketSlugs(bucketSlugs []string) *ListMonitorsParams {
	o.SetBucketSlugs(bucketSlugs)
	return o
}

// SetBucketSlugs adds the bucketSlugs to the list monitors params
func (o *ListMonitorsParams) SetBucketSlugs(bucketSlugs []string) {
	o.BucketSlugs = bucketSlugs
}

// WithCollectionSlugs adds the collectionSlugs to the list monitors params
func (o *ListMonitorsParams) WithCollectionSlugs(collectionSlugs []string) *ListMonitorsParams {
	o.SetCollectionSlugs(collectionSlugs)
	return o
}

// SetCollectionSlugs adds the collectionSlugs to the list monitors params
func (o *ListMonitorsParams) SetCollectionSlugs(collectionSlugs []string) {
	o.CollectionSlugs = collectionSlugs
}

// WithNames adds the names to the list monitors params
func (o *ListMonitorsParams) WithNames(names []string) *ListMonitorsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the list monitors params
func (o *ListMonitorsParams) SetNames(names []string) {
	o.Names = names
}

// WithPageMaxSize adds the pageMaxSize to the list monitors params
func (o *ListMonitorsParams) WithPageMaxSize(pageMaxSize *int64) *ListMonitorsParams {
	o.SetPageMaxSize(pageMaxSize)
	return o
}

// SetPageMaxSize adds the pageMaxSize to the list monitors params
func (o *ListMonitorsParams) SetPageMaxSize(pageMaxSize *int64) {
	o.PageMaxSize = pageMaxSize
}

// WithPageToken adds the pageToken to the list monitors params
func (o *ListMonitorsParams) WithPageToken(pageToken *string) *ListMonitorsParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list monitors params
func (o *ListMonitorsParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithSlugs adds the slugs to the list monitors params
func (o *ListMonitorsParams) WithSlugs(slugs []string) *ListMonitorsParams {
	o.SetSlugs(slugs)
	return o
}

// SetSlugs adds the slugs to the list monitors params
func (o *ListMonitorsParams) SetSlugs(slugs []string) {
	o.Slugs = slugs
}

// WithTeamSlugs adds the teamSlugs to the list monitors params
func (o *ListMonitorsParams) WithTeamSlugs(teamSlugs []string) *ListMonitorsParams {
	o.SetTeamSlugs(teamSlugs)
	return o
}

// SetTeamSlugs adds the teamSlugs to the list monitors params
func (o *ListMonitorsParams) SetTeamSlugs(teamSlugs []string) {
	o.TeamSlugs = teamSlugs
}

// WriteToRequest writes these params to a swagger request
func (o *ListMonitorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Slugs != nil {

		// binding items for slugs
		joinedSlugs := o.bindParamSlugs(reg)

		// query array param slugs
		if err := r.SetQueryParam("slugs", joinedSlugs...); err != nil {
			return err
		}
	}

	if o.TeamSlugs != nil {

		// binding items for team_slugs
		joinedTeamSlugs := o.bindParamTeamSlugs(reg)

		// query array param team_slugs
		if err := r.SetQueryParam("team_slugs", joinedTeamSlugs...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListMonitors binds the parameter bucket_slugs
func (o *ListMonitorsParams) bindParamBucketSlugs(formats strfmt.Registry) []string {
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

// bindParamListMonitors binds the parameter collection_slugs
func (o *ListMonitorsParams) bindParamCollectionSlugs(formats strfmt.Registry) []string {
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

// bindParamListMonitors binds the parameter names
func (o *ListMonitorsParams) bindParamNames(formats strfmt.Registry) []string {
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

// bindParamListMonitors binds the parameter slugs
func (o *ListMonitorsParams) bindParamSlugs(formats strfmt.Registry) []string {
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

// bindParamListMonitors binds the parameter team_slugs
func (o *ListMonitorsParams) bindParamTeamSlugs(formats strfmt.Registry) []string {
	teamSlugsIR := o.TeamSlugs

	var teamSlugsIC []string
	for _, teamSlugsIIR := range teamSlugsIR { // explode []string

		teamSlugsIIV := teamSlugsIIR // string as string
		teamSlugsIC = append(teamSlugsIC, teamSlugsIIV)
	}

	// items.CollectionFormat: "multi"
	teamSlugsIS := swag.JoinByFormat(teamSlugsIC, "multi")

	return teamSlugsIS
}
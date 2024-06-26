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
	"github.com/go-openapi/swag"
)

// NewListGrafanaDashboardsParams creates a new ListGrafanaDashboardsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListGrafanaDashboardsParams() *ListGrafanaDashboardsParams {
	return &ListGrafanaDashboardsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListGrafanaDashboardsParamsWithTimeout creates a new ListGrafanaDashboardsParams object
// with the ability to set a timeout on a request.
func NewListGrafanaDashboardsParamsWithTimeout(timeout time.Duration) *ListGrafanaDashboardsParams {
	return &ListGrafanaDashboardsParams{
		timeout: timeout,
	}
}

// NewListGrafanaDashboardsParamsWithContext creates a new ListGrafanaDashboardsParams object
// with the ability to set a context for a request.
func NewListGrafanaDashboardsParamsWithContext(ctx context.Context) *ListGrafanaDashboardsParams {
	return &ListGrafanaDashboardsParams{
		Context: ctx,
	}
}

// NewListGrafanaDashboardsParamsWithHTTPClient creates a new ListGrafanaDashboardsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListGrafanaDashboardsParamsWithHTTPClient(client *http.Client) *ListGrafanaDashboardsParams {
	return &ListGrafanaDashboardsParams{
		HTTPClient: client,
	}
}

/*
ListGrafanaDashboardsParams contains all the parameters to send to the API endpoint

	for the list grafana dashboards operation.

	Typically these are written to a http.Request.
*/
type ListGrafanaDashboardsParams struct {

	/* BucketSlugs.

	   Filters results by bucket_slug, where any GrafanaDashboard with a matching bucket_slug in the given list (and matches all other filters) is returned.
	*/
	BucketSlugs []string

	/* CollectionSlugs.

	   Filters results by collection_slug, where any GrafanaDashboard with a matching collection_slug in the given list (and matches all other filters) is returned.
	*/
	CollectionSlugs []string

	/* IncludeDashboardJSON.

	     Optional flag to populate the dashboard_json of the returned dashboards.
	By default, dashboard_json will be left empty.
	*/
	IncludeDashboardJSON *bool

	/* Names.

	   Filters results by name, where any GrafanaDashboard with a matching name in the given list (and matches all other filters) is returned.
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

	   Filters results by slug, where any GrafanaDashboard with a matching slug in the given list (and matches all other filters) is returned.
	*/
	Slugs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list grafana dashboards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGrafanaDashboardsParams) WithDefaults() *ListGrafanaDashboardsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list grafana dashboards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGrafanaDashboardsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) WithTimeout(timeout time.Duration) *ListGrafanaDashboardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) WithContext(ctx context.Context) *ListGrafanaDashboardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) WithHTTPClient(client *http.Client) *ListGrafanaDashboardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketSlugs adds the bucketSlugs to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) WithBucketSlugs(bucketSlugs []string) *ListGrafanaDashboardsParams {
	o.SetBucketSlugs(bucketSlugs)
	return o
}

// SetBucketSlugs adds the bucketSlugs to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) SetBucketSlugs(bucketSlugs []string) {
	o.BucketSlugs = bucketSlugs
}

// WithCollectionSlugs adds the collectionSlugs to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) WithCollectionSlugs(collectionSlugs []string) *ListGrafanaDashboardsParams {
	o.SetCollectionSlugs(collectionSlugs)
	return o
}

// SetCollectionSlugs adds the collectionSlugs to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) SetCollectionSlugs(collectionSlugs []string) {
	o.CollectionSlugs = collectionSlugs
}

// WithIncludeDashboardJSON adds the includeDashboardJSON to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) WithIncludeDashboardJSON(includeDashboardJSON *bool) *ListGrafanaDashboardsParams {
	o.SetIncludeDashboardJSON(includeDashboardJSON)
	return o
}

// SetIncludeDashboardJSON adds the includeDashboardJson to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) SetIncludeDashboardJSON(includeDashboardJSON *bool) {
	o.IncludeDashboardJSON = includeDashboardJSON
}

// WithNames adds the names to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) WithNames(names []string) *ListGrafanaDashboardsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) SetNames(names []string) {
	o.Names = names
}

// WithPageMaxSize adds the pageMaxSize to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) WithPageMaxSize(pageMaxSize *int64) *ListGrafanaDashboardsParams {
	o.SetPageMaxSize(pageMaxSize)
	return o
}

// SetPageMaxSize adds the pageMaxSize to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) SetPageMaxSize(pageMaxSize *int64) {
	o.PageMaxSize = pageMaxSize
}

// WithPageToken adds the pageToken to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) WithPageToken(pageToken *string) *ListGrafanaDashboardsParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithSlugs adds the slugs to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) WithSlugs(slugs []string) *ListGrafanaDashboardsParams {
	o.SetSlugs(slugs)
	return o
}

// SetSlugs adds the slugs to the list grafana dashboards params
func (o *ListGrafanaDashboardsParams) SetSlugs(slugs []string) {
	o.Slugs = slugs
}

// WriteToRequest writes these params to a swagger request
func (o *ListGrafanaDashboardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IncludeDashboardJSON != nil {

		// query param include_dashboard_json
		var qrIncludeDashboardJSON bool

		if o.IncludeDashboardJSON != nil {
			qrIncludeDashboardJSON = *o.IncludeDashboardJSON
		}
		qIncludeDashboardJSON := swag.FormatBool(qrIncludeDashboardJSON)
		if qIncludeDashboardJSON != "" {

			if err := r.SetQueryParam("include_dashboard_json", qIncludeDashboardJSON); err != nil {
				return err
			}
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

// bindParamListGrafanaDashboards binds the parameter bucket_slugs
func (o *ListGrafanaDashboardsParams) bindParamBucketSlugs(formats strfmt.Registry) []string {
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

// bindParamListGrafanaDashboards binds the parameter collection_slugs
func (o *ListGrafanaDashboardsParams) bindParamCollectionSlugs(formats strfmt.Registry) []string {
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

// bindParamListGrafanaDashboards binds the parameter names
func (o *ListGrafanaDashboardsParams) bindParamNames(formats strfmt.Registry) []string {
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

// bindParamListGrafanaDashboards binds the parameter slugs
func (o *ListGrafanaDashboardsParams) bindParamSlugs(formats strfmt.Registry) []string {
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

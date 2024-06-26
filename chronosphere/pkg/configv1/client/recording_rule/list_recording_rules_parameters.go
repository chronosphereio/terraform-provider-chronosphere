// Code generated by go-swagger; DO NOT EDIT.

package recording_rule

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

// NewListRecordingRulesParams creates a new ListRecordingRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListRecordingRulesParams() *ListRecordingRulesParams {
	return &ListRecordingRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListRecordingRulesParamsWithTimeout creates a new ListRecordingRulesParams object
// with the ability to set a timeout on a request.
func NewListRecordingRulesParamsWithTimeout(timeout time.Duration) *ListRecordingRulesParams {
	return &ListRecordingRulesParams{
		timeout: timeout,
	}
}

// NewListRecordingRulesParamsWithContext creates a new ListRecordingRulesParams object
// with the ability to set a context for a request.
func NewListRecordingRulesParamsWithContext(ctx context.Context) *ListRecordingRulesParams {
	return &ListRecordingRulesParams{
		Context: ctx,
	}
}

// NewListRecordingRulesParamsWithHTTPClient creates a new ListRecordingRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListRecordingRulesParamsWithHTTPClient(client *http.Client) *ListRecordingRulesParams {
	return &ListRecordingRulesParams{
		HTTPClient: client,
	}
}

/*
ListRecordingRulesParams contains all the parameters to send to the API endpoint

	for the list recording rules operation.

	Typically these are written to a http.Request.
*/
type ListRecordingRulesParams struct {

	/* BucketSlugs.

	   The execution_groups filter cannot be used when a bucket_slug filter is provided.
	*/
	BucketSlugs []string

	/* ExecutionGroups.

	   The bucket_slugs filter cannot be used when an execution_group filter is provided.
	*/
	ExecutionGroups []string

	/* Names.

	   Filters results by name, where any RecordingRule with a matching name in the given list (and matches all other filters) is returned.
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

	   Filters results by slug, where any RecordingRule with a matching slug in the given list (and matches all other filters) is returned.
	*/
	Slugs []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list recording rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRecordingRulesParams) WithDefaults() *ListRecordingRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list recording rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRecordingRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list recording rules params
func (o *ListRecordingRulesParams) WithTimeout(timeout time.Duration) *ListRecordingRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list recording rules params
func (o *ListRecordingRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list recording rules params
func (o *ListRecordingRulesParams) WithContext(ctx context.Context) *ListRecordingRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list recording rules params
func (o *ListRecordingRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list recording rules params
func (o *ListRecordingRulesParams) WithHTTPClient(client *http.Client) *ListRecordingRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list recording rules params
func (o *ListRecordingRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketSlugs adds the bucketSlugs to the list recording rules params
func (o *ListRecordingRulesParams) WithBucketSlugs(bucketSlugs []string) *ListRecordingRulesParams {
	o.SetBucketSlugs(bucketSlugs)
	return o
}

// SetBucketSlugs adds the bucketSlugs to the list recording rules params
func (o *ListRecordingRulesParams) SetBucketSlugs(bucketSlugs []string) {
	o.BucketSlugs = bucketSlugs
}

// WithExecutionGroups adds the executionGroups to the list recording rules params
func (o *ListRecordingRulesParams) WithExecutionGroups(executionGroups []string) *ListRecordingRulesParams {
	o.SetExecutionGroups(executionGroups)
	return o
}

// SetExecutionGroups adds the executionGroups to the list recording rules params
func (o *ListRecordingRulesParams) SetExecutionGroups(executionGroups []string) {
	o.ExecutionGroups = executionGroups
}

// WithNames adds the names to the list recording rules params
func (o *ListRecordingRulesParams) WithNames(names []string) *ListRecordingRulesParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the list recording rules params
func (o *ListRecordingRulesParams) SetNames(names []string) {
	o.Names = names
}

// WithPageMaxSize adds the pageMaxSize to the list recording rules params
func (o *ListRecordingRulesParams) WithPageMaxSize(pageMaxSize *int64) *ListRecordingRulesParams {
	o.SetPageMaxSize(pageMaxSize)
	return o
}

// SetPageMaxSize adds the pageMaxSize to the list recording rules params
func (o *ListRecordingRulesParams) SetPageMaxSize(pageMaxSize *int64) {
	o.PageMaxSize = pageMaxSize
}

// WithPageToken adds the pageToken to the list recording rules params
func (o *ListRecordingRulesParams) WithPageToken(pageToken *string) *ListRecordingRulesParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list recording rules params
func (o *ListRecordingRulesParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithSlugs adds the slugs to the list recording rules params
func (o *ListRecordingRulesParams) WithSlugs(slugs []string) *ListRecordingRulesParams {
	o.SetSlugs(slugs)
	return o
}

// SetSlugs adds the slugs to the list recording rules params
func (o *ListRecordingRulesParams) SetSlugs(slugs []string) {
	o.Slugs = slugs
}

// WriteToRequest writes these params to a swagger request
func (o *ListRecordingRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ExecutionGroups != nil {

		// binding items for execution_groups
		joinedExecutionGroups := o.bindParamExecutionGroups(reg)

		// query array param execution_groups
		if err := r.SetQueryParam("execution_groups", joinedExecutionGroups...); err != nil {
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

// bindParamListRecordingRules binds the parameter bucket_slugs
func (o *ListRecordingRulesParams) bindParamBucketSlugs(formats strfmt.Registry) []string {
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

// bindParamListRecordingRules binds the parameter execution_groups
func (o *ListRecordingRulesParams) bindParamExecutionGroups(formats strfmt.Registry) []string {
	executionGroupsIR := o.ExecutionGroups

	var executionGroupsIC []string
	for _, executionGroupsIIR := range executionGroupsIR { // explode []string

		executionGroupsIIV := executionGroupsIIR // string as string
		executionGroupsIC = append(executionGroupsIC, executionGroupsIIV)
	}

	// items.CollectionFormat: "multi"
	executionGroupsIS := swag.JoinByFormat(executionGroupsIC, "multi")

	return executionGroupsIS
}

// bindParamListRecordingRules binds the parameter names
func (o *ListRecordingRulesParams) bindParamNames(formats strfmt.Registry) []string {
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

// bindParamListRecordingRules binds the parameter slugs
func (o *ListRecordingRulesParams) bindParamSlugs(formats strfmt.Registry) []string {
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

// Code generated by go-swagger; DO NOT EDIT.

package link_template

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

// NewListLinkTemplatesParams creates a new ListLinkTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListLinkTemplatesParams() *ListLinkTemplatesParams {
	return &ListLinkTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListLinkTemplatesParamsWithTimeout creates a new ListLinkTemplatesParams object
// with the ability to set a timeout on a request.
func NewListLinkTemplatesParamsWithTimeout(timeout time.Duration) *ListLinkTemplatesParams {
	return &ListLinkTemplatesParams{
		timeout: timeout,
	}
}

// NewListLinkTemplatesParamsWithContext creates a new ListLinkTemplatesParams object
// with the ability to set a context for a request.
func NewListLinkTemplatesParamsWithContext(ctx context.Context) *ListLinkTemplatesParams {
	return &ListLinkTemplatesParams{
		Context: ctx,
	}
}

// NewListLinkTemplatesParamsWithHTTPClient creates a new ListLinkTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListLinkTemplatesParamsWithHTTPClient(client *http.Client) *ListLinkTemplatesParams {
	return &ListLinkTemplatesParams{
		HTTPClient: client,
	}
}

/*
ListLinkTemplatesParams contains all the parameters to send to the API endpoint

	for the list link templates operation.

	Typically these are written to a http.Request.
*/
type ListLinkTemplatesParams struct {

	/* Names.

	   Filters results by name, where any LinkTemplate with a matching name in the given list (and matches all other filters) is returned.
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

	   Filters results by slug, where any LinkTemplate with a matching slug in the given list (and matches all other filters) is returned.
	*/
	Slugs []string

	/* UIComponents.

	   Filters results by ui_component, where any LinkTemplate with a matching ui_component in the given list (and matches all other filters) is returned.
	*/
	UIComponents []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list link templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLinkTemplatesParams) WithDefaults() *ListLinkTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list link templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListLinkTemplatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list link templates params
func (o *ListLinkTemplatesParams) WithTimeout(timeout time.Duration) *ListLinkTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list link templates params
func (o *ListLinkTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list link templates params
func (o *ListLinkTemplatesParams) WithContext(ctx context.Context) *ListLinkTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list link templates params
func (o *ListLinkTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list link templates params
func (o *ListLinkTemplatesParams) WithHTTPClient(client *http.Client) *ListLinkTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list link templates params
func (o *ListLinkTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNames adds the names to the list link templates params
func (o *ListLinkTemplatesParams) WithNames(names []string) *ListLinkTemplatesParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the list link templates params
func (o *ListLinkTemplatesParams) SetNames(names []string) {
	o.Names = names
}

// WithPageMaxSize adds the pageMaxSize to the list link templates params
func (o *ListLinkTemplatesParams) WithPageMaxSize(pageMaxSize *int64) *ListLinkTemplatesParams {
	o.SetPageMaxSize(pageMaxSize)
	return o
}

// SetPageMaxSize adds the pageMaxSize to the list link templates params
func (o *ListLinkTemplatesParams) SetPageMaxSize(pageMaxSize *int64) {
	o.PageMaxSize = pageMaxSize
}

// WithPageToken adds the pageToken to the list link templates params
func (o *ListLinkTemplatesParams) WithPageToken(pageToken *string) *ListLinkTemplatesParams {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the list link templates params
func (o *ListLinkTemplatesParams) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithSlugs adds the slugs to the list link templates params
func (o *ListLinkTemplatesParams) WithSlugs(slugs []string) *ListLinkTemplatesParams {
	o.SetSlugs(slugs)
	return o
}

// SetSlugs adds the slugs to the list link templates params
func (o *ListLinkTemplatesParams) SetSlugs(slugs []string) {
	o.Slugs = slugs
}

// WithUIComponents adds the uIComponents to the list link templates params
func (o *ListLinkTemplatesParams) WithUIComponents(uIComponents []string) *ListLinkTemplatesParams {
	o.SetUIComponents(uIComponents)
	return o
}

// SetUIComponents adds the uiComponents to the list link templates params
func (o *ListLinkTemplatesParams) SetUIComponents(uIComponents []string) {
	o.UIComponents = uIComponents
}

// WriteToRequest writes these params to a swagger request
func (o *ListLinkTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.UIComponents != nil {

		// binding items for ui_components
		joinedUIComponents := o.bindParamUIComponents(reg)

		// query array param ui_components
		if err := r.SetQueryParam("ui_components", joinedUIComponents...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListLinkTemplates binds the parameter names
func (o *ListLinkTemplatesParams) bindParamNames(formats strfmt.Registry) []string {
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

// bindParamListLinkTemplates binds the parameter slugs
func (o *ListLinkTemplatesParams) bindParamSlugs(formats strfmt.Registry) []string {
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

// bindParamListLinkTemplates binds the parameter ui_components
func (o *ListLinkTemplatesParams) bindParamUIComponents(formats strfmt.Registry) []string {
	uIComponentsIR := o.UIComponents

	var uIComponentsIC []string
	for _, uIComponentsIIR := range uIComponentsIR { // explode []string

		uIComponentsIIV := uIComponentsIIR // string as string
		uIComponentsIC = append(uIComponentsIC, uIComponentsIIV)
	}

	// items.CollectionFormat: "multi"
	uIComponentsIS := swag.JoinByFormat(uIComponentsIC, "multi")

	return uIComponentsIS
}

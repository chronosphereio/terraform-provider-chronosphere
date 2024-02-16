// Code generated by go-swagger; DO NOT EDIT.

package grafana_dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new grafana dashboard API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for grafana dashboard API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateGrafanaDashboard(params *CreateGrafanaDashboardParams, opts ...ClientOption) (*CreateGrafanaDashboardOK, error)

	DeleteGrafanaDashboard(params *DeleteGrafanaDashboardParams, opts ...ClientOption) (*DeleteGrafanaDashboardOK, error)

	ListGrafanaDashboards(params *ListGrafanaDashboardsParams, opts ...ClientOption) (*ListGrafanaDashboardsOK, error)

	ReadGrafanaDashboard(params *ReadGrafanaDashboardParams, opts ...ClientOption) (*ReadGrafanaDashboardOK, error)

	UpdateGrafanaDashboard(params *UpdateGrafanaDashboardParams, opts ...ClientOption) (*UpdateGrafanaDashboardOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateGrafanaDashboard create grafana dashboard API
*/
func (a *Client) CreateGrafanaDashboard(params *CreateGrafanaDashboardParams, opts ...ClientOption) (*CreateGrafanaDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateGrafanaDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateGrafanaDashboard",
		Method:             "POST",
		PathPattern:        "/api/v1/config/grafana-dashboards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateGrafanaDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateGrafanaDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateGrafanaDashboardDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteGrafanaDashboard delete grafana dashboard API
*/
func (a *Client) DeleteGrafanaDashboard(params *DeleteGrafanaDashboardParams, opts ...ClientOption) (*DeleteGrafanaDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGrafanaDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteGrafanaDashboard",
		Method:             "DELETE",
		PathPattern:        "/api/v1/config/grafana-dashboards/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteGrafanaDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteGrafanaDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteGrafanaDashboardDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGrafanaDashboards list grafana dashboards API
*/
func (a *Client) ListGrafanaDashboards(params *ListGrafanaDashboardsParams, opts ...ClientOption) (*ListGrafanaDashboardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGrafanaDashboardsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListGrafanaDashboards",
		Method:             "GET",
		PathPattern:        "/api/v1/config/grafana-dashboards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListGrafanaDashboardsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGrafanaDashboardsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGrafanaDashboardsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadGrafanaDashboard read grafana dashboard API
*/
func (a *Client) ReadGrafanaDashboard(params *ReadGrafanaDashboardParams, opts ...ClientOption) (*ReadGrafanaDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadGrafanaDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadGrafanaDashboard",
		Method:             "GET",
		PathPattern:        "/api/v1/config/grafana-dashboards/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadGrafanaDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadGrafanaDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadGrafanaDashboardDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateGrafanaDashboard update grafana dashboard API
*/
func (a *Client) UpdateGrafanaDashboard(params *UpdateGrafanaDashboardParams, opts ...ClientOption) (*UpdateGrafanaDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGrafanaDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateGrafanaDashboard",
		Method:             "PUT",
		PathPattern:        "/api/v1/config/grafana-dashboards/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateGrafanaDashboardReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateGrafanaDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateGrafanaDashboardDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
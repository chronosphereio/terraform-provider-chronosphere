// Code generated by go-swagger; DO NOT EDIT.

package resource_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new resource pools API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for resource pools API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateResourcePools(params *CreateResourcePoolsParams, opts ...ClientOption) (*CreateResourcePoolsOK, error)

	DeleteResourcePools(params *DeleteResourcePoolsParams, opts ...ClientOption) (*DeleteResourcePoolsOK, error)

	ReadResourcePools(params *ReadResourcePoolsParams, opts ...ClientOption) (*ReadResourcePoolsOK, error)

	UpdateResourcePools(params *UpdateResourcePoolsParams, opts ...ClientOption) (*UpdateResourcePoolsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateResourcePools ResourcePools CRUD (subset for singleton objects)
*/
func (a *Client) CreateResourcePools(params *CreateResourcePoolsParams, opts ...ClientOption) (*CreateResourcePoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateResourcePoolsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateResourcePools",
		Method:             "POST",
		PathPattern:        "/api/v1/config/resource-pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateResourcePoolsReader{formats: a.formats},
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
	success, ok := result.(*CreateResourcePoolsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateResourcePoolsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteResourcePools delete resource pools API
*/
func (a *Client) DeleteResourcePools(params *DeleteResourcePoolsParams, opts ...ClientOption) (*DeleteResourcePoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteResourcePoolsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteResourcePools",
		Method:             "DELETE",
		PathPattern:        "/api/v1/config/resource-pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteResourcePoolsReader{formats: a.formats},
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
	success, ok := result.(*DeleteResourcePoolsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteResourcePoolsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadResourcePools read resource pools API
*/
func (a *Client) ReadResourcePools(params *ReadResourcePoolsParams, opts ...ClientOption) (*ReadResourcePoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadResourcePoolsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadResourcePools",
		Method:             "GET",
		PathPattern:        "/api/v1/config/resource-pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadResourcePoolsReader{formats: a.formats},
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
	success, ok := result.(*ReadResourcePoolsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadResourcePoolsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateResourcePools update resource pools API
*/
func (a *Client) UpdateResourcePools(params *UpdateResourcePoolsParams, opts ...ClientOption) (*UpdateResourcePoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateResourcePoolsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateResourcePools",
		Method:             "PUT",
		PathPattern:        "/api/v1/config/resource-pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateResourcePoolsReader{formats: a.formats},
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
	success, ok := result.(*UpdateResourcePoolsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateResourcePoolsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

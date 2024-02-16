// Code generated by go-swagger; DO NOT EDIT.

package service_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service account API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service account API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateServiceAccount(params *CreateServiceAccountParams, opts ...ClientOption) (*CreateServiceAccountOK, error)

	DeleteServiceAccount(params *DeleteServiceAccountParams, opts ...ClientOption) (*DeleteServiceAccountOK, error)

	ListServiceAccounts(params *ListServiceAccountsParams, opts ...ClientOption) (*ListServiceAccountsOK, error)

	ReadServiceAccount(params *ReadServiceAccountParams, opts ...ClientOption) (*ReadServiceAccountOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateServiceAccount create service account API
*/
func (a *Client) CreateServiceAccount(params *CreateServiceAccountParams, opts ...ClientOption) (*CreateServiceAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateServiceAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateServiceAccount",
		Method:             "POST",
		PathPattern:        "/api/v1/config/service-accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateServiceAccountReader{formats: a.formats},
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
	success, ok := result.(*CreateServiceAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateServiceAccountDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteServiceAccount delete service account API
*/
func (a *Client) DeleteServiceAccount(params *DeleteServiceAccountParams, opts ...ClientOption) (*DeleteServiceAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteServiceAccount",
		Method:             "DELETE",
		PathPattern:        "/api/v1/config/service-accounts/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteServiceAccountReader{formats: a.formats},
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
	success, ok := result.(*DeleteServiceAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteServiceAccountDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListServiceAccounts list service accounts API
*/
func (a *Client) ListServiceAccounts(params *ListServiceAccountsParams, opts ...ClientOption) (*ListServiceAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServiceAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListServiceAccounts",
		Method:             "GET",
		PathPattern:        "/api/v1/config/service-accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListServiceAccountsReader{formats: a.formats},
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
	success, ok := result.(*ListServiceAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListServiceAccountsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadServiceAccount read service account API
*/
func (a *Client) ReadServiceAccount(params *ReadServiceAccountParams, opts ...ClientOption) (*ReadServiceAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadServiceAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadServiceAccount",
		Method:             "GET",
		PathPattern:        "/api/v1/config/service-accounts/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadServiceAccountReader{formats: a.formats},
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
	success, ok := result.(*ReadServiceAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadServiceAccountDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
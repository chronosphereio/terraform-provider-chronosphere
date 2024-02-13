// Code generated by go-swagger; DO NOT EDIT.

package collection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new collection API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for collection API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCollection(params *CreateCollectionParams, opts ...ClientOption) (*CreateCollectionOK, error)

	DeleteCollection(params *DeleteCollectionParams, opts ...ClientOption) (*DeleteCollectionOK, error)

	ListCollections(params *ListCollectionsParams, opts ...ClientOption) (*ListCollectionsOK, error)

	ReadCollection(params *ReadCollectionParams, opts ...ClientOption) (*ReadCollectionOK, error)

	UpdateCollection(params *UpdateCollectionParams, opts ...ClientOption) (*UpdateCollectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateCollection create collection API
*/
func (a *Client) CreateCollection(params *CreateCollectionParams, opts ...ClientOption) (*CreateCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCollectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateCollection",
		Method:             "POST",
		PathPattern:        "/api/v1/config/collections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateCollectionReader{formats: a.formats},
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
	success, ok := result.(*CreateCollectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateCollectionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteCollection delete collection API
*/
func (a *Client) DeleteCollection(params *DeleteCollectionParams, opts ...ClientOption) (*DeleteCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCollectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCollection",
		Method:             "DELETE",
		PathPattern:        "/api/v1/config/collections/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCollectionReader{formats: a.formats},
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
	success, ok := result.(*DeleteCollectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteCollectionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListCollections list collections API
*/
func (a *Client) ListCollections(params *ListCollectionsParams, opts ...ClientOption) (*ListCollectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCollectionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListCollections",
		Method:             "GET",
		PathPattern:        "/api/v1/config/collections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListCollectionsReader{formats: a.formats},
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
	success, ok := result.(*ListCollectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListCollectionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadCollection read collection API
*/
func (a *Client) ReadCollection(params *ReadCollectionParams, opts ...ClientOption) (*ReadCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadCollectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadCollection",
		Method:             "GET",
		PathPattern:        "/api/v1/config/collections/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadCollectionReader{formats: a.formats},
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
	success, ok := result.(*ReadCollectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadCollectionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateCollection update collection API
*/
func (a *Client) UpdateCollection(params *UpdateCollectionParams, opts ...ClientOption) (*UpdateCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCollectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateCollection",
		Method:             "PUT",
		PathPattern:        "/api/v1/config/collections/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCollectionReader{formats: a.formats},
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
	success, ok := result.(*UpdateCollectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateCollectionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

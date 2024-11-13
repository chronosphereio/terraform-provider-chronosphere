// Code generated by go-swagger; DO NOT EDIT.

package s_l_o

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new s l o API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for s l o API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSLO(params *CreateSLOParams, opts ...ClientOption) (*CreateSLOOK, error)

	DeleteSLO(params *DeleteSLOParams, opts ...ClientOption) (*DeleteSLOOK, error)

	ListSLOs(params *ListSLOsParams, opts ...ClientOption) (*ListSLOsOK, error)

	ReadSLO(params *ReadSLOParams, opts ...ClientOption) (*ReadSLOOK, error)

	UpdateSLO(params *UpdateSLOParams, opts ...ClientOption) (*UpdateSLOOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateSLO create s l o API
*/
func (a *Client) CreateSLO(params *CreateSLOParams, opts ...ClientOption) (*CreateSLOOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSLOParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateSLO",
		Method:             "POST",
		PathPattern:        "/api/unstable/config/slos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSLOReader{formats: a.formats},
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
	success, ok := result.(*CreateSLOOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSLODefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteSLO delete s l o API
*/
func (a *Client) DeleteSLO(params *DeleteSLOParams, opts ...ClientOption) (*DeleteSLOOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSLOParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteSLO",
		Method:             "DELETE",
		PathPattern:        "/api/unstable/config/slos/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSLOReader{formats: a.formats},
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
	success, ok := result.(*DeleteSLOOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSLODefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListSLOs list s l os API
*/
func (a *Client) ListSLOs(params *ListSLOsParams, opts ...ClientOption) (*ListSLOsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSLOsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListSLOs",
		Method:             "GET",
		PathPattern:        "/api/unstable/config/slos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSLOsReader{formats: a.formats},
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
	success, ok := result.(*ListSLOsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListSLOsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadSLO read s l o API
*/
func (a *Client) ReadSLO(params *ReadSLOParams, opts ...ClientOption) (*ReadSLOOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadSLOParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadSLO",
		Method:             "GET",
		PathPattern:        "/api/unstable/config/slos/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadSLOReader{formats: a.formats},
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
	success, ok := result.(*ReadSLOOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadSLODefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateSLO update s l o API
*/
func (a *Client) UpdateSLO(params *UpdateSLOParams, opts ...ClientOption) (*UpdateSLOOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSLOParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateSLO",
		Method:             "PUT",
		PathPattern:        "/api/unstable/config/slos/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSLOReader{formats: a.formats},
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
	success, ok := result.(*UpdateSLOOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSLODefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

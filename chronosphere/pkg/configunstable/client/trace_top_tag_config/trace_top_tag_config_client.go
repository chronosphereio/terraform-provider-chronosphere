// Code generated by go-swagger; DO NOT EDIT.

package trace_top_tag_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new trace top tag config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for trace top tag config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTraceTopTagConfig(params *CreateTraceTopTagConfigParams, opts ...ClientOption) (*CreateTraceTopTagConfigOK, error)

	DeleteTraceTopTagConfig(params *DeleteTraceTopTagConfigParams, opts ...ClientOption) (*DeleteTraceTopTagConfigOK, error)

	ReadTraceTopTagConfig(params *ReadTraceTopTagConfigParams, opts ...ClientOption) (*ReadTraceTopTagConfigOK, error)

	UpdateTraceTopTagConfig(params *UpdateTraceTopTagConfigParams, opts ...ClientOption) (*UpdateTraceTopTagConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateTraceTopTagConfig create trace top tag config API
*/
func (a *Client) CreateTraceTopTagConfig(params *CreateTraceTopTagConfigParams, opts ...ClientOption) (*CreateTraceTopTagConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTraceTopTagConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateTraceTopTagConfig",
		Method:             "POST",
		PathPattern:        "/api/unstable/config/trace-top-tag-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateTraceTopTagConfigReader{formats: a.formats},
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
	success, ok := result.(*CreateTraceTopTagConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateTraceTopTagConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteTraceTopTagConfig delete trace top tag config API
*/
func (a *Client) DeleteTraceTopTagConfig(params *DeleteTraceTopTagConfigParams, opts ...ClientOption) (*DeleteTraceTopTagConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTraceTopTagConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteTraceTopTagConfig",
		Method:             "DELETE",
		PathPattern:        "/api/unstable/config/trace-top-tag-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTraceTopTagConfigReader{formats: a.formats},
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
	success, ok := result.(*DeleteTraceTopTagConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteTraceTopTagConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadTraceTopTagConfig read trace top tag config API
*/
func (a *Client) ReadTraceTopTagConfig(params *ReadTraceTopTagConfigParams, opts ...ClientOption) (*ReadTraceTopTagConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadTraceTopTagConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadTraceTopTagConfig",
		Method:             "GET",
		PathPattern:        "/api/unstable/config/trace-top-tag-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadTraceTopTagConfigReader{formats: a.formats},
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
	success, ok := result.(*ReadTraceTopTagConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadTraceTopTagConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateTraceTopTagConfig update trace top tag config API
*/
func (a *Client) UpdateTraceTopTagConfig(params *UpdateTraceTopTagConfigParams, opts ...ClientOption) (*UpdateTraceTopTagConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTraceTopTagConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateTraceTopTagConfig",
		Method:             "PUT",
		PathPattern:        "/api/unstable/config/trace-top-tag-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateTraceTopTagConfigReader{formats: a.formats},
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
	success, ok := result.(*UpdateTraceTopTagConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateTraceTopTagConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

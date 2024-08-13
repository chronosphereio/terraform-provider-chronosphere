// Code generated by go-swagger; DO NOT EDIT.

package log_scale_action

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new log scale action API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for log scale action API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateLogScaleAction(params *CreateLogScaleActionParams, opts ...ClientOption) (*CreateLogScaleActionOK, error)

	DeleteLogScaleAction(params *DeleteLogScaleActionParams, opts ...ClientOption) (*DeleteLogScaleActionOK, error)

	ListLogScaleActions(params *ListLogScaleActionsParams, opts ...ClientOption) (*ListLogScaleActionsOK, error)

	ReadLogScaleAction(params *ReadLogScaleActionParams, opts ...ClientOption) (*ReadLogScaleActionOK, error)

	UpdateLogScaleAction(params *UpdateLogScaleActionParams, opts ...ClientOption) (*UpdateLogScaleActionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateLogScaleAction create log scale action API
*/
func (a *Client) CreateLogScaleAction(params *CreateLogScaleActionParams, opts ...ClientOption) (*CreateLogScaleActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLogScaleActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateLogScaleAction",
		Method:             "POST",
		PathPattern:        "/api/v1/config/log-scale-actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateLogScaleActionReader{formats: a.formats},
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
	success, ok := result.(*CreateLogScaleActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateLogScaleActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteLogScaleAction delete log scale action API
*/
func (a *Client) DeleteLogScaleAction(params *DeleteLogScaleActionParams, opts ...ClientOption) (*DeleteLogScaleActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLogScaleActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteLogScaleAction",
		Method:             "DELETE",
		PathPattern:        "/api/v1/config/log-scale-actions/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteLogScaleActionReader{formats: a.formats},
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
	success, ok := result.(*DeleteLogScaleActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteLogScaleActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListLogScaleActions list log scale actions API
*/
func (a *Client) ListLogScaleActions(params *ListLogScaleActionsParams, opts ...ClientOption) (*ListLogScaleActionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListLogScaleActionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListLogScaleActions",
		Method:             "GET",
		PathPattern:        "/api/v1/config/log-scale-actions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListLogScaleActionsReader{formats: a.formats},
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
	success, ok := result.(*ListLogScaleActionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListLogScaleActionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadLogScaleAction read log scale action API
*/
func (a *Client) ReadLogScaleAction(params *ReadLogScaleActionParams, opts ...ClientOption) (*ReadLogScaleActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadLogScaleActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadLogScaleAction",
		Method:             "GET",
		PathPattern:        "/api/v1/config/log-scale-actions/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadLogScaleActionReader{formats: a.formats},
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
	success, ok := result.(*ReadLogScaleActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadLogScaleActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateLogScaleAction update log scale action API
*/
func (a *Client) UpdateLogScaleAction(params *UpdateLogScaleActionParams, opts ...ClientOption) (*UpdateLogScaleActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLogScaleActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateLogScaleAction",
		Method:             "PUT",
		PathPattern:        "/api/v1/config/log-scale-actions/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateLogScaleActionReader{formats: a.formats},
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
	success, ok := result.(*UpdateLogScaleActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateLogScaleActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

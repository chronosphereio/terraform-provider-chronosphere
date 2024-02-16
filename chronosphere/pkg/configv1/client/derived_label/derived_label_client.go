// Code generated by go-swagger; DO NOT EDIT.

package derived_label

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new derived label API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for derived label API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDerivedLabel(params *CreateDerivedLabelParams, opts ...ClientOption) (*CreateDerivedLabelOK, error)

	DeleteDerivedLabel(params *DeleteDerivedLabelParams, opts ...ClientOption) (*DeleteDerivedLabelOK, error)

	ListDerivedLabels(params *ListDerivedLabelsParams, opts ...ClientOption) (*ListDerivedLabelsOK, error)

	ReadDerivedLabel(params *ReadDerivedLabelParams, opts ...ClientOption) (*ReadDerivedLabelOK, error)

	UpdateDerivedLabel(params *UpdateDerivedLabelParams, opts ...ClientOption) (*UpdateDerivedLabelOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateDerivedLabel create derived label API
*/
func (a *Client) CreateDerivedLabel(params *CreateDerivedLabelParams, opts ...ClientOption) (*CreateDerivedLabelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDerivedLabelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateDerivedLabel",
		Method:             "POST",
		PathPattern:        "/api/v1/config/derived-labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDerivedLabelReader{formats: a.formats},
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
	success, ok := result.(*CreateDerivedLabelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateDerivedLabelDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteDerivedLabel delete derived label API
*/
func (a *Client) DeleteDerivedLabel(params *DeleteDerivedLabelParams, opts ...ClientOption) (*DeleteDerivedLabelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDerivedLabelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteDerivedLabel",
		Method:             "DELETE",
		PathPattern:        "/api/v1/config/derived-labels/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDerivedLabelReader{formats: a.formats},
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
	success, ok := result.(*DeleteDerivedLabelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteDerivedLabelDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListDerivedLabels list derived labels API
*/
func (a *Client) ListDerivedLabels(params *ListDerivedLabelsParams, opts ...ClientOption) (*ListDerivedLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDerivedLabelsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListDerivedLabels",
		Method:             "GET",
		PathPattern:        "/api/v1/config/derived-labels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListDerivedLabelsReader{formats: a.formats},
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
	success, ok := result.(*ListDerivedLabelsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDerivedLabelsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadDerivedLabel read derived label API
*/
func (a *Client) ReadDerivedLabel(params *ReadDerivedLabelParams, opts ...ClientOption) (*ReadDerivedLabelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadDerivedLabelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadDerivedLabel",
		Method:             "GET",
		PathPattern:        "/api/v1/config/derived-labels/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadDerivedLabelReader{formats: a.formats},
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
	success, ok := result.(*ReadDerivedLabelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadDerivedLabelDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateDerivedLabel update derived label API
*/
func (a *Client) UpdateDerivedLabel(params *UpdateDerivedLabelParams, opts ...ClientOption) (*UpdateDerivedLabelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDerivedLabelParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateDerivedLabel",
		Method:             "PUT",
		PathPattern:        "/api/v1/config/derived-labels/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDerivedLabelReader{formats: a.formats},
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
	success, ok := result.(*UpdateDerivedLabelOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateDerivedLabelDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
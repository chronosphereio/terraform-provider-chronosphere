// Code generated by go-swagger; DO NOT EDIT.

package rollup_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new rollup rule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rollup rule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRollupRule(params *CreateRollupRuleParams, opts ...ClientOption) (*CreateRollupRuleOK, error)

	DeleteRollupRule(params *DeleteRollupRuleParams, opts ...ClientOption) (*DeleteRollupRuleOK, error)

	ListRollupRules(params *ListRollupRulesParams, opts ...ClientOption) (*ListRollupRulesOK, error)

	ReadRollupRule(params *ReadRollupRuleParams, opts ...ClientOption) (*ReadRollupRuleOK, error)

	UpdateRollupRule(params *UpdateRollupRuleParams, opts ...ClientOption) (*UpdateRollupRuleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateRollupRule create rollup rule API
*/
func (a *Client) CreateRollupRule(params *CreateRollupRuleParams, opts ...ClientOption) (*CreateRollupRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRollupRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateRollupRule",
		Method:             "POST",
		PathPattern:        "/api/v1/config/rollup-rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRollupRuleReader{formats: a.formats},
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
	success, ok := result.(*CreateRollupRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateRollupRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteRollupRule delete rollup rule API
*/
func (a *Client) DeleteRollupRule(params *DeleteRollupRuleParams, opts ...ClientOption) (*DeleteRollupRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRollupRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteRollupRule",
		Method:             "DELETE",
		PathPattern:        "/api/v1/config/rollup-rules/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRollupRuleReader{formats: a.formats},
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
	success, ok := result.(*DeleteRollupRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteRollupRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListRollupRules list rollup rules API
*/
func (a *Client) ListRollupRules(params *ListRollupRulesParams, opts ...ClientOption) (*ListRollupRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRollupRulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListRollupRules",
		Method:             "GET",
		PathPattern:        "/api/v1/config/rollup-rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListRollupRulesReader{formats: a.formats},
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
	success, ok := result.(*ListRollupRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListRollupRulesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadRollupRule read rollup rule API
*/
func (a *Client) ReadRollupRule(params *ReadRollupRuleParams, opts ...ClientOption) (*ReadRollupRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadRollupRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadRollupRule",
		Method:             "GET",
		PathPattern:        "/api/v1/config/rollup-rules/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadRollupRuleReader{formats: a.formats},
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
	success, ok := result.(*ReadRollupRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadRollupRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateRollupRule update rollup rule API
*/
func (a *Client) UpdateRollupRule(params *UpdateRollupRuleParams, opts ...ClientOption) (*UpdateRollupRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRollupRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateRollupRule",
		Method:             "PUT",
		PathPattern:        "/api/v1/config/rollup-rules/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateRollupRuleReader{formats: a.formats},
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
	success, ok := result.(*UpdateRollupRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateRollupRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
// Code generated by go-swagger; DO NOT EDIT.

package trace_metrics_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new trace metrics rule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for trace metrics rule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTraceMetricsRule(params *CreateTraceMetricsRuleParams, opts ...ClientOption) (*CreateTraceMetricsRuleOK, error)

	DeleteTraceMetricsRule(params *DeleteTraceMetricsRuleParams, opts ...ClientOption) (*DeleteTraceMetricsRuleOK, error)

	ListTraceMetricsRules(params *ListTraceMetricsRulesParams, opts ...ClientOption) (*ListTraceMetricsRulesOK, error)

	ReadTraceMetricsRule(params *ReadTraceMetricsRuleParams, opts ...ClientOption) (*ReadTraceMetricsRuleOK, error)

	UpdateTraceMetricsRule(params *UpdateTraceMetricsRuleParams, opts ...ClientOption) (*UpdateTraceMetricsRuleOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	CreateTraceMetricsRule ***

Trace Metrics Rules
***
*/
func (a *Client) CreateTraceMetricsRule(params *CreateTraceMetricsRuleParams, opts ...ClientOption) (*CreateTraceMetricsRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTraceMetricsRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateTraceMetricsRule",
		Method:             "POST",
		PathPattern:        "/api/v1/config/trace-metrics-rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateTraceMetricsRuleReader{formats: a.formats},
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
	success, ok := result.(*CreateTraceMetricsRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateTraceMetricsRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteTraceMetricsRule delete trace metrics rule API
*/
func (a *Client) DeleteTraceMetricsRule(params *DeleteTraceMetricsRuleParams, opts ...ClientOption) (*DeleteTraceMetricsRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTraceMetricsRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteTraceMetricsRule",
		Method:             "DELETE",
		PathPattern:        "/api/v1/config/trace-metrics-rules/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTraceMetricsRuleReader{formats: a.formats},
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
	success, ok := result.(*DeleteTraceMetricsRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteTraceMetricsRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListTraceMetricsRules list trace metrics rules API
*/
func (a *Client) ListTraceMetricsRules(params *ListTraceMetricsRulesParams, opts ...ClientOption) (*ListTraceMetricsRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTraceMetricsRulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListTraceMetricsRules",
		Method:             "GET",
		PathPattern:        "/api/v1/config/trace-metrics-rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListTraceMetricsRulesReader{formats: a.formats},
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
	success, ok := result.(*ListTraceMetricsRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListTraceMetricsRulesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadTraceMetricsRule read trace metrics rule API
*/
func (a *Client) ReadTraceMetricsRule(params *ReadTraceMetricsRuleParams, opts ...ClientOption) (*ReadTraceMetricsRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadTraceMetricsRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadTraceMetricsRule",
		Method:             "GET",
		PathPattern:        "/api/v1/config/trace-metrics-rules/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadTraceMetricsRuleReader{formats: a.formats},
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
	success, ok := result.(*ReadTraceMetricsRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadTraceMetricsRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateTraceMetricsRule update trace metrics rule API
*/
func (a *Client) UpdateTraceMetricsRule(params *UpdateTraceMetricsRuleParams, opts ...ClientOption) (*UpdateTraceMetricsRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTraceMetricsRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateTraceMetricsRule",
		Method:             "PUT",
		PathPattern:        "/api/v1/config/trace-metrics-rules/{slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateTraceMetricsRuleReader{formats: a.formats},
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
	success, ok := result.(*UpdateTraceMetricsRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateTraceMetricsRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
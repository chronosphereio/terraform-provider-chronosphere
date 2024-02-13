// Code generated by go-swagger; DO NOT EDIT.

package trace_jaeger_remote_sampling_strategy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

// ListTraceJaegerRemoteSamplingStrategiesReader is a Reader for the ListTraceJaegerRemoteSamplingStrategies structure.
type ListTraceJaegerRemoteSamplingStrategiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTraceJaegerRemoteSamplingStrategiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTraceJaegerRemoteSamplingStrategiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListTraceJaegerRemoteSamplingStrategiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListTraceJaegerRemoteSamplingStrategiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListTraceJaegerRemoteSamplingStrategiesOK creates a ListTraceJaegerRemoteSamplingStrategiesOK with default headers values
func NewListTraceJaegerRemoteSamplingStrategiesOK() *ListTraceJaegerRemoteSamplingStrategiesOK {
	return &ListTraceJaegerRemoteSamplingStrategiesOK{}
}

/*
ListTraceJaegerRemoteSamplingStrategiesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListTraceJaegerRemoteSamplingStrategiesOK struct {
	Payload *models.ConfigunstableListTraceJaegerRemoteSamplingStrategiesResponse
}

// IsSuccess returns true when this list trace jaeger remote sampling strategies o k response has a 2xx status code
func (o *ListTraceJaegerRemoteSamplingStrategiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list trace jaeger remote sampling strategies o k response has a 3xx status code
func (o *ListTraceJaegerRemoteSamplingStrategiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list trace jaeger remote sampling strategies o k response has a 4xx status code
func (o *ListTraceJaegerRemoteSamplingStrategiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list trace jaeger remote sampling strategies o k response has a 5xx status code
func (o *ListTraceJaegerRemoteSamplingStrategiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list trace jaeger remote sampling strategies o k response a status code equal to that given
func (o *ListTraceJaegerRemoteSamplingStrategiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list trace jaeger remote sampling strategies o k response
func (o *ListTraceJaegerRemoteSamplingStrategiesOK) Code() int {
	return 200
}

func (o *ListTraceJaegerRemoteSamplingStrategiesOK) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies][%d] listTraceJaegerRemoteSamplingStrategiesOK  %+v", 200, o.Payload)
}

func (o *ListTraceJaegerRemoteSamplingStrategiesOK) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies][%d] listTraceJaegerRemoteSamplingStrategiesOK  %+v", 200, o.Payload)
}

func (o *ListTraceJaegerRemoteSamplingStrategiesOK) GetPayload() *models.ConfigunstableListTraceJaegerRemoteSamplingStrategiesResponse {
	return o.Payload
}

func (o *ListTraceJaegerRemoteSamplingStrategiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableListTraceJaegerRemoteSamplingStrategiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTraceJaegerRemoteSamplingStrategiesInternalServerError creates a ListTraceJaegerRemoteSamplingStrategiesInternalServerError with default headers values
func NewListTraceJaegerRemoteSamplingStrategiesInternalServerError() *ListTraceJaegerRemoteSamplingStrategiesInternalServerError {
	return &ListTraceJaegerRemoteSamplingStrategiesInternalServerError{}
}

/*
ListTraceJaegerRemoteSamplingStrategiesInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ListTraceJaegerRemoteSamplingStrategiesInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this list trace jaeger remote sampling strategies internal server error response has a 2xx status code
func (o *ListTraceJaegerRemoteSamplingStrategiesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list trace jaeger remote sampling strategies internal server error response has a 3xx status code
func (o *ListTraceJaegerRemoteSamplingStrategiesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list trace jaeger remote sampling strategies internal server error response has a 4xx status code
func (o *ListTraceJaegerRemoteSamplingStrategiesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list trace jaeger remote sampling strategies internal server error response has a 5xx status code
func (o *ListTraceJaegerRemoteSamplingStrategiesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list trace jaeger remote sampling strategies internal server error response a status code equal to that given
func (o *ListTraceJaegerRemoteSamplingStrategiesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list trace jaeger remote sampling strategies internal server error response
func (o *ListTraceJaegerRemoteSamplingStrategiesInternalServerError) Code() int {
	return 500
}

func (o *ListTraceJaegerRemoteSamplingStrategiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies][%d] listTraceJaegerRemoteSamplingStrategiesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListTraceJaegerRemoteSamplingStrategiesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies][%d] listTraceJaegerRemoteSamplingStrategiesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListTraceJaegerRemoteSamplingStrategiesInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListTraceJaegerRemoteSamplingStrategiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTraceJaegerRemoteSamplingStrategiesDefault creates a ListTraceJaegerRemoteSamplingStrategiesDefault with default headers values
func NewListTraceJaegerRemoteSamplingStrategiesDefault(code int) *ListTraceJaegerRemoteSamplingStrategiesDefault {
	return &ListTraceJaegerRemoteSamplingStrategiesDefault{
		_statusCode: code,
	}
}

/*
ListTraceJaegerRemoteSamplingStrategiesDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ListTraceJaegerRemoteSamplingStrategiesDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this list trace jaeger remote sampling strategies default response has a 2xx status code
func (o *ListTraceJaegerRemoteSamplingStrategiesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list trace jaeger remote sampling strategies default response has a 3xx status code
func (o *ListTraceJaegerRemoteSamplingStrategiesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list trace jaeger remote sampling strategies default response has a 4xx status code
func (o *ListTraceJaegerRemoteSamplingStrategiesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list trace jaeger remote sampling strategies default response has a 5xx status code
func (o *ListTraceJaegerRemoteSamplingStrategiesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list trace jaeger remote sampling strategies default response a status code equal to that given
func (o *ListTraceJaegerRemoteSamplingStrategiesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list trace jaeger remote sampling strategies default response
func (o *ListTraceJaegerRemoteSamplingStrategiesDefault) Code() int {
	return o._statusCode
}

func (o *ListTraceJaegerRemoteSamplingStrategiesDefault) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies][%d] ListTraceJaegerRemoteSamplingStrategies default  %+v", o._statusCode, o.Payload)
}

func (o *ListTraceJaegerRemoteSamplingStrategiesDefault) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies][%d] ListTraceJaegerRemoteSamplingStrategies default  %+v", o._statusCode, o.Payload)
}

func (o *ListTraceJaegerRemoteSamplingStrategiesDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ListTraceJaegerRemoteSamplingStrategiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

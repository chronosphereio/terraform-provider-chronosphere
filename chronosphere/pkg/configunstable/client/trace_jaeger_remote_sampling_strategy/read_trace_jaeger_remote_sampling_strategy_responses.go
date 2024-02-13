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

// ReadTraceJaegerRemoteSamplingStrategyReader is a Reader for the ReadTraceJaegerRemoteSamplingStrategy structure.
type ReadTraceJaegerRemoteSamplingStrategyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadTraceJaegerRemoteSamplingStrategyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadTraceJaegerRemoteSamplingStrategyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadTraceJaegerRemoteSamplingStrategyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadTraceJaegerRemoteSamplingStrategyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadTraceJaegerRemoteSamplingStrategyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadTraceJaegerRemoteSamplingStrategyOK creates a ReadTraceJaegerRemoteSamplingStrategyOK with default headers values
func NewReadTraceJaegerRemoteSamplingStrategyOK() *ReadTraceJaegerRemoteSamplingStrategyOK {
	return &ReadTraceJaegerRemoteSamplingStrategyOK{}
}

/*
ReadTraceJaegerRemoteSamplingStrategyOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadTraceJaegerRemoteSamplingStrategyOK struct {
	Payload *models.ConfigunstableReadTraceJaegerRemoteSamplingStrategyResponse
}

// IsSuccess returns true when this read trace jaeger remote sampling strategy o k response has a 2xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read trace jaeger remote sampling strategy o k response has a 3xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read trace jaeger remote sampling strategy o k response has a 4xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read trace jaeger remote sampling strategy o k response has a 5xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read trace jaeger remote sampling strategy o k response a status code equal to that given
func (o *ReadTraceJaegerRemoteSamplingStrategyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read trace jaeger remote sampling strategy o k response
func (o *ReadTraceJaegerRemoteSamplingStrategyOK) Code() int {
	return 200
}

func (o *ReadTraceJaegerRemoteSamplingStrategyOK) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] readTraceJaegerRemoteSamplingStrategyOK  %+v", 200, o.Payload)
}

func (o *ReadTraceJaegerRemoteSamplingStrategyOK) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] readTraceJaegerRemoteSamplingStrategyOK  %+v", 200, o.Payload)
}

func (o *ReadTraceJaegerRemoteSamplingStrategyOK) GetPayload() *models.ConfigunstableReadTraceJaegerRemoteSamplingStrategyResponse {
	return o.Payload
}

func (o *ReadTraceJaegerRemoteSamplingStrategyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableReadTraceJaegerRemoteSamplingStrategyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTraceJaegerRemoteSamplingStrategyNotFound creates a ReadTraceJaegerRemoteSamplingStrategyNotFound with default headers values
func NewReadTraceJaegerRemoteSamplingStrategyNotFound() *ReadTraceJaegerRemoteSamplingStrategyNotFound {
	return &ReadTraceJaegerRemoteSamplingStrategyNotFound{}
}

/*
ReadTraceJaegerRemoteSamplingStrategyNotFound describes a response with status code 404, with default header values.

Cannot read the TraceJaegerRemoteSamplingStrategy because the slug does not exist.
*/
type ReadTraceJaegerRemoteSamplingStrategyNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read trace jaeger remote sampling strategy not found response has a 2xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read trace jaeger remote sampling strategy not found response has a 3xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read trace jaeger remote sampling strategy not found response has a 4xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read trace jaeger remote sampling strategy not found response has a 5xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read trace jaeger remote sampling strategy not found response a status code equal to that given
func (o *ReadTraceJaegerRemoteSamplingStrategyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read trace jaeger remote sampling strategy not found response
func (o *ReadTraceJaegerRemoteSamplingStrategyNotFound) Code() int {
	return 404
}

func (o *ReadTraceJaegerRemoteSamplingStrategyNotFound) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] readTraceJaegerRemoteSamplingStrategyNotFound  %+v", 404, o.Payload)
}

func (o *ReadTraceJaegerRemoteSamplingStrategyNotFound) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] readTraceJaegerRemoteSamplingStrategyNotFound  %+v", 404, o.Payload)
}

func (o *ReadTraceJaegerRemoteSamplingStrategyNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadTraceJaegerRemoteSamplingStrategyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTraceJaegerRemoteSamplingStrategyInternalServerError creates a ReadTraceJaegerRemoteSamplingStrategyInternalServerError with default headers values
func NewReadTraceJaegerRemoteSamplingStrategyInternalServerError() *ReadTraceJaegerRemoteSamplingStrategyInternalServerError {
	return &ReadTraceJaegerRemoteSamplingStrategyInternalServerError{}
}

/*
ReadTraceJaegerRemoteSamplingStrategyInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadTraceJaegerRemoteSamplingStrategyInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read trace jaeger remote sampling strategy internal server error response has a 2xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read trace jaeger remote sampling strategy internal server error response has a 3xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read trace jaeger remote sampling strategy internal server error response has a 4xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read trace jaeger remote sampling strategy internal server error response has a 5xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read trace jaeger remote sampling strategy internal server error response a status code equal to that given
func (o *ReadTraceJaegerRemoteSamplingStrategyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read trace jaeger remote sampling strategy internal server error response
func (o *ReadTraceJaegerRemoteSamplingStrategyInternalServerError) Code() int {
	return 500
}

func (o *ReadTraceJaegerRemoteSamplingStrategyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] readTraceJaegerRemoteSamplingStrategyInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadTraceJaegerRemoteSamplingStrategyInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] readTraceJaegerRemoteSamplingStrategyInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadTraceJaegerRemoteSamplingStrategyInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadTraceJaegerRemoteSamplingStrategyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTraceJaegerRemoteSamplingStrategyDefault creates a ReadTraceJaegerRemoteSamplingStrategyDefault with default headers values
func NewReadTraceJaegerRemoteSamplingStrategyDefault(code int) *ReadTraceJaegerRemoteSamplingStrategyDefault {
	return &ReadTraceJaegerRemoteSamplingStrategyDefault{
		_statusCode: code,
	}
}

/*
ReadTraceJaegerRemoteSamplingStrategyDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadTraceJaegerRemoteSamplingStrategyDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read trace jaeger remote sampling strategy default response has a 2xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read trace jaeger remote sampling strategy default response has a 3xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read trace jaeger remote sampling strategy default response has a 4xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read trace jaeger remote sampling strategy default response has a 5xx status code
func (o *ReadTraceJaegerRemoteSamplingStrategyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read trace jaeger remote sampling strategy default response a status code equal to that given
func (o *ReadTraceJaegerRemoteSamplingStrategyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read trace jaeger remote sampling strategy default response
func (o *ReadTraceJaegerRemoteSamplingStrategyDefault) Code() int {
	return o._statusCode
}

func (o *ReadTraceJaegerRemoteSamplingStrategyDefault) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] ReadTraceJaegerRemoteSamplingStrategy default  %+v", o._statusCode, o.Payload)
}

func (o *ReadTraceJaegerRemoteSamplingStrategyDefault) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] ReadTraceJaegerRemoteSamplingStrategy default  %+v", o._statusCode, o.Payload)
}

func (o *ReadTraceJaegerRemoteSamplingStrategyDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadTraceJaegerRemoteSamplingStrategyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

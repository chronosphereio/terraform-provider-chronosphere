// Code generated by go-swagger; DO NOT EDIT.

package otel_metrics_ingestion

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

// ReadOtelMetricsIngestionReader is a Reader for the ReadOtelMetricsIngestion structure.
type ReadOtelMetricsIngestionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadOtelMetricsIngestionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadOtelMetricsIngestionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadOtelMetricsIngestionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadOtelMetricsIngestionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadOtelMetricsIngestionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadOtelMetricsIngestionOK creates a ReadOtelMetricsIngestionOK with default headers values
func NewReadOtelMetricsIngestionOK() *ReadOtelMetricsIngestionOK {
	return &ReadOtelMetricsIngestionOK{}
}

/*
ReadOtelMetricsIngestionOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadOtelMetricsIngestionOK struct {
	Payload *models.ConfigunstableReadOtelMetricsIngestionResponse
}

// IsSuccess returns true when this read otel metrics ingestion o k response has a 2xx status code
func (o *ReadOtelMetricsIngestionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read otel metrics ingestion o k response has a 3xx status code
func (o *ReadOtelMetricsIngestionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read otel metrics ingestion o k response has a 4xx status code
func (o *ReadOtelMetricsIngestionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read otel metrics ingestion o k response has a 5xx status code
func (o *ReadOtelMetricsIngestionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read otel metrics ingestion o k response a status code equal to that given
func (o *ReadOtelMetricsIngestionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read otel metrics ingestion o k response
func (o *ReadOtelMetricsIngestionOK) Code() int {
	return 200
}

func (o *ReadOtelMetricsIngestionOK) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/otel-metrics-ingestion][%d] readOtelMetricsIngestionOK  %+v", 200, o.Payload)
}

func (o *ReadOtelMetricsIngestionOK) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/otel-metrics-ingestion][%d] readOtelMetricsIngestionOK  %+v", 200, o.Payload)
}

func (o *ReadOtelMetricsIngestionOK) GetPayload() *models.ConfigunstableReadOtelMetricsIngestionResponse {
	return o.Payload
}

func (o *ReadOtelMetricsIngestionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableReadOtelMetricsIngestionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadOtelMetricsIngestionNotFound creates a ReadOtelMetricsIngestionNotFound with default headers values
func NewReadOtelMetricsIngestionNotFound() *ReadOtelMetricsIngestionNotFound {
	return &ReadOtelMetricsIngestionNotFound{}
}

/*
ReadOtelMetricsIngestionNotFound describes a response with status code 404, with default header values.

Cannot read the OtelMetricsIngestion because OtelMetricsIngestion has not been created.
*/
type ReadOtelMetricsIngestionNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read otel metrics ingestion not found response has a 2xx status code
func (o *ReadOtelMetricsIngestionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read otel metrics ingestion not found response has a 3xx status code
func (o *ReadOtelMetricsIngestionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read otel metrics ingestion not found response has a 4xx status code
func (o *ReadOtelMetricsIngestionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read otel metrics ingestion not found response has a 5xx status code
func (o *ReadOtelMetricsIngestionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read otel metrics ingestion not found response a status code equal to that given
func (o *ReadOtelMetricsIngestionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read otel metrics ingestion not found response
func (o *ReadOtelMetricsIngestionNotFound) Code() int {
	return 404
}

func (o *ReadOtelMetricsIngestionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/otel-metrics-ingestion][%d] readOtelMetricsIngestionNotFound  %+v", 404, o.Payload)
}

func (o *ReadOtelMetricsIngestionNotFound) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/otel-metrics-ingestion][%d] readOtelMetricsIngestionNotFound  %+v", 404, o.Payload)
}

func (o *ReadOtelMetricsIngestionNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadOtelMetricsIngestionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadOtelMetricsIngestionInternalServerError creates a ReadOtelMetricsIngestionInternalServerError with default headers values
func NewReadOtelMetricsIngestionInternalServerError() *ReadOtelMetricsIngestionInternalServerError {
	return &ReadOtelMetricsIngestionInternalServerError{}
}

/*
ReadOtelMetricsIngestionInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadOtelMetricsIngestionInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read otel metrics ingestion internal server error response has a 2xx status code
func (o *ReadOtelMetricsIngestionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read otel metrics ingestion internal server error response has a 3xx status code
func (o *ReadOtelMetricsIngestionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read otel metrics ingestion internal server error response has a 4xx status code
func (o *ReadOtelMetricsIngestionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read otel metrics ingestion internal server error response has a 5xx status code
func (o *ReadOtelMetricsIngestionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read otel metrics ingestion internal server error response a status code equal to that given
func (o *ReadOtelMetricsIngestionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read otel metrics ingestion internal server error response
func (o *ReadOtelMetricsIngestionInternalServerError) Code() int {
	return 500
}

func (o *ReadOtelMetricsIngestionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/otel-metrics-ingestion][%d] readOtelMetricsIngestionInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadOtelMetricsIngestionInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/otel-metrics-ingestion][%d] readOtelMetricsIngestionInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadOtelMetricsIngestionInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadOtelMetricsIngestionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadOtelMetricsIngestionDefault creates a ReadOtelMetricsIngestionDefault with default headers values
func NewReadOtelMetricsIngestionDefault(code int) *ReadOtelMetricsIngestionDefault {
	return &ReadOtelMetricsIngestionDefault{
		_statusCode: code,
	}
}

/*
ReadOtelMetricsIngestionDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadOtelMetricsIngestionDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read otel metrics ingestion default response has a 2xx status code
func (o *ReadOtelMetricsIngestionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read otel metrics ingestion default response has a 3xx status code
func (o *ReadOtelMetricsIngestionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read otel metrics ingestion default response has a 4xx status code
func (o *ReadOtelMetricsIngestionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read otel metrics ingestion default response has a 5xx status code
func (o *ReadOtelMetricsIngestionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read otel metrics ingestion default response a status code equal to that given
func (o *ReadOtelMetricsIngestionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read otel metrics ingestion default response
func (o *ReadOtelMetricsIngestionDefault) Code() int {
	return o._statusCode
}

func (o *ReadOtelMetricsIngestionDefault) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/otel-metrics-ingestion][%d] ReadOtelMetricsIngestion default  %+v", o._statusCode, o.Payload)
}

func (o *ReadOtelMetricsIngestionDefault) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/otel-metrics-ingestion][%d] ReadOtelMetricsIngestion default  %+v", o._statusCode, o.Payload)
}

func (o *ReadOtelMetricsIngestionDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadOtelMetricsIngestionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

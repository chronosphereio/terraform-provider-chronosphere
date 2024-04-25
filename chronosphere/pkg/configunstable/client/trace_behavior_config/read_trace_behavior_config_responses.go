// Code generated by go-swagger; DO NOT EDIT.

package trace_behavior_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

// ReadTraceBehaviorConfigReader is a Reader for the ReadTraceBehaviorConfig structure.
type ReadTraceBehaviorConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadTraceBehaviorConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadTraceBehaviorConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadTraceBehaviorConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadTraceBehaviorConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadTraceBehaviorConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadTraceBehaviorConfigOK creates a ReadTraceBehaviorConfigOK with default headers values
func NewReadTraceBehaviorConfigOK() *ReadTraceBehaviorConfigOK {
	return &ReadTraceBehaviorConfigOK{}
}

/*
ReadTraceBehaviorConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadTraceBehaviorConfigOK struct {
	Payload *models.ConfigunstableReadTraceBehaviorConfigResponse
}

// IsSuccess returns true when this read trace behavior config o k response has a 2xx status code
func (o *ReadTraceBehaviorConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read trace behavior config o k response has a 3xx status code
func (o *ReadTraceBehaviorConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read trace behavior config o k response has a 4xx status code
func (o *ReadTraceBehaviorConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read trace behavior config o k response has a 5xx status code
func (o *ReadTraceBehaviorConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read trace behavior config o k response a status code equal to that given
func (o *ReadTraceBehaviorConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read trace behavior config o k response
func (o *ReadTraceBehaviorConfigOK) Code() int {
	return 200
}

func (o *ReadTraceBehaviorConfigOK) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-behavior-config][%d] readTraceBehaviorConfigOK  %+v", 200, o.Payload)
}

func (o *ReadTraceBehaviorConfigOK) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-behavior-config][%d] readTraceBehaviorConfigOK  %+v", 200, o.Payload)
}

func (o *ReadTraceBehaviorConfigOK) GetPayload() *models.ConfigunstableReadTraceBehaviorConfigResponse {
	return o.Payload
}

func (o *ReadTraceBehaviorConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableReadTraceBehaviorConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTraceBehaviorConfigNotFound creates a ReadTraceBehaviorConfigNotFound with default headers values
func NewReadTraceBehaviorConfigNotFound() *ReadTraceBehaviorConfigNotFound {
	return &ReadTraceBehaviorConfigNotFound{}
}

/*
ReadTraceBehaviorConfigNotFound describes a response with status code 404, with default header values.

Cannot read the TraceBehaviorConfig because TraceBehaviorConfig has not been created.
*/
type ReadTraceBehaviorConfigNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read trace behavior config not found response has a 2xx status code
func (o *ReadTraceBehaviorConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read trace behavior config not found response has a 3xx status code
func (o *ReadTraceBehaviorConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read trace behavior config not found response has a 4xx status code
func (o *ReadTraceBehaviorConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read trace behavior config not found response has a 5xx status code
func (o *ReadTraceBehaviorConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read trace behavior config not found response a status code equal to that given
func (o *ReadTraceBehaviorConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read trace behavior config not found response
func (o *ReadTraceBehaviorConfigNotFound) Code() int {
	return 404
}

func (o *ReadTraceBehaviorConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-behavior-config][%d] readTraceBehaviorConfigNotFound  %+v", 404, o.Payload)
}

func (o *ReadTraceBehaviorConfigNotFound) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-behavior-config][%d] readTraceBehaviorConfigNotFound  %+v", 404, o.Payload)
}

func (o *ReadTraceBehaviorConfigNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadTraceBehaviorConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTraceBehaviorConfigInternalServerError creates a ReadTraceBehaviorConfigInternalServerError with default headers values
func NewReadTraceBehaviorConfigInternalServerError() *ReadTraceBehaviorConfigInternalServerError {
	return &ReadTraceBehaviorConfigInternalServerError{}
}

/*
ReadTraceBehaviorConfigInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadTraceBehaviorConfigInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read trace behavior config internal server error response has a 2xx status code
func (o *ReadTraceBehaviorConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read trace behavior config internal server error response has a 3xx status code
func (o *ReadTraceBehaviorConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read trace behavior config internal server error response has a 4xx status code
func (o *ReadTraceBehaviorConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read trace behavior config internal server error response has a 5xx status code
func (o *ReadTraceBehaviorConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read trace behavior config internal server error response a status code equal to that given
func (o *ReadTraceBehaviorConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read trace behavior config internal server error response
func (o *ReadTraceBehaviorConfigInternalServerError) Code() int {
	return 500
}

func (o *ReadTraceBehaviorConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-behavior-config][%d] readTraceBehaviorConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadTraceBehaviorConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-behavior-config][%d] readTraceBehaviorConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadTraceBehaviorConfigInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadTraceBehaviorConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadTraceBehaviorConfigDefault creates a ReadTraceBehaviorConfigDefault with default headers values
func NewReadTraceBehaviorConfigDefault(code int) *ReadTraceBehaviorConfigDefault {
	return &ReadTraceBehaviorConfigDefault{
		_statusCode: code,
	}
}

/*
ReadTraceBehaviorConfigDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadTraceBehaviorConfigDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read trace behavior config default response has a 2xx status code
func (o *ReadTraceBehaviorConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read trace behavior config default response has a 3xx status code
func (o *ReadTraceBehaviorConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read trace behavior config default response has a 4xx status code
func (o *ReadTraceBehaviorConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read trace behavior config default response has a 5xx status code
func (o *ReadTraceBehaviorConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read trace behavior config default response a status code equal to that given
func (o *ReadTraceBehaviorConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read trace behavior config default response
func (o *ReadTraceBehaviorConfigDefault) Code() int {
	return o._statusCode
}

func (o *ReadTraceBehaviorConfigDefault) Error() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-behavior-config][%d] ReadTraceBehaviorConfig default  %+v", o._statusCode, o.Payload)
}

func (o *ReadTraceBehaviorConfigDefault) String() string {
	return fmt.Sprintf("[GET /api/unstable/config/trace-behavior-config][%d] ReadTraceBehaviorConfig default  %+v", o._statusCode, o.Payload)
}

func (o *ReadTraceBehaviorConfigDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadTraceBehaviorConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

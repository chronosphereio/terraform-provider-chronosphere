// Code generated by go-swagger; DO NOT EDIT.

package mapping_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// ReadMappingRuleReader is a Reader for the ReadMappingRule structure.
type ReadMappingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadMappingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadMappingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadMappingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadMappingRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadMappingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadMappingRuleOK creates a ReadMappingRuleOK with default headers values
func NewReadMappingRuleOK() *ReadMappingRuleOK {
	return &ReadMappingRuleOK{}
}

/*
ReadMappingRuleOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadMappingRuleOK struct {
	Payload *models.Configv1ReadMappingRuleResponse
}

// IsSuccess returns true when this read mapping rule o k response has a 2xx status code
func (o *ReadMappingRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read mapping rule o k response has a 3xx status code
func (o *ReadMappingRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read mapping rule o k response has a 4xx status code
func (o *ReadMappingRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read mapping rule o k response has a 5xx status code
func (o *ReadMappingRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read mapping rule o k response a status code equal to that given
func (o *ReadMappingRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read mapping rule o k response
func (o *ReadMappingRuleOK) Code() int {
	return 200
}

func (o *ReadMappingRuleOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules/{slug}][%d] readMappingRuleOK  %+v", 200, o.Payload)
}

func (o *ReadMappingRuleOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules/{slug}][%d] readMappingRuleOK  %+v", 200, o.Payload)
}

func (o *ReadMappingRuleOK) GetPayload() *models.Configv1ReadMappingRuleResponse {
	return o.Payload
}

func (o *ReadMappingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ReadMappingRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadMappingRuleNotFound creates a ReadMappingRuleNotFound with default headers values
func NewReadMappingRuleNotFound() *ReadMappingRuleNotFound {
	return &ReadMappingRuleNotFound{}
}

/*
ReadMappingRuleNotFound describes a response with status code 404, with default header values.

Cannot read the MappingRule because the slug does not exist.
*/
type ReadMappingRuleNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read mapping rule not found response has a 2xx status code
func (o *ReadMappingRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read mapping rule not found response has a 3xx status code
func (o *ReadMappingRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read mapping rule not found response has a 4xx status code
func (o *ReadMappingRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read mapping rule not found response has a 5xx status code
func (o *ReadMappingRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read mapping rule not found response a status code equal to that given
func (o *ReadMappingRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read mapping rule not found response
func (o *ReadMappingRuleNotFound) Code() int {
	return 404
}

func (o *ReadMappingRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules/{slug}][%d] readMappingRuleNotFound  %+v", 404, o.Payload)
}

func (o *ReadMappingRuleNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules/{slug}][%d] readMappingRuleNotFound  %+v", 404, o.Payload)
}

func (o *ReadMappingRuleNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadMappingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadMappingRuleInternalServerError creates a ReadMappingRuleInternalServerError with default headers values
func NewReadMappingRuleInternalServerError() *ReadMappingRuleInternalServerError {
	return &ReadMappingRuleInternalServerError{}
}

/*
ReadMappingRuleInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadMappingRuleInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read mapping rule internal server error response has a 2xx status code
func (o *ReadMappingRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read mapping rule internal server error response has a 3xx status code
func (o *ReadMappingRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read mapping rule internal server error response has a 4xx status code
func (o *ReadMappingRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read mapping rule internal server error response has a 5xx status code
func (o *ReadMappingRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read mapping rule internal server error response a status code equal to that given
func (o *ReadMappingRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read mapping rule internal server error response
func (o *ReadMappingRuleInternalServerError) Code() int {
	return 500
}

func (o *ReadMappingRuleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules/{slug}][%d] readMappingRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadMappingRuleInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules/{slug}][%d] readMappingRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadMappingRuleInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadMappingRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadMappingRuleDefault creates a ReadMappingRuleDefault with default headers values
func NewReadMappingRuleDefault(code int) *ReadMappingRuleDefault {
	return &ReadMappingRuleDefault{
		_statusCode: code,
	}
}

/*
ReadMappingRuleDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadMappingRuleDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read mapping rule default response has a 2xx status code
func (o *ReadMappingRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read mapping rule default response has a 3xx status code
func (o *ReadMappingRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read mapping rule default response has a 4xx status code
func (o *ReadMappingRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read mapping rule default response has a 5xx status code
func (o *ReadMappingRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read mapping rule default response a status code equal to that given
func (o *ReadMappingRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read mapping rule default response
func (o *ReadMappingRuleDefault) Code() int {
	return o._statusCode
}

func (o *ReadMappingRuleDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules/{slug}][%d] ReadMappingRule default  %+v", o._statusCode, o.Payload)
}

func (o *ReadMappingRuleDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules/{slug}][%d] ReadMappingRule default  %+v", o._statusCode, o.Payload)
}

func (o *ReadMappingRuleDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadMappingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

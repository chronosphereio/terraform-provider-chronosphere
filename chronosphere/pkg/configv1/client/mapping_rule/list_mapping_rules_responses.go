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

// ListMappingRulesReader is a Reader for the ListMappingRules structure.
type ListMappingRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMappingRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMappingRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListMappingRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListMappingRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListMappingRulesOK creates a ListMappingRulesOK with default headers values
func NewListMappingRulesOK() *ListMappingRulesOK {
	return &ListMappingRulesOK{}
}

/*
ListMappingRulesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListMappingRulesOK struct {
	Payload *models.Configv1ListMappingRulesResponse
}

// IsSuccess returns true when this list mapping rules o k response has a 2xx status code
func (o *ListMappingRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list mapping rules o k response has a 3xx status code
func (o *ListMappingRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list mapping rules o k response has a 4xx status code
func (o *ListMappingRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list mapping rules o k response has a 5xx status code
func (o *ListMappingRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list mapping rules o k response a status code equal to that given
func (o *ListMappingRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list mapping rules o k response
func (o *ListMappingRulesOK) Code() int {
	return 200
}

func (o *ListMappingRulesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules][%d] listMappingRulesOK  %+v", 200, o.Payload)
}

func (o *ListMappingRulesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules][%d] listMappingRulesOK  %+v", 200, o.Payload)
}

func (o *ListMappingRulesOK) GetPayload() *models.Configv1ListMappingRulesResponse {
	return o.Payload
}

func (o *ListMappingRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ListMappingRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMappingRulesInternalServerError creates a ListMappingRulesInternalServerError with default headers values
func NewListMappingRulesInternalServerError() *ListMappingRulesInternalServerError {
	return &ListMappingRulesInternalServerError{}
}

/*
ListMappingRulesInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ListMappingRulesInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this list mapping rules internal server error response has a 2xx status code
func (o *ListMappingRulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list mapping rules internal server error response has a 3xx status code
func (o *ListMappingRulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list mapping rules internal server error response has a 4xx status code
func (o *ListMappingRulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list mapping rules internal server error response has a 5xx status code
func (o *ListMappingRulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list mapping rules internal server error response a status code equal to that given
func (o *ListMappingRulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list mapping rules internal server error response
func (o *ListMappingRulesInternalServerError) Code() int {
	return 500
}

func (o *ListMappingRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules][%d] listMappingRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListMappingRulesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules][%d] listMappingRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListMappingRulesInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListMappingRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMappingRulesDefault creates a ListMappingRulesDefault with default headers values
func NewListMappingRulesDefault(code int) *ListMappingRulesDefault {
	return &ListMappingRulesDefault{
		_statusCode: code,
	}
}

/*
ListMappingRulesDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ListMappingRulesDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this list mapping rules default response has a 2xx status code
func (o *ListMappingRulesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list mapping rules default response has a 3xx status code
func (o *ListMappingRulesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list mapping rules default response has a 4xx status code
func (o *ListMappingRulesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list mapping rules default response has a 5xx status code
func (o *ListMappingRulesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list mapping rules default response a status code equal to that given
func (o *ListMappingRulesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list mapping rules default response
func (o *ListMappingRulesDefault) Code() int {
	return o._statusCode
}

func (o *ListMappingRulesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules][%d] ListMappingRules default  %+v", o._statusCode, o.Payload)
}

func (o *ListMappingRulesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/mapping-rules][%d] ListMappingRules default  %+v", o._statusCode, o.Payload)
}

func (o *ListMappingRulesDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ListMappingRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

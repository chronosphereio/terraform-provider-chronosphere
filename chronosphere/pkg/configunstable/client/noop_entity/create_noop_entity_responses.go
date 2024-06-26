// Code generated by go-swagger; DO NOT EDIT.

package noop_entity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

// CreateNoopEntityReader is a Reader for the CreateNoopEntity structure.
type CreateNoopEntityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNoopEntityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNoopEntityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateNoopEntityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateNoopEntityConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateNoopEntityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateNoopEntityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateNoopEntityOK creates a CreateNoopEntityOK with default headers values
func NewCreateNoopEntityOK() *CreateNoopEntityOK {
	return &CreateNoopEntityOK{}
}

/*
CreateNoopEntityOK describes a response with status code 200, with default header values.

A successful response containing the created NoopEntity.
*/
type CreateNoopEntityOK struct {
	Payload *models.ConfigunstableCreateNoopEntityResponse
}

// IsSuccess returns true when this create noop entity o k response has a 2xx status code
func (o *CreateNoopEntityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create noop entity o k response has a 3xx status code
func (o *CreateNoopEntityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create noop entity o k response has a 4xx status code
func (o *CreateNoopEntityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create noop entity o k response has a 5xx status code
func (o *CreateNoopEntityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create noop entity o k response a status code equal to that given
func (o *CreateNoopEntityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create noop entity o k response
func (o *CreateNoopEntityOK) Code() int {
	return 200
}

func (o *CreateNoopEntityOK) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/noop-entities][%d] createNoopEntityOK  %+v", 200, o.Payload)
}

func (o *CreateNoopEntityOK) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/noop-entities][%d] createNoopEntityOK  %+v", 200, o.Payload)
}

func (o *CreateNoopEntityOK) GetPayload() *models.ConfigunstableCreateNoopEntityResponse {
	return o.Payload
}

func (o *CreateNoopEntityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableCreateNoopEntityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNoopEntityBadRequest creates a CreateNoopEntityBadRequest with default headers values
func NewCreateNoopEntityBadRequest() *CreateNoopEntityBadRequest {
	return &CreateNoopEntityBadRequest{}
}

/*
CreateNoopEntityBadRequest describes a response with status code 400, with default header values.

Cannot create the NoopEntity because the request is invalid.
*/
type CreateNoopEntityBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create noop entity bad request response has a 2xx status code
func (o *CreateNoopEntityBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create noop entity bad request response has a 3xx status code
func (o *CreateNoopEntityBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create noop entity bad request response has a 4xx status code
func (o *CreateNoopEntityBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create noop entity bad request response has a 5xx status code
func (o *CreateNoopEntityBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create noop entity bad request response a status code equal to that given
func (o *CreateNoopEntityBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create noop entity bad request response
func (o *CreateNoopEntityBadRequest) Code() int {
	return 400
}

func (o *CreateNoopEntityBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/noop-entities][%d] createNoopEntityBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNoopEntityBadRequest) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/noop-entities][%d] createNoopEntityBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNoopEntityBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateNoopEntityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNoopEntityConflict creates a CreateNoopEntityConflict with default headers values
func NewCreateNoopEntityConflict() *CreateNoopEntityConflict {
	return &CreateNoopEntityConflict{}
}

/*
CreateNoopEntityConflict describes a response with status code 409, with default header values.

Cannot create the NoopEntity because there is a conflict with an existing NoopEntity.
*/
type CreateNoopEntityConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create noop entity conflict response has a 2xx status code
func (o *CreateNoopEntityConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create noop entity conflict response has a 3xx status code
func (o *CreateNoopEntityConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create noop entity conflict response has a 4xx status code
func (o *CreateNoopEntityConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create noop entity conflict response has a 5xx status code
func (o *CreateNoopEntityConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create noop entity conflict response a status code equal to that given
func (o *CreateNoopEntityConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create noop entity conflict response
func (o *CreateNoopEntityConflict) Code() int {
	return 409
}

func (o *CreateNoopEntityConflict) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/noop-entities][%d] createNoopEntityConflict  %+v", 409, o.Payload)
}

func (o *CreateNoopEntityConflict) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/noop-entities][%d] createNoopEntityConflict  %+v", 409, o.Payload)
}

func (o *CreateNoopEntityConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateNoopEntityConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNoopEntityInternalServerError creates a CreateNoopEntityInternalServerError with default headers values
func NewCreateNoopEntityInternalServerError() *CreateNoopEntityInternalServerError {
	return &CreateNoopEntityInternalServerError{}
}

/*
CreateNoopEntityInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateNoopEntityInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create noop entity internal server error response has a 2xx status code
func (o *CreateNoopEntityInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create noop entity internal server error response has a 3xx status code
func (o *CreateNoopEntityInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create noop entity internal server error response has a 4xx status code
func (o *CreateNoopEntityInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create noop entity internal server error response has a 5xx status code
func (o *CreateNoopEntityInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create noop entity internal server error response a status code equal to that given
func (o *CreateNoopEntityInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create noop entity internal server error response
func (o *CreateNoopEntityInternalServerError) Code() int {
	return 500
}

func (o *CreateNoopEntityInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/noop-entities][%d] createNoopEntityInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateNoopEntityInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/noop-entities][%d] createNoopEntityInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateNoopEntityInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateNoopEntityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNoopEntityDefault creates a CreateNoopEntityDefault with default headers values
func NewCreateNoopEntityDefault(code int) *CreateNoopEntityDefault {
	return &CreateNoopEntityDefault{
		_statusCode: code,
	}
}

/*
CreateNoopEntityDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateNoopEntityDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create noop entity default response has a 2xx status code
func (o *CreateNoopEntityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create noop entity default response has a 3xx status code
func (o *CreateNoopEntityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create noop entity default response has a 4xx status code
func (o *CreateNoopEntityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create noop entity default response has a 5xx status code
func (o *CreateNoopEntityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create noop entity default response a status code equal to that given
func (o *CreateNoopEntityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create noop entity default response
func (o *CreateNoopEntityDefault) Code() int {
	return o._statusCode
}

func (o *CreateNoopEntityDefault) Error() string {
	return fmt.Sprintf("[POST /api/unstable/config/noop-entities][%d] CreateNoopEntity default  %+v", o._statusCode, o.Payload)
}

func (o *CreateNoopEntityDefault) String() string {
	return fmt.Sprintf("[POST /api/unstable/config/noop-entities][%d] CreateNoopEntity default  %+v", o._statusCode, o.Payload)
}

func (o *CreateNoopEntityDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateNoopEntityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

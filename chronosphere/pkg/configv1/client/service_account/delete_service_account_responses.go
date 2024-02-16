// Code generated by go-swagger; DO NOT EDIT.

package service_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// DeleteServiceAccountReader is a Reader for the DeleteServiceAccount structure.
type DeleteServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteServiceAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteServiceAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteServiceAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteServiceAccountOK creates a DeleteServiceAccountOK with default headers values
func NewDeleteServiceAccountOK() *DeleteServiceAccountOK {
	return &DeleteServiceAccountOK{}
}

/*
DeleteServiceAccountOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteServiceAccountOK struct {
	Payload models.Configv1DeleteServiceAccountResponse
}

// IsSuccess returns true when this delete service account o k response has a 2xx status code
func (o *DeleteServiceAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete service account o k response has a 3xx status code
func (o *DeleteServiceAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service account o k response has a 4xx status code
func (o *DeleteServiceAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete service account o k response has a 5xx status code
func (o *DeleteServiceAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service account o k response a status code equal to that given
func (o *DeleteServiceAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete service account o k response
func (o *DeleteServiceAccountOK) Code() int {
	return 200
}

func (o *DeleteServiceAccountOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/service-accounts/{slug}][%d] deleteServiceAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteServiceAccountOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/service-accounts/{slug}][%d] deleteServiceAccountOK  %+v", 200, o.Payload)
}

func (o *DeleteServiceAccountOK) GetPayload() models.Configv1DeleteServiceAccountResponse {
	return o.Payload
}

func (o *DeleteServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountNotFound creates a DeleteServiceAccountNotFound with default headers values
func NewDeleteServiceAccountNotFound() *DeleteServiceAccountNotFound {
	return &DeleteServiceAccountNotFound{}
}

/*
DeleteServiceAccountNotFound describes a response with status code 404, with default header values.

Cannot delete the ServiceAccount because the slug does not exist.
*/
type DeleteServiceAccountNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete service account not found response has a 2xx status code
func (o *DeleteServiceAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service account not found response has a 3xx status code
func (o *DeleteServiceAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service account not found response has a 4xx status code
func (o *DeleteServiceAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete service account not found response has a 5xx status code
func (o *DeleteServiceAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete service account not found response a status code equal to that given
func (o *DeleteServiceAccountNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete service account not found response
func (o *DeleteServiceAccountNotFound) Code() int {
	return 404
}

func (o *DeleteServiceAccountNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/service-accounts/{slug}][%d] deleteServiceAccountNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServiceAccountNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/service-accounts/{slug}][%d] deleteServiceAccountNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServiceAccountNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteServiceAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountInternalServerError creates a DeleteServiceAccountInternalServerError with default headers values
func NewDeleteServiceAccountInternalServerError() *DeleteServiceAccountInternalServerError {
	return &DeleteServiceAccountInternalServerError{}
}

/*
DeleteServiceAccountInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteServiceAccountInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete service account internal server error response has a 2xx status code
func (o *DeleteServiceAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete service account internal server error response has a 3xx status code
func (o *DeleteServiceAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete service account internal server error response has a 4xx status code
func (o *DeleteServiceAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete service account internal server error response has a 5xx status code
func (o *DeleteServiceAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete service account internal server error response a status code equal to that given
func (o *DeleteServiceAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete service account internal server error response
func (o *DeleteServiceAccountInternalServerError) Code() int {
	return 500
}

func (o *DeleteServiceAccountInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/service-accounts/{slug}][%d] deleteServiceAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteServiceAccountInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/service-accounts/{slug}][%d] deleteServiceAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteServiceAccountInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteServiceAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteServiceAccountDefault creates a DeleteServiceAccountDefault with default headers values
func NewDeleteServiceAccountDefault(code int) *DeleteServiceAccountDefault {
	return &DeleteServiceAccountDefault{
		_statusCode: code,
	}
}

/*
DeleteServiceAccountDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteServiceAccountDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete service account default response has a 2xx status code
func (o *DeleteServiceAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete service account default response has a 3xx status code
func (o *DeleteServiceAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete service account default response has a 4xx status code
func (o *DeleteServiceAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete service account default response has a 5xx status code
func (o *DeleteServiceAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete service account default response a status code equal to that given
func (o *DeleteServiceAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete service account default response
func (o *DeleteServiceAccountDefault) Code() int {
	return o._statusCode
}

func (o *DeleteServiceAccountDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/service-accounts/{slug}][%d] DeleteServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteServiceAccountDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/service-accounts/{slug}][%d] DeleteServiceAccount default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteServiceAccountDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteServiceAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
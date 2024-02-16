// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// DeleteTeamReader is a Reader for the DeleteTeam structure.
type DeleteTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTeamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTeamOK creates a DeleteTeamOK with default headers values
func NewDeleteTeamOK() *DeleteTeamOK {
	return &DeleteTeamOK{}
}

/*
DeleteTeamOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteTeamOK struct {
	Payload models.Configv1DeleteTeamResponse
}

// IsSuccess returns true when this delete team o k response has a 2xx status code
func (o *DeleteTeamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete team o k response has a 3xx status code
func (o *DeleteTeamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team o k response has a 4xx status code
func (o *DeleteTeamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team o k response has a 5xx status code
func (o *DeleteTeamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team o k response a status code equal to that given
func (o *DeleteTeamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete team o k response
func (o *DeleteTeamOK) Code() int {
	return 200
}

func (o *DeleteTeamOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/teams/{slug}][%d] deleteTeamOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/teams/{slug}][%d] deleteTeamOK  %+v", 200, o.Payload)
}

func (o *DeleteTeamOK) GetPayload() models.Configv1DeleteTeamResponse {
	return o.Payload
}

func (o *DeleteTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamNotFound creates a DeleteTeamNotFound with default headers values
func NewDeleteTeamNotFound() *DeleteTeamNotFound {
	return &DeleteTeamNotFound{}
}

/*
DeleteTeamNotFound describes a response with status code 404, with default header values.

Cannot delete the Team because the slug does not exist.
*/
type DeleteTeamNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete team not found response has a 2xx status code
func (o *DeleteTeamNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team not found response has a 3xx status code
func (o *DeleteTeamNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team not found response has a 4xx status code
func (o *DeleteTeamNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete team not found response has a 5xx status code
func (o *DeleteTeamNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete team not found response a status code equal to that given
func (o *DeleteTeamNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete team not found response
func (o *DeleteTeamNotFound) Code() int {
	return 404
}

func (o *DeleteTeamNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/teams/{slug}][%d] deleteTeamNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/teams/{slug}][%d] deleteTeamNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTeamNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamInternalServerError creates a DeleteTeamInternalServerError with default headers values
func NewDeleteTeamInternalServerError() *DeleteTeamInternalServerError {
	return &DeleteTeamInternalServerError{}
}

/*
DeleteTeamInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteTeamInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete team internal server error response has a 2xx status code
func (o *DeleteTeamInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete team internal server error response has a 3xx status code
func (o *DeleteTeamInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete team internal server error response has a 4xx status code
func (o *DeleteTeamInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete team internal server error response has a 5xx status code
func (o *DeleteTeamInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete team internal server error response a status code equal to that given
func (o *DeleteTeamInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete team internal server error response
func (o *DeleteTeamInternalServerError) Code() int {
	return 500
}

func (o *DeleteTeamInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/teams/{slug}][%d] deleteTeamInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTeamInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/teams/{slug}][%d] deleteTeamInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTeamInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteTeamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTeamDefault creates a DeleteTeamDefault with default headers values
func NewDeleteTeamDefault(code int) *DeleteTeamDefault {
	return &DeleteTeamDefault{
		_statusCode: code,
	}
}

/*
DeleteTeamDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteTeamDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete team default response has a 2xx status code
func (o *DeleteTeamDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete team default response has a 3xx status code
func (o *DeleteTeamDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete team default response has a 4xx status code
func (o *DeleteTeamDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete team default response has a 5xx status code
func (o *DeleteTeamDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete team default response a status code equal to that given
func (o *DeleteTeamDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete team default response
func (o *DeleteTeamDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTeamDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/teams/{slug}][%d] DeleteTeam default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTeamDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/teams/{slug}][%d] DeleteTeam default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTeamDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
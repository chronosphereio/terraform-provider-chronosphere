// Code generated by go-swagger; DO NOT EDIT.

package trace_behavior

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

// UpdateTraceBehaviorReader is a Reader for the UpdateTraceBehavior structure.
type UpdateTraceBehaviorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTraceBehaviorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTraceBehaviorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTraceBehaviorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTraceBehaviorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateTraceBehaviorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTraceBehaviorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateTraceBehaviorDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTraceBehaviorOK creates a UpdateTraceBehaviorOK with default headers values
func NewUpdateTraceBehaviorOK() *UpdateTraceBehaviorOK {
	return &UpdateTraceBehaviorOK{}
}

/*
UpdateTraceBehaviorOK describes a response with status code 200, with default header values.

A successful response containing the updated TraceBehavior.
*/
type UpdateTraceBehaviorOK struct {
	Payload *models.ConfigunstableUpdateTraceBehaviorResponse
}

// IsSuccess returns true when this update trace behavior o k response has a 2xx status code
func (o *UpdateTraceBehaviorOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update trace behavior o k response has a 3xx status code
func (o *UpdateTraceBehaviorOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace behavior o k response has a 4xx status code
func (o *UpdateTraceBehaviorOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trace behavior o k response has a 5xx status code
func (o *UpdateTraceBehaviorOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace behavior o k response a status code equal to that given
func (o *UpdateTraceBehaviorOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update trace behavior o k response
func (o *UpdateTraceBehaviorOK) Code() int {
	return 200
}

func (o *UpdateTraceBehaviorOK) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-behaviors/{slug}][%d] updateTraceBehaviorOK  %+v", 200, o.Payload)
}

func (o *UpdateTraceBehaviorOK) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-behaviors/{slug}][%d] updateTraceBehaviorOK  %+v", 200, o.Payload)
}

func (o *UpdateTraceBehaviorOK) GetPayload() *models.ConfigunstableUpdateTraceBehaviorResponse {
	return o.Payload
}

func (o *UpdateTraceBehaviorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableUpdateTraceBehaviorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceBehaviorBadRequest creates a UpdateTraceBehaviorBadRequest with default headers values
func NewUpdateTraceBehaviorBadRequest() *UpdateTraceBehaviorBadRequest {
	return &UpdateTraceBehaviorBadRequest{}
}

/*
UpdateTraceBehaviorBadRequest describes a response with status code 400, with default header values.

Cannot update the TraceBehavior because the request is invalid.
*/
type UpdateTraceBehaviorBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace behavior bad request response has a 2xx status code
func (o *UpdateTraceBehaviorBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace behavior bad request response has a 3xx status code
func (o *UpdateTraceBehaviorBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace behavior bad request response has a 4xx status code
func (o *UpdateTraceBehaviorBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trace behavior bad request response has a 5xx status code
func (o *UpdateTraceBehaviorBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace behavior bad request response a status code equal to that given
func (o *UpdateTraceBehaviorBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update trace behavior bad request response
func (o *UpdateTraceBehaviorBadRequest) Code() int {
	return 400
}

func (o *UpdateTraceBehaviorBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-behaviors/{slug}][%d] updateTraceBehaviorBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTraceBehaviorBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-behaviors/{slug}][%d] updateTraceBehaviorBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTraceBehaviorBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceBehaviorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceBehaviorNotFound creates a UpdateTraceBehaviorNotFound with default headers values
func NewUpdateTraceBehaviorNotFound() *UpdateTraceBehaviorNotFound {
	return &UpdateTraceBehaviorNotFound{}
}

/*
UpdateTraceBehaviorNotFound describes a response with status code 404, with default header values.

Cannot update the TraceBehavior because the slug does not exist.
*/
type UpdateTraceBehaviorNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace behavior not found response has a 2xx status code
func (o *UpdateTraceBehaviorNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace behavior not found response has a 3xx status code
func (o *UpdateTraceBehaviorNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace behavior not found response has a 4xx status code
func (o *UpdateTraceBehaviorNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trace behavior not found response has a 5xx status code
func (o *UpdateTraceBehaviorNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace behavior not found response a status code equal to that given
func (o *UpdateTraceBehaviorNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update trace behavior not found response
func (o *UpdateTraceBehaviorNotFound) Code() int {
	return 404
}

func (o *UpdateTraceBehaviorNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-behaviors/{slug}][%d] updateTraceBehaviorNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTraceBehaviorNotFound) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-behaviors/{slug}][%d] updateTraceBehaviorNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTraceBehaviorNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceBehaviorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceBehaviorConflict creates a UpdateTraceBehaviorConflict with default headers values
func NewUpdateTraceBehaviorConflict() *UpdateTraceBehaviorConflict {
	return &UpdateTraceBehaviorConflict{}
}

/*
UpdateTraceBehaviorConflict describes a response with status code 409, with default header values.

Cannot update the TraceBehavior because there is a conflict with an existing TraceBehavior.
*/
type UpdateTraceBehaviorConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace behavior conflict response has a 2xx status code
func (o *UpdateTraceBehaviorConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace behavior conflict response has a 3xx status code
func (o *UpdateTraceBehaviorConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace behavior conflict response has a 4xx status code
func (o *UpdateTraceBehaviorConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trace behavior conflict response has a 5xx status code
func (o *UpdateTraceBehaviorConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace behavior conflict response a status code equal to that given
func (o *UpdateTraceBehaviorConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update trace behavior conflict response
func (o *UpdateTraceBehaviorConflict) Code() int {
	return 409
}

func (o *UpdateTraceBehaviorConflict) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-behaviors/{slug}][%d] updateTraceBehaviorConflict  %+v", 409, o.Payload)
}

func (o *UpdateTraceBehaviorConflict) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-behaviors/{slug}][%d] updateTraceBehaviorConflict  %+v", 409, o.Payload)
}

func (o *UpdateTraceBehaviorConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceBehaviorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceBehaviorInternalServerError creates a UpdateTraceBehaviorInternalServerError with default headers values
func NewUpdateTraceBehaviorInternalServerError() *UpdateTraceBehaviorInternalServerError {
	return &UpdateTraceBehaviorInternalServerError{}
}

/*
UpdateTraceBehaviorInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateTraceBehaviorInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace behavior internal server error response has a 2xx status code
func (o *UpdateTraceBehaviorInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace behavior internal server error response has a 3xx status code
func (o *UpdateTraceBehaviorInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace behavior internal server error response has a 4xx status code
func (o *UpdateTraceBehaviorInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trace behavior internal server error response has a 5xx status code
func (o *UpdateTraceBehaviorInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update trace behavior internal server error response a status code equal to that given
func (o *UpdateTraceBehaviorInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update trace behavior internal server error response
func (o *UpdateTraceBehaviorInternalServerError) Code() int {
	return 500
}

func (o *UpdateTraceBehaviorInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-behaviors/{slug}][%d] updateTraceBehaviorInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTraceBehaviorInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-behaviors/{slug}][%d] updateTraceBehaviorInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTraceBehaviorInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceBehaviorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceBehaviorDefault creates a UpdateTraceBehaviorDefault with default headers values
func NewUpdateTraceBehaviorDefault(code int) *UpdateTraceBehaviorDefault {
	return &UpdateTraceBehaviorDefault{
		_statusCode: code,
	}
}

/*
UpdateTraceBehaviorDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateTraceBehaviorDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update trace behavior default response has a 2xx status code
func (o *UpdateTraceBehaviorDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update trace behavior default response has a 3xx status code
func (o *UpdateTraceBehaviorDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update trace behavior default response has a 4xx status code
func (o *UpdateTraceBehaviorDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update trace behavior default response has a 5xx status code
func (o *UpdateTraceBehaviorDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update trace behavior default response a status code equal to that given
func (o *UpdateTraceBehaviorDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update trace behavior default response
func (o *UpdateTraceBehaviorDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTraceBehaviorDefault) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-behaviors/{slug}][%d] UpdateTraceBehavior default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTraceBehaviorDefault) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-behaviors/{slug}][%d] UpdateTraceBehavior default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTraceBehaviorDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateTraceBehaviorDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

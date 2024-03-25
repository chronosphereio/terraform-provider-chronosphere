// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// DeleteMonitorReader is a Reader for the DeleteMonitor structure.
type DeleteMonitorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMonitorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMonitorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteMonitorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteMonitorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteMonitorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteMonitorDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMonitorOK creates a DeleteMonitorOK with default headers values
func NewDeleteMonitorOK() *DeleteMonitorOK {
	return &DeleteMonitorOK{}
}

/*
DeleteMonitorOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteMonitorOK struct {
	Payload models.Configv1DeleteMonitorResponse
}

// IsSuccess returns true when this delete monitor o k response has a 2xx status code
func (o *DeleteMonitorOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete monitor o k response has a 3xx status code
func (o *DeleteMonitorOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete monitor o k response has a 4xx status code
func (o *DeleteMonitorOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete monitor o k response has a 5xx status code
func (o *DeleteMonitorOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete monitor o k response a status code equal to that given
func (o *DeleteMonitorOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete monitor o k response
func (o *DeleteMonitorOK) Code() int {
	return 200
}

func (o *DeleteMonitorOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/monitors/{slug}][%d] deleteMonitorOK  %+v", 200, o.Payload)
}

func (o *DeleteMonitorOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/monitors/{slug}][%d] deleteMonitorOK  %+v", 200, o.Payload)
}

func (o *DeleteMonitorOK) GetPayload() models.Configv1DeleteMonitorResponse {
	return o.Payload
}

func (o *DeleteMonitorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMonitorBadRequest creates a DeleteMonitorBadRequest with default headers values
func NewDeleteMonitorBadRequest() *DeleteMonitorBadRequest {
	return &DeleteMonitorBadRequest{}
}

/*
DeleteMonitorBadRequest describes a response with status code 400, with default header values.

Cannot delete the Monitor because it is in use.
*/
type DeleteMonitorBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete monitor bad request response has a 2xx status code
func (o *DeleteMonitorBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete monitor bad request response has a 3xx status code
func (o *DeleteMonitorBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete monitor bad request response has a 4xx status code
func (o *DeleteMonitorBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete monitor bad request response has a 5xx status code
func (o *DeleteMonitorBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete monitor bad request response a status code equal to that given
func (o *DeleteMonitorBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete monitor bad request response
func (o *DeleteMonitorBadRequest) Code() int {
	return 400
}

func (o *DeleteMonitorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/monitors/{slug}][%d] deleteMonitorBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteMonitorBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/monitors/{slug}][%d] deleteMonitorBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteMonitorBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteMonitorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMonitorNotFound creates a DeleteMonitorNotFound with default headers values
func NewDeleteMonitorNotFound() *DeleteMonitorNotFound {
	return &DeleteMonitorNotFound{}
}

/*
DeleteMonitorNotFound describes a response with status code 404, with default header values.

Cannot delete the Monitor because the slug does not exist.
*/
type DeleteMonitorNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete monitor not found response has a 2xx status code
func (o *DeleteMonitorNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete monitor not found response has a 3xx status code
func (o *DeleteMonitorNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete monitor not found response has a 4xx status code
func (o *DeleteMonitorNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete monitor not found response has a 5xx status code
func (o *DeleteMonitorNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete monitor not found response a status code equal to that given
func (o *DeleteMonitorNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete monitor not found response
func (o *DeleteMonitorNotFound) Code() int {
	return 404
}

func (o *DeleteMonitorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/monitors/{slug}][%d] deleteMonitorNotFound  %+v", 404, o.Payload)
}

func (o *DeleteMonitorNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/monitors/{slug}][%d] deleteMonitorNotFound  %+v", 404, o.Payload)
}

func (o *DeleteMonitorNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteMonitorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMonitorInternalServerError creates a DeleteMonitorInternalServerError with default headers values
func NewDeleteMonitorInternalServerError() *DeleteMonitorInternalServerError {
	return &DeleteMonitorInternalServerError{}
}

/*
DeleteMonitorInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteMonitorInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete monitor internal server error response has a 2xx status code
func (o *DeleteMonitorInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete monitor internal server error response has a 3xx status code
func (o *DeleteMonitorInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete monitor internal server error response has a 4xx status code
func (o *DeleteMonitorInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete monitor internal server error response has a 5xx status code
func (o *DeleteMonitorInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete monitor internal server error response a status code equal to that given
func (o *DeleteMonitorInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete monitor internal server error response
func (o *DeleteMonitorInternalServerError) Code() int {
	return 500
}

func (o *DeleteMonitorInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/monitors/{slug}][%d] deleteMonitorInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteMonitorInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/monitors/{slug}][%d] deleteMonitorInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteMonitorInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteMonitorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMonitorDefault creates a DeleteMonitorDefault with default headers values
func NewDeleteMonitorDefault(code int) *DeleteMonitorDefault {
	return &DeleteMonitorDefault{
		_statusCode: code,
	}
}

/*
DeleteMonitorDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteMonitorDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete monitor default response has a 2xx status code
func (o *DeleteMonitorDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete monitor default response has a 3xx status code
func (o *DeleteMonitorDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete monitor default response has a 4xx status code
func (o *DeleteMonitorDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete monitor default response has a 5xx status code
func (o *DeleteMonitorDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete monitor default response a status code equal to that given
func (o *DeleteMonitorDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete monitor default response
func (o *DeleteMonitorDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMonitorDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/monitors/{slug}][%d] DeleteMonitor default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMonitorDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/monitors/{slug}][%d] DeleteMonitor default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMonitorDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteMonitorDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

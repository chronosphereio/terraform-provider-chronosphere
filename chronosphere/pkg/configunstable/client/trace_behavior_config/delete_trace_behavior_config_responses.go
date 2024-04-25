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

// DeleteTraceBehaviorConfigReader is a Reader for the DeleteTraceBehaviorConfig structure.
type DeleteTraceBehaviorConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTraceBehaviorConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTraceBehaviorConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTraceBehaviorConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTraceBehaviorConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTraceBehaviorConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTraceBehaviorConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTraceBehaviorConfigOK creates a DeleteTraceBehaviorConfigOK with default headers values
func NewDeleteTraceBehaviorConfigOK() *DeleteTraceBehaviorConfigOK {
	return &DeleteTraceBehaviorConfigOK{}
}

/*
DeleteTraceBehaviorConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteTraceBehaviorConfigOK struct {
	Payload models.ConfigunstableDeleteTraceBehaviorConfigResponse
}

// IsSuccess returns true when this delete trace behavior config o k response has a 2xx status code
func (o *DeleteTraceBehaviorConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete trace behavior config o k response has a 3xx status code
func (o *DeleteTraceBehaviorConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace behavior config o k response has a 4xx status code
func (o *DeleteTraceBehaviorConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete trace behavior config o k response has a 5xx status code
func (o *DeleteTraceBehaviorConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete trace behavior config o k response a status code equal to that given
func (o *DeleteTraceBehaviorConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete trace behavior config o k response
func (o *DeleteTraceBehaviorConfigOK) Code() int {
	return 200
}

func (o *DeleteTraceBehaviorConfigOK) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-behavior-config][%d] deleteTraceBehaviorConfigOK  %+v", 200, o.Payload)
}

func (o *DeleteTraceBehaviorConfigOK) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-behavior-config][%d] deleteTraceBehaviorConfigOK  %+v", 200, o.Payload)
}

func (o *DeleteTraceBehaviorConfigOK) GetPayload() models.ConfigunstableDeleteTraceBehaviorConfigResponse {
	return o.Payload
}

func (o *DeleteTraceBehaviorConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceBehaviorConfigBadRequest creates a DeleteTraceBehaviorConfigBadRequest with default headers values
func NewDeleteTraceBehaviorConfigBadRequest() *DeleteTraceBehaviorConfigBadRequest {
	return &DeleteTraceBehaviorConfigBadRequest{}
}

/*
DeleteTraceBehaviorConfigBadRequest describes a response with status code 400, with default header values.

Cannot delete the TraceBehaviorConfig because it is in use.
*/
type DeleteTraceBehaviorConfigBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete trace behavior config bad request response has a 2xx status code
func (o *DeleteTraceBehaviorConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete trace behavior config bad request response has a 3xx status code
func (o *DeleteTraceBehaviorConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace behavior config bad request response has a 4xx status code
func (o *DeleteTraceBehaviorConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete trace behavior config bad request response has a 5xx status code
func (o *DeleteTraceBehaviorConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete trace behavior config bad request response a status code equal to that given
func (o *DeleteTraceBehaviorConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete trace behavior config bad request response
func (o *DeleteTraceBehaviorConfigBadRequest) Code() int {
	return 400
}

func (o *DeleteTraceBehaviorConfigBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-behavior-config][%d] deleteTraceBehaviorConfigBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTraceBehaviorConfigBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-behavior-config][%d] deleteTraceBehaviorConfigBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTraceBehaviorConfigBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteTraceBehaviorConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceBehaviorConfigNotFound creates a DeleteTraceBehaviorConfigNotFound with default headers values
func NewDeleteTraceBehaviorConfigNotFound() *DeleteTraceBehaviorConfigNotFound {
	return &DeleteTraceBehaviorConfigNotFound{}
}

/*
DeleteTraceBehaviorConfigNotFound describes a response with status code 404, with default header values.

Cannot delete the TraceBehaviorConfig because the slug does not exist.
*/
type DeleteTraceBehaviorConfigNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete trace behavior config not found response has a 2xx status code
func (o *DeleteTraceBehaviorConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete trace behavior config not found response has a 3xx status code
func (o *DeleteTraceBehaviorConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace behavior config not found response has a 4xx status code
func (o *DeleteTraceBehaviorConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete trace behavior config not found response has a 5xx status code
func (o *DeleteTraceBehaviorConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete trace behavior config not found response a status code equal to that given
func (o *DeleteTraceBehaviorConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete trace behavior config not found response
func (o *DeleteTraceBehaviorConfigNotFound) Code() int {
	return 404
}

func (o *DeleteTraceBehaviorConfigNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-behavior-config][%d] deleteTraceBehaviorConfigNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTraceBehaviorConfigNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-behavior-config][%d] deleteTraceBehaviorConfigNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTraceBehaviorConfigNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteTraceBehaviorConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceBehaviorConfigInternalServerError creates a DeleteTraceBehaviorConfigInternalServerError with default headers values
func NewDeleteTraceBehaviorConfigInternalServerError() *DeleteTraceBehaviorConfigInternalServerError {
	return &DeleteTraceBehaviorConfigInternalServerError{}
}

/*
DeleteTraceBehaviorConfigInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteTraceBehaviorConfigInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete trace behavior config internal server error response has a 2xx status code
func (o *DeleteTraceBehaviorConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete trace behavior config internal server error response has a 3xx status code
func (o *DeleteTraceBehaviorConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace behavior config internal server error response has a 4xx status code
func (o *DeleteTraceBehaviorConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete trace behavior config internal server error response has a 5xx status code
func (o *DeleteTraceBehaviorConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete trace behavior config internal server error response a status code equal to that given
func (o *DeleteTraceBehaviorConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete trace behavior config internal server error response
func (o *DeleteTraceBehaviorConfigInternalServerError) Code() int {
	return 500
}

func (o *DeleteTraceBehaviorConfigInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-behavior-config][%d] deleteTraceBehaviorConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTraceBehaviorConfigInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-behavior-config][%d] deleteTraceBehaviorConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTraceBehaviorConfigInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteTraceBehaviorConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceBehaviorConfigDefault creates a DeleteTraceBehaviorConfigDefault with default headers values
func NewDeleteTraceBehaviorConfigDefault(code int) *DeleteTraceBehaviorConfigDefault {
	return &DeleteTraceBehaviorConfigDefault{
		_statusCode: code,
	}
}

/*
DeleteTraceBehaviorConfigDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteTraceBehaviorConfigDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete trace behavior config default response has a 2xx status code
func (o *DeleteTraceBehaviorConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete trace behavior config default response has a 3xx status code
func (o *DeleteTraceBehaviorConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete trace behavior config default response has a 4xx status code
func (o *DeleteTraceBehaviorConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete trace behavior config default response has a 5xx status code
func (o *DeleteTraceBehaviorConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete trace behavior config default response a status code equal to that given
func (o *DeleteTraceBehaviorConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete trace behavior config default response
func (o *DeleteTraceBehaviorConfigDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTraceBehaviorConfigDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-behavior-config][%d] DeleteTraceBehaviorConfig default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTraceBehaviorConfigDefault) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-behavior-config][%d] DeleteTraceBehaviorConfig default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTraceBehaviorConfigDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteTraceBehaviorConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

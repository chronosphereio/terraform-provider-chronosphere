// Code generated by go-swagger; DO NOT EDIT.

package log_control_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

// UpdateLogControlConfigReader is a Reader for the UpdateLogControlConfig structure.
type UpdateLogControlConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLogControlConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLogControlConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateLogControlConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateLogControlConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateLogControlConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateLogControlConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateLogControlConfigOK creates a UpdateLogControlConfigOK with default headers values
func NewUpdateLogControlConfigOK() *UpdateLogControlConfigOK {
	return &UpdateLogControlConfigOK{}
}

/*
UpdateLogControlConfigOK describes a response with status code 200, with default header values.

A successful response containing the updated LogControlConfig.
*/
type UpdateLogControlConfigOK struct {
	Payload *models.ConfigunstableUpdateLogControlConfigResponse
}

// IsSuccess returns true when this update log control config o k response has a 2xx status code
func (o *UpdateLogControlConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update log control config o k response has a 3xx status code
func (o *UpdateLogControlConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update log control config o k response has a 4xx status code
func (o *UpdateLogControlConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update log control config o k response has a 5xx status code
func (o *UpdateLogControlConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update log control config o k response a status code equal to that given
func (o *UpdateLogControlConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update log control config o k response
func (o *UpdateLogControlConfigOK) Code() int {
	return 200
}

func (o *UpdateLogControlConfigOK) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/log-control-config][%d] updateLogControlConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateLogControlConfigOK) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/log-control-config][%d] updateLogControlConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateLogControlConfigOK) GetPayload() *models.ConfigunstableUpdateLogControlConfigResponse {
	return o.Payload
}

func (o *UpdateLogControlConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableUpdateLogControlConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLogControlConfigBadRequest creates a UpdateLogControlConfigBadRequest with default headers values
func NewUpdateLogControlConfigBadRequest() *UpdateLogControlConfigBadRequest {
	return &UpdateLogControlConfigBadRequest{}
}

/*
UpdateLogControlConfigBadRequest describes a response with status code 400, with default header values.

Cannot update the LogControlConfig because the request is invalid.
*/
type UpdateLogControlConfigBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update log control config bad request response has a 2xx status code
func (o *UpdateLogControlConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update log control config bad request response has a 3xx status code
func (o *UpdateLogControlConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update log control config bad request response has a 4xx status code
func (o *UpdateLogControlConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update log control config bad request response has a 5xx status code
func (o *UpdateLogControlConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update log control config bad request response a status code equal to that given
func (o *UpdateLogControlConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update log control config bad request response
func (o *UpdateLogControlConfigBadRequest) Code() int {
	return 400
}

func (o *UpdateLogControlConfigBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/log-control-config][%d] updateLogControlConfigBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateLogControlConfigBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/log-control-config][%d] updateLogControlConfigBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateLogControlConfigBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateLogControlConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLogControlConfigNotFound creates a UpdateLogControlConfigNotFound with default headers values
func NewUpdateLogControlConfigNotFound() *UpdateLogControlConfigNotFound {
	return &UpdateLogControlConfigNotFound{}
}

/*
UpdateLogControlConfigNotFound describes a response with status code 404, with default header values.

Cannot update the LogControlConfig because LogControlConfig has not been created.
*/
type UpdateLogControlConfigNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update log control config not found response has a 2xx status code
func (o *UpdateLogControlConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update log control config not found response has a 3xx status code
func (o *UpdateLogControlConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update log control config not found response has a 4xx status code
func (o *UpdateLogControlConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update log control config not found response has a 5xx status code
func (o *UpdateLogControlConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update log control config not found response a status code equal to that given
func (o *UpdateLogControlConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update log control config not found response
func (o *UpdateLogControlConfigNotFound) Code() int {
	return 404
}

func (o *UpdateLogControlConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/log-control-config][%d] updateLogControlConfigNotFound  %+v", 404, o.Payload)
}

func (o *UpdateLogControlConfigNotFound) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/log-control-config][%d] updateLogControlConfigNotFound  %+v", 404, o.Payload)
}

func (o *UpdateLogControlConfigNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateLogControlConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLogControlConfigInternalServerError creates a UpdateLogControlConfigInternalServerError with default headers values
func NewUpdateLogControlConfigInternalServerError() *UpdateLogControlConfigInternalServerError {
	return &UpdateLogControlConfigInternalServerError{}
}

/*
UpdateLogControlConfigInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateLogControlConfigInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update log control config internal server error response has a 2xx status code
func (o *UpdateLogControlConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update log control config internal server error response has a 3xx status code
func (o *UpdateLogControlConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update log control config internal server error response has a 4xx status code
func (o *UpdateLogControlConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update log control config internal server error response has a 5xx status code
func (o *UpdateLogControlConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update log control config internal server error response a status code equal to that given
func (o *UpdateLogControlConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update log control config internal server error response
func (o *UpdateLogControlConfigInternalServerError) Code() int {
	return 500
}

func (o *UpdateLogControlConfigInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/log-control-config][%d] updateLogControlConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateLogControlConfigInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/log-control-config][%d] updateLogControlConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateLogControlConfigInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateLogControlConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLogControlConfigDefault creates a UpdateLogControlConfigDefault with default headers values
func NewUpdateLogControlConfigDefault(code int) *UpdateLogControlConfigDefault {
	return &UpdateLogControlConfigDefault{
		_statusCode: code,
	}
}

/*
UpdateLogControlConfigDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateLogControlConfigDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update log control config default response has a 2xx status code
func (o *UpdateLogControlConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update log control config default response has a 3xx status code
func (o *UpdateLogControlConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update log control config default response has a 4xx status code
func (o *UpdateLogControlConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update log control config default response has a 5xx status code
func (o *UpdateLogControlConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update log control config default response a status code equal to that given
func (o *UpdateLogControlConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update log control config default response
func (o *UpdateLogControlConfigDefault) Code() int {
	return o._statusCode
}

func (o *UpdateLogControlConfigDefault) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/log-control-config][%d] UpdateLogControlConfig default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateLogControlConfigDefault) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/log-control-config][%d] UpdateLogControlConfig default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateLogControlConfigDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateLogControlConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

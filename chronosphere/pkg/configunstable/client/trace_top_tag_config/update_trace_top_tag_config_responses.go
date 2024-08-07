// Code generated by go-swagger; DO NOT EDIT.

package trace_top_tag_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

// UpdateTraceTopTagConfigReader is a Reader for the UpdateTraceTopTagConfig structure.
type UpdateTraceTopTagConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTraceTopTagConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTraceTopTagConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTraceTopTagConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTraceTopTagConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTraceTopTagConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateTraceTopTagConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTraceTopTagConfigOK creates a UpdateTraceTopTagConfigOK with default headers values
func NewUpdateTraceTopTagConfigOK() *UpdateTraceTopTagConfigOK {
	return &UpdateTraceTopTagConfigOK{}
}

/*
UpdateTraceTopTagConfigOK describes a response with status code 200, with default header values.

A successful response containing the updated TraceTopTagConfig.
*/
type UpdateTraceTopTagConfigOK struct {
	Payload *models.ConfigunstableUpdateTraceTopTagConfigResponse
}

// IsSuccess returns true when this update trace top tag config o k response has a 2xx status code
func (o *UpdateTraceTopTagConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update trace top tag config o k response has a 3xx status code
func (o *UpdateTraceTopTagConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace top tag config o k response has a 4xx status code
func (o *UpdateTraceTopTagConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trace top tag config o k response has a 5xx status code
func (o *UpdateTraceTopTagConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace top tag config o k response a status code equal to that given
func (o *UpdateTraceTopTagConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update trace top tag config o k response
func (o *UpdateTraceTopTagConfigOK) Code() int {
	return 200
}

func (o *UpdateTraceTopTagConfigOK) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-top-tag-config][%d] updateTraceTopTagConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateTraceTopTagConfigOK) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-top-tag-config][%d] updateTraceTopTagConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateTraceTopTagConfigOK) GetPayload() *models.ConfigunstableUpdateTraceTopTagConfigResponse {
	return o.Payload
}

func (o *UpdateTraceTopTagConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableUpdateTraceTopTagConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceTopTagConfigBadRequest creates a UpdateTraceTopTagConfigBadRequest with default headers values
func NewUpdateTraceTopTagConfigBadRequest() *UpdateTraceTopTagConfigBadRequest {
	return &UpdateTraceTopTagConfigBadRequest{}
}

/*
UpdateTraceTopTagConfigBadRequest describes a response with status code 400, with default header values.

Cannot update the TraceTopTagConfig because the request is invalid.
*/
type UpdateTraceTopTagConfigBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace top tag config bad request response has a 2xx status code
func (o *UpdateTraceTopTagConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace top tag config bad request response has a 3xx status code
func (o *UpdateTraceTopTagConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace top tag config bad request response has a 4xx status code
func (o *UpdateTraceTopTagConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trace top tag config bad request response has a 5xx status code
func (o *UpdateTraceTopTagConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace top tag config bad request response a status code equal to that given
func (o *UpdateTraceTopTagConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update trace top tag config bad request response
func (o *UpdateTraceTopTagConfigBadRequest) Code() int {
	return 400
}

func (o *UpdateTraceTopTagConfigBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-top-tag-config][%d] updateTraceTopTagConfigBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTraceTopTagConfigBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-top-tag-config][%d] updateTraceTopTagConfigBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTraceTopTagConfigBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceTopTagConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceTopTagConfigNotFound creates a UpdateTraceTopTagConfigNotFound with default headers values
func NewUpdateTraceTopTagConfigNotFound() *UpdateTraceTopTagConfigNotFound {
	return &UpdateTraceTopTagConfigNotFound{}
}

/*
UpdateTraceTopTagConfigNotFound describes a response with status code 404, with default header values.

Cannot update the TraceTopTagConfig because TraceTopTagConfig has not been created.
*/
type UpdateTraceTopTagConfigNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace top tag config not found response has a 2xx status code
func (o *UpdateTraceTopTagConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace top tag config not found response has a 3xx status code
func (o *UpdateTraceTopTagConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace top tag config not found response has a 4xx status code
func (o *UpdateTraceTopTagConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trace top tag config not found response has a 5xx status code
func (o *UpdateTraceTopTagConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace top tag config not found response a status code equal to that given
func (o *UpdateTraceTopTagConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update trace top tag config not found response
func (o *UpdateTraceTopTagConfigNotFound) Code() int {
	return 404
}

func (o *UpdateTraceTopTagConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-top-tag-config][%d] updateTraceTopTagConfigNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTraceTopTagConfigNotFound) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-top-tag-config][%d] updateTraceTopTagConfigNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTraceTopTagConfigNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceTopTagConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceTopTagConfigInternalServerError creates a UpdateTraceTopTagConfigInternalServerError with default headers values
func NewUpdateTraceTopTagConfigInternalServerError() *UpdateTraceTopTagConfigInternalServerError {
	return &UpdateTraceTopTagConfigInternalServerError{}
}

/*
UpdateTraceTopTagConfigInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateTraceTopTagConfigInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace top tag config internal server error response has a 2xx status code
func (o *UpdateTraceTopTagConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace top tag config internal server error response has a 3xx status code
func (o *UpdateTraceTopTagConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace top tag config internal server error response has a 4xx status code
func (o *UpdateTraceTopTagConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trace top tag config internal server error response has a 5xx status code
func (o *UpdateTraceTopTagConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update trace top tag config internal server error response a status code equal to that given
func (o *UpdateTraceTopTagConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update trace top tag config internal server error response
func (o *UpdateTraceTopTagConfigInternalServerError) Code() int {
	return 500
}

func (o *UpdateTraceTopTagConfigInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-top-tag-config][%d] updateTraceTopTagConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTraceTopTagConfigInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-top-tag-config][%d] updateTraceTopTagConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTraceTopTagConfigInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceTopTagConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceTopTagConfigDefault creates a UpdateTraceTopTagConfigDefault with default headers values
func NewUpdateTraceTopTagConfigDefault(code int) *UpdateTraceTopTagConfigDefault {
	return &UpdateTraceTopTagConfigDefault{
		_statusCode: code,
	}
}

/*
UpdateTraceTopTagConfigDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateTraceTopTagConfigDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update trace top tag config default response has a 2xx status code
func (o *UpdateTraceTopTagConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update trace top tag config default response has a 3xx status code
func (o *UpdateTraceTopTagConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update trace top tag config default response has a 4xx status code
func (o *UpdateTraceTopTagConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update trace top tag config default response has a 5xx status code
func (o *UpdateTraceTopTagConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update trace top tag config default response a status code equal to that given
func (o *UpdateTraceTopTagConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update trace top tag config default response
func (o *UpdateTraceTopTagConfigDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTraceTopTagConfigDefault) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-top-tag-config][%d] UpdateTraceTopTagConfig default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTraceTopTagConfigDefault) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-top-tag-config][%d] UpdateTraceTopTagConfig default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTraceTopTagConfigDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateTraceTopTagConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

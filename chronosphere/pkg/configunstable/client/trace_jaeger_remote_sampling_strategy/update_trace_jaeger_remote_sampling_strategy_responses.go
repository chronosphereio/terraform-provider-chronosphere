// Code generated by go-swagger; DO NOT EDIT.

package trace_jaeger_remote_sampling_strategy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

// UpdateTraceJaegerRemoteSamplingStrategyReader is a Reader for the UpdateTraceJaegerRemoteSamplingStrategy structure.
type UpdateTraceJaegerRemoteSamplingStrategyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTraceJaegerRemoteSamplingStrategyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTraceJaegerRemoteSamplingStrategyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTraceJaegerRemoteSamplingStrategyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTraceJaegerRemoteSamplingStrategyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateTraceJaegerRemoteSamplingStrategyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTraceJaegerRemoteSamplingStrategyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateTraceJaegerRemoteSamplingStrategyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTraceJaegerRemoteSamplingStrategyOK creates a UpdateTraceJaegerRemoteSamplingStrategyOK with default headers values
func NewUpdateTraceJaegerRemoteSamplingStrategyOK() *UpdateTraceJaegerRemoteSamplingStrategyOK {
	return &UpdateTraceJaegerRemoteSamplingStrategyOK{}
}

/*
UpdateTraceJaegerRemoteSamplingStrategyOK describes a response with status code 200, with default header values.

A successful response containing the updated TraceJaegerRemoteSamplingStrategy.
*/
type UpdateTraceJaegerRemoteSamplingStrategyOK struct {
	Payload *models.ConfigunstableUpdateTraceJaegerRemoteSamplingStrategyResponse
}

// IsSuccess returns true when this update trace jaeger remote sampling strategy o k response has a 2xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update trace jaeger remote sampling strategy o k response has a 3xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace jaeger remote sampling strategy o k response has a 4xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trace jaeger remote sampling strategy o k response has a 5xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace jaeger remote sampling strategy o k response a status code equal to that given
func (o *UpdateTraceJaegerRemoteSamplingStrategyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update trace jaeger remote sampling strategy o k response
func (o *UpdateTraceJaegerRemoteSamplingStrategyOK) Code() int {
	return 200
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyOK) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] updateTraceJaegerRemoteSamplingStrategyOK  %+v", 200, o.Payload)
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyOK) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] updateTraceJaegerRemoteSamplingStrategyOK  %+v", 200, o.Payload)
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyOK) GetPayload() *models.ConfigunstableUpdateTraceJaegerRemoteSamplingStrategyResponse {
	return o.Payload
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableUpdateTraceJaegerRemoteSamplingStrategyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceJaegerRemoteSamplingStrategyBadRequest creates a UpdateTraceJaegerRemoteSamplingStrategyBadRequest with default headers values
func NewUpdateTraceJaegerRemoteSamplingStrategyBadRequest() *UpdateTraceJaegerRemoteSamplingStrategyBadRequest {
	return &UpdateTraceJaegerRemoteSamplingStrategyBadRequest{}
}

/*
UpdateTraceJaegerRemoteSamplingStrategyBadRequest describes a response with status code 400, with default header values.

Cannot update the TraceJaegerRemoteSamplingStrategy because the request is invalid.
*/
type UpdateTraceJaegerRemoteSamplingStrategyBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace jaeger remote sampling strategy bad request response has a 2xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace jaeger remote sampling strategy bad request response has a 3xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace jaeger remote sampling strategy bad request response has a 4xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trace jaeger remote sampling strategy bad request response has a 5xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace jaeger remote sampling strategy bad request response a status code equal to that given
func (o *UpdateTraceJaegerRemoteSamplingStrategyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update trace jaeger remote sampling strategy bad request response
func (o *UpdateTraceJaegerRemoteSamplingStrategyBadRequest) Code() int {
	return 400
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] updateTraceJaegerRemoteSamplingStrategyBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] updateTraceJaegerRemoteSamplingStrategyBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceJaegerRemoteSamplingStrategyNotFound creates a UpdateTraceJaegerRemoteSamplingStrategyNotFound with default headers values
func NewUpdateTraceJaegerRemoteSamplingStrategyNotFound() *UpdateTraceJaegerRemoteSamplingStrategyNotFound {
	return &UpdateTraceJaegerRemoteSamplingStrategyNotFound{}
}

/*
UpdateTraceJaegerRemoteSamplingStrategyNotFound describes a response with status code 404, with default header values.

Cannot update the TraceJaegerRemoteSamplingStrategy because the slug does not exist.
*/
type UpdateTraceJaegerRemoteSamplingStrategyNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace jaeger remote sampling strategy not found response has a 2xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace jaeger remote sampling strategy not found response has a 3xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace jaeger remote sampling strategy not found response has a 4xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trace jaeger remote sampling strategy not found response has a 5xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace jaeger remote sampling strategy not found response a status code equal to that given
func (o *UpdateTraceJaegerRemoteSamplingStrategyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update trace jaeger remote sampling strategy not found response
func (o *UpdateTraceJaegerRemoteSamplingStrategyNotFound) Code() int {
	return 404
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] updateTraceJaegerRemoteSamplingStrategyNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyNotFound) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] updateTraceJaegerRemoteSamplingStrategyNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceJaegerRemoteSamplingStrategyConflict creates a UpdateTraceJaegerRemoteSamplingStrategyConflict with default headers values
func NewUpdateTraceJaegerRemoteSamplingStrategyConflict() *UpdateTraceJaegerRemoteSamplingStrategyConflict {
	return &UpdateTraceJaegerRemoteSamplingStrategyConflict{}
}

/*
UpdateTraceJaegerRemoteSamplingStrategyConflict describes a response with status code 409, with default header values.

Cannot update the TraceJaegerRemoteSamplingStrategy because there is a conflict with an existing TraceJaegerRemoteSamplingStrategy.
*/
type UpdateTraceJaegerRemoteSamplingStrategyConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace jaeger remote sampling strategy conflict response has a 2xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace jaeger remote sampling strategy conflict response has a 3xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace jaeger remote sampling strategy conflict response has a 4xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trace jaeger remote sampling strategy conflict response has a 5xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update trace jaeger remote sampling strategy conflict response a status code equal to that given
func (o *UpdateTraceJaegerRemoteSamplingStrategyConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update trace jaeger remote sampling strategy conflict response
func (o *UpdateTraceJaegerRemoteSamplingStrategyConflict) Code() int {
	return 409
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyConflict) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] updateTraceJaegerRemoteSamplingStrategyConflict  %+v", 409, o.Payload)
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyConflict) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] updateTraceJaegerRemoteSamplingStrategyConflict  %+v", 409, o.Payload)
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceJaegerRemoteSamplingStrategyInternalServerError creates a UpdateTraceJaegerRemoteSamplingStrategyInternalServerError with default headers values
func NewUpdateTraceJaegerRemoteSamplingStrategyInternalServerError() *UpdateTraceJaegerRemoteSamplingStrategyInternalServerError {
	return &UpdateTraceJaegerRemoteSamplingStrategyInternalServerError{}
}

/*
UpdateTraceJaegerRemoteSamplingStrategyInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateTraceJaegerRemoteSamplingStrategyInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update trace jaeger remote sampling strategy internal server error response has a 2xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trace jaeger remote sampling strategy internal server error response has a 3xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trace jaeger remote sampling strategy internal server error response has a 4xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trace jaeger remote sampling strategy internal server error response has a 5xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update trace jaeger remote sampling strategy internal server error response a status code equal to that given
func (o *UpdateTraceJaegerRemoteSamplingStrategyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update trace jaeger remote sampling strategy internal server error response
func (o *UpdateTraceJaegerRemoteSamplingStrategyInternalServerError) Code() int {
	return 500
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] updateTraceJaegerRemoteSamplingStrategyInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] updateTraceJaegerRemoteSamplingStrategyInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTraceJaegerRemoteSamplingStrategyDefault creates a UpdateTraceJaegerRemoteSamplingStrategyDefault with default headers values
func NewUpdateTraceJaegerRemoteSamplingStrategyDefault(code int) *UpdateTraceJaegerRemoteSamplingStrategyDefault {
	return &UpdateTraceJaegerRemoteSamplingStrategyDefault{
		_statusCode: code,
	}
}

/*
UpdateTraceJaegerRemoteSamplingStrategyDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateTraceJaegerRemoteSamplingStrategyDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update trace jaeger remote sampling strategy default response has a 2xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update trace jaeger remote sampling strategy default response has a 3xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update trace jaeger remote sampling strategy default response has a 4xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update trace jaeger remote sampling strategy default response has a 5xx status code
func (o *UpdateTraceJaegerRemoteSamplingStrategyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update trace jaeger remote sampling strategy default response a status code equal to that given
func (o *UpdateTraceJaegerRemoteSamplingStrategyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update trace jaeger remote sampling strategy default response
func (o *UpdateTraceJaegerRemoteSamplingStrategyDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyDefault) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] UpdateTraceJaegerRemoteSamplingStrategy default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyDefault) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/trace-jaeger-remote-sampling-strategies/{slug}][%d] UpdateTraceJaegerRemoteSamplingStrategy default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateTraceJaegerRemoteSamplingStrategyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

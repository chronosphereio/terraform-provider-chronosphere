// Code generated by go-swagger; DO NOT EDIT.

package log_scale_alert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// ReadLogScaleAlertReader is a Reader for the ReadLogScaleAlert structure.
type ReadLogScaleAlertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadLogScaleAlertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadLogScaleAlertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadLogScaleAlertNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadLogScaleAlertInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadLogScaleAlertDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadLogScaleAlertOK creates a ReadLogScaleAlertOK with default headers values
func NewReadLogScaleAlertOK() *ReadLogScaleAlertOK {
	return &ReadLogScaleAlertOK{}
}

/*
ReadLogScaleAlertOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadLogScaleAlertOK struct {
	Payload *models.Configv1ReadLogScaleAlertResponse
}

// IsSuccess returns true when this read log scale alert o k response has a 2xx status code
func (o *ReadLogScaleAlertOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read log scale alert o k response has a 3xx status code
func (o *ReadLogScaleAlertOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read log scale alert o k response has a 4xx status code
func (o *ReadLogScaleAlertOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read log scale alert o k response has a 5xx status code
func (o *ReadLogScaleAlertOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read log scale alert o k response a status code equal to that given
func (o *ReadLogScaleAlertOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read log scale alert o k response
func (o *ReadLogScaleAlertOK) Code() int {
	return 200
}

func (o *ReadLogScaleAlertOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/log-scale-alerts/{slug}][%d] readLogScaleAlertOK  %+v", 200, o.Payload)
}

func (o *ReadLogScaleAlertOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/log-scale-alerts/{slug}][%d] readLogScaleAlertOK  %+v", 200, o.Payload)
}

func (o *ReadLogScaleAlertOK) GetPayload() *models.Configv1ReadLogScaleAlertResponse {
	return o.Payload
}

func (o *ReadLogScaleAlertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ReadLogScaleAlertResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadLogScaleAlertNotFound creates a ReadLogScaleAlertNotFound with default headers values
func NewReadLogScaleAlertNotFound() *ReadLogScaleAlertNotFound {
	return &ReadLogScaleAlertNotFound{}
}

/*
ReadLogScaleAlertNotFound describes a response with status code 404, with default header values.

Cannot read the LogScaleAlert because the slug does not exist.
*/
type ReadLogScaleAlertNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read log scale alert not found response has a 2xx status code
func (o *ReadLogScaleAlertNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read log scale alert not found response has a 3xx status code
func (o *ReadLogScaleAlertNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read log scale alert not found response has a 4xx status code
func (o *ReadLogScaleAlertNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read log scale alert not found response has a 5xx status code
func (o *ReadLogScaleAlertNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read log scale alert not found response a status code equal to that given
func (o *ReadLogScaleAlertNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read log scale alert not found response
func (o *ReadLogScaleAlertNotFound) Code() int {
	return 404
}

func (o *ReadLogScaleAlertNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/log-scale-alerts/{slug}][%d] readLogScaleAlertNotFound  %+v", 404, o.Payload)
}

func (o *ReadLogScaleAlertNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/config/log-scale-alerts/{slug}][%d] readLogScaleAlertNotFound  %+v", 404, o.Payload)
}

func (o *ReadLogScaleAlertNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadLogScaleAlertNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadLogScaleAlertInternalServerError creates a ReadLogScaleAlertInternalServerError with default headers values
func NewReadLogScaleAlertInternalServerError() *ReadLogScaleAlertInternalServerError {
	return &ReadLogScaleAlertInternalServerError{}
}

/*
ReadLogScaleAlertInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadLogScaleAlertInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read log scale alert internal server error response has a 2xx status code
func (o *ReadLogScaleAlertInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read log scale alert internal server error response has a 3xx status code
func (o *ReadLogScaleAlertInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read log scale alert internal server error response has a 4xx status code
func (o *ReadLogScaleAlertInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read log scale alert internal server error response has a 5xx status code
func (o *ReadLogScaleAlertInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read log scale alert internal server error response a status code equal to that given
func (o *ReadLogScaleAlertInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read log scale alert internal server error response
func (o *ReadLogScaleAlertInternalServerError) Code() int {
	return 500
}

func (o *ReadLogScaleAlertInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/log-scale-alerts/{slug}][%d] readLogScaleAlertInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadLogScaleAlertInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/log-scale-alerts/{slug}][%d] readLogScaleAlertInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadLogScaleAlertInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadLogScaleAlertInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadLogScaleAlertDefault creates a ReadLogScaleAlertDefault with default headers values
func NewReadLogScaleAlertDefault(code int) *ReadLogScaleAlertDefault {
	return &ReadLogScaleAlertDefault{
		_statusCode: code,
	}
}

/*
ReadLogScaleAlertDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadLogScaleAlertDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read log scale alert default response has a 2xx status code
func (o *ReadLogScaleAlertDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read log scale alert default response has a 3xx status code
func (o *ReadLogScaleAlertDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read log scale alert default response has a 4xx status code
func (o *ReadLogScaleAlertDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read log scale alert default response has a 5xx status code
func (o *ReadLogScaleAlertDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read log scale alert default response a status code equal to that given
func (o *ReadLogScaleAlertDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read log scale alert default response
func (o *ReadLogScaleAlertDefault) Code() int {
	return o._statusCode
}

func (o *ReadLogScaleAlertDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/log-scale-alerts/{slug}][%d] ReadLogScaleAlert default  %+v", o._statusCode, o.Payload)
}

func (o *ReadLogScaleAlertDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/log-scale-alerts/{slug}][%d] ReadLogScaleAlert default  %+v", o._statusCode, o.Payload)
}

func (o *ReadLogScaleAlertDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadLogScaleAlertDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

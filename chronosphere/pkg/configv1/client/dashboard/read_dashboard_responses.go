// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// ReadDashboardReader is a Reader for the ReadDashboard structure.
type ReadDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadDashboardOK creates a ReadDashboardOK with default headers values
func NewReadDashboardOK() *ReadDashboardOK {
	return &ReadDashboardOK{}
}

/*
ReadDashboardOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadDashboardOK struct {
	Payload *models.Configv1ReadDashboardResponse
}

// IsSuccess returns true when this read dashboard o k response has a 2xx status code
func (o *ReadDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read dashboard o k response has a 3xx status code
func (o *ReadDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read dashboard o k response has a 4xx status code
func (o *ReadDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read dashboard o k response has a 5xx status code
func (o *ReadDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read dashboard o k response a status code equal to that given
func (o *ReadDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read dashboard o k response
func (o *ReadDashboardOK) Code() int {
	return 200
}

func (o *ReadDashboardOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/dashboards/{slug}][%d] readDashboardOK  %+v", 200, o.Payload)
}

func (o *ReadDashboardOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/dashboards/{slug}][%d] readDashboardOK  %+v", 200, o.Payload)
}

func (o *ReadDashboardOK) GetPayload() *models.Configv1ReadDashboardResponse {
	return o.Payload
}

func (o *ReadDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ReadDashboardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDashboardNotFound creates a ReadDashboardNotFound with default headers values
func NewReadDashboardNotFound() *ReadDashboardNotFound {
	return &ReadDashboardNotFound{}
}

/*
ReadDashboardNotFound describes a response with status code 404, with default header values.

Cannot read the Dashboard because the slug does not exist.
*/
type ReadDashboardNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read dashboard not found response has a 2xx status code
func (o *ReadDashboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read dashboard not found response has a 3xx status code
func (o *ReadDashboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read dashboard not found response has a 4xx status code
func (o *ReadDashboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read dashboard not found response has a 5xx status code
func (o *ReadDashboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read dashboard not found response a status code equal to that given
func (o *ReadDashboardNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read dashboard not found response
func (o *ReadDashboardNotFound) Code() int {
	return 404
}

func (o *ReadDashboardNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/dashboards/{slug}][%d] readDashboardNotFound  %+v", 404, o.Payload)
}

func (o *ReadDashboardNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/config/dashboards/{slug}][%d] readDashboardNotFound  %+v", 404, o.Payload)
}

func (o *ReadDashboardNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDashboardInternalServerError creates a ReadDashboardInternalServerError with default headers values
func NewReadDashboardInternalServerError() *ReadDashboardInternalServerError {
	return &ReadDashboardInternalServerError{}
}

/*
ReadDashboardInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadDashboardInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read dashboard internal server error response has a 2xx status code
func (o *ReadDashboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read dashboard internal server error response has a 3xx status code
func (o *ReadDashboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read dashboard internal server error response has a 4xx status code
func (o *ReadDashboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read dashboard internal server error response has a 5xx status code
func (o *ReadDashboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read dashboard internal server error response a status code equal to that given
func (o *ReadDashboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read dashboard internal server error response
func (o *ReadDashboardInternalServerError) Code() int {
	return 500
}

func (o *ReadDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/dashboards/{slug}][%d] readDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadDashboardInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/dashboards/{slug}][%d] readDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadDashboardInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDashboardDefault creates a ReadDashboardDefault with default headers values
func NewReadDashboardDefault(code int) *ReadDashboardDefault {
	return &ReadDashboardDefault{
		_statusCode: code,
	}
}

/*
ReadDashboardDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadDashboardDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read dashboard default response has a 2xx status code
func (o *ReadDashboardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read dashboard default response has a 3xx status code
func (o *ReadDashboardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read dashboard default response has a 4xx status code
func (o *ReadDashboardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read dashboard default response has a 5xx status code
func (o *ReadDashboardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read dashboard default response a status code equal to that given
func (o *ReadDashboardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read dashboard default response
func (o *ReadDashboardDefault) Code() int {
	return o._statusCode
}

func (o *ReadDashboardDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/dashboards/{slug}][%d] ReadDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *ReadDashboardDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/dashboards/{slug}][%d] ReadDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *ReadDashboardDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

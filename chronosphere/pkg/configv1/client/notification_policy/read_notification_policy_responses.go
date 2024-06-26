// Code generated by go-swagger; DO NOT EDIT.

package notification_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// ReadNotificationPolicyReader is a Reader for the ReadNotificationPolicy structure.
type ReadNotificationPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadNotificationPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadNotificationPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReadNotificationPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadNotificationPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReadNotificationPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReadNotificationPolicyOK creates a ReadNotificationPolicyOK with default headers values
func NewReadNotificationPolicyOK() *ReadNotificationPolicyOK {
	return &ReadNotificationPolicyOK{}
}

/*
ReadNotificationPolicyOK describes a response with status code 200, with default header values.

A successful response.
*/
type ReadNotificationPolicyOK struct {
	Payload *models.Configv1ReadNotificationPolicyResponse
}

// IsSuccess returns true when this read notification policy o k response has a 2xx status code
func (o *ReadNotificationPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this read notification policy o k response has a 3xx status code
func (o *ReadNotificationPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read notification policy o k response has a 4xx status code
func (o *ReadNotificationPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this read notification policy o k response has a 5xx status code
func (o *ReadNotificationPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this read notification policy o k response a status code equal to that given
func (o *ReadNotificationPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the read notification policy o k response
func (o *ReadNotificationPolicyOK) Code() int {
	return 200
}

func (o *ReadNotificationPolicyOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies/{slug}][%d] readNotificationPolicyOK  %+v", 200, o.Payload)
}

func (o *ReadNotificationPolicyOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies/{slug}][%d] readNotificationPolicyOK  %+v", 200, o.Payload)
}

func (o *ReadNotificationPolicyOK) GetPayload() *models.Configv1ReadNotificationPolicyResponse {
	return o.Payload
}

func (o *ReadNotificationPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ReadNotificationPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadNotificationPolicyNotFound creates a ReadNotificationPolicyNotFound with default headers values
func NewReadNotificationPolicyNotFound() *ReadNotificationPolicyNotFound {
	return &ReadNotificationPolicyNotFound{}
}

/*
ReadNotificationPolicyNotFound describes a response with status code 404, with default header values.

Cannot read the NotificationPolicy because the slug does not exist.
*/
type ReadNotificationPolicyNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read notification policy not found response has a 2xx status code
func (o *ReadNotificationPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read notification policy not found response has a 3xx status code
func (o *ReadNotificationPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read notification policy not found response has a 4xx status code
func (o *ReadNotificationPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this read notification policy not found response has a 5xx status code
func (o *ReadNotificationPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this read notification policy not found response a status code equal to that given
func (o *ReadNotificationPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the read notification policy not found response
func (o *ReadNotificationPolicyNotFound) Code() int {
	return 404
}

func (o *ReadNotificationPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies/{slug}][%d] readNotificationPolicyNotFound  %+v", 404, o.Payload)
}

func (o *ReadNotificationPolicyNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies/{slug}][%d] readNotificationPolicyNotFound  %+v", 404, o.Payload)
}

func (o *ReadNotificationPolicyNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadNotificationPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadNotificationPolicyInternalServerError creates a ReadNotificationPolicyInternalServerError with default headers values
func NewReadNotificationPolicyInternalServerError() *ReadNotificationPolicyInternalServerError {
	return &ReadNotificationPolicyInternalServerError{}
}

/*
ReadNotificationPolicyInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ReadNotificationPolicyInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this read notification policy internal server error response has a 2xx status code
func (o *ReadNotificationPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this read notification policy internal server error response has a 3xx status code
func (o *ReadNotificationPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this read notification policy internal server error response has a 4xx status code
func (o *ReadNotificationPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this read notification policy internal server error response has a 5xx status code
func (o *ReadNotificationPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this read notification policy internal server error response a status code equal to that given
func (o *ReadNotificationPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the read notification policy internal server error response
func (o *ReadNotificationPolicyInternalServerError) Code() int {
	return 500
}

func (o *ReadNotificationPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies/{slug}][%d] readNotificationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadNotificationPolicyInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies/{slug}][%d] readNotificationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadNotificationPolicyInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ReadNotificationPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadNotificationPolicyDefault creates a ReadNotificationPolicyDefault with default headers values
func NewReadNotificationPolicyDefault(code int) *ReadNotificationPolicyDefault {
	return &ReadNotificationPolicyDefault{
		_statusCode: code,
	}
}

/*
ReadNotificationPolicyDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ReadNotificationPolicyDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this read notification policy default response has a 2xx status code
func (o *ReadNotificationPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this read notification policy default response has a 3xx status code
func (o *ReadNotificationPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this read notification policy default response has a 4xx status code
func (o *ReadNotificationPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this read notification policy default response has a 5xx status code
func (o *ReadNotificationPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this read notification policy default response a status code equal to that given
func (o *ReadNotificationPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the read notification policy default response
func (o *ReadNotificationPolicyDefault) Code() int {
	return o._statusCode
}

func (o *ReadNotificationPolicyDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies/{slug}][%d] ReadNotificationPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *ReadNotificationPolicyDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies/{slug}][%d] ReadNotificationPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *ReadNotificationPolicyDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ReadNotificationPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

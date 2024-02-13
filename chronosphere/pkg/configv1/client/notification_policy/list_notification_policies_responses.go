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

// ListNotificationPoliciesReader is a Reader for the ListNotificationPolicies structure.
type ListNotificationPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNotificationPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNotificationPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListNotificationPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListNotificationPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNotificationPoliciesOK creates a ListNotificationPoliciesOK with default headers values
func NewListNotificationPoliciesOK() *ListNotificationPoliciesOK {
	return &ListNotificationPoliciesOK{}
}

/*
ListNotificationPoliciesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListNotificationPoliciesOK struct {
	Payload *models.Configv1ListNotificationPoliciesResponse
}

// IsSuccess returns true when this list notification policies o k response has a 2xx status code
func (o *ListNotificationPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list notification policies o k response has a 3xx status code
func (o *ListNotificationPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list notification policies o k response has a 4xx status code
func (o *ListNotificationPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list notification policies o k response has a 5xx status code
func (o *ListNotificationPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list notification policies o k response a status code equal to that given
func (o *ListNotificationPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list notification policies o k response
func (o *ListNotificationPoliciesOK) Code() int {
	return 200
}

func (o *ListNotificationPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies][%d] listNotificationPoliciesOK  %+v", 200, o.Payload)
}

func (o *ListNotificationPoliciesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies][%d] listNotificationPoliciesOK  %+v", 200, o.Payload)
}

func (o *ListNotificationPoliciesOK) GetPayload() *models.Configv1ListNotificationPoliciesResponse {
	return o.Payload
}

func (o *ListNotificationPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1ListNotificationPoliciesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNotificationPoliciesInternalServerError creates a ListNotificationPoliciesInternalServerError with default headers values
func NewListNotificationPoliciesInternalServerError() *ListNotificationPoliciesInternalServerError {
	return &ListNotificationPoliciesInternalServerError{}
}

/*
ListNotificationPoliciesInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type ListNotificationPoliciesInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this list notification policies internal server error response has a 2xx status code
func (o *ListNotificationPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list notification policies internal server error response has a 3xx status code
func (o *ListNotificationPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list notification policies internal server error response has a 4xx status code
func (o *ListNotificationPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list notification policies internal server error response has a 5xx status code
func (o *ListNotificationPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list notification policies internal server error response a status code equal to that given
func (o *ListNotificationPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list notification policies internal server error response
func (o *ListNotificationPoliciesInternalServerError) Code() int {
	return 500
}

func (o *ListNotificationPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies][%d] listNotificationPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListNotificationPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies][%d] listNotificationPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListNotificationPoliciesInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *ListNotificationPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNotificationPoliciesDefault creates a ListNotificationPoliciesDefault with default headers values
func NewListNotificationPoliciesDefault(code int) *ListNotificationPoliciesDefault {
	return &ListNotificationPoliciesDefault{
		_statusCode: code,
	}
}

/*
ListNotificationPoliciesDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type ListNotificationPoliciesDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this list notification policies default response has a 2xx status code
func (o *ListNotificationPoliciesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list notification policies default response has a 3xx status code
func (o *ListNotificationPoliciesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list notification policies default response has a 4xx status code
func (o *ListNotificationPoliciesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list notification policies default response has a 5xx status code
func (o *ListNotificationPoliciesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list notification policies default response a status code equal to that given
func (o *ListNotificationPoliciesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list notification policies default response
func (o *ListNotificationPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *ListNotificationPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies][%d] ListNotificationPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *ListNotificationPoliciesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/config/notification-policies][%d] ListNotificationPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *ListNotificationPoliciesDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *ListNotificationPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

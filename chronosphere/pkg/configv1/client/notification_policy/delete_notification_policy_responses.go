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

// DeleteNotificationPolicyReader is a Reader for the DeleteNotificationPolicy structure.
type DeleteNotificationPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNotificationPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNotificationPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteNotificationPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteNotificationPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteNotificationPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNotificationPolicyOK creates a DeleteNotificationPolicyOK with default headers values
func NewDeleteNotificationPolicyOK() *DeleteNotificationPolicyOK {
	return &DeleteNotificationPolicyOK{}
}

/*
DeleteNotificationPolicyOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteNotificationPolicyOK struct {
	Payload models.Configv1DeleteNotificationPolicyResponse
}

// IsSuccess returns true when this delete notification policy o k response has a 2xx status code
func (o *DeleteNotificationPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete notification policy o k response has a 3xx status code
func (o *DeleteNotificationPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete notification policy o k response has a 4xx status code
func (o *DeleteNotificationPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete notification policy o k response has a 5xx status code
func (o *DeleteNotificationPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete notification policy o k response a status code equal to that given
func (o *DeleteNotificationPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete notification policy o k response
func (o *DeleteNotificationPolicyOK) Code() int {
	return 200
}

func (o *DeleteNotificationPolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notification-policies/{slug}][%d] deleteNotificationPolicyOK  %+v", 200, o.Payload)
}

func (o *DeleteNotificationPolicyOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notification-policies/{slug}][%d] deleteNotificationPolicyOK  %+v", 200, o.Payload)
}

func (o *DeleteNotificationPolicyOK) GetPayload() models.Configv1DeleteNotificationPolicyResponse {
	return o.Payload
}

func (o *DeleteNotificationPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNotificationPolicyNotFound creates a DeleteNotificationPolicyNotFound with default headers values
func NewDeleteNotificationPolicyNotFound() *DeleteNotificationPolicyNotFound {
	return &DeleteNotificationPolicyNotFound{}
}

/*
DeleteNotificationPolicyNotFound describes a response with status code 404, with default header values.

Cannot delete the NotificationPolicy because the slug does not exist.
*/
type DeleteNotificationPolicyNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete notification policy not found response has a 2xx status code
func (o *DeleteNotificationPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete notification policy not found response has a 3xx status code
func (o *DeleteNotificationPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete notification policy not found response has a 4xx status code
func (o *DeleteNotificationPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete notification policy not found response has a 5xx status code
func (o *DeleteNotificationPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete notification policy not found response a status code equal to that given
func (o *DeleteNotificationPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete notification policy not found response
func (o *DeleteNotificationPolicyNotFound) Code() int {
	return 404
}

func (o *DeleteNotificationPolicyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notification-policies/{slug}][%d] deleteNotificationPolicyNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNotificationPolicyNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notification-policies/{slug}][%d] deleteNotificationPolicyNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNotificationPolicyNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteNotificationPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNotificationPolicyInternalServerError creates a DeleteNotificationPolicyInternalServerError with default headers values
func NewDeleteNotificationPolicyInternalServerError() *DeleteNotificationPolicyInternalServerError {
	return &DeleteNotificationPolicyInternalServerError{}
}

/*
DeleteNotificationPolicyInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteNotificationPolicyInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete notification policy internal server error response has a 2xx status code
func (o *DeleteNotificationPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete notification policy internal server error response has a 3xx status code
func (o *DeleteNotificationPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete notification policy internal server error response has a 4xx status code
func (o *DeleteNotificationPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete notification policy internal server error response has a 5xx status code
func (o *DeleteNotificationPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete notification policy internal server error response a status code equal to that given
func (o *DeleteNotificationPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete notification policy internal server error response
func (o *DeleteNotificationPolicyInternalServerError) Code() int {
	return 500
}

func (o *DeleteNotificationPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notification-policies/{slug}][%d] deleteNotificationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNotificationPolicyInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notification-policies/{slug}][%d] deleteNotificationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNotificationPolicyInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteNotificationPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNotificationPolicyDefault creates a DeleteNotificationPolicyDefault with default headers values
func NewDeleteNotificationPolicyDefault(code int) *DeleteNotificationPolicyDefault {
	return &DeleteNotificationPolicyDefault{
		_statusCode: code,
	}
}

/*
DeleteNotificationPolicyDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteNotificationPolicyDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete notification policy default response has a 2xx status code
func (o *DeleteNotificationPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete notification policy default response has a 3xx status code
func (o *DeleteNotificationPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete notification policy default response has a 4xx status code
func (o *DeleteNotificationPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete notification policy default response has a 5xx status code
func (o *DeleteNotificationPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete notification policy default response a status code equal to that given
func (o *DeleteNotificationPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete notification policy default response
func (o *DeleteNotificationPolicyDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNotificationPolicyDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notification-policies/{slug}][%d] DeleteNotificationPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNotificationPolicyDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notification-policies/{slug}][%d] DeleteNotificationPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNotificationPolicyDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteNotificationPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

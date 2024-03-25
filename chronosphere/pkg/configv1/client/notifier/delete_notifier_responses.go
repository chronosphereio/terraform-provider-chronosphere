// Code generated by go-swagger; DO NOT EDIT.

package notifier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// DeleteNotifierReader is a Reader for the DeleteNotifier structure.
type DeleteNotifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNotifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNotifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNotifierBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteNotifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteNotifierInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteNotifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNotifierOK creates a DeleteNotifierOK with default headers values
func NewDeleteNotifierOK() *DeleteNotifierOK {
	return &DeleteNotifierOK{}
}

/*
DeleteNotifierOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteNotifierOK struct {
	Payload models.Configv1DeleteNotifierResponse
}

// IsSuccess returns true when this delete notifier o k response has a 2xx status code
func (o *DeleteNotifierOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete notifier o k response has a 3xx status code
func (o *DeleteNotifierOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete notifier o k response has a 4xx status code
func (o *DeleteNotifierOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete notifier o k response has a 5xx status code
func (o *DeleteNotifierOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete notifier o k response a status code equal to that given
func (o *DeleteNotifierOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete notifier o k response
func (o *DeleteNotifierOK) Code() int {
	return 200
}

func (o *DeleteNotifierOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notifiers/{slug}][%d] deleteNotifierOK  %+v", 200, o.Payload)
}

func (o *DeleteNotifierOK) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notifiers/{slug}][%d] deleteNotifierOK  %+v", 200, o.Payload)
}

func (o *DeleteNotifierOK) GetPayload() models.Configv1DeleteNotifierResponse {
	return o.Payload
}

func (o *DeleteNotifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNotifierBadRequest creates a DeleteNotifierBadRequest with default headers values
func NewDeleteNotifierBadRequest() *DeleteNotifierBadRequest {
	return &DeleteNotifierBadRequest{}
}

/*
DeleteNotifierBadRequest describes a response with status code 400, with default header values.

Cannot delete the Notifier because it is in use.
*/
type DeleteNotifierBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete notifier bad request response has a 2xx status code
func (o *DeleteNotifierBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete notifier bad request response has a 3xx status code
func (o *DeleteNotifierBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete notifier bad request response has a 4xx status code
func (o *DeleteNotifierBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete notifier bad request response has a 5xx status code
func (o *DeleteNotifierBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete notifier bad request response a status code equal to that given
func (o *DeleteNotifierBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete notifier bad request response
func (o *DeleteNotifierBadRequest) Code() int {
	return 400
}

func (o *DeleteNotifierBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notifiers/{slug}][%d] deleteNotifierBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteNotifierBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notifiers/{slug}][%d] deleteNotifierBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteNotifierBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteNotifierBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNotifierNotFound creates a DeleteNotifierNotFound with default headers values
func NewDeleteNotifierNotFound() *DeleteNotifierNotFound {
	return &DeleteNotifierNotFound{}
}

/*
DeleteNotifierNotFound describes a response with status code 404, with default header values.

Cannot delete the Notifier because the slug does not exist.
*/
type DeleteNotifierNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete notifier not found response has a 2xx status code
func (o *DeleteNotifierNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete notifier not found response has a 3xx status code
func (o *DeleteNotifierNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete notifier not found response has a 4xx status code
func (o *DeleteNotifierNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete notifier not found response has a 5xx status code
func (o *DeleteNotifierNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete notifier not found response a status code equal to that given
func (o *DeleteNotifierNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete notifier not found response
func (o *DeleteNotifierNotFound) Code() int {
	return 404
}

func (o *DeleteNotifierNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notifiers/{slug}][%d] deleteNotifierNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNotifierNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notifiers/{slug}][%d] deleteNotifierNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNotifierNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteNotifierNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNotifierInternalServerError creates a DeleteNotifierInternalServerError with default headers values
func NewDeleteNotifierInternalServerError() *DeleteNotifierInternalServerError {
	return &DeleteNotifierInternalServerError{}
}

/*
DeleteNotifierInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteNotifierInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete notifier internal server error response has a 2xx status code
func (o *DeleteNotifierInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete notifier internal server error response has a 3xx status code
func (o *DeleteNotifierInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete notifier internal server error response has a 4xx status code
func (o *DeleteNotifierInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete notifier internal server error response has a 5xx status code
func (o *DeleteNotifierInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete notifier internal server error response a status code equal to that given
func (o *DeleteNotifierInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete notifier internal server error response
func (o *DeleteNotifierInternalServerError) Code() int {
	return 500
}

func (o *DeleteNotifierInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notifiers/{slug}][%d] deleteNotifierInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNotifierInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notifiers/{slug}][%d] deleteNotifierInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteNotifierInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteNotifierInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNotifierDefault creates a DeleteNotifierDefault with default headers values
func NewDeleteNotifierDefault(code int) *DeleteNotifierDefault {
	return &DeleteNotifierDefault{
		_statusCode: code,
	}
}

/*
DeleteNotifierDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteNotifierDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete notifier default response has a 2xx status code
func (o *DeleteNotifierDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete notifier default response has a 3xx status code
func (o *DeleteNotifierDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete notifier default response has a 4xx status code
func (o *DeleteNotifierDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete notifier default response has a 5xx status code
func (o *DeleteNotifierDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete notifier default response a status code equal to that given
func (o *DeleteNotifierDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete notifier default response
func (o *DeleteNotifierDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNotifierDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notifiers/{slug}][%d] DeleteNotifier default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNotifierDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v1/config/notifiers/{slug}][%d] DeleteNotifier default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNotifierDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteNotifierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

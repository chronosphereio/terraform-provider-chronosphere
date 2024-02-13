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

// CreateNotificationPolicyReader is a Reader for the CreateNotificationPolicy structure.
type CreateNotificationPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNotificationPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNotificationPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateNotificationPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateNotificationPolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateNotificationPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateNotificationPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateNotificationPolicyOK creates a CreateNotificationPolicyOK with default headers values
func NewCreateNotificationPolicyOK() *CreateNotificationPolicyOK {
	return &CreateNotificationPolicyOK{}
}

/*
CreateNotificationPolicyOK describes a response with status code 200, with default header values.

A successful response containing the created NotificationPolicy.
*/
type CreateNotificationPolicyOK struct {
	Payload *models.Configv1CreateNotificationPolicyResponse
}

// IsSuccess returns true when this create notification policy o k response has a 2xx status code
func (o *CreateNotificationPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create notification policy o k response has a 3xx status code
func (o *CreateNotificationPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create notification policy o k response has a 4xx status code
func (o *CreateNotificationPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create notification policy o k response has a 5xx status code
func (o *CreateNotificationPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create notification policy o k response a status code equal to that given
func (o *CreateNotificationPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create notification policy o k response
func (o *CreateNotificationPolicyOK) Code() int {
	return 200
}

func (o *CreateNotificationPolicyOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/notification-policies][%d] createNotificationPolicyOK  %+v", 200, o.Payload)
}

func (o *CreateNotificationPolicyOK) String() string {
	return fmt.Sprintf("[POST /api/v1/config/notification-policies][%d] createNotificationPolicyOK  %+v", 200, o.Payload)
}

func (o *CreateNotificationPolicyOK) GetPayload() *models.Configv1CreateNotificationPolicyResponse {
	return o.Payload
}

func (o *CreateNotificationPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1CreateNotificationPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNotificationPolicyBadRequest creates a CreateNotificationPolicyBadRequest with default headers values
func NewCreateNotificationPolicyBadRequest() *CreateNotificationPolicyBadRequest {
	return &CreateNotificationPolicyBadRequest{}
}

/*
CreateNotificationPolicyBadRequest describes a response with status code 400, with default header values.

Cannot create the NotificationPolicy because the request is invalid.
*/
type CreateNotificationPolicyBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create notification policy bad request response has a 2xx status code
func (o *CreateNotificationPolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create notification policy bad request response has a 3xx status code
func (o *CreateNotificationPolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create notification policy bad request response has a 4xx status code
func (o *CreateNotificationPolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create notification policy bad request response has a 5xx status code
func (o *CreateNotificationPolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create notification policy bad request response a status code equal to that given
func (o *CreateNotificationPolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create notification policy bad request response
func (o *CreateNotificationPolicyBadRequest) Code() int {
	return 400
}

func (o *CreateNotificationPolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/notification-policies][%d] createNotificationPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNotificationPolicyBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/config/notification-policies][%d] createNotificationPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNotificationPolicyBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateNotificationPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNotificationPolicyConflict creates a CreateNotificationPolicyConflict with default headers values
func NewCreateNotificationPolicyConflict() *CreateNotificationPolicyConflict {
	return &CreateNotificationPolicyConflict{}
}

/*
CreateNotificationPolicyConflict describes a response with status code 409, with default header values.

Cannot create the NotificationPolicy because there is a conflict with an existing NotificationPolicy.
*/
type CreateNotificationPolicyConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create notification policy conflict response has a 2xx status code
func (o *CreateNotificationPolicyConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create notification policy conflict response has a 3xx status code
func (o *CreateNotificationPolicyConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create notification policy conflict response has a 4xx status code
func (o *CreateNotificationPolicyConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create notification policy conflict response has a 5xx status code
func (o *CreateNotificationPolicyConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create notification policy conflict response a status code equal to that given
func (o *CreateNotificationPolicyConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create notification policy conflict response
func (o *CreateNotificationPolicyConflict) Code() int {
	return 409
}

func (o *CreateNotificationPolicyConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/notification-policies][%d] createNotificationPolicyConflict  %+v", 409, o.Payload)
}

func (o *CreateNotificationPolicyConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/config/notification-policies][%d] createNotificationPolicyConflict  %+v", 409, o.Payload)
}

func (o *CreateNotificationPolicyConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateNotificationPolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNotificationPolicyInternalServerError creates a CreateNotificationPolicyInternalServerError with default headers values
func NewCreateNotificationPolicyInternalServerError() *CreateNotificationPolicyInternalServerError {
	return &CreateNotificationPolicyInternalServerError{}
}

/*
CreateNotificationPolicyInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateNotificationPolicyInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create notification policy internal server error response has a 2xx status code
func (o *CreateNotificationPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create notification policy internal server error response has a 3xx status code
func (o *CreateNotificationPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create notification policy internal server error response has a 4xx status code
func (o *CreateNotificationPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create notification policy internal server error response has a 5xx status code
func (o *CreateNotificationPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create notification policy internal server error response a status code equal to that given
func (o *CreateNotificationPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create notification policy internal server error response
func (o *CreateNotificationPolicyInternalServerError) Code() int {
	return 500
}

func (o *CreateNotificationPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/notification-policies][%d] createNotificationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateNotificationPolicyInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/config/notification-policies][%d] createNotificationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateNotificationPolicyInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateNotificationPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNotificationPolicyDefault creates a CreateNotificationPolicyDefault with default headers values
func NewCreateNotificationPolicyDefault(code int) *CreateNotificationPolicyDefault {
	return &CreateNotificationPolicyDefault{
		_statusCode: code,
	}
}

/*
CreateNotificationPolicyDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateNotificationPolicyDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create notification policy default response has a 2xx status code
func (o *CreateNotificationPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create notification policy default response has a 3xx status code
func (o *CreateNotificationPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create notification policy default response has a 4xx status code
func (o *CreateNotificationPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create notification policy default response has a 5xx status code
func (o *CreateNotificationPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create notification policy default response a status code equal to that given
func (o *CreateNotificationPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create notification policy default response
func (o *CreateNotificationPolicyDefault) Code() int {
	return o._statusCode
}

func (o *CreateNotificationPolicyDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/notification-policies][%d] CreateNotificationPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *CreateNotificationPolicyDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/config/notification-policies][%d] CreateNotificationPolicy default  %+v", o._statusCode, o.Payload)
}

func (o *CreateNotificationPolicyDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateNotificationPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

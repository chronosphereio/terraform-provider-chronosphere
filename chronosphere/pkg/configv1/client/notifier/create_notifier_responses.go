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

// CreateNotifierReader is a Reader for the CreateNotifier structure.
type CreateNotifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNotifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNotifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateNotifierBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateNotifierConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateNotifierInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateNotifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateNotifierOK creates a CreateNotifierOK with default headers values
func NewCreateNotifierOK() *CreateNotifierOK {
	return &CreateNotifierOK{}
}

/*
CreateNotifierOK describes a response with status code 200, with default header values.

A successful response containing the created Notifier.
*/
type CreateNotifierOK struct {
	Payload *models.Configv1CreateNotifierResponse
}

// IsSuccess returns true when this create notifier o k response has a 2xx status code
func (o *CreateNotifierOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create notifier o k response has a 3xx status code
func (o *CreateNotifierOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create notifier o k response has a 4xx status code
func (o *CreateNotifierOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create notifier o k response has a 5xx status code
func (o *CreateNotifierOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create notifier o k response a status code equal to that given
func (o *CreateNotifierOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create notifier o k response
func (o *CreateNotifierOK) Code() int {
	return 200
}

func (o *CreateNotifierOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/notifiers][%d] createNotifierOK  %+v", 200, o.Payload)
}

func (o *CreateNotifierOK) String() string {
	return fmt.Sprintf("[POST /api/v1/config/notifiers][%d] createNotifierOK  %+v", 200, o.Payload)
}

func (o *CreateNotifierOK) GetPayload() *models.Configv1CreateNotifierResponse {
	return o.Payload
}

func (o *CreateNotifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1CreateNotifierResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNotifierBadRequest creates a CreateNotifierBadRequest with default headers values
func NewCreateNotifierBadRequest() *CreateNotifierBadRequest {
	return &CreateNotifierBadRequest{}
}

/*
CreateNotifierBadRequest describes a response with status code 400, with default header values.

Cannot create the Notifier because the request is invalid.
*/
type CreateNotifierBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create notifier bad request response has a 2xx status code
func (o *CreateNotifierBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create notifier bad request response has a 3xx status code
func (o *CreateNotifierBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create notifier bad request response has a 4xx status code
func (o *CreateNotifierBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create notifier bad request response has a 5xx status code
func (o *CreateNotifierBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create notifier bad request response a status code equal to that given
func (o *CreateNotifierBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create notifier bad request response
func (o *CreateNotifierBadRequest) Code() int {
	return 400
}

func (o *CreateNotifierBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/notifiers][%d] createNotifierBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNotifierBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v1/config/notifiers][%d] createNotifierBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNotifierBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateNotifierBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNotifierConflict creates a CreateNotifierConflict with default headers values
func NewCreateNotifierConflict() *CreateNotifierConflict {
	return &CreateNotifierConflict{}
}

/*
CreateNotifierConflict describes a response with status code 409, with default header values.

Cannot create the Notifier because there is a conflict with an existing Notifier.
*/
type CreateNotifierConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create notifier conflict response has a 2xx status code
func (o *CreateNotifierConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create notifier conflict response has a 3xx status code
func (o *CreateNotifierConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create notifier conflict response has a 4xx status code
func (o *CreateNotifierConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create notifier conflict response has a 5xx status code
func (o *CreateNotifierConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create notifier conflict response a status code equal to that given
func (o *CreateNotifierConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create notifier conflict response
func (o *CreateNotifierConflict) Code() int {
	return 409
}

func (o *CreateNotifierConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/notifiers][%d] createNotifierConflict  %+v", 409, o.Payload)
}

func (o *CreateNotifierConflict) String() string {
	return fmt.Sprintf("[POST /api/v1/config/notifiers][%d] createNotifierConflict  %+v", 409, o.Payload)
}

func (o *CreateNotifierConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateNotifierConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNotifierInternalServerError creates a CreateNotifierInternalServerError with default headers values
func NewCreateNotifierInternalServerError() *CreateNotifierInternalServerError {
	return &CreateNotifierInternalServerError{}
}

/*
CreateNotifierInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type CreateNotifierInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this create notifier internal server error response has a 2xx status code
func (o *CreateNotifierInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create notifier internal server error response has a 3xx status code
func (o *CreateNotifierInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create notifier internal server error response has a 4xx status code
func (o *CreateNotifierInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create notifier internal server error response has a 5xx status code
func (o *CreateNotifierInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create notifier internal server error response a status code equal to that given
func (o *CreateNotifierInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create notifier internal server error response
func (o *CreateNotifierInternalServerError) Code() int {
	return 500
}

func (o *CreateNotifierInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/notifiers][%d] createNotifierInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateNotifierInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v1/config/notifiers][%d] createNotifierInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateNotifierInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *CreateNotifierInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNotifierDefault creates a CreateNotifierDefault with default headers values
func NewCreateNotifierDefault(code int) *CreateNotifierDefault {
	return &CreateNotifierDefault{
		_statusCode: code,
	}
}

/*
CreateNotifierDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type CreateNotifierDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this create notifier default response has a 2xx status code
func (o *CreateNotifierDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create notifier default response has a 3xx status code
func (o *CreateNotifierDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create notifier default response has a 4xx status code
func (o *CreateNotifierDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create notifier default response has a 5xx status code
func (o *CreateNotifierDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create notifier default response a status code equal to that given
func (o *CreateNotifierDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create notifier default response
func (o *CreateNotifierDefault) Code() int {
	return o._statusCode
}

func (o *CreateNotifierDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/config/notifiers][%d] CreateNotifier default  %+v", o._statusCode, o.Payload)
}

func (o *CreateNotifierDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/config/notifiers][%d] CreateNotifier default  %+v", o._statusCode, o.Payload)
}

func (o *CreateNotifierDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreateNotifierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
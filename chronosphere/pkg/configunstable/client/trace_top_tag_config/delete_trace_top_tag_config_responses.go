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

// DeleteTraceTopTagConfigReader is a Reader for the DeleteTraceTopTagConfig structure.
type DeleteTraceTopTagConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTraceTopTagConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTraceTopTagConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTraceTopTagConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTraceTopTagConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTraceTopTagConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTraceTopTagConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTraceTopTagConfigOK creates a DeleteTraceTopTagConfigOK with default headers values
func NewDeleteTraceTopTagConfigOK() *DeleteTraceTopTagConfigOK {
	return &DeleteTraceTopTagConfigOK{}
}

/*
DeleteTraceTopTagConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteTraceTopTagConfigOK struct {
	Payload models.ConfigunstableDeleteTraceTopTagConfigResponse
}

// IsSuccess returns true when this delete trace top tag config o k response has a 2xx status code
func (o *DeleteTraceTopTagConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete trace top tag config o k response has a 3xx status code
func (o *DeleteTraceTopTagConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace top tag config o k response has a 4xx status code
func (o *DeleteTraceTopTagConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete trace top tag config o k response has a 5xx status code
func (o *DeleteTraceTopTagConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete trace top tag config o k response a status code equal to that given
func (o *DeleteTraceTopTagConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete trace top tag config o k response
func (o *DeleteTraceTopTagConfigOK) Code() int {
	return 200
}

func (o *DeleteTraceTopTagConfigOK) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-top-tag-config][%d] deleteTraceTopTagConfigOK  %+v", 200, o.Payload)
}

func (o *DeleteTraceTopTagConfigOK) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-top-tag-config][%d] deleteTraceTopTagConfigOK  %+v", 200, o.Payload)
}

func (o *DeleteTraceTopTagConfigOK) GetPayload() models.ConfigunstableDeleteTraceTopTagConfigResponse {
	return o.Payload
}

func (o *DeleteTraceTopTagConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceTopTagConfigBadRequest creates a DeleteTraceTopTagConfigBadRequest with default headers values
func NewDeleteTraceTopTagConfigBadRequest() *DeleteTraceTopTagConfigBadRequest {
	return &DeleteTraceTopTagConfigBadRequest{}
}

/*
DeleteTraceTopTagConfigBadRequest describes a response with status code 400, with default header values.

Cannot delete the TraceTopTagConfig because it is in use.
*/
type DeleteTraceTopTagConfigBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete trace top tag config bad request response has a 2xx status code
func (o *DeleteTraceTopTagConfigBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete trace top tag config bad request response has a 3xx status code
func (o *DeleteTraceTopTagConfigBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace top tag config bad request response has a 4xx status code
func (o *DeleteTraceTopTagConfigBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete trace top tag config bad request response has a 5xx status code
func (o *DeleteTraceTopTagConfigBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete trace top tag config bad request response a status code equal to that given
func (o *DeleteTraceTopTagConfigBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete trace top tag config bad request response
func (o *DeleteTraceTopTagConfigBadRequest) Code() int {
	return 400
}

func (o *DeleteTraceTopTagConfigBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-top-tag-config][%d] deleteTraceTopTagConfigBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTraceTopTagConfigBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-top-tag-config][%d] deleteTraceTopTagConfigBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTraceTopTagConfigBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteTraceTopTagConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceTopTagConfigNotFound creates a DeleteTraceTopTagConfigNotFound with default headers values
func NewDeleteTraceTopTagConfigNotFound() *DeleteTraceTopTagConfigNotFound {
	return &DeleteTraceTopTagConfigNotFound{}
}

/*
DeleteTraceTopTagConfigNotFound describes a response with status code 404, with default header values.

Cannot delete the TraceTopTagConfig because the slug does not exist.
*/
type DeleteTraceTopTagConfigNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete trace top tag config not found response has a 2xx status code
func (o *DeleteTraceTopTagConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete trace top tag config not found response has a 3xx status code
func (o *DeleteTraceTopTagConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace top tag config not found response has a 4xx status code
func (o *DeleteTraceTopTagConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete trace top tag config not found response has a 5xx status code
func (o *DeleteTraceTopTagConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete trace top tag config not found response a status code equal to that given
func (o *DeleteTraceTopTagConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete trace top tag config not found response
func (o *DeleteTraceTopTagConfigNotFound) Code() int {
	return 404
}

func (o *DeleteTraceTopTagConfigNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-top-tag-config][%d] deleteTraceTopTagConfigNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTraceTopTagConfigNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-top-tag-config][%d] deleteTraceTopTagConfigNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTraceTopTagConfigNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteTraceTopTagConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceTopTagConfigInternalServerError creates a DeleteTraceTopTagConfigInternalServerError with default headers values
func NewDeleteTraceTopTagConfigInternalServerError() *DeleteTraceTopTagConfigInternalServerError {
	return &DeleteTraceTopTagConfigInternalServerError{}
}

/*
DeleteTraceTopTagConfigInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type DeleteTraceTopTagConfigInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this delete trace top tag config internal server error response has a 2xx status code
func (o *DeleteTraceTopTagConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete trace top tag config internal server error response has a 3xx status code
func (o *DeleteTraceTopTagConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete trace top tag config internal server error response has a 4xx status code
func (o *DeleteTraceTopTagConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete trace top tag config internal server error response has a 5xx status code
func (o *DeleteTraceTopTagConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete trace top tag config internal server error response a status code equal to that given
func (o *DeleteTraceTopTagConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete trace top tag config internal server error response
func (o *DeleteTraceTopTagConfigInternalServerError) Code() int {
	return 500
}

func (o *DeleteTraceTopTagConfigInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-top-tag-config][%d] deleteTraceTopTagConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTraceTopTagConfigInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-top-tag-config][%d] deleteTraceTopTagConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTraceTopTagConfigInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *DeleteTraceTopTagConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTraceTopTagConfigDefault creates a DeleteTraceTopTagConfigDefault with default headers values
func NewDeleteTraceTopTagConfigDefault(code int) *DeleteTraceTopTagConfigDefault {
	return &DeleteTraceTopTagConfigDefault{
		_statusCode: code,
	}
}

/*
DeleteTraceTopTagConfigDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type DeleteTraceTopTagConfigDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this delete trace top tag config default response has a 2xx status code
func (o *DeleteTraceTopTagConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete trace top tag config default response has a 3xx status code
func (o *DeleteTraceTopTagConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete trace top tag config default response has a 4xx status code
func (o *DeleteTraceTopTagConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete trace top tag config default response has a 5xx status code
func (o *DeleteTraceTopTagConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete trace top tag config default response a status code equal to that given
func (o *DeleteTraceTopTagConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete trace top tag config default response
func (o *DeleteTraceTopTagConfigDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTraceTopTagConfigDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-top-tag-config][%d] DeleteTraceTopTagConfig default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTraceTopTagConfigDefault) String() string {
	return fmt.Sprintf("[DELETE /api/unstable/config/trace-top-tag-config][%d] DeleteTraceTopTagConfig default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTraceTopTagConfigDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteTraceTopTagConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package recording_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
)

// UpdateRecordingRuleReader is a Reader for the UpdateRecordingRule structure.
type UpdateRecordingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRecordingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRecordingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRecordingRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRecordingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateRecordingRuleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRecordingRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateRecordingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRecordingRuleOK creates a UpdateRecordingRuleOK with default headers values
func NewUpdateRecordingRuleOK() *UpdateRecordingRuleOK {
	return &UpdateRecordingRuleOK{}
}

/*
UpdateRecordingRuleOK describes a response with status code 200, with default header values.

A successful response containing the updated RecordingRule.
*/
type UpdateRecordingRuleOK struct {
	Payload *models.Configv1UpdateRecordingRuleResponse
}

// IsSuccess returns true when this update recording rule o k response has a 2xx status code
func (o *UpdateRecordingRuleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update recording rule o k response has a 3xx status code
func (o *UpdateRecordingRuleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update recording rule o k response has a 4xx status code
func (o *UpdateRecordingRuleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update recording rule o k response has a 5xx status code
func (o *UpdateRecordingRuleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update recording rule o k response a status code equal to that given
func (o *UpdateRecordingRuleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update recording rule o k response
func (o *UpdateRecordingRuleOK) Code() int {
	return 200
}

func (o *UpdateRecordingRuleOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/recording-rules/{slug}][%d] updateRecordingRuleOK  %+v", 200, o.Payload)
}

func (o *UpdateRecordingRuleOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/recording-rules/{slug}][%d] updateRecordingRuleOK  %+v", 200, o.Payload)
}

func (o *UpdateRecordingRuleOK) GetPayload() *models.Configv1UpdateRecordingRuleResponse {
	return o.Payload
}

func (o *UpdateRecordingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1UpdateRecordingRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRecordingRuleBadRequest creates a UpdateRecordingRuleBadRequest with default headers values
func NewUpdateRecordingRuleBadRequest() *UpdateRecordingRuleBadRequest {
	return &UpdateRecordingRuleBadRequest{}
}

/*
UpdateRecordingRuleBadRequest describes a response with status code 400, with default header values.

Cannot update the RecordingRule because the request is invalid.
*/
type UpdateRecordingRuleBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update recording rule bad request response has a 2xx status code
func (o *UpdateRecordingRuleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update recording rule bad request response has a 3xx status code
func (o *UpdateRecordingRuleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update recording rule bad request response has a 4xx status code
func (o *UpdateRecordingRuleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update recording rule bad request response has a 5xx status code
func (o *UpdateRecordingRuleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update recording rule bad request response a status code equal to that given
func (o *UpdateRecordingRuleBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update recording rule bad request response
func (o *UpdateRecordingRuleBadRequest) Code() int {
	return 400
}

func (o *UpdateRecordingRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/recording-rules/{slug}][%d] updateRecordingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRecordingRuleBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/recording-rules/{slug}][%d] updateRecordingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRecordingRuleBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateRecordingRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRecordingRuleNotFound creates a UpdateRecordingRuleNotFound with default headers values
func NewUpdateRecordingRuleNotFound() *UpdateRecordingRuleNotFound {
	return &UpdateRecordingRuleNotFound{}
}

/*
UpdateRecordingRuleNotFound describes a response with status code 404, with default header values.

Cannot update the RecordingRule because the slug does not exist.
*/
type UpdateRecordingRuleNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update recording rule not found response has a 2xx status code
func (o *UpdateRecordingRuleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update recording rule not found response has a 3xx status code
func (o *UpdateRecordingRuleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update recording rule not found response has a 4xx status code
func (o *UpdateRecordingRuleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update recording rule not found response has a 5xx status code
func (o *UpdateRecordingRuleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update recording rule not found response a status code equal to that given
func (o *UpdateRecordingRuleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update recording rule not found response
func (o *UpdateRecordingRuleNotFound) Code() int {
	return 404
}

func (o *UpdateRecordingRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/recording-rules/{slug}][%d] updateRecordingRuleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRecordingRuleNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/recording-rules/{slug}][%d] updateRecordingRuleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRecordingRuleNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateRecordingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRecordingRuleConflict creates a UpdateRecordingRuleConflict with default headers values
func NewUpdateRecordingRuleConflict() *UpdateRecordingRuleConflict {
	return &UpdateRecordingRuleConflict{}
}

/*
UpdateRecordingRuleConflict describes a response with status code 409, with default header values.

Cannot update the RecordingRule because there is a conflict with an existing RecordingRule.
*/
type UpdateRecordingRuleConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update recording rule conflict response has a 2xx status code
func (o *UpdateRecordingRuleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update recording rule conflict response has a 3xx status code
func (o *UpdateRecordingRuleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update recording rule conflict response has a 4xx status code
func (o *UpdateRecordingRuleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update recording rule conflict response has a 5xx status code
func (o *UpdateRecordingRuleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update recording rule conflict response a status code equal to that given
func (o *UpdateRecordingRuleConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update recording rule conflict response
func (o *UpdateRecordingRuleConflict) Code() int {
	return 409
}

func (o *UpdateRecordingRuleConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/recording-rules/{slug}][%d] updateRecordingRuleConflict  %+v", 409, o.Payload)
}

func (o *UpdateRecordingRuleConflict) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/recording-rules/{slug}][%d] updateRecordingRuleConflict  %+v", 409, o.Payload)
}

func (o *UpdateRecordingRuleConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateRecordingRuleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRecordingRuleInternalServerError creates a UpdateRecordingRuleInternalServerError with default headers values
func NewUpdateRecordingRuleInternalServerError() *UpdateRecordingRuleInternalServerError {
	return &UpdateRecordingRuleInternalServerError{}
}

/*
UpdateRecordingRuleInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateRecordingRuleInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update recording rule internal server error response has a 2xx status code
func (o *UpdateRecordingRuleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update recording rule internal server error response has a 3xx status code
func (o *UpdateRecordingRuleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update recording rule internal server error response has a 4xx status code
func (o *UpdateRecordingRuleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update recording rule internal server error response has a 5xx status code
func (o *UpdateRecordingRuleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update recording rule internal server error response a status code equal to that given
func (o *UpdateRecordingRuleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update recording rule internal server error response
func (o *UpdateRecordingRuleInternalServerError) Code() int {
	return 500
}

func (o *UpdateRecordingRuleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/recording-rules/{slug}][%d] updateRecordingRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRecordingRuleInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/recording-rules/{slug}][%d] updateRecordingRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRecordingRuleInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateRecordingRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRecordingRuleDefault creates a UpdateRecordingRuleDefault with default headers values
func NewUpdateRecordingRuleDefault(code int) *UpdateRecordingRuleDefault {
	return &UpdateRecordingRuleDefault{
		_statusCode: code,
	}
}

/*
UpdateRecordingRuleDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateRecordingRuleDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update recording rule default response has a 2xx status code
func (o *UpdateRecordingRuleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update recording rule default response has a 3xx status code
func (o *UpdateRecordingRuleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update recording rule default response has a 4xx status code
func (o *UpdateRecordingRuleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update recording rule default response has a 5xx status code
func (o *UpdateRecordingRuleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update recording rule default response a status code equal to that given
func (o *UpdateRecordingRuleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update recording rule default response
func (o *UpdateRecordingRuleDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRecordingRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/recording-rules/{slug}][%d] UpdateRecordingRule default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRecordingRuleDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/recording-rules/{slug}][%d] UpdateRecordingRule default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRecordingRuleDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateRecordingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateRecordingRuleBody update recording rule body
swagger:model UpdateRecordingRuleBody
*/
type UpdateRecordingRuleBody struct {

	// If true, the RecordingRule will be created if it does not already exist, identified by slug. If false, an error will be returned if the RecordingRule does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// If true, the RecordingRule will not be created nor updated, and no response RecordingRule will be returned. The response will return an error if the given RecordingRule is invalid.
	DryRun bool `json:"dry_run,omitempty"`

	// recording rule
	RecordingRule *models.Configv1RecordingRule `json:"recording_rule,omitempty"`
}

// Validate validates this update recording rule body
func (o *UpdateRecordingRuleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRecordingRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateRecordingRuleBody) validateRecordingRule(formats strfmt.Registry) error {
	if swag.IsZero(o.RecordingRule) { // not required
		return nil
	}

	if o.RecordingRule != nil {
		if err := o.RecordingRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "recording_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "recording_rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update recording rule body based on the context it is used
func (o *UpdateRecordingRuleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRecordingRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateRecordingRuleBody) contextValidateRecordingRule(ctx context.Context, formats strfmt.Registry) error {

	if o.RecordingRule != nil {
		if err := o.RecordingRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "recording_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "recording_rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateRecordingRuleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateRecordingRuleBody) UnmarshalBinary(b []byte) error {
	var res UpdateRecordingRuleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
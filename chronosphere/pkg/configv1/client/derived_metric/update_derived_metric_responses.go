// Code generated by go-swagger; DO NOT EDIT.

package derived_metric

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

// UpdateDerivedMetricReader is a Reader for the UpdateDerivedMetric structure.
type UpdateDerivedMetricReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDerivedMetricReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDerivedMetricOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDerivedMetricBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDerivedMetricNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateDerivedMetricConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateDerivedMetricInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateDerivedMetricDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDerivedMetricOK creates a UpdateDerivedMetricOK with default headers values
func NewUpdateDerivedMetricOK() *UpdateDerivedMetricOK {
	return &UpdateDerivedMetricOK{}
}

/*
UpdateDerivedMetricOK describes a response with status code 200, with default header values.

A successful response containing the updated DerivedMetric.
*/
type UpdateDerivedMetricOK struct {
	Payload *models.Configv1UpdateDerivedMetricResponse
}

// IsSuccess returns true when this update derived metric o k response has a 2xx status code
func (o *UpdateDerivedMetricOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update derived metric o k response has a 3xx status code
func (o *UpdateDerivedMetricOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update derived metric o k response has a 4xx status code
func (o *UpdateDerivedMetricOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update derived metric o k response has a 5xx status code
func (o *UpdateDerivedMetricOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update derived metric o k response a status code equal to that given
func (o *UpdateDerivedMetricOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update derived metric o k response
func (o *UpdateDerivedMetricOK) Code() int {
	return 200
}

func (o *UpdateDerivedMetricOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/derived-metrics/{slug}][%d] updateDerivedMetricOK  %+v", 200, o.Payload)
}

func (o *UpdateDerivedMetricOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/derived-metrics/{slug}][%d] updateDerivedMetricOK  %+v", 200, o.Payload)
}

func (o *UpdateDerivedMetricOK) GetPayload() *models.Configv1UpdateDerivedMetricResponse {
	return o.Payload
}

func (o *UpdateDerivedMetricOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Configv1UpdateDerivedMetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDerivedMetricBadRequest creates a UpdateDerivedMetricBadRequest with default headers values
func NewUpdateDerivedMetricBadRequest() *UpdateDerivedMetricBadRequest {
	return &UpdateDerivedMetricBadRequest{}
}

/*
UpdateDerivedMetricBadRequest describes a response with status code 400, with default header values.

Cannot update the DerivedMetric because the request is invalid.
*/
type UpdateDerivedMetricBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update derived metric bad request response has a 2xx status code
func (o *UpdateDerivedMetricBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update derived metric bad request response has a 3xx status code
func (o *UpdateDerivedMetricBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update derived metric bad request response has a 4xx status code
func (o *UpdateDerivedMetricBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update derived metric bad request response has a 5xx status code
func (o *UpdateDerivedMetricBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update derived metric bad request response a status code equal to that given
func (o *UpdateDerivedMetricBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update derived metric bad request response
func (o *UpdateDerivedMetricBadRequest) Code() int {
	return 400
}

func (o *UpdateDerivedMetricBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/derived-metrics/{slug}][%d] updateDerivedMetricBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDerivedMetricBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/derived-metrics/{slug}][%d] updateDerivedMetricBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDerivedMetricBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateDerivedMetricBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDerivedMetricNotFound creates a UpdateDerivedMetricNotFound with default headers values
func NewUpdateDerivedMetricNotFound() *UpdateDerivedMetricNotFound {
	return &UpdateDerivedMetricNotFound{}
}

/*
UpdateDerivedMetricNotFound describes a response with status code 404, with default header values.

Cannot update the DerivedMetric because the slug does not exist.
*/
type UpdateDerivedMetricNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update derived metric not found response has a 2xx status code
func (o *UpdateDerivedMetricNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update derived metric not found response has a 3xx status code
func (o *UpdateDerivedMetricNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update derived metric not found response has a 4xx status code
func (o *UpdateDerivedMetricNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update derived metric not found response has a 5xx status code
func (o *UpdateDerivedMetricNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update derived metric not found response a status code equal to that given
func (o *UpdateDerivedMetricNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update derived metric not found response
func (o *UpdateDerivedMetricNotFound) Code() int {
	return 404
}

func (o *UpdateDerivedMetricNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/derived-metrics/{slug}][%d] updateDerivedMetricNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDerivedMetricNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/derived-metrics/{slug}][%d] updateDerivedMetricNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDerivedMetricNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateDerivedMetricNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDerivedMetricConflict creates a UpdateDerivedMetricConflict with default headers values
func NewUpdateDerivedMetricConflict() *UpdateDerivedMetricConflict {
	return &UpdateDerivedMetricConflict{}
}

/*
UpdateDerivedMetricConflict describes a response with status code 409, with default header values.

Cannot update the DerivedMetric because there is a conflict with an existing DerivedMetric.
*/
type UpdateDerivedMetricConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update derived metric conflict response has a 2xx status code
func (o *UpdateDerivedMetricConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update derived metric conflict response has a 3xx status code
func (o *UpdateDerivedMetricConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update derived metric conflict response has a 4xx status code
func (o *UpdateDerivedMetricConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update derived metric conflict response has a 5xx status code
func (o *UpdateDerivedMetricConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update derived metric conflict response a status code equal to that given
func (o *UpdateDerivedMetricConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update derived metric conflict response
func (o *UpdateDerivedMetricConflict) Code() int {
	return 409
}

func (o *UpdateDerivedMetricConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/derived-metrics/{slug}][%d] updateDerivedMetricConflict  %+v", 409, o.Payload)
}

func (o *UpdateDerivedMetricConflict) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/derived-metrics/{slug}][%d] updateDerivedMetricConflict  %+v", 409, o.Payload)
}

func (o *UpdateDerivedMetricConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateDerivedMetricConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDerivedMetricInternalServerError creates a UpdateDerivedMetricInternalServerError with default headers values
func NewUpdateDerivedMetricInternalServerError() *UpdateDerivedMetricInternalServerError {
	return &UpdateDerivedMetricInternalServerError{}
}

/*
UpdateDerivedMetricInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateDerivedMetricInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update derived metric internal server error response has a 2xx status code
func (o *UpdateDerivedMetricInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update derived metric internal server error response has a 3xx status code
func (o *UpdateDerivedMetricInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update derived metric internal server error response has a 4xx status code
func (o *UpdateDerivedMetricInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update derived metric internal server error response has a 5xx status code
func (o *UpdateDerivedMetricInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update derived metric internal server error response a status code equal to that given
func (o *UpdateDerivedMetricInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update derived metric internal server error response
func (o *UpdateDerivedMetricInternalServerError) Code() int {
	return 500
}

func (o *UpdateDerivedMetricInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/derived-metrics/{slug}][%d] updateDerivedMetricInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDerivedMetricInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/derived-metrics/{slug}][%d] updateDerivedMetricInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDerivedMetricInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateDerivedMetricInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDerivedMetricDefault creates a UpdateDerivedMetricDefault with default headers values
func NewUpdateDerivedMetricDefault(code int) *UpdateDerivedMetricDefault {
	return &UpdateDerivedMetricDefault{
		_statusCode: code,
	}
}

/*
UpdateDerivedMetricDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateDerivedMetricDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update derived metric default response has a 2xx status code
func (o *UpdateDerivedMetricDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update derived metric default response has a 3xx status code
func (o *UpdateDerivedMetricDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update derived metric default response has a 4xx status code
func (o *UpdateDerivedMetricDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update derived metric default response has a 5xx status code
func (o *UpdateDerivedMetricDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update derived metric default response a status code equal to that given
func (o *UpdateDerivedMetricDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update derived metric default response
func (o *UpdateDerivedMetricDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDerivedMetricDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/config/derived-metrics/{slug}][%d] UpdateDerivedMetric default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDerivedMetricDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/config/derived-metrics/{slug}][%d] UpdateDerivedMetric default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDerivedMetricDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateDerivedMetricDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateDerivedMetricBody update derived metric body
swagger:model UpdateDerivedMetricBody
*/
type UpdateDerivedMetricBody struct {

	// If true, the DerivedMetric will be created if it does not already exist, identified by slug. If false, an error will be returned if the DerivedMetric does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// derived metric
	DerivedMetric *models.Configv1DerivedMetric `json:"derived_metric,omitempty"`

	// If true, the DerivedMetric will not be created nor updated, and no response DerivedMetric will be returned. The response will return an error if the given DerivedMetric is invalid.
	DryRun bool `json:"dry_run,omitempty"`
}

// Validate validates this update derived metric body
func (o *UpdateDerivedMetricBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDerivedMetric(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDerivedMetricBody) validateDerivedMetric(formats strfmt.Registry) error {
	if swag.IsZero(o.DerivedMetric) { // not required
		return nil
	}

	if o.DerivedMetric != nil {
		if err := o.DerivedMetric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "derived_metric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "derived_metric")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update derived metric body based on the context it is used
func (o *UpdateDerivedMetricBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDerivedMetric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDerivedMetricBody) contextValidateDerivedMetric(ctx context.Context, formats strfmt.Registry) error {

	if o.DerivedMetric != nil {
		if err := o.DerivedMetric.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "derived_metric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "derived_metric")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDerivedMetricBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDerivedMetricBody) UnmarshalBinary(b []byte) error {
	var res UpdateDerivedMetricBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package saved_trace_search

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

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configunstable/models"
)

// UpdateSavedTraceSearchReader is a Reader for the UpdateSavedTraceSearch structure.
type UpdateSavedTraceSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSavedTraceSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSavedTraceSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSavedTraceSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateSavedTraceSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateSavedTraceSearchConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSavedTraceSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateSavedTraceSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSavedTraceSearchOK creates a UpdateSavedTraceSearchOK with default headers values
func NewUpdateSavedTraceSearchOK() *UpdateSavedTraceSearchOK {
	return &UpdateSavedTraceSearchOK{}
}

/*
UpdateSavedTraceSearchOK describes a response with status code 200, with default header values.

A successful response containing the updated SavedTraceSearch.
*/
type UpdateSavedTraceSearchOK struct {
	Payload *models.ConfigunstableUpdateSavedTraceSearchResponse
}

// IsSuccess returns true when this update saved trace search o k response has a 2xx status code
func (o *UpdateSavedTraceSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update saved trace search o k response has a 3xx status code
func (o *UpdateSavedTraceSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update saved trace search o k response has a 4xx status code
func (o *UpdateSavedTraceSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update saved trace search o k response has a 5xx status code
func (o *UpdateSavedTraceSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update saved trace search o k response a status code equal to that given
func (o *UpdateSavedTraceSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update saved trace search o k response
func (o *UpdateSavedTraceSearchOK) Code() int {
	return 200
}

func (o *UpdateSavedTraceSearchOK) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/saved-trace-searches/{slug}][%d] updateSavedTraceSearchOK  %+v", 200, o.Payload)
}

func (o *UpdateSavedTraceSearchOK) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/saved-trace-searches/{slug}][%d] updateSavedTraceSearchOK  %+v", 200, o.Payload)
}

func (o *UpdateSavedTraceSearchOK) GetPayload() *models.ConfigunstableUpdateSavedTraceSearchResponse {
	return o.Payload
}

func (o *UpdateSavedTraceSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigunstableUpdateSavedTraceSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSavedTraceSearchBadRequest creates a UpdateSavedTraceSearchBadRequest with default headers values
func NewUpdateSavedTraceSearchBadRequest() *UpdateSavedTraceSearchBadRequest {
	return &UpdateSavedTraceSearchBadRequest{}
}

/*
UpdateSavedTraceSearchBadRequest describes a response with status code 400, with default header values.

Cannot update the SavedTraceSearch because the request is invalid.
*/
type UpdateSavedTraceSearchBadRequest struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update saved trace search bad request response has a 2xx status code
func (o *UpdateSavedTraceSearchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update saved trace search bad request response has a 3xx status code
func (o *UpdateSavedTraceSearchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update saved trace search bad request response has a 4xx status code
func (o *UpdateSavedTraceSearchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update saved trace search bad request response has a 5xx status code
func (o *UpdateSavedTraceSearchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update saved trace search bad request response a status code equal to that given
func (o *UpdateSavedTraceSearchBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update saved trace search bad request response
func (o *UpdateSavedTraceSearchBadRequest) Code() int {
	return 400
}

func (o *UpdateSavedTraceSearchBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/saved-trace-searches/{slug}][%d] updateSavedTraceSearchBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSavedTraceSearchBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/saved-trace-searches/{slug}][%d] updateSavedTraceSearchBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSavedTraceSearchBadRequest) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateSavedTraceSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSavedTraceSearchNotFound creates a UpdateSavedTraceSearchNotFound with default headers values
func NewUpdateSavedTraceSearchNotFound() *UpdateSavedTraceSearchNotFound {
	return &UpdateSavedTraceSearchNotFound{}
}

/*
UpdateSavedTraceSearchNotFound describes a response with status code 404, with default header values.

Cannot update the SavedTraceSearch because the slug does not exist.
*/
type UpdateSavedTraceSearchNotFound struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update saved trace search not found response has a 2xx status code
func (o *UpdateSavedTraceSearchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update saved trace search not found response has a 3xx status code
func (o *UpdateSavedTraceSearchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update saved trace search not found response has a 4xx status code
func (o *UpdateSavedTraceSearchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update saved trace search not found response has a 5xx status code
func (o *UpdateSavedTraceSearchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update saved trace search not found response a status code equal to that given
func (o *UpdateSavedTraceSearchNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update saved trace search not found response
func (o *UpdateSavedTraceSearchNotFound) Code() int {
	return 404
}

func (o *UpdateSavedTraceSearchNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/saved-trace-searches/{slug}][%d] updateSavedTraceSearchNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSavedTraceSearchNotFound) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/saved-trace-searches/{slug}][%d] updateSavedTraceSearchNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSavedTraceSearchNotFound) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateSavedTraceSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSavedTraceSearchConflict creates a UpdateSavedTraceSearchConflict with default headers values
func NewUpdateSavedTraceSearchConflict() *UpdateSavedTraceSearchConflict {
	return &UpdateSavedTraceSearchConflict{}
}

/*
UpdateSavedTraceSearchConflict describes a response with status code 409, with default header values.

Cannot update the SavedTraceSearch because there is a conflict with an existing SavedTraceSearch.
*/
type UpdateSavedTraceSearchConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update saved trace search conflict response has a 2xx status code
func (o *UpdateSavedTraceSearchConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update saved trace search conflict response has a 3xx status code
func (o *UpdateSavedTraceSearchConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update saved trace search conflict response has a 4xx status code
func (o *UpdateSavedTraceSearchConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update saved trace search conflict response has a 5xx status code
func (o *UpdateSavedTraceSearchConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update saved trace search conflict response a status code equal to that given
func (o *UpdateSavedTraceSearchConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update saved trace search conflict response
func (o *UpdateSavedTraceSearchConflict) Code() int {
	return 409
}

func (o *UpdateSavedTraceSearchConflict) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/saved-trace-searches/{slug}][%d] updateSavedTraceSearchConflict  %+v", 409, o.Payload)
}

func (o *UpdateSavedTraceSearchConflict) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/saved-trace-searches/{slug}][%d] updateSavedTraceSearchConflict  %+v", 409, o.Payload)
}

func (o *UpdateSavedTraceSearchConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateSavedTraceSearchConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSavedTraceSearchInternalServerError creates a UpdateSavedTraceSearchInternalServerError with default headers values
func NewUpdateSavedTraceSearchInternalServerError() *UpdateSavedTraceSearchInternalServerError {
	return &UpdateSavedTraceSearchInternalServerError{}
}

/*
UpdateSavedTraceSearchInternalServerError describes a response with status code 500, with default header values.

An unexpected error response.
*/
type UpdateSavedTraceSearchInternalServerError struct {
	Payload *models.APIError
}

// IsSuccess returns true when this update saved trace search internal server error response has a 2xx status code
func (o *UpdateSavedTraceSearchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update saved trace search internal server error response has a 3xx status code
func (o *UpdateSavedTraceSearchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update saved trace search internal server error response has a 4xx status code
func (o *UpdateSavedTraceSearchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update saved trace search internal server error response has a 5xx status code
func (o *UpdateSavedTraceSearchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update saved trace search internal server error response a status code equal to that given
func (o *UpdateSavedTraceSearchInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update saved trace search internal server error response
func (o *UpdateSavedTraceSearchInternalServerError) Code() int {
	return 500
}

func (o *UpdateSavedTraceSearchInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/saved-trace-searches/{slug}][%d] updateSavedTraceSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSavedTraceSearchInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/saved-trace-searches/{slug}][%d] updateSavedTraceSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSavedTraceSearchInternalServerError) GetPayload() *models.APIError {
	return o.Payload
}

func (o *UpdateSavedTraceSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSavedTraceSearchDefault creates a UpdateSavedTraceSearchDefault with default headers values
func NewUpdateSavedTraceSearchDefault(code int) *UpdateSavedTraceSearchDefault {
	return &UpdateSavedTraceSearchDefault{
		_statusCode: code,
	}
}

/*
UpdateSavedTraceSearchDefault describes a response with status code -1, with default header values.

An undefined error response.
*/
type UpdateSavedTraceSearchDefault struct {
	_statusCode int

	Payload models.GenericError
}

// IsSuccess returns true when this update saved trace search default response has a 2xx status code
func (o *UpdateSavedTraceSearchDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update saved trace search default response has a 3xx status code
func (o *UpdateSavedTraceSearchDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update saved trace search default response has a 4xx status code
func (o *UpdateSavedTraceSearchDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update saved trace search default response has a 5xx status code
func (o *UpdateSavedTraceSearchDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update saved trace search default response a status code equal to that given
func (o *UpdateSavedTraceSearchDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update saved trace search default response
func (o *UpdateSavedTraceSearchDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSavedTraceSearchDefault) Error() string {
	return fmt.Sprintf("[PUT /api/unstable/config/saved-trace-searches/{slug}][%d] UpdateSavedTraceSearch default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSavedTraceSearchDefault) String() string {
	return fmt.Sprintf("[PUT /api/unstable/config/saved-trace-searches/{slug}][%d] UpdateSavedTraceSearch default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSavedTraceSearchDefault) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateSavedTraceSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateSavedTraceSearchBody update saved trace search body
swagger:model UpdateSavedTraceSearchBody
*/
type UpdateSavedTraceSearchBody struct {

	// If true, the SavedTraceSearch will be created if it does not already exist, identified by slug. If false, an error will be returned if the SavedTraceSearch does not already exist.
	CreateIfMissing bool `json:"create_if_missing,omitempty"`

	// saved trace search
	SavedTraceSearch *models.ConfigunstableSavedTraceSearch `json:"saved_trace_search,omitempty"`
}

// Validate validates this update saved trace search body
func (o *UpdateSavedTraceSearchBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSavedTraceSearch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateSavedTraceSearchBody) validateSavedTraceSearch(formats strfmt.Registry) error {
	if swag.IsZero(o.SavedTraceSearch) { // not required
		return nil
	}

	if o.SavedTraceSearch != nil {
		if err := o.SavedTraceSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "saved_trace_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "saved_trace_search")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update saved trace search body based on the context it is used
func (o *UpdateSavedTraceSearchBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSavedTraceSearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateSavedTraceSearchBody) contextValidateSavedTraceSearch(ctx context.Context, formats strfmt.Registry) error {

	if o.SavedTraceSearch != nil {
		if err := o.SavedTraceSearch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "saved_trace_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "saved_trace_search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateSavedTraceSearchBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateSavedTraceSearchBody) UnmarshalBinary(b []byte) error {
	var res UpdateSavedTraceSearchBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

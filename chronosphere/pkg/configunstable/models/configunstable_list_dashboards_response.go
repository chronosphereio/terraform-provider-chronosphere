// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigunstableListDashboardsResponse configunstable list dashboards response
//
// swagger:model configunstableListDashboardsResponse
type ConfigunstableListDashboardsResponse struct {

	// dashboards
	Dashboards []*ConfigunstableDashboard `json:"dashboards"`

	// page
	Page *Configv1PageResult `json:"page,omitempty"`
}

// Validate validates this configunstable list dashboards response
func (m *ConfigunstableListDashboardsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableListDashboardsResponse) validateDashboards(formats strfmt.Registry) error {
	if swag.IsZero(m.Dashboards) { // not required
		return nil
	}

	for i := 0; i < len(m.Dashboards); i++ {
		if swag.IsZero(m.Dashboards[i]) { // not required
			continue
		}

		if m.Dashboards[i] != nil {
			if err := m.Dashboards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dashboards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dashboards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigunstableListDashboardsResponse) validatePage(formats strfmt.Registry) error {
	if swag.IsZero(m.Page) { // not required
		return nil
	}

	if m.Page != nil {
		if err := m.Page.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this configunstable list dashboards response based on the context it is used
func (m *ConfigunstableListDashboardsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDashboards(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigunstableListDashboardsResponse) contextValidateDashboards(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Dashboards); i++ {

		if m.Dashboards[i] != nil {
			if err := m.Dashboards[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dashboards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dashboards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigunstableListDashboardsResponse) contextValidatePage(ctx context.Context, formats strfmt.Registry) error {

	if m.Page != nil {
		if err := m.Page.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigunstableListDashboardsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigunstableListDashboardsResponse) UnmarshalBinary(b []byte) error {
	var res ConfigunstableListDashboardsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

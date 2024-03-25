// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1SendNotificationV1Request v1 send notification v1 request
//
// swagger:model v1SendNotificationV1Request
type V1SendNotificationV1Request struct {

	// notification
	Notification *V1Notification `json:"notification,omitempty"`
}

// Validate validates this v1 send notification v1 request
func (m *V1SendNotificationV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SendNotificationV1Request) validateNotification(formats strfmt.Registry) error {
	if swag.IsZero(m.Notification) { // not required
		return nil
	}

	if m.Notification != nil {
		if err := m.Notification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notification")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 send notification v1 request based on the context it is used
func (m *V1SendNotificationV1Request) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNotification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SendNotificationV1Request) contextValidateNotification(ctx context.Context, formats strfmt.Registry) error {

	if m.Notification != nil {
		if err := m.Notification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SendNotificationV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SendNotificationV1Request) UnmarshalBinary(b []byte) error {
	var res V1SendNotificationV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
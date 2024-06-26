// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1UserNotification v1 user notification
//
// swagger:model v1UserNotification
type V1UserNotification struct {

	// message
	Message string `json:"message,omitempty"`

	// notification Id
	NotificationID string `json:"notificationId,omitempty"`
}

// Validate validates this v1 user notification
func (m *V1UserNotification) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 user notification based on context it is used
func (m *V1UserNotification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1UserNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UserNotification) UnmarshalBinary(b []byte) error {
	var res V1UserNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

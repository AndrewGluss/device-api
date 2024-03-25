// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ActDeviceApiv1Status act device apiv1 status
//
// swagger:model act_device_apiv1Status
type ActDeviceApiv1Status string

func NewActDeviceApiv1Status(value ActDeviceApiv1Status) *ActDeviceApiv1Status {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ActDeviceApiv1Status.
func (m ActDeviceApiv1Status) Pointer() *ActDeviceApiv1Status {
	return &m
}

const (

	// ActDeviceApiv1StatusSTATUSCREATED captures enum value "STATUS_CREATED"
	ActDeviceApiv1StatusSTATUSCREATED ActDeviceApiv1Status = "STATUS_CREATED"

	// ActDeviceApiv1StatusSTATUSINPROGRESS captures enum value "STATUS_IN_PROGRESS"
	ActDeviceApiv1StatusSTATUSINPROGRESS ActDeviceApiv1Status = "STATUS_IN_PROGRESS"

	// ActDeviceApiv1StatusSTATUSDELIVERED captures enum value "STATUS_DELIVERED"
	ActDeviceApiv1StatusSTATUSDELIVERED ActDeviceApiv1Status = "STATUS_DELIVERED"
)

// for schema
var actDeviceApiv1StatusEnum []interface{}

func init() {
	var res []ActDeviceApiv1Status
	if err := json.Unmarshal([]byte(`["STATUS_CREATED","STATUS_IN_PROGRESS","STATUS_DELIVERED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		actDeviceApiv1StatusEnum = append(actDeviceApiv1StatusEnum, v)
	}
}

func (m ActDeviceApiv1Status) validateActDeviceApiv1StatusEnum(path, location string, value ActDeviceApiv1Status) error {
	if err := validate.EnumCase(path, location, value, actDeviceApiv1StatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this act device apiv1 status
func (m ActDeviceApiv1Status) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateActDeviceApiv1StatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this act device apiv1 status based on context it is used
func (m ActDeviceApiv1Status) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
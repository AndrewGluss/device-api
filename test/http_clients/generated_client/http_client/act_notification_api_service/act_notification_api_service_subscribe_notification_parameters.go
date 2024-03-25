// Code generated by go-swagger; DO NOT EDIT.

package act_notification_api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewActNotificationAPIServiceSubscribeNotificationParams creates a new ActNotificationAPIServiceSubscribeNotificationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActNotificationAPIServiceSubscribeNotificationParams() *ActNotificationAPIServiceSubscribeNotificationParams {
	return &ActNotificationAPIServiceSubscribeNotificationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActNotificationAPIServiceSubscribeNotificationParamsWithTimeout creates a new ActNotificationAPIServiceSubscribeNotificationParams object
// with the ability to set a timeout on a request.
func NewActNotificationAPIServiceSubscribeNotificationParamsWithTimeout(timeout time.Duration) *ActNotificationAPIServiceSubscribeNotificationParams {
	return &ActNotificationAPIServiceSubscribeNotificationParams{
		timeout: timeout,
	}
}

// NewActNotificationAPIServiceSubscribeNotificationParamsWithContext creates a new ActNotificationAPIServiceSubscribeNotificationParams object
// with the ability to set a context for a request.
func NewActNotificationAPIServiceSubscribeNotificationParamsWithContext(ctx context.Context) *ActNotificationAPIServiceSubscribeNotificationParams {
	return &ActNotificationAPIServiceSubscribeNotificationParams{
		Context: ctx,
	}
}

// NewActNotificationAPIServiceSubscribeNotificationParamsWithHTTPClient creates a new ActNotificationAPIServiceSubscribeNotificationParams object
// with the ability to set a custom HTTPClient for a request.
func NewActNotificationAPIServiceSubscribeNotificationParamsWithHTTPClient(client *http.Client) *ActNotificationAPIServiceSubscribeNotificationParams {
	return &ActNotificationAPIServiceSubscribeNotificationParams{
		HTTPClient: client,
	}
}

/*
ActNotificationAPIServiceSubscribeNotificationParams contains all the parameters to send to the API endpoint

	for the act notification Api service subscribe notification operation.

	Typically these are written to a http.Request.
*/
type ActNotificationAPIServiceSubscribeNotificationParams struct {

	// DeviceID.
	//
	// Format: uint64
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the act notification Api service subscribe notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActNotificationAPIServiceSubscribeNotificationParams) WithDefaults() *ActNotificationAPIServiceSubscribeNotificationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the act notification Api service subscribe notification params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActNotificationAPIServiceSubscribeNotificationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the act notification Api service subscribe notification params
func (o *ActNotificationAPIServiceSubscribeNotificationParams) WithTimeout(timeout time.Duration) *ActNotificationAPIServiceSubscribeNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the act notification Api service subscribe notification params
func (o *ActNotificationAPIServiceSubscribeNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the act notification Api service subscribe notification params
func (o *ActNotificationAPIServiceSubscribeNotificationParams) WithContext(ctx context.Context) *ActNotificationAPIServiceSubscribeNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the act notification Api service subscribe notification params
func (o *ActNotificationAPIServiceSubscribeNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the act notification Api service subscribe notification params
func (o *ActNotificationAPIServiceSubscribeNotificationParams) WithHTTPClient(client *http.Client) *ActNotificationAPIServiceSubscribeNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the act notification Api service subscribe notification params
func (o *ActNotificationAPIServiceSubscribeNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the act notification Api service subscribe notification params
func (o *ActNotificationAPIServiceSubscribeNotificationParams) WithDeviceID(deviceID string) *ActNotificationAPIServiceSubscribeNotificationParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the act notification Api service subscribe notification params
func (o *ActNotificationAPIServiceSubscribeNotificationParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *ActNotificationAPIServiceSubscribeNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package act_device_api_service

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

// NewActDeviceAPIServiceUpdateDeviceV1Params creates a new ActDeviceAPIServiceUpdateDeviceV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActDeviceAPIServiceUpdateDeviceV1Params() *ActDeviceAPIServiceUpdateDeviceV1Params {
	return &ActDeviceAPIServiceUpdateDeviceV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewActDeviceAPIServiceUpdateDeviceV1ParamsWithTimeout creates a new ActDeviceAPIServiceUpdateDeviceV1Params object
// with the ability to set a timeout on a request.
func NewActDeviceAPIServiceUpdateDeviceV1ParamsWithTimeout(timeout time.Duration) *ActDeviceAPIServiceUpdateDeviceV1Params {
	return &ActDeviceAPIServiceUpdateDeviceV1Params{
		timeout: timeout,
	}
}

// NewActDeviceAPIServiceUpdateDeviceV1ParamsWithContext creates a new ActDeviceAPIServiceUpdateDeviceV1Params object
// with the ability to set a context for a request.
func NewActDeviceAPIServiceUpdateDeviceV1ParamsWithContext(ctx context.Context) *ActDeviceAPIServiceUpdateDeviceV1Params {
	return &ActDeviceAPIServiceUpdateDeviceV1Params{
		Context: ctx,
	}
}

// NewActDeviceAPIServiceUpdateDeviceV1ParamsWithHTTPClient creates a new ActDeviceAPIServiceUpdateDeviceV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewActDeviceAPIServiceUpdateDeviceV1ParamsWithHTTPClient(client *http.Client) *ActDeviceAPIServiceUpdateDeviceV1Params {
	return &ActDeviceAPIServiceUpdateDeviceV1Params{
		HTTPClient: client,
	}
}

/*
ActDeviceAPIServiceUpdateDeviceV1Params contains all the parameters to send to the API endpoint

	for the act device Api service update device v1 operation.

	Typically these are written to a http.Request.
*/
type ActDeviceAPIServiceUpdateDeviceV1Params struct {

	// Body.
	Body ActDeviceAPIServiceUpdateDeviceV1Body

	// DeviceID.
	//
	// Format: uint64
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the act device Api service update device v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActDeviceAPIServiceUpdateDeviceV1Params) WithDefaults() *ActDeviceAPIServiceUpdateDeviceV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the act device Api service update device v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActDeviceAPIServiceUpdateDeviceV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the act device Api service update device v1 params
func (o *ActDeviceAPIServiceUpdateDeviceV1Params) WithTimeout(timeout time.Duration) *ActDeviceAPIServiceUpdateDeviceV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the act device Api service update device v1 params
func (o *ActDeviceAPIServiceUpdateDeviceV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the act device Api service update device v1 params
func (o *ActDeviceAPIServiceUpdateDeviceV1Params) WithContext(ctx context.Context) *ActDeviceAPIServiceUpdateDeviceV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the act device Api service update device v1 params
func (o *ActDeviceAPIServiceUpdateDeviceV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the act device Api service update device v1 params
func (o *ActDeviceAPIServiceUpdateDeviceV1Params) WithHTTPClient(client *http.Client) *ActDeviceAPIServiceUpdateDeviceV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the act device Api service update device v1 params
func (o *ActDeviceAPIServiceUpdateDeviceV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the act device Api service update device v1 params
func (o *ActDeviceAPIServiceUpdateDeviceV1Params) WithBody(body ActDeviceAPIServiceUpdateDeviceV1Body) *ActDeviceAPIServiceUpdateDeviceV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the act device Api service update device v1 params
func (o *ActDeviceAPIServiceUpdateDeviceV1Params) SetBody(body ActDeviceAPIServiceUpdateDeviceV1Body) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the act device Api service update device v1 params
func (o *ActDeviceAPIServiceUpdateDeviceV1Params) WithDeviceID(deviceID string) *ActDeviceAPIServiceUpdateDeviceV1Params {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the act device Api service update device v1 params
func (o *ActDeviceAPIServiceUpdateDeviceV1Params) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *ActDeviceAPIServiceUpdateDeviceV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
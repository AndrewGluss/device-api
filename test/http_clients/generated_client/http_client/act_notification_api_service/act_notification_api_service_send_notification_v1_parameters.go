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

	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_clients/generated_client/models"
)

// NewActNotificationAPIServiceSendNotificationV1Params creates a new ActNotificationAPIServiceSendNotificationV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActNotificationAPIServiceSendNotificationV1Params() *ActNotificationAPIServiceSendNotificationV1Params {
	return &ActNotificationAPIServiceSendNotificationV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewActNotificationAPIServiceSendNotificationV1ParamsWithTimeout creates a new ActNotificationAPIServiceSendNotificationV1Params object
// with the ability to set a timeout on a request.
func NewActNotificationAPIServiceSendNotificationV1ParamsWithTimeout(timeout time.Duration) *ActNotificationAPIServiceSendNotificationV1Params {
	return &ActNotificationAPIServiceSendNotificationV1Params{
		timeout: timeout,
	}
}

// NewActNotificationAPIServiceSendNotificationV1ParamsWithContext creates a new ActNotificationAPIServiceSendNotificationV1Params object
// with the ability to set a context for a request.
func NewActNotificationAPIServiceSendNotificationV1ParamsWithContext(ctx context.Context) *ActNotificationAPIServiceSendNotificationV1Params {
	return &ActNotificationAPIServiceSendNotificationV1Params{
		Context: ctx,
	}
}

// NewActNotificationAPIServiceSendNotificationV1ParamsWithHTTPClient creates a new ActNotificationAPIServiceSendNotificationV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewActNotificationAPIServiceSendNotificationV1ParamsWithHTTPClient(client *http.Client) *ActNotificationAPIServiceSendNotificationV1Params {
	return &ActNotificationAPIServiceSendNotificationV1Params{
		HTTPClient: client,
	}
}

/*
ActNotificationAPIServiceSendNotificationV1Params contains all the parameters to send to the API endpoint

	for the act notification Api service send notification v1 operation.

	Typically these are written to a http.Request.
*/
type ActNotificationAPIServiceSendNotificationV1Params struct {

	// Body.
	Body *models.V1SendNotificationV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the act notification Api service send notification v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActNotificationAPIServiceSendNotificationV1Params) WithDefaults() *ActNotificationAPIServiceSendNotificationV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the act notification Api service send notification v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActNotificationAPIServiceSendNotificationV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the act notification Api service send notification v1 params
func (o *ActNotificationAPIServiceSendNotificationV1Params) WithTimeout(timeout time.Duration) *ActNotificationAPIServiceSendNotificationV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the act notification Api service send notification v1 params
func (o *ActNotificationAPIServiceSendNotificationV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the act notification Api service send notification v1 params
func (o *ActNotificationAPIServiceSendNotificationV1Params) WithContext(ctx context.Context) *ActNotificationAPIServiceSendNotificationV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the act notification Api service send notification v1 params
func (o *ActNotificationAPIServiceSendNotificationV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the act notification Api service send notification v1 params
func (o *ActNotificationAPIServiceSendNotificationV1Params) WithHTTPClient(client *http.Client) *ActNotificationAPIServiceSendNotificationV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the act notification Api service send notification v1 params
func (o *ActNotificationAPIServiceSendNotificationV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the act notification Api service send notification v1 params
func (o *ActNotificationAPIServiceSendNotificationV1Params) WithBody(body *models.V1SendNotificationV1Request) *ActNotificationAPIServiceSendNotificationV1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the act notification Api service send notification v1 params
func (o *ActNotificationAPIServiceSendNotificationV1Params) SetBody(body *models.V1SendNotificationV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ActNotificationAPIServiceSendNotificationV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

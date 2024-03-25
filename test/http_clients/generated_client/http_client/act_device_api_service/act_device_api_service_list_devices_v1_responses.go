// Code generated by go-swagger; DO NOT EDIT.

package act_device_api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_clients/generated_client/models"
)

// ActDeviceAPIServiceListDevicesV1Reader is a Reader for the ActDeviceAPIServiceListDevicesV1 structure.
type ActDeviceAPIServiceListDevicesV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActDeviceAPIServiceListDevicesV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActDeviceAPIServiceListDevicesV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActDeviceAPIServiceListDevicesV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActDeviceAPIServiceListDevicesV1OK creates a ActDeviceAPIServiceListDevicesV1OK with default headers values
func NewActDeviceAPIServiceListDevicesV1OK() *ActDeviceAPIServiceListDevicesV1OK {
	return &ActDeviceAPIServiceListDevicesV1OK{}
}

/*
ActDeviceAPIServiceListDevicesV1OK describes a response with status code 200, with default header values.

A successful response.
*/
type ActDeviceAPIServiceListDevicesV1OK struct {
	Payload *models.V1ListDevicesV1Response
}

// IsSuccess returns true when this act device Api service list devices v1 o k response has a 2xx status code
func (o *ActDeviceAPIServiceListDevicesV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this act device Api service list devices v1 o k response has a 3xx status code
func (o *ActDeviceAPIServiceListDevicesV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this act device Api service list devices v1 o k response has a 4xx status code
func (o *ActDeviceAPIServiceListDevicesV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this act device Api service list devices v1 o k response has a 5xx status code
func (o *ActDeviceAPIServiceListDevicesV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this act device Api service list devices v1 o k response a status code equal to that given
func (o *ActDeviceAPIServiceListDevicesV1OK) IsCode(code int) bool {
	return code == 200
}

func (o *ActDeviceAPIServiceListDevicesV1OK) Error() string {
	return fmt.Sprintf("[GET /api/v1/devices][%d] actDeviceApiServiceListDevicesV1OK  %+v", 200, o.Payload)
}

func (o *ActDeviceAPIServiceListDevicesV1OK) String() string {
	return fmt.Sprintf("[GET /api/v1/devices][%d] actDeviceApiServiceListDevicesV1OK  %+v", 200, o.Payload)
}

func (o *ActDeviceAPIServiceListDevicesV1OK) GetPayload() *models.V1ListDevicesV1Response {
	return o.Payload
}

func (o *ActDeviceAPIServiceListDevicesV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ListDevicesV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActDeviceAPIServiceListDevicesV1Default creates a ActDeviceAPIServiceListDevicesV1Default with default headers values
func NewActDeviceAPIServiceListDevicesV1Default(code int) *ActDeviceAPIServiceListDevicesV1Default {
	return &ActDeviceAPIServiceListDevicesV1Default{
		_statusCode: code,
	}
}

/*
ActDeviceAPIServiceListDevicesV1Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ActDeviceAPIServiceListDevicesV1Default struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the act device Api service list devices v1 default response
func (o *ActDeviceAPIServiceListDevicesV1Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this act device Api service list devices v1 default response has a 2xx status code
func (o *ActDeviceAPIServiceListDevicesV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this act device Api service list devices v1 default response has a 3xx status code
func (o *ActDeviceAPIServiceListDevicesV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this act device Api service list devices v1 default response has a 4xx status code
func (o *ActDeviceAPIServiceListDevicesV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this act device Api service list devices v1 default response has a 5xx status code
func (o *ActDeviceAPIServiceListDevicesV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this act device Api service list devices v1 default response a status code equal to that given
func (o *ActDeviceAPIServiceListDevicesV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ActDeviceAPIServiceListDevicesV1Default) Error() string {
	return fmt.Sprintf("[GET /api/v1/devices][%d] ActDeviceApiService_ListDevicesV1 default  %+v", o._statusCode, o.Payload)
}

func (o *ActDeviceAPIServiceListDevicesV1Default) String() string {
	return fmt.Sprintf("[GET /api/v1/devices][%d] ActDeviceApiService_ListDevicesV1 default  %+v", o._statusCode, o.Payload)
}

func (o *ActDeviceAPIServiceListDevicesV1Default) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ActDeviceAPIServiceListDevicesV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
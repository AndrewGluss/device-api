//go:build http_test || allure_test

package http_test

import (
	"context"
	"fmt"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_clients/generated_client/http_client"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_clients/generated_client/http_client/act_device_api_service"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_clients/generated_client/models"
	"testing"
)

func Test_HTTP_Create_Device(t *testing.T) {
	runner.Run(t, "Http Create Device", func(t provider.T) {
		t.Title("Create Device with http request")
		t.Description("Test act_device_api. Create new device")
		cfg := http_client.DefaultTransportConfig()
		cfg.Host = "localhost:8080"
		var client *http_client.DeviceAPI
		var newDevice *models.V1CreateDeviceV1Request
		ctx := context.Background()
		var respCreate *act_device_api_service.ActDeviceAPIServiceCreateDeviceV1OK
		var respDescribe *act_device_api_service.ActDeviceAPIServiceDescribeDeviceV1OK
		t.WithNewStep("Prepare client", func(sCtx provider.StepCtx) {
			client = http_client.NewHTTPClientWithConfig(nil, cfg)
			newDevice = &models.V1CreateDeviceV1Request{
				Platform: "Android",
				UserID:   "777",
			}
		})
		t.WithNewStep("Create Device", func(sCtx provider.StepCtx) {
			// Act
			respCreate, _ = client.ActDeviceAPIService.ActDeviceAPIServiceCreateDeviceV1(
				&act_device_api_service.ActDeviceAPIServiceCreateDeviceV1Params{
					Body:    newDevice,
					Context: ctx,
				},
			)
			sCtx.WithNewAttachment("Request", allure.Text, []byte(fmt.Sprintf("%+v", respCreate)))
			describeID := respCreate.Payload.DeviceID
			respDescribe, _ = client.ActDeviceAPIService.ActDeviceAPIServiceDescribeDeviceV1(
				&act_device_api_service.ActDeviceAPIServiceDescribeDeviceV1Params{
					DeviceID: describeID,
					Context:  ctx,
				},
			)
			sCtx.WithNewAttachment("Response", allure.Text, []byte(fmt.Sprintf("%+v", respDescribe.Payload.Value)))
			sCtx.Require().Equal(newDevice.Platform, respDescribe.Payload.Value.Platform)
			sCtx.Require().Equal(newDevice.UserID, respDescribe.Payload.Value.UserID)
		})
	})
}

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
	"testing"
)

func Test_HTTP_Update_Device(t *testing.T) {
	runner.Run(t, "Http Update Device", func(t provider.T) {
		t.Title("Update Device with http request")
		t.Description("Test act_device_api. update existed device")
		cfg := http_client.DefaultTransportConfig()
		cfg.Host = "localhost:8080"
		var client *http_client.DeviceAPI

		var respUpdate *act_device_api_service.ActDeviceAPIServiceUpdateDeviceV1OK
		var respDescribe *act_device_api_service.ActDeviceAPIServiceDescribeDeviceV1OK
		t.WithNewStep("Prepare client", func(sCtx provider.StepCtx) {

			client = http_client.NewHTTPClientWithConfig(nil, cfg)
		})

		t.WithNewStep("Request to update", func(sCtx provider.StepCtx) {
			updateDeviceID := "7"
			updateBody := act_device_api_service.ActDeviceAPIServiceUpdateDeviceV1Body{
				Platform: "Linux",
				UserID:   "5553535",
			}
			ctx := context.Background()

			respUpdate, _ = client.ActDeviceAPIService.ActDeviceAPIServiceUpdateDeviceV1(
				&act_device_api_service.ActDeviceAPIServiceUpdateDeviceV1Params{
					DeviceID: updateDeviceID,
					Body:     updateBody,
					Context:  ctx,
				})
			sCtx.WithNewAttachment("Request", allure.Text, []byte(fmt.Sprintf("%+v", respUpdate)))

			describeID := updateDeviceID
			respDescribe, _ = client.ActDeviceAPIService.ActDeviceAPIServiceDescribeDeviceV1(
				&act_device_api_service.ActDeviceAPIServiceDescribeDeviceV1Params{
					DeviceID: describeID,
					Context:  context.Background(),
				},
			)
			sCtx.WithNewAttachment("Response", allure.Text, []byte(fmt.Sprintf("%+v", respDescribe.Payload.Value)))
			sCtx.Require().Equal(updateBody.Platform, respDescribe.Payload.Value.Platform)
			sCtx.Require().Equal(updateBody.UserID, respDescribe.Payload.Value.UserID)
			sCtx.Require().Equal(respUpdate.Payload.Success, true)
		})
	})
}

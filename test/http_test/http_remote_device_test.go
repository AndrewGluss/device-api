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

func Test_HTTP_Remove_Device(t *testing.T) {
	runner.Run(t, "Http Remove Device", func(t provider.T) {
		t.Title("Remove Device with http request")
		t.Description("Test act_device_api. remove existed device")
		cfg := http_client.DefaultTransportConfig()
		cfg.Host = "localhost:8080"
		var client *http_client.DeviceAPI
		ctx := context.Background()
		var respRemote *act_device_api_service.ActDeviceAPIServiceRemoveDeviceV1OK
		var respDescribe *act_device_api_service.ActDeviceAPIServiceDescribeDeviceV1OK
		t.WithNewStep("Prepare client", func(sCtx provider.StepCtx) {

			client = http_client.NewHTTPClientWithConfig(nil, cfg)
		})
		t.WithNewStep("Request", func(sCtx provider.StepCtx) {
			remoteDevice := "2"
			// Act
			respRemote, _ = client.ActDeviceAPIService.ActDeviceAPIServiceRemoveDeviceV1(
				&act_device_api_service.ActDeviceAPIServiceRemoveDeviceV1Params{
					DeviceID: remoteDevice,
					Context:  ctx,
				},
			)
			sCtx.WithNewAttachment("Request", allure.Text, []byte(fmt.Sprintf("%+v", respRemote)))

			describeID := remoteDevice
			respDescribe, _ = client.ActDeviceAPIService.ActDeviceAPIServiceDescribeDeviceV1(
				&act_device_api_service.ActDeviceAPIServiceDescribeDeviceV1Params{
					DeviceID: describeID,
					Context:  context.Background(),
				},
			)
			sCtx.Require().Equal(respRemote.Payload.Found, true)
			sCtx.Require().Equal(respDescribe.IsServerError(), false)
		})
	})
}

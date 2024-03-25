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

func Test_HTTP_Describe_Device(t *testing.T) {
	runner.Run(t, "Http Describe Device", func(t provider.T) {
		t.Title("Create Device with http request")
		t.Description("Test act_device_api. Create new device")
		cfg := http_client.DefaultTransportConfig()
		cfg.Host = "localhost:8080"
		var client *http_client.DeviceAPI
		var respDescribe *act_device_api_service.ActDeviceAPIServiceDescribeDeviceV1OK
		var describeID string
		ctx := context.Background()
		t.WithNewStep("Prepare client", func(sCtx provider.StepCtx) {

			client = http_client.NewHTTPClientWithConfig(nil, cfg)

		})
		t.WithNewStep("Describe Device", func(sCtx provider.StepCtx) {
			describeID = "10"

			// Act

			respDescribe, _ = client.ActDeviceAPIService.ActDeviceAPIServiceDescribeDeviceV1(
				&act_device_api_service.ActDeviceAPIServiceDescribeDeviceV1Params{
					DeviceID: describeID,
					Context:  ctx,
				},
			)
			sCtx.WithNewAttachment("Response", allure.Text, []byte(fmt.Sprintf("%+v", respDescribe.Payload.Value)))

			t.Require().Equal(respDescribe.Payload.Value.ID, describeID)
		})
	})
}

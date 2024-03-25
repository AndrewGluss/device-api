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

func Test_HTTP_List_Devices(t *testing.T) {
	runner.Run(t, "Http List Devices", func(t provider.T) {
		t.Title("List Devices with http request")
		t.Description("Test act_device_api. List devices")
		cfg := http_client.DefaultTransportConfig()
		cfg.Host = "localhost:8080"
		var client *http_client.DeviceAPI
		ctx := context.Background()
		var respList *act_device_api_service.ActDeviceAPIServiceListDevicesV1OK
		var err error
		t.WithNewStep("Prepare to request", func(sCtx provider.StepCtx) {

			client = http_client.NewHTTPClientWithConfig(nil, cfg)
		})
		// Act
		t.WithNewStep("List Devices request", func(sCtx provider.StepCtx) {
			page := "1"
			perPage := "10"
			respList, err = client.ActDeviceAPIService.ActDeviceAPIServiceListDevicesV1(
				&act_device_api_service.ActDeviceAPIServiceListDevicesV1Params{
					Page:    &page,
					PerPage: &perPage,
					Context: ctx,
				},
			)
			sCtx.WithNewAttachment("Response", allure.Text, []byte(fmt.Sprintf("%+v", respList.Payload.Items)))
			sCtx.Require().NoError(err)
			sCtx.Require().NotEmpty(respList)
		})

	})
}

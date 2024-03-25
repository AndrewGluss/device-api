//go:build grpc_test || allure_test

package grpc_test

import (
	"context"
	"fmt"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	act_device_api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"
	"google.golang.org/grpc"
	"testing"
)

func Test_gRPC_List_Devices(t *testing.T) {
	runner.Run(t, "Grpc List Devices", func(t provider.T) {
		t.Title("List Devices with http request")
		t.Description("Test act_device_api. List devices")
		var conn *grpc.ClientConn
		var err error
		var deviceAPIClient act_device_api.ActDeviceApiServiceClient
		var response *act_device_api.ListDevicesV1Response
		ctx := context.Background()

		t.WithNewStep("Prepare client", func(sCtx provider.StepCtx) {
			const host string = "localhost:8082"

			conn, err = grpc.Dial(host, grpc.WithInsecure())

			if err != nil {
				t.Fatal("Error to connect")
			}
			if err != nil {
				t.Fatal("Error to connect")
			}
		})

		t.WithNewStep("Request to list devices", func(sCtx provider.StepCtx) {
			// Arrange

			deviceAPIClient = act_device_api.NewActDeviceApiServiceClient(conn)
			request := &act_device_api.ListDevicesV1Request{
				Page:    uint64(1),
				PerPage: uint64(10),
			}
			sCtx.WithNewAttachment("Request", allure.Text, []byte(fmt.Sprintf("%+v", request)))
			//Act

			response, _ = deviceAPIClient.ListDevicesV1(ctx, request)
			sCtx.WithNewAttachment("Response", allure.Text, []byte(fmt.Sprintf("%+v", response)))
			//Assert

			t.Require().NotEqual(len(response.Items), 0)
		})
	})
}

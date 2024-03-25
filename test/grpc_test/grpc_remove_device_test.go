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

func Test_gRPC_Remove_Device(t *testing.T) {
	runner.Run(t, "Grpc Remove Device", func(t provider.T) {
		t.Title("Remove Device with http request")
		t.Description("Test act_device_api. remove existed device")
		var conn *grpc.ClientConn
		var err error
		var deviceAPIClient act_device_api.ActDeviceApiServiceClient
		var response *act_device_api.RemoveDeviceV1Response
		var respDescribe *act_device_api.DescribeDeviceV1Response
		ctx := context.Background()
		t.WithNewStep("Prepare client", func(sCtx provider.StepCtx) {
			const host string = "localhost:8082"

			conn, err = grpc.Dial(host, grpc.WithInsecure())

			if err != nil {
				t.Fatal("Error to connect")
			}
		})

		// Arrange
		t.WithNewStep("Request to remove", func(sCtx provider.StepCtx) {
			deviceAPIClient = act_device_api.NewActDeviceApiServiceClient(conn)
			request := &act_device_api.RemoveDeviceV1Request{
				DeviceId: uint64(1),
			}
			sCtx.WithNewAttachment("Request", allure.Text, []byte(fmt.Sprintf("%+v", request)))
			//Act

			response, _ = deviceAPIClient.RemoveDeviceV1(ctx, request)
			deviceID := request.DeviceId
			requestDescribe := &act_device_api.DescribeDeviceV1Request{
				DeviceId: deviceID,
			}
			sCtx.WithNewAttachment("Response", allure.Text, []byte(fmt.Sprintf("%+v", response)))
			respDescribe, _ = deviceAPIClient.DescribeDeviceV1(ctx, requestDescribe)

			//Assert

			sCtx.Require().Equal(response.Found, true)
			sCtx.Require().Equal(respDescribe.Validate(), nil)
		})
	})
}

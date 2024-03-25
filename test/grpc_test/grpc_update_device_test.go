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

func Test_gRPC_Update_Device(t *testing.T) {
	runner.Run(t, "Grpc Update Device", func(t provider.T) {
		t.Title("Update Device with grpc request")
		t.Description("Test act_device_api. update existed device")
		var conn *grpc.ClientConn
		var err error
		var deviceAPIClient act_device_api.ActDeviceApiServiceClient
		var request *act_device_api.UpdateDeviceV1Request
		var response *act_device_api.UpdateDeviceV1Response
		var respDescribe *act_device_api.DescribeDeviceV1Response
		var deviceID uint64
		ctx := context.Background()

		t.WithNewStep("Prepare client", func(sCtx provider.StepCtx) {
			const host string = "localhost:8082"

			conn, err = grpc.Dial(host, grpc.WithInsecure())

			if err != nil {
				t.Fatal("Error to connect")
			}
		})
		t.WithNewStep("Request to update", func(sCtx provider.StepCtx) {
			deviceAPIClient = act_device_api.NewActDeviceApiServiceClient(conn)
			request = &act_device_api.UpdateDeviceV1Request{
				DeviceId: uint64(8),
				Platform: "Android",
				UserId:   uint64(696),
			}
			sCtx.WithNewAttachment("Request", allure.Text, []byte(fmt.Sprintf("%+v", request)))
			//Act

			response, err = deviceAPIClient.UpdateDeviceV1(ctx, request)

			sCtx.WithNewAttachment("Response", allure.Text, []byte(fmt.Sprintf("%+v", response)))

			deviceID = request.DeviceId
			requestDescribe := &act_device_api.DescribeDeviceV1Request{
				DeviceId: deviceID,
			}
			respDescribe, err = deviceAPIClient.DescribeDeviceV1(ctx, requestDescribe)
		})

		//Assert
		t.Require().Equal(response.Success, true)
		t.Require().Equal(request.Platform, respDescribe.Value.Platform)
		t.Require().Equal(request.UserId, respDescribe.Value.UserId)
		t.Require().Equal(deviceID, respDescribe.Value.Id)
	})
}

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

func Test_gRPC_Create_Device(t *testing.T) {
	runner.Run(t, "Grpc Create Device", func(t provider.T) {
		t.Title("Create Device with http request")
		t.Description("Test act_device_api. Create new device")
		var conn *grpc.ClientConn
		var err error
		var deviceAPIClient act_device_api.ActDeviceApiServiceClient
		var request *act_device_api.CreateDeviceV1Request
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

		t.WithNewStep("Request to create", func(sCtx provider.StepCtx) {
			deviceAPIClient = act_device_api.NewActDeviceApiServiceClient(conn)
			request = &act_device_api.CreateDeviceV1Request{
				Platform: "iOS",
				UserId:   uint64(666),
			}
			sCtx.WithNewAttachment("Request", allure.Text, []byte(fmt.Sprintf("%+v", request)))
			//Act

			response, _ := deviceAPIClient.CreateDeviceV1(ctx, request)
			deviceID = response.DeviceId
			requestDescribe := &act_device_api.DescribeDeviceV1Request{
				DeviceId: deviceID,
			}
			sCtx.WithNewAttachment("Response", allure.Text, []byte(fmt.Sprintf("%+v", response)))
			respDescribe, err = deviceAPIClient.DescribeDeviceV1(ctx, requestDescribe)

			//Assert
			sCtx.Require().Equal(request.Platform, respDescribe.Value.Platform)
			sCtx.Require().Equal(request.UserId, respDescribe.Value.UserId)
			sCtx.Require().Equal(deviceID, respDescribe.Value.Id)
		})

	})
}

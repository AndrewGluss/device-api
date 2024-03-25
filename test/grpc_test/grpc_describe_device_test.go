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

func Test_gRPC_Describe_Device(t *testing.T) {
	runner.Run(t, "Grpc Describe Device", func(t provider.T) {
		t.Title("Create Device with http request")
		t.Description("Test act_device_api. Create new device")
		var conn *grpc.ClientConn
		var err error
		var deviceAPIClient act_device_api.ActDeviceApiServiceClient
		var request *act_device_api.DescribeDeviceV1Request
		var response *act_device_api.DescribeDeviceV1Response
		ctx := context.Background()
		t.WithNewStep("Prepare client", func(sCtx provider.StepCtx) {
			const host string = "localhost:8082"

			conn, err = grpc.Dial(host, grpc.WithInsecure())

			if err != nil {
				t.Fatal("Error to connect")
			}
		})
		t.WithNewStep("Request to describe device", func(sCtx provider.StepCtx) {
			deviceAPIClient = act_device_api.NewActDeviceApiServiceClient(conn)
			request = &act_device_api.DescribeDeviceV1Request{
				DeviceId: uint64(9),
			}
			sCtx.WithNewAttachment("Request", allure.Text, []byte(fmt.Sprintf("%+v", request)))
			//Act

			response, err = deviceAPIClient.DescribeDeviceV1(ctx, request)
			sCtx.WithNewAttachment("Response", allure.Text, []byte(fmt.Sprintf("%+v", response)))
			//Assert

			sCtx.Require().Equal(response.Value.Id, request.DeviceId)
		})
	})
}

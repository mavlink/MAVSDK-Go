package telemetry

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/grpc"
)

//Service interface for intertelemetry with telemetry interface of drone
type Service interface {
	getResult() TelemetryResult_Result
	InitTelemetry(*grpc.ClientConn) TelemetryServiceClient
	Position() (*PositionResponse, error)
}

//ServiceImpl creates client and implements telemetryService
type ServiceImpl struct {
	Name             string
	Client           *grpc.ClientConn
	telemetryService TelemetryServiceClient
	positionResponse PositionResponse
}

func (a *ServiceImpl) translateFromGRPC(response *TelemetryResult) TelemetryResult_Result {
	return response.Result
}

func (a *ServiceImpl) translateToGRPC() *Position {
	return a.positionResponse.Position
}

func (a *ServiceImpl) getResult(response *TelemetryResult) TelemetryResult_Result {
	return a.translateFromGRPC(response)
}

//InitTelemetry initializes telemetry client
func (a *ServiceImpl) InitTelemetry() {
	a.telemetryService = NewTelemetryServiceClient(a.Client)
}

//Position acts as a wrapper for Position request from telemetry grpc package
func (a *ServiceImpl) Position() {
	request := &SubscribePositionRequest{}
	ctx := context.Background()
	stream, err := a.telemetryService.SubscribePosition(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}
	for {
		m := &PositionResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m.String())
	}
	// return response, err
}

//SetRatePosition acts as a wrapper for Position request from telemetry grpc package
func (a *ServiceImpl) SetRatePosition(rateHz float64) {
	request := &SetRatePositionRequest{}
	request.RateHz = rateHz
	ctx := context.Background()
	response, err := a.telemetryService.SetRatePosition(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}
	result := a.getResult(response.TelemetryResult)
	if result != TelemetryResult_RESULT_SUCCESS {
		fmt.Println("Error while setting rate position")
	}

}

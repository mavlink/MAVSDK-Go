package sample

import (
	"context"
	"fmt"
	"io"

	"github.com/ykhedar/MAVSDK-Go/Sources/telemetry"
	"google.golang.org/grpc"
)

//TelemetryService interface for intertelemetry with telemetry interface of drone
type TelemetryService interface {
	getResult() telemetry.TelemetryResult_Result
	InitTelemetry(*grpc.ClientConn) telemetry.TelemetryServiceClient
	Position() (*telemetry.PositionResponse, error)
}

//TelemetryServiceImpl creates client and implements telemetryService
type TelemetryServiceImpl struct {
	Name             string
	Client           *grpc.ClientConn
	telemetryService telemetry.TelemetryServiceClient
	positionResponse telemetry.PositionResponse
}

func (a *TelemetryServiceImpl) translateFromGRPC(response *telemetry.PositionResponse) *telemetry.Position {
	return response.Position
}

func (a *TelemetryServiceImpl) translateToGRPC() *telemetry.Position {
	return a.positionResponse.Position
}

func (a *TelemetryServiceImpl) getResult(response *telemetry.PositionResponse) *telemetry.Position {
	return a.translateFromGRPC(response)
}

//InitTelemetry initializes telemetry client
func (a *TelemetryServiceImpl) InitTelemetry() {
	a.telemetryService = telemetry.NewTelemetryServiceClient(a.Client)
}

//Position acts as a wrapper for Position request from telemetry grpc package
func (a *TelemetryServiceImpl) Position() {
	request := &telemetry.SubscribePositionRequest{}
	ctx := context.Background()
	stream, err := a.telemetryService.SubscribePosition(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}
	for {
		m := &telemetry.PositionResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", a.getResult(m).String())
	}
	// return response, err
}

// package telemery

import (
	"context"
	"fmt"
	"io"
)

//Service interface for intertelemetry with telemetry interface of drone
type Service interface {
	Result() TelemetryResult_Result
}

//ServiceImpl creates client and implements telemetryService
type ServiceImpl struct {
	Client TelemetryServiceClient
}

//SubscribePosition acts as a wrapper for Position request from telemetry grpc package
func (s *ServiceImpl) SubscribePosition() <-chan string {
	request := &SubscribePositionRequest{}
	ctx := context.Background()
	stream, err := s.Client.SubscribePosition(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
	}
	ch := make(chan string)
	go func() {
		defer close(ch)
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
			m.GetPosition()
			fmt.Printf("message %v\n", m.String())
			// ch <- m.String()
		}
	}()
	return ch
}

//SetRatePosition acts as a wrapper for Position request from telemetry grpc package
func (s *ServiceImpl) SetRatePosition(rateHz float64) {
	request := &SetRatePositionRequest{}
	request.RateHz = rateHz
	ctx := context.Background()

	response, err := s.Client.SetRatePosition(ctx, request)

	if err != nil {
		fmt.Printf("Unable to SetRatePosition %v\n", err)
	}
	result := response.TelemetryResult
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Println("Error while setting rate position")
	}
	fmt.Printf("Position set\n")
}

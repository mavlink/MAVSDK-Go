package transponder

import (
	"context"
	"fmt"
	"io"

	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client TransponderServiceClient
}

/*
   Subscribe to 'transponder' updates.


*/

func (a *ServiceImpl) Transponder(ctx context.Context) (<-chan *AdsbVehicle, error) {
	ch := make(chan *AdsbVehicle)
	request := &SubscribeTransponderRequest{}
	stream, err := a.Client.SubscribeTransponder(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				m := &TransponderResponse{}
				err := stream.RecvMsg(m)
				if err == io.EOF {
					return
				}
				if err != nil {
					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
						return
					}
					fmt.Printf("Unable to receive message: %v\n", err)
					break
				}
				ch <- m.GetTransponder()
			}
		}
	}()
	return ch, nil
}

/*
   Set rate to 'transponder' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateTransponder(ctx context.Context, rateHz float64) (*SetRateTransponderResponse, error) {

	request := &SetRateTransponderRequest{}
	request.RateHz = rateHz
	response, err := s.Client.SetRateTransponder(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

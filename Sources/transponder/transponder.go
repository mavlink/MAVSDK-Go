package transponder

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc/status"
	codes "google.golang.org/grpc/codes"
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
			m := &TransponderResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Transponder messages, err: %v", err)
				break
			}
			ch <- m.GetTransponder()
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

package winch

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc/status"
	codes "google.golang.org/grpc/codes"
)

type ServiceImpl struct {
	Client WinchServiceClient
}

/*
   Subscribe to 'winch status' updates.


*/

func (a *ServiceImpl) Status(ctx context.Context) (<-chan *Status, error) {
	ch := make(chan *Status)
	request := &SubscribeStatusRequest{}
	stream, err := a.Client.SubscribeStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &StatusResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Status messages, err: %v", err)
				break
			}
			ch <- m.GetStatus()
		}
	}()
	return ch, nil
}

/*
   Allow motor to freewheel.

   Parameters
   ----------
   instance uint32


*/

func (s *ServiceImpl) Relax(ctx context.Context, instance uint32) (*RelaxResponse, error) {

	request := &RelaxRequest{}
	request.Instance = instance
	response, err := s.Client.Relax(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Wind or unwind specified length of line, optionally using specified rate.

   Parameters
   ----------
   instance uint32

   lengthM float32

   rateMS float32


*/

func (s *ServiceImpl) RelativeLengthControl(ctx context.Context, instance uint32, lengthM float32, rateMS float32) (*RelativeLengthControlResponse, error) {

	request := &RelativeLengthControlRequest{}
	request.Instance = instance
	request.LengthM = lengthM
	request.RateMS = rateMS
	response, err := s.Client.RelativeLengthControl(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Wind or unwind line at specified rate.

   Parameters
   ----------
   instance uint32

   rateMS float32


*/

func (s *ServiceImpl) RateControl(ctx context.Context, instance uint32, rateMS float32) (*RateControlResponse, error) {

	request := &RateControlRequest{}
	request.Instance = instance
	request.RateMS = rateMS
	response, err := s.Client.RateControl(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Perform the locking sequence to relieve motor while in the fully retracted position.

   Parameters
   ----------
   instance uint32


*/

func (s *ServiceImpl) Lock(ctx context.Context, instance uint32) (*LockResponse, error) {

	request := &LockRequest{}
	request.Instance = instance
	response, err := s.Client.Lock(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Sequence of drop, slow down, touch down, reel up, lock.

   Parameters
   ----------
   instance uint32


*/

func (s *ServiceImpl) Deliver(ctx context.Context, instance uint32) (*DeliverResponse, error) {

	request := &DeliverRequest{}
	request.Instance = instance
	response, err := s.Client.Deliver(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Engage motor and hold current position.

   Parameters
   ----------
   instance uint32


*/

func (s *ServiceImpl) Hold(ctx context.Context, instance uint32) (*HoldResponse, error) {

	request := &HoldRequest{}
	request.Instance = instance
	response, err := s.Client.Hold(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Return the reel to the fully retracted position.

   Parameters
   ----------
   instance uint32


*/

func (s *ServiceImpl) Retract(ctx context.Context, instance uint32) (*RetractResponse, error) {

	request := &RetractRequest{}
	request.Instance = instance
	response, err := s.Client.Retract(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Load the reel with line.

   The winch will calculate the total loaded length and stop when the tension exceeds a threshold.

   Parameters
   ----------
   instance uint32


*/

func (s *ServiceImpl) LoadLine(ctx context.Context, instance uint32) (*LoadLineResponse, error) {

	request := &LoadLineRequest{}
	request.Instance = instance
	response, err := s.Client.LoadLine(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Spool out the entire length of the line.

   Parameters
   ----------
   instance uint32


*/

func (s *ServiceImpl) AbandonLine(ctx context.Context, instance uint32) (*AbandonLineResponse, error) {

	request := &AbandonLineRequest{}
	request.Instance = instance
	response, err := s.Client.AbandonLine(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Spools out just enough to present the hook to the user to load the payload.

   Parameters
   ----------
   instance uint32


*/

func (s *ServiceImpl) LoadPayload(ctx context.Context, instance uint32) (*LoadPayloadResponse, error) {

	request := &LoadPayloadRequest{}
	request.Instance = instance
	response, err := s.Client.LoadPayload(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

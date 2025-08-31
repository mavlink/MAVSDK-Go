package winch

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client WinchServiceClient
}

/*
Status Subscribe to 'winch status' updates.
*/
func (a *ServiceImpl) Status(
	ctx context.Context,

) (<-chan *Status, error) {
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
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive Status messages, err: %v", err)
			}
			ch <- m.GetStatus()
		}
	}()
	return ch, nil
}

/*
Relax Allow motor to freewheel.
*/
func (s *ServiceImpl) Relax(
	ctx context.Context,
	instance uint32,

) (*RelaxResponse, error) {
	request := &RelaxRequest{
		Instance: instance,
	}
	response, err := s.Client.Relax(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
RelativeLengthControl Wind or unwind specified length of line, optionally using specified rate.
*/
func (s *ServiceImpl) RelativeLengthControl(
	ctx context.Context,
	instance uint32,
	lengthM float32,
	rateMS float32,

) (*RelativeLengthControlResponse, error) {
	request := &RelativeLengthControlRequest{
		Instance: instance,
		LengthM:  lengthM,
		RateMS:   rateMS,
	}
	response, err := s.Client.RelativeLengthControl(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
RateControl Wind or unwind line at specified rate.
*/
func (s *ServiceImpl) RateControl(
	ctx context.Context,
	instance uint32,
	rateMS float32,

) (*RateControlResponse, error) {
	request := &RateControlRequest{
		Instance: instance,
		RateMS:   rateMS,
	}
	response, err := s.Client.RateControl(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
Lock Perform the locking sequence to relieve motor while in the fully retracted position.
*/
func (s *ServiceImpl) Lock(
	ctx context.Context,
	instance uint32,

) (*LockResponse, error) {
	request := &LockRequest{
		Instance: instance,
	}
	response, err := s.Client.Lock(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
Deliver Sequence of drop, slow down, touch down, reel up, lock.
*/
func (s *ServiceImpl) Deliver(
	ctx context.Context,
	instance uint32,

) (*DeliverResponse, error) {
	request := &DeliverRequest{
		Instance: instance,
	}
	response, err := s.Client.Deliver(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
Hold Engage motor and hold current position.
*/
func (s *ServiceImpl) Hold(
	ctx context.Context,
	instance uint32,

) (*HoldResponse, error) {
	request := &HoldRequest{
		Instance: instance,
	}
	response, err := s.Client.Hold(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
Retract Return the reel to the fully retracted position.
*/
func (s *ServiceImpl) Retract(
	ctx context.Context,
	instance uint32,

) (*RetractResponse, error) {
	request := &RetractRequest{
		Instance: instance,
	}
	response, err := s.Client.Retract(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
LoadLine Load the reel with line.

	The winch will calculate the total loaded length and stop when the tension exceeds a threshold.
*/
func (s *ServiceImpl) LoadLine(
	ctx context.Context,
	instance uint32,

) (*LoadLineResponse, error) {
	request := &LoadLineRequest{
		Instance: instance,
	}
	response, err := s.Client.LoadLine(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
AbandonLine Spool out the entire length of the line.
*/
func (s *ServiceImpl) AbandonLine(
	ctx context.Context,
	instance uint32,

) (*AbandonLineResponse, error) {
	request := &AbandonLineRequest{
		Instance: instance,
	}
	response, err := s.Client.AbandonLine(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
LoadPayload Spools out just enough to present the hook to the user to load the payload.
*/
func (s *ServiceImpl) LoadPayload(
	ctx context.Context,
	instance uint32,

) (*LoadPayloadResponse, error) {
	request := &LoadPayloadRequest{
		Instance: instance,
	}
	response, err := s.Client.LoadPayload(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

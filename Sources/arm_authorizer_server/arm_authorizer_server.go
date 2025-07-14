package arm_authorizer_server

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client ArmAuthorizerServerServiceClient
}

/*
ArmAuthorization Subscribe to arm authorization request messages. Each request received should respond to using RespondArmAuthorization
*/
func (a *ServiceImpl) ArmAuthorization(
	ctx context.Context,

) (<-chan uint32, error) {
	ch := make(chan uint32)
	request := &SubscribeArmAuthorizationRequest{}
	stream, err := a.Client.SubscribeArmAuthorization(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ArmAuthorizationResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive ArmAuthorization messages, err: %v", err)
			}
			ch <- m.GetSystemId()
		}
	}()
	return ch, nil
}

/*
AcceptArmAuthorization Authorize arm for the specific time
*/
func (s *ServiceImpl) AcceptArmAuthorization(
	ctx context.Context,
	validTimeS int32,

) (*AcceptArmAuthorizationResponse, error) {
	request := &AcceptArmAuthorizationRequest{
		ValidTimeS: validTimeS,
	}
	response, err := s.Client.AcceptArmAuthorization(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
RejectArmAuthorization Reject arm authorization request
*/
func (s *ServiceImpl) RejectArmAuthorization(
	ctx context.Context,
	temporarily bool,
	reason *RejectionReason,

	extraInfo int32,

) (*RejectArmAuthorizationResponse, error) {
	request := &RejectArmAuthorizationRequest{
		Temporarily: temporarily,
		Reason:      *reason,
		ExtraInfo:   extraInfo,
	}
	response, err := s.Client.RejectArmAuthorization(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

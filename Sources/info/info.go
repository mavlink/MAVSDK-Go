package info

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client InfoServiceClient
}

/*
GetFlightInformation Get flight information of the system.
*/
func (s *ServiceImpl) GetFlightInformation(
	ctx context.Context,

) (*GetFlightInformationResponse, error) {
	request := &GetFlightInformationRequest{}
	response, err := s.Client.GetFlightInformation(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
GetIdentification Get the identification of the system.
*/
func (s *ServiceImpl) GetIdentification(
	ctx context.Context,

) (*GetIdentificationResponse, error) {
	request := &GetIdentificationRequest{}
	response, err := s.Client.GetIdentification(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
GetProduct Get product information of the system.
*/
func (s *ServiceImpl) GetProduct(
	ctx context.Context,

) (*GetProductResponse, error) {
	request := &GetProductRequest{}
	response, err := s.Client.GetProduct(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
GetVersion Get the version information of the system.
*/
func (s *ServiceImpl) GetVersion(
	ctx context.Context,

) (*GetVersionResponse, error) {
	request := &GetVersionRequest{}
	response, err := s.Client.GetVersion(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
GetSpeedFactor Get the speed factor of a simulation (with lockstep a simulation can run faster or slower than realtime).
*/
func (s *ServiceImpl) GetSpeedFactor(
	ctx context.Context,

) (*GetSpeedFactorResponse, error) {
	request := &GetSpeedFactorRequest{}
	response, err := s.Client.GetSpeedFactor(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
FlightInformation Subscribe to 'flight information' updates.
*/
func (a *ServiceImpl) FlightInformation(
	ctx context.Context,

) (<-chan *FlightInfo, error) {
	ch := make(chan *FlightInfo)
	request := &SubscribeFlightInformationRequest{}
	stream, err := a.Client.SubscribeFlightInformation(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &FlightInformationResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive FlightInformation messages, err: %v", err)
			}
			ch <- m.GetFlightInfo()
		}
	}()
	return ch, nil
}

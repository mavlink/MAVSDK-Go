package action_server

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client ActionServerServiceClient
}

/*
ArmDisarm Subscribe to ARM/DISARM commands
*/
func (a *ServiceImpl) ArmDisarm(
	ctx context.Context,

) (<-chan *ArmDisarm, error) {
	ch := make(chan *ArmDisarm)
	request := &SubscribeArmDisarmRequest{}
	stream, err := a.Client.SubscribeArmDisarm(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ArmDisarmResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive ArmDisarm messages, err: %v", err)
			}
			ch <- m.GetArm()
		}
	}()
	return ch, nil
}

/*
FlightModeChange Subscribe to DO_SET_MODE
*/
func (a *ServiceImpl) FlightModeChange(
	ctx context.Context,

) (<-chan FlightMode, error) {
	ch := make(chan FlightMode)
	request := &SubscribeFlightModeChangeRequest{}
	stream, err := a.Client.SubscribeFlightModeChange(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &FlightModeChangeResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive FlightModeChange messages, err: %v", err)
			}
			ch <- m.GetFlightMode()
		}
	}()
	return ch, nil
}

/*
Takeoff Subscribe to takeoff command
*/
func (a *ServiceImpl) Takeoff(
	ctx context.Context,

) (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeTakeoffRequest{}
	stream, err := a.Client.SubscribeTakeoff(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &TakeoffResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Takeoff messages, err: %v", err)
			}
			ch <- m.GetTakeoff()
		}
	}()
	return ch, nil
}

/*
Land Subscribe to land command
*/
func (a *ServiceImpl) Land(
	ctx context.Context,

) (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeLandRequest{}
	stream, err := a.Client.SubscribeLand(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &LandResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Land messages, err: %v", err)
			}
			ch <- m.GetLand()
		}
	}()
	return ch, nil
}

/*
Reboot Subscribe to reboot command
*/
func (a *ServiceImpl) Reboot(
	ctx context.Context,

) (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeRebootRequest{}
	stream, err := a.Client.SubscribeReboot(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &RebootResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Reboot messages, err: %v", err)
			}
			ch <- m.GetReboot()
		}
	}()
	return ch, nil
}

/*
Shutdown Subscribe to shutdown command
*/
func (a *ServiceImpl) Shutdown(
	ctx context.Context,

) (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeShutdownRequest{}
	stream, err := a.Client.SubscribeShutdown(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ShutdownResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Shutdown messages, err: %v", err)
			}
			ch <- m.GetShutdown()
		}
	}()
	return ch, nil
}

/*
Terminate Subscribe to terminate command
*/
func (a *ServiceImpl) Terminate(
	ctx context.Context,

) (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeTerminateRequest{}
	stream, err := a.Client.SubscribeTerminate(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &TerminateResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Terminate messages, err: %v", err)
			}
			ch <- m.GetTerminate()
		}
	}()
	return ch, nil
}

/*
SetAllowTakeoff Can the vehicle takeoff
*/
func (s *ServiceImpl) SetAllowTakeoff(
	ctx context.Context,
	allowTakeoff bool,

) (*SetAllowTakeoffResponse, error) {
	request := &SetAllowTakeoffRequest{
		AllowTakeoff: allowTakeoff,
	}
	response, err := s.Client.SetAllowTakeoff(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetArmable Can the vehicle arm when requested
*/
func (s *ServiceImpl) SetArmable(
	ctx context.Context,
	armable bool,
	forceArmable bool,

) (*SetArmableResponse, error) {
	request := &SetArmableRequest{
		Armable:      armable,
		ForceArmable: forceArmable,
	}
	response, err := s.Client.SetArmable(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetDisarmable Can the vehicle disarm when requested
*/
func (s *ServiceImpl) SetDisarmable(
	ctx context.Context,
	disarmable bool,
	forceDisarmable bool,

) (*SetDisarmableResponse, error) {
	request := &SetDisarmableRequest{
		Disarmable:      disarmable,
		ForceDisarmable: forceDisarmable,
	}
	response, err := s.Client.SetDisarmable(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetAllowableFlightModes Set which modes the vehicle can transition to (Manual always allowed)
*/
func (s *ServiceImpl) SetAllowableFlightModes(
	ctx context.Context,
	flightModes *AllowableFlightModes,

) (*SetAllowableFlightModesResponse, error) {
	request := &SetAllowableFlightModesRequest{
		FlightModes: flightModes,
	}
	response, err := s.Client.SetAllowableFlightModes(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
GetAllowableFlightModes Get which modes the vehicle can transition to (Manual always allowed)
*/
func (s *ServiceImpl) GetAllowableFlightModes(
	ctx context.Context,

) (*GetAllowableFlightModesResponse, error) {
	request := &GetAllowableFlightModesRequest{}
	response, err := s.Client.GetAllowableFlightModes(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

package action_server

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct {
	Client ActionServerServiceClient
}

/*
   Subscribe to ARM/DISARM commands


*/

func (a *ServiceImpl) ArmDisarm() (<-chan *ArmDisarm, error) {
	ch := make(chan *ArmDisarm)
	request := &SubscribeArmDisarmRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetArm()
		}
	}()
	return ch, nil
}

/*
   Subscribe to DO_SET_MODE


*/

func (a *ServiceImpl) FlightModeChange() (<-chan FlightMode, error) {
	ch := make(chan FlightMode)
	request := &SubscribeFlightModeChangeRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetFlightMode()
		}
	}()
	return ch, nil
}

/*
   Subscribe to takeoff command


*/

func (a *ServiceImpl) Takeoff() (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeTakeoffRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetTakeoff()
		}
	}()
	return ch, nil
}

/*
   Subscribe to land command


*/

func (a *ServiceImpl) Land() (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeLandRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetLand()
		}
	}()
	return ch, nil
}

/*
   Subscribe to reboot command


*/

func (a *ServiceImpl) Reboot() (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeRebootRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetReboot()
		}
	}()
	return ch, nil
}

/*
   Subscribe to shutdown command


*/

func (a *ServiceImpl) Shutdown() (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeShutdownRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetShutdown()
		}
	}()
	return ch, nil
}

/*
   Subscribe to terminate command


*/

func (a *ServiceImpl) Terminate() (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeTerminateRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetTerminate()
		}
	}()
	return ch, nil
}

/*
   Can the vehicle takeoff

   Parameters
   ----------
   allowTakeoff bool


*/

func (s *ServiceImpl) SetAllowTakeoff(allowTakeoff bool) (*SetAllowTakeoffResponse, error) {

	request := &SetAllowTakeoffRequest{}
	ctx := context.Background()
	request.AllowTakeoff = allowTakeoff
	response, err := s.Client.SetAllowTakeoff(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Can the vehicle arm when requested

   Parameters
   ----------
   armable bool

   forceArmable bool


*/

func (s *ServiceImpl) SetArmable(armable bool, forceArmable bool) (*SetArmableResponse, error) {

	request := &SetArmableRequest{}
	ctx := context.Background()
	request.Armable = armable
	request.ForceArmable = forceArmable
	response, err := s.Client.SetArmable(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Can the vehicle disarm when requested

   Parameters
   ----------
   disarmable bool

   forceDisarmable bool


*/

func (s *ServiceImpl) SetDisarmable(disarmable bool, forceDisarmable bool) (*SetDisarmableResponse, error) {

	request := &SetDisarmableRequest{}
	ctx := context.Background()
	request.Disarmable = disarmable
	request.ForceDisarmable = forceDisarmable
	response, err := s.Client.SetDisarmable(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Set which modes the vehicle can transition to (Manual always allowed)

   Parameters
   ----------
   flightModes *AllowableFlightModes



*/

func (s *ServiceImpl) SetAllowableFlightModes(flightModes *AllowableFlightModes) (*SetAllowableFlightModesResponse, error) {

	request := &SetAllowableFlightModesRequest{}
	ctx := context.Background()
	request.FlightModes = flightModes

	response, err := s.Client.SetAllowableFlightModes(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Get which modes the vehicle can transition to (Manual always allowed)



   Returns
   -------
   False
   FlightModes : AllowableFlightModes


*/

func (s *ServiceImpl) GetAllowableFlightModes() (*GetAllowableFlightModesResponse, error) {
	request := &GetAllowableFlightModesRequest{}
	ctx := context.Background()
	response, err := s.Client.GetAllowableFlightModes(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

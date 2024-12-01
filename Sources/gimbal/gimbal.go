package gimbal

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client GimbalServiceClient
}

/*
SetAngles Set gimbal roll, pitch and yaw angles.

	This sets the desired roll, pitch and yaw angles of a gimbal.
	Will return when the command is accepted, however, it might
	take the gimbal longer to actually be set to the new angles.
*/
func (s *ServiceImpl) SetAngles(
	ctx context.Context,
	rollDeg float32,
	pitchDeg float32,
	yawDeg float32,

) (*SetAnglesResponse, error) {
	request := &SetAnglesRequest{
		RollDeg:  rollDeg,
		PitchDeg: pitchDeg,
		YawDeg:   yawDeg,
	}
	response, err := s.Client.SetAngles(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetPitchAndYaw Set gimbal pitch and yaw angles.

	This sets the desired pitch and yaw angles of a gimbal.
	Will return when the command is accepted, however, it might
	take the gimbal longer to actually be set to the new angles.
*/
func (s *ServiceImpl) SetPitchAndYaw(
	ctx context.Context,
	pitchDeg float32,
	yawDeg float32,

) (*SetPitchAndYawResponse, error) {
	request := &SetPitchAndYawRequest{
		PitchDeg: pitchDeg,
		YawDeg:   yawDeg,
	}
	response, err := s.Client.SetPitchAndYaw(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetPitchRateAndYawRate Set gimbal angular rates around pitch and yaw axes.

	This sets the desired angular rates around pitch and yaw axes of a gimbal.
	Will return when the command is accepted, however, it might
	take the gimbal longer to actually reach the angular rate.
*/
func (s *ServiceImpl) SetPitchRateAndYawRate(
	ctx context.Context,
	pitchRateDegS float32,
	yawRateDegS float32,

) (*SetPitchRateAndYawRateResponse, error) {
	request := &SetPitchRateAndYawRateRequest{
		PitchRateDegS: pitchRateDegS,
		YawRateDegS:   yawRateDegS,
	}
	response, err := s.Client.SetPitchRateAndYawRate(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetMode Set gimbal mode.

	This sets the desired yaw mode of a gimbal.
	Will return when the command is accepted. However, it might
	take the gimbal longer to actually be set to the new angles.
*/
func (s *ServiceImpl) SetMode(
	ctx context.Context,
	gimbalMode *GimbalMode,

) (*SetModeResponse, error) {
	request := &SetModeRequest{
		GimbalMode: *gimbalMode,
	}
	response, err := s.Client.SetMode(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetRoiLocation Set gimbal region of interest (ROI).

	This sets a region of interest that the gimbal will point to.
	The gimbal will continue to point to the specified region until it
	receives a new command.
	The function will return when the command is accepted, however, it might
	take the gimbal longer to actually rotate to the ROI.
*/
func (s *ServiceImpl) SetRoiLocation(
	ctx context.Context,
	latitudeDeg float64,
	longitudeDeg float64,
	altitudeM float32,

) (*SetRoiLocationResponse, error) {
	request := &SetRoiLocationRequest{
		LatitudeDeg:  latitudeDeg,
		LongitudeDeg: longitudeDeg,
		AltitudeM:    altitudeM,
	}
	response, err := s.Client.SetRoiLocation(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
TakeControl Take control.

	There can be only two components in control of a gimbal at any given time.
	One with "primary" control, and one with "secondary" control. The way the
	secondary control is implemented is not specified and hence depends on the
	vehicle.

	Components are expected to be cooperative, which means that they can
	override each other and should therefore do it carefully.
*/
func (s *ServiceImpl) TakeControl(
	ctx context.Context,
	controlMode *ControlMode,

) (*TakeControlResponse, error) {
	request := &TakeControlRequest{
		ControlMode: *controlMode,
	}
	response, err := s.Client.TakeControl(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ReleaseControl Release control.

	Release control, such that other components can control the gimbal.
*/
func (s *ServiceImpl) ReleaseControl(
	ctx context.Context,

) (*ReleaseControlResponse, error) {
	request := &ReleaseControlRequest{}
	response, err := s.Client.ReleaseControl(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
Control Subscribe to control status updates.

	This allows a component to know if it has primary, secondary or
	no control over the gimbal. Also, it gives the system and component ids
	of the other components in control (if any).
*/
func (a *ServiceImpl) Control(
	ctx context.Context,

) (<-chan *ControlStatus, error) {
	ch := make(chan *ControlStatus)
	request := &SubscribeControlRequest{}
	stream, err := a.Client.SubscribeControl(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ControlResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Control messages, err: %v", err)
			}
			ch <- m.GetControlStatus()
		}
	}()
	return ch, nil
}

/*
Attitude Subscribe to attitude updates.

	This gets you the gimbal's attitude and angular rate.
*/
func (a *ServiceImpl) Attitude(
	ctx context.Context,

) (<-chan *Attitude, error) {
	ch := make(chan *Attitude)
	request := &SubscribeAttitudeRequest{}
	stream, err := a.Client.SubscribeAttitude(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &AttitudeResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Attitude messages, err: %v", err)
			}
			ch <- m.GetAttitude()
		}
	}()
	return ch, nil
}

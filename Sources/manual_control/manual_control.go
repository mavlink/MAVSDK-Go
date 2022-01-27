package manual_control

import (
	"context"
)

type ServiceImpl struct {
	Client ManualControlServiceClient
}

/*
   Start position control using e.g. joystick input.

   Requires manual control input to be sent regularly already.
   Requires a valid position using e.g. GPS, external vision, or optical flow.


*/

func (s *ServiceImpl) StartPositionControl() (*StartPositionControlResponse, error) {

	request := &StartPositionControlRequest{}
	ctx := context.Background()
	response, err := s.Client.StartPositionControl(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Start altitude control

   Requires manual control input to be sent regularly already.
   Does not require a  valid position e.g. GPS.


*/

func (s *ServiceImpl) StartAltitudeControl() (*StartAltitudeControlResponse, error) {

	request := &StartAltitudeControlRequest{}
	ctx := context.Background()
	response, err := s.Client.StartAltitudeControl(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Set manual control input

   The manual control input needs to be sent at a rate high enough to prevent
   triggering of RC loss, a good minimum rate is 10 Hz.

   Parameters
   ----------
   x float32

   y float32

   z float32

   r float32


*/

func (s *ServiceImpl) SetManualControlInput(x float32, y float32, z float32, r float32) (*SetManualControlInputResponse, error) {

	request := &SetManualControlInputRequest{}
	ctx := context.Background()
	request.X = x
	request.Y = y
	request.Z = z
	request.R = r
	response, err := s.Client.SetManualControlInput(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

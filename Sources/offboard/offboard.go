package offboard

import (
	"context"
)

type ServiceImpl struct {
	Client OffboardServiceClient
}

/*
   Start offboard control.


*/

func (s *ServiceImpl) Start() (*StartResponse, error) {

	request := &StartRequest{}
	ctx := context.Background()
	response, err := s.Client.Start(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Stop offboard control.

   The vehicle will be put into Hold mode: https://docs.px4.io/en/flight_modes/hold.html


*/

func (s *ServiceImpl) Stop() (*StopResponse, error) {

	request := &StopRequest{}
	ctx := context.Background()
	response, err := s.Client.Stop(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Check if offboard control is active.

   True means that the vehicle is in offboard mode and we are actively sending
   setpoints.



   Returns
   -------
   False
   IsActive : bool
        True if offboard is active


*/

func (s *ServiceImpl) IsActive() (*IsActiveResponse, error) {
	request := &IsActiveRequest{}
	ctx := context.Background()
	response, err := s.Client.IsActive(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
   Set the attitude in terms of roll, pitch and yaw in degrees with thrust.

   Parameters
   ----------
   attitude *Attitude



*/

func (s *ServiceImpl) SetAttitude(attitude *Attitude) (*SetAttitudeResponse, error) {

	request := &SetAttitudeRequest{}
	ctx := context.Background()
	request.Attitude = attitude

	response, err := s.Client.SetAttitude(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Set direct actuator control values to groups #0 and #1.

   First 8 controls will go to control group 0, the following 8 controls to control group 1 (if
   actuator_control.num_controls more than 8).

   Parameters
   ----------
   actuatorControl *ActuatorControl



*/

func (s *ServiceImpl) SetActuatorControl(actuatorControl *ActuatorControl) (*SetActuatorControlResponse, error) {

	request := &SetActuatorControlRequest{}
	ctx := context.Background()
	request.ActuatorControl = actuatorControl

	response, err := s.Client.SetActuatorControl(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Set the attitude rate in terms of pitch, roll and yaw angular rate along with thrust.

   Parameters
   ----------
   attitudeRate *AttitudeRate



*/

func (s *ServiceImpl) SetAttitudeRate(attitudeRate *AttitudeRate) (*SetAttitudeRateResponse, error) {

	request := &SetAttitudeRateRequest{}
	ctx := context.Background()
	request.AttitudeRate = attitudeRate

	response, err := s.Client.SetAttitudeRate(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Set the position in NED coordinates and yaw.

   Parameters
   ----------
   positionNedYaw *PositionNedYaw



*/

func (s *ServiceImpl) SetPositionNed(positionNedYaw *PositionNedYaw) (*SetPositionNedResponse, error) {

	request := &SetPositionNedRequest{}
	ctx := context.Background()
	request.PositionNedYaw = positionNedYaw

	response, err := s.Client.SetPositionNed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Set the position in Global coordinates (latitude, longitude, altitude) and yaw

   Parameters
   ----------
   positionGlobalYaw *PositionGlobalYaw



*/

func (s *ServiceImpl) SetPositionGlobal(positionGlobalYaw *PositionGlobalYaw) (*SetPositionGlobalResponse, error) {

	request := &SetPositionGlobalRequest{}
	ctx := context.Background()
	request.PositionGlobalYaw = positionGlobalYaw

	response, err := s.Client.SetPositionGlobal(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Set the velocity in body coordinates and yaw angular rate. Not available for fixed-wing aircraft.

   Parameters
   ----------
   velocityBodyYawspeed *VelocityBodyYawspeed



*/

func (s *ServiceImpl) SetVelocityBody(velocityBodyYawspeed *VelocityBodyYawspeed) (*SetVelocityBodyResponse, error) {

	request := &SetVelocityBodyRequest{}
	ctx := context.Background()
	request.VelocityBodyYawspeed = velocityBodyYawspeed

	response, err := s.Client.SetVelocityBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Set the velocity in NED coordinates and yaw. Not available for fixed-wing aircraft.

   Parameters
   ----------
   velocityNedYaw *VelocityNedYaw



*/

func (s *ServiceImpl) SetVelocityNed(velocityNedYaw *VelocityNedYaw) (*SetVelocityNedResponse, error) {

	request := &SetVelocityNedRequest{}
	ctx := context.Background()
	request.VelocityNedYaw = velocityNedYaw

	response, err := s.Client.SetVelocityNed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Set the position in NED coordinates, with the velocity to be used as feed-forward.

   Parameters
   ----------
   positionNedYaw *PositionNedYaw


   velocityNedYaw *VelocityNedYaw



*/

func (s *ServiceImpl) SetPositionVelocityNed(positionNedYaw *PositionNedYaw, velocityNedYaw *VelocityNedYaw) (*SetPositionVelocityNedResponse, error) {

	request := &SetPositionVelocityNedRequest{}
	ctx := context.Background()
	request.PositionNedYaw = positionNedYaw

	request.VelocityNedYaw = velocityNedYaw

	response, err := s.Client.SetPositionVelocityNed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Set the acceleration in NED coordinates.

   Parameters
   ----------
   accelerationNed *AccelerationNed



*/

func (s *ServiceImpl) SetAccelerationNed(accelerationNed *AccelerationNed) (*SetAccelerationNedResponse, error) {

	request := &SetAccelerationNedRequest{}
	ctx := context.Background()
	request.AccelerationNed = accelerationNed

	response, err := s.Client.SetAccelerationNed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

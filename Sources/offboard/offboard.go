package offboard

import (
	"context"
)

type ServiceImpl struct {
	Client OffboardServiceClient
}

/*
Start Start offboard control.
*/
func (s *ServiceImpl) Start(
	ctx context.Context,

) (*StartResponse, error) {
	request := &StartRequest{}
	response, err := s.Client.Start(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
Stop Stop offboard control.

	The vehicle will be put into Hold mode: https://docs.px4.io/en/flight_modes/hold.html
*/
func (s *ServiceImpl) Stop(
	ctx context.Context,

) (*StopResponse, error) {
	request := &StopRequest{}
	response, err := s.Client.Stop(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
IsActive Check if offboard control is active.

	True means that the vehicle is in offboard mode and we are actively sending
	setpoints.
*/
func (s *ServiceImpl) IsActive(
	ctx context.Context,

) (*IsActiveResponse, error) {
	request := &IsActiveRequest{}
	response, err := s.Client.IsActive(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetAttitude Set the attitude in terms of roll, pitch and yaw in degrees with thrust.
*/
func (s *ServiceImpl) SetAttitude(
	ctx context.Context,
	attitude *Attitude,

) (*SetAttitudeResponse, error) {
	request := &SetAttitudeRequest{
		Attitude: attitude,
	}
	response, err := s.Client.SetAttitude(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetActuatorControl Set direct actuator control values to groups #0 and #1.

	First 8 controls will go to control group 0, the following 8 controls to control group 1 (if
	actuator_control.num_controls more than 8).
*/
func (s *ServiceImpl) SetActuatorControl(
	ctx context.Context,
	actuatorControl *ActuatorControl,

) (*SetActuatorControlResponse, error) {
	request := &SetActuatorControlRequest{
		ActuatorControl: actuatorControl,
	}
	response, err := s.Client.SetActuatorControl(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetAttitudeRate Set the attitude rate in terms of pitch, roll and yaw angular rate along with thrust.
*/
func (s *ServiceImpl) SetAttitudeRate(
	ctx context.Context,
	attitudeRate *AttitudeRate,

) (*SetAttitudeRateResponse, error) {
	request := &SetAttitudeRateRequest{
		AttitudeRate: attitudeRate,
	}
	response, err := s.Client.SetAttitudeRate(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetPositionNed Set the position in NED coordinates and yaw.
*/
func (s *ServiceImpl) SetPositionNed(
	ctx context.Context,
	positionNedYaw *PositionNedYaw,

) (*SetPositionNedResponse, error) {
	request := &SetPositionNedRequest{
		PositionNedYaw: positionNedYaw,
	}
	response, err := s.Client.SetPositionNed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetPositionGlobal Set the position in Global coordinates (latitude, longitude, altitude) and yaw
*/
func (s *ServiceImpl) SetPositionGlobal(
	ctx context.Context,
	positionGlobalYaw *PositionGlobalYaw,

) (*SetPositionGlobalResponse, error) {
	request := &SetPositionGlobalRequest{
		PositionGlobalYaw: positionGlobalYaw,
	}
	response, err := s.Client.SetPositionGlobal(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetVelocityBody Set the velocity in body coordinates and yaw angular rate. Not available for fixed-wing aircraft.
*/
func (s *ServiceImpl) SetVelocityBody(
	ctx context.Context,
	velocityBodyYawspeed *VelocityBodyYawspeed,

) (*SetVelocityBodyResponse, error) {
	request := &SetVelocityBodyRequest{
		VelocityBodyYawspeed: velocityBodyYawspeed,
	}
	response, err := s.Client.SetVelocityBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetVelocityNed Set the velocity in NED coordinates and yaw. Not available for fixed-wing aircraft.
*/
func (s *ServiceImpl) SetVelocityNed(
	ctx context.Context,
	velocityNedYaw *VelocityNedYaw,

) (*SetVelocityNedResponse, error) {
	request := &SetVelocityNedRequest{
		VelocityNedYaw: velocityNedYaw,
	}
	response, err := s.Client.SetVelocityNed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetPositionVelocityNed Set the position in NED coordinates, with the velocity to be used as feed-forward.
*/
func (s *ServiceImpl) SetPositionVelocityNed(
	ctx context.Context,
	positionNedYaw *PositionNedYaw,

	velocityNedYaw *VelocityNedYaw,

) (*SetPositionVelocityNedResponse, error) {
	request := &SetPositionVelocityNedRequest{
		PositionNedYaw: positionNedYaw,

		VelocityNedYaw: velocityNedYaw,
	}
	response, err := s.Client.SetPositionVelocityNed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetPositionVelocityAccelerationNed Set the position, velocity and acceleration in NED coordinates, with velocity and acceleration used as feed-forward.
*/
func (s *ServiceImpl) SetPositionVelocityAccelerationNed(
	ctx context.Context,
	positionNedYaw *PositionNedYaw,

	velocityNedYaw *VelocityNedYaw,

	accelerationNed *AccelerationNed,

) (*SetPositionVelocityAccelerationNedResponse, error) {
	request := &SetPositionVelocityAccelerationNedRequest{
		PositionNedYaw: positionNedYaw,

		VelocityNedYaw: velocityNedYaw,

		AccelerationNed: accelerationNed,
	}
	response, err := s.Client.SetPositionVelocityAccelerationNed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetAccelerationNed Set the acceleration in NED coordinates.
*/
func (s *ServiceImpl) SetAccelerationNed(
	ctx context.Context,
	accelerationNed *AccelerationNed,

) (*SetAccelerationNedResponse, error) {
	request := &SetAccelerationNedRequest{
		AccelerationNed: accelerationNed,
	}
	response, err := s.Client.SetAccelerationNed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

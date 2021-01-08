package action

import (
	"context"
)

type ServiceImpl struct {
	Client ActionServiceClient
}

/*
   Send command to arm the drone.

   Arming a drone normally causes motors to spin at idle.
   Before arming take all safety precautions and stand clear of the drone!


*/

func (s *ServiceImpl) Arm(ctx context.Context) (*ArmResponse, error) {

	request := &ArmRequest{}
	response, err := s.Client.Arm(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Send command to disarm the drone.

   This will disarm a drone that considers itself landed. If flying, the drone should
   reject the disarm command. Disarming means that all motors will stop.


*/

func (s *ServiceImpl) Disarm(ctx context.Context) (*DisarmResponse, error) {

	request := &DisarmRequest{}
	response, err := s.Client.Disarm(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Send command to take off and hover.

   This switches the drone into position control mode and commands
   it to take off and hover at the takeoff altitude.

   Note that the vehicle must be armed before it can take off.


*/

func (s *ServiceImpl) Takeoff(ctx context.Context) (*TakeoffResponse, error) {

	request := &TakeoffRequest{}
	response, err := s.Client.Takeoff(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Send command to land at the current position.

   This switches the drone to 'Land' flight mode.


*/

func (s *ServiceImpl) Land(ctx context.Context) (*LandResponse, error) {

	request := &LandRequest{}
	response, err := s.Client.Land(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Send command to reboot the drone components.

   This will reboot the autopilot, companion computer, camera and gimbal.


*/

func (s *ServiceImpl) Reboot(ctx context.Context) (*RebootResponse, error) {

	request := &RebootRequest{}
	response, err := s.Client.Reboot(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Send command to shut down the drone components.

   This will shut down the autopilot, onboard computer, camera and gimbal.
   This command should only be used when the autopilot is disarmed and autopilots commonly
   reject it if they are not already ready to shut down.


*/

func (s *ServiceImpl) Shutdown(ctx context.Context) (*ShutdownResponse, error) {

	request := &ShutdownRequest{}
	response, err := s.Client.Shutdown(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Send command to terminate the drone.

   This will run the terminate routine as configured on the drone (e.g. disarm and open the parachute).


*/

func (s *ServiceImpl) Terminate(ctx context.Context) (*TerminateResponse, error) {

	request := &TerminateRequest{}
	response, err := s.Client.Terminate(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Send command to kill the drone.

   This will disarm a drone irrespective of whether it is landed or flying.
   Note that the drone will fall out of the sky if this command is used while flying.


*/

func (s *ServiceImpl) Kill(ctx context.Context) (*KillResponse, error) {

	request := &KillRequest{}
	response, err := s.Client.Kill(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Send command to return to the launch (takeoff) position and land.

   This switches the drone into [Return mode](https://docs.px4.io/master/en/flight_modes/return.html) which
   generally means it will rise up to a certain altitude to clear any obstacles before heading
   back to the launch (takeoff) position and land there.


*/

func (s *ServiceImpl) ReturnToLaunch(ctx context.Context) (*ReturnToLaunchResponse, error) {

	request := &ReturnToLaunchRequest{}
	response, err := s.Client.ReturnToLaunch(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Send command to move the vehicle to a specific global position.

   The latitude and longitude are given in degrees (WGS84 frame) and the altitude
   in meters AMSL (above mean sea level).

   The yaw angle is in degrees (frame is NED, 0 is North, positive is clockwise).

   Parameters
   ----------
   latitudeDeg float64

   longitudeDeg float64

   absoluteAltitudeM float32

   yawDeg float32


*/

func (s *ServiceImpl) GotoLocation(ctx context.Context, latitudeDeg float64, longitudeDeg float64, absoluteAltitudeM float32, yawDeg float32) (*GotoLocationResponse, error) {

	request := &GotoLocationRequest{}
	request.LatitudeDeg = latitudeDeg
	request.LongitudeDeg = longitudeDeg
	request.AbsoluteAltitudeM = absoluteAltitudeM
	request.YawDeg = yawDeg
	response, err := s.Client.GotoLocation(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Send command to transition the drone to fixedwing.

   The associated action will only be executed for VTOL vehicles (on other vehicle types the
   command will fail). The command will succeed if called when the vehicle
   is already in fixedwing mode.


*/

func (s *ServiceImpl) TransitionToFixedwing(ctx context.Context) (*TransitionToFixedwingResponse, error) {

	request := &TransitionToFixedwingRequest{}
	response, err := s.Client.TransitionToFixedwing(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Send command to transition the drone to multicopter.

   The associated action will only be executed for VTOL vehicles (on other vehicle types the
   command will fail). The command will succeed if called when the vehicle
   is already in multicopter mode.


*/

func (s *ServiceImpl) TransitionToMulticopter(ctx context.Context) (*TransitionToMulticopterResponse, error) {

	request := &TransitionToMulticopterRequest{}
	response, err := s.Client.TransitionToMulticopter(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Get the takeoff altitude (in meters above ground).



   Returns
   -------
   False
   Altitude : float32
        Takeoff altitude relative to ground/takeoff location (in meters)


*/

func (s *ServiceImpl) GetTakeoffAltitude(ctx context.Context) (*GetTakeoffAltitudeResponse, error) {
	request := &GetTakeoffAltitudeRequest{}
	response, err := s.Client.GetTakeoffAltitude(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
   Set takeoff altitude (in meters above ground).

   Parameters
   ----------
   altitude float32


*/

func (s *ServiceImpl) SetTakeoffAltitude(ctx context.Context, altitude float32) (*SetTakeoffAltitudeResponse, error) {

	request := &SetTakeoffAltitudeRequest{}
	request.Altitude = altitude
	response, err := s.Client.SetTakeoffAltitude(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Get the vehicle maximum speed (in metres/second).



   Returns
   -------
   False
   Speed : float32
        Maximum speed (in metres/second)


*/

func (s *ServiceImpl) GetMaximumSpeed(ctx context.Context) (*GetMaximumSpeedResponse, error) {
	request := &GetMaximumSpeedRequest{}
	response, err := s.Client.GetMaximumSpeed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
   Set vehicle maximum speed (in metres/second).

   Parameters
   ----------
   speed float32


*/

func (s *ServiceImpl) SetMaximumSpeed(ctx context.Context, speed float32) (*SetMaximumSpeedResponse, error) {

	request := &SetMaximumSpeedRequest{}
	request.Speed = speed
	response, err := s.Client.SetMaximumSpeed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Get the return to launch minimum return altitude (in meters).



   Returns
   -------
   False
   RelativeAltitudeM : float32
        Return altitude relative to takeoff location (in meters)


*/

func (s *ServiceImpl) GetReturnToLaunchAltitude(ctx context.Context) (*GetReturnToLaunchAltitudeResponse, error) {
	request := &GetReturnToLaunchAltitudeRequest{}
	response, err := s.Client.GetReturnToLaunchAltitude(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
   Set the return to launch minimum return altitude (in meters).

   Parameters
   ----------
   relativeAltitudeM float32


*/

func (s *ServiceImpl) SetReturnToLaunchAltitude(ctx context.Context, relativeAltitudeM float32) (*SetReturnToLaunchAltitudeResponse, error) {

	request := &SetReturnToLaunchAltitudeRequest{}
	request.RelativeAltitudeM = relativeAltitudeM
	response, err := s.Client.SetReturnToLaunchAltitude(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

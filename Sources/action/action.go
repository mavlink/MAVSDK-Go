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
   Send command to force-arm the drone without any checks.

   Attention: this is not to be used for normal flying but only bench tests!

   Arming a drone normally causes motors to spin at idle.
   Before arming take all safety precautions and stand clear of the drone!


*/

func (s *ServiceImpl) ArmForce(ctx context.Context) (*ArmForceResponse, error) {

	request := &ArmForceRequest{}
	response, err := s.Client.ArmForce(ctx, request)
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
   Send command do orbit to the drone.

   This will run the orbit routine with the given parameters.

   Parameters
   ----------
   radiusM float32

   velocityMs float32

   yawBehavior *OrbitYawBehavior


   latitudeDeg float64

   longitudeDeg float64

   absoluteAltitudeM float64


*/

func (s *ServiceImpl) DoOrbit(ctx context.Context, radiusM float32, velocityMs float32, yawBehavior *OrbitYawBehavior, latitudeDeg float64, longitudeDeg float64, absoluteAltitudeM float64) (*DoOrbitResponse, error) {

	request := &DoOrbitRequest{}
	request.RadiusM = radiusM
	request.VelocityMs = velocityMs
	request.YawBehavior = *yawBehavior
	request.LatitudeDeg = latitudeDeg
	request.LongitudeDeg = longitudeDeg
	request.AbsoluteAltitudeM = absoluteAltitudeM
	response, err := s.Client.DoOrbit(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Send command to hold position (a.k.a. "Loiter").

   Sends a command to drone to change to Hold flight mode, causing the
   vehicle to stop and maintain its current GPS position and altitude.

   Note: this command is specific to the PX4 Autopilot flight stack as
   it implies a change to a PX4-specific mode.


*/

func (s *ServiceImpl) Hold(ctx context.Context) (*HoldResponse, error) {

	request := &HoldRequest{}
	response, err := s.Client.Hold(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Send command to set the value of an actuator.

   Note that the index of the actuator starts at 1 and that the value goes from -1 to 1.

   Parameters
   ----------
   index int32

   value float32


*/

func (s *ServiceImpl) SetActuator(ctx context.Context, index int32, value float32) (*SetActuatorResponse, error) {

	request := &SetActuatorRequest{}
	request.Index = index
	request.Value = value
	response, err := s.Client.SetActuator(ctx, request)
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

/*
   Set current speed.

   This will set the speed during a mission, reposition, and similar.
   It is ephemeral, so not stored on the drone and does not survive a reboot.

   Parameters
   ----------
   speedMS float32


*/

func (s *ServiceImpl) SetCurrentSpeed(ctx context.Context, speedMS float32) (*SetCurrentSpeedResponse, error) {

	request := &SetCurrentSpeedRequest{}
	request.SpeedMS = speedMS
	response, err := s.Client.SetCurrentSpeed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

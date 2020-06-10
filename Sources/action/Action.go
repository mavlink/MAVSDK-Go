package action

import (
	"context"
	"fmt"
)

type ServiceImpl struct {
	Client ActionServiceClient
}

/*
   Send command to arm the drone.

   Arming a drone normally causes motors to spin at idle.
   Before arming take all safety precautions and stand clear of the drone!


*/

func (s *ServiceImpl) Arm() {

	request := &ArmRequest{}
	ctx := context.Background()
	response, err := s.Client.Arm(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing Arm grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for Arm")
	}

}

/*
   Send command to disarm the drone.

   This will disarm a drone that considers itself landed. If flying, the drone should
   reject the disarm command. Disarming means that all motors will stop.


*/

func (s *ServiceImpl) Disarm() {

	request := &DisarmRequest{}
	ctx := context.Background()
	response, err := s.Client.Disarm(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing Disarm grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for Disarm")
	}

}

/*
   Send command to take off and hover.

   This switches the drone into position control mode and commands
   it to take off and hover at the takeoff altitude.

   Note that the vehicle must be armed before it can take off.


*/

func (s *ServiceImpl) Takeoff() {

	request := &TakeoffRequest{}
	ctx := context.Background()
	response, err := s.Client.Takeoff(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing Takeoff grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for Takeoff")
	}

}

/*
   Send command to land at the current position.

   This switches the drone to 'Land' flight mode.


*/

func (s *ServiceImpl) Land() {

	request := &LandRequest{}
	ctx := context.Background()
	response, err := s.Client.Land(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing Land grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for Land")
	}

}

/*
   Send command to reboot the drone components.

   This will reboot the autopilot, companion computer, camera and gimbal.


*/

func (s *ServiceImpl) Reboot() {

	request := &RebootRequest{}
	ctx := context.Background()
	response, err := s.Client.Reboot(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing Reboot grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for Reboot")
	}

}

/*
   Send command to shut down the drone components.

   This will shut down the autopilot, onboard computer, camera and gimbal.
   This command should only be used when the autopilot is disarmed and autopilots commonly
   reject it if they are not already ready to shut down.


*/

func (s *ServiceImpl) Shutdown() {

	request := &ShutdownRequest{}
	ctx := context.Background()
	response, err := s.Client.Shutdown(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing Shutdown grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for Shutdown")
	}

}

/*
   Send command to kill the drone.

   This will disarm a drone irrespective of whether it is landed or flying.
   Note that the drone will fall out of the sky if this command is used while flying.


*/

func (s *ServiceImpl) Kill() {

	request := &KillRequest{}
	ctx := context.Background()
	response, err := s.Client.Kill(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing Kill grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for Kill")
	}

}

/*
   Send command to return to the launch (takeoff) position and land.

   This switches the drone into [Return mode](https://docs.px4.io/master/en/flight_modes/return.html) which
   generally means it will rise up to a certain altitude to clear any obstacles before heading
   back to the launch (takeoff) position and land there.


*/

func (s *ServiceImpl) ReturnToLaunch() {

	request := &ReturnToLaunchRequest{}
	ctx := context.Background()
	response, err := s.Client.ReturnToLaunch(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing ReturnToLaunch grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for ReturnToLaunch")
	}

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

func (s *ServiceImpl) GotoLocation(latitudeDeg float64, longitudeDeg float64, absoluteAltitudeM float32, yawDeg float32) {

	request := &GotoLocationRequest{}
	ctx := context.Background()
	request.LatitudeDeg = latitudeDeg
	request.LongitudeDeg = longitudeDeg
	request.AbsoluteAltitudeM = absoluteAltitudeM
	request.YawDeg = yawDeg
	response, err := s.Client.GotoLocation(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing GotoLocation grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for GotoLocation")
	}

}

/*
   Send command to transition the drone to fixedwing.

   The associated action will only be executed for VTOL vehicles (on other vehicle types the
   command will fail). The command will succeed if called when the vehicle
   is already in fixedwing mode.


*/

func (s *ServiceImpl) TransitionToFixedwing() {

	request := &TransitionToFixedwingRequest{}
	ctx := context.Background()
	response, err := s.Client.TransitionToFixedwing(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing TransitionToFixedwing grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for TransitionToFixedwing")
	}

}

/*
   Send command to transition the drone to multicopter.

   The associated action will only be executed for VTOL vehicles (on other vehicle types the
   command will fail). The command will succeed if called when the vehicle
   is already in multicopter mode.


*/

func (s *ServiceImpl) TransitionToMulticopter() {

	request := &TransitionToMulticopterRequest{}
	ctx := context.Background()
	response, err := s.Client.TransitionToMulticopter(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing TransitionToMulticopter grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for TransitionToMulticopter")
	}

}

/*
   Get the takeoff altitude (in meters above ground).



   Returns
   -------
   False
   Altitude : float32
        Takeoff altitude relative to ground/takeoff location (in meters)


*/

func (s *ServiceImpl) GetTakeoffAltitude() float32 {
	request := &GetTakeoffAltitudeRequest{}
	ctx := context.Background()
	response, err := s.Client.GetTakeoffAltitude(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while getting GetTakeoffAltitude")
	}

	return response.GetAltitude()

}

/*
   Set takeoff altitude (in meters above ground).

   Parameters
   ----------
   altitude float32


*/

func (s *ServiceImpl) SetTakeoffAltitude(altitude float32) {

	request := &SetTakeoffAltitudeRequest{}
	ctx := context.Background()
	request.Altitude = altitude
	response, err := s.Client.SetTakeoffAltitude(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetTakeoffAltitude grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetTakeoffAltitude")
	}

}

/*
   Get the vehicle maximum speed (in metres/second).



   Returns
   -------
   False
   Speed : float32
        Maximum speed (in metres/second)


*/

func (s *ServiceImpl) GetMaximumSpeed() float32 {
	request := &GetMaximumSpeedRequest{}
	ctx := context.Background()
	response, err := s.Client.GetMaximumSpeed(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while getting GetMaximumSpeed")
	}

	return response.GetSpeed()

}

/*
   Set vehicle maximum speed (in metres/second).

   Parameters
   ----------
   speed float32


*/

func (s *ServiceImpl) SetMaximumSpeed(speed float32) {

	request := &SetMaximumSpeedRequest{}
	ctx := context.Background()
	request.Speed = speed
	response, err := s.Client.SetMaximumSpeed(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetMaximumSpeed grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetMaximumSpeed")
	}

}

/*
   Get the return to launch minimum return altitude (in meters).



   Returns
   -------
   False
   RelativeAltitudeM : float32
        Return altitude relative to takeoff location (in meters)


*/

func (s *ServiceImpl) GetReturnToLaunchAltitude() float32 {
	request := &GetReturnToLaunchAltitudeRequest{}
	ctx := context.Background()
	response, err := s.Client.GetReturnToLaunchAltitude(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while getting GetReturnToLaunchAltitude")
	}

	return response.GetRelativeAltitudeM()

}

/*
   Set the return to launch minimum return altitude (in meters).

   Parameters
   ----------
   relativeAltitudeM float32


*/

func (s *ServiceImpl) SetReturnToLaunchAltitude(relativeAltitudeM float32) {

	request := &SetReturnToLaunchAltitudeRequest{}
	ctx := context.Background()
	request.RelativeAltitudeM = relativeAltitudeM
	response, err := s.Client.SetReturnToLaunchAltitude(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetReturnToLaunchAltitude grpc %v\n", err)
	}

	result := response.GetActionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetReturnToLaunchAltitude")
	}

}

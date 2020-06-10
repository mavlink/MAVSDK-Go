package telemetry

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct {
	Client TelemetryServiceClient
}

/*
   Subscribe to 'position' updates.


*/
func (a *ServiceImpl) Position() {
	request := &SubscribePositionRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribePosition(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &PositionResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'home position' updates.


*/
func (a *ServiceImpl) Home() {
	request := &SubscribeHomeRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeHome(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &HomeResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to in-air updates.


*/
func (a *ServiceImpl) InAir() {
	request := &SubscribeInAirRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeInAir(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &InAirResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to landed state updates


*/
func (a *ServiceImpl) LandedState() {
	request := &SubscribeLandedStateRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeLandedState(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &LandedStateResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to armed updates.


*/
func (a *ServiceImpl) Armed() {
	request := &SubscribeArmedRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeArmed(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &ArmedResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'attitude' updates (quaternion).


*/
func (a *ServiceImpl) AttitudeQuaternion() {
	request := &SubscribeAttitudeQuaternionRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeAttitudeQuaternion(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &AttitudeQuaternionResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'attitude' updates (Euler).


*/
func (a *ServiceImpl) AttitudeEuler() {
	request := &SubscribeAttitudeEulerRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeAttitudeEuler(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &AttitudeEulerResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'attitude' updates (angular velocity)


*/
func (a *ServiceImpl) AttitudeAngularVelocityBody() {
	request := &SubscribeAttitudeAngularVelocityBodyRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeAttitudeAngularVelocityBody(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &AttitudeAngularVelocityBodyResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'camera attitude' updates (quaternion).


*/
func (a *ServiceImpl) CameraAttitudeQuaternion() {
	request := &SubscribeCameraAttitudeQuaternionRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeCameraAttitudeQuaternion(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &CameraAttitudeQuaternionResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'camera attitude' updates (Euler).


*/
func (a *ServiceImpl) CameraAttitudeEuler() {
	request := &SubscribeCameraAttitudeEulerRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeCameraAttitudeEuler(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &CameraAttitudeEulerResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'ground speed' updates (NED).


*/
func (a *ServiceImpl) VelocityNed() {
	request := &SubscribeVelocityNedRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeVelocityNed(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &VelocityNedResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'GPS info' updates.


*/
func (a *ServiceImpl) GpsInfo() {
	request := &SubscribeGpsInfoRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeGpsInfo(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &GpsInfoResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'battery' updates.


*/
func (a *ServiceImpl) Battery() {
	request := &SubscribeBatteryRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeBattery(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &BatteryResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'flight mode' updates.


*/
func (a *ServiceImpl) FlightMode() {
	request := &SubscribeFlightModeRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeFlightMode(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &FlightModeResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'health' updates.


*/
func (a *ServiceImpl) Health() {
	request := &SubscribeHealthRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeHealth(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &HealthResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'RC status' updates.


*/
func (a *ServiceImpl) RcStatus() {
	request := &SubscribeRcStatusRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeRcStatus(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &RcStatusResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'status text' updates.


*/
func (a *ServiceImpl) StatusText() {
	request := &SubscribeStatusTextRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeStatusText(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &StatusTextResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'actuator control target' updates.


*/
func (a *ServiceImpl) ActuatorControlTarget() {
	request := &SubscribeActuatorControlTargetRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeActuatorControlTarget(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &ActuatorControlTargetResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'actuator output status' updates.


*/
func (a *ServiceImpl) ActuatorOutputStatus() {
	request := &SubscribeActuatorOutputStatusRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeActuatorOutputStatus(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &ActuatorOutputStatusResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'odometry' updates.


*/
func (a *ServiceImpl) Odometry() {
	request := &SubscribeOdometryRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeOdometry(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &OdometryResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'position velocity' updates.


*/
func (a *ServiceImpl) PositionVelocityNed() {
	request := &SubscribePositionVelocityNedRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribePositionVelocityNed(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &PositionVelocityNedResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'ground truth' updates.


*/
func (a *ServiceImpl) GroundTruth() {
	request := &SubscribeGroundTruthRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeGroundTruth(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &GroundTruthResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'fixedwing metrics' updates.


*/
func (a *ServiceImpl) FixedwingMetrics() {
	request := &SubscribeFixedwingMetricsRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeFixedwingMetrics(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &FixedwingMetricsResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'IMU' updates.


*/
func (a *ServiceImpl) Imu() {
	request := &SubscribeImuRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeImu(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &ImuResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'HealthAllOk' updates.


*/
func (a *ServiceImpl) HealthAllOk() {
	request := &SubscribeHealthAllOkRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeHealthAllOk(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &HealthAllOkResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Subscribe to 'unix epoch time' updates.


*/
func (a *ServiceImpl) UnixEpochTime() {
	request := &SubscribeUnixEpochTimeRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeUnixEpochTime(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &UnixEpochTimeResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Set rate to 'position' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRatePosition(rateHz float64) {

	request := &SetRatePositionRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRatePosition(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRatePosition grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRatePosition")
	}

}

/*
   Set rate to 'home position' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateHome(rateHz float64) {

	request := &SetRateHomeRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateHome(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateHome grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateHome")
	}

}

/*
   Set rate to in-air updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateInAir(rateHz float64) {

	request := &SetRateInAirRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateInAir(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateInAir grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateInAir")
	}

}

/*
   Set rate to landed state updates

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateLandedState(rateHz float64) {

	request := &SetRateLandedStateRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateLandedState(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateLandedState grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateLandedState")
	}

}

/*
   Set rate to 'attitude' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateAttitude(rateHz float64) {

	request := &SetRateAttitudeRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateAttitude(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateAttitude grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateAttitude")
	}

}

/*
   Set rate of camera attitude updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateCameraAttitude(rateHz float64) {

	request := &SetRateCameraAttitudeRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateCameraAttitude(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateCameraAttitude grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateCameraAttitude")
	}

}

/*
   Set rate to 'ground speed' updates (NED).

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateVelocityNed(rateHz float64) {

	request := &SetRateVelocityNedRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateVelocityNed(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateVelocityNed grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateVelocityNed")
	}

}

/*
   Set rate to 'GPS info' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateGpsInfo(rateHz float64) {

	request := &SetRateGpsInfoRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateGpsInfo(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateGpsInfo grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateGpsInfo")
	}

}

/*
   Set rate to 'battery' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateBattery(rateHz float64) {

	request := &SetRateBatteryRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateBattery(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateBattery grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateBattery")
	}

}

/*
   Set rate to 'RC status' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateRcStatus(rateHz float64) {

	request := &SetRateRcStatusRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateRcStatus(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateRcStatus grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateRcStatus")
	}

}

/*
   Set rate to 'actuator control target' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateActuatorControlTarget(rateHz float64) {

	request := &SetRateActuatorControlTargetRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateActuatorControlTarget(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateActuatorControlTarget grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateActuatorControlTarget")
	}

}

/*
   Set rate to 'actuator output status' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateActuatorOutputStatus(rateHz float64) {

	request := &SetRateActuatorOutputStatusRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateActuatorOutputStatus(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateActuatorOutputStatus grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateActuatorOutputStatus")
	}

}

/*
   Set rate to 'odometry' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateOdometry(rateHz float64) {

	request := &SetRateOdometryRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateOdometry(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateOdometry grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateOdometry")
	}

}

/*
   Set rate to 'position velocity' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRatePositionVelocityNed(rateHz float64) {

	request := &SetRatePositionVelocityNedRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRatePositionVelocityNed(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRatePositionVelocityNed grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRatePositionVelocityNed")
	}

}

/*
   Set rate to 'ground truth' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateGroundTruth(rateHz float64) {

	request := &SetRateGroundTruthRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateGroundTruth(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateGroundTruth grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateGroundTruth")
	}

}

/*
   Set rate to 'fixedwing metrics' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateFixedwingMetrics(rateHz float64) {

	request := &SetRateFixedwingMetricsRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateFixedwingMetrics(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateFixedwingMetrics grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateFixedwingMetrics")
	}

}

/*
   Set rate to 'IMU' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateImu(rateHz float64) {

	request := &SetRateImuRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateImu(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateImu grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateImu")
	}

}

/*
   Set rate to 'unix epoch time' updates.

   Parameters
   ----------
   rateHz float64


*/

func (s *ServiceImpl) SetRateUnixEpochTime(rateHz float64) {

	request := &SetRateUnixEpochTimeRequest{}
	ctx := context.Background()
	request.RateHz = rateHz
	response, err := s.Client.SetRateUnixEpochTime(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetRateUnixEpochTime grpc %v\n", err)
	}

	result := response.GetTelemetryResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != TelemetryResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetRateUnixEpochTime")
	}

}

package telemetry_server

import (
	"context"
)

type ServiceImpl struct {
	Client TelemetryServerServiceClient
}

/*
PublishPosition Publish to 'position' updates.
*/
func (s *ServiceImpl) PublishPosition(
	ctx context.Context,
	position *Position,

	velocityNed *VelocityNed,

	heading *Heading,

) (*PublishPositionResponse, error) {
	request := &PublishPositionRequest{
		Position: position,

		VelocityNed: velocityNed,

		Heading: heading,
	}
	response, err := s.Client.PublishPosition(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishHome Publish to 'home position' updates.
*/
func (s *ServiceImpl) PublishHome(
	ctx context.Context,
	home *Position,

) (*PublishHomeResponse, error) {
	request := &PublishHomeRequest{
		Home: home,
	}
	response, err := s.Client.PublishHome(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishSysStatus Publish 'sys status' updates.
*/
func (s *ServiceImpl) PublishSysStatus(
	ctx context.Context,
	battery *Battery,

	rcReceiverStatus bool,
	gyroStatus bool,
	accelStatus bool,
	magStatus bool,
	gpsStatus bool,

) (*PublishSysStatusResponse, error) {
	request := &PublishSysStatusRequest{
		Battery: battery,

		RcReceiverStatus: rcReceiverStatus,
		GyroStatus:       gyroStatus,
		AccelStatus:      accelStatus,
		MagStatus:        magStatus,
		GpsStatus:        gpsStatus,
	}
	response, err := s.Client.PublishSysStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishExtendedSysState Publish 'extended sys state' updates.
*/
func (s *ServiceImpl) PublishExtendedSysState(
	ctx context.Context,
	vtolState *VtolState,

	landedState *LandedState,

) (*PublishExtendedSysStateResponse, error) {
	request := &PublishExtendedSysStateRequest{
		VtolState:   *vtolState,
		LandedState: *landedState,
	}
	response, err := s.Client.PublishExtendedSysState(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishRawGps Publish to 'Raw GPS' updates.
*/
func (s *ServiceImpl) PublishRawGps(
	ctx context.Context,
	rawGps *RawGps,

	gpsInfo *GpsInfo,

) (*PublishRawGpsResponse, error) {
	request := &PublishRawGpsRequest{
		RawGps: rawGps,

		GpsInfo: gpsInfo,
	}
	response, err := s.Client.PublishRawGps(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishBattery Publish to 'battery' updates.
*/
func (s *ServiceImpl) PublishBattery(
	ctx context.Context,
	battery *Battery,

) (*PublishBatteryResponse, error) {
	request := &PublishBatteryRequest{
		Battery: battery,
	}
	response, err := s.Client.PublishBattery(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishStatusText Publish to 'status text' updates.
*/
func (s *ServiceImpl) PublishStatusText(
	ctx context.Context,
	statusText *StatusText,

) (*PublishStatusTextResponse, error) {
	request := &PublishStatusTextRequest{
		StatusText: statusText,
	}
	response, err := s.Client.PublishStatusText(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishOdometry Publish to 'odometry' updates.
*/
func (s *ServiceImpl) PublishOdometry(
	ctx context.Context,
	odometry *Odometry,

) (*PublishOdometryResponse, error) {
	request := &PublishOdometryRequest{
		Odometry: odometry,
	}
	response, err := s.Client.PublishOdometry(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishPositionVelocityNed Publish to 'position velocity' updates.
*/
func (s *ServiceImpl) PublishPositionVelocityNed(
	ctx context.Context,
	positionVelocityNed *PositionVelocityNed,

) (*PublishPositionVelocityNedResponse, error) {
	request := &PublishPositionVelocityNedRequest{
		PositionVelocityNed: positionVelocityNed,
	}
	response, err := s.Client.PublishPositionVelocityNed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishGroundTruth Publish to 'ground truth' updates.
*/
func (s *ServiceImpl) PublishGroundTruth(
	ctx context.Context,
	groundTruth *GroundTruth,

) (*PublishGroundTruthResponse, error) {
	request := &PublishGroundTruthRequest{
		GroundTruth: groundTruth,
	}
	response, err := s.Client.PublishGroundTruth(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishImu Publish to 'IMU' updates (in SI units in NED body frame).
*/
func (s *ServiceImpl) PublishImu(
	ctx context.Context,
	imu *Imu,

) (*PublishImuResponse, error) {
	request := &PublishImuRequest{
		Imu: imu,
	}
	response, err := s.Client.PublishImu(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishScaledImu Publish to 'Scaled IMU' updates.
*/
func (s *ServiceImpl) PublishScaledImu(
	ctx context.Context,
	imu *Imu,

) (*PublishScaledImuResponse, error) {
	request := &PublishScaledImuRequest{
		Imu: imu,
	}
	response, err := s.Client.PublishScaledImu(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishRawImu Publish to 'Raw IMU' updates.
*/
func (s *ServiceImpl) PublishRawImu(
	ctx context.Context,
	imu *Imu,

) (*PublishRawImuResponse, error) {
	request := &PublishRawImuRequest{
		Imu: imu,
	}
	response, err := s.Client.PublishRawImu(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishUnixEpochTime Publish to 'unix epoch time' updates.
*/
func (s *ServiceImpl) PublishUnixEpochTime(
	ctx context.Context,
	timeUs uint64,

) (*PublishUnixEpochTimeResponse, error) {
	request := &PublishUnixEpochTimeRequest{
		TimeUs: timeUs,
	}
	response, err := s.Client.PublishUnixEpochTime(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PublishDistanceSensor Publish to "distance sensor" updates.
*/
func (s *ServiceImpl) PublishDistanceSensor(
	ctx context.Context,
	distanceSensor *DistanceSensor,

) (*PublishDistanceSensorResponse, error) {
	request := &PublishDistanceSensorRequest{
		DistanceSensor: distanceSensor,
	}
	response, err := s.Client.PublishDistanceSensor(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

package telemetry_server

import (
	"context"
)

type ServiceImpl struct {
	Client TelemetryServerServiceClient
}

/*
   Publish to 'position' updates.

   Parameters
   ----------
   position *Position


   velocityNed *VelocityNed


   heading *Heading



*/

func (s *ServiceImpl) PublishPosition(ctx context.Context, position *Position, velocityNed *VelocityNed, heading *Heading) (*PublishPositionResponse, error) {

	request := &PublishPositionRequest{}
	request.Position = position

	request.VelocityNed = velocityNed

	request.Heading = heading

	response, err := s.Client.PublishPosition(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish to 'home position' updates.

   Parameters
   ----------
   home *Position



*/

func (s *ServiceImpl) PublishHome(ctx context.Context, home *Position) (*PublishHomeResponse, error) {

	request := &PublishHomeRequest{}
	request.Home = home

	response, err := s.Client.PublishHome(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish 'sys status' updates.

   Parameters
   ----------
   battery *Battery


   rcReceiverStatus bool

   gyroStatus bool

   accelStatus bool

   magStatus bool

   gpsStatus bool


*/

func (s *ServiceImpl) PublishSysStatus(ctx context.Context, battery *Battery, rcReceiverStatus bool, gyroStatus bool, accelStatus bool, magStatus bool, gpsStatus bool) (*PublishSysStatusResponse, error) {

	request := &PublishSysStatusRequest{}
	request.Battery = battery

	request.RcReceiverStatus = rcReceiverStatus
	request.GyroStatus = gyroStatus
	request.AccelStatus = accelStatus
	request.MagStatus = magStatus
	request.GpsStatus = gpsStatus
	response, err := s.Client.PublishSysStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish 'extended sys state' updates.

   Parameters
   ----------
   vtolState *VtolState


   landedState *LandedState



*/

func (s *ServiceImpl) PublishExtendedSysState(ctx context.Context, vtolState *VtolState, landedState *LandedState) (*PublishExtendedSysStateResponse, error) {

	request := &PublishExtendedSysStateRequest{}
	request.VtolState = *vtolState
	request.LandedState = *landedState
	response, err := s.Client.PublishExtendedSysState(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish to 'Raw GPS' updates.

   Parameters
   ----------
   rawGps *RawGps


   gpsInfo *GpsInfo



*/

func (s *ServiceImpl) PublishRawGps(ctx context.Context, rawGps *RawGps, gpsInfo *GpsInfo) (*PublishRawGpsResponse, error) {

	request := &PublishRawGpsRequest{}
	request.RawGps = rawGps

	request.GpsInfo = gpsInfo

	response, err := s.Client.PublishRawGps(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish to 'battery' updates.

   Parameters
   ----------
   battery *Battery



*/

func (s *ServiceImpl) PublishBattery(ctx context.Context, battery *Battery) (*PublishBatteryResponse, error) {

	request := &PublishBatteryRequest{}
	request.Battery = battery

	response, err := s.Client.PublishBattery(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish to 'status text' updates.

   Parameters
   ----------
   statusText *StatusText



*/

func (s *ServiceImpl) PublishStatusText(ctx context.Context, statusText *StatusText) (*PublishStatusTextResponse, error) {

	request := &PublishStatusTextRequest{}
	request.StatusText = statusText

	response, err := s.Client.PublishStatusText(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish to 'odometry' updates.

   Parameters
   ----------
   odometry *Odometry



*/

func (s *ServiceImpl) PublishOdometry(ctx context.Context, odometry *Odometry) (*PublishOdometryResponse, error) {

	request := &PublishOdometryRequest{}
	request.Odometry = odometry

	response, err := s.Client.PublishOdometry(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish to 'position velocity' updates.

   Parameters
   ----------
   positionVelocityNed *PositionVelocityNed



*/

func (s *ServiceImpl) PublishPositionVelocityNed(ctx context.Context, positionVelocityNed *PositionVelocityNed) (*PublishPositionVelocityNedResponse, error) {

	request := &PublishPositionVelocityNedRequest{}
	request.PositionVelocityNed = positionVelocityNed

	response, err := s.Client.PublishPositionVelocityNed(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish to 'ground truth' updates.

   Parameters
   ----------
   groundTruth *GroundTruth



*/

func (s *ServiceImpl) PublishGroundTruth(ctx context.Context, groundTruth *GroundTruth) (*PublishGroundTruthResponse, error) {

	request := &PublishGroundTruthRequest{}
	request.GroundTruth = groundTruth

	response, err := s.Client.PublishGroundTruth(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish to 'IMU' updates (in SI units in NED body frame).

   Parameters
   ----------
   imu *Imu



*/

func (s *ServiceImpl) PublishImu(ctx context.Context, imu *Imu) (*PublishImuResponse, error) {

	request := &PublishImuRequest{}
	request.Imu = imu

	response, err := s.Client.PublishImu(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish to 'Scaled IMU' updates.

   Parameters
   ----------
   imu *Imu



*/

func (s *ServiceImpl) PublishScaledImu(ctx context.Context, imu *Imu) (*PublishScaledImuResponse, error) {

	request := &PublishScaledImuRequest{}
	request.Imu = imu

	response, err := s.Client.PublishScaledImu(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish to 'Raw IMU' updates.

   Parameters
   ----------
   imu *Imu



*/

func (s *ServiceImpl) PublishRawImu(ctx context.Context, imu *Imu) (*PublishRawImuResponse, error) {

	request := &PublishRawImuRequest{}
	request.Imu = imu

	response, err := s.Client.PublishRawImu(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish to 'unix epoch time' updates.

   Parameters
   ----------
   timeUs uint64


*/

func (s *ServiceImpl) PublishUnixEpochTime(ctx context.Context, timeUs uint64) (*PublishUnixEpochTimeResponse, error) {

	request := &PublishUnixEpochTimeRequest{}
	request.TimeUs = timeUs
	response, err := s.Client.PublishUnixEpochTime(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Publish to "distance sensor" updates.

   Parameters
   ----------
   distanceSensor *DistanceSensor



*/

func (s *ServiceImpl) PublishDistanceSensor(ctx context.Context, distanceSensor *DistanceSensor) (*PublishDistanceSensorResponse, error) {

	request := &PublishDistanceSensorRequest{}
	request.DistanceSensor = distanceSensor

	response, err := s.Client.PublishDistanceSensor(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

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

func (s *ServiceImpl) PublishPosition(position *Position, velocityNed *VelocityNed, heading *Heading) (*PublishPositionResponse, error) {

	request := &PublishPositionRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) PublishHome(home *Position) (*PublishHomeResponse, error) {

	request := &PublishHomeRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) PublishSysStatus(battery *Battery, rcReceiverStatus bool, gyroStatus bool, accelStatus bool, magStatus bool, gpsStatus bool) (*PublishSysStatusResponse, error) {

	request := &PublishSysStatusRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) PublishExtendedSysState(vtolState *VtolState, landedState *LandedState) (*PublishExtendedSysStateResponse, error) {

	request := &PublishExtendedSysStateRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) PublishRawGps(rawGps *RawGps, gpsInfo *GpsInfo) (*PublishRawGpsResponse, error) {

	request := &PublishRawGpsRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) PublishBattery(battery *Battery) (*PublishBatteryResponse, error) {

	request := &PublishBatteryRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) PublishStatusText(statusText *StatusText) (*PublishStatusTextResponse, error) {

	request := &PublishStatusTextRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) PublishOdometry(odometry *Odometry) (*PublishOdometryResponse, error) {

	request := &PublishOdometryRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) PublishPositionVelocityNed(positionVelocityNed *PositionVelocityNed) (*PublishPositionVelocityNedResponse, error) {

	request := &PublishPositionVelocityNedRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) PublishGroundTruth(groundTruth *GroundTruth) (*PublishGroundTruthResponse, error) {

	request := &PublishGroundTruthRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) PublishImu(imu *Imu) (*PublishImuResponse, error) {

	request := &PublishImuRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) PublishScaledImu(imu *Imu) (*PublishScaledImuResponse, error) {

	request := &PublishScaledImuRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) PublishRawImu(imu *Imu) (*PublishRawImuResponse, error) {

	request := &PublishRawImuRequest{}
	ctx := context.Background()
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

func (s *ServiceImpl) PublishUnixEpochTime(timeUs uint64) (*PublishUnixEpochTimeResponse, error) {

	request := &PublishUnixEpochTimeRequest{}
	ctx := context.Background()
	request.TimeUs = timeUs
	response, err := s.Client.PublishUnixEpochTime(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

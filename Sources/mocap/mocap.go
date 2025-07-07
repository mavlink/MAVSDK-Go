package mocap

import (
	"context"
)

type ServiceImpl struct {
	Client MocapServiceClient
}

/*
SetVisionPositionEstimate Send Global position/attitude estimate from a vision source.
*/
func (s *ServiceImpl) SetVisionPositionEstimate(
	ctx context.Context,
	visionPositionEstimate *VisionPositionEstimate,

) (*SetVisionPositionEstimateResponse, error) {
	request := &SetVisionPositionEstimateRequest{
		VisionPositionEstimate: visionPositionEstimate,
	}
	response, err := s.Client.SetVisionPositionEstimate(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetVisionSpeedEstimate Send Global speed estimate from a vision source.
*/
func (s *ServiceImpl) SetVisionSpeedEstimate(
	ctx context.Context,
	visionSpeedEstimate *VisionSpeedEstimate,

) (*SetVisionSpeedEstimateResponse, error) {
	request := &SetVisionSpeedEstimateRequest{
		VisionSpeedEstimate: visionSpeedEstimate,
	}
	response, err := s.Client.SetVisionSpeedEstimate(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetAttitudePositionMocap Send motion capture attitude and position.
*/
func (s *ServiceImpl) SetAttitudePositionMocap(
	ctx context.Context,
	attitudePositionMocap *AttitudePositionMocap,

) (*SetAttitudePositionMocapResponse, error) {
	request := &SetAttitudePositionMocapRequest{
		AttitudePositionMocap: attitudePositionMocap,
	}
	response, err := s.Client.SetAttitudePositionMocap(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetOdometry Send odometry information with an external interface.
*/
func (s *ServiceImpl) SetOdometry(
	ctx context.Context,
	odometry *Odometry,

) (*SetOdometryResponse, error) {
	request := &SetOdometryRequest{
		Odometry: odometry,
	}
	response, err := s.Client.SetOdometry(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

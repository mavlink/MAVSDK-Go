package follow_me

import (
	"context"
)

type ServiceImpl struct {
	Client FollowMeServiceClient
}

/*
GetConfig Get current configuration.
*/
func (s *ServiceImpl) GetConfig(
	ctx context.Context,

) (*GetConfigResponse, error) {
	request := &GetConfigRequest{}
	response, err := s.Client.GetConfig(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetConfig Apply configuration by sending it to the system.
*/
func (s *ServiceImpl) SetConfig(
	ctx context.Context,
	config *Config,

) (*SetConfigResponse, error) {
	request := &SetConfigRequest{
		Config: config,
	}
	response, err := s.Client.SetConfig(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
IsActive Check if FollowMe is active.
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
SetTargetLocation Set location of the moving target.
*/
func (s *ServiceImpl) SetTargetLocation(
	ctx context.Context,
	location *TargetLocation,

) (*SetTargetLocationResponse, error) {
	request := &SetTargetLocationRequest{
		Location: location,
	}
	response, err := s.Client.SetTargetLocation(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
GetLastLocation Get the last location of the target.
*/
func (s *ServiceImpl) GetLastLocation(
	ctx context.Context,

) (*GetLastLocationResponse, error) {
	request := &GetLastLocationRequest{}
	response, err := s.Client.GetLastLocation(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
Start Start FollowMe mode.
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
Stop Stop FollowMe mode.
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

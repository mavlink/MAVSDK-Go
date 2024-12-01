package gripper

import (
	"context"
)

type ServiceImpl struct {
	Client GripperServiceClient
}

/*
Grab Gripper grab cargo.
*/
func (s *ServiceImpl) Grab(
	ctx context.Context,
	instance uint32,

) (*GrabResponse, error) {
	request := &GrabRequest{
		Instance: instance,
	}
	response, err := s.Client.Grab(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
Release Gripper release cargo.
*/
func (s *ServiceImpl) Release(
	ctx context.Context,
	instance uint32,

) (*ReleaseResponse, error) {
	request := &ReleaseRequest{
		Instance: instance,
	}
	response, err := s.Client.Release(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

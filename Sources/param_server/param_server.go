package param_server

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client ParamServerServiceClient
}

/*
RetrieveParamInt Retrieve an int parameter.

	If the type is wrong, the result will be `WRONG_TYPE`.
*/
func (s *ServiceImpl) RetrieveParamInt(
	ctx context.Context,
	name string,

) (*RetrieveParamIntResponse, error) {
	request := &RetrieveParamIntRequest{
		Name: name,
	}
	response, err := s.Client.RetrieveParamInt(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ProvideParamInt Provide an int parameter.

	If the type is wrong, the result will be `WRONG_TYPE`.
*/
func (s *ServiceImpl) ProvideParamInt(
	ctx context.Context,
	name string,
	value int32,

) (*ProvideParamIntResponse, error) {
	request := &ProvideParamIntRequest{
		Name:  name,
		Value: value,
	}
	response, err := s.Client.ProvideParamInt(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
RetrieveParamFloat Retrieve a float parameter.

	If the type is wrong, the result will be `WRONG_TYPE`.
*/
func (s *ServiceImpl) RetrieveParamFloat(
	ctx context.Context,
	name string,

) (*RetrieveParamFloatResponse, error) {
	request := &RetrieveParamFloatRequest{
		Name: name,
	}
	response, err := s.Client.RetrieveParamFloat(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ProvideParamFloat Provide a float parameter.

	If the type is wrong, the result will be `WRONG_TYPE`.
*/
func (s *ServiceImpl) ProvideParamFloat(
	ctx context.Context,
	name string,
	value float32,

) (*ProvideParamFloatResponse, error) {
	request := &ProvideParamFloatRequest{
		Name:  name,
		Value: value,
	}
	response, err := s.Client.ProvideParamFloat(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
RetrieveParamCustom Retrieve a custom parameter.

	If the type is wrong, the result will be `WRONG_TYPE`.
*/
func (s *ServiceImpl) RetrieveParamCustom(
	ctx context.Context,
	name string,

) (*RetrieveParamCustomResponse, error) {
	request := &RetrieveParamCustomRequest{
		Name: name,
	}
	response, err := s.Client.RetrieveParamCustom(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ProvideParamCustom Provide a custom parameter.

	If the type is wrong, the result will be `WRONG_TYPE`.
*/
func (s *ServiceImpl) ProvideParamCustom(
	ctx context.Context,
	name string,
	value string,

) (*ProvideParamCustomResponse, error) {
	request := &ProvideParamCustomRequest{
		Name:  name,
		Value: value,
	}
	response, err := s.Client.ProvideParamCustom(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
RetrieveAllParams Retrieve all parameters.
*/
func (s *ServiceImpl) RetrieveAllParams(
	ctx context.Context,

) (*RetrieveAllParamsResponse, error) {
	request := &RetrieveAllParamsRequest{}
	response, err := s.Client.RetrieveAllParams(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ChangedParamInt Subscribe to changed int param.
*/
func (a *ServiceImpl) ChangedParamInt(
	ctx context.Context,

) (<-chan *IntParam, error) {
	ch := make(chan *IntParam)
	request := &SubscribeChangedParamIntRequest{}
	stream, err := a.Client.SubscribeChangedParamInt(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ChangedParamIntResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive ChangedParamInt messages, err: %v", err)
			}
			ch <- m.GetParam()
		}
	}()
	return ch, nil
}

/*
ChangedParamFloat Subscribe to changed float param.
*/
func (a *ServiceImpl) ChangedParamFloat(
	ctx context.Context,

) (<-chan *FloatParam, error) {
	ch := make(chan *FloatParam)
	request := &SubscribeChangedParamFloatRequest{}
	stream, err := a.Client.SubscribeChangedParamFloat(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ChangedParamFloatResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive ChangedParamFloat messages, err: %v", err)
			}
			ch <- m.GetParam()
		}
	}()
	return ch, nil
}

/*
ChangedParamCustom Subscribe to changed custom param.
*/
func (a *ServiceImpl) ChangedParamCustom(
	ctx context.Context,

) (<-chan *CustomParam, error) {
	ch := make(chan *CustomParam)
	request := &SubscribeChangedParamCustomRequest{}
	stream, err := a.Client.SubscribeChangedParamCustom(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ChangedParamCustomResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive ChangedParamCustom messages, err: %v", err)
			}
			ch <- m.GetParam()
		}
	}()
	return ch, nil
}

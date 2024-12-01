package param

import (
	"context"
)

type ServiceImpl struct {
	Client ParamServiceClient
}

/*
GetParamInt Get an int parameter.

	If the type is wrong, the result will be `WRONG_TYPE`.
*/
func (s *ServiceImpl) GetParamInt(
	ctx context.Context,
	name string,

) (*GetParamIntResponse, error) {
	request := &GetParamIntRequest{
		Name: name,
	}
	response, err := s.Client.GetParamInt(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetParamInt Set an int parameter.

	If the type is wrong, the result will be `WRONG_TYPE`.
*/
func (s *ServiceImpl) SetParamInt(
	ctx context.Context,
	name string,
	value int32,

) (*SetParamIntResponse, error) {
	request := &SetParamIntRequest{
		Name:  name,
		Value: value,
	}
	response, err := s.Client.SetParamInt(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
GetParamFloat Get a float parameter.

	If the type is wrong, the result will be `WRONG_TYPE`.
*/
func (s *ServiceImpl) GetParamFloat(
	ctx context.Context,
	name string,

) (*GetParamFloatResponse, error) {
	request := &GetParamFloatRequest{
		Name: name,
	}
	response, err := s.Client.GetParamFloat(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetParamFloat Set a float parameter.

	If the type is wrong, the result will be `WRONG_TYPE`.
*/
func (s *ServiceImpl) SetParamFloat(
	ctx context.Context,
	name string,
	value float32,

) (*SetParamFloatResponse, error) {
	request := &SetParamFloatRequest{
		Name:  name,
		Value: value,
	}
	response, err := s.Client.SetParamFloat(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
GetParamCustom Get a custom parameter.

	If the type is wrong, the result will be `WRONG_TYPE`.
*/
func (s *ServiceImpl) GetParamCustom(
	ctx context.Context,
	name string,

) (*GetParamCustomResponse, error) {
	request := &GetParamCustomRequest{
		Name: name,
	}
	response, err := s.Client.GetParamCustom(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetParamCustom Set a custom parameter.

	If the type is wrong, the result will be `WRONG_TYPE`.
*/
func (s *ServiceImpl) SetParamCustom(
	ctx context.Context,
	name string,
	value string,

) (*SetParamCustomResponse, error) {
	request := &SetParamCustomRequest{
		Name:  name,
		Value: value,
	}
	response, err := s.Client.SetParamCustom(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
GetAllParams Get all parameters.
*/
func (s *ServiceImpl) GetAllParams(
	ctx context.Context,

) (*GetAllParamsResponse, error) {
	request := &GetAllParamsRequest{}
	response, err := s.Client.GetAllParams(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SelectComponent Select component ID of parameter component to talk to and param protocol version.

	Default is the autopilot component (1), and Version (0).
*/
func (s *ServiceImpl) SelectComponent(
	ctx context.Context,
	componentId int32,
	protocolVersion *ProtocolVersion,

) (*SelectComponentResponse, error) {
	request := &SelectComponentRequest{
		ComponentId:     componentId,
		ProtocolVersion: *protocolVersion,
	}
	response, err := s.Client.SelectComponent(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

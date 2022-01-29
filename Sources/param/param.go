package param

import (
	"context"
)

type ServiceImpl struct {
	Client ParamServiceClient
}

/*
   Get an int parameter.

   If the type is wrong, the result will be `WRONG_TYPE`.

   Parameters
   ----------
   name string

   Returns
   -------
   False
   Value : int32
        Value of the requested parameter


*/

func (s *ServiceImpl) GetParamInt(ctx context.Context, name string) (*GetParamIntResponse, error) {
	request := &GetParamIntRequest{}
	request.Name = name
	response, err := s.Client.GetParamInt(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
   Set an int parameter.

   If the type is wrong, the result will be `WRONG_TYPE`.

   Parameters
   ----------
   name string

   value int32


*/

func (s *ServiceImpl) SetParamInt(ctx context.Context, name string, value int32) (*SetParamIntResponse, error) {

	request := &SetParamIntRequest{}
	request.Name = name
	request.Value = value
	response, err := s.Client.SetParamInt(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Get a float parameter.

   If the type is wrong, the result will be `WRONG_TYPE`.

   Parameters
   ----------
   name string

   Returns
   -------
   False
   Value : float32
        Value of the requested parameter


*/

func (s *ServiceImpl) GetParamFloat(ctx context.Context, name string) (*GetParamFloatResponse, error) {
	request := &GetParamFloatRequest{}
	request.Name = name
	response, err := s.Client.GetParamFloat(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
   Set a float parameter.

   If the type is wrong, the result will be `WRONG_TYPE`.

   Parameters
   ----------
   name string

   value float32


*/

func (s *ServiceImpl) SetParamFloat(ctx context.Context, name string, value float32) (*SetParamFloatResponse, error) {

	request := &SetParamFloatRequest{}
	request.Name = name
	request.Value = value
	response, err := s.Client.SetParamFloat(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Get all parameters.



   Returns
   -------
   False
   Params : AllParams
        Collection of all parameters


*/

func (s *ServiceImpl) GetAllParams(ctx context.Context) (*GetAllParamsResponse, error) {
	request := &GetAllParamsRequest{}
	response, err := s.Client.GetAllParams(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

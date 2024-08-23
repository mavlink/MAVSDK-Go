package param_server

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct {
	Client ParamServerServiceClient
}

/*
   Retrieve an int parameter.

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

func (s *ServiceImpl) RetrieveParamInt(ctx context.Context, name string) (*RetrieveParamIntResponse, error) {
	request := &RetrieveParamIntRequest{}
	request.Name = name
	response, err := s.Client.RetrieveParamInt(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
   Provide an int parameter.

   If the type is wrong, the result will be `WRONG_TYPE`.

   Parameters
   ----------
   name string

   value int32


*/

func (s *ServiceImpl) ProvideParamInt(ctx context.Context, name string, value int32) (*ProvideParamIntResponse, error) {

	request := &ProvideParamIntRequest{}
	request.Name = name
	request.Value = value
	response, err := s.Client.ProvideParamInt(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Retrieve a float parameter.

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

func (s *ServiceImpl) RetrieveParamFloat(ctx context.Context, name string) (*RetrieveParamFloatResponse, error) {
	request := &RetrieveParamFloatRequest{}
	request.Name = name
	response, err := s.Client.RetrieveParamFloat(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
   Provide a float parameter.

   If the type is wrong, the result will be `WRONG_TYPE`.

   Parameters
   ----------
   name string

   value float32


*/

func (s *ServiceImpl) ProvideParamFloat(ctx context.Context, name string, value float32) (*ProvideParamFloatResponse, error) {

	request := &ProvideParamFloatRequest{}
	request.Name = name
	request.Value = value
	response, err := s.Client.ProvideParamFloat(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Retrieve a custom parameter.

   If the type is wrong, the result will be `WRONG_TYPE`.

   Parameters
   ----------
   name string

   Returns
   -------
   False
   Value : string
        Value of the requested parameter


*/

func (s *ServiceImpl) RetrieveParamCustom(ctx context.Context, name string) (*RetrieveParamCustomResponse, error) {
	request := &RetrieveParamCustomRequest{}
	request.Name = name
	response, err := s.Client.RetrieveParamCustom(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
   Provide a custom parameter.

   If the type is wrong, the result will be `WRONG_TYPE`.

   Parameters
   ----------
   name string

   value string


*/

func (s *ServiceImpl) ProvideParamCustom(ctx context.Context, name string, value string) (*ProvideParamCustomResponse, error) {

	request := &ProvideParamCustomRequest{}
	request.Name = name
	request.Value = value
	response, err := s.Client.ProvideParamCustom(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Retrieve all parameters.



   Returns
   -------
   False
   Params : AllParams
        Collection of all parameters


*/

func (s *ServiceImpl) RetrieveAllParams(ctx context.Context) (*RetrieveAllParamsResponse, error) {
	request := &RetrieveAllParamsRequest{}
	response, err := s.Client.RetrieveAllParams(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
   Subscribe to changed int param.


*/

func (a *ServiceImpl) ChangedParamInt(ctx context.Context) (<-chan *IntParam, error) {
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetParam()
		}
	}()
	return ch, nil
}

/*
   Subscribe to changed float param.


*/

func (a *ServiceImpl) ChangedParamFloat(ctx context.Context) (<-chan *FloatParam, error) {
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetParam()
		}
	}()
	return ch, nil
}

/*
   Subscribe to changed custom param.


*/

func (a *ServiceImpl) ChangedParamCustom(ctx context.Context) (<-chan *CustomParam, error) {
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetParam()
		}
	}()
	return ch, nil
}

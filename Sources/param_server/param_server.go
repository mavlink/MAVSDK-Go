package param_server

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
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


    func(s *ServiceImpl)RetrieveParamInt(ctx context.Context, name string) (*RetrieveParamIntResponse, error){
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

    func(s *ServiceImpl)ProvideParamInt(ctx context.Context, name string, value int32)(*ProvideParamIntResponse, error){
        
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


    func(s *ServiceImpl)RetrieveParamFloat(ctx context.Context, name string) (*RetrieveParamFloatResponse, error){
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

    func(s *ServiceImpl)ProvideParamFloat(ctx context.Context, name string, value float32)(*ProvideParamFloatResponse, error){
        
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
         Retrieve all parameters.

         

         Returns
         -------
         False
         Params : AllParams
              Collection of all parameters

         
    */


    func(s *ServiceImpl)RetrieveAllParams(ctx context.Context, ) (*RetrieveAllParamsResponse, error){
        request := &RetrieveAllParamsRequest{}
    	response, err := s.Client.RetrieveAllParams(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
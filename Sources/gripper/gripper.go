package gripper

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client GripperServiceClient
}
    /*
         Gripper grab cargo.

         Parameters
         ----------
         instance uint32

         
    */

    func(s *ServiceImpl)Grab(ctx context.Context, instance uint32)(*GrabResponse, error){
        
        request := &GrabRequest{}
    	request.Instance = instance
        response, err := s.Client.Grab(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Gripper release cargo.

         Parameters
         ----------
         instance uint32

         
    */

    func(s *ServiceImpl)Release(ctx context.Context, instance uint32)(*ReleaseResponse, error){
        
        request := &ReleaseRequest{}
    	request.Instance = instance
        response, err := s.Client.Release(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
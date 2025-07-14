package arm_authorizer_server

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client ArmAuthorizerServerServiceClient
}

     /*
         Subscribe to arm authorization request messages. Each request received should respond to using RespondArmAuthorization

         
    */

    func (a *ServiceImpl) ArmAuthorization(ctx context.Context, ) (<-chan   uint32  , error){
    		ch := make(chan   uint32  )
    		request := &SubscribeArmAuthorizationRequest{}
    		stream, err := a.Client.SubscribeArmAuthorization(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &ArmAuthorizationResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive ArmAuthorization messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetSystemId()
    			}
    		}()	
    	return ch, nil
    }
    /*
         Authorize arm for the specific time

         Parameters
         ----------
         validTimeS int32

         
    */

    func(s *ServiceImpl)AcceptArmAuthorization(ctx context.Context, validTimeS int32)(*AcceptArmAuthorizationResponse, error){
        
        request := &AcceptArmAuthorizationRequest{}
    	request.ValidTimeS = validTimeS
        response, err := s.Client.AcceptArmAuthorization(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Reject arm authorization request

         Parameters
         ----------
         temporarily bool

         reason *RejectionReason 
            

         extraInfo int32

         
    */

    func(s *ServiceImpl)RejectArmAuthorization(ctx context.Context, temporarily bool, reason *RejectionReason , extraInfo int32)(*RejectArmAuthorizationResponse, error){
        
        request := &RejectArmAuthorizationRequest{}
    	request.Temporarily = temporarily
        request.Reason = *reason
        request.ExtraInfo = extraInfo
        response, err := s.Client.RejectArmAuthorization(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
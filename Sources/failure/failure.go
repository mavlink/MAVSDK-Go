package failure

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client FailureServiceClient
}
    /*
         Injects a failure.

         Parameters
         ----------
         failureUnit *FailureUnit 
            

         failureType *FailureType 
            

         instance int32

         
    */

    func(s *ServiceImpl)Inject(ctx context.Context, failureUnit *FailureUnit , failureType *FailureType , instance int32)(*InjectResponse, error){
        
        request := &InjectRequest{}
    	request.FailureUnit = *failureUnit
        request.FailureType = *failureType
        request.Instance = instance
        response, err := s.Client.Inject(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
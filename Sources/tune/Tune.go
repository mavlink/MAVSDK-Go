package tune

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client TuneServiceClient
}
    /*
         Send a tune to be played by the system.

         Parameters
         ----------
         tuneDescription *TuneDescription 
            

         
    */

    func(s *ServiceImpl)PlayTune(tuneDescription *TuneDescription )(*PlayTuneResponse, error){
        
        request := &PlayTuneRequest{}
        ctx:= context.Background()
         request.TuneDescription = tuneDescription
            
        response, err := s.Client.PlayTune(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
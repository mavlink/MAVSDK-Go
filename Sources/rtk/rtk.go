package rtk

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client RtkServiceClient
}
    /*
         Send RTCM data.

         Parameters
         ----------
         rtcmData *RtcmData 
            

         
    */

    func(s *ServiceImpl)SendRtcmData(ctx context.Context, rtcmData *RtcmData )(*SendRtcmDataResponse, error){
        
        request := &SendRtcmDataRequest{}
    	request.RtcmData = rtcmData
            
        response, err := s.Client.SendRtcmData(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
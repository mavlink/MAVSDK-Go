package geofence

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client GeofenceServiceClient
}

    func(s *ServiceImpl)UploadGeofence(, polygons []*Polygons){
     request := &UploadGeofenceRequest{}
     ctx:= context.Background()
         request.Polygons = 
            
        response, err := s.Client.UploadGeofence(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing UploadGeofence grpc %v\n", err)
    	}
        fmt.Printf("")
        
        result := response.GetGeofenceResult()
        fmt.Printf("result %v\n", result)
        if result.Result != GeofenceResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for UploadGeofence")
        }
        
    }

       
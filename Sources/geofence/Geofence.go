package geofence

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client GeofenceServiceClient
}
    /*
         Upload a geofence.

         Polygons are uploaded to a drone. Once uploaded, the geofence will remain
         on the drone even if a connection is lost.

         Parameters
         ----------
         polygons []*Polygon

         
    */

    func(s *ServiceImpl)UploadGeofence(polygons []*Polygon){
        
        request := &UploadGeofenceRequest{}
        ctx:= context.Background()
         request.Polygons = polygons
            
        response, err := s.Client.UploadGeofence(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing UploadGeofence grpc %v\n", err)
    	}
        
        result := response.GetGeofenceResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != GeofenceResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for UploadGeofence")
        }
        
    }

       
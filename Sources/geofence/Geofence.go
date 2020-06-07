import (
	"context"
	"fmt"
	"io"
)

type Service interface{
    Result() GeofenceResult_Result

}

type ServiceImpl struct{
    Client GeofenceServiceClient
}

    func(s *ServiceImpl)upload_geofence(self, polygons []*Polygons){
     request = &UploadGeofenceRequest{}
         request.Polygons = 
            
        response, err = s.Client.UploadGeofence(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetGeofenceResult()
        fmt.Printf("result %v\n", result)
        if result.Result != GeofenceResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading UploadGeofence")
        }
        
    }

       
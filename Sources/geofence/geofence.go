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
         Upload geofences.

         Polygon and Circular geofences are uploaded to a drone. Once uploaded, the geofence will remain
         on the drone even if a connection is lost.

         Parameters
         ----------
         geofenceData *GeofenceData 
            

         
    */

    func(s *ServiceImpl)UploadGeofence(ctx context.Context, geofenceData *GeofenceData )(*UploadGeofenceResponse, error){
        
        request := &UploadGeofenceRequest{}
    	request.GeofenceData = geofenceData
            
        response, err := s.Client.UploadGeofence(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Clear all geofences saved on the vehicle.

         
    */

    func(s *ServiceImpl)ClearGeofence(ctx context.Context, )(*ClearGeofenceResponse, error){
        
        request := &ClearGeofenceRequest{}
    	response, err := s.Client.ClearGeofence(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
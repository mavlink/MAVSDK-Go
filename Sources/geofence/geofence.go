package geofence

import (
	"context"
	"fmt"
)

//Service interface for intertelemetry with telemetry interface of drone
type Service interface {
	Result() GeofenceResult
}

//ServiceImpl creates client and implements telemetryService
type ServiceImpl struct {
	Client GeofenceServiceClient
}

func (s *ServiceImpl) translateFromGRPC(response *GeofenceResult) GeofenceResult_Result {
	return response.Result
}

//UploadGeofence acts as a wrapper for UploadGeofence request from geofence grpc package
func (s *ServiceImpl) UploadGeofence(polygons []*Polygon) {
	request := &UploadGeofenceRequest{}
	request.Polygons = polygons

	ctx := context.Background()
	response, err := s.Client.UploadGeofence(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
	}
	result := response.GetGeofenceResult()
	fmt.Printf("result %v\n", result)
	if result.Result != GeofenceResult_RESULT_SUCCESS {
		fmt.Printf("Error while uploading Geofence")
	}
}

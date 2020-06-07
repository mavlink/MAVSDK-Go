package geofence

import (
	"context"
	"fmt"
)

type Service interface {
	Result() GeofenceResult_Result
}

type ServiceImpl struct {
	Client GeofenceServiceClient
}

func (s *ServiceImpl) upload_geofence(polygons []*Polygon) {
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
		fmt.Printf("Error while uploading UploadGeofence")
	}
}

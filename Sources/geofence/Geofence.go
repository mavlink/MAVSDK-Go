package geofence

import (
	"context"
	"fmt"
)

type ServiceImpl struct {
	Client GeofenceServiceClient
}

func (s *ServiceImpl) UploadGeofence(polygons []*Polygon) {
	request := &UploadGeofenceRequest{}
	ctx := context.Background()
	request.Polygons = polygons

	response, err := s.Client.UploadGeofence(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing UploadGeofence grpc %v\n", err)
	}

	result := response.GetGeofenceResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != GeofenceResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for UploadGeofence")
	}

}

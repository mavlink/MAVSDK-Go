package geofence

import (
	"context"
)

type ServiceImpl struct {
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

func (s *ServiceImpl) UploadGeofence(ctx context.Context, polygons []*Polygon) (*UploadGeofenceResponse, error) {

	request := &UploadGeofenceRequest{}
	request.Polygons = polygons

	response, err := s.Client.UploadGeofence(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

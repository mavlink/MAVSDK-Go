package geofence

import (
	"context"
)

type ServiceImpl struct {
	Client GeofenceServiceClient
}

/*
UploadGeofence Upload geofences.

	Polygon and Circular geofences are uploaded to a drone. Once uploaded, the geofence will remain
	on the drone even if a connection is lost.
*/
func (s *ServiceImpl) UploadGeofence(
	ctx context.Context,
	geofenceData *GeofenceData,

) (*UploadGeofenceResponse, error) {
	request := &UploadGeofenceRequest{
		GeofenceData: geofenceData,
	}
	response, err := s.Client.UploadGeofence(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ClearGeofence Clear all geofences saved on the vehicle.
*/
func (s *ServiceImpl) ClearGeofence(
	ctx context.Context,

) (*ClearGeofenceResponse, error) {
	request := &ClearGeofenceRequest{}
	response, err := s.Client.ClearGeofence(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

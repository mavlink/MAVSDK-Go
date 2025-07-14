package mission_raw

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client MissionRawServiceClient
}

/*
UploadMission Upload a list of raw mission items to the system.

	The raw mission items are uploaded to a drone. Once uploaded the mission
	can be started and executed even if the connection is lost.
*/
func (s *ServiceImpl) UploadMission(
	ctx context.Context,
	missionItems []*MissionItem,

) (*UploadMissionResponse, error) {
	request := &UploadMissionRequest{
		MissionItems: missionItems,
	}
	response, err := s.Client.UploadMission(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
UploadGeofence Upload a list of geofence items to the system.
*/
func (s *ServiceImpl) UploadGeofence(
	ctx context.Context,
	missionItems []*MissionItem,

) (*UploadGeofenceResponse, error) {
	request := &UploadGeofenceRequest{
		MissionItems: missionItems,
	}
	response, err := s.Client.UploadGeofence(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
UploadRallyPoints Upload a list of rally point items to the system.
*/
func (s *ServiceImpl) UploadRallyPoints(
	ctx context.Context,
	missionItems []*MissionItem,

) (*UploadRallyPointsResponse, error) {
	request := &UploadRallyPointsRequest{
		MissionItems: missionItems,
	}
	response, err := s.Client.UploadRallyPoints(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
CancelMissionUpload Cancel an ongoing mission upload.
*/
func (s *ServiceImpl) CancelMissionUpload(
	ctx context.Context,

) (*CancelMissionUploadResponse, error) {
	request := &CancelMissionUploadRequest{}
	response, err := s.Client.CancelMissionUpload(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
DownloadMission Download a list of raw mission items from the system (asynchronous).
*/
func (s *ServiceImpl) DownloadMission(
	ctx context.Context,

) (*DownloadMissionResponse, error) {
	request := &DownloadMissionRequest{}
	response, err := s.Client.DownloadMission(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
DownloadGeofence Download a list of raw geofence items from the system (asynchronous).
*/
func (s *ServiceImpl) DownloadGeofence(
	ctx context.Context,

) (*DownloadGeofenceResponse, error) {
	request := &DownloadGeofenceRequest{}
	response, err := s.Client.DownloadGeofence(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
DownloadRallypoints Download a list of raw rallypoint items from the system (asynchronous).
*/
func (s *ServiceImpl) DownloadRallypoints(
	ctx context.Context,

) (*DownloadRallypointsResponse, error) {
	request := &DownloadRallypointsRequest{}
	response, err := s.Client.DownloadRallypoints(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
CancelMissionDownload Cancel an ongoing mission download.
*/
func (s *ServiceImpl) CancelMissionDownload(
	ctx context.Context,

) (*CancelMissionDownloadResponse, error) {
	request := &CancelMissionDownloadRequest{}
	response, err := s.Client.CancelMissionDownload(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
StartMission Start the mission.

	A mission must be uploaded to the vehicle before this can be called.
*/
func (s *ServiceImpl) StartMission(
	ctx context.Context,

) (*StartMissionResponse, error) {
	request := &StartMissionRequest{}
	response, err := s.Client.StartMission(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PauseMission Pause the mission.

	Pausing the mission puts the vehicle into
	[HOLD mode](https://docs.px4.io/en/flight_modes/hold.html).
	A multicopter should just hover at the spot while a fixedwing vehicle should loiter
	around the location where it paused.
*/
func (s *ServiceImpl) PauseMission(
	ctx context.Context,

) (*PauseMissionResponse, error) {
	request := &PauseMissionRequest{}
	response, err := s.Client.PauseMission(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ClearMission Clear the mission saved on the vehicle.
*/
func (s *ServiceImpl) ClearMission(
	ctx context.Context,

) (*ClearMissionResponse, error) {
	request := &ClearMissionRequest{}
	response, err := s.Client.ClearMission(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetCurrentMissionItem Sets the raw mission item index to go to.

	By setting the current index to 0, the mission is restarted from the beginning. If it is set
	to a specific index of a raw mission item, the mission will be set to this item.
*/
func (s *ServiceImpl) SetCurrentMissionItem(
	ctx context.Context,
	index int32,

) (*SetCurrentMissionItemResponse, error) {
	request := &SetCurrentMissionItemRequest{
		Index: index,
	}
	response, err := s.Client.SetCurrentMissionItem(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
MissionProgress Subscribe to mission progress updates.
*/
func (a *ServiceImpl) MissionProgress(
	ctx context.Context,

) (<-chan *MissionProgress, error) {
	ch := make(chan *MissionProgress)
	request := &SubscribeMissionProgressRequest{}
	stream, err := a.Client.SubscribeMissionProgress(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &MissionProgressResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive MissionProgress messages, err: %v", err)
			}
			ch <- m.GetMissionProgress()
		}
	}()
	return ch, nil
}

/*
MissionChanged Subscribes to mission changed.

	This notification can be used to be informed if a ground station has
	been uploaded or changed by a ground station or companion computer.

	@param callback Callback to notify about change.
*/
func (a *ServiceImpl) MissionChanged(
	ctx context.Context,

) (<-chan bool, error) {
	ch := make(chan bool)
	request := &SubscribeMissionChangedRequest{}
	stream, err := a.Client.SubscribeMissionChanged(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &MissionChangedResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive MissionChanged messages, err: %v", err)
			}
			ch <- m.GetMissionChanged()
		}
	}()
	return ch, nil
}

/*
ImportQgroundcontrolMission Import a QGroundControl missions in JSON .plan format, from a file.

	Supported:
	- Waypoints
	- Survey
	Not supported:
	- Structure Scan
*/
func (s *ServiceImpl) ImportQgroundcontrolMission(
	ctx context.Context,
	qgcPlanPath string,

) (*ImportQgroundcontrolMissionResponse, error) {
	request := &ImportQgroundcontrolMissionRequest{
		QgcPlanPath: qgcPlanPath,
	}
	response, err := s.Client.ImportQgroundcontrolMission(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ImportQgroundcontrolMissionFromString Import a QGroundControl missions in JSON .plan format, from a string.

	Supported:
	- Waypoints
	- Survey
	Not supported:
	- Structure Scan
*/
func (s *ServiceImpl) ImportQgroundcontrolMissionFromString(
	ctx context.Context,
	qgcPlan string,

) (*ImportQgroundcontrolMissionFromStringResponse, error) {
	request := &ImportQgroundcontrolMissionFromStringRequest{
		QgcPlan: qgcPlan,
	}
	response, err := s.Client.ImportQgroundcontrolMissionFromString(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
IsMissionFinished Check if the mission is finished.

	Returns true if the mission is finished, false otherwise.
*/
func (s *ServiceImpl) IsMissionFinished(
	ctx context.Context,

) (*IsMissionFinishedResponse, error) {
	request := &IsMissionFinishedRequest{}
	response, err := s.Client.IsMissionFinished(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

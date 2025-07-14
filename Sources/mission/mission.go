package mission

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client MissionServiceClient
}

/*
UploadMission Upload a list of mission items to the system.

	The mission items are uploaded to a drone. Once uploaded the mission can be started and
	executed even if the connection is lost.
*/
func (s *ServiceImpl) UploadMission(
	ctx context.Context,
	missionPlan *MissionPlan,

) (*UploadMissionResponse, error) {
	request := &UploadMissionRequest{
		MissionPlan: missionPlan,
	}
	response, err := s.Client.UploadMission(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
UploadMissionWithProgress Upload a list of mission items to the system and report upload progress.

	The mission items are uploaded to a drone. Once uploaded the mission can be started and
	executed even if the connection is lost.
*/
func (a *ServiceImpl) UploadMissionWithProgress(
	ctx context.Context,
	missionPlan *MissionPlan,

) (<-chan *ProgressData, error) {
	ch := make(chan *ProgressData)
	request := &SubscribeUploadMissionWithProgressRequest{
		MissionPlan: missionPlan,
	}
	stream, err := a.Client.SubscribeUploadMissionWithProgress(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &UploadMissionWithProgressResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive UploadMissionWithProgress messages, err: %v", err)
			}
			ch <- m.GetProgressData()
		}
	}()
	return ch, nil
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
DownloadMission Download a list of mission items from the system (asynchronous).

	Will fail if any of the downloaded mission items are not supported
	by the MAVSDK API.
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
DownloadMissionWithProgress Download a list of mission items from the system (asynchronous) and report progress.

	Will fail if any of the downloaded mission items are not supported
	by the MAVSDK API.
*/
func (a *ServiceImpl) DownloadMissionWithProgress(
	ctx context.Context,

) (<-chan *ProgressDataOrMission, error) {
	ch := make(chan *ProgressDataOrMission)
	request := &SubscribeDownloadMissionWithProgressRequest{}
	stream, err := a.Client.SubscribeDownloadMissionWithProgress(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &DownloadMissionWithProgressResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive DownloadMissionWithProgress messages, err: %v", err)
			}
			ch <- m.GetProgressData()
		}
	}()
	return ch, nil
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
SetCurrentMissionItem Sets the mission item index to go to.

	By setting the current index to 0, the mission is restarted from the beginning. If it is set
	to a specific index of a mission item, the mission will be set to this item.

	Note that this is not necessarily true for general missions using MAVLink if loop counters
	are used.
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
IsMissionFinished Check if the mission has been finished.
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
GetReturnToLaunchAfterMission Get whether to trigger Return-to-Launch (RTL) after mission is complete.

	Before getting this option, it needs to be set, or a mission
	needs to be downloaded.
*/
func (s *ServiceImpl) GetReturnToLaunchAfterMission(
	ctx context.Context,

) (*GetReturnToLaunchAfterMissionResponse, error) {
	request := &GetReturnToLaunchAfterMissionRequest{}
	response, err := s.Client.GetReturnToLaunchAfterMission(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetReturnToLaunchAfterMission Set whether to trigger Return-to-Launch (RTL) after the mission is complete.

	This will only take effect for the next mission upload, meaning that
	the mission may have to be uploaded again.
*/
func (s *ServiceImpl) SetReturnToLaunchAfterMission(
	ctx context.Context,
	enable bool,

) (*SetReturnToLaunchAfterMissionResponse, error) {
	request := &SetReturnToLaunchAfterMissionRequest{
		Enable: enable,
	}
	response, err := s.Client.SetReturnToLaunchAfterMission(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

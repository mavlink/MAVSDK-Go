package mission

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct {
	Client MissionServiceClient
}

/*
   Upload a list of mission items to the system.

   The mission items are uploaded to a drone. Once uploaded the mission can be started and
   executed even if the connection is lost.

   Parameters
   ----------
   missionPlan *MissionPlan



*/

func (s *ServiceImpl) UploadMission(missionPlan *MissionPlan) {

	request := &UploadMissionRequest{}
	ctx := context.Background()
	request.MissionPlan = missionPlan

	response, err := s.Client.UploadMission(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing UploadMission grpc %v\n", err)
	}

	result := response.GetMissionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != MissionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for UploadMission")
	}

}

/*
   Cancel an ongoing mission upload.


*/

func (s *ServiceImpl) CancelMissionUpload() {

	request := &CancelMissionUploadRequest{}
	ctx := context.Background()
	response, err := s.Client.CancelMissionUpload(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing CancelMissionUpload grpc %v\n", err)
	}

	result := response.GetMissionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != MissionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for CancelMissionUpload")
	}

}

/*
   Download a list of mission items from the system (asynchronous).

   Will fail if any of the downloaded mission items are not supported
   by the MAVSDK API.



   Returns
   -------
   False
   MissionPlan : MissionPlan
        The mission plan


*/

func (s *ServiceImpl) DownloadMission() *MissionPlan {
	request := &DownloadMissionRequest{}
	ctx := context.Background()
	response, err := s.Client.DownloadMission(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
	}

	result := response.GetMissionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != MissionResult_RESULT_SUCCESS {
		fmt.Printf("Error while getting DownloadMission")
	}

	return response.GetMissionPlan()

}

/*
   Cancel an ongoing mission download.


*/

func (s *ServiceImpl) CancelMissionDownload() {

	request := &CancelMissionDownloadRequest{}
	ctx := context.Background()
	response, err := s.Client.CancelMissionDownload(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing CancelMissionDownload grpc %v\n", err)
	}

	result := response.GetMissionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != MissionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for CancelMissionDownload")
	}

}

/*
   Start the mission.

   A mission must be uploaded to the vehicle before this can be called.


*/

func (s *ServiceImpl) StartMission() {

	request := &StartMissionRequest{}
	ctx := context.Background()
	response, err := s.Client.StartMission(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing StartMission grpc %v\n", err)
	}

	result := response.GetMissionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != MissionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for StartMission")
	}

}

/*
   Pause the mission.

   Pausing the mission puts the vehicle into
   [HOLD mode](https://docs.px4.io/en/flight_modes/hold.html).
   A multicopter should just hover at the spot while a fixedwing vehicle should loiter
   around the location where it paused.


*/

func (s *ServiceImpl) PauseMission() {

	request := &PauseMissionRequest{}
	ctx := context.Background()
	response, err := s.Client.PauseMission(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing PauseMission grpc %v\n", err)
	}

	result := response.GetMissionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != MissionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for PauseMission")
	}

}

/*
   Clear the mission saved on the vehicle.


*/

func (s *ServiceImpl) ClearMission() {

	request := &ClearMissionRequest{}
	ctx := context.Background()
	response, err := s.Client.ClearMission(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing ClearMission grpc %v\n", err)
	}

	result := response.GetMissionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != MissionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for ClearMission")
	}

}

/*
   Sets the mission item index to go to.

   By setting the current index to 0, the mission is restarted from the beginning. If it is set
   to a specific index of a mission item, the mission will be set to this item.

   Note that this is not necessarily true for general missions using MAVLink if loop counters
   are used.

   Parameters
   ----------
   index int32


*/

func (s *ServiceImpl) SetCurrentMissionItem(index int32) {

	request := &SetCurrentMissionItemRequest{}
	ctx := context.Background()
	request.Index = index
	response, err := s.Client.SetCurrentMissionItem(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetCurrentMissionItem grpc %v\n", err)
	}

	result := response.GetMissionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != MissionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetCurrentMissionItem")
	}

}

/*
   Check if the mission has been finished.



   Returns
   -------
   False
   IsFinished : bool
        True if the mission is finished and the last mission item has been reached


*/

func (s *ServiceImpl) IsMissionFinished() bool {
	request := &IsMissionFinishedRequest{}
	ctx := context.Background()
	response, err := s.Client.IsMissionFinished(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
	}

	result := response.GetMissionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != MissionResult_RESULT_SUCCESS {
		fmt.Printf("Error while getting IsMissionFinished")
	}

	return response.GetIsFinished()

}

/*
   Subscribe to mission progress updates.


*/
func (a *ServiceImpl) MissionProgress() {
	request := &SubscribeMissionProgressRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeMissionProgress(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	for {
		m := &MissionProgressResponse{}
		err := stream.RecvMsg(m)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Unable to receive message %v", err)
			break
		}
		fmt.Printf("message %v\n", m)
	}
}

/*
   Get whether to trigger Return-to-Launch (RTL) after mission is complete.

   Before getting this option, it needs to be set, or a mission
   needs to be downloaded.



   Returns
   -------
   False
   Enable : bool
        If true, trigger an RTL at the end of the mission


*/

func (s *ServiceImpl) GetReturnToLaunchAfterMission() bool {
	request := &GetReturnToLaunchAfterMissionRequest{}
	ctx := context.Background()
	response, err := s.Client.GetReturnToLaunchAfterMission(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
	}

	result := response.GetMissionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != MissionResult_RESULT_SUCCESS {
		fmt.Printf("Error while getting GetReturnToLaunchAfterMission")
	}

	return response.GetEnable()

}

/*
   Set whether to trigger Return-to-Launch (RTL) after the mission is complete.

   This will only take effect for the next mission upload, meaning that
   the mission may have to be uploaded again.

   Parameters
   ----------
   enable bool


*/

func (s *ServiceImpl) SetReturnToLaunchAfterMission(enable bool) {

	request := &SetReturnToLaunchAfterMissionRequest{}
	ctx := context.Background()
	request.Enable = enable
	response, err := s.Client.SetReturnToLaunchAfterMission(ctx, request)
	if err != nil {
		fmt.Printf("Error while performing SetReturnToLaunchAfterMission grpc %v\n", err)
	}

	result := response.GetMissionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != MissionResult_RESULT_SUCCESS {
		fmt.Printf("Error while extracting result for SetReturnToLaunchAfterMission")
	}

}

/*
   Import a QGroundControl (QGC) mission plan.

   The method will fail if any of the imported mission items are not supported
   by the MAVSDK API.

   Parameters
   ----------
   qgcPlanPath string

   Returns
   -------
   False
   MissionPlan : MissionPlan
        The mission plan


*/

func (s *ServiceImpl) ImportQgroundcontrolMission(qgcPlanPath string) *MissionPlan {
	request := &ImportQgroundcontrolMissionRequest{}
	ctx := context.Background()
	request.QgcPlanPath = qgcPlanPath
	response, err := s.Client.ImportQgroundcontrolMission(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
	}

	result := response.GetMissionResult()
	fmt.Printf("result %v\n", result.String())
	if result.Result != MissionResult_RESULT_SUCCESS {
		fmt.Printf("Error while getting ImportQgroundcontrolMission")
	}

	return response.GetMissionPlan()

}

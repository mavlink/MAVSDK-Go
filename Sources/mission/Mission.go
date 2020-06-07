import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client MissionServiceClient
}

    func(s *ServiceImpl)upload_mission(self, mission_plan []*MissionPlan){
     request = &UploadMissionRequest{}
         request.MissionPlan = 
            
        response, err = s.Client.UploadMission(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetMissionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading UploadMission")
        }
        
    }

       

    func(s *ServiceImpl)cancel_mission_upload(self){
     request = &CancelMissionUploadRequest{}
         response, err = s.Client.CancelMissionUpload(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetMissionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading CancelMissionUpload")
        }
        
    }

       

    func(s *ServiceImpl)download_mission(self)([]*DownloadMission){
     request = &DownloadMissionRequest{}
         response, err = s.Client.DownloadMission(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetMissionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading DownloadMission")
        }
        

        
    }

       

    func(s *ServiceImpl)cancel_mission_download(self){
     request = &CancelMissionDownloadRequest{}
         response, err = s.Client.CancelMissionDownload(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetMissionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading CancelMissionDownload")
        }
        
    }

       

    func(s *ServiceImpl)start_mission(self){
     request = &StartMissionRequest{}
         response, err = s.Client.StartMission(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetMissionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading StartMission")
        }
        
    }

       

    func(s *ServiceImpl)pause_mission(self){
     request = &PauseMissionRequest{}
         response, err = s.Client.PauseMission(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetMissionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading PauseMission")
        }
        
    }

       

    func(s *ServiceImpl)clear_mission(self){
     request = &ClearMissionRequest{}
         response, err = s.Client.ClearMission(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetMissionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading ClearMission")
        }
        
    }

       

    func(s *ServiceImpl)set_current_mission_item(self, index []*Index){
     request = &SetCurrentMissionItemRequest{}
         request.index = index
        response, err = s.Client.SetCurrentMissionItem(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetMissionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetCurrentMissionItem")
        }
        
    }

       

    func(s *ServiceImpl)is_mission_finished(self)([]*IsMissionFinished){
     request = &IsMissionFinishedRequest{}
         response, err = s.Client.IsMissionFinished(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetMissionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading IsMissionFinished")
        }
        

        return response.GetIsMissionFinished()
        
    }

       

    func (a *ServiceImpl) MissionProgress(){
    	request := &MissionProgressRequest{}
    		ctx := context.Background()
    		stream, err := a.Mission.SubscribeMissionProgress(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &mission_progressResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			fmt.Printf("message %v\n", m.GetMissionProgress())
    		}	
    }

    func(s *ServiceImpl)get_return_to_launch_after_mission(self)([]*GetReturnToLaunchAfterMission){
     request = &GetReturnToLaunchAfterMissionRequest{}
         response, err = s.Client.GetReturnToLaunchAfterMission(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetMissionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading GetReturnToLaunchAfterMission")
        }
        

        return response.GetGetReturnToLaunchAfterMission()
        
    }

       

    func(s *ServiceImpl)set_return_to_launch_after_mission(self, enable []*Enable){
     request = &SetReturnToLaunchAfterMissionRequest{}
         request.enable = enable
        response, err = s.Client.SetReturnToLaunchAfterMission(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetMissionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetReturnToLaunchAfterMission")
        }
        
    }

       

    func(s *ServiceImpl)import_qgroundcontrol_mission(self, qgc_plan_path []*QgcPlanPath)([]*ImportQgroundcontrolMission){
     request = &ImportQgroundcontrolMissionRequest{}
         request.qgc_plan_path = qgc_plan_path
        response, err = s.Client.ImportQgroundcontrolMission(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetMissionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading ImportQgroundcontrolMission")
        }
        

        
    }

       
package mission

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client MissionServiceClient
}

    func(s *ServiceImpl)UploadMission( mission_plan MissionPlan){
        request := &UploadMissionRequest{}
        ctx:= context.Background()
         request.MissionPlan = mission_plan
            
        response, err := s.Client.UploadMission(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing UploadMission grpc %v\n", err)
    	}
        
        result := response.GetMissionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for UploadMission")
        }
        
    }

       

    func(s *ServiceImpl)CancelMissionUpload(){
        request := &CancelMissionUploadRequest{}
        ctx:= context.Background()
         response, err := s.Client.CancelMissionUpload(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing CancelMissionUpload grpc %v\n", err)
    	}
        
        result := response.GetMissionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for CancelMissionUpload")
        }
        
    }

       

    func(s *ServiceImpl)DownloadMission() (*DownloadMissionResponse){
        request := &DownloadMissionRequest{}
        ctx:= context.Background()
         response, err := s.Client.DownloadMission(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result := response.GetMissionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while getting DownloadMission")
        }
        
        return response

    }

       

    func(s *ServiceImpl)CancelMissionDownload(){
        request := &CancelMissionDownloadRequest{}
        ctx:= context.Background()
         response, err := s.Client.CancelMissionDownload(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing CancelMissionDownload grpc %v\n", err)
    	}
        
        result := response.GetMissionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for CancelMissionDownload")
        }
        
    }

       

    func(s *ServiceImpl)StartMission(){
        request := &StartMissionRequest{}
        ctx:= context.Background()
         response, err := s.Client.StartMission(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing StartMission grpc %v\n", err)
    	}
        
        result := response.GetMissionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for StartMission")
        }
        
    }

       

    func(s *ServiceImpl)PauseMission(){
        request := &PauseMissionRequest{}
        ctx:= context.Background()
         response, err := s.Client.PauseMission(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing PauseMission grpc %v\n", err)
    	}
        
        result := response.GetMissionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for PauseMission")
        }
        
    }

       

    func(s *ServiceImpl)ClearMission(){
        request := &ClearMissionRequest{}
        ctx:= context.Background()
         response, err := s.Client.ClearMission(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing ClearMission grpc %v\n", err)
    	}
        
        result := response.GetMissionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for ClearMission")
        }
        
    }

       

    func(s *ServiceImpl)SetCurrentMissionItem( index Index){
        request := &SetCurrentMissionItemRequest{}
        ctx:= context.Background()
         request.Index = index
        response, err := s.Client.SetCurrentMissionItem(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetCurrentMissionItem grpc %v\n", err)
    	}
        
        result := response.GetMissionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetCurrentMissionItem")
        }
        
    }

       

    func(s *ServiceImpl)IsMissionFinished() (*IsMissionFinishedResponse){
        request := &IsMissionFinishedRequest{}
        ctx:= context.Background()
         response, err := s.Client.IsMissionFinished(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result := response.GetMissionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while getting IsMissionFinished")
        }
        
        return response

    }

       

    func (a *ServiceImpl) MissionProgress(){
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

    func(s *ServiceImpl)GetReturnToLaunchAfterMission() (*GetReturnToLaunchAfterMissionResponse){
        request := &GetReturnToLaunchAfterMissionRequest{}
        ctx:= context.Background()
         response, err := s.Client.GetReturnToLaunchAfterMission(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result := response.GetMissionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while getting GetReturnToLaunchAfterMission")
        }
        
        return response

    }

       

    func(s *ServiceImpl)SetReturnToLaunchAfterMission( enable Enable){
        request := &SetReturnToLaunchAfterMissionRequest{}
        ctx:= context.Background()
         request.Enable = enable
        response, err := s.Client.SetReturnToLaunchAfterMission(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetReturnToLaunchAfterMission grpc %v\n", err)
    	}
        
        result := response.GetMissionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetReturnToLaunchAfterMission")
        }
        
    }

       

    func(s *ServiceImpl)ImportQgroundcontrolMission( qgc_plan_path QgcPlanPath) (*ImportQgroundcontrolMissionResponse){
        request := &ImportQgroundcontrolMissionRequest{}
        ctx:= context.Background()
         request.QgcPlanPath = qgc_plan_path
        response, err := s.Client.ImportQgroundcontrolMission(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result := response.GetMissionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != MissionResult_RESULT_SUCCESS{
            fmt.Printf("Error while getting ImportQgroundcontrolMission")
        }
        
        return response

    }

       
package mission

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
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

    func(s *ServiceImpl)UploadMission(missionPlan *MissionPlan )(*UploadMissionResponse, error){
        
        request := &UploadMissionRequest{}
        ctx:= context.Background()
         request.MissionPlan = missionPlan
            
        response, err := s.Client.UploadMission(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       

     /*
         Upload a list of mission items to the system and report upload progress.

         The mission items are uploaded to a drone. Once uploaded the mission can be started and
         executed even if the connection is lost.

         Parameters
         ----------
         missionPlan *MissionPlan 
    */

    func (a *ServiceImpl) UploadMissionWithProgress(missionPlan *MissionPlan ) (<-chan  *ProgressData , error){
    		ch := make(chan  *ProgressData )
    		request := &SubscribeUploadMissionWithProgressRequest{}
    		request.MissionPlan = 
            
        ctx := context.Background()
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
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetProgressData()
    		}
    	}()	
    	return ch, nil
    }
    /*
         Cancel an ongoing mission upload.

         
    */

    func(s *ServiceImpl)CancelMissionUpload()(*CancelMissionUploadResponse, error){
        
        request := &CancelMissionUploadRequest{}
        ctx:= context.Background()
         response, err := s.Client.CancelMissionUpload(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
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


    func(s *ServiceImpl)DownloadMission() (*DownloadMissionResponse, error){
        request := &DownloadMissionRequest{}
        ctx:= context.Background()
         response, err := s.Client.DownloadMission(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       

     /*
         Download a list of mission items from the system (asynchronous) and report progress.

         Will fail if any of the downloaded mission items are not supported
         by the MAVSDK API.

         
    */

    func (a *ServiceImpl) DownloadMissionWithProgress() (<-chan  *ProgressDataOrMission , error){
    		ch := make(chan  *ProgressDataOrMission )
    		request := &SubscribeDownloadMissionWithProgressRequest{}
    		ctx := context.Background()
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
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetProgressData()
    		}
    	}()	
    	return ch, nil
    }
    /*
         Cancel an ongoing mission download.

         
    */

    func(s *ServiceImpl)CancelMissionDownload()(*CancelMissionDownloadResponse, error){
        
        request := &CancelMissionDownloadRequest{}
        ctx:= context.Background()
         response, err := s.Client.CancelMissionDownload(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Start the mission.

         A mission must be uploaded to the vehicle before this can be called.

         
    */

    func(s *ServiceImpl)StartMission()(*StartMissionResponse, error){
        
        request := &StartMissionRequest{}
        ctx:= context.Background()
         response, err := s.Client.StartMission(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Pause the mission.

         Pausing the mission puts the vehicle into
         [HOLD mode](https://docs.px4.io/en/flight_modes/hold.html).
         A multicopter should just hover at the spot while a fixedwing vehicle should loiter
         around the location where it paused.

         
    */

    func(s *ServiceImpl)PauseMission()(*PauseMissionResponse, error){
        
        request := &PauseMissionRequest{}
        ctx:= context.Background()
         response, err := s.Client.PauseMission(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Clear the mission saved on the vehicle.

         
    */

    func(s *ServiceImpl)ClearMission()(*ClearMissionResponse, error){
        
        request := &ClearMissionRequest{}
        ctx:= context.Background()
         response, err := s.Client.ClearMission(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
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

    func(s *ServiceImpl)SetCurrentMissionItem(index int32)(*SetCurrentMissionItemResponse, error){
        
        request := &SetCurrentMissionItemRequest{}
        ctx:= context.Background()
         request.Index = index
        response, err := s.Client.SetCurrentMissionItem(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Check if the mission has been finished.

         

         Returns
         -------
         False
         IsFinished : bool
              True if the mission is finished and the last mission item has been reached

         
    */


    func(s *ServiceImpl)IsMissionFinished() (*IsMissionFinishedResponse, error){
        request := &IsMissionFinishedRequest{}
        ctx:= context.Background()
         response, err := s.Client.IsMissionFinished(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       

     /*
         Subscribe to mission progress updates.

         
    */

    func (a *ServiceImpl) MissionProgress() (<-chan  *MissionProgress , error){
    		ch := make(chan  *MissionProgress )
    		request := &SubscribeMissionProgressRequest{}
    		ctx := context.Background()
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
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetMissionProgress()
    		}
    	}()	
    	return ch, nil
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


    func(s *ServiceImpl)GetReturnToLaunchAfterMission() (*GetReturnToLaunchAfterMissionResponse, error){
        request := &GetReturnToLaunchAfterMissionRequest{}
        ctx:= context.Background()
         response, err := s.Client.GetReturnToLaunchAfterMission(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Set whether to trigger Return-to-Launch (RTL) after the mission is complete.

         This will only take effect for the next mission upload, meaning that
         the mission may have to be uploaded again.

         Parameters
         ----------
         enable bool

         
    */

    func(s *ServiceImpl)SetReturnToLaunchAfterMission(enable bool)(*SetReturnToLaunchAfterMissionResponse, error){
        
        request := &SetReturnToLaunchAfterMissionRequest{}
        ctx:= context.Background()
         request.Enable = enable
        response, err := s.Client.SetReturnToLaunchAfterMission(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
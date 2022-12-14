package mission_raw

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client MissionRawServiceClient
}
    /*
         Upload a list of raw mission items to the system.

         The raw mission items are uploaded to a drone. Once uploaded the mission
         can be started and executed even if the connection is lost.

         Parameters
         ----------
         missionItems []*MissionItem

         
    */

    func(s *ServiceImpl)UploadMission(ctx context.Context, missionItems []*MissionItem)(*UploadMissionResponse, error){
        
        request := &UploadMissionRequest{}
    	request.MissionItems = missionItems
            
        response, err := s.Client.UploadMission(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Cancel an ongoing mission upload.

         
    */

    func(s *ServiceImpl)CancelMissionUpload(ctx context.Context, )(*CancelMissionUploadResponse, error){
        
        request := &CancelMissionUploadRequest{}
    	response, err := s.Client.CancelMissionUpload(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Download a list of raw mission items from the system (asynchronous).

         

         Returns
         -------
         True
         MissionItems : []*MissionItem
              The mission items

         
    */


    func(s *ServiceImpl)DownloadMission(ctx context.Context, ) (*DownloadMissionResponse, error){
        request := &DownloadMissionRequest{}
    	response, err := s.Client.DownloadMission(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Cancel an ongoing mission download.

         
    */

    func(s *ServiceImpl)CancelMissionDownload(ctx context.Context, )(*CancelMissionDownloadResponse, error){
        
        request := &CancelMissionDownloadRequest{}
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

    func(s *ServiceImpl)StartMission(ctx context.Context, )(*StartMissionResponse, error){
        
        request := &StartMissionRequest{}
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

    func(s *ServiceImpl)PauseMission(ctx context.Context, )(*PauseMissionResponse, error){
        
        request := &PauseMissionRequest{}
    	response, err := s.Client.PauseMission(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Clear the mission saved on the vehicle.

         
    */

    func(s *ServiceImpl)ClearMission(ctx context.Context, )(*ClearMissionResponse, error){
        
        request := &ClearMissionRequest{}
    	response, err := s.Client.ClearMission(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Sets the raw mission item index to go to.

         By setting the current index to 0, the mission is restarted from the beginning. If it is set
         to a specific index of a raw mission item, the mission will be set to this item.

         Parameters
         ----------
         index int32

         
    */

    func(s *ServiceImpl)SetCurrentMissionItem(ctx context.Context, index int32)(*SetCurrentMissionItemResponse, error){
        
        request := &SetCurrentMissionItemRequest{}
    	request.Index = index
        response, err := s.Client.SetCurrentMissionItem(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       

     /*
         Subscribe to mission progress updates.

         
    */

    func (a *ServiceImpl) MissionProgress(ctx context.Context, ) (<-chan  *MissionProgress , error){
    		ch := make(chan  *MissionProgress )
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
         *
         Subscribes to mission changed.

         This notification can be used to be informed if a ground station has
         been uploaded or changed by a ground station or companion computer.

         @param callback Callback to notify about change.

         
    */

    func (a *ServiceImpl) MissionChanged(ctx context.Context, ) (<-chan   bool  , error){
    		ch := make(chan   bool  )
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
    					break
    				}
    				if err != nil {
    					fmt.Printf("Unable to receive message %v", err)
    					break
    				}
    				ch <- m.GetMissionChanged()
    			}
    		}()	
    	return ch, nil
    }
    /*
         Import a QGroundControl missions in JSON .plan format.

         Supported:
         - Waypoints
         - Survey
         Not supported:
         - Structure Scan

         Parameters
         ----------
         qgcPlanPath string

         Returns
         -------
         False
         MissionImportData : MissionImportData
              The imported mission data

         
    */


    func(s *ServiceImpl)ImportQgroundcontrolMission(ctx context.Context, qgcPlanPath string) (*ImportQgroundcontrolMissionResponse, error){
        request := &ImportQgroundcontrolMissionRequest{}
    	request.QgcPlanPath = qgcPlanPath
        response, err := s.Client.ImportQgroundcontrolMission(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
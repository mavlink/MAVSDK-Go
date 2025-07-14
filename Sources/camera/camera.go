package camera

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client CameraServiceClient
}
    /*
         Take one photo.

         Parameters
         ----------
         componentId int32

         
    */

    func(s *ServiceImpl)TakePhoto(ctx context.Context, componentId int32)(*TakePhotoResponse, error){
        
        request := &TakePhotoRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.TakePhoto(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Start photo timelapse with a given interval.

         Parameters
         ----------
         componentId int32

         intervalS float32

         
    */

    func(s *ServiceImpl)StartPhotoInterval(ctx context.Context, componentId int32, intervalS float32)(*StartPhotoIntervalResponse, error){
        
        request := &StartPhotoIntervalRequest{}
    	request.ComponentId = componentId
        request.IntervalS = intervalS
        response, err := s.Client.StartPhotoInterval(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Stop a running photo timelapse.

         Parameters
         ----------
         componentId int32

         
    */

    func(s *ServiceImpl)StopPhotoInterval(ctx context.Context, componentId int32)(*StopPhotoIntervalResponse, error){
        
        request := &StopPhotoIntervalRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.StopPhotoInterval(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Start a video recording.

         Parameters
         ----------
         componentId int32

         
    */

    func(s *ServiceImpl)StartVideo(ctx context.Context, componentId int32)(*StartVideoResponse, error){
        
        request := &StartVideoRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.StartVideo(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Stop a running video recording.

         Parameters
         ----------
         componentId int32

         
    */

    func(s *ServiceImpl)StopVideo(ctx context.Context, componentId int32)(*StopVideoResponse, error){
        
        request := &StopVideoRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.StopVideo(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Start video streaming.

         Parameters
         ----------
         componentId int32

         streamId int32

         
    */

    func(s *ServiceImpl)StartVideoStreaming(ctx context.Context, componentId int32, streamId int32)(*StartVideoStreamingResponse, error){
        
        request := &StartVideoStreamingRequest{}
    	request.ComponentId = componentId
        request.StreamId = streamId
        response, err := s.Client.StartVideoStreaming(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Stop current video streaming.

         Parameters
         ----------
         componentId int32

         streamId int32

         
    */

    func(s *ServiceImpl)StopVideoStreaming(ctx context.Context, componentId int32, streamId int32)(*StopVideoStreamingResponse, error){
        
        request := &StopVideoStreamingRequest{}
    	request.ComponentId = componentId
        request.StreamId = streamId
        response, err := s.Client.StopVideoStreaming(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set camera mode.

         Parameters
         ----------
         componentId int32

         mode *Mode 
            

         
    */

    func(s *ServiceImpl)SetMode(ctx context.Context, componentId int32, mode *Mode )(*SetModeResponse, error){
        
        request := &SetModeRequest{}
    	request.ComponentId = componentId
        request.Mode = *mode
        response, err := s.Client.SetMode(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         List photos available on the camera.

         Note that this might need to be called initially to set the PhotosRange accordingly.
         Once set to 'all' rather than 'since connection', it will try to request the previous
         images over time.

         Parameters
         ----------
         componentId int32photosRange *PhotosRange 
            

         Returns
         -------
         True
         CaptureInfos : []*CaptureInfo
              List of capture infos (representing the photos)

         
    */


    func(s *ServiceImpl)ListPhotos(ctx context.Context, componentId int32, photosRange *PhotosRange ) (*ListPhotosResponse, error){
        request := &ListPhotosRequest{}
    	request.ComponentId = componentId
        request.PhotosRange = *photosRange
        response, err := s.Client.ListPhotos(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       

     /*
         Subscribe to list of cameras.

         This allows to find out what cameras are connected to the system.
         Based on the camera ID, we can then address a specific camera.

         
    */

    func (a *ServiceImpl) CameraList(ctx context.Context, ) (<-chan  *CameraList , error){
    		ch := make(chan  *CameraList )
    		request := &SubscribeCameraListRequest{}
    		stream, err := a.Client.SubscribeCameraList(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &CameraListResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive CameraList messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetCameraList()
    			}
    		}()	
    	return ch, nil
    }

     /*
         Subscribe to camera mode updates.

         
    */

    func (a *ServiceImpl) Mode(ctx context.Context, ) (<-chan  *ModeUpdate , error){
    		ch := make(chan  *ModeUpdate )
    		request := &SubscribeModeRequest{}
    		stream, err := a.Client.SubscribeMode(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &ModeResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive Mode messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetUpdate()
    			}
    		}()	
    	return ch, nil
    }
    /*
         Get camera mode.

         Parameters
         ----------
         componentId int32

         Returns
         -------
         False
         Mode : Mode
              Mode

         
    */


    func(s *ServiceImpl)GetMode(ctx context.Context, componentId int32) (*GetModeResponse, error){
        request := &GetModeRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.GetMode(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       

     /*
         Subscribe to video stream info updates.

         
    */

    func (a *ServiceImpl) VideoStreamInfo(ctx context.Context, ) (<-chan  *VideoStreamUpdate , error){
    		ch := make(chan  *VideoStreamUpdate )
    		request := &SubscribeVideoStreamInfoRequest{}
    		stream, err := a.Client.SubscribeVideoStreamInfo(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &VideoStreamInfoResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive VideoStreamInfo messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetUpdate()
    			}
    		}()	
    	return ch, nil
    }
    /*
         Get video stream info.

         Parameters
         ----------
         componentId int32

         Returns
         -------
         False
         VideoStreamInfo : VideoStreamInfo
              Video stream info

         
    */


    func(s *ServiceImpl)GetVideoStreamInfo(ctx context.Context, componentId int32) (*GetVideoStreamInfoResponse, error){
        request := &GetVideoStreamInfoRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.GetVideoStreamInfo(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       

     /*
         Subscribe to capture info updates.

         
    */

    func (a *ServiceImpl) CaptureInfo(ctx context.Context, ) (<-chan  *CaptureInfo , error){
    		ch := make(chan  *CaptureInfo )
    		request := &SubscribeCaptureInfoRequest{}
    		stream, err := a.Client.SubscribeCaptureInfo(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &CaptureInfoResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive CaptureInfo messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetCaptureInfo()
    			}
    		}()	
    	return ch, nil
    }

     /*
         Subscribe to camera's storage status updates.

         
    */

    func (a *ServiceImpl) Storage(ctx context.Context, ) (<-chan  *StorageUpdate , error){
    		ch := make(chan  *StorageUpdate )
    		request := &SubscribeStorageRequest{}
    		stream, err := a.Client.SubscribeStorage(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &StorageResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive Storage messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetUpdate()
    			}
    		}()	
    	return ch, nil
    }
    /*
         Get camera's storage status.

         Parameters
         ----------
         componentId int32

         Returns
         -------
         False
         Storage : Storage
              Camera's storage status

         
    */


    func(s *ServiceImpl)GetStorage(ctx context.Context, componentId int32) (*GetStorageResponse, error){
        request := &GetStorageRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.GetStorage(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       

     /*
         Get the list of current camera settings.

         
    */

    func (a *ServiceImpl) CurrentSettings(ctx context.Context, ) (<-chan  *CurrentSettingsUpdate , error){
    		ch := make(chan  *CurrentSettingsUpdate )
    		request := &SubscribeCurrentSettingsRequest{}
    		stream, err := a.Client.SubscribeCurrentSettings(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &CurrentSettingsResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive CurrentSettings messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetUpdate()
    			}
    		}()	
    	return ch, nil
    }
    /*
         Get current settings.

         Parameters
         ----------
         componentId int32

         Returns
         -------
         True
         CurrentSettings : []*Setting
              List of current settings

         
    */


    func(s *ServiceImpl)GetCurrentSettings(ctx context.Context, componentId int32) (*GetCurrentSettingsResponse, error){
        request := &GetCurrentSettingsRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.GetCurrentSettings(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       

     /*
         Get the list of settings that can be changed.

         
    */

    func (a *ServiceImpl) PossibleSettingOptions(ctx context.Context, ) (<-chan  *PossibleSettingOptionsUpdate , error){
    		ch := make(chan  *PossibleSettingOptionsUpdate )
    		request := &SubscribePossibleSettingOptionsRequest{}
    		stream, err := a.Client.SubscribePossibleSettingOptions(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &PossibleSettingOptionsResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive PossibleSettingOptions messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetUpdate()
    			}
    		}()	
    	return ch, nil
    }
    /*
         Get possible setting options.

         Parameters
         ----------
         componentId int32

         Returns
         -------
         True
         SettingOptions : []*SettingOptions
              List of settings that can be changed

         
    */


    func(s *ServiceImpl)GetPossibleSettingOptions(ctx context.Context, componentId int32) (*GetPossibleSettingOptionsResponse, error){
        request := &GetPossibleSettingOptionsRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.GetPossibleSettingOptions(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Set a setting to some value.

         Only setting_id of setting and option_id of option needs to be set.

         Parameters
         ----------
         componentId int32

         setting *Setting 
            

         
    */

    func(s *ServiceImpl)SetSetting(ctx context.Context, componentId int32, setting *Setting )(*SetSettingResponse, error){
        
        request := &SetSettingRequest{}
    	request.ComponentId = componentId
        request.Setting = setting
            
        response, err := s.Client.SetSetting(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Get a setting.

         Only setting_id of setting needs to be set.

         Parameters
         ----------
         componentId int32setting *Setting 
            

         Returns
         -------
         False
         Setting : Setting
              Setting

         
    */


    func(s *ServiceImpl)GetSetting(ctx context.Context, componentId int32, setting *Setting ) (*GetSettingResponse, error){
        request := &GetSettingRequest{}
    	request.ComponentId = componentId
        request.Setting = setting
            
        response, err := s.Client.GetSetting(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Format storage (e.g. SD card) in camera.

         This will delete all content of the camera storage!

         Parameters
         ----------
         componentId int32

         storageId int32

         
    */

    func(s *ServiceImpl)FormatStorage(ctx context.Context, componentId int32, storageId int32)(*FormatStorageResponse, error){
        
        request := &FormatStorageRequest{}
    	request.ComponentId = componentId
        request.StorageId = storageId
        response, err := s.Client.FormatStorage(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Reset all settings in camera.

         This will reset all camera settings to default value

         Parameters
         ----------
         componentId int32

         
    */

    func(s *ServiceImpl)ResetSettings(ctx context.Context, componentId int32)(*ResetSettingsResponse, error){
        
        request := &ResetSettingsRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.ResetSettings(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Start zooming in.

         Parameters
         ----------
         componentId int32

         
    */

    func(s *ServiceImpl)ZoomInStart(ctx context.Context, componentId int32)(*ZoomInStartResponse, error){
        
        request := &ZoomInStartRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.ZoomInStart(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Start zooming out.

         Parameters
         ----------
         componentId int32

         
    */

    func(s *ServiceImpl)ZoomOutStart(ctx context.Context, componentId int32)(*ZoomOutStartResponse, error){
        
        request := &ZoomOutStartRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.ZoomOutStart(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Stop zooming.

         Parameters
         ----------
         componentId int32

         
    */

    func(s *ServiceImpl)ZoomStop(ctx context.Context, componentId int32)(*ZoomStopResponse, error){
        
        request := &ZoomStopRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.ZoomStop(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Zoom to value as proportion of full camera range (percentage between 0.0 and 100.0).

         Parameters
         ----------
         componentId int32

         range float32

         
    */

    func(s *ServiceImpl)ZoomRange(ctx context.Context, componentId int32, range float32)(*ZoomRangeResponse, error){
        
        request := &ZoomRangeRequest{}
    	request.ComponentId = componentId
        request.Range = range
        response, err := s.Client.ZoomRange(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Track point.

         Parameters
         ----------
         componentId int32

         pointX float32

         pointY float32

         radius float32

         
    */

    func(s *ServiceImpl)TrackPoint(ctx context.Context, componentId int32, pointX float32, pointY float32, radius float32)(*TrackPointResponse, error){
        
        request := &TrackPointRequest{}
    	request.ComponentId = componentId
        request.PointX = pointX
        request.PointY = pointY
        request.Radius = radius
        response, err := s.Client.TrackPoint(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Track rectangle.

         Parameters
         ----------
         componentId int32

         topLeftX float32

         topLeftY float32

         bottomRightX float32

         bottomRightY float32

         
    */

    func(s *ServiceImpl)TrackRectangle(ctx context.Context, componentId int32, topLeftX float32, topLeftY float32, bottomRightX float32, bottomRightY float32)(*TrackRectangleResponse, error){
        
        request := &TrackRectangleRequest{}
    	request.ComponentId = componentId
        request.TopLeftX = topLeftX
        request.TopLeftY = topLeftY
        request.BottomRightX = bottomRightX
        request.BottomRightY = bottomRightY
        response, err := s.Client.TrackRectangle(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Stop tracking.

         Parameters
         ----------
         componentId int32

         
    */

    func(s *ServiceImpl)TrackStop(ctx context.Context, componentId int32)(*TrackStopResponse, error){
        
        request := &TrackStopRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.TrackStop(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Start focusing in.

         Parameters
         ----------
         componentId int32

         
    */

    func(s *ServiceImpl)FocusInStart(ctx context.Context, componentId int32)(*FocusInStartResponse, error){
        
        request := &FocusInStartRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.FocusInStart(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Start focusing out.

         Parameters
         ----------
         componentId int32

         
    */

    func(s *ServiceImpl)FocusOutStart(ctx context.Context, componentId int32)(*FocusOutStartResponse, error){
        
        request := &FocusOutStartRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.FocusOutStart(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Stop focus.

         Parameters
         ----------
         componentId int32

         
    */

    func(s *ServiceImpl)FocusStop(ctx context.Context, componentId int32)(*FocusStopResponse, error){
        
        request := &FocusStopRequest{}
    	request.ComponentId = componentId
        response, err := s.Client.FocusStop(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Focus with range value of full range (value between 0.0 and 100.0).

         Parameters
         ----------
         componentId int32

         range float32

         
    */

    func(s *ServiceImpl)FocusRange(ctx context.Context, componentId int32, range float32)(*FocusRangeResponse, error){
        
        request := &FocusRangeRequest{}
    	request.ComponentId = componentId
        request.Range = range
        response, err := s.Client.FocusRange(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
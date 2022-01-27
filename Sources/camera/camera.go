package camera

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct {
	Client CameraServiceClient
}

/*
   Prepare the camera plugin (e.g. download the camera definition, etc).


*/

func (s *ServiceImpl) Prepare() (*PrepareResponse, error) {

	request := &PrepareRequest{}
	ctx := context.Background()
	response, err := s.Client.Prepare(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Take one photo.


*/

func (s *ServiceImpl) TakePhoto() (*TakePhotoResponse, error) {

	request := &TakePhotoRequest{}
	ctx := context.Background()
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
   intervalS float32


*/

func (s *ServiceImpl) StartPhotoInterval(intervalS float32) (*StartPhotoIntervalResponse, error) {

	request := &StartPhotoIntervalRequest{}
	ctx := context.Background()
	request.IntervalS = intervalS
	response, err := s.Client.StartPhotoInterval(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Stop a running photo timelapse.


*/

func (s *ServiceImpl) StopPhotoInterval() (*StopPhotoIntervalResponse, error) {

	request := &StopPhotoIntervalRequest{}
	ctx := context.Background()
	response, err := s.Client.StopPhotoInterval(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Start a video recording.


*/

func (s *ServiceImpl) StartVideo() (*StartVideoResponse, error) {

	request := &StartVideoRequest{}
	ctx := context.Background()
	response, err := s.Client.StartVideo(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Stop a running video recording.


*/

func (s *ServiceImpl) StopVideo() (*StopVideoResponse, error) {

	request := &StopVideoRequest{}
	ctx := context.Background()
	response, err := s.Client.StopVideo(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Start video streaming.


*/

func (s *ServiceImpl) StartVideoStreaming() (*StartVideoStreamingResponse, error) {

	request := &StartVideoStreamingRequest{}
	ctx := context.Background()
	response, err := s.Client.StartVideoStreaming(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Stop current video streaming.


*/

func (s *ServiceImpl) StopVideoStreaming() (*StopVideoStreamingResponse, error) {

	request := &StopVideoStreamingRequest{}
	ctx := context.Background()
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
   mode *Mode



*/

func (s *ServiceImpl) SetMode(mode *Mode) (*SetModeResponse, error) {

	request := &SetModeRequest{}
	ctx := context.Background()
	request.Mode = *mode
	response, err := s.Client.SetMode(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   List photos available on the camera.

   Parameters
   ----------
   photosRange *PhotosRange


   Returns
   -------
   True
   CaptureInfos : []*CaptureInfo
        List of capture infos (representing the photos)


*/

func (s *ServiceImpl) ListPhotos(photosRange *PhotosRange) (*ListPhotosResponse, error) {
	request := &ListPhotosRequest{}
	ctx := context.Background()
	request.PhotosRange = *photosRange
	response, err := s.Client.ListPhotos(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

/*
   Subscribe to camera mode updates.


*/

func (a *ServiceImpl) Mode() (<-chan Mode, error) {
	ch := make(chan Mode)
	request := &SubscribeModeRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetMode()
		}
	}()
	return ch, nil
}

/*
   Subscribe to camera information updates.


*/

func (a *ServiceImpl) Information() (<-chan *Information, error) {
	ch := make(chan *Information)
	request := &SubscribeInformationRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeInformation(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &InformationResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetInformation()
		}
	}()
	return ch, nil
}

/*
   Subscribe to video stream info updates.


*/

func (a *ServiceImpl) VideoStreamInfo() (<-chan *VideoStreamInfo, error) {
	ch := make(chan *VideoStreamInfo)
	request := &SubscribeVideoStreamInfoRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetVideoStreamInfo()
		}
	}()
	return ch, nil
}

/*
   Subscribe to capture info updates.


*/

func (a *ServiceImpl) CaptureInfo() (<-chan *CaptureInfo, error) {
	ch := make(chan *CaptureInfo)
	request := &SubscribeCaptureInfoRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetCaptureInfo()
		}
	}()
	return ch, nil
}

/*
   Subscribe to camera status updates.


*/

func (a *ServiceImpl) Status() (<-chan *Status, error) {
	ch := make(chan *Status)
	request := &SubscribeStatusRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &StatusResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetCameraStatus()
		}
	}()
	return ch, nil
}

/*
   Get the list of current camera settings.


*/

func (a *ServiceImpl) CurrentSettings() (<-chan []*Setting, error) {
	ch := make(chan []*Setting)
	request := &SubscribeCurrentSettingsRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetCurrentSettings()
		}
	}()
	return ch, nil
}

/*
   Get the list of settings that can be changed.


*/

func (a *ServiceImpl) PossibleSettingOptions() (<-chan []*SettingOptions, error) {
	ch := make(chan []*SettingOptions)
	request := &SubscribePossibleSettingOptionsRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetSettingOptions()
		}
	}()
	return ch, nil
}

/*
   Set a setting to some value.

   Only setting_id of setting and option_id of option needs to be set.

   Parameters
   ----------
   setting *Setting



*/

func (s *ServiceImpl) SetSetting(setting *Setting) (*SetSettingResponse, error) {

	request := &SetSettingRequest{}
	ctx := context.Background()
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
   setting *Setting


   Returns
   -------
   False
   Setting : Setting
        Setting


*/

func (s *ServiceImpl) GetSetting(setting *Setting) (*GetSettingResponse, error) {
	request := &GetSettingRequest{}
	ctx := context.Background()
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


*/

func (s *ServiceImpl) FormatStorage() (*FormatStorageResponse, error) {

	request := &FormatStorageRequest{}
	ctx := context.Background()
	response, err := s.Client.FormatStorage(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

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

func (s *ServiceImpl) Prepare(ctx context.Context) (*PrepareResponse, error) {

	request := &PrepareRequest{}
	response, err := s.Client.Prepare(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Take one photo.


*/

func (s *ServiceImpl) TakePhoto(ctx context.Context) (*TakePhotoResponse, error) {

	request := &TakePhotoRequest{}
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

func (s *ServiceImpl) StartPhotoInterval(ctx context.Context, intervalS float32) (*StartPhotoIntervalResponse, error) {

	request := &StartPhotoIntervalRequest{}
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

func (s *ServiceImpl) StopPhotoInterval(ctx context.Context) (*StopPhotoIntervalResponse, error) {

	request := &StopPhotoIntervalRequest{}
	response, err := s.Client.StopPhotoInterval(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Start a video recording.


*/

func (s *ServiceImpl) StartVideo(ctx context.Context) (*StartVideoResponse, error) {

	request := &StartVideoRequest{}
	response, err := s.Client.StartVideo(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Stop a running video recording.


*/

func (s *ServiceImpl) StopVideo(ctx context.Context) (*StopVideoResponse, error) {

	request := &StopVideoRequest{}
	response, err := s.Client.StopVideo(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Start video streaming.


*/

func (s *ServiceImpl) StartVideoStreaming(ctx context.Context) (*StartVideoStreamingResponse, error) {

	request := &StartVideoStreamingRequest{}
	response, err := s.Client.StartVideoStreaming(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Stop current video streaming.


*/

func (s *ServiceImpl) StopVideoStreaming(ctx context.Context) (*StopVideoStreamingResponse, error) {

	request := &StopVideoStreamingRequest{}
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

func (s *ServiceImpl) SetMode(ctx context.Context, mode *Mode) (*SetModeResponse, error) {

	request := &SetModeRequest{}
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

func (s *ServiceImpl) ListPhotos(ctx context.Context, photosRange *PhotosRange) (*ListPhotosResponse, error) {
	request := &ListPhotosRequest{}
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

func (a *ServiceImpl) Mode(ctx context.Context) (<-chan Mode, error) {
	ch := make(chan Mode)
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

func (a *ServiceImpl) Information(ctx context.Context) (<-chan *Information, error) {
	ch := make(chan *Information)
	request := &SubscribeInformationRequest{}
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

func (a *ServiceImpl) VideoStreamInfo(ctx context.Context) (<-chan *VideoStreamInfo, error) {
	ch := make(chan *VideoStreamInfo)
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

func (a *ServiceImpl) CaptureInfo(ctx context.Context) (<-chan *CaptureInfo, error) {
	ch := make(chan *CaptureInfo)
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

func (a *ServiceImpl) Status(ctx context.Context) (<-chan *Status, error) {
	ch := make(chan *Status)
	request := &SubscribeStatusRequest{}
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

func (a *ServiceImpl) CurrentSettings(ctx context.Context) (<-chan []*Setting, error) {
	ch := make(chan []*Setting)
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

func (a *ServiceImpl) PossibleSettingOptions(ctx context.Context) (<-chan []*SettingOptions, error) {
	ch := make(chan []*SettingOptions)
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

func (s *ServiceImpl) SetSetting(ctx context.Context, setting *Setting) (*SetSettingResponse, error) {

	request := &SetSettingRequest{}
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

func (s *ServiceImpl) GetSetting(ctx context.Context, setting *Setting) (*GetSettingResponse, error) {
	request := &GetSettingRequest{}
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

func (s *ServiceImpl) FormatStorage(ctx context.Context) (*FormatStorageResponse, error) {

	request := &FormatStorageRequest{}
	response, err := s.Client.FormatStorage(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

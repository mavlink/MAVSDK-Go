package camera

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client CameraServiceClient
}

/*
Prepare Prepare the camera plugin (e.g. download the camera definition, etc).
*/
func (s *ServiceImpl) Prepare(
	ctx context.Context,

) (*PrepareResponse, error) {
	request := &PrepareRequest{}
	response, err := s.Client.Prepare(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
TakePhoto Take one photo.
*/
func (s *ServiceImpl) TakePhoto(
	ctx context.Context,

) (*TakePhotoResponse, error) {
	request := &TakePhotoRequest{}
	response, err := s.Client.TakePhoto(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
StartPhotoInterval Start photo timelapse with a given interval.
*/
func (s *ServiceImpl) StartPhotoInterval(
	ctx context.Context,
	intervalS float32,

) (*StartPhotoIntervalResponse, error) {
	request := &StartPhotoIntervalRequest{
		IntervalS: intervalS,
	}
	response, err := s.Client.StartPhotoInterval(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
StopPhotoInterval Stop a running photo timelapse.
*/
func (s *ServiceImpl) StopPhotoInterval(
	ctx context.Context,

) (*StopPhotoIntervalResponse, error) {
	request := &StopPhotoIntervalRequest{}
	response, err := s.Client.StopPhotoInterval(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
StartVideo Start a video recording.
*/
func (s *ServiceImpl) StartVideo(
	ctx context.Context,

) (*StartVideoResponse, error) {
	request := &StartVideoRequest{}
	response, err := s.Client.StartVideo(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
StopVideo Stop a running video recording.
*/
func (s *ServiceImpl) StopVideo(
	ctx context.Context,

) (*StopVideoResponse, error) {
	request := &StopVideoRequest{}
	response, err := s.Client.StopVideo(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
StartVideoStreaming Start video streaming.
*/
func (s *ServiceImpl) StartVideoStreaming(
	ctx context.Context,
	streamId int32,

) (*StartVideoStreamingResponse, error) {
	request := &StartVideoStreamingRequest{
		StreamId: streamId,
	}
	response, err := s.Client.StartVideoStreaming(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
StopVideoStreaming Stop current video streaming.
*/
func (s *ServiceImpl) StopVideoStreaming(
	ctx context.Context,
	streamId int32,

) (*StopVideoStreamingResponse, error) {
	request := &StopVideoStreamingRequest{
		StreamId: streamId,
	}
	response, err := s.Client.StopVideoStreaming(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetMode Set camera mode.
*/
func (s *ServiceImpl) SetMode(
	ctx context.Context,
	mode *Mode,

) (*SetModeResponse, error) {
	request := &SetModeRequest{
		Mode: *mode,
	}
	response, err := s.Client.SetMode(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ListPhotos List photos available on the camera.
*/
func (s *ServiceImpl) ListPhotos(
	ctx context.Context,
	photosRange *PhotosRange,

) (*ListPhotosResponse, error) {
	request := &ListPhotosRequest{
		PhotosRange: *photosRange,
	}
	response, err := s.Client.ListPhotos(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
Mode Subscribe to camera mode updates.
*/
func (a *ServiceImpl) Mode(
	ctx context.Context,

) (<-chan Mode, error) {
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
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Mode messages, err: %v", err)
			}
			ch <- m.GetMode()
		}
	}()
	return ch, nil
}

/*
Information Subscribe to camera information updates.
*/
func (a *ServiceImpl) Information(
	ctx context.Context,

) (<-chan *Information, error) {
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
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Information messages, err: %v", err)
			}
			ch <- m.GetInformation()
		}
	}()
	return ch, nil
}

/*
VideoStreamInfo Subscribe to video stream info updates.
*/
func (a *ServiceImpl) VideoStreamInfo(
	ctx context.Context,

) (<-chan *VideoStreamInfo, error) {
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
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive VideoStreamInfo messages, err: %v", err)
			}
			ch <- m.GetVideoStreamInfo()
		}
	}()
	return ch, nil
}

/*
CaptureInfo Subscribe to capture info updates.
*/
func (a *ServiceImpl) CaptureInfo(
	ctx context.Context,

) (<-chan *CaptureInfo, error) {
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
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive CaptureInfo messages, err: %v", err)
			}
			ch <- m.GetCaptureInfo()
		}
	}()
	return ch, nil
}

/*
Status Subscribe to camera status updates.
*/
func (a *ServiceImpl) Status(
	ctx context.Context,

) (<-chan *Status, error) {
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
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive Status messages, err: %v", err)
			}
			ch <- m.GetCameraStatus()
		}
	}()
	return ch, nil
}

/*
CurrentSettings Get the list of current camera settings.
*/
func (a *ServiceImpl) CurrentSettings(
	ctx context.Context,

) (<-chan []*Setting, error) {
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
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive CurrentSettings messages, err: %v", err)
			}
			ch <- m.GetCurrentSettings()
		}
	}()
	return ch, nil
}

/*
PossibleSettingOptions Get the list of settings that can be changed.
*/
func (a *ServiceImpl) PossibleSettingOptions(
	ctx context.Context,

) (<-chan []*SettingOptions, error) {
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
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive PossibleSettingOptions messages, err: %v", err)
			}
			ch <- m.GetSettingOptions()
		}
	}()
	return ch, nil
}

/*
SetSetting Set a setting to some value.

	Only setting_id of setting and option_id of option needs to be set.
*/
func (s *ServiceImpl) SetSetting(
	ctx context.Context,
	setting *Setting,

) (*SetSettingResponse, error) {
	request := &SetSettingRequest{
		Setting: setting,
	}
	response, err := s.Client.SetSetting(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
GetSetting Get a setting.

	Only setting_id of setting needs to be set.
*/
func (s *ServiceImpl) GetSetting(
	ctx context.Context,
	setting *Setting,

) (*GetSettingResponse, error) {
	request := &GetSettingRequest{
		Setting: setting,
	}
	response, err := s.Client.GetSetting(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
FormatStorage Format storage (e.g. SD card) in camera.

	This will delete all content of the camera storage!
*/
func (s *ServiceImpl) FormatStorage(
	ctx context.Context,
	storageId int32,

) (*FormatStorageResponse, error) {
	request := &FormatStorageRequest{
		StorageId: storageId,
	}
	response, err := s.Client.FormatStorage(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SelectCamera Select current camera .

	Bind the plugin instance to a specific camera_id
*/
func (s *ServiceImpl) SelectCamera(
	ctx context.Context,
	cameraId int32,

) (*SelectCameraResponse, error) {
	request := &SelectCameraRequest{
		CameraId: cameraId,
	}
	response, err := s.Client.SelectCamera(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ResetSettings Reset all settings in camera.

	This will reset all camera settings to default value
*/
func (s *ServiceImpl) ResetSettings(
	ctx context.Context,

) (*ResetSettingsResponse, error) {
	request := &ResetSettingsRequest{}
	response, err := s.Client.ResetSettings(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ZoomInStart Start zooming in.
*/
func (s *ServiceImpl) ZoomInStart(
	ctx context.Context,

) (*ZoomInStartResponse, error) {
	request := &ZoomInStartRequest{}
	response, err := s.Client.ZoomInStart(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ZoomOutStart Start zooming out.
*/
func (s *ServiceImpl) ZoomOutStart(
	ctx context.Context,

) (*ZoomOutStartResponse, error) {
	request := &ZoomOutStartRequest{}
	response, err := s.Client.ZoomOutStart(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ZoomStop Stop zooming.
*/
func (s *ServiceImpl) ZoomStop(
	ctx context.Context,

) (*ZoomStopResponse, error) {
	request := &ZoomStopRequest{}
	response, err := s.Client.ZoomStop(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ZoomRange Zoom to value as proportion of full camera range (percentage between 0.0 and 100.0).
*/
func (s *ServiceImpl) ZoomRange(
	ctx context.Context,

	rangeVar float32,

) (*ZoomRangeResponse, error) {
	request := &ZoomRangeRequest{

		Range: rangeVar,
	}
	response, err := s.Client.ZoomRange(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
TrackPoint Track point.
*/
func (s *ServiceImpl) TrackPoint(
	ctx context.Context,
	pointX float32,
	pointY float32,
	radius float32,

) (*TrackPointResponse, error) {
	request := &TrackPointRequest{
		PointX: pointX,
		PointY: pointY,
		Radius: radius,
	}
	response, err := s.Client.TrackPoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
TrackRectangle Track rectangle.
*/
func (s *ServiceImpl) TrackRectangle(
	ctx context.Context,
	topLeftX float32,
	topLeftY float32,
	bottomRightX float32,
	bottomRightY float32,

) (*TrackRectangleResponse, error) {
	request := &TrackRectangleRequest{
		TopLeftX:     topLeftX,
		TopLeftY:     topLeftY,
		BottomRightX: bottomRightX,
		BottomRightY: bottomRightY,
	}
	response, err := s.Client.TrackRectangle(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
TrackStop Stop tracking.
*/
func (s *ServiceImpl) TrackStop(
	ctx context.Context,

) (*TrackStopResponse, error) {
	request := &TrackStopRequest{}
	response, err := s.Client.TrackStop(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
FocusInStart Start focusing in.
*/
func (s *ServiceImpl) FocusInStart(
	ctx context.Context,

) (*FocusInStartResponse, error) {
	request := &FocusInStartRequest{}
	response, err := s.Client.FocusInStart(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
FocusOutStart Start focusing out.
*/
func (s *ServiceImpl) FocusOutStart(
	ctx context.Context,

) (*FocusOutStartResponse, error) {
	request := &FocusOutStartRequest{}
	response, err := s.Client.FocusOutStart(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
FocusStop Stop focus.
*/
func (s *ServiceImpl) FocusStop(
	ctx context.Context,

) (*FocusStopResponse, error) {
	request := &FocusStopRequest{}
	response, err := s.Client.FocusStop(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
FocusRange Focus with range value of full range (value between 0.0 and 100.0).
*/
func (s *ServiceImpl) FocusRange(
	ctx context.Context,

	rangeVar float32,

) (*FocusRangeResponse, error) {
	request := &FocusRangeRequest{

		Range: rangeVar,
	}
	response, err := s.Client.FocusRange(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

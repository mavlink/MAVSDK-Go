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
TakePhoto Take one photo.
*/
func (s *ServiceImpl) TakePhoto(
	ctx context.Context,
	componentId int32,

) (*TakePhotoResponse, error) {
	request := &TakePhotoRequest{
		ComponentId: componentId,
	}
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
	componentId int32,
	intervalS float32,

) (*StartPhotoIntervalResponse, error) {
	request := &StartPhotoIntervalRequest{
		ComponentId: componentId,
		IntervalS:   intervalS,
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
	componentId int32,

) (*StopPhotoIntervalResponse, error) {
	request := &StopPhotoIntervalRequest{
		ComponentId: componentId,
	}
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
	componentId int32,

) (*StartVideoResponse, error) {
	request := &StartVideoRequest{
		ComponentId: componentId,
	}
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
	componentId int32,

) (*StopVideoResponse, error) {
	request := &StopVideoRequest{
		ComponentId: componentId,
	}
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
	componentId int32,
	streamId int32,

) (*StartVideoStreamingResponse, error) {
	request := &StartVideoStreamingRequest{
		ComponentId: componentId,
		StreamId:    streamId,
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
	componentId int32,
	streamId int32,

) (*StopVideoStreamingResponse, error) {
	request := &StopVideoStreamingRequest{
		ComponentId: componentId,
		StreamId:    streamId,
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
	componentId int32,
	mode *Mode,

) (*SetModeResponse, error) {
	request := &SetModeRequest{
		ComponentId: componentId,
		Mode:        *mode,
	}
	response, err := s.Client.SetMode(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ListPhotos List photos available on the camera.

	Note that this might need to be called initially to set the PhotosRange accordingly.
	Once set to 'all' rather than 'since connection', it will try to request the previous
	images over time.
*/
func (s *ServiceImpl) ListPhotos(
	ctx context.Context,
	componentId int32,
	photosRange *PhotosRange,

) (*ListPhotosResponse, error) {
	request := &ListPhotosRequest{
		ComponentId: componentId,
		PhotosRange: *photosRange,
	}
	response, err := s.Client.ListPhotos(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
CameraList Subscribe to list of cameras.

	This allows to find out what cameras are connected to the system.
	Based on the camera ID, we can then address a specific camera.
*/
func (a *ServiceImpl) CameraList(
	ctx context.Context,

) (<-chan *CameraList, error) {
	ch := make(chan *CameraList)
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
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive CameraList messages, err: %v", err)
			}
			ch <- m.GetCameraList()
		}
	}()
	return ch, nil
}

/*
Mode Subscribe to camera mode updates.
*/
func (a *ServiceImpl) Mode(
	ctx context.Context,

) (<-chan *ModeUpdate, error) {
	ch := make(chan *ModeUpdate)
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
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive Mode messages, err: %v", err)
			}
			ch <- m.GetUpdate()
		}
	}()
	return ch, nil
}

/*
GetMode Get camera mode.
*/
func (s *ServiceImpl) GetMode(
	ctx context.Context,
	componentId int32,

) (*GetModeResponse, error) {
	request := &GetModeRequest{
		ComponentId: componentId,
	}
	response, err := s.Client.GetMode(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
VideoStreamInfo Subscribe to video stream info updates.
*/
func (a *ServiceImpl) VideoStreamInfo(
	ctx context.Context,

) (<-chan *VideoStreamUpdate, error) {
	ch := make(chan *VideoStreamUpdate)
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
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive VideoStreamInfo messages, err: %v", err)
			}
			ch <- m.GetUpdate()
		}
	}()
	return ch, nil
}

/*
GetVideoStreamInfo Get video stream info.
*/
func (s *ServiceImpl) GetVideoStreamInfo(
	ctx context.Context,
	componentId int32,

) (*GetVideoStreamInfoResponse, error) {
	request := &GetVideoStreamInfoRequest{
		ComponentId: componentId,
	}
	response, err := s.Client.GetVideoStreamInfo(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
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
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
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
Storage Subscribe to camera's storage status updates.
*/
func (a *ServiceImpl) Storage(
	ctx context.Context,

) (<-chan *StorageUpdate, error) {
	ch := make(chan *StorageUpdate)
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
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive Storage messages, err: %v", err)
			}
			ch <- m.GetUpdate()
		}
	}()
	return ch, nil
}

/*
GetStorage Get camera's storage status.
*/
func (s *ServiceImpl) GetStorage(
	ctx context.Context,
	componentId int32,

) (*GetStorageResponse, error) {
	request := &GetStorageRequest{
		ComponentId: componentId,
	}
	response, err := s.Client.GetStorage(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
CurrentSettings Get the list of current camera settings.
*/
func (a *ServiceImpl) CurrentSettings(
	ctx context.Context,

) (<-chan *CurrentSettingsUpdate, error) {
	ch := make(chan *CurrentSettingsUpdate)
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
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive CurrentSettings messages, err: %v", err)
			}
			ch <- m.GetUpdate()
		}
	}()
	return ch, nil
}

/*
GetCurrentSettings Get current settings.
*/
func (s *ServiceImpl) GetCurrentSettings(
	ctx context.Context,
	componentId int32,

) (*GetCurrentSettingsResponse, error) {
	request := &GetCurrentSettingsRequest{
		ComponentId: componentId,
	}
	response, err := s.Client.GetCurrentSettings(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
PossibleSettingOptions Get the list of settings that can be changed.
*/
func (a *ServiceImpl) PossibleSettingOptions(
	ctx context.Context,

) (<-chan *PossibleSettingOptionsUpdate, error) {
	ch := make(chan *PossibleSettingOptionsUpdate)
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
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive PossibleSettingOptions messages, err: %v", err)
			}
			ch <- m.GetUpdate()
		}
	}()
	return ch, nil
}

/*
GetPossibleSettingOptions Get possible setting options.
*/
func (s *ServiceImpl) GetPossibleSettingOptions(
	ctx context.Context,
	componentId int32,

) (*GetPossibleSettingOptionsResponse, error) {
	request := &GetPossibleSettingOptionsRequest{
		ComponentId: componentId,
	}
	response, err := s.Client.GetPossibleSettingOptions(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetSetting Set a setting to some value.

	Only setting_id of setting and option_id of option needs to be set.
*/
func (s *ServiceImpl) SetSetting(
	ctx context.Context,
	componentId int32,
	setting *Setting,

) (*SetSettingResponse, error) {
	request := &SetSettingRequest{
		ComponentId: componentId,
		Setting:     setting,
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
	componentId int32,
	setting *Setting,

) (*GetSettingResponse, error) {
	request := &GetSettingRequest{
		ComponentId: componentId,
		Setting:     setting,
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
	componentId int32,
	storageId int32,

) (*FormatStorageResponse, error) {
	request := &FormatStorageRequest{
		ComponentId: componentId,
		StorageId:   storageId,
	}
	response, err := s.Client.FormatStorage(ctx, request)
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
	componentId int32,

) (*ResetSettingsResponse, error) {
	request := &ResetSettingsRequest{
		ComponentId: componentId,
	}
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
	componentId int32,

) (*ZoomInStartResponse, error) {
	request := &ZoomInStartRequest{
		ComponentId: componentId,
	}
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
	componentId int32,

) (*ZoomOutStartResponse, error) {
	request := &ZoomOutStartRequest{
		ComponentId: componentId,
	}
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
	componentId int32,

) (*ZoomStopResponse, error) {
	request := &ZoomStopRequest{
		ComponentId: componentId,
	}
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
	componentId int32,

	rangeVar float32,

) (*ZoomRangeResponse, error) {
	request := &ZoomRangeRequest{
		ComponentId: componentId,

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
	componentId int32,
	pointX float32,
	pointY float32,
	radius float32,

) (*TrackPointResponse, error) {
	request := &TrackPointRequest{
		ComponentId: componentId,
		PointX:      pointX,
		PointY:      pointY,
		Radius:      radius,
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
	componentId int32,
	topLeftX float32,
	topLeftY float32,
	bottomRightX float32,
	bottomRightY float32,

) (*TrackRectangleResponse, error) {
	request := &TrackRectangleRequest{
		ComponentId:  componentId,
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
	componentId int32,

) (*TrackStopResponse, error) {
	request := &TrackStopRequest{
		ComponentId: componentId,
	}
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
	componentId int32,

) (*FocusInStartResponse, error) {
	request := &FocusInStartRequest{
		ComponentId: componentId,
	}
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
	componentId int32,

) (*FocusOutStartResponse, error) {
	request := &FocusOutStartRequest{
		ComponentId: componentId,
	}
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
	componentId int32,

) (*FocusStopResponse, error) {
	request := &FocusStopRequest{
		ComponentId: componentId,
	}
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
	componentId int32,

	rangeVar float32,

) (*FocusRangeResponse, error) {
	request := &FocusRangeRequest{
		ComponentId: componentId,

		Range: rangeVar,
	}
	response, err := s.Client.FocusRange(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

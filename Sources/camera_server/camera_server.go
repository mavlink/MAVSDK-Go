package camera_server

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client CameraServerServiceClient
}

/*
SetInformation Sets the camera information. This must be called as soon as the camera server is created.
*/
func (s *ServiceImpl) SetInformation(
	ctx context.Context,
	information *Information,

) (*SetInformationResponse, error) {
	request := &SetInformationRequest{
		Information: information,
	}
	response, err := s.Client.SetInformation(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetVideoStreaming Sets video streaming settings.
*/
func (s *ServiceImpl) SetVideoStreaming(
	ctx context.Context,
	videoStreaming *VideoStreaming,

) (*SetVideoStreamingResponse, error) {
	request := &SetVideoStreamingRequest{
		VideoStreaming: videoStreaming,
	}
	response, err := s.Client.SetVideoStreaming(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetInProgress Sets image capture in progress status flags. This should be set to true when the camera is busy taking a photo and false when it is done.
*/
func (s *ServiceImpl) SetInProgress(
	ctx context.Context,
	inProgress bool,

) (*SetInProgressResponse, error) {
	request := &SetInProgressRequest{
		InProgress: inProgress,
	}
	response, err := s.Client.SetInProgress(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
TakePhoto Subscribe to image capture requests. Each request received should respond to using RespondTakePhoto.
*/
func (a *ServiceImpl) TakePhoto(
	ctx context.Context,

) (<-chan int32, error) {
	ch := make(chan int32)
	request := &SubscribeTakePhotoRequest{}
	stream, err := a.Client.SubscribeTakePhoto(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &TakePhotoResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive TakePhoto messages, err: %v", err)
			}
			ch <- m.GetIndex()
		}
	}()
	return ch, nil
}

/*
RespondTakePhoto Respond to an image capture request from SubscribeTakePhoto.
*/
func (s *ServiceImpl) RespondTakePhoto(
	ctx context.Context,
	takePhotoFeedback *CameraFeedback,

	captureInfo *CaptureInfo,

) (*RespondTakePhotoResponse, error) {
	request := &RespondTakePhotoRequest{
		TakePhotoFeedback: *takePhotoFeedback,
		CaptureInfo:       captureInfo,
	}
	response, err := s.Client.RespondTakePhoto(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
StartVideo Subscribe to start video requests. Each request received should respond to using RespondStartVideo
*/
func (a *ServiceImpl) StartVideo(
	ctx context.Context,

) (<-chan int32, error) {
	ch := make(chan int32)
	request := &SubscribeStartVideoRequest{}
	stream, err := a.Client.SubscribeStartVideo(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &StartVideoResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive StartVideo messages, err: %v", err)
			}
			ch <- m.GetStreamId()
		}
	}()
	return ch, nil
}

/*
RespondStartVideo Subscribe to stop video requests. Each request received should respond using StopVideoResponse
*/
func (s *ServiceImpl) RespondStartVideo(
	ctx context.Context,
	startVideoFeedback *CameraFeedback,

) (*RespondStartVideoResponse, error) {
	request := &RespondStartVideoRequest{
		StartVideoFeedback: *startVideoFeedback,
	}
	response, err := s.Client.RespondStartVideo(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
StopVideo Subscribe to stop video requests. Each request received should response to using RespondStopVideo
*/
func (a *ServiceImpl) StopVideo(
	ctx context.Context,

) (<-chan int32, error) {
	ch := make(chan int32)
	request := &SubscribeStopVideoRequest{}
	stream, err := a.Client.SubscribeStopVideo(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &StopVideoResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive StopVideo messages, err: %v", err)
			}
			ch <- m.GetStreamId()
		}
	}()
	return ch, nil
}

/*
RespondStopVideo Respond to stop video request from SubscribeStopVideo.
*/
func (s *ServiceImpl) RespondStopVideo(
	ctx context.Context,
	stopVideoFeedback *CameraFeedback,

) (*RespondStopVideoResponse, error) {
	request := &RespondStopVideoRequest{
		StopVideoFeedback: *stopVideoFeedback,
	}
	response, err := s.Client.RespondStopVideo(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
StartVideoStreaming Subscribe to start video streaming requests. Each request received should response to using RespondStartVideoStreaming
*/
func (a *ServiceImpl) StartVideoStreaming(
	ctx context.Context,

) (<-chan int32, error) {
	ch := make(chan int32)
	request := &SubscribeStartVideoStreamingRequest{}
	stream, err := a.Client.SubscribeStartVideoStreaming(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &StartVideoStreamingResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive StartVideoStreaming messages, err: %v", err)
			}
			ch <- m.GetStreamId()
		}
	}()
	return ch, nil
}

/*
RespondStartVideoStreaming Respond to start video streaming from SubscribeStartVideoStreaming.
*/
func (s *ServiceImpl) RespondStartVideoStreaming(
	ctx context.Context,
	startVideoStreamingFeedback *CameraFeedback,

) (*RespondStartVideoStreamingResponse, error) {
	request := &RespondStartVideoStreamingRequest{
		StartVideoStreamingFeedback: *startVideoStreamingFeedback,
	}
	response, err := s.Client.RespondStartVideoStreaming(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
StopVideoStreaming Subscribe to stop video streaming requests. Each request received should response to using RespondStopVideoStreaming
*/
func (a *ServiceImpl) StopVideoStreaming(
	ctx context.Context,

) (<-chan int32, error) {
	ch := make(chan int32)
	request := &SubscribeStopVideoStreamingRequest{}
	stream, err := a.Client.SubscribeStopVideoStreaming(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &StopVideoStreamingResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive StopVideoStreaming messages, err: %v", err)
			}
			ch <- m.GetStreamId()
		}
	}()
	return ch, nil
}

/*
RespondStopVideoStreaming Respond to stop video streaming from SubscribeStopVideoStreaming.
*/
func (s *ServiceImpl) RespondStopVideoStreaming(
	ctx context.Context,
	stopVideoStreamingFeedback *CameraFeedback,

) (*RespondStopVideoStreamingResponse, error) {
	request := &RespondStopVideoStreamingRequest{
		StopVideoStreamingFeedback: *stopVideoStreamingFeedback,
	}
	response, err := s.Client.RespondStopVideoStreaming(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetMode Subscribe to set camera mode requests. Each request received should response to using RespondSetMode
*/
func (a *ServiceImpl) SetMode(
	ctx context.Context,

) (<-chan Mode, error) {
	ch := make(chan Mode)
	request := &SubscribeSetModeRequest{}
	stream, err := a.Client.SubscribeSetMode(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &SetModeResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive SetMode messages, err: %v", err)
			}
			ch <- m.GetMode()
		}
	}()
	return ch, nil
}

/*
RespondSetMode Respond to set camera mode from SubscribeSetMode.
*/
func (s *ServiceImpl) RespondSetMode(
	ctx context.Context,
	setModeFeedback *CameraFeedback,

) (*RespondSetModeResponse, error) {
	request := &RespondSetModeRequest{
		SetModeFeedback: *setModeFeedback,
	}
	response, err := s.Client.RespondSetMode(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
StorageInformation Subscribe to camera storage information requests. Each request received should response to using RespondStorageInformation
*/
func (a *ServiceImpl) StorageInformation(
	ctx context.Context,

) (<-chan int32, error) {
	ch := make(chan int32)
	request := &SubscribeStorageInformationRequest{}
	stream, err := a.Client.SubscribeStorageInformation(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &StorageInformationResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive StorageInformation messages, err: %v", err)
			}
			ch <- m.GetStorageId()
		}
	}()
	return ch, nil
}

/*
RespondStorageInformation Respond to camera storage information from SubscribeStorageInformation.
*/
func (s *ServiceImpl) RespondStorageInformation(
	ctx context.Context,
	storageInformationFeedback *CameraFeedback,

	storageInformation *StorageInformation,

) (*RespondStorageInformationResponse, error) {
	request := &RespondStorageInformationRequest{
		StorageInformationFeedback: *storageInformationFeedback,
		StorageInformation:         storageInformation,
	}
	response, err := s.Client.RespondStorageInformation(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
CaptureStatus Subscribe to camera capture status requests. Each request received should response to using RespondCaptureStatus
*/
func (a *ServiceImpl) CaptureStatus(
	ctx context.Context,

) (<-chan int32, error) {
	ch := make(chan int32)
	request := &SubscribeCaptureStatusRequest{}
	stream, err := a.Client.SubscribeCaptureStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &CaptureStatusResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive CaptureStatus messages, err: %v", err)
			}
			ch <- m.GetReserved()
		}
	}()
	return ch, nil
}

/*
RespondCaptureStatus Respond to camera capture status from SubscribeCaptureStatus.
*/
func (s *ServiceImpl) RespondCaptureStatus(
	ctx context.Context,
	captureStatusFeedback *CameraFeedback,

	captureStatus *CaptureStatus,

) (*RespondCaptureStatusResponse, error) {
	request := &RespondCaptureStatusRequest{
		CaptureStatusFeedback: *captureStatusFeedback,
		CaptureStatus:         captureStatus,
	}
	response, err := s.Client.RespondCaptureStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
FormatStorage Subscribe to format storage requests. Each request received should response to using RespondFormatStorage
*/
func (a *ServiceImpl) FormatStorage(
	ctx context.Context,

) (<-chan int32, error) {
	ch := make(chan int32)
	request := &SubscribeFormatStorageRequest{}
	stream, err := a.Client.SubscribeFormatStorage(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &FormatStorageResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive FormatStorage messages, err: %v", err)
			}
			ch <- m.GetStorageId()
		}
	}()
	return ch, nil
}

/*
RespondFormatStorage Respond to format storage from SubscribeFormatStorage.
*/
func (s *ServiceImpl) RespondFormatStorage(
	ctx context.Context,
	formatStorageFeedback *CameraFeedback,

) (*RespondFormatStorageResponse, error) {
	request := &RespondFormatStorageRequest{
		FormatStorageFeedback: *formatStorageFeedback,
	}
	response, err := s.Client.RespondFormatStorage(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ResetSettings Subscribe to reset settings requests. Each request received should response to using RespondResetSettings
*/
func (a *ServiceImpl) ResetSettings(
	ctx context.Context,

) (<-chan int32, error) {
	ch := make(chan int32)
	request := &SubscribeResetSettingsRequest{}
	stream, err := a.Client.SubscribeResetSettings(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ResetSettingsResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive ResetSettings messages, err: %v", err)
			}
			ch <- m.GetReserved()
		}
	}()
	return ch, nil
}

/*
RespondResetSettings Respond to reset settings from SubscribeResetSettings.
*/
func (s *ServiceImpl) RespondResetSettings(
	ctx context.Context,
	resetSettingsFeedback *CameraFeedback,

) (*RespondResetSettingsResponse, error) {
	request := &RespondResetSettingsRequest{
		ResetSettingsFeedback: *resetSettingsFeedback,
	}
	response, err := s.Client.RespondResetSettings(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ZoomInStart Subscribe to zoom in start command
*/
func (a *ServiceImpl) ZoomInStart(
	ctx context.Context,

) (<-chan int32, error) {
	ch := make(chan int32)
	request := &SubscribeZoomInStartRequest{}
	stream, err := a.Client.SubscribeZoomInStart(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ZoomInStartResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive ZoomInStart messages, err: %v", err)
			}
			ch <- m.GetReserved()
		}
	}()
	return ch, nil
}

/*
RespondZoomInStart Respond to zoom in start.
*/
func (s *ServiceImpl) RespondZoomInStart(
	ctx context.Context,
	zoomInStartFeedback *CameraFeedback,

) (*RespondZoomInStartResponse, error) {
	request := &RespondZoomInStartRequest{
		ZoomInStartFeedback: *zoomInStartFeedback,
	}
	response, err := s.Client.RespondZoomInStart(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ZoomOutStart Subscribe to zoom out start command
*/
func (a *ServiceImpl) ZoomOutStart(
	ctx context.Context,

) (<-chan int32, error) {
	ch := make(chan int32)
	request := &SubscribeZoomOutStartRequest{}
	stream, err := a.Client.SubscribeZoomOutStart(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ZoomOutStartResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive ZoomOutStart messages, err: %v", err)
			}
			ch <- m.GetReserved()
		}
	}()
	return ch, nil
}

/*
RespondZoomOutStart Respond to zoom out start.
*/
func (s *ServiceImpl) RespondZoomOutStart(
	ctx context.Context,
	zoomOutStartFeedback *CameraFeedback,

) (*RespondZoomOutStartResponse, error) {
	request := &RespondZoomOutStartRequest{
		ZoomOutStartFeedback: *zoomOutStartFeedback,
	}
	response, err := s.Client.RespondZoomOutStart(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ZoomStop Subscribe to zoom stop command
*/
func (a *ServiceImpl) ZoomStop(
	ctx context.Context,

) (<-chan int32, error) {
	ch := make(chan int32)
	request := &SubscribeZoomStopRequest{}
	stream, err := a.Client.SubscribeZoomStop(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ZoomStopResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive ZoomStop messages, err: %v", err)
			}
			ch <- m.GetReserved()
		}
	}()
	return ch, nil
}

/*
RespondZoomStop Respond to zoom stop.
*/
func (s *ServiceImpl) RespondZoomStop(
	ctx context.Context,
	zoomStopFeedback *CameraFeedback,

) (*RespondZoomStopResponse, error) {
	request := &RespondZoomStopRequest{
		ZoomStopFeedback: *zoomStopFeedback,
	}
	response, err := s.Client.RespondZoomStop(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
ZoomRange Subscribe to zoom range command
*/
func (a *ServiceImpl) ZoomRange(
	ctx context.Context,

) (<-chan float32, error) {
	ch := make(chan float32)
	request := &SubscribeZoomRangeRequest{}
	stream, err := a.Client.SubscribeZoomRange(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ZoomRangeResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive ZoomRange messages, err: %v", err)
			}
			ch <- m.GetFactor()
		}
	}()
	return ch, nil
}

/*
RespondZoomRange Respond to zoom range.
*/
func (s *ServiceImpl) RespondZoomRange(
	ctx context.Context,
	zoomRangeFeedback *CameraFeedback,

) (*RespondZoomRangeResponse, error) {
	request := &RespondZoomRangeRequest{
		ZoomRangeFeedback: *zoomRangeFeedback,
	}
	response, err := s.Client.RespondZoomRange(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetTrackingRectangleStatus Set/update the current rectangle tracking status.
*/
func (s *ServiceImpl) SetTrackingRectangleStatus(
	ctx context.Context,
	trackedRectangle *TrackRectangle,

) (*SetTrackingRectangleStatusResponse, error) {
	request := &SetTrackingRectangleStatusRequest{
		TrackedRectangle: trackedRectangle,
	}
	response, err := s.Client.SetTrackingRectangleStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
SetTrackingOffStatus Set the current tracking status to off.
*/
func (s *ServiceImpl) SetTrackingOffStatus(
	ctx context.Context,

) (*SetTrackingOffStatusResponse, error) {
	request := &SetTrackingOffStatusRequest{}
	response, err := s.Client.SetTrackingOffStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
TrackingPointCommand Subscribe to incoming tracking point command.
*/
func (a *ServiceImpl) TrackingPointCommand(
	ctx context.Context,

) (<-chan *TrackPoint, error) {
	ch := make(chan *TrackPoint)
	request := &SubscribeTrackingPointCommandRequest{}
	stream, err := a.Client.SubscribeTrackingPointCommand(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &TrackingPointCommandResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive TrackingPointCommand messages, err: %v", err)
			}
			ch <- m.GetTrackPoint()
		}
	}()
	return ch, nil
}

/*
TrackingRectangleCommand Subscribe to incoming tracking rectangle command.
*/
func (a *ServiceImpl) TrackingRectangleCommand(
	ctx context.Context,

) (<-chan *TrackRectangle, error) {
	ch := make(chan *TrackRectangle)
	request := &SubscribeTrackingRectangleCommandRequest{}
	stream, err := a.Client.SubscribeTrackingRectangleCommand(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &TrackingRectangleCommandResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive TrackingRectangleCommand messages, err: %v", err)
			}
			ch <- m.GetTrackRectangle()
		}
	}()
	return ch, nil
}

/*
TrackingOffCommand Subscribe to incoming tracking off command.
*/
func (a *ServiceImpl) TrackingOffCommand(
	ctx context.Context,

) (<-chan int32, error) {
	ch := make(chan int32)
	request := &SubscribeTrackingOffCommandRequest{}
	stream, err := a.Client.SubscribeTrackingOffCommand(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &TrackingOffCommandResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive TrackingOffCommand messages, err: %v", err)
			}
			ch <- m.GetDummy()
		}
	}()
	return ch, nil
}

/*
RespondTrackingPointCommand Respond to an incoming tracking point command.
*/
func (s *ServiceImpl) RespondTrackingPointCommand(
	ctx context.Context,
	stopVideoFeedback *CameraFeedback,

) (*RespondTrackingPointCommandResponse, error) {
	request := &RespondTrackingPointCommandRequest{
		StopVideoFeedback: *stopVideoFeedback,
	}
	response, err := s.Client.RespondTrackingPointCommand(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
RespondTrackingRectangleCommand Respond to an incoming tracking rectangle command.
*/
func (s *ServiceImpl) RespondTrackingRectangleCommand(
	ctx context.Context,
	stopVideoFeedback *CameraFeedback,

) (*RespondTrackingRectangleCommandResponse, error) {
	request := &RespondTrackingRectangleCommandRequest{
		StopVideoFeedback: *stopVideoFeedback,
	}
	response, err := s.Client.RespondTrackingRectangleCommand(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
RespondTrackingOffCommand Respond to an incoming tracking off command.
*/
func (s *ServiceImpl) RespondTrackingOffCommand(
	ctx context.Context,
	stopVideoFeedback *CameraFeedback,

) (*RespondTrackingOffCommandResponse, error) {
	request := &RespondTrackingOffCommandRequest{
		StopVideoFeedback: *stopVideoFeedback,
	}
	response, err := s.Client.RespondTrackingOffCommand(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

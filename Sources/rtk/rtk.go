package rtk

import (
	"context"
)

type ServiceImpl struct {
	Client RtkServiceClient
}

/*
SendRtcmData Send RTCM data.
*/
func (s *ServiceImpl) SendRtcmData(
	ctx context.Context,
	rtcmData *RtcmData,

) (*SendRtcmDataResponse, error) {
	request := &SendRtcmDataRequest{
		RtcmData: rtcmData,
	}
	response, err := s.Client.SendRtcmData(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

package server_utility

import (
	"context"
)

type ServiceImpl struct {
	Client ServerUtilityServiceClient
}

/*
SendStatusText Sends a statustext.
*/
func (s *ServiceImpl) SendStatusText(
	ctx context.Context,

	typeVar StatusTextType,
	text string,

) (*SendStatusTextResponse, error) {
	request := &SendStatusTextRequest{

		Type: typeVar,
		Text: text,
	}
	response, err := s.Client.SendStatusText(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

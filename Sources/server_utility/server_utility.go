package server_utility

import (
	"context"
)

type ServiceImpl struct {
	Client ServerUtilityServiceClient
}

/*
   Sends a statustext.

   Parameters
   ----------
   type *StatusTextType


   text string


*/

func (s *ServiceImpl) SendStatusText(ctx context.Context, statusTextType *StatusTextType, text string) (*SendStatusTextResponse, error) {

	request := &SendStatusTextRequest{}
	request.Type = *statusTextType
	request.Text = text
	response, err := s.Client.SendStatusText(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

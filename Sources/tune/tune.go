package tune

import (
	"context"
)

type ServiceImpl struct {
	Client TuneServiceClient
}

/*
   Send a tune to be played by the system.

   Parameters
   ----------
   tuneDescription *TuneDescription



*/

func (s *ServiceImpl) PlayTune(ctx context.Context, tuneDescription *TuneDescription) (*PlayTuneResponse, error) {

	request := &PlayTuneRequest{}
	request.TuneDescription = tuneDescription

	response, err := s.Client.PlayTune(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

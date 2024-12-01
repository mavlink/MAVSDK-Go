package failure

import (
	"context"
)

type ServiceImpl struct {
	Client FailureServiceClient
}

/*
Inject Injects a failure.
*/
func (s *ServiceImpl) Inject(
	ctx context.Context,
	failureUnit *FailureUnit,

	failureType *FailureType,

	instance int32,

) (*InjectResponse, error) {
	request := &InjectRequest{
		FailureUnit: *failureUnit,
		FailureType: *failureType,
		Instance:    instance,
	}
	response, err := s.Client.Inject(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

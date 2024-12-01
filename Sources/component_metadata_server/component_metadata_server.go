package component_metadata_server

import (
	"context"
)

type ServiceImpl struct {
	Client ComponentMetadataServerServiceClient
}

/*
SetMetadata Provide metadata (can only be called once)
*/
func (s *ServiceImpl) SetMetadata(
	ctx context.Context,
	metadata []*Metadata,

) (*SetMetadataResponse, error) {
	request := &SetMetadataRequest{
		Metadata: metadata,
	}
	response, err := s.Client.SetMetadata(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

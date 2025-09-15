package component_metadata

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client ComponentMetadataServiceClient
}

/*
RequestComponent Request metadata from a specific component. This is used to start requesting metadata from a component.

	The metadata can later be accessed via subscription (see below) or GetMetadata.
*/
func (s *ServiceImpl) RequestComponent(
	ctx context.Context,
	compid uint32,

) (*RequestComponentResponse, error) {
	request := &RequestComponentRequest{
		Compid: compid,
	}
	response, err := s.Client.RequestComponent(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
RequestAutopilotComponent Request metadata from the autopilot component. This is used to start requesting metadata from the autopilot.

	The metadata can later be accessed via subscription (see below) or GetMetadata.
*/
func (s *ServiceImpl) RequestAutopilotComponent(
	ctx context.Context,

) (*RequestAutopilotComponentResponse, error) {
	request := &RequestAutopilotComponentRequest{}
	response, err := s.Client.RequestAutopilotComponent(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
MetadataAvailable Register a callback that gets called when metadata is available
*/
func (a *ServiceImpl) MetadataAvailable(
	ctx context.Context,

) (<-chan *MetadataUpdate, error) {
	ch := make(chan *MetadataUpdate)
	request := &SubscribeMetadataAvailableRequest{}
	stream, err := a.Client.SubscribeMetadataAvailable(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &MetadataAvailableResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && (s.Code() == codes.Canceled || s.Code() == codes.Unimplemented) {
					return
				}
				log.Fatalf("Unable to receive MetadataAvailable messages, err: %v", err)
			}
			ch <- m.GetData()
		}
	}()
	return ch, nil
}

/*
GetMetadata Access metadata. This can be used if you know the metadata is available already, otherwise use

	the subscription to get notified when it becomes available.
*/
func (s *ServiceImpl) GetMetadata(
	ctx context.Context,
	compid uint32,
	metadataType *MetadataType,

) (*GetMetadataResponse, error) {
	request := &GetMetadataRequest{
		Compid:       compid,
		MetadataType: *metadataType,
	}
	response, err := s.Client.GetMetadata(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

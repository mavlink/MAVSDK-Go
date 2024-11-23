package mission_raw_server

import (
	"context"
	"io"
	"log"

	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceImpl struct {
	Client MissionRawServerServiceClient
}

/*
   Subscribe to when a new mission is uploaded (asynchronous).


*/

func (a *ServiceImpl) IncomingMission(ctx context.Context) (<-chan *MissionPlan, error) {
	ch := make(chan *MissionPlan)
	request := &SubscribeIncomingMissionRequest{}
	stream, err := a.Client.SubscribeIncomingMission(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &IncomingMissionResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive IncomingMission messages, err: %v", err)
				break
			}
			ch <- m.GetMissionPlan()
		}
	}()
	return ch, nil
}

/*
   Subscribe to when a new current item is set


*/

func (a *ServiceImpl) CurrentItemChanged(ctx context.Context) (<-chan *MissionItem, error) {
	ch := make(chan *MissionItem)
	request := &SubscribeCurrentItemChangedRequest{}
	stream, err := a.Client.SubscribeCurrentItemChanged(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &CurrentItemChangedResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive CurrentItemChanged messages, err: %v", err)
				break
			}
			ch <- m.GetMissionItem()
		}
	}()
	return ch, nil
}

/*
   Set Current item as completed


*/

func (s *ServiceImpl) SetCurrentItemComplete(ctx context.Context) (*SetCurrentItemCompleteResponse, error) {

	request := &SetCurrentItemCompleteRequest{}
	response, err := s.Client.SetCurrentItemComplete(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Subscribe when a MISSION_CLEAR_ALL is received


*/

func (a *ServiceImpl) ClearAll(ctx context.Context) (<-chan uint32, error) {
	ch := make(chan uint32)
	request := &SubscribeClearAllRequest{}
	stream, err := a.Client.SubscribeClearAll(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ClearAllResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				return
			}
			if err != nil {
				if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
					return
				}
				log.Fatalf("Unable to receive ClearAll messages, err: %v", err)
				break
			}
			ch <- m.GetClearType()
		}
	}()
	return ch, nil
}

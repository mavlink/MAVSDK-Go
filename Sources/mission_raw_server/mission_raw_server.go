package mission_raw_server

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct {
	Client MissionRawServerServiceClient
}

/*
   Subscribe to when a new mission is uploaded (asynchronous).


*/

func (a *ServiceImpl) IncomingMission() (<-chan *MissionPlan, error) {
	ch := make(chan *MissionPlan)
	request := &SubscribeIncomingMissionRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
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

func (a *ServiceImpl) CurrentItemChanged() (<-chan *MissionItem, error) {
	ch := make(chan *MissionItem)
	request := &SubscribeCurrentItemChangedRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
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

func (s *ServiceImpl) SetCurrentItemComplete() (*SetCurrentItemCompleteResponse, error) {

	request := &SetCurrentItemCompleteRequest{}
	ctx := context.Background()
	response, err := s.Client.SetCurrentItemComplete(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/*
   Subscribe when a MISSION_CLEAR_ALL is received


*/

func (a *ServiceImpl) ClearAll() (<-chan uint32, error) {
	ch := make(chan uint32)
	request := &SubscribeClearAllRequest{}
	ctx := context.Background()
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
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetClearType()
		}
	}()
	return ch, nil
}

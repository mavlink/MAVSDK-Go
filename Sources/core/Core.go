package core

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct {
	Client CoreServiceClient
}

/*
   Subscribe to 'connection state' updates.


*/

func (a *ServiceImpl) ConnectionState() (<-chan *ConnectionState, error) {
	ch := make(chan *ConnectionState)
	request := &SubscribeConnectionStateRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeConnectionState(ctx, request)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for {
			m := &ConnectionStateResponse{}
			err := stream.RecvMsg(m)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Unable to receive message %v", err)
				break
			}
			ch <- m.GetConnectionState()
		}
	}()
	return ch, nil
}

/*
   Get a list of currently running plugins.



   Returns
   -------
   True
   PluginInfo : []*PluginInfo
        Plugin info


*/

func (s *ServiceImpl) ListRunningPlugins() (*ListRunningPluginsResponse, error) {
	request := &ListRunningPluginsRequest{}
	ctx := context.Background()
	response, err := s.Client.ListRunningPlugins(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil

}

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

func (a *ServiceImpl) ConnectionState() <-chan *ConnectionState {
	ch := make(chan *ConnectionState)
	request := &SubscribeConnectionStateRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeConnectionState(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
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
	return ch
}

/*
   Get a list of currently running plugins.



   Returns
   -------
   True
   PluginInfo : []*PluginInfo
        Plugin info


*/

func (s *ServiceImpl) ListRunningPlugins() []*PluginInfo {
	request := &ListRunningPluginsRequest{}
	ctx := context.Background()
	response, err := s.Client.ListRunningPlugins(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
	}

	return response.GetPluginInfo()

}

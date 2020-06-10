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
func (a *ServiceImpl) ConnectionState() {
	request := &SubscribeConnectionStateRequest{}
	ctx := context.Background()
	stream, err := a.Client.SubscribeConnectionState(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

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
		fmt.Printf("message %v\n", m)
	}
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

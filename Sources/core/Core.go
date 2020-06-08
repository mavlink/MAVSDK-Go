package core

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client CoreServiceClient
}

    func (a *ServiceImpl) ConnectionState(){
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
    			fmt.Printf("message %v\n", m.GetConnectionState())
    		}	
    }

    func(s *ServiceImpl)ListRunningPlugins() (*ListRunningPluginsResponse){
        request := &ListRunningPluginsRequest{}
        ctx:= context.Background()
         response, err := s.Client.ListRunningPlugins(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        

        
    }

       
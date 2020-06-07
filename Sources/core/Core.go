import (
	"context"
	"fmt"
	"io"
)

type Service interface{
    Result() CoreResult_Result

}

type ServiceImpl struct{
    Client CoreServiceClient
}

    func (a *ServiceImpl) ConnectionState(){
    	request := &ConnectionStateRequest{}
    		ctx := context.Background()
    		stream, err := a.Core.SubscribeConnectionState(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &connection_stateResponse{}
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

    func(s *ServiceImpl)list_running_plugins(self)([]*ListRunningPlugins){
     request = &ListRunningPluginsRequest{}
         response, err = s.Client.ListRunningPlugins(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        

        
    }

       
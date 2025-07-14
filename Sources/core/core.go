package core

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client CoreServiceClient
}

     /*
         Subscribe to 'connection state' updates.

         
    */

    func (a *ServiceImpl) ConnectionState(ctx context.Context, ) (<-chan  *ConnectionState , error){
    		ch := make(chan  *ConnectionState )
    		request := &SubscribeConnectionStateRequest{}
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
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive ConnectionState messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetConnectionState()
    			}
    		}()	
    	return ch, nil
    }
    /*
         Set timeout of MAVLink transfers.

         The default timeout used is generally (0.5 seconds) seconds.
         If MAVSDK is used on the same host this timeout can be reduced, while
         if MAVSDK has to communicate over links with high latency it might
         need to be increased to prevent timeouts.

         Parameters
         ----------
         timeoutS float64

         
    */

    func(s *ServiceImpl)SetMavlinkTimeout(ctx context.Context, timeoutS float64)(*SetMavlinkTimeoutResponse, error){
        
        request := &SetMavlinkTimeoutRequest{}
    	request.TimeoutS = timeoutS
        response, err := s.Client.SetMavlinkTimeout(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
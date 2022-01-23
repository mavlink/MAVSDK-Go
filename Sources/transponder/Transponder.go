package transponder

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client TransponderServiceClient
}

     /*
         Subscribe to 'transponder' updates.

         
    */

    func (a *ServiceImpl) Transponder() (<-chan  *AdsbVehicle , error){
    		ch := make(chan  *AdsbVehicle )
    		request := &SubscribeTransponderRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeTransponder(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &TransponderResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetTransponder()
    		}
    	}()	
    	return ch, nil
    }
    /*
         Set rate to 'transponder' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateTransponder(rateHz float64)(*SetRateTransponderResponse, error){
        
        request := &SetRateTransponderRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateTransponder(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
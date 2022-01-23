package tracking_server

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client TrackingServerServiceClient
}
    /*
         Set/update the current point tracking status.

         Parameters
         ----------
         trackedPoint *TrackPoint 
            

         
    */

    func(s *ServiceImpl)SetTrackingPointStatus(trackedPoint *TrackPoint )(*SetTrackingPointStatusResponse, error){
        
        request := &SetTrackingPointStatusRequest{}
        ctx:= context.Background()
         request.TrackedPoint = trackedPoint
            
        response, err := s.Client.SetTrackingPointStatus(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set/update the current rectangle tracking status.

         Parameters
         ----------
         trackedRectangle *TrackRectangle 
            

         
    */

    func(s *ServiceImpl)SetTrackingRectangleStatus(trackedRectangle *TrackRectangle )(*SetTrackingRectangleStatusResponse, error){
        
        request := &SetTrackingRectangleStatusRequest{}
        ctx:= context.Background()
         request.TrackedRectangle = trackedRectangle
            
        response, err := s.Client.SetTrackingRectangleStatus(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set the current tracking status to off.

         
    */

    func(s *ServiceImpl)SetTrackingOffStatus()(*SetTrackingOffStatusResponse, error){
        
        request := &SetTrackingOffStatusRequest{}
        ctx:= context.Background()
         response, err := s.Client.SetTrackingOffStatus(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       

     /*
         Subscribe to incoming tracking point command.

         
    */

    func (a *ServiceImpl) TrackingPointCommand() (<-chan  *TrackPoint , error){
    		ch := make(chan  *TrackPoint )
    		request := &SubscribeTrackingPointCommandRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeTrackingPointCommand(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &TrackingPointCommandResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetTrackPoint()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to incoming tracking rectangle command.

         
    */

    func (a *ServiceImpl) TrackingRectangleCommand() (<-chan  *TrackRectangle , error){
    		ch := make(chan  *TrackRectangle )
    		request := &SubscribeTrackingRectangleCommandRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeTrackingRectangleCommand(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &TrackingRectangleCommandResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetTrackRectangle()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to incoming tracking off command.

         
    */

    func (a *ServiceImpl) TrackingOffCommand() (<-chan   int32  , error){
    		ch := make(chan   int32  )
    		request := &SubscribeTrackingOffCommandRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeTrackingOffCommand(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &TrackingOffCommandResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetDummy()
    		}
    	}()	
    	return ch, nil
    }
    /*
         Respond to an incoming tracking point command.

         Parameters
         ----------
         commandAnswer *CommandAnswer 
            

         
    */

    func(s *ServiceImpl)RespondTrackingPointCommand(commandAnswer *CommandAnswer )(*RespondTrackingPointCommandResponse, error){
        
        request := &RespondTrackingPointCommandRequest{}
        ctx:= context.Background()
         request.CommandAnswer = commandAnswer
            
        response, err := s.Client.RespondTrackingPointCommand(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Respond to an incoming tracking rectangle command.

         Parameters
         ----------
         commandAnswer *CommandAnswer 
            

         
    */

    func(s *ServiceImpl)RespondTrackingRectangleCommand(commandAnswer *CommandAnswer )(*RespondTrackingRectangleCommandResponse, error){
        
        request := &RespondTrackingRectangleCommandRequest{}
        ctx:= context.Background()
         request.CommandAnswer = commandAnswer
            
        response, err := s.Client.RespondTrackingRectangleCommand(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Respond to an incoming tracking off command.

         Parameters
         ----------
         commandAnswer *CommandAnswer 
            

         
    */

    func(s *ServiceImpl)RespondTrackingOffCommand(commandAnswer *CommandAnswer )(*RespondTrackingOffCommandResponse, error){
        
        request := &RespondTrackingOffCommandRequest{}
        ctx:= context.Background()
         request.CommandAnswer = commandAnswer
            
        response, err := s.Client.RespondTrackingOffCommand(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
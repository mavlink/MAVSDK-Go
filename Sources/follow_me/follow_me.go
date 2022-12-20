package follow_me

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client FollowMeServiceClient
}
    /*
         Get current configuration.

         

         Returns
         -------
         False
         Config : Config
              The current configuration

         
    */


    func(s *ServiceImpl)GetConfig(ctx context.Context, ) (*GetConfigResponse, error){
        request := &GetConfigRequest{}
    	response, err := s.Client.GetConfig(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Apply configuration by sending it to the system.

         Parameters
         ----------
         config *Config 
            

         
    */

    func(s *ServiceImpl)SetConfig(ctx context.Context, config *Config )(*SetConfigResponse, error){
        
        request := &SetConfigRequest{}
    	request.Config = config
            
        response, err := s.Client.SetConfig(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Check if FollowMe is active.

         

         Returns
         -------
         False
         IsActive : bool
              Whether follow me is active or not

         
    */


    func(s *ServiceImpl)IsActive(ctx context.Context, ) (*IsActiveResponse, error){
        request := &IsActiveRequest{}
    	response, err := s.Client.IsActive(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Set location of the moving target.

         Parameters
         ----------
         location *TargetLocation 
            

         
    */

    func(s *ServiceImpl)SetTargetLocation(ctx context.Context, location *TargetLocation )(*SetTargetLocationResponse, error){
        
        request := &SetTargetLocationRequest{}
    	request.Location = location
            
        response, err := s.Client.SetTargetLocation(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Get the last location of the target.

         

         Returns
         -------
         False
         Location : TargetLocation
              The last target location that was set

         
    */


    func(s *ServiceImpl)GetLastLocation(ctx context.Context, ) (*GetLastLocationResponse, error){
        request := &GetLastLocationRequest{}
    	response, err := s.Client.GetLastLocation(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
    /*
         Start FollowMe mode.

         
    */

    func(s *ServiceImpl)Start(ctx context.Context, )(*StartResponse, error){
        
        request := &StartRequest{}
    	response, err := s.Client.Start(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Stop FollowMe mode.

         
    */

    func(s *ServiceImpl)Stop(ctx context.Context, )(*StopResponse, error){
        
        request := &StopRequest{}
    	response, err := s.Client.Stop(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
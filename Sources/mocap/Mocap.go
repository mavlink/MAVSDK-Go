package mocap

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client MocapServiceClient
}
    /*
         Send Global position/attitude estimate from a vision source.

         Parameters
         ----------
         visionPositionEstimate *VisionPositionEstimate 
            

         
    */

    func(s *ServiceImpl)SetVisionPositionEstimate(visionPositionEstimate *VisionPositionEstimate )(*SetVisionPositionEstimateResponse, error){
        
        request := &SetVisionPositionEstimateRequest{}
        ctx:= context.Background()
         request.VisionPositionEstimate = visionPositionEstimate
            
        response, err := s.Client.SetVisionPositionEstimate(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Send motion capture attitude and position.

         Parameters
         ----------
         attitudePositionMocap *AttitudePositionMocap 
            

         
    */

    func(s *ServiceImpl)SetAttitudePositionMocap(attitudePositionMocap *AttitudePositionMocap )(*SetAttitudePositionMocapResponse, error){
        
        request := &SetAttitudePositionMocapRequest{}
        ctx:= context.Background()
         request.AttitudePositionMocap = attitudePositionMocap
            
        response, err := s.Client.SetAttitudePositionMocap(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Send odometry information with an external interface.

         Parameters
         ----------
         odometry *Odometry 
            

         
    */

    func(s *ServiceImpl)SetOdometry(odometry *Odometry )(*SetOdometryResponse, error){
        
        request := &SetOdometryRequest{}
        ctx:= context.Background()
         request.Odometry = odometry
            
        response, err := s.Client.SetOdometry(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
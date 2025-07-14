package gimbal

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client GimbalServiceClient
}
    /*
         Set gimbal roll, pitch and yaw angles.

         This sets the desired roll, pitch and yaw angles of a gimbal.
         Will return when the command is accepted, however, it might
         take the gimbal longer to actually be set to the new angles.

         Note that the roll angle needs to be set to 0 when send_mode is Once.

         Parameters
         ----------
         gimbalId int32

         rollDeg float32

         pitchDeg float32

         yawDeg float32

         gimbalMode *GimbalMode 
            

         sendMode *SendMode 
            

         
    */

    func(s *ServiceImpl)SetAngles(ctx context.Context, gimbalId int32, rollDeg float32, pitchDeg float32, yawDeg float32, gimbalMode *GimbalMode , sendMode *SendMode )(*SetAnglesResponse, error){
        
        request := &SetAnglesRequest{}
    	request.GimbalId = gimbalId
        request.RollDeg = rollDeg
        request.PitchDeg = pitchDeg
        request.YawDeg = yawDeg
        request.GimbalMode = *gimbalMode
        request.SendMode = *sendMode
        response, err := s.Client.SetAngles(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set gimbal angular rates.

         This sets the desired angular rates around roll, pitch and yaw axes of a gimbal.
         Will return when the command is accepted, however, it might
         take the gimbal longer to actually reach the angular rate.

         Note that the roll angle needs to be set to 0 when send_mode is Once.

         Parameters
         ----------
         gimbalId int32

         rollRateDegS float32

         pitchRateDegS float32

         yawRateDegS float32

         gimbalMode *GimbalMode 
            

         sendMode *SendMode 
            

         
    */

    func(s *ServiceImpl)SetAngularRates(ctx context.Context, gimbalId int32, rollRateDegS float32, pitchRateDegS float32, yawRateDegS float32, gimbalMode *GimbalMode , sendMode *SendMode )(*SetAngularRatesResponse, error){
        
        request := &SetAngularRatesRequest{}
    	request.GimbalId = gimbalId
        request.RollRateDegS = rollRateDegS
        request.PitchRateDegS = pitchRateDegS
        request.YawRateDegS = yawRateDegS
        request.GimbalMode = *gimbalMode
        request.SendMode = *sendMode
        response, err := s.Client.SetAngularRates(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set gimbal region of interest (ROI).

         This sets a region of interest that the gimbal will point to.
         The gimbal will continue to point to the specified region until it
         receives a new command.
         The function will return when the command is accepted, however, it might
         take the gimbal longer to actually rotate to the ROI.

         Parameters
         ----------
         gimbalId int32

         latitudeDeg float64

         longitudeDeg float64

         altitudeM float32

         
    */

    func(s *ServiceImpl)SetRoiLocation(ctx context.Context, gimbalId int32, latitudeDeg float64, longitudeDeg float64, altitudeM float32)(*SetRoiLocationResponse, error){
        
        request := &SetRoiLocationRequest{}
    	request.GimbalId = gimbalId
        request.LatitudeDeg = latitudeDeg
        request.LongitudeDeg = longitudeDeg
        request.AltitudeM = altitudeM
        response, err := s.Client.SetRoiLocation(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Take control.

         There can be only two components in control of a gimbal at any given time.
         One with "primary" control, and one with "secondary" control. The way the
         secondary control is implemented is not specified and hence depends on the
         vehicle.

         Components are expected to be cooperative, which means that they can
         override each other and should therefore do it carefully.

         Parameters
         ----------
         gimbalId int32

         controlMode *ControlMode 
            

         
    */

    func(s *ServiceImpl)TakeControl(ctx context.Context, gimbalId int32, controlMode *ControlMode )(*TakeControlResponse, error){
        
        request := &TakeControlRequest{}
    	request.GimbalId = gimbalId
        request.ControlMode = *controlMode
        response, err := s.Client.TakeControl(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Release control.

         Release control, such that other components can control the gimbal.

         Parameters
         ----------
         gimbalId int32

         
    */

    func(s *ServiceImpl)ReleaseControl(ctx context.Context, gimbalId int32)(*ReleaseControlResponse, error){
        
        request := &ReleaseControlRequest{}
    	request.GimbalId = gimbalId
        response, err := s.Client.ReleaseControl(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       

     /*
         Subscribe to list of gimbals.

         This allows to find out what gimbals are connected to the system.
         Based on the gimbal ID, we can then address a specific gimbal.

         
    */

    func (a *ServiceImpl) GimbalList(ctx context.Context, ) (<-chan  *GimbalList , error){
    		ch := make(chan  *GimbalList )
    		request := &SubscribeGimbalListRequest{}
    		stream, err := a.Client.SubscribeGimbalList(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &GimbalListResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive GimbalList messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetGimbalList()
    			}
    		}()	
    	return ch, nil
    }

     /*
         Subscribe to control status updates.

         This allows a component to know if it has primary, secondary or
         no control over the gimbal. Also, it gives the system and component ids
         of the other components in control (if any).

         
    */

    func (a *ServiceImpl) ControlStatus(ctx context.Context, ) (<-chan  *ControlStatus , error){
    		ch := make(chan  *ControlStatus )
    		request := &SubscribeControlStatusRequest{}
    		stream, err := a.Client.SubscribeControlStatus(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &ControlStatusResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive ControlStatus messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetControlStatus()
    			}
    		}()	
    	return ch, nil
    }
    /*
         Get control status for specific gimbal.

         Parameters
         ----------
         gimbalId int32

         Returns
         -------
         False
         ControlStatus : ControlStatus
              Control status

         
    */


    func(s *ServiceImpl)GetControlStatus(ctx context.Context, gimbalId int32) (*GetControlStatusResponse, error){
        request := &GetControlStatusRequest{}
    	request.GimbalId = gimbalId
        response, err := s.Client.GetControlStatus(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       

     /*
         Subscribe to attitude updates.

         This gets you the gimbal's attitude and angular rate.

         
    */

    func (a *ServiceImpl) Attitude(ctx context.Context, ) (<-chan  *Attitude , error){
    		ch := make(chan  *Attitude )
    		request := &SubscribeAttitudeRequest{}
    		stream, err := a.Client.SubscribeAttitude(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    		go func() {
    			defer close(ch)
    			for {
    				m := &AttitudeResponse{}
    				err := stream.RecvMsg(m)
    				if err == io.EOF {
    					return
    				}
    				if err != nil {
    					if s, ok := status.FromError(err); ok && s.Code() == codes.Canceled {
    						return
    					}
    					fmt.Printf("Unable to receive Attitude messages, err: %v\n", err)
    					break
    				}
    				ch <- m.GetAttitude()
    			}
    		}()	
    	return ch, nil
    }
    /*
         Get attitude for specific gimbal.

         Parameters
         ----------
         gimbalId int32

         Returns
         -------
         False
         Attitude : Attitude
              The attitude

         
    */


    func(s *ServiceImpl)GetAttitude(ctx context.Context, gimbalId int32) (*GetAttitudeResponse, error){
        request := &GetAttitudeRequest{}
    	request.GimbalId = gimbalId
        response, err := s.Client.GetAttitude(ctx, request)
        if err != nil {
    		return nil, err
    	}
        return response, nil

    }

       
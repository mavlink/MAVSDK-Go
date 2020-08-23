package telemetry

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client TelemetryServiceClient
}

     /*
         Subscribe to 'position' updates.

         
    */

    func (a *ServiceImpl) Position() (<-chan  *Position , error){
    		ch := make(chan  *Position )
    		request := &SubscribePositionRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribePosition(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &PositionResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetPosition()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'home position' updates.

         
    */

    func (a *ServiceImpl) Home() (<-chan  *Position , error){
    		ch := make(chan  *Position )
    		request := &SubscribeHomeRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeHome(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &HomeResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetHome()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to in-air updates.

         
    */

    func (a *ServiceImpl) InAir() (<-chan   bool  , error){
    		ch := make(chan   bool  )
    		request := &SubscribeInAirRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeInAir(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &InAirResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetIsInAir()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to landed state updates

         
    */

    func (a *ServiceImpl) LandedState() (<-chan   LandedState  , error){
    		ch := make(chan   LandedState  )
    		request := &SubscribeLandedStateRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeLandedState(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &LandedStateResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetLandedState()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to armed updates.

         
    */

    func (a *ServiceImpl) Armed() (<-chan   bool  , error){
    		ch := make(chan   bool  )
    		request := &SubscribeArmedRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeArmed(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &ArmedResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetIsArmed()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'attitude' updates (quaternion).

         
    */

    func (a *ServiceImpl) AttitudeQuaternion() (<-chan  *Quaternion , error){
    		ch := make(chan  *Quaternion )
    		request := &SubscribeAttitudeQuaternionRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeAttitudeQuaternion(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &AttitudeQuaternionResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetAttitudeQuaternion()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'attitude' updates (Euler).

         
    */

    func (a *ServiceImpl) AttitudeEuler() (<-chan  *EulerAngle , error){
    		ch := make(chan  *EulerAngle )
    		request := &SubscribeAttitudeEulerRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeAttitudeEuler(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &AttitudeEulerResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetAttitudeEuler()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'attitude' updates (angular velocity)

         
    */

    func (a *ServiceImpl) AttitudeAngularVelocityBody() (<-chan  *AngularVelocityBody , error){
    		ch := make(chan  *AngularVelocityBody )
    		request := &SubscribeAttitudeAngularVelocityBodyRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeAttitudeAngularVelocityBody(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &AttitudeAngularVelocityBodyResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetAttitudeAngularVelocityBody()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'camera attitude' updates (quaternion).

         
    */

    func (a *ServiceImpl) CameraAttitudeQuaternion() (<-chan  *Quaternion , error){
    		ch := make(chan  *Quaternion )
    		request := &SubscribeCameraAttitudeQuaternionRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeCameraAttitudeQuaternion(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &CameraAttitudeQuaternionResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetAttitudeQuaternion()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'camera attitude' updates (Euler).

         
    */

    func (a *ServiceImpl) CameraAttitudeEuler() (<-chan  *EulerAngle , error){
    		ch := make(chan  *EulerAngle )
    		request := &SubscribeCameraAttitudeEulerRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeCameraAttitudeEuler(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &CameraAttitudeEulerResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetAttitudeEuler()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'ground speed' updates (NED).

         
    */

    func (a *ServiceImpl) VelocityNed() (<-chan  *VelocityNed , error){
    		ch := make(chan  *VelocityNed )
    		request := &SubscribeVelocityNedRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeVelocityNed(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &VelocityNedResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetVelocityNed()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'GPS info' updates.

         
    */

    func (a *ServiceImpl) GpsInfo() (<-chan  *GpsInfo , error){
    		ch := make(chan  *GpsInfo )
    		request := &SubscribeGpsInfoRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeGpsInfo(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &GpsInfoResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetGpsInfo()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'battery' updates.

         
    */

    func (a *ServiceImpl) Battery() (<-chan  *Battery , error){
    		ch := make(chan  *Battery )
    		request := &SubscribeBatteryRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeBattery(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &BatteryResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetBattery()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'flight mode' updates.

         
    */

    func (a *ServiceImpl) FlightMode() (<-chan   FlightMode  , error){
    		ch := make(chan   FlightMode  )
    		request := &SubscribeFlightModeRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeFlightMode(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &FlightModeResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetFlightMode()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'health' updates.

         
    */

    func (a *ServiceImpl) Health() (<-chan  *Health , error){
    		ch := make(chan  *Health )
    		request := &SubscribeHealthRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeHealth(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &HealthResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetHealth()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'RC status' updates.

         
    */

    func (a *ServiceImpl) RcStatus() (<-chan  *RcStatus , error){
    		ch := make(chan  *RcStatus )
    		request := &SubscribeRcStatusRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeRcStatus(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &RcStatusResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetRcStatus()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'status text' updates.

         
    */

    func (a *ServiceImpl) StatusText() (<-chan  *StatusText , error){
    		ch := make(chan  *StatusText )
    		request := &SubscribeStatusTextRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeStatusText(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &StatusTextResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetStatusText()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'actuator control target' updates.

         
    */

    func (a *ServiceImpl) ActuatorControlTarget() (<-chan  *ActuatorControlTarget , error){
    		ch := make(chan  *ActuatorControlTarget )
    		request := &SubscribeActuatorControlTargetRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeActuatorControlTarget(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &ActuatorControlTargetResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetActuatorControlTarget()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'actuator output status' updates.

         
    */

    func (a *ServiceImpl) ActuatorOutputStatus() (<-chan  *ActuatorOutputStatus , error){
    		ch := make(chan  *ActuatorOutputStatus )
    		request := &SubscribeActuatorOutputStatusRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeActuatorOutputStatus(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &ActuatorOutputStatusResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetActuatorOutputStatus()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'odometry' updates.

         
    */

    func (a *ServiceImpl) Odometry() (<-chan  *Odometry , error){
    		ch := make(chan  *Odometry )
    		request := &SubscribeOdometryRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeOdometry(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &OdometryResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetOdometry()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'position velocity' updates.

         
    */

    func (a *ServiceImpl) PositionVelocityNed() (<-chan  *PositionVelocityNed , error){
    		ch := make(chan  *PositionVelocityNed )
    		request := &SubscribePositionVelocityNedRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribePositionVelocityNed(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &PositionVelocityNedResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetPositionVelocityNed()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'ground truth' updates.

         
    */

    func (a *ServiceImpl) GroundTruth() (<-chan  *GroundTruth , error){
    		ch := make(chan  *GroundTruth )
    		request := &SubscribeGroundTruthRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeGroundTruth(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &GroundTruthResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetGroundTruth()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'fixedwing metrics' updates.

         
    */

    func (a *ServiceImpl) FixedwingMetrics() (<-chan  *FixedwingMetrics , error){
    		ch := make(chan  *FixedwingMetrics )
    		request := &SubscribeFixedwingMetricsRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeFixedwingMetrics(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &FixedwingMetricsResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetFixedwingMetrics()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'IMU' updates.

         
    */

    func (a *ServiceImpl) Imu() (<-chan  *Imu , error){
    		ch := make(chan  *Imu )
    		request := &SubscribeImuRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeImu(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &ImuResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetImu()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'HealthAllOk' updates.

         
    */

    func (a *ServiceImpl) HealthAllOk() (<-chan   bool  , error){
    		ch := make(chan   bool  )
    		request := &SubscribeHealthAllOkRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeHealthAllOk(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &HealthAllOkResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetIsHealthAllOk()
    		}
    	}()	
    	return ch, nil
    }

     /*
         Subscribe to 'unix epoch time' updates.

         
    */

    func (a *ServiceImpl) UnixEpochTime() (<-chan   uint64  , error){
    		ch := make(chan   uint64  )
    		request := &SubscribeUnixEpochTimeRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeUnixEpochTime(ctx, request)
    		if err != nil {
    			return nil, err
    		}
    	go func() {
    		defer close(ch)
    		for {
    			m := &UnixEpochTimeResponse{}
    			err := stream.RecvMsg(m)
    			if err == io.EOF {
    				break
    			}
    			if err != nil {
    				fmt.Printf("Unable to receive message %v", err)
    				break
    			}
    			ch <- m.GetTimeUs()
    		}
    	}()	
    	return ch, nil
    }
    /*
         Set rate to 'position' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRatePosition(rateHz float64)(*SetRatePositionResponse, error){
        
        request := &SetRatePositionRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRatePosition(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'home position' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateHome(rateHz float64)(*SetRateHomeResponse, error){
        
        request := &SetRateHomeRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateHome(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to in-air updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateInAir(rateHz float64)(*SetRateInAirResponse, error){
        
        request := &SetRateInAirRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateInAir(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to landed state updates

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateLandedState(rateHz float64)(*SetRateLandedStateResponse, error){
        
        request := &SetRateLandedStateRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateLandedState(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'attitude' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateAttitude(rateHz float64)(*SetRateAttitudeResponse, error){
        
        request := &SetRateAttitudeRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateAttitude(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate of camera attitude updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateCameraAttitude(rateHz float64)(*SetRateCameraAttitudeResponse, error){
        
        request := &SetRateCameraAttitudeRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateCameraAttitude(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'ground speed' updates (NED).

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateVelocityNed(rateHz float64)(*SetRateVelocityNedResponse, error){
        
        request := &SetRateVelocityNedRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateVelocityNed(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'GPS info' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateGpsInfo(rateHz float64)(*SetRateGpsInfoResponse, error){
        
        request := &SetRateGpsInfoRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateGpsInfo(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'battery' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateBattery(rateHz float64)(*SetRateBatteryResponse, error){
        
        request := &SetRateBatteryRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateBattery(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'RC status' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateRcStatus(rateHz float64)(*SetRateRcStatusResponse, error){
        
        request := &SetRateRcStatusRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateRcStatus(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'actuator control target' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateActuatorControlTarget(rateHz float64)(*SetRateActuatorControlTargetResponse, error){
        
        request := &SetRateActuatorControlTargetRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateActuatorControlTarget(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'actuator output status' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateActuatorOutputStatus(rateHz float64)(*SetRateActuatorOutputStatusResponse, error){
        
        request := &SetRateActuatorOutputStatusRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateActuatorOutputStatus(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'odometry' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateOdometry(rateHz float64)(*SetRateOdometryResponse, error){
        
        request := &SetRateOdometryRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateOdometry(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'position velocity' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRatePositionVelocityNed(rateHz float64)(*SetRatePositionVelocityNedResponse, error){
        
        request := &SetRatePositionVelocityNedRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRatePositionVelocityNed(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'ground truth' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateGroundTruth(rateHz float64)(*SetRateGroundTruthResponse, error){
        
        request := &SetRateGroundTruthRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateGroundTruth(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'fixedwing metrics' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateFixedwingMetrics(rateHz float64)(*SetRateFixedwingMetricsResponse, error){
        
        request := &SetRateFixedwingMetricsRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateFixedwingMetrics(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'IMU' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateImu(rateHz float64)(*SetRateImuResponse, error){
        
        request := &SetRateImuRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateImu(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
    /*
         Set rate to 'unix epoch time' updates.

         Parameters
         ----------
         rateHz float64

         
    */

    func(s *ServiceImpl)SetRateUnixEpochTime(rateHz float64)(*SetRateUnixEpochTimeResponse, error){
        
        request := &SetRateUnixEpochTimeRequest{}
        ctx:= context.Background()
         request.RateHz = rateHz
        response, err := s.Client.SetRateUnixEpochTime(ctx, request)
        if err != nil {
    		return nil, err
        }
        return response, nil
    }

       
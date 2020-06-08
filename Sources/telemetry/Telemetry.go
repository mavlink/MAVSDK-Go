package telemetry

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client TelemetryServiceClient
}

    func (a *ServiceImpl) Position(){
    	request := &SubscribePositionRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribePosition(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetPosition())
    		}	
    }

    func (a *ServiceImpl) Home(){
    	request := &SubscribeHomeRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeHome(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetHome())
    		}	
    }

    func (a *ServiceImpl) InAir(){
    	request := &SubscribeInAirRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeInAir(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetInAir())
    		}	
    }

    func (a *ServiceImpl) LandedState(){
    	request := &SubscribeLandedStateRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeLandedState(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetLandedState())
    		}	
    }

    func (a *ServiceImpl) Armed(){
    	request := &SubscribeArmedRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeArmed(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetArmed())
    		}	
    }

    func (a *ServiceImpl) AttitudeQuaternion(){
    	request := &SubscribeAttitudeQuaternionRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeAttitudeQuaternion(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetAttitudeQuaternion())
    		}	
    }

    func (a *ServiceImpl) AttitudeEuler(){
    	request := &SubscribeAttitudeEulerRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeAttitudeEuler(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetAttitudeEuler())
    		}	
    }

    func (a *ServiceImpl) AttitudeAngularVelocityBody(){
    	request := &SubscribeAttitudeAngularVelocityBodyRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeAttitudeAngularVelocityBody(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetAttitudeAngularVelocityBody())
    		}	
    }

    func (a *ServiceImpl) CameraAttitudeQuaternion(){
    	request := &SubscribeCameraAttitudeQuaternionRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeCameraAttitudeQuaternion(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetCameraAttitudeQuaternion())
    		}	
    }

    func (a *ServiceImpl) CameraAttitudeEuler(){
    	request := &SubscribeCameraAttitudeEulerRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeCameraAttitudeEuler(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetCameraAttitudeEuler())
    		}	
    }

    func (a *ServiceImpl) VelocityNed(){
    	request := &SubscribeVelocityNedRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeVelocityNed(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetVelocityNed())
    		}	
    }

    func (a *ServiceImpl) GpsInfo(){
    	request := &SubscribeGpsInfoRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeGpsInfo(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetGpsInfo())
    		}	
    }

    func (a *ServiceImpl) Battery(){
    	request := &SubscribeBatteryRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeBattery(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetBattery())
    		}	
    }

    func (a *ServiceImpl) FlightMode(){
    	request := &SubscribeFlightModeRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeFlightMode(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetFlightMode())
    		}	
    }

    func (a *ServiceImpl) Health(){
    	request := &SubscribeHealthRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeHealth(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetHealth())
    		}	
    }

    func (a *ServiceImpl) RcStatus(){
    	request := &SubscribeRcStatusRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeRcStatus(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetRcStatus())
    		}	
    }

    func (a *ServiceImpl) StatusText(){
    	request := &SubscribeStatusTextRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeStatusText(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetStatusText())
    		}	
    }

    func (a *ServiceImpl) ActuatorControlTarget(){
    	request := &SubscribeActuatorControlTargetRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeActuatorControlTarget(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetActuatorControlTarget())
    		}	
    }

    func (a *ServiceImpl) ActuatorOutputStatus(){
    	request := &SubscribeActuatorOutputStatusRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeActuatorOutputStatus(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetActuatorOutputStatus())
    		}	
    }

    func (a *ServiceImpl) Odometry(){
    	request := &SubscribeOdometryRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeOdometry(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetOdometry())
    		}	
    }

    func (a *ServiceImpl) PositionVelocityNed(){
    	request := &SubscribePositionVelocityNedRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribePositionVelocityNed(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetPositionVelocityNed())
    		}	
    }

    func (a *ServiceImpl) GroundTruth(){
    	request := &SubscribeGroundTruthRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeGroundTruth(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetGroundTruth())
    		}	
    }

    func (a *ServiceImpl) FixedwingMetrics(){
    	request := &SubscribeFixedwingMetricsRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeFixedwingMetrics(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetFixedwingMetrics())
    		}	
    }

    func (a *ServiceImpl) Imu(){
    	request := &SubscribeImuRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeImu(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetImu())
    		}	
    }

    func (a *ServiceImpl) HealthAllOk(){
    	request := &SubscribeHealthAllOkRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeHealthAllOk(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetHealthAllOk())
    		}	
    }

    func (a *ServiceImpl) UnixEpochTime(){
    	request := &SubscribeUnixEpochTimeRequest{}
    		ctx := context.Background()
    		stream, err := a.Client.SubscribeUnixEpochTime(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

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
    			fmt.Printf("message %v\n", m.GetUnixEpochTime())
    		}	
    }

    func(s *ServiceImpl)SetRatePosition( rate_hz RateHz){
        request := &SetRatePositionRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRatePosition(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRatePosition grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRatePosition")
        }
        
    }

       

    func(s *ServiceImpl)SetRateHome( rate_hz RateHz){
        request := &SetRateHomeRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateHome(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateHome grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateHome")
        }
        
    }

       

    func(s *ServiceImpl)SetRateInAir( rate_hz RateHz){
        request := &SetRateInAirRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateInAir(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateInAir grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateInAir")
        }
        
    }

       

    func(s *ServiceImpl)SetRateLandedState( rate_hz RateHz){
        request := &SetRateLandedStateRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateLandedState(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateLandedState grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateLandedState")
        }
        
    }

       

    func(s *ServiceImpl)SetRateAttitude( rate_hz RateHz){
        request := &SetRateAttitudeRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateAttitude(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateAttitude grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateAttitude")
        }
        
    }

       

    func(s *ServiceImpl)SetRateCameraAttitude( rate_hz RateHz){
        request := &SetRateCameraAttitudeRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateCameraAttitude(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateCameraAttitude grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateCameraAttitude")
        }
        
    }

       

    func(s *ServiceImpl)SetRateVelocityNed( rate_hz RateHz){
        request := &SetRateVelocityNedRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateVelocityNed(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateVelocityNed grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateVelocityNed")
        }
        
    }

       

    func(s *ServiceImpl)SetRateGpsInfo( rate_hz RateHz){
        request := &SetRateGpsInfoRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateGpsInfo(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateGpsInfo grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateGpsInfo")
        }
        
    }

       

    func(s *ServiceImpl)SetRateBattery( rate_hz RateHz){
        request := &SetRateBatteryRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateBattery(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateBattery grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateBattery")
        }
        
    }

       

    func(s *ServiceImpl)SetRateRcStatus( rate_hz RateHz){
        request := &SetRateRcStatusRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateRcStatus(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateRcStatus grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateRcStatus")
        }
        
    }

       

    func(s *ServiceImpl)SetRateActuatorControlTarget( rate_hz RateHz){
        request := &SetRateActuatorControlTargetRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateActuatorControlTarget(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateActuatorControlTarget grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateActuatorControlTarget")
        }
        
    }

       

    func(s *ServiceImpl)SetRateActuatorOutputStatus( rate_hz RateHz){
        request := &SetRateActuatorOutputStatusRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateActuatorOutputStatus(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateActuatorOutputStatus grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateActuatorOutputStatus")
        }
        
    }

       

    func(s *ServiceImpl)SetRateOdometry( rate_hz RateHz){
        request := &SetRateOdometryRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateOdometry(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateOdometry grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateOdometry")
        }
        
    }

       

    func(s *ServiceImpl)SetRatePositionVelocityNed( rate_hz RateHz){
        request := &SetRatePositionVelocityNedRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRatePositionVelocityNed(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRatePositionVelocityNed grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRatePositionVelocityNed")
        }
        
    }

       

    func(s *ServiceImpl)SetRateGroundTruth( rate_hz RateHz){
        request := &SetRateGroundTruthRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateGroundTruth(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateGroundTruth grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateGroundTruth")
        }
        
    }

       

    func(s *ServiceImpl)SetRateFixedwingMetrics( rate_hz RateHz){
        request := &SetRateFixedwingMetricsRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateFixedwingMetrics(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateFixedwingMetrics grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateFixedwingMetrics")
        }
        
    }

       

    func(s *ServiceImpl)SetRateImu( rate_hz RateHz){
        request := &SetRateImuRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateImu(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateImu grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateImu")
        }
        
    }

       

    func(s *ServiceImpl)SetRateUnixEpochTime( rate_hz RateHz){
        request := &SetRateUnixEpochTimeRequest{}
        ctx:= context.Background()
         request.RateHz = rate_hz
        response, err := s.Client.SetRateUnixEpochTime(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetRateUnixEpochTime grpc %v\n", err)
    	}
        
        result := response.GetTelemetryResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetRateUnixEpochTime")
        }
        
    }

       
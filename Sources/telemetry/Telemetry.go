import (
	"context"
	"fmt"
	"io"
)

type Service interface{
    Result() TelemetryResult_Result

}

type ServiceImpl struct{
    Client TelemetryServiceClient
}

    func (a *ServiceImpl) Position(){
    	request := &PositionRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribePosition(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &positionResponse{}
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
    	request := &HomeRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeHome(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &homeResponse{}
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
    	request := &InAirRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeInAir(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &in_airResponse{}
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
    	request := &LandedStateRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeLandedState(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &landed_stateResponse{}
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
    	request := &ArmedRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeArmed(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &armedResponse{}
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
    	request := &AttitudeQuaternionRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeAttitudeQuaternion(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &attitude_quaternionResponse{}
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
    	request := &AttitudeEulerRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeAttitudeEuler(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &attitude_eulerResponse{}
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
    	request := &AttitudeAngularVelocityBodyRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeAttitudeAngularVelocityBody(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &attitude_angular_velocity_bodyResponse{}
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
    	request := &CameraAttitudeQuaternionRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeCameraAttitudeQuaternion(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &camera_attitude_quaternionResponse{}
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
    	request := &CameraAttitudeEulerRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeCameraAttitudeEuler(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &camera_attitude_eulerResponse{}
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
    	request := &VelocityNedRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeVelocityNed(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &velocity_nedResponse{}
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
    	request := &GpsInfoRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeGpsInfo(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &gps_infoResponse{}
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
    	request := &BatteryRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeBattery(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &batteryResponse{}
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
    	request := &FlightModeRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeFlightMode(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &flight_modeResponse{}
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
    	request := &HealthRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeHealth(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &healthResponse{}
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
    	request := &RcStatusRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeRcStatus(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &rc_statusResponse{}
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
    	request := &StatusTextRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeStatusText(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &status_textResponse{}
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
    	request := &ActuatorControlTargetRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeActuatorControlTarget(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &actuator_control_targetResponse{}
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
    	request := &ActuatorOutputStatusRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeActuatorOutputStatus(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &actuator_output_statusResponse{}
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
    	request := &OdometryRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeOdometry(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &odometryResponse{}
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
    	request := &PositionVelocityNedRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribePositionVelocityNed(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &position_velocity_nedResponse{}
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
    	request := &GroundTruthRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeGroundTruth(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &ground_truthResponse{}
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
    	request := &FixedwingMetricsRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeFixedwingMetrics(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &fixedwing_metricsResponse{}
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
    	request := &ImuRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeImu(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &imuResponse{}
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
    	request := &HealthAllOkRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeHealthAllOk(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &health_all_okResponse{}
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
    	request := &UnixEpochTimeRequest{}
    		ctx := context.Background()
    		stream, err := a.Telemetry.SubscribeUnixEpochTime(ctx, request)
    		if err != nil {
    			fmt.Printf("Unable to subscribe %v\n", err)
    		}

    		for {
    			m := &unix_epoch_timeResponse{}
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

    func(s *ServiceImpl)set_rate_position(self, rate_hz []*RateHz){
     request = &SetRatePositionRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRatePosition(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRatePosition")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_home(self, rate_hz []*RateHz){
     request = &SetRateHomeRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateHome(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateHome")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_in_air(self, rate_hz []*RateHz){
     request = &SetRateInAirRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateInAir(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateInAir")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_landed_state(self, rate_hz []*RateHz){
     request = &SetRateLandedStateRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateLandedState(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateLandedState")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_attitude(self, rate_hz []*RateHz){
     request = &SetRateAttitudeRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateAttitude(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateAttitude")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_camera_attitude(self, rate_hz []*RateHz){
     request = &SetRateCameraAttitudeRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateCameraAttitude(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateCameraAttitude")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_velocity_ned(self, rate_hz []*RateHz){
     request = &SetRateVelocityNedRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateVelocityNed(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateVelocityNed")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_gps_info(self, rate_hz []*RateHz){
     request = &SetRateGpsInfoRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateGpsInfo(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateGpsInfo")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_battery(self, rate_hz []*RateHz){
     request = &SetRateBatteryRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateBattery(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateBattery")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_rc_status(self, rate_hz []*RateHz){
     request = &SetRateRcStatusRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateRcStatus(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateRcStatus")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_actuator_control_target(self, rate_hz []*RateHz){
     request = &SetRateActuatorControlTargetRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateActuatorControlTarget(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateActuatorControlTarget")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_actuator_output_status(self, rate_hz []*RateHz){
     request = &SetRateActuatorOutputStatusRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateActuatorOutputStatus(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateActuatorOutputStatus")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_odometry(self, rate_hz []*RateHz){
     request = &SetRateOdometryRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateOdometry(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateOdometry")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_position_velocity_ned(self, rate_hz []*RateHz){
     request = &SetRatePositionVelocityNedRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRatePositionVelocityNed(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRatePositionVelocityNed")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_ground_truth(self, rate_hz []*RateHz){
     request = &SetRateGroundTruthRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateGroundTruth(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateGroundTruth")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_fixedwing_metrics(self, rate_hz []*RateHz){
     request = &SetRateFixedwingMetricsRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateFixedwingMetrics(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateFixedwingMetrics")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_imu(self, rate_hz []*RateHz){
     request = &SetRateImuRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateImu(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateImu")
        }
        
    }

       

    func(s *ServiceImpl)set_rate_unix_epoch_time(self, rate_hz []*RateHz){
     request = &SetRateUnixEpochTimeRequest{}
         request.rate_hz = rate_hz
        response, err = s.Client.SetRateUnixEpochTime(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetTelemetryResult()
        fmt.Printf("result %v\n", result)
        if result.Result != TelemetryResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetRateUnixEpochTime")
        }
        
    }

       
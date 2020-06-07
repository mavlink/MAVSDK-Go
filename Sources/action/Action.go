import (
	"context"
	"fmt"
	"io"
)

type Service interface{
    Result() ActionResult_Result

}

type ServiceImpl struct{
    Client ActionServiceClient
}

    func(s *ServiceImpl)arm(self){
     request = &ArmRequest{}
         response, err = s.Client.Arm(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading Arm")
        }
        
    }

       

    func(s *ServiceImpl)disarm(self){
     request = &DisarmRequest{}
         response, err = s.Client.Disarm(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading Disarm")
        }
        
    }

       

    func(s *ServiceImpl)takeoff(self){
     request = &TakeoffRequest{}
         response, err = s.Client.Takeoff(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading Takeoff")
        }
        
    }

       

    func(s *ServiceImpl)land(self){
     request = &LandRequest{}
         response, err = s.Client.Land(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading Land")
        }
        
    }

       

    func(s *ServiceImpl)reboot(self){
     request = &RebootRequest{}
         response, err = s.Client.Reboot(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading Reboot")
        }
        
    }

       

    func(s *ServiceImpl)shutdown(self){
     request = &ShutdownRequest{}
         response, err = s.Client.Shutdown(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading Shutdown")
        }
        
    }

       

    func(s *ServiceImpl)kill(self){
     request = &KillRequest{}
         response, err = s.Client.Kill(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading Kill")
        }
        
    }

       

    func(s *ServiceImpl)return_to_launch(self){
     request = &ReturnToLaunchRequest{}
         response, err = s.Client.ReturnToLaunch(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading ReturnToLaunch")
        }
        
    }

       

    func(s *ServiceImpl)goto_location(self, latitude_deg []*LatitudeDeg, longitude_deg []*LongitudeDeg, absolute_altitude_m []*AbsoluteAltitudeM, yaw_deg []*YawDeg){
     request = &GotoLocationRequest{}
         request.latitude_deg = latitude_deg
        request.longitude_deg = longitude_deg
        request.absolute_altitude_m = absolute_altitude_m
        request.yaw_deg = yaw_deg
        response, err = s.Client.GotoLocation(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading GotoLocation")
        }
        
    }

       

    func(s *ServiceImpl)transition_to_fixedwing(self){
     request = &TransitionToFixedwingRequest{}
         response, err = s.Client.TransitionToFixedwing(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading TransitionToFixedwing")
        }
        
    }

       

    func(s *ServiceImpl)transition_to_multicopter(self){
     request = &TransitionToMulticopterRequest{}
         response, err = s.Client.TransitionToMulticopter(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading TransitionToMulticopter")
        }
        
    }

       

    func(s *ServiceImpl)get_takeoff_altitude(self)([]*GetTakeoffAltitude){
     request = &GetTakeoffAltitudeRequest{}
         response, err = s.Client.GetTakeoffAltitude(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading GetTakeoffAltitude")
        }
        

        return response.GetGetTakeoffAltitude()
        
    }

       

    func(s *ServiceImpl)set_takeoff_altitude(self, altitude []*Altitude){
     request = &SetTakeoffAltitudeRequest{}
         request.altitude = altitude
        response, err = s.Client.SetTakeoffAltitude(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetTakeoffAltitude")
        }
        
    }

       

    func(s *ServiceImpl)get_maximum_speed(self)([]*GetMaximumSpeed){
     request = &GetMaximumSpeedRequest{}
         response, err = s.Client.GetMaximumSpeed(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading GetMaximumSpeed")
        }
        

        return response.GetGetMaximumSpeed()
        
    }

       

    func(s *ServiceImpl)set_maximum_speed(self, speed []*Speed){
     request = &SetMaximumSpeedRequest{}
         request.speed = speed
        response, err = s.Client.SetMaximumSpeed(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetMaximumSpeed")
        }
        
    }

       

    func(s *ServiceImpl)get_return_to_launch_altitude(self)([]*GetReturnToLaunchAltitude){
     request = &GetReturnToLaunchAltitudeRequest{}
         response, err = s.Client.GetReturnToLaunchAltitude(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading GetReturnToLaunchAltitude")
        }
        

        return response.GetGetReturnToLaunchAltitude()
        
    }

       

    func(s *ServiceImpl)set_return_to_launch_altitude(self, relative_altitude_m []*RelativeAltitudeM){
     request = &SetReturnToLaunchAltitudeRequest{}
         request.relative_altitude_m = relative_altitude_m
        response, err = s.Client.SetReturnToLaunchAltitude(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result = response.GetActionResult()
        fmt.Printf("result %v\n", result)
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while uploading SetReturnToLaunchAltitude")
        }
        
    }

       
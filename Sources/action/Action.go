package action

import (
	"context"
	"fmt"
	"io"
)

type ServiceImpl struct{
    Client ActionServiceClient
}

    func(s *ServiceImpl)Arm(){
        request := &ArmRequest{}
        ctx:= context.Background()
         response, err := s.Client.Arm(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing Arm grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for Arm")
        }
        
    }

       

    func(s *ServiceImpl)Disarm(){
        request := &DisarmRequest{}
        ctx:= context.Background()
         response, err := s.Client.Disarm(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing Disarm grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for Disarm")
        }
        
    }

       

    func(s *ServiceImpl)Takeoff(){
        request := &TakeoffRequest{}
        ctx:= context.Background()
         response, err := s.Client.Takeoff(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing Takeoff grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for Takeoff")
        }
        
    }

       

    func(s *ServiceImpl)Land(){
        request := &LandRequest{}
        ctx:= context.Background()
         response, err := s.Client.Land(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing Land grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for Land")
        }
        
    }

       

    func(s *ServiceImpl)Reboot(){
        request := &RebootRequest{}
        ctx:= context.Background()
         response, err := s.Client.Reboot(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing Reboot grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for Reboot")
        }
        
    }

       

    func(s *ServiceImpl)Shutdown(){
        request := &ShutdownRequest{}
        ctx:= context.Background()
         response, err := s.Client.Shutdown(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing Shutdown grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for Shutdown")
        }
        
    }

       

    func(s *ServiceImpl)Kill(){
        request := &KillRequest{}
        ctx:= context.Background()
         response, err := s.Client.Kill(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing Kill grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for Kill")
        }
        
    }

       

    func(s *ServiceImpl)ReturnToLaunch(){
        request := &ReturnToLaunchRequest{}
        ctx:= context.Background()
         response, err := s.Client.ReturnToLaunch(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing ReturnToLaunch grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for ReturnToLaunch")
        }
        
    }

       

    func(s *ServiceImpl)GotoLocation( latitude_deg LatitudeDeg longitude_deg LongitudeDeg absolute_altitude_m AbsoluteAltitudeM yaw_deg YawDeg){
        request := &GotoLocationRequest{}
        ctx:= context.Background()
         request.LatitudeDeg = latitude_deg
        request.LongitudeDeg = longitude_deg
        request.AbsoluteAltitudeM = absolute_altitude_m
        request.YawDeg = yaw_deg
        response, err := s.Client.GotoLocation(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing GotoLocation grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for GotoLocation")
        }
        
    }

       

    func(s *ServiceImpl)TransitionToFixedwing(){
        request := &TransitionToFixedwingRequest{}
        ctx:= context.Background()
         response, err := s.Client.TransitionToFixedwing(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing TransitionToFixedwing grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for TransitionToFixedwing")
        }
        
    }

       

    func(s *ServiceImpl)TransitionToMulticopter(){
        request := &TransitionToMulticopterRequest{}
        ctx:= context.Background()
         response, err := s.Client.TransitionToMulticopter(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing TransitionToMulticopter grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for TransitionToMulticopter")
        }
        
    }

       

    func(s *ServiceImpl)GetTakeoffAltitude() (*GetTakeoffAltitudeResponse){
        request := &GetTakeoffAltitudeRequest{}
        ctx:= context.Background()
         response, err := s.Client.GetTakeoffAltitude(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while getting GetTakeoffAltitude")
        }
        
        return response

    }

       

    func(s *ServiceImpl)SetTakeoffAltitude( altitude Altitude){
        request := &SetTakeoffAltitudeRequest{}
        ctx:= context.Background()
         request.Altitude = altitude
        response, err := s.Client.SetTakeoffAltitude(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetTakeoffAltitude grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetTakeoffAltitude")
        }
        
    }

       

    func(s *ServiceImpl)GetMaximumSpeed() (*GetMaximumSpeedResponse){
        request := &GetMaximumSpeedRequest{}
        ctx:= context.Background()
         response, err := s.Client.GetMaximumSpeed(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while getting GetMaximumSpeed")
        }
        
        return response

    }

       

    func(s *ServiceImpl)SetMaximumSpeed( speed Speed){
        request := &SetMaximumSpeedRequest{}
        ctx:= context.Background()
         request.Speed = speed
        response, err := s.Client.SetMaximumSpeed(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetMaximumSpeed grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetMaximumSpeed")
        }
        
    }

       

    func(s *ServiceImpl)GetReturnToLaunchAltitude() (*GetReturnToLaunchAltitudeResponse){
        request := &GetReturnToLaunchAltitudeRequest{}
        ctx:= context.Background()
         response, err := s.Client.GetReturnToLaunchAltitude(ctx, request)
        if err != nil {
    		fmt.Printf("Unable to subscribe to position grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while getting GetReturnToLaunchAltitude")
        }
        
        return response

    }

       

    func(s *ServiceImpl)SetReturnToLaunchAltitude( relative_altitude_m RelativeAltitudeM){
        request := &SetReturnToLaunchAltitudeRequest{}
        ctx:= context.Background()
         request.RelativeAltitudeM = relative_altitude_m
        response, err := s.Client.SetReturnToLaunchAltitude(ctx, request)
        if err != nil {
    		fmt.Printf("Error while performing SetReturnToLaunchAltitude grpc %v\n", err)
    	}
        
        result := response.GetActionResult()
        fmt.Printf("result %v\n",  result.String())
        if result.Result != ActionResult_RESULT_SUCCESS{
            fmt.Printf("Error while extracting result for SetReturnToLaunchAltitude")
        }
        
    }

       
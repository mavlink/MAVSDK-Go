// package action

import (
	"context"
	"fmt"
)

// //Service interface for ActionService
// type Service interface {
// 	Result()
// }

// //ServiceImpl creates client and implements Service
// type ServiceImpl struct {
// 	Client ActionServiceClient
// }

//GotoLocation acts as a wrapper for GotoLocation request from action grpc package
func (a *ServiceImpl) GotoLocation(latitude_deg, longitude_deg float64, absolute_altitude_m, yaw_deg float32) {
	request := &GotoLocationRequest{}
	request.LatitudeDeg = latitude_deg
	request.LongitudeDeg = longitude_deg
	request.AbsoluteAltitudeM = absolute_altitude_m
	request.YawDeg = yaw_deg
	ctx := context.Background()
	response, err := a.Client.GotoLocation(ctx, request)
	if err != nil {
		fmt.Printf("Set send GotoLocation command %v\n", err)
	}
	result := response.GetActionResult()
	if result.Result != ActionResult_RESULT_SUCCESS {
		fmt.Printf("Error while setting new location")
	}
	fmt.Printf("Response %v\n", response)
}

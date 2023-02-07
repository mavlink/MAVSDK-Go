package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mavlink/MAVSDK-Go/Sources/action"
	"github.com/mavlink/MAVSDK-Go/Sources/core"
	"github.com/mavlink/MAVSDK-Go/Sources/geofence"
	"github.com/mavlink/MAVSDK-Go/Sources/telemetry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//Drone creates a drone object to interact with drone related plugins
type Drone struct {
	port         string
	mavsdkServer string
	action       action.ServiceImpl
	core         core.ServiceImpl
	telemetry    telemetry.ServiceImpl
	geofence     geofence.ServiceImpl
}

//Connect Starts a mavsdk server and create a connection to it
func (s *Drone) Connect() {
	grpcConnection := s.connectToMAVSDKServer()
	s.InitPlugins(grpcConnection)

}

//InitPlugins initializes all the plugins
func (s *Drone) InitPlugins(cc *grpc.ClientConn) {

	s.telemetry = telemetry.ServiceImpl{
		Client: telemetry.NewTelemetryServiceClient(cc),
	}
	s.core = core.ServiceImpl{
		Client: core.NewCoreServiceClient(cc),
	}
	s.action = action.ServiceImpl{
		Client: action.NewActionServiceClient(cc),
	}
	s.action = action.ServiceImpl{
		Client: action.NewActionServiceClient(cc),
	}
	s.geofence = geofence.ServiceImpl{
		Client: geofence.NewGeofenceServiceClient(cc),
	}
}

func (s *Drone) connectToMAVSDKServer() *grpc.ClientConn {
	dialoption := grpc.WithTransportCredentials(insecure.NewCredentials())

	serverAddr := s.mavsdkServer + ":" + s.port
	cc, err := grpc.Dial(serverAddr, dialoption)
	if err != nil {
		fmt.Printf("Error while dialing %v", err)
	}
	grpc.ConnectionTimeout(5)
	return cc
}

func main() {
	drone := &Drone{port: "50051", mavsdkServer: "127.0.0.1"}
	drone.Connect()
	drone.action.Arm(context.Background())
	drone.action.Takeoff(context.Background())
	drone.action.Land(context.Background())
	drone.core.ConnectionState(context.Background())
	lat := 47.3977508
	lon := 8.5456074
	p1 := &geofence.Point{
		LatitudeDeg:  lat - 0.0001,
		LongitudeDeg: lon - 0.0001,
	}
	p2 := &geofence.Point{
		LatitudeDeg:  lat + 0.0001,
		LongitudeDeg: lon - 0.0001,
	}
	p3 := &geofence.Point{
		LatitudeDeg:  lat + 0.0001,
		LongitudeDeg: lon + 0.0001,
	}
	p4 := &geofence.Point{
		LatitudeDeg:  lat - 0.0001,
		LongitudeDeg: lon + 0.0001,
	}
	// this is not a test or verification package. this only checks the sanity of geofence api
	polygon := &geofence.Polygon{
		Points:    []*geofence.Point{p1, p2, p3, p4},
		FenceType: geofence.FenceType_FENCE_TYPE_EXCLUSION}
	response, err := drone.geofence.UploadGeofence(context.Background(), &geofence.GeofenceData{
		Polygons: []*geofence.Polygon{polygon},
	})
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
	log.Printf("response %v", response)
}

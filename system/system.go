package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ykhedar/MAVSDK-Go/sources/core"
	"github.com/ykhedar/MAVSDK-Go/sources/geofence"
	"github.com/ykhedar/MAVSDK-Go/sources/log_files"
	"github.com/ykhedar/MAVSDK-Go/sources/mission"
	"github.com/ykhedar/MAVSDK-Go/sources/sample"
	"google.golang.org/grpc"
)

//System provides interface to interract with a drone
type System interface {
	Connect()
	initPlugins()
	startMAVSDKServer()
	connectToMAVSDKServer() *grpc.ClientConn
}

//Drone creates a drone object to interact with drone related plugins
type Drone struct {
	port         string
	mavsdkServer string
	core         core.CoreServiceClient
	geofence     geofence.GeofenceServiceClient
	logFiles     log_files.LogFilesServiceClient
	mission      mission.MissionServiceClient
	telemetry    sample.TelemetryServiceImpl
}

//Connect Starts a mavsdk server and create a connection to it
func (s *Drone) Connect() {
	//start mavsdk server
	// s.startMAVSDKServer()
	grpcConnection := s.connectToMAVSDKServer()
	s.initPlugins(grpcConnection)

}

//initPlugins initializes all the plugins
func (s *Drone) initPlugins(cc *grpc.ClientConn) {

	s.telemetry = sample.TelemetryServiceImpl{
		Name:   "action",
		Client: cc,
	}
	s.core = core.NewCoreServiceClient(cc)
	s.geofence = geofence.NewGeofenceServiceClient(cc)
	s.logFiles = log_files.NewLogFilesServiceClient(cc)
	s.mission = mission.NewMissionServiceClient(cc)
}

func (s *Drone) startMAVSDKServer() {
	cmd := exec.Command("mavsdk_server_win32", "-p", "50051")
	print("Command [%s]", cmd)
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		print("Error starting mavsdkserver %s", err)
	}
}

func (s *Drone) connectToMAVSDKServer() *grpc.ClientConn {
	dialoption := grpc.WithInsecure()

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
	drone.telemetry.InitTelemetry()
	drone.telemetry.Position()

}

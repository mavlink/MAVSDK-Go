package core

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

//Service interface for intertelemetry with telemetry interface of drone
type Service interface {
	getResult()
	InitTelemetry(*grpc.ClientConn) CoreServiceClient
}

//ServiceImpl creates client and implements telemetryService
type ServiceImpl struct {
	Name        string
	Client      *grpc.ClientConn
	coreService CoreServiceClient
}

//InitCore initializes telemetry client
func (a *ServiceImpl) InitCore() {
	a.coreService = NewCoreServiceClient(a.Client)
}

//ListRunningPlugins acts as a wrapper for Position request from telemetry grpc package
func (a *ServiceImpl) ListRunningPlugins() {
	request := &ListRunningPluginsRequest{}
	ctx := context.Background()
	response, err := a.coreService.ListRunningPlugins(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}
	pluginInfo := []*PluginInfo{}
	for _, plugin := range response.PluginInfo {
		pluginInfo = append(pluginInfo, plugin)
	}
	fmt.Printf("plugins %v\n", pluginInfo)
	// return response, err
}

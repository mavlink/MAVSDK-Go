package core

import (
	"context"
	"fmt"
)

//Service interface for intertelemetry with telemetry interface of drone
type Service interface {
	Result()
}

//ServiceImpl creates client and implements telemetryService
type ServiceImpl struct {
	Client CoreServiceClient
}

//ListRunningPlugins acts as a wrapper for Position request from telemetry grpc package
func (a *ServiceImpl) ListRunningPlugins() []*PluginInfo {
	request := &ListRunningPluginsRequest{}
	ctx := context.Background()
	response, err := a.Client.ListRunningPlugins(ctx, request)
	if err != nil {
		fmt.Printf("Unable to subscribe %v\n", err)
	}

	fmt.Printf("plugins %v\n", response.GetPluginInfo())
	// return response, err
	return response.GetPluginInfo()
}

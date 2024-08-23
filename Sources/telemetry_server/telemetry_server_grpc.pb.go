// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: telemetry_server.proto

package telemetry_server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TelemetryServerService_PublishPosition_FullMethodName            = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishPosition"
	TelemetryServerService_PublishHome_FullMethodName                = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishHome"
	TelemetryServerService_PublishSysStatus_FullMethodName           = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishSysStatus"
	TelemetryServerService_PublishExtendedSysState_FullMethodName    = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishExtendedSysState"
	TelemetryServerService_PublishRawGps_FullMethodName              = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishRawGps"
	TelemetryServerService_PublishBattery_FullMethodName             = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishBattery"
	TelemetryServerService_PublishStatusText_FullMethodName          = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishStatusText"
	TelemetryServerService_PublishOdometry_FullMethodName            = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishOdometry"
	TelemetryServerService_PublishPositionVelocityNed_FullMethodName = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishPositionVelocityNed"
	TelemetryServerService_PublishGroundTruth_FullMethodName         = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishGroundTruth"
	TelemetryServerService_PublishImu_FullMethodName                 = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishImu"
	TelemetryServerService_PublishScaledImu_FullMethodName           = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishScaledImu"
	TelemetryServerService_PublishRawImu_FullMethodName              = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishRawImu"
	TelemetryServerService_PublishUnixEpochTime_FullMethodName       = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishUnixEpochTime"
	TelemetryServerService_PublishDistanceSensor_FullMethodName      = "/mavsdk.rpc.telemetry_server.TelemetryServerService/PublishDistanceSensor"
)

// TelemetryServerServiceClient is the client API for TelemetryServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Allow users to provide vehicle telemetry and state information
// (e.g. battery, GPS, RC connection, flight mode etc.) and set telemetry update rates.
type TelemetryServerServiceClient interface {
	// Publish to 'position' updates.
	PublishPosition(ctx context.Context, in *PublishPositionRequest, opts ...grpc.CallOption) (*PublishPositionResponse, error)
	// Publish to 'home position' updates.
	PublishHome(ctx context.Context, in *PublishHomeRequest, opts ...grpc.CallOption) (*PublishHomeResponse, error)
	// Publish 'sys status' updates.
	PublishSysStatus(ctx context.Context, in *PublishSysStatusRequest, opts ...grpc.CallOption) (*PublishSysStatusResponse, error)
	// Publish 'extended sys state' updates.
	PublishExtendedSysState(ctx context.Context, in *PublishExtendedSysStateRequest, opts ...grpc.CallOption) (*PublishExtendedSysStateResponse, error)
	// Publish to 'Raw GPS' updates.
	PublishRawGps(ctx context.Context, in *PublishRawGpsRequest, opts ...grpc.CallOption) (*PublishRawGpsResponse, error)
	// Publish to 'battery' updates.
	PublishBattery(ctx context.Context, in *PublishBatteryRequest, opts ...grpc.CallOption) (*PublishBatteryResponse, error)
	// Publish to 'status text' updates.
	PublishStatusText(ctx context.Context, in *PublishStatusTextRequest, opts ...grpc.CallOption) (*PublishStatusTextResponse, error)
	// Publish to 'odometry' updates.
	PublishOdometry(ctx context.Context, in *PublishOdometryRequest, opts ...grpc.CallOption) (*PublishOdometryResponse, error)
	// Publish to 'position velocity' updates.
	PublishPositionVelocityNed(ctx context.Context, in *PublishPositionVelocityNedRequest, opts ...grpc.CallOption) (*PublishPositionVelocityNedResponse, error)
	// Publish to 'ground truth' updates.
	PublishGroundTruth(ctx context.Context, in *PublishGroundTruthRequest, opts ...grpc.CallOption) (*PublishGroundTruthResponse, error)
	// Publish to 'IMU' updates (in SI units in NED body frame).
	PublishImu(ctx context.Context, in *PublishImuRequest, opts ...grpc.CallOption) (*PublishImuResponse, error)
	// Publish to 'Scaled IMU' updates.
	PublishScaledImu(ctx context.Context, in *PublishScaledImuRequest, opts ...grpc.CallOption) (*PublishScaledImuResponse, error)
	// Publish to 'Raw IMU' updates.
	PublishRawImu(ctx context.Context, in *PublishRawImuRequest, opts ...grpc.CallOption) (*PublishRawImuResponse, error)
	// Publish to 'unix epoch time' updates.
	PublishUnixEpochTime(ctx context.Context, in *PublishUnixEpochTimeRequest, opts ...grpc.CallOption) (*PublishUnixEpochTimeResponse, error)
	// Publish to "distance sensor" updates.
	PublishDistanceSensor(ctx context.Context, in *PublishDistanceSensorRequest, opts ...grpc.CallOption) (*PublishDistanceSensorResponse, error)
}

type telemetryServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTelemetryServerServiceClient(cc grpc.ClientConnInterface) TelemetryServerServiceClient {
	return &telemetryServerServiceClient{cc}
}

func (c *telemetryServerServiceClient) PublishPosition(ctx context.Context, in *PublishPositionRequest, opts ...grpc.CallOption) (*PublishPositionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishPositionResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishPosition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishHome(ctx context.Context, in *PublishHomeRequest, opts ...grpc.CallOption) (*PublishHomeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishHomeResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishHome_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishSysStatus(ctx context.Context, in *PublishSysStatusRequest, opts ...grpc.CallOption) (*PublishSysStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishSysStatusResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishSysStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishExtendedSysState(ctx context.Context, in *PublishExtendedSysStateRequest, opts ...grpc.CallOption) (*PublishExtendedSysStateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishExtendedSysStateResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishExtendedSysState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishRawGps(ctx context.Context, in *PublishRawGpsRequest, opts ...grpc.CallOption) (*PublishRawGpsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishRawGpsResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishRawGps_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishBattery(ctx context.Context, in *PublishBatteryRequest, opts ...grpc.CallOption) (*PublishBatteryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishBatteryResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishBattery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishStatusText(ctx context.Context, in *PublishStatusTextRequest, opts ...grpc.CallOption) (*PublishStatusTextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishStatusTextResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishStatusText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishOdometry(ctx context.Context, in *PublishOdometryRequest, opts ...grpc.CallOption) (*PublishOdometryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishOdometryResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishOdometry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishPositionVelocityNed(ctx context.Context, in *PublishPositionVelocityNedRequest, opts ...grpc.CallOption) (*PublishPositionVelocityNedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishPositionVelocityNedResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishPositionVelocityNed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishGroundTruth(ctx context.Context, in *PublishGroundTruthRequest, opts ...grpc.CallOption) (*PublishGroundTruthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishGroundTruthResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishGroundTruth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishImu(ctx context.Context, in *PublishImuRequest, opts ...grpc.CallOption) (*PublishImuResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishImuResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishImu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishScaledImu(ctx context.Context, in *PublishScaledImuRequest, opts ...grpc.CallOption) (*PublishScaledImuResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishScaledImuResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishScaledImu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishRawImu(ctx context.Context, in *PublishRawImuRequest, opts ...grpc.CallOption) (*PublishRawImuResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishRawImuResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishRawImu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishUnixEpochTime(ctx context.Context, in *PublishUnixEpochTimeRequest, opts ...grpc.CallOption) (*PublishUnixEpochTimeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishUnixEpochTimeResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishUnixEpochTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telemetryServerServiceClient) PublishDistanceSensor(ctx context.Context, in *PublishDistanceSensorRequest, opts ...grpc.CallOption) (*PublishDistanceSensorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishDistanceSensorResponse)
	err := c.cc.Invoke(ctx, TelemetryServerService_PublishDistanceSensor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelemetryServerServiceServer is the server API for TelemetryServerService service.
// All implementations must embed UnimplementedTelemetryServerServiceServer
// for forward compatibility.
//
// Allow users to provide vehicle telemetry and state information
// (e.g. battery, GPS, RC connection, flight mode etc.) and set telemetry update rates.
type TelemetryServerServiceServer interface {
	// Publish to 'position' updates.
	PublishPosition(context.Context, *PublishPositionRequest) (*PublishPositionResponse, error)
	// Publish to 'home position' updates.
	PublishHome(context.Context, *PublishHomeRequest) (*PublishHomeResponse, error)
	// Publish 'sys status' updates.
	PublishSysStatus(context.Context, *PublishSysStatusRequest) (*PublishSysStatusResponse, error)
	// Publish 'extended sys state' updates.
	PublishExtendedSysState(context.Context, *PublishExtendedSysStateRequest) (*PublishExtendedSysStateResponse, error)
	// Publish to 'Raw GPS' updates.
	PublishRawGps(context.Context, *PublishRawGpsRequest) (*PublishRawGpsResponse, error)
	// Publish to 'battery' updates.
	PublishBattery(context.Context, *PublishBatteryRequest) (*PublishBatteryResponse, error)
	// Publish to 'status text' updates.
	PublishStatusText(context.Context, *PublishStatusTextRequest) (*PublishStatusTextResponse, error)
	// Publish to 'odometry' updates.
	PublishOdometry(context.Context, *PublishOdometryRequest) (*PublishOdometryResponse, error)
	// Publish to 'position velocity' updates.
	PublishPositionVelocityNed(context.Context, *PublishPositionVelocityNedRequest) (*PublishPositionVelocityNedResponse, error)
	// Publish to 'ground truth' updates.
	PublishGroundTruth(context.Context, *PublishGroundTruthRequest) (*PublishGroundTruthResponse, error)
	// Publish to 'IMU' updates (in SI units in NED body frame).
	PublishImu(context.Context, *PublishImuRequest) (*PublishImuResponse, error)
	// Publish to 'Scaled IMU' updates.
	PublishScaledImu(context.Context, *PublishScaledImuRequest) (*PublishScaledImuResponse, error)
	// Publish to 'Raw IMU' updates.
	PublishRawImu(context.Context, *PublishRawImuRequest) (*PublishRawImuResponse, error)
	// Publish to 'unix epoch time' updates.
	PublishUnixEpochTime(context.Context, *PublishUnixEpochTimeRequest) (*PublishUnixEpochTimeResponse, error)
	// Publish to "distance sensor" updates.
	PublishDistanceSensor(context.Context, *PublishDistanceSensorRequest) (*PublishDistanceSensorResponse, error)
	mustEmbedUnimplementedTelemetryServerServiceServer()
}

// UnimplementedTelemetryServerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTelemetryServerServiceServer struct{}

func (UnimplementedTelemetryServerServiceServer) PublishPosition(context.Context, *PublishPositionRequest) (*PublishPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishPosition not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishHome(context.Context, *PublishHomeRequest) (*PublishHomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishHome not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishSysStatus(context.Context, *PublishSysStatusRequest) (*PublishSysStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishSysStatus not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishExtendedSysState(context.Context, *PublishExtendedSysStateRequest) (*PublishExtendedSysStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishExtendedSysState not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishRawGps(context.Context, *PublishRawGpsRequest) (*PublishRawGpsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishRawGps not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishBattery(context.Context, *PublishBatteryRequest) (*PublishBatteryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishBattery not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishStatusText(context.Context, *PublishStatusTextRequest) (*PublishStatusTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishStatusText not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishOdometry(context.Context, *PublishOdometryRequest) (*PublishOdometryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishOdometry not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishPositionVelocityNed(context.Context, *PublishPositionVelocityNedRequest) (*PublishPositionVelocityNedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishPositionVelocityNed not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishGroundTruth(context.Context, *PublishGroundTruthRequest) (*PublishGroundTruthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishGroundTruth not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishImu(context.Context, *PublishImuRequest) (*PublishImuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishImu not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishScaledImu(context.Context, *PublishScaledImuRequest) (*PublishScaledImuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishScaledImu not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishRawImu(context.Context, *PublishRawImuRequest) (*PublishRawImuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishRawImu not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishUnixEpochTime(context.Context, *PublishUnixEpochTimeRequest) (*PublishUnixEpochTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishUnixEpochTime not implemented")
}
func (UnimplementedTelemetryServerServiceServer) PublishDistanceSensor(context.Context, *PublishDistanceSensorRequest) (*PublishDistanceSensorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishDistanceSensor not implemented")
}
func (UnimplementedTelemetryServerServiceServer) mustEmbedUnimplementedTelemetryServerServiceServer() {
}
func (UnimplementedTelemetryServerServiceServer) testEmbeddedByValue() {}

// UnsafeTelemetryServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TelemetryServerServiceServer will
// result in compilation errors.
type UnsafeTelemetryServerServiceServer interface {
	mustEmbedUnimplementedTelemetryServerServiceServer()
}

func RegisterTelemetryServerServiceServer(s grpc.ServiceRegistrar, srv TelemetryServerServiceServer) {
	// If the following call pancis, it indicates UnimplementedTelemetryServerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TelemetryServerService_ServiceDesc, srv)
}

func _TelemetryServerService_PublishPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishPosition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishPosition(ctx, req.(*PublishPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishHome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishHome(ctx, req.(*PublishHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishSysStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishSysStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishSysStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishSysStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishSysStatus(ctx, req.(*PublishSysStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishExtendedSysState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishExtendedSysStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishExtendedSysState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishExtendedSysState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishExtendedSysState(ctx, req.(*PublishExtendedSysStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishRawGps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRawGpsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishRawGps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishRawGps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishRawGps(ctx, req.(*PublishRawGpsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishBattery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishBatteryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishBattery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishBattery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishBattery(ctx, req.(*PublishBatteryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishStatusText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishStatusTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishStatusText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishStatusText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishStatusText(ctx, req.(*PublishStatusTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishOdometry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishOdometryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishOdometry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishOdometry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishOdometry(ctx, req.(*PublishOdometryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishPositionVelocityNed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishPositionVelocityNedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishPositionVelocityNed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishPositionVelocityNed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishPositionVelocityNed(ctx, req.(*PublishPositionVelocityNedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishGroundTruth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishGroundTruthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishGroundTruth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishGroundTruth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishGroundTruth(ctx, req.(*PublishGroundTruthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishImu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishImuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishImu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishImu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishImu(ctx, req.(*PublishImuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishScaledImu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishScaledImuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishScaledImu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishScaledImu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishScaledImu(ctx, req.(*PublishScaledImuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishRawImu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRawImuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishRawImu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishRawImu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishRawImu(ctx, req.(*PublishRawImuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishUnixEpochTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishUnixEpochTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishUnixEpochTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishUnixEpochTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishUnixEpochTime(ctx, req.(*PublishUnixEpochTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelemetryServerService_PublishDistanceSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishDistanceSensorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServerServiceServer).PublishDistanceSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TelemetryServerService_PublishDistanceSensor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServerServiceServer).PublishDistanceSensor(ctx, req.(*PublishDistanceSensorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TelemetryServerService_ServiceDesc is the grpc.ServiceDesc for TelemetryServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TelemetryServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.telemetry_server.TelemetryServerService",
	HandlerType: (*TelemetryServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishPosition",
			Handler:    _TelemetryServerService_PublishPosition_Handler,
		},
		{
			MethodName: "PublishHome",
			Handler:    _TelemetryServerService_PublishHome_Handler,
		},
		{
			MethodName: "PublishSysStatus",
			Handler:    _TelemetryServerService_PublishSysStatus_Handler,
		},
		{
			MethodName: "PublishExtendedSysState",
			Handler:    _TelemetryServerService_PublishExtendedSysState_Handler,
		},
		{
			MethodName: "PublishRawGps",
			Handler:    _TelemetryServerService_PublishRawGps_Handler,
		},
		{
			MethodName: "PublishBattery",
			Handler:    _TelemetryServerService_PublishBattery_Handler,
		},
		{
			MethodName: "PublishStatusText",
			Handler:    _TelemetryServerService_PublishStatusText_Handler,
		},
		{
			MethodName: "PublishOdometry",
			Handler:    _TelemetryServerService_PublishOdometry_Handler,
		},
		{
			MethodName: "PublishPositionVelocityNed",
			Handler:    _TelemetryServerService_PublishPositionVelocityNed_Handler,
		},
		{
			MethodName: "PublishGroundTruth",
			Handler:    _TelemetryServerService_PublishGroundTruth_Handler,
		},
		{
			MethodName: "PublishImu",
			Handler:    _TelemetryServerService_PublishImu_Handler,
		},
		{
			MethodName: "PublishScaledImu",
			Handler:    _TelemetryServerService_PublishScaledImu_Handler,
		},
		{
			MethodName: "PublishRawImu",
			Handler:    _TelemetryServerService_PublishRawImu_Handler,
		},
		{
			MethodName: "PublishUnixEpochTime",
			Handler:    _TelemetryServerService_PublishUnixEpochTime_Handler,
		},
		{
			MethodName: "PublishDistanceSensor",
			Handler:    _TelemetryServerService_PublishDistanceSensor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "telemetry_server.proto",
}

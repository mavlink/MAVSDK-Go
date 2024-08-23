// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: action_server.proto

package action_server

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
	ActionServerService_SubscribeArmDisarm_FullMethodName        = "/mavsdk.rpc.action_server.ActionServerService/SubscribeArmDisarm"
	ActionServerService_SubscribeFlightModeChange_FullMethodName = "/mavsdk.rpc.action_server.ActionServerService/SubscribeFlightModeChange"
	ActionServerService_SubscribeTakeoff_FullMethodName          = "/mavsdk.rpc.action_server.ActionServerService/SubscribeTakeoff"
	ActionServerService_SubscribeLand_FullMethodName             = "/mavsdk.rpc.action_server.ActionServerService/SubscribeLand"
	ActionServerService_SubscribeReboot_FullMethodName           = "/mavsdk.rpc.action_server.ActionServerService/SubscribeReboot"
	ActionServerService_SubscribeShutdown_FullMethodName         = "/mavsdk.rpc.action_server.ActionServerService/SubscribeShutdown"
	ActionServerService_SubscribeTerminate_FullMethodName        = "/mavsdk.rpc.action_server.ActionServerService/SubscribeTerminate"
	ActionServerService_SetAllowTakeoff_FullMethodName           = "/mavsdk.rpc.action_server.ActionServerService/SetAllowTakeoff"
	ActionServerService_SetArmable_FullMethodName                = "/mavsdk.rpc.action_server.ActionServerService/SetArmable"
	ActionServerService_SetDisarmable_FullMethodName             = "/mavsdk.rpc.action_server.ActionServerService/SetDisarmable"
	ActionServerService_SetAllowableFlightModes_FullMethodName   = "/mavsdk.rpc.action_server.ActionServerService/SetAllowableFlightModes"
	ActionServerService_GetAllowableFlightModes_FullMethodName   = "/mavsdk.rpc.action_server.ActionServerService/GetAllowableFlightModes"
)

// ActionServerServiceClient is the client API for ActionServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Provide vehicle actions (as a server) such as arming, taking off, and landing.
type ActionServerServiceClient interface {
	// Subscribe to ARM/DISARM commands
	SubscribeArmDisarm(ctx context.Context, in *SubscribeArmDisarmRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ArmDisarmResponse], error)
	// Subscribe to DO_SET_MODE
	SubscribeFlightModeChange(ctx context.Context, in *SubscribeFlightModeChangeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FlightModeChangeResponse], error)
	// Subscribe to takeoff command
	SubscribeTakeoff(ctx context.Context, in *SubscribeTakeoffRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TakeoffResponse], error)
	// Subscribe to land command
	SubscribeLand(ctx context.Context, in *SubscribeLandRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[LandResponse], error)
	// Subscribe to reboot command
	SubscribeReboot(ctx context.Context, in *SubscribeRebootRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RebootResponse], error)
	// Subscribe to shutdown command
	SubscribeShutdown(ctx context.Context, in *SubscribeShutdownRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ShutdownResponse], error)
	// Subscribe to terminate command
	SubscribeTerminate(ctx context.Context, in *SubscribeTerminateRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TerminateResponse], error)
	// Can the vehicle takeoff
	SetAllowTakeoff(ctx context.Context, in *SetAllowTakeoffRequest, opts ...grpc.CallOption) (*SetAllowTakeoffResponse, error)
	// Can the vehicle arm when requested
	SetArmable(ctx context.Context, in *SetArmableRequest, opts ...grpc.CallOption) (*SetArmableResponse, error)
	// Can the vehicle disarm when requested
	SetDisarmable(ctx context.Context, in *SetDisarmableRequest, opts ...grpc.CallOption) (*SetDisarmableResponse, error)
	// Set which modes the vehicle can transition to (Manual always allowed)
	SetAllowableFlightModes(ctx context.Context, in *SetAllowableFlightModesRequest, opts ...grpc.CallOption) (*SetAllowableFlightModesResponse, error)
	// Get which modes the vehicle can transition to (Manual always allowed)
	GetAllowableFlightModes(ctx context.Context, in *GetAllowableFlightModesRequest, opts ...grpc.CallOption) (*GetAllowableFlightModesResponse, error)
}

type actionServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActionServerServiceClient(cc grpc.ClientConnInterface) ActionServerServiceClient {
	return &actionServerServiceClient{cc}
}

func (c *actionServerServiceClient) SubscribeArmDisarm(ctx context.Context, in *SubscribeArmDisarmRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ArmDisarmResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ActionServerService_ServiceDesc.Streams[0], ActionServerService_SubscribeArmDisarm_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeArmDisarmRequest, ArmDisarmResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeArmDisarmClient = grpc.ServerStreamingClient[ArmDisarmResponse]

func (c *actionServerServiceClient) SubscribeFlightModeChange(ctx context.Context, in *SubscribeFlightModeChangeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FlightModeChangeResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ActionServerService_ServiceDesc.Streams[1], ActionServerService_SubscribeFlightModeChange_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeFlightModeChangeRequest, FlightModeChangeResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeFlightModeChangeClient = grpc.ServerStreamingClient[FlightModeChangeResponse]

func (c *actionServerServiceClient) SubscribeTakeoff(ctx context.Context, in *SubscribeTakeoffRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TakeoffResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ActionServerService_ServiceDesc.Streams[2], ActionServerService_SubscribeTakeoff_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeTakeoffRequest, TakeoffResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeTakeoffClient = grpc.ServerStreamingClient[TakeoffResponse]

func (c *actionServerServiceClient) SubscribeLand(ctx context.Context, in *SubscribeLandRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[LandResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ActionServerService_ServiceDesc.Streams[3], ActionServerService_SubscribeLand_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeLandRequest, LandResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeLandClient = grpc.ServerStreamingClient[LandResponse]

func (c *actionServerServiceClient) SubscribeReboot(ctx context.Context, in *SubscribeRebootRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RebootResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ActionServerService_ServiceDesc.Streams[4], ActionServerService_SubscribeReboot_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeRebootRequest, RebootResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeRebootClient = grpc.ServerStreamingClient[RebootResponse]

func (c *actionServerServiceClient) SubscribeShutdown(ctx context.Context, in *SubscribeShutdownRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ShutdownResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ActionServerService_ServiceDesc.Streams[5], ActionServerService_SubscribeShutdown_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeShutdownRequest, ShutdownResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeShutdownClient = grpc.ServerStreamingClient[ShutdownResponse]

func (c *actionServerServiceClient) SubscribeTerminate(ctx context.Context, in *SubscribeTerminateRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[TerminateResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ActionServerService_ServiceDesc.Streams[6], ActionServerService_SubscribeTerminate_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeTerminateRequest, TerminateResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeTerminateClient = grpc.ServerStreamingClient[TerminateResponse]

func (c *actionServerServiceClient) SetAllowTakeoff(ctx context.Context, in *SetAllowTakeoffRequest, opts ...grpc.CallOption) (*SetAllowTakeoffResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetAllowTakeoffResponse)
	err := c.cc.Invoke(ctx, ActionServerService_SetAllowTakeoff_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServerServiceClient) SetArmable(ctx context.Context, in *SetArmableRequest, opts ...grpc.CallOption) (*SetArmableResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetArmableResponse)
	err := c.cc.Invoke(ctx, ActionServerService_SetArmable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServerServiceClient) SetDisarmable(ctx context.Context, in *SetDisarmableRequest, opts ...grpc.CallOption) (*SetDisarmableResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetDisarmableResponse)
	err := c.cc.Invoke(ctx, ActionServerService_SetDisarmable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServerServiceClient) SetAllowableFlightModes(ctx context.Context, in *SetAllowableFlightModesRequest, opts ...grpc.CallOption) (*SetAllowableFlightModesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetAllowableFlightModesResponse)
	err := c.cc.Invoke(ctx, ActionServerService_SetAllowableFlightModes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServerServiceClient) GetAllowableFlightModes(ctx context.Context, in *GetAllowableFlightModesRequest, opts ...grpc.CallOption) (*GetAllowableFlightModesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllowableFlightModesResponse)
	err := c.cc.Invoke(ctx, ActionServerService_GetAllowableFlightModes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionServerServiceServer is the server API for ActionServerService service.
// All implementations must embed UnimplementedActionServerServiceServer
// for forward compatibility.
//
// Provide vehicle actions (as a server) such as arming, taking off, and landing.
type ActionServerServiceServer interface {
	// Subscribe to ARM/DISARM commands
	SubscribeArmDisarm(*SubscribeArmDisarmRequest, grpc.ServerStreamingServer[ArmDisarmResponse]) error
	// Subscribe to DO_SET_MODE
	SubscribeFlightModeChange(*SubscribeFlightModeChangeRequest, grpc.ServerStreamingServer[FlightModeChangeResponse]) error
	// Subscribe to takeoff command
	SubscribeTakeoff(*SubscribeTakeoffRequest, grpc.ServerStreamingServer[TakeoffResponse]) error
	// Subscribe to land command
	SubscribeLand(*SubscribeLandRequest, grpc.ServerStreamingServer[LandResponse]) error
	// Subscribe to reboot command
	SubscribeReboot(*SubscribeRebootRequest, grpc.ServerStreamingServer[RebootResponse]) error
	// Subscribe to shutdown command
	SubscribeShutdown(*SubscribeShutdownRequest, grpc.ServerStreamingServer[ShutdownResponse]) error
	// Subscribe to terminate command
	SubscribeTerminate(*SubscribeTerminateRequest, grpc.ServerStreamingServer[TerminateResponse]) error
	// Can the vehicle takeoff
	SetAllowTakeoff(context.Context, *SetAllowTakeoffRequest) (*SetAllowTakeoffResponse, error)
	// Can the vehicle arm when requested
	SetArmable(context.Context, *SetArmableRequest) (*SetArmableResponse, error)
	// Can the vehicle disarm when requested
	SetDisarmable(context.Context, *SetDisarmableRequest) (*SetDisarmableResponse, error)
	// Set which modes the vehicle can transition to (Manual always allowed)
	SetAllowableFlightModes(context.Context, *SetAllowableFlightModesRequest) (*SetAllowableFlightModesResponse, error)
	// Get which modes the vehicle can transition to (Manual always allowed)
	GetAllowableFlightModes(context.Context, *GetAllowableFlightModesRequest) (*GetAllowableFlightModesResponse, error)
	mustEmbedUnimplementedActionServerServiceServer()
}

// UnimplementedActionServerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedActionServerServiceServer struct{}

func (UnimplementedActionServerServiceServer) SubscribeArmDisarm(*SubscribeArmDisarmRequest, grpc.ServerStreamingServer[ArmDisarmResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeArmDisarm not implemented")
}
func (UnimplementedActionServerServiceServer) SubscribeFlightModeChange(*SubscribeFlightModeChangeRequest, grpc.ServerStreamingServer[FlightModeChangeResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeFlightModeChange not implemented")
}
func (UnimplementedActionServerServiceServer) SubscribeTakeoff(*SubscribeTakeoffRequest, grpc.ServerStreamingServer[TakeoffResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeTakeoff not implemented")
}
func (UnimplementedActionServerServiceServer) SubscribeLand(*SubscribeLandRequest, grpc.ServerStreamingServer[LandResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeLand not implemented")
}
func (UnimplementedActionServerServiceServer) SubscribeReboot(*SubscribeRebootRequest, grpc.ServerStreamingServer[RebootResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeReboot not implemented")
}
func (UnimplementedActionServerServiceServer) SubscribeShutdown(*SubscribeShutdownRequest, grpc.ServerStreamingServer[ShutdownResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeShutdown not implemented")
}
func (UnimplementedActionServerServiceServer) SubscribeTerminate(*SubscribeTerminateRequest, grpc.ServerStreamingServer[TerminateResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeTerminate not implemented")
}
func (UnimplementedActionServerServiceServer) SetAllowTakeoff(context.Context, *SetAllowTakeoffRequest) (*SetAllowTakeoffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAllowTakeoff not implemented")
}
func (UnimplementedActionServerServiceServer) SetArmable(context.Context, *SetArmableRequest) (*SetArmableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetArmable not implemented")
}
func (UnimplementedActionServerServiceServer) SetDisarmable(context.Context, *SetDisarmableRequest) (*SetDisarmableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDisarmable not implemented")
}
func (UnimplementedActionServerServiceServer) SetAllowableFlightModes(context.Context, *SetAllowableFlightModesRequest) (*SetAllowableFlightModesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAllowableFlightModes not implemented")
}
func (UnimplementedActionServerServiceServer) GetAllowableFlightModes(context.Context, *GetAllowableFlightModesRequest) (*GetAllowableFlightModesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllowableFlightModes not implemented")
}
func (UnimplementedActionServerServiceServer) mustEmbedUnimplementedActionServerServiceServer() {}
func (UnimplementedActionServerServiceServer) testEmbeddedByValue()                             {}

// UnsafeActionServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionServerServiceServer will
// result in compilation errors.
type UnsafeActionServerServiceServer interface {
	mustEmbedUnimplementedActionServerServiceServer()
}

func RegisterActionServerServiceServer(s grpc.ServiceRegistrar, srv ActionServerServiceServer) {
	// If the following call pancis, it indicates UnimplementedActionServerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ActionServerService_ServiceDesc, srv)
}

func _ActionServerService_SubscribeArmDisarm_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeArmDisarmRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActionServerServiceServer).SubscribeArmDisarm(m, &grpc.GenericServerStream[SubscribeArmDisarmRequest, ArmDisarmResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeArmDisarmServer = grpc.ServerStreamingServer[ArmDisarmResponse]

func _ActionServerService_SubscribeFlightModeChange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeFlightModeChangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActionServerServiceServer).SubscribeFlightModeChange(m, &grpc.GenericServerStream[SubscribeFlightModeChangeRequest, FlightModeChangeResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeFlightModeChangeServer = grpc.ServerStreamingServer[FlightModeChangeResponse]

func _ActionServerService_SubscribeTakeoff_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeTakeoffRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActionServerServiceServer).SubscribeTakeoff(m, &grpc.GenericServerStream[SubscribeTakeoffRequest, TakeoffResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeTakeoffServer = grpc.ServerStreamingServer[TakeoffResponse]

func _ActionServerService_SubscribeLand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeLandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActionServerServiceServer).SubscribeLand(m, &grpc.GenericServerStream[SubscribeLandRequest, LandResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeLandServer = grpc.ServerStreamingServer[LandResponse]

func _ActionServerService_SubscribeReboot_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRebootRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActionServerServiceServer).SubscribeReboot(m, &grpc.GenericServerStream[SubscribeRebootRequest, RebootResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeRebootServer = grpc.ServerStreamingServer[RebootResponse]

func _ActionServerService_SubscribeShutdown_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeShutdownRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActionServerServiceServer).SubscribeShutdown(m, &grpc.GenericServerStream[SubscribeShutdownRequest, ShutdownResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeShutdownServer = grpc.ServerStreamingServer[ShutdownResponse]

func _ActionServerService_SubscribeTerminate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeTerminateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ActionServerServiceServer).SubscribeTerminate(m, &grpc.GenericServerStream[SubscribeTerminateRequest, TerminateResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ActionServerService_SubscribeTerminateServer = grpc.ServerStreamingServer[TerminateResponse]

func _ActionServerService_SetAllowTakeoff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAllowTakeoffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServerServiceServer).SetAllowTakeoff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionServerService_SetAllowTakeoff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServerServiceServer).SetAllowTakeoff(ctx, req.(*SetAllowTakeoffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionServerService_SetArmable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetArmableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServerServiceServer).SetArmable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionServerService_SetArmable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServerServiceServer).SetArmable(ctx, req.(*SetArmableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionServerService_SetDisarmable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDisarmableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServerServiceServer).SetDisarmable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionServerService_SetDisarmable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServerServiceServer).SetDisarmable(ctx, req.(*SetDisarmableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionServerService_SetAllowableFlightModes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAllowableFlightModesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServerServiceServer).SetAllowableFlightModes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionServerService_SetAllowableFlightModes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServerServiceServer).SetAllowableFlightModes(ctx, req.(*SetAllowableFlightModesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionServerService_GetAllowableFlightModes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllowableFlightModesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServerServiceServer).GetAllowableFlightModes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActionServerService_GetAllowableFlightModes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServerServiceServer).GetAllowableFlightModes(ctx, req.(*GetAllowableFlightModesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActionServerService_ServiceDesc is the grpc.ServiceDesc for ActionServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActionServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.action_server.ActionServerService",
	HandlerType: (*ActionServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetAllowTakeoff",
			Handler:    _ActionServerService_SetAllowTakeoff_Handler,
		},
		{
			MethodName: "SetArmable",
			Handler:    _ActionServerService_SetArmable_Handler,
		},
		{
			MethodName: "SetDisarmable",
			Handler:    _ActionServerService_SetDisarmable_Handler,
		},
		{
			MethodName: "SetAllowableFlightModes",
			Handler:    _ActionServerService_SetAllowableFlightModes_Handler,
		},
		{
			MethodName: "GetAllowableFlightModes",
			Handler:    _ActionServerService_GetAllowableFlightModes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeArmDisarm",
			Handler:       _ActionServerService_SubscribeArmDisarm_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeFlightModeChange",
			Handler:       _ActionServerService_SubscribeFlightModeChange_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeTakeoff",
			Handler:       _ActionServerService_SubscribeTakeoff_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeLand",
			Handler:       _ActionServerService_SubscribeLand_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeReboot",
			Handler:       _ActionServerService_SubscribeReboot_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeShutdown",
			Handler:       _ActionServerService_SubscribeShutdown_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeTerminate",
			Handler:       _ActionServerService_SubscribeTerminate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "action_server.proto",
}

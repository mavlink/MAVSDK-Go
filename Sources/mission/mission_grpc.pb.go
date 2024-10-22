// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: mission.proto

package mission

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
	MissionService_UploadMission_FullMethodName                        = "/mavsdk.rpc.mission.MissionService/UploadMission"
	MissionService_SubscribeUploadMissionWithProgress_FullMethodName   = "/mavsdk.rpc.mission.MissionService/SubscribeUploadMissionWithProgress"
	MissionService_CancelMissionUpload_FullMethodName                  = "/mavsdk.rpc.mission.MissionService/CancelMissionUpload"
	MissionService_DownloadMission_FullMethodName                      = "/mavsdk.rpc.mission.MissionService/DownloadMission"
	MissionService_SubscribeDownloadMissionWithProgress_FullMethodName = "/mavsdk.rpc.mission.MissionService/SubscribeDownloadMissionWithProgress"
	MissionService_CancelMissionDownload_FullMethodName                = "/mavsdk.rpc.mission.MissionService/CancelMissionDownload"
	MissionService_StartMission_FullMethodName                         = "/mavsdk.rpc.mission.MissionService/StartMission"
	MissionService_PauseMission_FullMethodName                         = "/mavsdk.rpc.mission.MissionService/PauseMission"
	MissionService_ClearMission_FullMethodName                         = "/mavsdk.rpc.mission.MissionService/ClearMission"
	MissionService_SetCurrentMissionItem_FullMethodName                = "/mavsdk.rpc.mission.MissionService/SetCurrentMissionItem"
	MissionService_IsMissionFinished_FullMethodName                    = "/mavsdk.rpc.mission.MissionService/IsMissionFinished"
	MissionService_SubscribeMissionProgress_FullMethodName             = "/mavsdk.rpc.mission.MissionService/SubscribeMissionProgress"
	MissionService_GetReturnToLaunchAfterMission_FullMethodName        = "/mavsdk.rpc.mission.MissionService/GetReturnToLaunchAfterMission"
	MissionService_SetReturnToLaunchAfterMission_FullMethodName        = "/mavsdk.rpc.mission.MissionService/SetReturnToLaunchAfterMission"
)

// MissionServiceClient is the client API for MissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Enable waypoint missions.
type MissionServiceClient interface {
	// Upload a list of mission items to the system.
	//
	// The mission items are uploaded to a drone. Once uploaded the mission can be started and
	// executed even if the connection is lost.
	UploadMission(ctx context.Context, in *UploadMissionRequest, opts ...grpc.CallOption) (*UploadMissionResponse, error)
	// Upload a list of mission items to the system and report upload progress.
	//
	// The mission items are uploaded to a drone. Once uploaded the mission can be started and
	// executed even if the connection is lost.
	SubscribeUploadMissionWithProgress(ctx context.Context, in *SubscribeUploadMissionWithProgressRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UploadMissionWithProgressResponse], error)
	// Cancel an ongoing mission upload.
	CancelMissionUpload(ctx context.Context, in *CancelMissionUploadRequest, opts ...grpc.CallOption) (*CancelMissionUploadResponse, error)
	// Download a list of mission items from the system (asynchronous).
	//
	// Will fail if any of the downloaded mission items are not supported
	// by the MAVSDK API.
	DownloadMission(ctx context.Context, in *DownloadMissionRequest, opts ...grpc.CallOption) (*DownloadMissionResponse, error)
	// Download a list of mission items from the system (asynchronous) and report progress.
	//
	// Will fail if any of the downloaded mission items are not supported
	// by the MAVSDK API.
	SubscribeDownloadMissionWithProgress(ctx context.Context, in *SubscribeDownloadMissionWithProgressRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadMissionWithProgressResponse], error)
	// Cancel an ongoing mission download.
	CancelMissionDownload(ctx context.Context, in *CancelMissionDownloadRequest, opts ...grpc.CallOption) (*CancelMissionDownloadResponse, error)
	// Start the mission.
	//
	// A mission must be uploaded to the vehicle before this can be called.
	StartMission(ctx context.Context, in *StartMissionRequest, opts ...grpc.CallOption) (*StartMissionResponse, error)
	// Pause the mission.
	//
	// Pausing the mission puts the vehicle into
	// [HOLD mode](https://docs.px4.io/en/flight_modes/hold.html).
	// A multicopter should just hover at the spot while a fixedwing vehicle should loiter
	// around the location where it paused.
	PauseMission(ctx context.Context, in *PauseMissionRequest, opts ...grpc.CallOption) (*PauseMissionResponse, error)
	// Clear the mission saved on the vehicle.
	ClearMission(ctx context.Context, in *ClearMissionRequest, opts ...grpc.CallOption) (*ClearMissionResponse, error)
	// Sets the mission item index to go to.
	//
	// By setting the current index to 0, the mission is restarted from the beginning. If it is set
	// to a specific index of a mission item, the mission will be set to this item.
	//
	// Note that this is not necessarily true for general missions using MAVLink if loop counters
	// are used.
	SetCurrentMissionItem(ctx context.Context, in *SetCurrentMissionItemRequest, opts ...grpc.CallOption) (*SetCurrentMissionItemResponse, error)
	// Check if the mission has been finished.
	IsMissionFinished(ctx context.Context, in *IsMissionFinishedRequest, opts ...grpc.CallOption) (*IsMissionFinishedResponse, error)
	// Subscribe to mission progress updates.
	SubscribeMissionProgress(ctx context.Context, in *SubscribeMissionProgressRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MissionProgressResponse], error)
	// Get whether to trigger Return-to-Launch (RTL) after mission is complete.
	//
	// Before getting this option, it needs to be set, or a mission
	// needs to be downloaded.
	GetReturnToLaunchAfterMission(ctx context.Context, in *GetReturnToLaunchAfterMissionRequest, opts ...grpc.CallOption) (*GetReturnToLaunchAfterMissionResponse, error)
	// Set whether to trigger Return-to-Launch (RTL) after the mission is complete.
	//
	// This will only take effect for the next mission upload, meaning that
	// the mission may have to be uploaded again.
	SetReturnToLaunchAfterMission(ctx context.Context, in *SetReturnToLaunchAfterMissionRequest, opts ...grpc.CallOption) (*SetReturnToLaunchAfterMissionResponse, error)
}

type missionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMissionServiceClient(cc grpc.ClientConnInterface) MissionServiceClient {
	return &missionServiceClient{cc}
}

func (c *missionServiceClient) UploadMission(ctx context.Context, in *UploadMissionRequest, opts ...grpc.CallOption) (*UploadMissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadMissionResponse)
	err := c.cc.Invoke(ctx, MissionService_UploadMission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) SubscribeUploadMissionWithProgress(ctx context.Context, in *SubscribeUploadMissionWithProgressRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[UploadMissionWithProgressResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MissionService_ServiceDesc.Streams[0], MissionService_SubscribeUploadMissionWithProgress_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeUploadMissionWithProgressRequest, UploadMissionWithProgressResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MissionService_SubscribeUploadMissionWithProgressClient = grpc.ServerStreamingClient[UploadMissionWithProgressResponse]

func (c *missionServiceClient) CancelMissionUpload(ctx context.Context, in *CancelMissionUploadRequest, opts ...grpc.CallOption) (*CancelMissionUploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelMissionUploadResponse)
	err := c.cc.Invoke(ctx, MissionService_CancelMissionUpload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) DownloadMission(ctx context.Context, in *DownloadMissionRequest, opts ...grpc.CallOption) (*DownloadMissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DownloadMissionResponse)
	err := c.cc.Invoke(ctx, MissionService_DownloadMission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) SubscribeDownloadMissionWithProgress(ctx context.Context, in *SubscribeDownloadMissionWithProgressRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DownloadMissionWithProgressResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MissionService_ServiceDesc.Streams[1], MissionService_SubscribeDownloadMissionWithProgress_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeDownloadMissionWithProgressRequest, DownloadMissionWithProgressResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MissionService_SubscribeDownloadMissionWithProgressClient = grpc.ServerStreamingClient[DownloadMissionWithProgressResponse]

func (c *missionServiceClient) CancelMissionDownload(ctx context.Context, in *CancelMissionDownloadRequest, opts ...grpc.CallOption) (*CancelMissionDownloadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelMissionDownloadResponse)
	err := c.cc.Invoke(ctx, MissionService_CancelMissionDownload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) StartMission(ctx context.Context, in *StartMissionRequest, opts ...grpc.CallOption) (*StartMissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartMissionResponse)
	err := c.cc.Invoke(ctx, MissionService_StartMission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) PauseMission(ctx context.Context, in *PauseMissionRequest, opts ...grpc.CallOption) (*PauseMissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PauseMissionResponse)
	err := c.cc.Invoke(ctx, MissionService_PauseMission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) ClearMission(ctx context.Context, in *ClearMissionRequest, opts ...grpc.CallOption) (*ClearMissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClearMissionResponse)
	err := c.cc.Invoke(ctx, MissionService_ClearMission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) SetCurrentMissionItem(ctx context.Context, in *SetCurrentMissionItemRequest, opts ...grpc.CallOption) (*SetCurrentMissionItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetCurrentMissionItemResponse)
	err := c.cc.Invoke(ctx, MissionService_SetCurrentMissionItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) IsMissionFinished(ctx context.Context, in *IsMissionFinishedRequest, opts ...grpc.CallOption) (*IsMissionFinishedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsMissionFinishedResponse)
	err := c.cc.Invoke(ctx, MissionService_IsMissionFinished_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) SubscribeMissionProgress(ctx context.Context, in *SubscribeMissionProgressRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MissionProgressResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MissionService_ServiceDesc.Streams[2], MissionService_SubscribeMissionProgress_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeMissionProgressRequest, MissionProgressResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MissionService_SubscribeMissionProgressClient = grpc.ServerStreamingClient[MissionProgressResponse]

func (c *missionServiceClient) GetReturnToLaunchAfterMission(ctx context.Context, in *GetReturnToLaunchAfterMissionRequest, opts ...grpc.CallOption) (*GetReturnToLaunchAfterMissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReturnToLaunchAfterMissionResponse)
	err := c.cc.Invoke(ctx, MissionService_GetReturnToLaunchAfterMission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) SetReturnToLaunchAfterMission(ctx context.Context, in *SetReturnToLaunchAfterMissionRequest, opts ...grpc.CallOption) (*SetReturnToLaunchAfterMissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetReturnToLaunchAfterMissionResponse)
	err := c.cc.Invoke(ctx, MissionService_SetReturnToLaunchAfterMission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MissionServiceServer is the server API for MissionService service.
// All implementations must embed UnimplementedMissionServiceServer
// for forward compatibility.
//
// Enable waypoint missions.
type MissionServiceServer interface {
	// Upload a list of mission items to the system.
	//
	// The mission items are uploaded to a drone. Once uploaded the mission can be started and
	// executed even if the connection is lost.
	UploadMission(context.Context, *UploadMissionRequest) (*UploadMissionResponse, error)
	// Upload a list of mission items to the system and report upload progress.
	//
	// The mission items are uploaded to a drone. Once uploaded the mission can be started and
	// executed even if the connection is lost.
	SubscribeUploadMissionWithProgress(*SubscribeUploadMissionWithProgressRequest, grpc.ServerStreamingServer[UploadMissionWithProgressResponse]) error
	// Cancel an ongoing mission upload.
	CancelMissionUpload(context.Context, *CancelMissionUploadRequest) (*CancelMissionUploadResponse, error)
	// Download a list of mission items from the system (asynchronous).
	//
	// Will fail if any of the downloaded mission items are not supported
	// by the MAVSDK API.
	DownloadMission(context.Context, *DownloadMissionRequest) (*DownloadMissionResponse, error)
	// Download a list of mission items from the system (asynchronous) and report progress.
	//
	// Will fail if any of the downloaded mission items are not supported
	// by the MAVSDK API.
	SubscribeDownloadMissionWithProgress(*SubscribeDownloadMissionWithProgressRequest, grpc.ServerStreamingServer[DownloadMissionWithProgressResponse]) error
	// Cancel an ongoing mission download.
	CancelMissionDownload(context.Context, *CancelMissionDownloadRequest) (*CancelMissionDownloadResponse, error)
	// Start the mission.
	//
	// A mission must be uploaded to the vehicle before this can be called.
	StartMission(context.Context, *StartMissionRequest) (*StartMissionResponse, error)
	// Pause the mission.
	//
	// Pausing the mission puts the vehicle into
	// [HOLD mode](https://docs.px4.io/en/flight_modes/hold.html).
	// A multicopter should just hover at the spot while a fixedwing vehicle should loiter
	// around the location where it paused.
	PauseMission(context.Context, *PauseMissionRequest) (*PauseMissionResponse, error)
	// Clear the mission saved on the vehicle.
	ClearMission(context.Context, *ClearMissionRequest) (*ClearMissionResponse, error)
	// Sets the mission item index to go to.
	//
	// By setting the current index to 0, the mission is restarted from the beginning. If it is set
	// to a specific index of a mission item, the mission will be set to this item.
	//
	// Note that this is not necessarily true for general missions using MAVLink if loop counters
	// are used.
	SetCurrentMissionItem(context.Context, *SetCurrentMissionItemRequest) (*SetCurrentMissionItemResponse, error)
	// Check if the mission has been finished.
	IsMissionFinished(context.Context, *IsMissionFinishedRequest) (*IsMissionFinishedResponse, error)
	// Subscribe to mission progress updates.
	SubscribeMissionProgress(*SubscribeMissionProgressRequest, grpc.ServerStreamingServer[MissionProgressResponse]) error
	// Get whether to trigger Return-to-Launch (RTL) after mission is complete.
	//
	// Before getting this option, it needs to be set, or a mission
	// needs to be downloaded.
	GetReturnToLaunchAfterMission(context.Context, *GetReturnToLaunchAfterMissionRequest) (*GetReturnToLaunchAfterMissionResponse, error)
	// Set whether to trigger Return-to-Launch (RTL) after the mission is complete.
	//
	// This will only take effect for the next mission upload, meaning that
	// the mission may have to be uploaded again.
	SetReturnToLaunchAfterMission(context.Context, *SetReturnToLaunchAfterMissionRequest) (*SetReturnToLaunchAfterMissionResponse, error)
	mustEmbedUnimplementedMissionServiceServer()
}

// UnimplementedMissionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMissionServiceServer struct{}

func (UnimplementedMissionServiceServer) UploadMission(context.Context, *UploadMissionRequest) (*UploadMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadMission not implemented")
}
func (UnimplementedMissionServiceServer) SubscribeUploadMissionWithProgress(*SubscribeUploadMissionWithProgressRequest, grpc.ServerStreamingServer[UploadMissionWithProgressResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeUploadMissionWithProgress not implemented")
}
func (UnimplementedMissionServiceServer) CancelMissionUpload(context.Context, *CancelMissionUploadRequest) (*CancelMissionUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelMissionUpload not implemented")
}
func (UnimplementedMissionServiceServer) DownloadMission(context.Context, *DownloadMissionRequest) (*DownloadMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadMission not implemented")
}
func (UnimplementedMissionServiceServer) SubscribeDownloadMissionWithProgress(*SubscribeDownloadMissionWithProgressRequest, grpc.ServerStreamingServer[DownloadMissionWithProgressResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeDownloadMissionWithProgress not implemented")
}
func (UnimplementedMissionServiceServer) CancelMissionDownload(context.Context, *CancelMissionDownloadRequest) (*CancelMissionDownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelMissionDownload not implemented")
}
func (UnimplementedMissionServiceServer) StartMission(context.Context, *StartMissionRequest) (*StartMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMission not implemented")
}
func (UnimplementedMissionServiceServer) PauseMission(context.Context, *PauseMissionRequest) (*PauseMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseMission not implemented")
}
func (UnimplementedMissionServiceServer) ClearMission(context.Context, *ClearMissionRequest) (*ClearMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearMission not implemented")
}
func (UnimplementedMissionServiceServer) SetCurrentMissionItem(context.Context, *SetCurrentMissionItemRequest) (*SetCurrentMissionItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCurrentMissionItem not implemented")
}
func (UnimplementedMissionServiceServer) IsMissionFinished(context.Context, *IsMissionFinishedRequest) (*IsMissionFinishedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsMissionFinished not implemented")
}
func (UnimplementedMissionServiceServer) SubscribeMissionProgress(*SubscribeMissionProgressRequest, grpc.ServerStreamingServer[MissionProgressResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeMissionProgress not implemented")
}
func (UnimplementedMissionServiceServer) GetReturnToLaunchAfterMission(context.Context, *GetReturnToLaunchAfterMissionRequest) (*GetReturnToLaunchAfterMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReturnToLaunchAfterMission not implemented")
}
func (UnimplementedMissionServiceServer) SetReturnToLaunchAfterMission(context.Context, *SetReturnToLaunchAfterMissionRequest) (*SetReturnToLaunchAfterMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetReturnToLaunchAfterMission not implemented")
}
func (UnimplementedMissionServiceServer) mustEmbedUnimplementedMissionServiceServer() {}
func (UnimplementedMissionServiceServer) testEmbeddedByValue()                        {}

// UnsafeMissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MissionServiceServer will
// result in compilation errors.
type UnsafeMissionServiceServer interface {
	mustEmbedUnimplementedMissionServiceServer()
}

func RegisterMissionServiceServer(s grpc.ServiceRegistrar, srv MissionServiceServer) {
	// If the following call pancis, it indicates UnimplementedMissionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MissionService_ServiceDesc, srv)
}

func _MissionService_UploadMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).UploadMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_UploadMission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).UploadMission(ctx, req.(*UploadMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_SubscribeUploadMissionWithProgress_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeUploadMissionWithProgressRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MissionServiceServer).SubscribeUploadMissionWithProgress(m, &grpc.GenericServerStream[SubscribeUploadMissionWithProgressRequest, UploadMissionWithProgressResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MissionService_SubscribeUploadMissionWithProgressServer = grpc.ServerStreamingServer[UploadMissionWithProgressResponse]

func _MissionService_CancelMissionUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelMissionUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).CancelMissionUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_CancelMissionUpload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).CancelMissionUpload(ctx, req.(*CancelMissionUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_DownloadMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).DownloadMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_DownloadMission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).DownloadMission(ctx, req.(*DownloadMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_SubscribeDownloadMissionWithProgress_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeDownloadMissionWithProgressRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MissionServiceServer).SubscribeDownloadMissionWithProgress(m, &grpc.GenericServerStream[SubscribeDownloadMissionWithProgressRequest, DownloadMissionWithProgressResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MissionService_SubscribeDownloadMissionWithProgressServer = grpc.ServerStreamingServer[DownloadMissionWithProgressResponse]

func _MissionService_CancelMissionDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelMissionDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).CancelMissionDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_CancelMissionDownload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).CancelMissionDownload(ctx, req.(*CancelMissionDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_StartMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).StartMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_StartMission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).StartMission(ctx, req.(*StartMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_PauseMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).PauseMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_PauseMission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).PauseMission(ctx, req.(*PauseMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_ClearMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).ClearMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_ClearMission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).ClearMission(ctx, req.(*ClearMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_SetCurrentMissionItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCurrentMissionItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).SetCurrentMissionItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_SetCurrentMissionItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).SetCurrentMissionItem(ctx, req.(*SetCurrentMissionItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_IsMissionFinished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsMissionFinishedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).IsMissionFinished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_IsMissionFinished_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).IsMissionFinished(ctx, req.(*IsMissionFinishedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_SubscribeMissionProgress_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeMissionProgressRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MissionServiceServer).SubscribeMissionProgress(m, &grpc.GenericServerStream[SubscribeMissionProgressRequest, MissionProgressResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MissionService_SubscribeMissionProgressServer = grpc.ServerStreamingServer[MissionProgressResponse]

func _MissionService_GetReturnToLaunchAfterMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReturnToLaunchAfterMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).GetReturnToLaunchAfterMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_GetReturnToLaunchAfterMission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).GetReturnToLaunchAfterMission(ctx, req.(*GetReturnToLaunchAfterMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_SetReturnToLaunchAfterMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReturnToLaunchAfterMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).SetReturnToLaunchAfterMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_SetReturnToLaunchAfterMission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).SetReturnToLaunchAfterMission(ctx, req.(*SetReturnToLaunchAfterMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MissionService_ServiceDesc is the grpc.ServiceDesc for MissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.mission.MissionService",
	HandlerType: (*MissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadMission",
			Handler:    _MissionService_UploadMission_Handler,
		},
		{
			MethodName: "CancelMissionUpload",
			Handler:    _MissionService_CancelMissionUpload_Handler,
		},
		{
			MethodName: "DownloadMission",
			Handler:    _MissionService_DownloadMission_Handler,
		},
		{
			MethodName: "CancelMissionDownload",
			Handler:    _MissionService_CancelMissionDownload_Handler,
		},
		{
			MethodName: "StartMission",
			Handler:    _MissionService_StartMission_Handler,
		},
		{
			MethodName: "PauseMission",
			Handler:    _MissionService_PauseMission_Handler,
		},
		{
			MethodName: "ClearMission",
			Handler:    _MissionService_ClearMission_Handler,
		},
		{
			MethodName: "SetCurrentMissionItem",
			Handler:    _MissionService_SetCurrentMissionItem_Handler,
		},
		{
			MethodName: "IsMissionFinished",
			Handler:    _MissionService_IsMissionFinished_Handler,
		},
		{
			MethodName: "GetReturnToLaunchAfterMission",
			Handler:    _MissionService_GetReturnToLaunchAfterMission_Handler,
		},
		{
			MethodName: "SetReturnToLaunchAfterMission",
			Handler:    _MissionService_SetReturnToLaunchAfterMission_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeUploadMissionWithProgress",
			Handler:       _MissionService_SubscribeUploadMissionWithProgress_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeDownloadMissionWithProgress",
			Handler:       _MissionService_SubscribeDownloadMissionWithProgress_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeMissionProgress",
			Handler:       _MissionService_SubscribeMissionProgress_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mission.proto",
}

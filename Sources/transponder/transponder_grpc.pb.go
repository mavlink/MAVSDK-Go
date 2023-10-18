// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: transponder.proto

package transponder

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TransponderService_SubscribeTransponder_FullMethodName = "/mavsdk.rpc.transponder.TransponderService/SubscribeTransponder"
	TransponderService_SetRateTransponder_FullMethodName   = "/mavsdk.rpc.transponder.TransponderService/SetRateTransponder"
)

// TransponderServiceClient is the client API for TransponderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransponderServiceClient interface {
	// Subscribe to 'transponder' updates.
	SubscribeTransponder(ctx context.Context, in *SubscribeTransponderRequest, opts ...grpc.CallOption) (TransponderService_SubscribeTransponderClient, error)
	// Set rate to 'transponder' updates.
	SetRateTransponder(ctx context.Context, in *SetRateTransponderRequest, opts ...grpc.CallOption) (*SetRateTransponderResponse, error)
}

type transponderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransponderServiceClient(cc grpc.ClientConnInterface) TransponderServiceClient {
	return &transponderServiceClient{cc}
}

func (c *transponderServiceClient) SubscribeTransponder(ctx context.Context, in *SubscribeTransponderRequest, opts ...grpc.CallOption) (TransponderService_SubscribeTransponderClient, error) {
	stream, err := c.cc.NewStream(ctx, &TransponderService_ServiceDesc.Streams[0], TransponderService_SubscribeTransponder_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transponderServiceSubscribeTransponderClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TransponderService_SubscribeTransponderClient interface {
	Recv() (*TransponderResponse, error)
	grpc.ClientStream
}

type transponderServiceSubscribeTransponderClient struct {
	grpc.ClientStream
}

func (x *transponderServiceSubscribeTransponderClient) Recv() (*TransponderResponse, error) {
	m := new(TransponderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transponderServiceClient) SetRateTransponder(ctx context.Context, in *SetRateTransponderRequest, opts ...grpc.CallOption) (*SetRateTransponderResponse, error) {
	out := new(SetRateTransponderResponse)
	err := c.cc.Invoke(ctx, TransponderService_SetRateTransponder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransponderServiceServer is the server API for TransponderService service.
// All implementations must embed UnimplementedTransponderServiceServer
// for forward compatibility
type TransponderServiceServer interface {
	// Subscribe to 'transponder' updates.
	SubscribeTransponder(*SubscribeTransponderRequest, TransponderService_SubscribeTransponderServer) error
	// Set rate to 'transponder' updates.
	SetRateTransponder(context.Context, *SetRateTransponderRequest) (*SetRateTransponderResponse, error)
	mustEmbedUnimplementedTransponderServiceServer()
}

// UnimplementedTransponderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransponderServiceServer struct {
}

func (UnimplementedTransponderServiceServer) SubscribeTransponder(*SubscribeTransponderRequest, TransponderService_SubscribeTransponderServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeTransponder not implemented")
}
func (UnimplementedTransponderServiceServer) SetRateTransponder(context.Context, *SetRateTransponderRequest) (*SetRateTransponderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRateTransponder not implemented")
}
func (UnimplementedTransponderServiceServer) mustEmbedUnimplementedTransponderServiceServer() {}

// UnsafeTransponderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransponderServiceServer will
// result in compilation errors.
type UnsafeTransponderServiceServer interface {
	mustEmbedUnimplementedTransponderServiceServer()
}

func RegisterTransponderServiceServer(s grpc.ServiceRegistrar, srv TransponderServiceServer) {
	s.RegisterService(&TransponderService_ServiceDesc, srv)
}

func _TransponderService_SubscribeTransponder_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeTransponderRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransponderServiceServer).SubscribeTransponder(m, &transponderServiceSubscribeTransponderServer{stream})
}

type TransponderService_SubscribeTransponderServer interface {
	Send(*TransponderResponse) error
	grpc.ServerStream
}

type transponderServiceSubscribeTransponderServer struct {
	grpc.ServerStream
}

func (x *transponderServiceSubscribeTransponderServer) Send(m *TransponderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TransponderService_SetRateTransponder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRateTransponderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransponderServiceServer).SetRateTransponder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransponderService_SetRateTransponder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransponderServiceServer).SetRateTransponder(ctx, req.(*SetRateTransponderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransponderService_ServiceDesc is the grpc.ServiceDesc for TransponderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransponderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.transponder.TransponderService",
	HandlerType: (*TransponderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetRateTransponder",
			Handler:    _TransponderService_SetRateTransponder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeTransponder",
			Handler:       _TransponderService_SubscribeTransponder_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transponder.proto",
}

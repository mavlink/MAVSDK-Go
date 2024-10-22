// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: failure.proto

package failure

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
	FailureService_Inject_FullMethodName = "/mavsdk.rpc.failure.FailureService/Inject"
)

// FailureServiceClient is the client API for FailureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Inject failures into system to test failsafes.
type FailureServiceClient interface {
	// Injects a failure.
	Inject(ctx context.Context, in *InjectRequest, opts ...grpc.CallOption) (*InjectResponse, error)
}

type failureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFailureServiceClient(cc grpc.ClientConnInterface) FailureServiceClient {
	return &failureServiceClient{cc}
}

func (c *failureServiceClient) Inject(ctx context.Context, in *InjectRequest, opts ...grpc.CallOption) (*InjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InjectResponse)
	err := c.cc.Invoke(ctx, FailureService_Inject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FailureServiceServer is the server API for FailureService service.
// All implementations must embed UnimplementedFailureServiceServer
// for forward compatibility.
//
// Inject failures into system to test failsafes.
type FailureServiceServer interface {
	// Injects a failure.
	Inject(context.Context, *InjectRequest) (*InjectResponse, error)
	mustEmbedUnimplementedFailureServiceServer()
}

// UnimplementedFailureServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFailureServiceServer struct{}

func (UnimplementedFailureServiceServer) Inject(context.Context, *InjectRequest) (*InjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inject not implemented")
}
func (UnimplementedFailureServiceServer) mustEmbedUnimplementedFailureServiceServer() {}
func (UnimplementedFailureServiceServer) testEmbeddedByValue()                        {}

// UnsafeFailureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FailureServiceServer will
// result in compilation errors.
type UnsafeFailureServiceServer interface {
	mustEmbedUnimplementedFailureServiceServer()
}

func RegisterFailureServiceServer(s grpc.ServiceRegistrar, srv FailureServiceServer) {
	// If the following call pancis, it indicates UnimplementedFailureServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FailureService_ServiceDesc, srv)
}

func _FailureService_Inject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FailureServiceServer).Inject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FailureService_Inject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FailureServiceServer).Inject(ctx, req.(*InjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FailureService_ServiceDesc is the grpc.ServiceDesc for FailureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FailureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.failure.FailureService",
	HandlerType: (*FailureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Inject",
			Handler:    _FailureService_Inject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "failure.proto",
}

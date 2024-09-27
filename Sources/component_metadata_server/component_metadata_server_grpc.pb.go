// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: component_metadata_server.proto

package component_metadata_server

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
	ComponentMetadataServerService_SetMetadata_FullMethodName = "/mavsdk.rpc.component_metadata_server.ComponentMetadataServerService/SetMetadata"
)

// ComponentMetadataServerServiceClient is the client API for ComponentMetadataServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Provide component metadata json definitions, such as parameters.
type ComponentMetadataServerServiceClient interface {
	// Provide metadata (can only be called once)
	SetMetadata(ctx context.Context, in *SetMetadataRequest, opts ...grpc.CallOption) (*SetMetadataResponse, error)
}

type componentMetadataServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComponentMetadataServerServiceClient(cc grpc.ClientConnInterface) ComponentMetadataServerServiceClient {
	return &componentMetadataServerServiceClient{cc}
}

func (c *componentMetadataServerServiceClient) SetMetadata(ctx context.Context, in *SetMetadataRequest, opts ...grpc.CallOption) (*SetMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetMetadataResponse)
	err := c.cc.Invoke(ctx, ComponentMetadataServerService_SetMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComponentMetadataServerServiceServer is the server API for ComponentMetadataServerService service.
// All implementations must embed UnimplementedComponentMetadataServerServiceServer
// for forward compatibility.
//
// Provide component metadata json definitions, such as parameters.
type ComponentMetadataServerServiceServer interface {
	// Provide metadata (can only be called once)
	SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error)
	mustEmbedUnimplementedComponentMetadataServerServiceServer()
}

// UnimplementedComponentMetadataServerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedComponentMetadataServerServiceServer struct{}

func (UnimplementedComponentMetadataServerServiceServer) SetMetadata(context.Context, *SetMetadataRequest) (*SetMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMetadata not implemented")
}
func (UnimplementedComponentMetadataServerServiceServer) mustEmbedUnimplementedComponentMetadataServerServiceServer() {
}
func (UnimplementedComponentMetadataServerServiceServer) testEmbeddedByValue() {}

// UnsafeComponentMetadataServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComponentMetadataServerServiceServer will
// result in compilation errors.
type UnsafeComponentMetadataServerServiceServer interface {
	mustEmbedUnimplementedComponentMetadataServerServiceServer()
}

func RegisterComponentMetadataServerServiceServer(s grpc.ServiceRegistrar, srv ComponentMetadataServerServiceServer) {
	// If the following call pancis, it indicates UnimplementedComponentMetadataServerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ComponentMetadataServerService_ServiceDesc, srv)
}

func _ComponentMetadataServerService_SetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentMetadataServerServiceServer).SetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ComponentMetadataServerService_SetMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentMetadataServerServiceServer).SetMetadata(ctx, req.(*SetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComponentMetadataServerService_ServiceDesc is the grpc.ServiceDesc for ComponentMetadataServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComponentMetadataServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.component_metadata_server.ComponentMetadataServerService",
	HandlerType: (*ComponentMetadataServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetMetadata",
			Handler:    _ComponentMetadataServerService_SetMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "component_metadata_server.proto",
}
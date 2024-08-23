// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: ftp_server.proto

package ftp_server

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
	FtpServerService_SetRootDir_FullMethodName = "/mavsdk.rpc.ftp_server.FtpServerService/SetRootDir"
)

// FtpServerServiceClient is the client API for FtpServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Provide files or directories to transfer.
type FtpServerServiceClient interface {
	// Set root directory.
	//
	// This is the directory that can then be accessed by a client.
	// The directory needs to exist when this is called.
	// The permissions are the same as the file permission for the user running the server.
	// The root directory can't be changed while an FTP process is in progress.
	SetRootDir(ctx context.Context, in *SetRootDirRequest, opts ...grpc.CallOption) (*SetRootDirResponse, error)
}

type ftpServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFtpServerServiceClient(cc grpc.ClientConnInterface) FtpServerServiceClient {
	return &ftpServerServiceClient{cc}
}

func (c *ftpServerServiceClient) SetRootDir(ctx context.Context, in *SetRootDirRequest, opts ...grpc.CallOption) (*SetRootDirResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetRootDirResponse)
	err := c.cc.Invoke(ctx, FtpServerService_SetRootDir_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FtpServerServiceServer is the server API for FtpServerService service.
// All implementations must embed UnimplementedFtpServerServiceServer
// for forward compatibility.
//
// Provide files or directories to transfer.
type FtpServerServiceServer interface {
	// Set root directory.
	//
	// This is the directory that can then be accessed by a client.
	// The directory needs to exist when this is called.
	// The permissions are the same as the file permission for the user running the server.
	// The root directory can't be changed while an FTP process is in progress.
	SetRootDir(context.Context, *SetRootDirRequest) (*SetRootDirResponse, error)
	mustEmbedUnimplementedFtpServerServiceServer()
}

// UnimplementedFtpServerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFtpServerServiceServer struct{}

func (UnimplementedFtpServerServiceServer) SetRootDir(context.Context, *SetRootDirRequest) (*SetRootDirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRootDir not implemented")
}
func (UnimplementedFtpServerServiceServer) mustEmbedUnimplementedFtpServerServiceServer() {}
func (UnimplementedFtpServerServiceServer) testEmbeddedByValue()                          {}

// UnsafeFtpServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FtpServerServiceServer will
// result in compilation errors.
type UnsafeFtpServerServiceServer interface {
	mustEmbedUnimplementedFtpServerServiceServer()
}

func RegisterFtpServerServiceServer(s grpc.ServiceRegistrar, srv FtpServerServiceServer) {
	// If the following call pancis, it indicates UnimplementedFtpServerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FtpServerService_ServiceDesc, srv)
}

func _FtpServerService_SetRootDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRootDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FtpServerServiceServer).SetRootDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FtpServerService_SetRootDir_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FtpServerServiceServer).SetRootDir(ctx, req.(*SetRootDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FtpServerService_ServiceDesc is the grpc.ServiceDesc for FtpServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FtpServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.ftp_server.FtpServerService",
	HandlerType: (*FtpServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetRootDir",
			Handler:    _FtpServerService_SetRootDir_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ftp_server.proto",
}

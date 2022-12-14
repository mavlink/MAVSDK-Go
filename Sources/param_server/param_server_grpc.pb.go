// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: param_server.proto

package param_server

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

// ParamServerServiceClient is the client API for ParamServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParamServerServiceClient interface {
	// Retrieve an int parameter.
	//
	// If the type is wrong, the result will be `WRONG_TYPE`.
	RetrieveParamInt(ctx context.Context, in *RetrieveParamIntRequest, opts ...grpc.CallOption) (*RetrieveParamIntResponse, error)
	// Provide an int parameter.
	//
	// If the type is wrong, the result will be `WRONG_TYPE`.
	ProvideParamInt(ctx context.Context, in *ProvideParamIntRequest, opts ...grpc.CallOption) (*ProvideParamIntResponse, error)
	// Retrieve a float parameter.
	//
	// If the type is wrong, the result will be `WRONG_TYPE`.
	RetrieveParamFloat(ctx context.Context, in *RetrieveParamFloatRequest, opts ...grpc.CallOption) (*RetrieveParamFloatResponse, error)
	// Provide a float parameter.
	//
	// If the type is wrong, the result will be `WRONG_TYPE`.
	ProvideParamFloat(ctx context.Context, in *ProvideParamFloatRequest, opts ...grpc.CallOption) (*ProvideParamFloatResponse, error)
	// Retrieve all parameters.
	RetrieveAllParams(ctx context.Context, in *RetrieveAllParamsRequest, opts ...grpc.CallOption) (*RetrieveAllParamsResponse, error)
}

type paramServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParamServerServiceClient(cc grpc.ClientConnInterface) ParamServerServiceClient {
	return &paramServerServiceClient{cc}
}

func (c *paramServerServiceClient) RetrieveParamInt(ctx context.Context, in *RetrieveParamIntRequest, opts ...grpc.CallOption) (*RetrieveParamIntResponse, error) {
	out := new(RetrieveParamIntResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.param_server.ParamServerService/RetrieveParamInt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramServerServiceClient) ProvideParamInt(ctx context.Context, in *ProvideParamIntRequest, opts ...grpc.CallOption) (*ProvideParamIntResponse, error) {
	out := new(ProvideParamIntResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.param_server.ParamServerService/ProvideParamInt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramServerServiceClient) RetrieveParamFloat(ctx context.Context, in *RetrieveParamFloatRequest, opts ...grpc.CallOption) (*RetrieveParamFloatResponse, error) {
	out := new(RetrieveParamFloatResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.param_server.ParamServerService/RetrieveParamFloat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramServerServiceClient) ProvideParamFloat(ctx context.Context, in *ProvideParamFloatRequest, opts ...grpc.CallOption) (*ProvideParamFloatResponse, error) {
	out := new(ProvideParamFloatResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.param_server.ParamServerService/ProvideParamFloat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paramServerServiceClient) RetrieveAllParams(ctx context.Context, in *RetrieveAllParamsRequest, opts ...grpc.CallOption) (*RetrieveAllParamsResponse, error) {
	out := new(RetrieveAllParamsResponse)
	err := c.cc.Invoke(ctx, "/mavsdk.rpc.param_server.ParamServerService/RetrieveAllParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParamServerServiceServer is the server API for ParamServerService service.
// All implementations must embed UnimplementedParamServerServiceServer
// for forward compatibility
type ParamServerServiceServer interface {
	// Retrieve an int parameter.
	//
	// If the type is wrong, the result will be `WRONG_TYPE`.
	RetrieveParamInt(context.Context, *RetrieveParamIntRequest) (*RetrieveParamIntResponse, error)
	// Provide an int parameter.
	//
	// If the type is wrong, the result will be `WRONG_TYPE`.
	ProvideParamInt(context.Context, *ProvideParamIntRequest) (*ProvideParamIntResponse, error)
	// Retrieve a float parameter.
	//
	// If the type is wrong, the result will be `WRONG_TYPE`.
	RetrieveParamFloat(context.Context, *RetrieveParamFloatRequest) (*RetrieveParamFloatResponse, error)
	// Provide a float parameter.
	//
	// If the type is wrong, the result will be `WRONG_TYPE`.
	ProvideParamFloat(context.Context, *ProvideParamFloatRequest) (*ProvideParamFloatResponse, error)
	// Retrieve all parameters.
	RetrieveAllParams(context.Context, *RetrieveAllParamsRequest) (*RetrieveAllParamsResponse, error)
	mustEmbedUnimplementedParamServerServiceServer()
}

// UnimplementedParamServerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedParamServerServiceServer struct {
}

func (UnimplementedParamServerServiceServer) RetrieveParamInt(context.Context, *RetrieveParamIntRequest) (*RetrieveParamIntResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveParamInt not implemented")
}
func (UnimplementedParamServerServiceServer) ProvideParamInt(context.Context, *ProvideParamIntRequest) (*ProvideParamIntResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvideParamInt not implemented")
}
func (UnimplementedParamServerServiceServer) RetrieveParamFloat(context.Context, *RetrieveParamFloatRequest) (*RetrieveParamFloatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveParamFloat not implemented")
}
func (UnimplementedParamServerServiceServer) ProvideParamFloat(context.Context, *ProvideParamFloatRequest) (*ProvideParamFloatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvideParamFloat not implemented")
}
func (UnimplementedParamServerServiceServer) RetrieveAllParams(context.Context, *RetrieveAllParamsRequest) (*RetrieveAllParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveAllParams not implemented")
}
func (UnimplementedParamServerServiceServer) mustEmbedUnimplementedParamServerServiceServer() {}

// UnsafeParamServerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParamServerServiceServer will
// result in compilation errors.
type UnsafeParamServerServiceServer interface {
	mustEmbedUnimplementedParamServerServiceServer()
}

func RegisterParamServerServiceServer(s grpc.ServiceRegistrar, srv ParamServerServiceServer) {
	s.RegisterService(&ParamServerService_ServiceDesc, srv)
}

func _ParamServerService_RetrieveParamInt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveParamIntRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamServerServiceServer).RetrieveParamInt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.param_server.ParamServerService/RetrieveParamInt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamServerServiceServer).RetrieveParamInt(ctx, req.(*RetrieveParamIntRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamServerService_ProvideParamInt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvideParamIntRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamServerServiceServer).ProvideParamInt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.param_server.ParamServerService/ProvideParamInt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamServerServiceServer).ProvideParamInt(ctx, req.(*ProvideParamIntRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamServerService_RetrieveParamFloat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveParamFloatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamServerServiceServer).RetrieveParamFloat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.param_server.ParamServerService/RetrieveParamFloat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamServerServiceServer).RetrieveParamFloat(ctx, req.(*RetrieveParamFloatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamServerService_ProvideParamFloat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvideParamFloatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamServerServiceServer).ProvideParamFloat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.param_server.ParamServerService/ProvideParamFloat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamServerServiceServer).ProvideParamFloat(ctx, req.(*ProvideParamFloatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParamServerService_RetrieveAllParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveAllParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParamServerServiceServer).RetrieveAllParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mavsdk.rpc.param_server.ParamServerService/RetrieveAllParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParamServerServiceServer).RetrieveAllParams(ctx, req.(*RetrieveAllParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ParamServerService_ServiceDesc is the grpc.ServiceDesc for ParamServerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParamServerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mavsdk.rpc.param_server.ParamServerService",
	HandlerType: (*ParamServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveParamInt",
			Handler:    _ParamServerService_RetrieveParamInt_Handler,
		},
		{
			MethodName: "ProvideParamInt",
			Handler:    _ParamServerService_ProvideParamInt_Handler,
		},
		{
			MethodName: "RetrieveParamFloat",
			Handler:    _ParamServerService_RetrieveParamFloat_Handler,
		},
		{
			MethodName: "ProvideParamFloat",
			Handler:    _ParamServerService_ProvideParamFloat_Handler,
		},
		{
			MethodName: "RetrieveAllParams",
			Handler:    _ParamServerService_RetrieveAllParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "param_server.proto",
}

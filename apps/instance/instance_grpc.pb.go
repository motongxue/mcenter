// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: apps/instance/pb/instance.proto

package instance

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// 实例注册
	RegistryInstance(ctx context.Context, in *RegistryRequest, opts ...grpc.CallOption) (*Instance, error)
	// 实例上报心跳
	Heartbeat(ctx context.Context, opts ...grpc.CallOption) (Service_HeartbeatClient, error)
	// 实例注销
	UnRegistry(ctx context.Context, in *UnregistryRequest, opts ...grpc.CallOption) (*Instance, error)
	// 实例搜索, 用于客户端做服务发现
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*InstanceSet, error)
	// 查询实例详情
	DescribeInstance(ctx context.Context, in *DescribeInstanceRequest, opts ...grpc.CallOption) (*Instance, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) RegistryInstance(ctx context.Context, in *RegistryRequest, opts ...grpc.CallOption) (*Instance, error) {
	out := new(Instance)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.instance.Service/RegistryInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Heartbeat(ctx context.Context, opts ...grpc.CallOption) (Service_HeartbeatClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[0], "/infraboard.mcenter.instance.Service/Heartbeat", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceHeartbeatClient{stream}
	return x, nil
}

type Service_HeartbeatClient interface {
	Send(*HeartbeatRequest) error
	Recv() (*HeartbeatResponse, error)
	grpc.ClientStream
}

type serviceHeartbeatClient struct {
	grpc.ClientStream
}

func (x *serviceHeartbeatClient) Send(m *HeartbeatRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serviceHeartbeatClient) Recv() (*HeartbeatResponse, error) {
	m := new(HeartbeatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) UnRegistry(ctx context.Context, in *UnregistryRequest, opts ...grpc.CallOption) (*Instance, error) {
	out := new(Instance)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.instance.Service/UnRegistry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*InstanceSet, error) {
	out := new(InstanceSet)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.instance.Service/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeInstance(ctx context.Context, in *DescribeInstanceRequest, opts ...grpc.CallOption) (*Instance, error) {
	out := new(Instance)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.instance.Service/DescribeInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// 实例注册
	RegistryInstance(context.Context, *RegistryRequest) (*Instance, error)
	// 实例上报心跳
	Heartbeat(Service_HeartbeatServer) error
	// 实例注销
	UnRegistry(context.Context, *UnregistryRequest) (*Instance, error)
	// 实例搜索, 用于客户端做服务发现
	Search(context.Context, *SearchRequest) (*InstanceSet, error)
	// 查询实例详情
	DescribeInstance(context.Context, *DescribeInstanceRequest) (*Instance, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) RegistryInstance(context.Context, *RegistryRequest) (*Instance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistryInstance not implemented")
}
func (UnimplementedServiceServer) Heartbeat(Service_HeartbeatServer) error {
	return status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedServiceServer) UnRegistry(context.Context, *UnregistryRequest) (*Instance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegistry not implemented")
}
func (UnimplementedServiceServer) Search(context.Context, *SearchRequest) (*InstanceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedServiceServer) DescribeInstance(context.Context, *DescribeInstanceRequest) (*Instance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeInstance not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_RegistryInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RegistryInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.instance.Service/RegistryInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RegistryInstance(ctx, req.(*RegistryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Heartbeat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServiceServer).Heartbeat(&serviceHeartbeatServer{stream})
}

type Service_HeartbeatServer interface {
	Send(*HeartbeatResponse) error
	Recv() (*HeartbeatRequest, error)
	grpc.ServerStream
}

type serviceHeartbeatServer struct {
	grpc.ServerStream
}

func (x *serviceHeartbeatServer) Send(m *HeartbeatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serviceHeartbeatServer) Recv() (*HeartbeatRequest, error) {
	m := new(HeartbeatRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Service_UnRegistry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregistryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UnRegistry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.instance.Service/UnRegistry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UnRegistry(ctx, req.(*UnregistryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.instance.Service/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.instance.Service/DescribeInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeInstance(ctx, req.(*DescribeInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mcenter.instance.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegistryInstance",
			Handler:    _Service_RegistryInstance_Handler,
		},
		{
			MethodName: "UnRegistry",
			Handler:    _Service_UnRegistry_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Service_Search_Handler,
		},
		{
			MethodName: "DescribeInstance",
			Handler:    _Service_DescribeInstance_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Heartbeat",
			Handler:       _Service_Heartbeat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "apps/instance/pb/instance.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: appuser/gateway/v1/role/user/user.proto

package user

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
	Gateway_CreateUser_FullMethodName    = "/appuser.gateway.role.user.v1.Gateway/CreateUser"
	Gateway_GetUsers_FullMethodName      = "/appuser.gateway.role.user.v1.Gateway/GetUsers"
	Gateway_DeleteUser_FullMethodName    = "/appuser.gateway.role.user.v1.Gateway/DeleteUser"
	Gateway_DeleteAppUser_FullMethodName = "/appuser.gateway.role.user.v1.Gateway/DeleteAppUser"
	Gateway_CreateAppUser_FullMethodName = "/appuser.gateway.role.user.v1.Gateway/CreateAppUser"
	Gateway_GetAppUsers_FullMethodName   = "/appuser.gateway.role.user.v1.Gateway/GetAppUsers"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	DeleteAppUser(ctx context.Context, in *DeleteAppUserRequest, opts ...grpc.CallOption) (*DeleteAppUserResponse, error)
	CreateAppUser(ctx context.Context, in *CreateAppUserRequest, opts ...grpc.CallOption) (*CreateAppUserResponse, error)
	GetAppUsers(ctx context.Context, in *GetAppUsersRequest, opts ...grpc.CallOption) (*GetAppUsersResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, Gateway_GetUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteAppUser(ctx context.Context, in *DeleteAppUserRequest, opts ...grpc.CallOption) (*DeleteAppUserResponse, error) {
	out := new(DeleteAppUserResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteAppUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppUser(ctx context.Context, in *CreateAppUserRequest, opts ...grpc.CallOption) (*CreateAppUserResponse, error) {
	out := new(CreateAppUserResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAppUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppUsers(ctx context.Context, in *GetAppUsersRequest, opts ...grpc.CallOption) (*GetAppUsersResponse, error) {
	out := new(GetAppUsersResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	DeleteAppUser(context.Context, *DeleteAppUserRequest) (*DeleteAppUserResponse, error)
	CreateAppUser(context.Context, *CreateAppUserRequest) (*CreateAppUserResponse, error)
	GetAppUsers(context.Context, *GetAppUsersRequest) (*GetAppUsersResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct{}

func (UnimplementedGatewayServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func (UnimplementedGatewayServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}

func (UnimplementedGatewayServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func (UnimplementedGatewayServer) DeleteAppUser(context.Context, *DeleteAppUserRequest) (*DeleteAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAppUser not implemented")
}

func (UnimplementedGatewayServer) CreateAppUser(context.Context, *CreateAppUserRequest) (*CreateAppUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppUser not implemented")
}

func (UnimplementedGatewayServer) GetAppUsers(context.Context, *GetAppUsersRequest) (*GetAppUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppUsers not implemented")
}
func (UnimplementedGatewayServer) mustEmbedUnimplementedGatewayServer() {}

// UnsafeGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServer will
// result in compilation errors.
type UnsafeGatewayServer interface {
	mustEmbedUnimplementedGatewayServer()
}

func RegisterGatewayServer(s grpc.ServiceRegistrar, srv GatewayServer) {
	s.RegisterService(&Gateway_ServiceDesc, srv)
}

func _Gateway_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteAppUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteAppUser(ctx, req.(*DeleteAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAppUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppUser(ctx, req.(*CreateAppUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppUsers(ctx, req.(*GetAppUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.gateway.role.user.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Gateway_CreateUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Gateway_GetUsers_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Gateway_DeleteUser_Handler,
		},
		{
			MethodName: "DeleteAppUser",
			Handler:    _Gateway_DeleteAppUser_Handler,
		},
		{
			MethodName: "CreateAppUser",
			Handler:    _Gateway_CreateAppUser_Handler,
		},
		{
			MethodName: "GetAppUsers",
			Handler:    _Gateway_GetAppUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appuser/gateway/v1/role/user/user.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: appuser/gateway/v1/admin/admin.proto

package admin

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
	Gateway_CreateAdminApps_FullMethodName    = "/appuser.gateway.admin.v1.Gateway/CreateAdminApps"
	Gateway_CreateGenesisRoles_FullMethodName = "/appuser.gateway.admin.v1.Gateway/CreateGenesisRoles"
	Gateway_CreateGenesisUser_FullMethodName  = "/appuser.gateway.admin.v1.Gateway/CreateGenesisUser"
	Gateway_GetAdminApps_FullMethodName       = "/appuser.gateway.admin.v1.Gateway/GetAdminApps"
	Gateway_GetGenesisRoles_FullMethodName    = "/appuser.gateway.admin.v1.Gateway/GetGenesisRoles"
	Gateway_GetGenesisUsers_FullMethodName    = "/appuser.gateway.admin.v1.Gateway/GetGenesisUsers"
	Gateway_AuthorizeGenesis_FullMethodName   = "/appuser.gateway.admin.v1.Gateway/AuthorizeGenesis"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	// VPN only apis
	// Create genesis / admin apps
	CreateAdminApps(ctx context.Context, in *CreateAdminAppsRequest, opts ...grpc.CallOption) (*CreateAdminAppsResponse, error)
	// Create genesis / admin role
	CreateGenesisRoles(ctx context.Context, in *CreateGenesisRolesRequest, opts ...grpc.CallOption) (*CreateGenesisRolesResponse, error)
	// Create genesis / admin user
	CreateGenesisUser(ctx context.Context, in *CreateGenesisUserRequest, opts ...grpc.CallOption) (*CreateGenesisUserResponse, error)
	GetAdminApps(ctx context.Context, in *GetAdminAppsRequest, opts ...grpc.CallOption) (*GetAdminAppsResponse, error)
	GetGenesisRoles(ctx context.Context, in *GetGenesisRolesRequest, opts ...grpc.CallOption) (*GetGenesisRolesResponse, error)
	GetGenesisUsers(ctx context.Context, in *GetGenesisUsersRequest, opts ...grpc.CallOption) (*GetGenesisUsersResponse, error)
	AuthorizeGenesis(ctx context.Context, in *AuthorizeGenesisRequest, opts ...grpc.CallOption) (*AuthorizeGenesisResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateAdminApps(ctx context.Context, in *CreateAdminAppsRequest, opts ...grpc.CallOption) (*CreateAdminAppsResponse, error) {
	out := new(CreateAdminAppsResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAdminApps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateGenesisRoles(ctx context.Context, in *CreateGenesisRolesRequest, opts ...grpc.CallOption) (*CreateGenesisRolesResponse, error) {
	out := new(CreateGenesisRolesResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateGenesisRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateGenesisUser(ctx context.Context, in *CreateGenesisUserRequest, opts ...grpc.CallOption) (*CreateGenesisUserResponse, error) {
	out := new(CreateGenesisUserResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateGenesisUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAdminApps(ctx context.Context, in *GetAdminAppsRequest, opts ...grpc.CallOption) (*GetAdminAppsResponse, error) {
	out := new(GetAdminAppsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAdminApps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetGenesisRoles(ctx context.Context, in *GetGenesisRolesRequest, opts ...grpc.CallOption) (*GetGenesisRolesResponse, error) {
	out := new(GetGenesisRolesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetGenesisRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetGenesisUsers(ctx context.Context, in *GetGenesisUsersRequest, opts ...grpc.CallOption) (*GetGenesisUsersResponse, error) {
	out := new(GetGenesisUsersResponse)
	err := c.cc.Invoke(ctx, Gateway_GetGenesisUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AuthorizeGenesis(ctx context.Context, in *AuthorizeGenesisRequest, opts ...grpc.CallOption) (*AuthorizeGenesisResponse, error) {
	out := new(AuthorizeGenesisResponse)
	err := c.cc.Invoke(ctx, Gateway_AuthorizeGenesis_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	// VPN only apis
	// Create genesis / admin apps
	CreateAdminApps(context.Context, *CreateAdminAppsRequest) (*CreateAdminAppsResponse, error)
	// Create genesis / admin role
	CreateGenesisRoles(context.Context, *CreateGenesisRolesRequest) (*CreateGenesisRolesResponse, error)
	// Create genesis / admin user
	CreateGenesisUser(context.Context, *CreateGenesisUserRequest) (*CreateGenesisUserResponse, error)
	GetAdminApps(context.Context, *GetAdminAppsRequest) (*GetAdminAppsResponse, error)
	GetGenesisRoles(context.Context, *GetGenesisRolesRequest) (*GetGenesisRolesResponse, error)
	GetGenesisUsers(context.Context, *GetGenesisUsersRequest) (*GetGenesisUsersResponse, error)
	AuthorizeGenesis(context.Context, *AuthorizeGenesisRequest) (*AuthorizeGenesisResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct{}

func (UnimplementedGatewayServer) CreateAdminApps(context.Context, *CreateAdminAppsRequest) (*CreateAdminAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdminApps not implemented")
}

func (UnimplementedGatewayServer) CreateGenesisRoles(context.Context, *CreateGenesisRolesRequest) (*CreateGenesisRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGenesisRoles not implemented")
}

func (UnimplementedGatewayServer) CreateGenesisUser(context.Context, *CreateGenesisUserRequest) (*CreateGenesisUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGenesisUser not implemented")
}

func (UnimplementedGatewayServer) GetAdminApps(context.Context, *GetAdminAppsRequest) (*GetAdminAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminApps not implemented")
}

func (UnimplementedGatewayServer) GetGenesisRoles(context.Context, *GetGenesisRolesRequest) (*GetGenesisRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenesisRoles not implemented")
}

func (UnimplementedGatewayServer) GetGenesisUsers(context.Context, *GetGenesisUsersRequest) (*GetGenesisUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenesisUsers not implemented")
}

func (UnimplementedGatewayServer) AuthorizeGenesis(context.Context, *AuthorizeGenesisRequest) (*AuthorizeGenesisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeGenesis not implemented")
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

func _Gateway_CreateAdminApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdminAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAdminApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAdminApps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAdminApps(ctx, req.(*CreateAdminAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateGenesisRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGenesisRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateGenesisRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateGenesisRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateGenesisRoles(ctx, req.(*CreateGenesisRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateGenesisUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGenesisUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateGenesisUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateGenesisUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateGenesisUser(ctx, req.(*CreateGenesisUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAdminApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAdminApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAdminApps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAdminApps(ctx, req.(*GetAdminAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetGenesisRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGenesisRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetGenesisRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetGenesisRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetGenesisRoles(ctx, req.(*GetGenesisRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetGenesisUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGenesisUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetGenesisUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetGenesisUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetGenesisUsers(ctx, req.(*GetGenesisUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AuthorizeGenesis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeGenesisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AuthorizeGenesis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AuthorizeGenesis_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AuthorizeGenesis(ctx, req.(*AuthorizeGenesisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appuser.gateway.admin.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAdminApps",
			Handler:    _Gateway_CreateAdminApps_Handler,
		},
		{
			MethodName: "CreateGenesisRoles",
			Handler:    _Gateway_CreateGenesisRoles_Handler,
		},
		{
			MethodName: "CreateGenesisUser",
			Handler:    _Gateway_CreateGenesisUser_Handler,
		},
		{
			MethodName: "GetAdminApps",
			Handler:    _Gateway_GetAdminApps_Handler,
		},
		{
			MethodName: "GetGenesisRoles",
			Handler:    _Gateway_GetGenesisRoles_Handler,
		},
		{
			MethodName: "GetGenesisUsers",
			Handler:    _Gateway_GetGenesisUsers_Handler,
		},
		{
			MethodName: "AuthorizeGenesis",
			Handler:    _Gateway_AuthorizeGenesis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appuser/gateway/v1/admin/admin.proto",
}

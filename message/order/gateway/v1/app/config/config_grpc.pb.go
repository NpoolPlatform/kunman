// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: order/gateway/v1/app/config/config.proto

package config

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
	Gateway_CreateAppConfig_FullMethodName      = "/order.gateway.app.config.v1.Gateway/CreateAppConfig"
	Gateway_UpdateAppConfig_FullMethodName      = "/order.gateway.app.config.v1.Gateway/UpdateAppConfig"
	Gateway_GetAppConfig_FullMethodName         = "/order.gateway.app.config.v1.Gateway/GetAppConfig"
	Gateway_AdminCreateAppConfig_FullMethodName = "/order.gateway.app.config.v1.Gateway/AdminCreateAppConfig"
	Gateway_AdminUpdateAppConfig_FullMethodName = "/order.gateway.app.config.v1.Gateway/AdminUpdateAppConfig"
	Gateway_AdminGetAppConfigs_FullMethodName   = "/order.gateway.app.config.v1.Gateway/AdminGetAppConfigs"
	Gateway_AdminDeleteAppConfig_FullMethodName = "/order.gateway.app.config.v1.Gateway/AdminDeleteAppConfig"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateAppConfig(ctx context.Context, in *CreateAppConfigRequest, opts ...grpc.CallOption) (*CreateAppConfigResponse, error)
	UpdateAppConfig(ctx context.Context, in *UpdateAppConfigRequest, opts ...grpc.CallOption) (*UpdateAppConfigResponse, error)
	GetAppConfig(ctx context.Context, in *GetAppConfigRequest, opts ...grpc.CallOption) (*GetAppConfigResponse, error)
	// Admin apis
	AdminCreateAppConfig(ctx context.Context, in *AdminCreateAppConfigRequest, opts ...grpc.CallOption) (*AdminCreateAppConfigResponse, error)
	AdminUpdateAppConfig(ctx context.Context, in *AdminUpdateAppConfigRequest, opts ...grpc.CallOption) (*AdminUpdateAppConfigResponse, error)
	AdminGetAppConfigs(ctx context.Context, in *AdminGetAppConfigsRequest, opts ...grpc.CallOption) (*AdminGetAppConfigsResponse, error)
	AdminDeleteAppConfig(ctx context.Context, in *AdminDeleteAppConfigRequest, opts ...grpc.CallOption) (*AdminDeleteAppConfigResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateAppConfig(ctx context.Context, in *CreateAppConfigRequest, opts ...grpc.CallOption) (*CreateAppConfigResponse, error) {
	out := new(CreateAppConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppConfig(ctx context.Context, in *UpdateAppConfigRequest, opts ...grpc.CallOption) (*UpdateAppConfigResponse, error) {
	out := new(UpdateAppConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppConfig(ctx context.Context, in *GetAppConfigRequest, opts ...grpc.CallOption) (*GetAppConfigResponse, error) {
	out := new(GetAppConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminCreateAppConfig(ctx context.Context, in *AdminCreateAppConfigRequest, opts ...grpc.CallOption) (*AdminCreateAppConfigResponse, error) {
	out := new(AdminCreateAppConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminCreateAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdateAppConfig(ctx context.Context, in *AdminUpdateAppConfigRequest, opts ...grpc.CallOption) (*AdminUpdateAppConfigResponse, error) {
	out := new(AdminUpdateAppConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminUpdateAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetAppConfigs(ctx context.Context, in *AdminGetAppConfigsRequest, opts ...grpc.CallOption) (*AdminGetAppConfigsResponse, error) {
	out := new(AdminGetAppConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminGetAppConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeleteAppConfig(ctx context.Context, in *AdminDeleteAppConfigRequest, opts ...grpc.CallOption) (*AdminDeleteAppConfigResponse, error) {
	out := new(AdminDeleteAppConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminDeleteAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateAppConfig(context.Context, *CreateAppConfigRequest) (*CreateAppConfigResponse, error)
	UpdateAppConfig(context.Context, *UpdateAppConfigRequest) (*UpdateAppConfigResponse, error)
	GetAppConfig(context.Context, *GetAppConfigRequest) (*GetAppConfigResponse, error)
	// Admin apis
	AdminCreateAppConfig(context.Context, *AdminCreateAppConfigRequest) (*AdminCreateAppConfigResponse, error)
	AdminUpdateAppConfig(context.Context, *AdminUpdateAppConfigRequest) (*AdminUpdateAppConfigResponse, error)
	AdminGetAppConfigs(context.Context, *AdminGetAppConfigsRequest) (*AdminGetAppConfigsResponse, error)
	AdminDeleteAppConfig(context.Context, *AdminDeleteAppConfigRequest) (*AdminDeleteAppConfigResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct{}

func (UnimplementedGatewayServer) CreateAppConfig(context.Context, *CreateAppConfigRequest) (*CreateAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppConfig not implemented")
}

func (UnimplementedGatewayServer) UpdateAppConfig(context.Context, *UpdateAppConfigRequest) (*UpdateAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppConfig not implemented")
}

func (UnimplementedGatewayServer) GetAppConfig(context.Context, *GetAppConfigRequest) (*GetAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppConfig not implemented")
}

func (UnimplementedGatewayServer) AdminCreateAppConfig(context.Context, *AdminCreateAppConfigRequest) (*AdminCreateAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateAppConfig not implemented")
}

func (UnimplementedGatewayServer) AdminUpdateAppConfig(context.Context, *AdminUpdateAppConfigRequest) (*AdminUpdateAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateAppConfig not implemented")
}

func (UnimplementedGatewayServer) AdminGetAppConfigs(context.Context, *AdminGetAppConfigsRequest) (*AdminGetAppConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetAppConfigs not implemented")
}

func (UnimplementedGatewayServer) AdminDeleteAppConfig(context.Context, *AdminDeleteAppConfigRequest) (*AdminDeleteAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteAppConfig not implemented")
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

func _Gateway_CreateAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppConfig(ctx, req.(*CreateAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppConfig(ctx, req.(*UpdateAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppConfig(ctx, req.(*GetAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminCreateAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreateAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminCreateAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreateAppConfig(ctx, req.(*AdminCreateAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdateAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdateAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminUpdateAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdateAppConfig(ctx, req.(*AdminUpdateAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetAppConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetAppConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetAppConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminGetAppConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetAppConfigs(ctx, req.(*AdminGetAppConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeleteAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeleteAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminDeleteAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeleteAppConfig(ctx, req.(*AdminDeleteAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.gateway.app.config.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppConfig",
			Handler:    _Gateway_CreateAppConfig_Handler,
		},
		{
			MethodName: "UpdateAppConfig",
			Handler:    _Gateway_UpdateAppConfig_Handler,
		},
		{
			MethodName: "GetAppConfig",
			Handler:    _Gateway_GetAppConfig_Handler,
		},
		{
			MethodName: "AdminCreateAppConfig",
			Handler:    _Gateway_AdminCreateAppConfig_Handler,
		},
		{
			MethodName: "AdminUpdateAppConfig",
			Handler:    _Gateway_AdminUpdateAppConfig_Handler,
		},
		{
			MethodName: "AdminGetAppConfigs",
			Handler:    _Gateway_AdminGetAppConfigs_Handler,
		},
		{
			MethodName: "AdminDeleteAppConfig",
			Handler:    _Gateway_AdminDeleteAppConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/gateway/v1/app/config/config.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: inspire/gateway/v1/app/good/commission/config/config.proto

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
	Gateway_CreateAppGoodCommissionConfig_FullMethodName      = "/inspire.gateway.app.good.commission.config.v1.Gateway/CreateAppGoodCommissionConfig"
	Gateway_AdminCreateAppGoodCommissionConfig_FullMethodName = "/inspire.gateway.app.good.commission.config.v1.Gateway/AdminCreateAppGoodCommissionConfig"
	Gateway_UpdateAppGoodCommissionConfig_FullMethodName      = "/inspire.gateway.app.good.commission.config.v1.Gateway/UpdateAppGoodCommissionConfig"
	Gateway_AdminUpdateAppGoodCommissionConfig_FullMethodName = "/inspire.gateway.app.good.commission.config.v1.Gateway/AdminUpdateAppGoodCommissionConfig"
	Gateway_GetAppGoodCommissionConfigs_FullMethodName        = "/inspire.gateway.app.good.commission.config.v1.Gateway/GetAppGoodCommissionConfigs"
	Gateway_AdminGetAppGoodCommissionConfigs_FullMethodName   = "/inspire.gateway.app.good.commission.config.v1.Gateway/AdminGetAppGoodCommissionConfigs"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateAppGoodCommissionConfig(ctx context.Context, in *CreateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*CreateAppGoodCommissionConfigResponse, error)
	AdminCreateAppGoodCommissionConfig(ctx context.Context, in *AdminCreateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*AdminCreateAppGoodCommissionConfigResponse, error)
	UpdateAppGoodCommissionConfig(ctx context.Context, in *UpdateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*UpdateAppGoodCommissionConfigResponse, error)
	AdminUpdateAppGoodCommissionConfig(ctx context.Context, in *AdminUpdateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*AdminUpdateAppGoodCommissionConfigResponse, error)
	GetAppGoodCommissionConfigs(ctx context.Context, in *GetAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*GetAppGoodCommissionConfigsResponse, error)
	AdminGetAppGoodCommissionConfigs(ctx context.Context, in *AdminGetAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*AdminGetAppGoodCommissionConfigsResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateAppGoodCommissionConfig(ctx context.Context, in *CreateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*CreateAppGoodCommissionConfigResponse, error) {
	out := new(CreateAppGoodCommissionConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAppGoodCommissionConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminCreateAppGoodCommissionConfig(ctx context.Context, in *AdminCreateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*AdminCreateAppGoodCommissionConfigResponse, error) {
	out := new(AdminCreateAppGoodCommissionConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminCreateAppGoodCommissionConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppGoodCommissionConfig(ctx context.Context, in *UpdateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*UpdateAppGoodCommissionConfigResponse, error) {
	out := new(UpdateAppGoodCommissionConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateAppGoodCommissionConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdateAppGoodCommissionConfig(ctx context.Context, in *AdminUpdateAppGoodCommissionConfigRequest, opts ...grpc.CallOption) (*AdminUpdateAppGoodCommissionConfigResponse, error) {
	out := new(AdminUpdateAppGoodCommissionConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminUpdateAppGoodCommissionConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppGoodCommissionConfigs(ctx context.Context, in *GetAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*GetAppGoodCommissionConfigsResponse, error) {
	out := new(GetAppGoodCommissionConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppGoodCommissionConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetAppGoodCommissionConfigs(ctx context.Context, in *AdminGetAppGoodCommissionConfigsRequest, opts ...grpc.CallOption) (*AdminGetAppGoodCommissionConfigsResponse, error) {
	out := new(AdminGetAppGoodCommissionConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminGetAppGoodCommissionConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateAppGoodCommissionConfig(context.Context, *CreateAppGoodCommissionConfigRequest) (*CreateAppGoodCommissionConfigResponse, error)
	AdminCreateAppGoodCommissionConfig(context.Context, *AdminCreateAppGoodCommissionConfigRequest) (*AdminCreateAppGoodCommissionConfigResponse, error)
	UpdateAppGoodCommissionConfig(context.Context, *UpdateAppGoodCommissionConfigRequest) (*UpdateAppGoodCommissionConfigResponse, error)
	AdminUpdateAppGoodCommissionConfig(context.Context, *AdminUpdateAppGoodCommissionConfigRequest) (*AdminUpdateAppGoodCommissionConfigResponse, error)
	GetAppGoodCommissionConfigs(context.Context, *GetAppGoodCommissionConfigsRequest) (*GetAppGoodCommissionConfigsResponse, error)
	AdminGetAppGoodCommissionConfigs(context.Context, *AdminGetAppGoodCommissionConfigsRequest) (*AdminGetAppGoodCommissionConfigsResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct{}

func (UnimplementedGatewayServer) CreateAppGoodCommissionConfig(context.Context, *CreateAppGoodCommissionConfigRequest) (*CreateAppGoodCommissionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppGoodCommissionConfig not implemented")
}

func (UnimplementedGatewayServer) AdminCreateAppGoodCommissionConfig(context.Context, *AdminCreateAppGoodCommissionConfigRequest) (*AdminCreateAppGoodCommissionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateAppGoodCommissionConfig not implemented")
}

func (UnimplementedGatewayServer) UpdateAppGoodCommissionConfig(context.Context, *UpdateAppGoodCommissionConfigRequest) (*UpdateAppGoodCommissionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppGoodCommissionConfig not implemented")
}

func (UnimplementedGatewayServer) AdminUpdateAppGoodCommissionConfig(context.Context, *AdminUpdateAppGoodCommissionConfigRequest) (*AdminUpdateAppGoodCommissionConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateAppGoodCommissionConfig not implemented")
}

func (UnimplementedGatewayServer) GetAppGoodCommissionConfigs(context.Context, *GetAppGoodCommissionConfigsRequest) (*GetAppGoodCommissionConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppGoodCommissionConfigs not implemented")
}

func (UnimplementedGatewayServer) AdminGetAppGoodCommissionConfigs(context.Context, *AdminGetAppGoodCommissionConfigsRequest) (*AdminGetAppGoodCommissionConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetAppGoodCommissionConfigs not implemented")
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

func _Gateway_CreateAppGoodCommissionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppGoodCommissionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppGoodCommissionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAppGoodCommissionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppGoodCommissionConfig(ctx, req.(*CreateAppGoodCommissionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminCreateAppGoodCommissionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateAppGoodCommissionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreateAppGoodCommissionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminCreateAppGoodCommissionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreateAppGoodCommissionConfig(ctx, req.(*AdminCreateAppGoodCommissionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppGoodCommissionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppGoodCommissionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppGoodCommissionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateAppGoodCommissionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppGoodCommissionConfig(ctx, req.(*UpdateAppGoodCommissionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdateAppGoodCommissionConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateAppGoodCommissionConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdateAppGoodCommissionConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminUpdateAppGoodCommissionConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdateAppGoodCommissionConfig(ctx, req.(*AdminUpdateAppGoodCommissionConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppGoodCommissionConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppGoodCommissionConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppGoodCommissionConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppGoodCommissionConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppGoodCommissionConfigs(ctx, req.(*GetAppGoodCommissionConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetAppGoodCommissionConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetAppGoodCommissionConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetAppGoodCommissionConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminGetAppGoodCommissionConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetAppGoodCommissionConfigs(ctx, req.(*AdminGetAppGoodCommissionConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.gateway.app.good.commission.config.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAppGoodCommissionConfig",
			Handler:    _Gateway_CreateAppGoodCommissionConfig_Handler,
		},
		{
			MethodName: "AdminCreateAppGoodCommissionConfig",
			Handler:    _Gateway_AdminCreateAppGoodCommissionConfig_Handler,
		},
		{
			MethodName: "UpdateAppGoodCommissionConfig",
			Handler:    _Gateway_UpdateAppGoodCommissionConfig_Handler,
		},
		{
			MethodName: "AdminUpdateAppGoodCommissionConfig",
			Handler:    _Gateway_AdminUpdateAppGoodCommissionConfig_Handler,
		},
		{
			MethodName: "GetAppGoodCommissionConfigs",
			Handler:    _Gateway_GetAppGoodCommissionConfigs_Handler,
		},
		{
			MethodName: "AdminGetAppGoodCommissionConfigs",
			Handler:    _Gateway_AdminGetAppGoodCommissionConfigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inspire/gateway/v1/app/good/commission/config/config.proto",
}

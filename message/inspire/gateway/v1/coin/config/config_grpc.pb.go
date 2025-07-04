// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: inspire/gateway/v1/coin/config/config.proto

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
	Gateway_AdminCreateCoinConfig_FullMethodName = "/inspire.gateway.coin.config.v1.Gateway/AdminCreateCoinConfig"
	Gateway_AdminUpdateCoinConfig_FullMethodName = "/inspire.gateway.coin.config.v1.Gateway/AdminUpdateCoinConfig"
	Gateway_AdminGetCoinConfigs_FullMethodName   = "/inspire.gateway.coin.config.v1.Gateway/AdminGetCoinConfigs"
	Gateway_AdminDeleteCoinConfig_FullMethodName = "/inspire.gateway.coin.config.v1.Gateway/AdminDeleteCoinConfig"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	AdminCreateCoinConfig(ctx context.Context, in *AdminCreateCoinConfigRequest, opts ...grpc.CallOption) (*AdminCreateCoinConfigResponse, error)
	AdminUpdateCoinConfig(ctx context.Context, in *AdminUpdateCoinConfigRequest, opts ...grpc.CallOption) (*AdminUpdateCoinConfigResponse, error)
	AdminGetCoinConfigs(ctx context.Context, in *AdminGetCoinConfigsRequest, opts ...grpc.CallOption) (*AdminGetCoinConfigsResponse, error)
	AdminDeleteCoinConfig(ctx context.Context, in *AdminDeleteCoinConfigRequest, opts ...grpc.CallOption) (*AdminDeleteCoinConfigResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) AdminCreateCoinConfig(ctx context.Context, in *AdminCreateCoinConfigRequest, opts ...grpc.CallOption) (*AdminCreateCoinConfigResponse, error) {
	out := new(AdminCreateCoinConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminCreateCoinConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdateCoinConfig(ctx context.Context, in *AdminUpdateCoinConfigRequest, opts ...grpc.CallOption) (*AdminUpdateCoinConfigResponse, error) {
	out := new(AdminUpdateCoinConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminUpdateCoinConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetCoinConfigs(ctx context.Context, in *AdminGetCoinConfigsRequest, opts ...grpc.CallOption) (*AdminGetCoinConfigsResponse, error) {
	out := new(AdminGetCoinConfigsResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminGetCoinConfigs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeleteCoinConfig(ctx context.Context, in *AdminDeleteCoinConfigRequest, opts ...grpc.CallOption) (*AdminDeleteCoinConfigResponse, error) {
	out := new(AdminDeleteCoinConfigResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminDeleteCoinConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	AdminCreateCoinConfig(context.Context, *AdminCreateCoinConfigRequest) (*AdminCreateCoinConfigResponse, error)
	AdminUpdateCoinConfig(context.Context, *AdminUpdateCoinConfigRequest) (*AdminUpdateCoinConfigResponse, error)
	AdminGetCoinConfigs(context.Context, *AdminGetCoinConfigsRequest) (*AdminGetCoinConfigsResponse, error)
	AdminDeleteCoinConfig(context.Context, *AdminDeleteCoinConfigRequest) (*AdminDeleteCoinConfigResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct{}

func (UnimplementedGatewayServer) AdminCreateCoinConfig(context.Context, *AdminCreateCoinConfigRequest) (*AdminCreateCoinConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateCoinConfig not implemented")
}

func (UnimplementedGatewayServer) AdminUpdateCoinConfig(context.Context, *AdminUpdateCoinConfigRequest) (*AdminUpdateCoinConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateCoinConfig not implemented")
}

func (UnimplementedGatewayServer) AdminGetCoinConfigs(context.Context, *AdminGetCoinConfigsRequest) (*AdminGetCoinConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetCoinConfigs not implemented")
}

func (UnimplementedGatewayServer) AdminDeleteCoinConfig(context.Context, *AdminDeleteCoinConfigRequest) (*AdminDeleteCoinConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteCoinConfig not implemented")
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

func _Gateway_AdminCreateCoinConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateCoinConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreateCoinConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminCreateCoinConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreateCoinConfig(ctx, req.(*AdminCreateCoinConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdateCoinConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateCoinConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdateCoinConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminUpdateCoinConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdateCoinConfig(ctx, req.(*AdminUpdateCoinConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetCoinConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetCoinConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetCoinConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminGetCoinConfigs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetCoinConfigs(ctx, req.(*AdminGetCoinConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeleteCoinConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteCoinConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeleteCoinConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminDeleteCoinConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeleteCoinConfig(ctx, req.(*AdminDeleteCoinConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inspire.gateway.coin.config.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminCreateCoinConfig",
			Handler:    _Gateway_AdminCreateCoinConfig_Handler,
		},
		{
			MethodName: "AdminUpdateCoinConfig",
			Handler:    _Gateway_AdminUpdateCoinConfig_Handler,
		},
		{
			MethodName: "AdminGetCoinConfigs",
			Handler:    _Gateway_AdminGetCoinConfigs_Handler,
		},
		{
			MethodName: "AdminDeleteCoinConfig",
			Handler:    _Gateway_AdminDeleteCoinConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inspire/gateway/v1/coin/config/config.proto",
}

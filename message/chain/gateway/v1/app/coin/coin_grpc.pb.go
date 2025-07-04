// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: chain/gateway/v1/app/coin/coin.proto

package coin

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
	Gateway_CreateCoin_FullMethodName  = "/chain.gateway.app.coin.v1.Gateway/CreateCoin"
	Gateway_GetCoins_FullMethodName    = "/chain.gateway.app.coin.v1.Gateway/GetCoins"
	Gateway_GetAppCoins_FullMethodName = "/chain.gateway.app.coin.v1.Gateway/GetAppCoins"
	Gateway_UpdateCoin_FullMethodName  = "/chain.gateway.app.coin.v1.Gateway/UpdateCoin"
	Gateway_DeleteCoin_FullMethodName  = "/chain.gateway.app.coin.v1.Gateway/DeleteCoin"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateCoin(ctx context.Context, in *CreateCoinRequest, opts ...grpc.CallOption) (*CreateCoinResponse, error)
	GetCoins(ctx context.Context, in *GetCoinsRequest, opts ...grpc.CallOption) (*GetCoinsResponse, error)
	GetAppCoins(ctx context.Context, in *GetAppCoinsRequest, opts ...grpc.CallOption) (*GetAppCoinsResponse, error)
	UpdateCoin(ctx context.Context, in *UpdateCoinRequest, opts ...grpc.CallOption) (*UpdateCoinResponse, error)
	DeleteCoin(ctx context.Context, in *DeleteCoinRequest, opts ...grpc.CallOption) (*DeleteCoinResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateCoin(ctx context.Context, in *CreateCoinRequest, opts ...grpc.CallOption) (*CreateCoinResponse, error) {
	out := new(CreateCoinResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetCoins(ctx context.Context, in *GetCoinsRequest, opts ...grpc.CallOption) (*GetCoinsResponse, error) {
	out := new(GetCoinsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetCoins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppCoins(ctx context.Context, in *GetAppCoinsRequest, opts ...grpc.CallOption) (*GetAppCoinsResponse, error) {
	out := new(GetAppCoinsResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppCoins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateCoin(ctx context.Context, in *UpdateCoinRequest, opts ...grpc.CallOption) (*UpdateCoinResponse, error) {
	out := new(UpdateCoinResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteCoin(ctx context.Context, in *DeleteCoinRequest, opts ...grpc.CallOption) (*DeleteCoinResponse, error) {
	out := new(DeleteCoinResponse)
	err := c.cc.Invoke(ctx, Gateway_DeleteCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateCoin(context.Context, *CreateCoinRequest) (*CreateCoinResponse, error)
	GetCoins(context.Context, *GetCoinsRequest) (*GetCoinsResponse, error)
	GetAppCoins(context.Context, *GetAppCoinsRequest) (*GetAppCoinsResponse, error)
	UpdateCoin(context.Context, *UpdateCoinRequest) (*UpdateCoinResponse, error)
	DeleteCoin(context.Context, *DeleteCoinRequest) (*DeleteCoinResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct{}

func (UnimplementedGatewayServer) CreateCoin(context.Context, *CreateCoinRequest) (*CreateCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoin not implemented")
}

func (UnimplementedGatewayServer) GetCoins(context.Context, *GetCoinsRequest) (*GetCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoins not implemented")
}

func (UnimplementedGatewayServer) GetAppCoins(context.Context, *GetAppCoinsRequest) (*GetAppCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppCoins not implemented")
}

func (UnimplementedGatewayServer) UpdateCoin(context.Context, *UpdateCoinRequest) (*UpdateCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoin not implemented")
}

func (UnimplementedGatewayServer) DeleteCoin(context.Context, *DeleteCoinRequest) (*DeleteCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoin not implemented")
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

func _Gateway_CreateCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateCoin(ctx, req.(*CreateCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetCoins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetCoins(ctx, req.(*GetCoinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppCoinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppCoins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppCoins(ctx, req.(*GetAppCoinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateCoin(ctx, req.(*UpdateCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_DeleteCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteCoin(ctx, req.(*DeleteCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chain.gateway.app.coin.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoin",
			Handler:    _Gateway_CreateCoin_Handler,
		},
		{
			MethodName: "GetCoins",
			Handler:    _Gateway_GetCoins_Handler,
		},
		{
			MethodName: "GetAppCoins",
			Handler:    _Gateway_GetAppCoins_Handler,
		},
		{
			MethodName: "UpdateCoin",
			Handler:    _Gateway_UpdateCoin_Handler,
		},
		{
			MethodName: "DeleteCoin",
			Handler:    _Gateway_DeleteCoin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chain/gateway/v1/app/coin/coin.proto",
}

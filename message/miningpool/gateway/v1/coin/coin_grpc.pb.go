// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: miningpool/gateway/v1/coin/coin.proto

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
	Gateway_AdminCreateCoin_FullMethodName = "/miningpool.gateway.coin.v1.Gateway/AdminCreateCoin"
	Gateway_AdminUpdateCoin_FullMethodName = "/miningpool.gateway.coin.v1.Gateway/AdminUpdateCoin"
	Gateway_AdminGetCoins_FullMethodName   = "/miningpool.gateway.coin.v1.Gateway/AdminGetCoins"
	Gateway_AdminDeleteCoin_FullMethodName = "/miningpool.gateway.coin.v1.Gateway/AdminDeleteCoin"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	AdminCreateCoin(ctx context.Context, in *AdminCreateCoinRequest, opts ...grpc.CallOption) (*AdminCreateCoinResponse, error)
	AdminUpdateCoin(ctx context.Context, in *AdminUpdateCoinRequest, opts ...grpc.CallOption) (*AdminUpdateCoinResponse, error)
	AdminGetCoins(ctx context.Context, in *AdminGetCoinsRequest, opts ...grpc.CallOption) (*AdminGetCoinsResponse, error)
	AdminDeleteCoin(ctx context.Context, in *AdminDeleteCoinRequest, opts ...grpc.CallOption) (*AdminDeleteCoinResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) AdminCreateCoin(ctx context.Context, in *AdminCreateCoinRequest, opts ...grpc.CallOption) (*AdminCreateCoinResponse, error) {
	out := new(AdminCreateCoinResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminCreateCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdateCoin(ctx context.Context, in *AdminUpdateCoinRequest, opts ...grpc.CallOption) (*AdminUpdateCoinResponse, error) {
	out := new(AdminUpdateCoinResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminUpdateCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetCoins(ctx context.Context, in *AdminGetCoinsRequest, opts ...grpc.CallOption) (*AdminGetCoinsResponse, error) {
	out := new(AdminGetCoinsResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminGetCoins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeleteCoin(ctx context.Context, in *AdminDeleteCoinRequest, opts ...grpc.CallOption) (*AdminDeleteCoinResponse, error) {
	out := new(AdminDeleteCoinResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminDeleteCoin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	AdminCreateCoin(context.Context, *AdminCreateCoinRequest) (*AdminCreateCoinResponse, error)
	AdminUpdateCoin(context.Context, *AdminUpdateCoinRequest) (*AdminUpdateCoinResponse, error)
	AdminGetCoins(context.Context, *AdminGetCoinsRequest) (*AdminGetCoinsResponse, error)
	AdminDeleteCoin(context.Context, *AdminDeleteCoinRequest) (*AdminDeleteCoinResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct{}

func (UnimplementedGatewayServer) AdminCreateCoin(context.Context, *AdminCreateCoinRequest) (*AdminCreateCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateCoin not implemented")
}

func (UnimplementedGatewayServer) AdminUpdateCoin(context.Context, *AdminUpdateCoinRequest) (*AdminUpdateCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateCoin not implemented")
}

func (UnimplementedGatewayServer) AdminGetCoins(context.Context, *AdminGetCoinsRequest) (*AdminGetCoinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetCoins not implemented")
}

func (UnimplementedGatewayServer) AdminDeleteCoin(context.Context, *AdminDeleteCoinRequest) (*AdminDeleteCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteCoin not implemented")
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

func _Gateway_AdminCreateCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreateCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminCreateCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreateCoin(ctx, req.(*AdminCreateCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdateCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdateCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminUpdateCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdateCoin(ctx, req.(*AdminUpdateCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetCoinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminGetCoins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetCoins(ctx, req.(*AdminGetCoinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeleteCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteCoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeleteCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminDeleteCoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeleteCoin(ctx, req.(*AdminDeleteCoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miningpool.gateway.coin.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminCreateCoin",
			Handler:    _Gateway_AdminCreateCoin_Handler,
		},
		{
			MethodName: "AdminUpdateCoin",
			Handler:    _Gateway_AdminUpdateCoin_Handler,
		},
		{
			MethodName: "AdminGetCoins",
			Handler:    _Gateway_AdminGetCoins_Handler,
		},
		{
			MethodName: "AdminDeleteCoin",
			Handler:    _Gateway_AdminDeleteCoin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "miningpool/gateway/v1/coin/coin.proto",
}

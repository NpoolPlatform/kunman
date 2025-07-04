// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: good/gateway/v1/device/manufacturer/manufacturer.proto

package manufacturer

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
	Gateway_AdminCreateManufacturer_FullMethodName = "/good.gateway.device.manufacturer.v1.Gateway/AdminCreateManufacturer"
	Gateway_AdminUpdateManufacturer_FullMethodName = "/good.gateway.device.manufacturer.v1.Gateway/AdminUpdateManufacturer"
	Gateway_GetManufacturers_FullMethodName        = "/good.gateway.device.manufacturer.v1.Gateway/GetManufacturers"
	Gateway_AdminDeleteManufacturer_FullMethodName = "/good.gateway.device.manufacturer.v1.Gateway/AdminDeleteManufacturer"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	AdminCreateManufacturer(ctx context.Context, in *AdminCreateManufacturerRequest, opts ...grpc.CallOption) (*AdminCreateManufacturerResponse, error)
	AdminUpdateManufacturer(ctx context.Context, in *AdminUpdateManufacturerRequest, opts ...grpc.CallOption) (*AdminUpdateManufacturerResponse, error)
	GetManufacturers(ctx context.Context, in *GetManufacturersRequest, opts ...grpc.CallOption) (*GetManufacturersResponse, error)
	AdminDeleteManufacturer(ctx context.Context, in *AdminDeleteManufacturerRequest, opts ...grpc.CallOption) (*AdminDeleteManufacturerResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) AdminCreateManufacturer(ctx context.Context, in *AdminCreateManufacturerRequest, opts ...grpc.CallOption) (*AdminCreateManufacturerResponse, error) {
	out := new(AdminCreateManufacturerResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminCreateManufacturer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdateManufacturer(ctx context.Context, in *AdminUpdateManufacturerRequest, opts ...grpc.CallOption) (*AdminUpdateManufacturerResponse, error) {
	out := new(AdminUpdateManufacturerResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminUpdateManufacturer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetManufacturers(ctx context.Context, in *GetManufacturersRequest, opts ...grpc.CallOption) (*GetManufacturersResponse, error) {
	out := new(GetManufacturersResponse)
	err := c.cc.Invoke(ctx, Gateway_GetManufacturers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeleteManufacturer(ctx context.Context, in *AdminDeleteManufacturerRequest, opts ...grpc.CallOption) (*AdminDeleteManufacturerResponse, error) {
	out := new(AdminDeleteManufacturerResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminDeleteManufacturer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	AdminCreateManufacturer(context.Context, *AdminCreateManufacturerRequest) (*AdminCreateManufacturerResponse, error)
	AdminUpdateManufacturer(context.Context, *AdminUpdateManufacturerRequest) (*AdminUpdateManufacturerResponse, error)
	GetManufacturers(context.Context, *GetManufacturersRequest) (*GetManufacturersResponse, error)
	AdminDeleteManufacturer(context.Context, *AdminDeleteManufacturerRequest) (*AdminDeleteManufacturerResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct{}

func (UnimplementedGatewayServer) AdminCreateManufacturer(context.Context, *AdminCreateManufacturerRequest) (*AdminCreateManufacturerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateManufacturer not implemented")
}

func (UnimplementedGatewayServer) AdminUpdateManufacturer(context.Context, *AdminUpdateManufacturerRequest) (*AdminUpdateManufacturerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateManufacturer not implemented")
}

func (UnimplementedGatewayServer) GetManufacturers(context.Context, *GetManufacturersRequest) (*GetManufacturersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManufacturers not implemented")
}

func (UnimplementedGatewayServer) AdminDeleteManufacturer(context.Context, *AdminDeleteManufacturerRequest) (*AdminDeleteManufacturerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteManufacturer not implemented")
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

func _Gateway_AdminCreateManufacturer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateManufacturerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreateManufacturer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminCreateManufacturer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreateManufacturer(ctx, req.(*AdminCreateManufacturerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdateManufacturer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateManufacturerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdateManufacturer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminUpdateManufacturer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdateManufacturer(ctx, req.(*AdminUpdateManufacturerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetManufacturers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManufacturersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetManufacturers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetManufacturers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetManufacturers(ctx, req.(*GetManufacturersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeleteManufacturer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteManufacturerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeleteManufacturer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminDeleteManufacturer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeleteManufacturer(ctx, req.(*AdminDeleteManufacturerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.device.manufacturer.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminCreateManufacturer",
			Handler:    _Gateway_AdminCreateManufacturer_Handler,
		},
		{
			MethodName: "AdminUpdateManufacturer",
			Handler:    _Gateway_AdminUpdateManufacturer_Handler,
		},
		{
			MethodName: "GetManufacturers",
			Handler:    _Gateway_GetManufacturers_Handler,
		},
		{
			MethodName: "AdminDeleteManufacturer",
			Handler:    _Gateway_AdminDeleteManufacturer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "good/gateway/v1/device/manufacturer/manufacturer.proto",
}

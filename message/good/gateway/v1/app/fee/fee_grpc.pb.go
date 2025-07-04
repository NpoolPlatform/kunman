// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: good/gateway/v1/app/fee/fee.proto

package fee

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
	Gateway_GetAppFees_FullMethodName        = "/good.gateway.app.fee.v1.Gateway/GetAppFees"
	Gateway_GetAppFee_FullMethodName         = "/good.gateway.app.fee.v1.Gateway/GetAppFee"
	Gateway_UpdateAppFee_FullMethodName      = "/good.gateway.app.fee.v1.Gateway/UpdateAppFee"
	Gateway_AdminCreateAppFee_FullMethodName = "/good.gateway.app.fee.v1.Gateway/AdminCreateAppFee"
	Gateway_AdminGetAppFees_FullMethodName   = "/good.gateway.app.fee.v1.Gateway/AdminGetAppFees"
	Gateway_AdminUpdateAppFee_FullMethodName = "/good.gateway.app.fee.v1.Gateway/AdminUpdateAppFee"
	Gateway_AdminDeleteAppFee_FullMethodName = "/good.gateway.app.fee.v1.Gateway/AdminDeleteAppFee"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	GetAppFees(ctx context.Context, in *GetAppFeesRequest, opts ...grpc.CallOption) (*GetAppFeesResponse, error)
	GetAppFee(ctx context.Context, in *GetAppFeeRequest, opts ...grpc.CallOption) (*GetAppFeeResponse, error)
	UpdateAppFee(ctx context.Context, in *UpdateAppFeeRequest, opts ...grpc.CallOption) (*UpdateAppFeeResponse, error)
	AdminCreateAppFee(ctx context.Context, in *AdminCreateAppFeeRequest, opts ...grpc.CallOption) (*AdminCreateAppFeeResponse, error)
	AdminGetAppFees(ctx context.Context, in *AdminGetAppFeesRequest, opts ...grpc.CallOption) (*AdminGetAppFeesResponse, error)
	AdminUpdateAppFee(ctx context.Context, in *AdminUpdateAppFeeRequest, opts ...grpc.CallOption) (*AdminUpdateAppFeeResponse, error)
	AdminDeleteAppFee(ctx context.Context, in *AdminDeleteAppFeeRequest, opts ...grpc.CallOption) (*AdminDeleteAppFeeResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetAppFees(ctx context.Context, in *GetAppFeesRequest, opts ...grpc.CallOption) (*GetAppFeesResponse, error) {
	out := new(GetAppFeesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppFees_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppFee(ctx context.Context, in *GetAppFeeRequest, opts ...grpc.CallOption) (*GetAppFeeResponse, error) {
	out := new(GetAppFeeResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppFee(ctx context.Context, in *UpdateAppFeeRequest, opts ...grpc.CallOption) (*UpdateAppFeeResponse, error) {
	out := new(UpdateAppFeeResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateAppFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminCreateAppFee(ctx context.Context, in *AdminCreateAppFeeRequest, opts ...grpc.CallOption) (*AdminCreateAppFeeResponse, error) {
	out := new(AdminCreateAppFeeResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminCreateAppFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminGetAppFees(ctx context.Context, in *AdminGetAppFeesRequest, opts ...grpc.CallOption) (*AdminGetAppFeesResponse, error) {
	out := new(AdminGetAppFeesResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminGetAppFees_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminUpdateAppFee(ctx context.Context, in *AdminUpdateAppFeeRequest, opts ...grpc.CallOption) (*AdminUpdateAppFeeResponse, error) {
	out := new(AdminUpdateAppFeeResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminUpdateAppFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) AdminDeleteAppFee(ctx context.Context, in *AdminDeleteAppFeeRequest, opts ...grpc.CallOption) (*AdminDeleteAppFeeResponse, error) {
	out := new(AdminDeleteAppFeeResponse)
	err := c.cc.Invoke(ctx, Gateway_AdminDeleteAppFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	GetAppFees(context.Context, *GetAppFeesRequest) (*GetAppFeesResponse, error)
	GetAppFee(context.Context, *GetAppFeeRequest) (*GetAppFeeResponse, error)
	UpdateAppFee(context.Context, *UpdateAppFeeRequest) (*UpdateAppFeeResponse, error)
	AdminCreateAppFee(context.Context, *AdminCreateAppFeeRequest) (*AdminCreateAppFeeResponse, error)
	AdminGetAppFees(context.Context, *AdminGetAppFeesRequest) (*AdminGetAppFeesResponse, error)
	AdminUpdateAppFee(context.Context, *AdminUpdateAppFeeRequest) (*AdminUpdateAppFeeResponse, error)
	AdminDeleteAppFee(context.Context, *AdminDeleteAppFeeRequest) (*AdminDeleteAppFeeResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct{}

func (UnimplementedGatewayServer) GetAppFees(context.Context, *GetAppFeesRequest) (*GetAppFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppFees not implemented")
}

func (UnimplementedGatewayServer) GetAppFee(context.Context, *GetAppFeeRequest) (*GetAppFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppFee not implemented")
}

func (UnimplementedGatewayServer) UpdateAppFee(context.Context, *UpdateAppFeeRequest) (*UpdateAppFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppFee not implemented")
}

func (UnimplementedGatewayServer) AdminCreateAppFee(context.Context, *AdminCreateAppFeeRequest) (*AdminCreateAppFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminCreateAppFee not implemented")
}

func (UnimplementedGatewayServer) AdminGetAppFees(context.Context, *AdminGetAppFeesRequest) (*AdminGetAppFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminGetAppFees not implemented")
}

func (UnimplementedGatewayServer) AdminUpdateAppFee(context.Context, *AdminUpdateAppFeeRequest) (*AdminUpdateAppFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUpdateAppFee not implemented")
}

func (UnimplementedGatewayServer) AdminDeleteAppFee(context.Context, *AdminDeleteAppFeeRequest) (*AdminDeleteAppFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDeleteAppFee not implemented")
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

func _Gateway_GetAppFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppFees_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppFees(ctx, req.(*GetAppFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppFee(ctx, req.(*GetAppFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateAppFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppFee(ctx, req.(*UpdateAppFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminCreateAppFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminCreateAppFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminCreateAppFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminCreateAppFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminCreateAppFee(ctx, req.(*AdminCreateAppFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminGetAppFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGetAppFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminGetAppFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminGetAppFees_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminGetAppFees(ctx, req.(*AdminGetAppFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminUpdateAppFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUpdateAppFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminUpdateAppFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminUpdateAppFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminUpdateAppFee(ctx, req.(*AdminUpdateAppFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_AdminDeleteAppFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDeleteAppFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).AdminDeleteAppFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_AdminDeleteAppFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).AdminDeleteAppFee(ctx, req.(*AdminDeleteAppFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "good.gateway.app.fee.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppFees",
			Handler:    _Gateway_GetAppFees_Handler,
		},
		{
			MethodName: "GetAppFee",
			Handler:    _Gateway_GetAppFee_Handler,
		},
		{
			MethodName: "UpdateAppFee",
			Handler:    _Gateway_UpdateAppFee_Handler,
		},
		{
			MethodName: "AdminCreateAppFee",
			Handler:    _Gateway_AdminCreateAppFee_Handler,
		},
		{
			MethodName: "AdminGetAppFees",
			Handler:    _Gateway_AdminGetAppFees_Handler,
		},
		{
			MethodName: "AdminUpdateAppFee",
			Handler:    _Gateway_AdminUpdateAppFee_Handler,
		},
		{
			MethodName: "AdminDeleteAppFee",
			Handler:    _Gateway_AdminDeleteAppFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "good/gateway/v1/app/fee/fee.proto",
}

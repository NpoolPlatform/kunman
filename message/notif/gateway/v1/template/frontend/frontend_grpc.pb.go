// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: notif/gateway/v1/template/frontend/frontend.proto

package frontend

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
	Gateway_CreateFrontendTemplate_FullMethodName    = "/notif.gateway.template.frontend.v1.Gateway/CreateFrontendTemplate"
	Gateway_CreateAppFrontendTemplate_FullMethodName = "/notif.gateway.template.frontend.v1.Gateway/CreateAppFrontendTemplate"
	Gateway_GetFrontendTemplate_FullMethodName       = "/notif.gateway.template.frontend.v1.Gateway/GetFrontendTemplate"
	Gateway_GetFrontendTemplates_FullMethodName      = "/notif.gateway.template.frontend.v1.Gateway/GetFrontendTemplates"
	Gateway_GetAppFrontendTemplates_FullMethodName   = "/notif.gateway.template.frontend.v1.Gateway/GetAppFrontendTemplates"
	Gateway_UpdateFrontendTemplate_FullMethodName    = "/notif.gateway.template.frontend.v1.Gateway/UpdateFrontendTemplate"
	Gateway_UpdateAppFrontendTemplate_FullMethodName = "/notif.gateway.template.frontend.v1.Gateway/UpdateAppFrontendTemplate"
)

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayClient interface {
	CreateFrontendTemplate(ctx context.Context, in *CreateFrontendTemplateRequest, opts ...grpc.CallOption) (*CreateFrontendTemplateResponse, error)
	CreateAppFrontendTemplate(ctx context.Context, in *CreateAppFrontendTemplateRequest, opts ...grpc.CallOption) (*CreateAppFrontendTemplateResponse, error)
	GetFrontendTemplate(ctx context.Context, in *GetFrontendTemplateRequest, opts ...grpc.CallOption) (*GetFrontendTemplateResponse, error)
	GetFrontendTemplates(ctx context.Context, in *GetFrontendTemplatesRequest, opts ...grpc.CallOption) (*GetFrontendTemplatesResponse, error)
	GetAppFrontendTemplates(ctx context.Context, in *GetAppFrontendTemplatesRequest, opts ...grpc.CallOption) (*GetAppFrontendTemplatesResponse, error)
	UpdateFrontendTemplate(ctx context.Context, in *UpdateFrontendTemplateRequest, opts ...grpc.CallOption) (*UpdateFrontendTemplateResponse, error)
	UpdateAppFrontendTemplate(ctx context.Context, in *UpdateAppFrontendTemplateRequest, opts ...grpc.CallOption) (*UpdateAppFrontendTemplateResponse, error)
}

type gatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayClient(cc grpc.ClientConnInterface) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateFrontendTemplate(ctx context.Context, in *CreateFrontendTemplateRequest, opts ...grpc.CallOption) (*CreateFrontendTemplateResponse, error) {
	out := new(CreateFrontendTemplateResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateFrontendTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) CreateAppFrontendTemplate(ctx context.Context, in *CreateAppFrontendTemplateRequest, opts ...grpc.CallOption) (*CreateAppFrontendTemplateResponse, error) {
	out := new(CreateAppFrontendTemplateResponse)
	err := c.cc.Invoke(ctx, Gateway_CreateAppFrontendTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetFrontendTemplate(ctx context.Context, in *GetFrontendTemplateRequest, opts ...grpc.CallOption) (*GetFrontendTemplateResponse, error) {
	out := new(GetFrontendTemplateResponse)
	err := c.cc.Invoke(ctx, Gateway_GetFrontendTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetFrontendTemplates(ctx context.Context, in *GetFrontendTemplatesRequest, opts ...grpc.CallOption) (*GetFrontendTemplatesResponse, error) {
	out := new(GetFrontendTemplatesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetFrontendTemplates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetAppFrontendTemplates(ctx context.Context, in *GetAppFrontendTemplatesRequest, opts ...grpc.CallOption) (*GetAppFrontendTemplatesResponse, error) {
	out := new(GetAppFrontendTemplatesResponse)
	err := c.cc.Invoke(ctx, Gateway_GetAppFrontendTemplates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateFrontendTemplate(ctx context.Context, in *UpdateFrontendTemplateRequest, opts ...grpc.CallOption) (*UpdateFrontendTemplateResponse, error) {
	out := new(UpdateFrontendTemplateResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateFrontendTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) UpdateAppFrontendTemplate(ctx context.Context, in *UpdateAppFrontendTemplateRequest, opts ...grpc.CallOption) (*UpdateAppFrontendTemplateResponse, error) {
	out := new(UpdateAppFrontendTemplateResponse)
	err := c.cc.Invoke(ctx, Gateway_UpdateAppFrontendTemplate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
// All implementations must embed UnimplementedGatewayServer
// for forward compatibility
type GatewayServer interface {
	CreateFrontendTemplate(context.Context, *CreateFrontendTemplateRequest) (*CreateFrontendTemplateResponse, error)
	CreateAppFrontendTemplate(context.Context, *CreateAppFrontendTemplateRequest) (*CreateAppFrontendTemplateResponse, error)
	GetFrontendTemplate(context.Context, *GetFrontendTemplateRequest) (*GetFrontendTemplateResponse, error)
	GetFrontendTemplates(context.Context, *GetFrontendTemplatesRequest) (*GetFrontendTemplatesResponse, error)
	GetAppFrontendTemplates(context.Context, *GetAppFrontendTemplatesRequest) (*GetAppFrontendTemplatesResponse, error)
	UpdateFrontendTemplate(context.Context, *UpdateFrontendTemplateRequest) (*UpdateFrontendTemplateResponse, error)
	UpdateAppFrontendTemplate(context.Context, *UpdateAppFrontendTemplateRequest) (*UpdateAppFrontendTemplateResponse, error)
	mustEmbedUnimplementedGatewayServer()
}

// UnimplementedGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct{}

func (UnimplementedGatewayServer) CreateFrontendTemplate(context.Context, *CreateFrontendTemplateRequest) (*CreateFrontendTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFrontendTemplate not implemented")
}

func (UnimplementedGatewayServer) CreateAppFrontendTemplate(context.Context, *CreateAppFrontendTemplateRequest) (*CreateAppFrontendTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAppFrontendTemplate not implemented")
}

func (UnimplementedGatewayServer) GetFrontendTemplate(context.Context, *GetFrontendTemplateRequest) (*GetFrontendTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFrontendTemplate not implemented")
}

func (UnimplementedGatewayServer) GetFrontendTemplates(context.Context, *GetFrontendTemplatesRequest) (*GetFrontendTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFrontendTemplates not implemented")
}

func (UnimplementedGatewayServer) GetAppFrontendTemplates(context.Context, *GetAppFrontendTemplatesRequest) (*GetAppFrontendTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppFrontendTemplates not implemented")
}

func (UnimplementedGatewayServer) UpdateFrontendTemplate(context.Context, *UpdateFrontendTemplateRequest) (*UpdateFrontendTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFrontendTemplate not implemented")
}

func (UnimplementedGatewayServer) UpdateAppFrontendTemplate(context.Context, *UpdateAppFrontendTemplateRequest) (*UpdateAppFrontendTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppFrontendTemplate not implemented")
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

func _Gateway_CreateFrontendTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFrontendTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateFrontendTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateFrontendTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateFrontendTemplate(ctx, req.(*CreateFrontendTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_CreateAppFrontendTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppFrontendTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateAppFrontendTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_CreateAppFrontendTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateAppFrontendTemplate(ctx, req.(*CreateAppFrontendTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetFrontendTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFrontendTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetFrontendTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetFrontendTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetFrontendTemplate(ctx, req.(*GetFrontendTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetFrontendTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFrontendTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetFrontendTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetFrontendTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetFrontendTemplates(ctx, req.(*GetFrontendTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetAppFrontendTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppFrontendTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetAppFrontendTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_GetAppFrontendTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetAppFrontendTemplates(ctx, req.(*GetAppFrontendTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateFrontendTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFrontendTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateFrontendTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateFrontendTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateFrontendTemplate(ctx, req.(*UpdateFrontendTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_UpdateAppFrontendTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppFrontendTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).UpdateAppFrontendTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gateway_UpdateAppFrontendTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).UpdateAppFrontendTemplate(ctx, req.(*UpdateAppFrontendTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gateway_ServiceDesc is the grpc.ServiceDesc for Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notif.gateway.template.frontend.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFrontendTemplate",
			Handler:    _Gateway_CreateFrontendTemplate_Handler,
		},
		{
			MethodName: "CreateAppFrontendTemplate",
			Handler:    _Gateway_CreateAppFrontendTemplate_Handler,
		},
		{
			MethodName: "GetFrontendTemplate",
			Handler:    _Gateway_GetFrontendTemplate_Handler,
		},
		{
			MethodName: "GetFrontendTemplates",
			Handler:    _Gateway_GetFrontendTemplates_Handler,
		},
		{
			MethodName: "GetAppFrontendTemplates",
			Handler:    _Gateway_GetAppFrontendTemplates_Handler,
		},
		{
			MethodName: "UpdateFrontendTemplate",
			Handler:    _Gateway_UpdateFrontendTemplate_Handler,
		},
		{
			MethodName: "UpdateAppFrontendTemplate",
			Handler:    _Gateway_UpdateAppFrontendTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notif/gateway/v1/template/frontend/frontend.proto",
}

package api

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/sphinx/proxy"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// Server ..
type Server struct {
	proxy.UnimplementedSphinxProxyServer
}

func Register(server grpc.ServiceRegistrar) {
	proxy.RegisterSphinxProxyServer(server, &Server{})
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return proxy.RegisterSphinxProxyHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

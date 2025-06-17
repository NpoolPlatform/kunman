package appconfig

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/order/gateway/v1/app/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	config.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	config.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return config.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

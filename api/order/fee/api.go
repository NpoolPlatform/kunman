package feeorder

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/order/gateway/v1/fee"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fee.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	fee.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return fee.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

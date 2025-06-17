package order

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/order/gateway/v1/order"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	order.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	order.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return order.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

package subscription

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	subscription.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	subscription.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return subscription.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

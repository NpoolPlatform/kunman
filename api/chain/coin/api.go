package coin

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/chain/gateway/v1/coin"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	coin.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	coin.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return coin.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

package feed

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/chain/gateway/v1/coin/currency/feed"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	feed.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	feed.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return feed.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

package coinfiat

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/chain/gateway/v1/coin/fiat"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fiat.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	fiat.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return fiat.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

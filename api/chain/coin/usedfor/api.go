package coinusedfor

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/chain/gateway/v1/coin/usedfor"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	usedfor.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	usedfor.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return usedfor.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

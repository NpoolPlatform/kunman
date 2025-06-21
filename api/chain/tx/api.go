package tx

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/chain/gateway/v1/tx"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	tx.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	tx.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return tx.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

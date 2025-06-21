package description

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/coin/description"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	description.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	description.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return description.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

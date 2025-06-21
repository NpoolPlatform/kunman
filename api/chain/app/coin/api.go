package appcoin

import (
	"context"

	appcoin1 "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/coin"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appcoin1.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	appcoin1.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return appcoin1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

package api

import (
	npool "github.com/NpoolPlatform/kunman/message/agi/gateway/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}

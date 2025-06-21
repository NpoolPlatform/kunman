package appfiat

import (
	"context"

	appfiat1 "github.com/NpoolPlatform/kunman/message/chain/gateway/v1/app/fiat"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appfiat1.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	appfiat1.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return appfiat1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

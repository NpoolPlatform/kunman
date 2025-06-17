package topmostgood

import (
	"context"

	topmostgood "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost/good"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	topmostgood.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	topmostgood.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := topmostgood.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

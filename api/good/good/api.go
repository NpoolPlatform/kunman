package good

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/good/gateway/v1/good"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	good.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	good.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := good.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

package default1

import (
	"context"

	default1 "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/default"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	default1.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	default1.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := default1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

package malfunction

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/malfunction"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	malfunction.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	malfunction.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := malfunction.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

package topmost

import (
	"context"

	topmost "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/topmost"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	topmost.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	topmost.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := topmost.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

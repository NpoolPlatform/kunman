package label

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/label"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	label.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	label.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := label.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

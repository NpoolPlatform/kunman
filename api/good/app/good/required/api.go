package required

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/required"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	required.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	required.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := required.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

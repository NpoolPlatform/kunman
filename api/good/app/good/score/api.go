package score

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/score"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	score.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	score.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := score.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

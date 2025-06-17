package oneshot

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription/oneshot"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	oneshot.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	oneshot.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := oneshot.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

package change

import (
	"context"

	change "github.com/NpoolPlatform/kunman/message/billing/gateway/v1/user/subscription/change"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	change.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	change.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := change.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

package addon

import (
	"context"

	addon "github.com/NpoolPlatform/kunman/message/billing/gateway/v1/addon"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	addon.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	addon.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := addon.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

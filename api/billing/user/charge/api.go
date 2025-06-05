package charge

import (
	"context"

	charge "github.com/NpoolPlatform/kunman/message/billing/gw/v1/user/charge"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	charge.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	charge.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := charge.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

package exchange

import (
	"context"

	exchange "github.com/NpoolPlatform/kunman/message/billing/gw/v1/credit/exchange"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	exchange.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	exchange.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := exchange.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

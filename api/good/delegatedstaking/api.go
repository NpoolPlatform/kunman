package delegatedstaking

import (
	"context"

	delegatedstaking "github.com/NpoolPlatform/kunman/message/good/gateway/v1/delegatedstaking"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	delegatedstaking.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	delegatedstaking.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := delegatedstaking.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

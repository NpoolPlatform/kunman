package goodcoin

import (
	"context"

	goodcoin "github.com/NpoolPlatform/kunman/message/good/gateway/v1/good/coin"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	goodcoin.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	goodcoin.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := goodcoin.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

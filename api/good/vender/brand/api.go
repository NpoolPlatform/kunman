package brand

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/good/gateway/v1/vender/brand"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	brand.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	brand.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := brand.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

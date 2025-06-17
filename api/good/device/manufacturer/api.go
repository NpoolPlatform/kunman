package manufacturer

import (
	"context"

	manufacturer "github.com/NpoolPlatform/kunman/message/good/gateway/v1/device/manufacturer"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	manufacturer.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	manufacturer.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := manufacturer.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

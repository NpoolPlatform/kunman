package compensate

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental/compensate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	compensate.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	compensate.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return compensate.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

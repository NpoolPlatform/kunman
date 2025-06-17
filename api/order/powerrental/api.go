package powerrental

import (
	"context"

	"github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	powerrental.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	powerrental.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return powerrental.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

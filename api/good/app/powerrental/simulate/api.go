package simulate

import (
	"context"

	simulate1 "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/powerrental/simulate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	simulate1.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	simulate1.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := simulate1.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

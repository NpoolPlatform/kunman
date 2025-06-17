package displaycolor

import (
	"context"

	displaycolor "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/color"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	displaycolor.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	displaycolor.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := displaycolor.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

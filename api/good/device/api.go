package devicetype

import (
	"context"

	devicetype "github.com/NpoolPlatform/kunman/message/good/gateway/v1/device"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	devicetype.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	devicetype.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := devicetype.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

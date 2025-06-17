package displayname

import (
	"context"

	displayname "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/good/display/name"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	displayname.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	displayname.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := displayname.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

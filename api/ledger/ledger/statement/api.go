package statement

import (
	"context"

	statement "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1/ledger/statement"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	statement.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	statement.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := statement.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

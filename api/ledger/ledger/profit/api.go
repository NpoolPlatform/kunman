package profit

import (
	"context"

	profit "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1/ledger/profit"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	profit.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	profit.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := profit.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

package deposit

import (
	"context"

	deposit "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1/ledger/deposit"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	deposit.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	deposit.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := deposit.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

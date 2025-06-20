package ledger

import (
	"context"

	ledger "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1/ledger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ledger.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	ledger.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := ledger.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

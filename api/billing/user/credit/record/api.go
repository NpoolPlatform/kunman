package record

import (
	"context"

	record "github.com/NpoolPlatform/kunman/message/billing/gw/v1/user/credit/record"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	record.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	record.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := record.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

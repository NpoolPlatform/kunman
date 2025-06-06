package api

import (
	agi "github.com/NpoolPlatform/kunman/api/agi"

	npool "github.com/NpoolPlatform/kunman/message/billing/gateway/v1"
)

type Server struct {
	npool.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	agi.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := agi.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}

	return nil
}

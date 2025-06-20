package api

import (
	"context"

	agi "github.com/NpoolPlatform/kunman/api/agi"
	good "github.com/NpoolPlatform/kunman/api/good"
	ledger "github.com/NpoolPlatform/kunman/api/ledger"
	order "github.com/NpoolPlatform/kunman/api/order"

	kunman "github.com/NpoolPlatform/kunman/message/version"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	kunman.UnimplementedKunmanServer
}

func Register(server grpc.ServiceRegistrar) {
	kunman.RegisterKunmanServer(server, &Server{})
	agi.Register(server)
	good.Register(server)
	order.Register(server)
	ledger.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := kunman.RegisterKunmanHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := agi.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := good.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := order.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := ledger.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}

	return nil
}

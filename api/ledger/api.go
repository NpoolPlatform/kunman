package api

import (
	ledger "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1"

	ledger1 "github.com/NpoolPlatform/kunman/api/ledger/ledger"
	deposit "github.com/NpoolPlatform/kunman/api/ledger/ledger/deposit"
	profit "github.com/NpoolPlatform/kunman/api/ledger/ledger/profit"
	statement "github.com/NpoolPlatform/kunman/api/ledger/ledger/statement"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ledger.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	ledger.RegisterGatewayServer(server, &Server{})
	ledger1.Register(server)
	deposit.Register(server)
	profit.Register(server)
	statement.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := ledger1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := deposit.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := profit.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := statement.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

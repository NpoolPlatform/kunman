package api

import (
	"context"

	addon "github.com/NpoolPlatform/kunman/api/billing/addon"
	exchange "github.com/NpoolPlatform/kunman/api/billing/credit/exchange"
	subscription "github.com/NpoolPlatform/kunman/api/billing/subscription"
	usercharge "github.com/NpoolPlatform/kunman/api/billing/user/charge"
	record "github.com/NpoolPlatform/kunman/api/billing/user/credit/record"
	usersubscription "github.com/NpoolPlatform/kunman/api/billing/user/subscription"
	usersubscriptionchange "github.com/NpoolPlatform/kunman/api/billing/user/subscription/change"
	npool "github.com/NpoolPlatform/kunman/message/billing/gw/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterGatewayServer(server, &Server{})
	addon.Register(server)
	subscription.Register(server)
	exchange.Register(server)
	usersubscription.Register(server)
	record.Register(server)
	usersubscriptionchange.Register(server)
	usercharge.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := npool.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := addon.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := subscription.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := exchange.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := usersubscription.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := record.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := usersubscriptionchange.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := usercharge.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}

	return nil
}

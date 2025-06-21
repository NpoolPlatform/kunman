package api

import (
	"context"

	chaingw "github.com/NpoolPlatform/kunman/message/chain/gateway/v1"

	appcoin "github.com/NpoolPlatform/kunman/api/chain/app/coin"
	"github.com/NpoolPlatform/kunman/api/chain/app/coin/description"
	"github.com/NpoolPlatform/kunman/api/chain/chain"
	"github.com/NpoolPlatform/kunman/api/chain/coin"
	coincurrency "github.com/NpoolPlatform/kunman/api/chain/coin/currency"
	coincurrencyfeed "github.com/NpoolPlatform/kunman/api/chain/coin/currency/feed"
	coincurrencyhis "github.com/NpoolPlatform/kunman/api/chain/coin/currency/history"
	coinfiat "github.com/NpoolPlatform/kunman/api/chain/coin/fiat"
	coinfiatcurrencyhis "github.com/NpoolPlatform/kunman/api/chain/coin/fiat/currency/history"
	coinusedfor "github.com/NpoolPlatform/kunman/api/chain/coin/usedfor"
	"github.com/NpoolPlatform/kunman/api/chain/fiat"
	fiatcurrency "github.com/NpoolPlatform/kunman/api/chain/fiat/currency"
	fiatcurrencyfeed "github.com/NpoolPlatform/kunman/api/chain/fiat/currency/feed"
	fiatcurrencyhis "github.com/NpoolPlatform/kunman/api/chain/fiat/currency/history"
	"github.com/NpoolPlatform/kunman/api/chain/tx"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	chaingw.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	chaingw.RegisterGatewayServer(server, &Server{})
	coin.Register(server)
	coinfiat.Register(server)
	coincurrency.Register(server)
	coincurrencyfeed.Register(server)
	coincurrencyhis.Register(server)
	coinfiatcurrencyhis.Register(server)
	appcoin.Register(server)
	tx.Register(server)
	description.Register(server)
	fiat.Register(server)
	fiatcurrency.Register(server)
	fiatcurrencyfeed.Register(server)
	fiatcurrencyhis.Register(server)
	chain.Register(server)
	coinusedfor.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := chaingw.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := coin.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := coinfiat.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := coincurrency.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := coincurrencyfeed.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := coincurrencyhis.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := coinfiatcurrencyhis.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appcoin.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := tx.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := description.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := fiat.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := fiatcurrency.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := fiatcurrencyfeed.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := fiatcurrencyhis.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := chain.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := coinusedfor.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

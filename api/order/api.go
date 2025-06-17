package api

import (
	appconfig1 "github.com/NpoolPlatform/kunman/api/order/app/config"
	compensate1 "github.com/NpoolPlatform/kunman/api/order/compensate"
	feeorder1 "github.com/NpoolPlatform/kunman/api/order/fee"
	order1 "github.com/NpoolPlatform/kunman/api/order/order"
	ordercoupon1 "github.com/NpoolPlatform/kunman/api/order/order/coupon"
	outofgas1 "github.com/NpoolPlatform/kunman/api/order/outofgas"
	powerrentalorder1 "github.com/NpoolPlatform/kunman/api/order/powerrental"
	powerrentalcompensate1 "github.com/NpoolPlatform/kunman/api/order/powerrental/compensate"
	powerrentaloutofgas1 "github.com/NpoolPlatform/kunman/api/order/powerrental/outofgas"
	subscriptionorder1 "github.com/NpoolPlatform/kunman/api/order/subscription"
	order "github.com/NpoolPlatform/kunman/message/order/gateway/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	order.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	order.RegisterGatewayServer(server, &Server{})
	feeorder1.Register(server)
	order1.Register(server)
	ordercoupon1.Register(server)
	compensate1.Register(server)
	powerrentalcompensate1.Register(server)
	powerrentaloutofgas1.Register(server)
	outofgas1.Register(server)
	powerrentalorder1.Register(server)
	subscriptionorder1.Register(server)
	appconfig1.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := order1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := ordercoupon1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := feeorder1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := compensate1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := powerrentalcompensate1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := powerrentaloutofgas1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := outofgas1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := powerrentalorder1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := subscriptionorder1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appconfig1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}

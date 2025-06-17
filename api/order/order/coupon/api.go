package ordercoupon

import (
	"context"

	ordercoupon "github.com/NpoolPlatform/kunman/message/order/gateway/v1/order/coupon"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ordercoupon.UnimplementedGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	ordercoupon.RegisterGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return ordercoupon.RegisterGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

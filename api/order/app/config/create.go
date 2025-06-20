//nolint:dupl
package appconfig

import (
	"context"

	config1 "github.com/NpoolPlatform/kunman/gateway/order/app/config"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/app/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) CreateAppConfig(ctx context.Context, in *npool.CreateAppConfigRequest) (*npool.CreateAppConfigResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithAppID(&in.AppID, true),
		config1.WithEnableSimulateOrder(in.EnableSimulateOrder, false),
		config1.WithSimulateOrderCouponMode(in.SimulateOrderCouponMode, false),
		config1.WithSimulateOrderCouponProbability(in.SimulateOrderCouponProbability, false),
		config1.WithSimulateOrderCashableProfitProbability(in.SimulateOrderCashableProfitProbability, false),
		config1.WithMaxUnpaidOrders(in.MaxUnpaidOrders, false),
		config1.WithMaxTypedCouponsPerOrder(in.MaxTypedCouponsPerOrder, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAppConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateAppConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAppConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppConfigResponse{
		Info: info,
	}, nil
}

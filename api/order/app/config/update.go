//nolint:dupl
package appconfig

import (
	"context"

	config1 "github.com/NpoolPlatform/kunman/gateway/order/app/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/app/config"
)

func (s *Server) UpdateAppConfig(ctx context.Context, in *npool.UpdateAppConfigRequest) (*npool.UpdateAppConfigResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithID(&in.ID, true),
		config1.WithEntID(&in.EntID, true),
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
			"UpdateAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateAppConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAppConfigResponse{
		Info: info,
	}, nil
}

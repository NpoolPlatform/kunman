//nolint:dupl
package powerrental

import (
	"context"

	ordercommon "github.com/NpoolPlatform/kunman/api/order/order/common"
	powerrental1 "github.com/NpoolPlatform/kunman/gateway/order/powerrental"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) AdminCreatePowerRentalOrder(ctx context.Context, in *npool.AdminCreatePowerRentalOrderRequest) (*npool.AdminCreatePowerRentalOrderResponse, error) {
	if err := ordercommon.ValidateAdminCreateOrderType(in.GetOrderType()); err != nil {
		logger.Sugar().Errorw(
			"AdminCreatePowerRentalOrder",
			"In", in,
		)
		return &npool.AdminCreatePowerRentalOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithAppID(&in.TargetAppID, true),
		powerrental1.WithUserID(&in.TargetUserID, true),
		powerrental1.WithAppGoodID(&in.AppGoodID, true),
		powerrental1.WithOrderType(&in.OrderType, true),
		powerrental1.WithCreateMethod(func() *types.OrderCreateMethod { e := types.OrderCreateMethod_OrderCreatedByAdmin; return &e }(), true),
		powerrental1.WithDurationSeconds(in.DurationSeconds, false),
		powerrental1.WithUnits(in.Units, true),
		powerrental1.WithAppSpotUnits(in.AppSpotUnits, false),
		powerrental1.WithAppGoodStockID(&in.AppGoodStockID, true),
		powerrental1.WithInvestmentType(&in.InvestmentType, true),
		powerrental1.WithOrderBenefitReqs(in.OrderBenefitAccounts),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreatePowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreatePowerRentalOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreatePowerRentalOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreatePowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreatePowerRentalOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminCreatePowerRentalOrderResponse{
		Info: info,
	}, nil
}

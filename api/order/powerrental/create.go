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

func (s *Server) CreatePowerRentalOrder(ctx context.Context, in *npool.CreatePowerRentalOrderRequest) (*npool.CreatePowerRentalOrderResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithAppID(&in.AppID, true),
		powerrental1.WithUserID(&in.UserID, true),
		powerrental1.WithAppGoodID(&in.AppGoodID, true),
		powerrental1.WithOrderType(func() *types.OrderType { e := types.OrderType_Normal; return &e }(), true),
		powerrental1.WithCreateMethod(func() *types.OrderCreateMethod { e := types.OrderCreateMethod_OrderCreatedByPurchase; return &e }(), true),
		powerrental1.WithDurationSeconds(in.DurationSeconds, false),
		powerrental1.WithUnits(in.Units, false),
		powerrental1.WithAppSpotUnits(in.AppSpotUnits, false),
		powerrental1.WithPaymentBalances(in.Balances, true),
		powerrental1.WithPaymentTransferCoinTypeID(in.PaymentTransferCoinTypeID, false),
		powerrental1.WithCouponIDs(in.CouponIDs, true),
		powerrental1.WithFeeAppGoodIDs(in.FeeAppGoodIDs, true),
		powerrental1.WithFeeDurationSeconds(in.FeeDurationSeconds, false),
		powerrental1.WithFeeAutoDeduction(in.FeeAutoDeduction, false),
		powerrental1.WithAppGoodStockID(in.AppGoodStockID, false),
		powerrental1.WithInvestmentType(&in.InvestmentType, true),
		powerrental1.WithSimulate(in.Simulate, false),
		powerrental1.WithOrderBenefitReqs(in.OrderBenefitAccounts),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePowerRentalOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreatePowerRentalOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePowerRentalOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreatePowerRentalOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateUserPowerRentalOrder(ctx context.Context, in *npool.CreateUserPowerRentalOrderRequest) (*npool.CreateUserPowerRentalOrderResponse, error) {
	if err := ordercommon.ValidateAdminCreateOrderType(in.GetOrderType()); err != nil {
		logger.Sugar().Errorw(
			"CreateUserPowerRentalOrder",
			"In", in,
		)
		return &npool.CreateUserPowerRentalOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithAppID(&in.AppID, true),
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
			"CreateUserPowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserPowerRentalOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreatePowerRentalOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserPowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserPowerRentalOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateUserPowerRentalOrderResponse{
		Info: info,
	}, nil
}

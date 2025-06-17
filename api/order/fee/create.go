//nolint:dupl
package feeorder

import (
	"context"

	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/fee"
	ordercommon "github.com/NpoolPlatform/kunman/api/order/order/common"
	feeorder1 "github.com/NpoolPlatform/kunman/gateway/order/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) CreateFeeOrder(ctx context.Context, in *npool.CreateFeeOrderRequest) (*npool.CreateFeeOrderResponse, error) {
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithAppID(&in.AppID, true),
		feeorder1.WithUserID(&in.UserID, true),
		feeorder1.WithAppGoodID(&in.AppGoodID, true),
		feeorder1.WithParentOrderID(&in.ParentOrderID, true),
		feeorder1.WithOrderType(func() *types.OrderType { e := types.OrderType_Normal; return &e }(), true),
		feeorder1.WithCreateMethod(func() *types.OrderCreateMethod { e := types.OrderCreateMethod_OrderCreatedByPurchase; return &e }(), true),
		feeorder1.WithDurationSeconds(&in.DurationSeconds, true),
		feeorder1.WithPaymentBalances(in.Balances, true),
		feeorder1.WithPaymentTransferCoinTypeID(in.PaymentTransferCoinTypeID, false),
		feeorder1.WithCouponIDs(in.CouponIDs, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFeeOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateFeeOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFeeOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFeeOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateUserFeeOrder(ctx context.Context, in *npool.CreateUserFeeOrderRequest) (*npool.CreateUserFeeOrderResponse, error) {
	if err := ordercommon.ValidateAdminCreateOrderType(in.GetOrderType()); err != nil {
		logger.Sugar().Errorw(
			"CreateUserFeeOrder",
			"In", in,
		)
		return &npool.CreateUserFeeOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithAppID(&in.AppID, true),
		feeorder1.WithUserID(&in.TargetUserID, true),
		feeorder1.WithAppGoodID(&in.AppGoodID, true),
		feeorder1.WithParentOrderID(&in.ParentOrderID, true),
		feeorder1.WithOrderType(&in.OrderType, true),
		feeorder1.WithCreateMethod(func() *types.OrderCreateMethod { e := types.OrderCreateMethod_OrderCreatedByAdmin; return &e }(), true),
		feeorder1.WithDurationSeconds(&in.DurationSeconds, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserFeeOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateFeeOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUserFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserFeeOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateUserFeeOrderResponse{
		Info: info,
	}, nil
}

package feeorder

import (
	"context"

	feeorder1 "github.com/NpoolPlatform/kunman/gateway/order/fee"
	types "github.com/NpoolPlatform/kunman/message/basetypes/order/v1"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) CreateFeeOrders(ctx context.Context, in *npool.CreateFeeOrdersRequest) (*npool.CreateFeeOrdersResponse, error) {
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithAppID(&in.AppID, true),
		feeorder1.WithUserID(&in.UserID, true),
		feeorder1.WithAppGoodIDs(in.AppGoodIDs, true),
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
			"CreateFeeOrders",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFeeOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateFeeOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFeeOrders",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFeeOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFeeOrdersResponse{
		Infos: infos,
	}, nil
}

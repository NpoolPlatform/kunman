//nolint:dupl
package feeorder

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/fee"
	feeorder1 "github.com/NpoolPlatform/kunman/gateway/order/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) UpdateFeeOrder(ctx context.Context, in *npool.UpdateFeeOrderRequest) (*npool.UpdateFeeOrderResponse, error) {
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithID(&in.ID, true),
		feeorder1.WithEntID(&in.EntID, true),
		feeorder1.WithAppID(&in.AppID, true),
		feeorder1.WithUserID(&in.UserID, true),
		feeorder1.WithOrderID(&in.OrderID, true),
		feeorder1.WithPaymentBalances(in.Balances, false),
		feeorder1.WithPaymentTransferCoinTypeID(in.PaymentTransferCoinTypeID, false),
		feeorder1.WithUserSetPaid(in.Paid, false),
		feeorder1.WithUserSetCanceled(in.Canceled, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFeeOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateFeeOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFeeOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateFeeOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateUserFeeOrder(ctx context.Context, in *npool.UpdateUserFeeOrderRequest) (*npool.UpdateUserFeeOrderResponse, error) {
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithID(&in.ID, true),
		feeorder1.WithEntID(&in.EntID, true),
		feeorder1.WithAppID(&in.AppID, true),
		feeorder1.WithUserID(&in.TargetUserID, true),
		feeorder1.WithOrderID(&in.OrderID, true),
		feeorder1.WithAdminSetCanceled(in.Canceled, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUserFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserFeeOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateFeeOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUserFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserFeeOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateUserFeeOrderResponse{
		Info: info,
	}, nil
}

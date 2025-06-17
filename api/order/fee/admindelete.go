package feeorder

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/fee"
	feeorder1 "github.com/NpoolPlatform/kunman/gateway/order/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) AdminDeleteFeeOrder(ctx context.Context, in *npool.AdminDeleteFeeOrderRequest) (*npool.AdminDeleteFeeOrderResponse, error) {
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithID(&in.ID, true),
		feeorder1.WithEntID(&in.EntID, true),
		feeorder1.WithAppID(&in.TargetAppID, true),
		feeorder1.WithUserID(&in.TargetUserID, true),
		feeorder1.WithOrderID(&in.OrderID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteFeeOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteFeeOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteFeeOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminDeleteFeeOrderResponse{
		Info: info,
	}, nil
}

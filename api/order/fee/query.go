package feeorder

import (
	"context"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/fee"
	feeorder1 "github.com/NpoolPlatform/kunman/gateway/order/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) GetFeeOrder(ctx context.Context, in *npool.GetFeeOrderRequest) (*npool.GetFeeOrderResponse, error) {
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithAppID(&in.AppID, true),
		feeorder1.WithUserID(&in.UserID, true),
		feeorder1.WithOrderID(&in.OrderID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeeOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.GetFeeOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeeOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFeeOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) GetFeeOrders(ctx context.Context, in *npool.GetFeeOrdersRequest) (*npool.GetFeeOrdersResponse, error) {
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithAppID(&in.AppID, true),
		feeorder1.WithUserID(in.TargetUserID, false),
		feeorder1.WithAppGoodID(in.AppGoodID, false),
		feeorder1.WithOffset(in.Offset),
		feeorder1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFeeOrders",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeeOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetFeeOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFeeOrders",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeeOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFeeOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetMyFeeOrders(ctx context.Context, in *npool.GetMyFeeOrdersRequest) (*npool.GetMyFeeOrdersResponse, error) {
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithAppID(&in.AppID, true),
		feeorder1.WithUserID(&in.UserID, true),
		feeorder1.WithAppGoodID(in.AppGoodID, false),
		feeorder1.WithOffset(in.Offset),
		feeorder1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMyFeeOrders",
			"In", in,
			"Error", err,
		)
		return &npool.GetMyFeeOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetFeeOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMyFeeOrders",
			"In", in,
			"Error", err,
		)
		return &npool.GetMyFeeOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetMyFeeOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}

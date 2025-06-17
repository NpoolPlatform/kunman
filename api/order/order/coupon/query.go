package ordercoupon

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/order/coupon"
	ordercoupon1 "github.com/NpoolPlatform/kunman/gateway/order/order/coupon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOrderCoupons(ctx context.Context, in *npool.GetOrderCouponsRequest) (*npool.GetOrderCouponsResponse, error) {
	handler, err := ordercoupon1.NewHandler(
		ctx,
		ordercoupon1.WithAppID(&in.AppID, true),
		ordercoupon1.WithUserID(in.TargetUserID, false),
		ordercoupon1.WithOffset(in.GetOffset()),
		ordercoupon1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderCoupons",
			"In", in,
			"Error", err,
		)
		return &npool.GetOrderCouponsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetOrderCoupons(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderCoupons",
			"In", in,
			"Error", err,
		)
		return &npool.GetOrderCouponsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderCouponsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetMyOrderCoupons(ctx context.Context, in *npool.GetMyOrderCouponsRequest) (*npool.GetMyOrderCouponsResponse, error) {
	handler, err := ordercoupon1.NewHandler(
		ctx,
		ordercoupon1.WithAppID(&in.AppID, true),
		ordercoupon1.WithUserID(&in.UserID, true),
		ordercoupon1.WithOffset(in.GetOffset()),
		ordercoupon1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMyOrderCoupons",
			"In", in,
			"Error", err,
		)
		return &npool.GetMyOrderCouponsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetOrderCoupons(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMyOrderCoupons",
			"In", in,
			"Error", err,
		)
		return &npool.GetMyOrderCouponsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetMyOrderCouponsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

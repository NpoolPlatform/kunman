package ordercoupon

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/logger"

	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/order/coupon"
	ordercoupon1 "github.com/NpoolPlatform/kunman/gateway/order/order/coupon"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AdminGetOrderCoupons(ctx context.Context, in *npool.AdminGetOrderCouponsRequest) (*npool.AdminGetOrderCouponsResponse, error) {
	handler, err := ordercoupon1.NewHandler(
		ctx,
		ordercoupon1.WithAppID(in.TargetAppID, false),
		ordercoupon1.WithOffset(in.GetOffset()),
		ordercoupon1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetOrderCoupons",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetOrderCouponsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetOrderCoupons(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetOrderCoupons",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetOrderCouponsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminGetOrderCouponsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

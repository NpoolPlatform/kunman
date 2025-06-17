package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/order/subscription"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) AdminGetSubscriptionOrders(ctx context.Context, in *npool.AdminGetSubscriptionOrdersRequest) (*npool.AdminGetSubscriptionOrdersResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(in.TargetAppID, false),
		subscription1.WithGoodID(in.GoodID, false),
		subscription1.WithOffset(in.Offset),
		subscription1.WithLimit(in.Limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetSubscriptionOrders",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetSubscriptionOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetSubscriptionOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminGetSubscriptionOrders",
			"In", in,
			"Error", err,
		)
		return &npool.AdminGetSubscriptionOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminGetSubscriptionOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}

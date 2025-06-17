package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/order/subscription"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) AdminDeleteSubscriptionOrder(ctx context.Context, in *npool.AdminDeleteSubscriptionOrderRequest) (*npool.AdminDeleteSubscriptionOrderResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(&in.ID, true),
		subscription1.WithEntID(&in.EntID, true),
		subscription1.WithAppID(&in.TargetAppID, true),
		subscription1.WithUserID(&in.TargetUserID, true),
		subscription1.WithOrderID(&in.OrderID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteSubscriptionOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteSubscriptionOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteSubscriptionOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminDeleteSubscriptionOrderResponse{
		Info: info,
	}, nil
}

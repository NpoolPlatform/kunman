//nolint:dupl
package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/order/subscription"
	npool "github.com/NpoolPlatform/kunman/message/order/gateway/v1/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
)

func (s *Server) AdminUpdateSubscriptionOrder(ctx context.Context, in *npool.AdminUpdateSubscriptionOrderRequest) (*npool.AdminUpdateSubscriptionOrderResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(&in.ID, true),
		subscription1.WithEntID(&in.EntID, true),
		subscription1.WithAppID(&in.TargetAppID, true),
		subscription1.WithUserID(&in.TargetUserID, true),
		subscription1.WithOrderID(&in.OrderID, true),
		subscription1.WithAdminSetCanceled(in.Canceled, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateSubscriptionOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateSubscriptionOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateSubscriptionOrder",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateSubscriptionOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.AdminUpdateSubscriptionOrderResponse{
		Info: info,
	}, nil
}

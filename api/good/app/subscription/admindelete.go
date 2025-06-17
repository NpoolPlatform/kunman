package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/good/app/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription"
)

func (s *Server) AdminDeleteAppSubscription(ctx context.Context, in *npool.AdminDeleteAppSubscriptionRequest) (*npool.AdminDeleteAppSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(&in.ID, true),
		subscription1.WithEntID(&in.EntID, true),
		subscription1.WithAppID(&in.TargetAppID, true),
		subscription1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteSubscription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminDeleteAppSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminDeleteAppSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminDeleteAppSubscriptionResponse{
		Info: info,
	}, nil
}

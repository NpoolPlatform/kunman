package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/good/app/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription"
)

func (s *Server) UpdateAppSubscription(ctx context.Context, in *npool.UpdateAppSubscriptionRequest) (*npool.UpdateAppSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(&in.ID, true),
		subscription1.WithEntID(&in.EntID, true),
		subscription1.WithAppID(&in.AppID, true),
		subscription1.WithAppGoodID(&in.AppGoodID, true),

		subscription1.WithName(in.Name, false),
		subscription1.WithBanner(in.Banner, false),

		subscription1.WithEnableSetCommission(in.EnableSetCommission, false),
		subscription1.WithUSDPrice(in.USDPrice, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateSubscription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateAppSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateAppSubscriptionResponse{
		Info: info,
	}, nil
}

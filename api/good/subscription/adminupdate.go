package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/good/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/subscription"
)

func (s *Server) AdminUpdateSubscription(ctx context.Context, in *npool.AdminUpdateSubscriptionRequest) (*npool.AdminUpdateSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(&in.ID, true),
		subscription1.WithEntID(&in.EntID, true),
		subscription1.WithGoodID(&in.GoodID, true),
		subscription1.WithName(in.Name, false),
		subscription1.WithDurationDisplayType(in.DurationDisplayType, false),
		subscription1.WithDurationUnits(in.DurationUnits, false),
		subscription1.WithDurationQuota(in.DurationQuota, false),
		subscription1.WithDailyBonusQuota(in.DailyBonusQuota, false),
		subscription1.WithUSDPrice(in.USDPrice, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateSubscription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminUpdateSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminUpdateSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminUpdateSubscriptionResponse{
		Info: info,
	}, nil
}

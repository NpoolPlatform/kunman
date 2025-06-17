package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/good/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/subscription"
)

func (s *Server) AdminCreateSubscription(ctx context.Context, in *npool.AdminCreateSubscriptionRequest) (*npool.AdminCreateSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithDurationDisplayType(&in.DurationDisplayType, true),
		subscription1.WithName(&in.Name, true),
		subscription1.WithDurationUnits(in.DurationUnits, false),
		subscription1.WithDurationQuota(&in.DurationQuota, false),
		subscription1.WithDailyBonusQuota(in.DailyBonusQuota, false),
		subscription1.WithUSDPrice(&in.USDPrice, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateSubscription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateSubscriptionResponse{
		Info: info,
	}, nil
}

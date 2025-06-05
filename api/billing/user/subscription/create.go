package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/billing/user/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gateway/v1/user/subscription"
)

func (s *Server) AdminCreateSubscription(ctx context.Context, in *npool.AdminCreateSubscriptionRequest) (*npool.AdminCreateSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(&in.TargetAppID, true),
		subscription1.WithUserID(&in.TargetUserID, true),
		subscription1.WithPackageID(in.PackageID, false),
		subscription1.WithStartAt(in.StartAt, false),
		subscription1.WithEndAt(in.EndAt, false),
		subscription1.WithUsageState(&in.UsageState, true),
		subscription1.WithSubscriptionCredit(&in.SubscriptionCredit, true),
		subscription1.WithAddonCredit(&in.AddonCredit, true),
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

package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/billing/user/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gateway/v1/user/subscription"
)

func (s *Server) AdminUpdateSubscription(ctx context.Context, in *npool.AdminUpdateSubscriptionRequest) (*npool.AdminUpdateSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(&in.ID, true),
		subscription1.WithEntID(&in.EntID, true),
		subscription1.WithAppID(&in.TargetAppID, true),
		subscription1.WithUserID(&in.TargetUserID, true),
		subscription1.WithStartAt(in.StartAt, false),
		subscription1.WithEndAt(in.EndAt, false),
		subscription1.WithUsageState(in.UsageState, false),
		subscription1.WithSubscriptionCredit(in.SubscriptionCredit, false),
		subscription1.WithAddonCredit(in.AddonCredit, false),
		subscription1.WithPackageID(in.PackageID, false),
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

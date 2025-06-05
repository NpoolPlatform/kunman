package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/billing/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gw/v1/subscription"
)

func (s *Server) AdminUpdateSubscription(ctx context.Context, in *npool.AdminUpdateSubscriptionRequest) (*npool.AdminUpdateSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithID(&in.ID, true),
		subscription1.WithEntID(&in.EntID, true),
		subscription1.WithAppID(&in.TargetAppID, false),
		subscription1.WithPackageName(in.PackageName, false),
		subscription1.WithUsdPrice(in.UsdPrice, false),
		subscription1.WithDescription(in.Description, false),
		subscription1.WithSortOrder(in.SortOrder, false),
		subscription1.WithCredit(in.Credit, false),
		subscription1.WithResetType(in.ResetType, false),
		subscription1.WithQPSLimit(in.QPSLimit, false),
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

package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/billing/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/billing/gateway/v1/subscription"
)

func (s *Server) AdminCreateSubscription(ctx context.Context, in *npool.AdminCreateSubscriptionRequest) (*npool.AdminCreateSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(&in.TargetAppID, true),
		subscription1.WithPackageName(in.PackageName, true),
		subscription1.WithUsdPrice(in.UsdPrice, true),
		subscription1.WithDescription(in.Description, true),
		subscription1.WithSortOrder(in.SortOrder, true),
		subscription1.WithPackageType(in.PackageType, true),
		subscription1.WithCredit(in.Credit, true),
		subscription1.WithResetType(in.ResetType, true),
		subscription1.WithQPSLimit(in.QPSLimit, true),
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

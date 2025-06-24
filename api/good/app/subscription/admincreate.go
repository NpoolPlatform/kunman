package subscription

import (
	"context"

	subscription1 "github.com/NpoolPlatform/kunman/gateway/good/app/subscription"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/subscription"
)

func (s *Server) AdminCreateAppSubscription(ctx context.Context, in *npool.AdminCreateAppSubscriptionRequest) (*npool.AdminCreateAppSubscriptionResponse, error) {
	handler, err := subscription1.NewHandler(
		ctx,
		subscription1.WithAppID(&in.TargetAppID, true),
		subscription1.WithGoodID(&in.GoodID, true),

		subscription1.WithName(&in.Name, true),
		subscription1.WithBanner(in.Banner, false),

		subscription1.WithEnableSetCommission(in.EnableSetCommission, false),
		subscription1.WithUSDPrice(&in.USDPrice, true),

		subscription1.WithProductID(in.ProductID, false),
		subscription1.WithTrialUnits(in.TrialUnits, false),
		subscription1.WithTrialUSDPrice(in.TrialUSDPrice, false),

		subscription1.WithPriceFiatID(in.PriceFiatID, false),
		subscription1.WithFiatPrice(in.FiatPrice, false),
		subscription1.WithTrialFiatPrice(in.TrialFiatPrice, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateSubscription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppSubscription",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppSubscriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateAppSubscriptionResponse{
		Info: info,
	}, nil
}

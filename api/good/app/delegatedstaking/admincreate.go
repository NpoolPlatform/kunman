package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/kunman/gateway/good/app/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/kunman/message/good/gateway/v1/app/delegatedstaking"
)

func (s *Server) AdminCreateAppDelegatedStaking(ctx context.Context, in *npool.AdminCreateAppDelegatedStakingRequest) (*npool.AdminCreateAppDelegatedStakingResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithAppID(&in.TargetAppID, true),
		delegatedstaking1.WithGoodID(&in.GoodID, true),

		delegatedstaking1.WithPurchasable(in.Purchasable, false),
		delegatedstaking1.WithEnableProductPage(in.EnableProductPage, false),
		delegatedstaking1.WithProductPage(in.ProductPage, false),
		delegatedstaking1.WithOnline(in.Online, false),
		delegatedstaking1.WithVisible(in.Visible, false),
		delegatedstaking1.WithName(&in.Name, true),
		delegatedstaking1.WithDisplayIndex(in.DisplayIndex, false),
		delegatedstaking1.WithBanner(in.Banner, false),

		delegatedstaking1.WithServiceStartAt(&in.ServiceStartAt, true),
		delegatedstaking1.WithStartMode(in.StartMode, false),
		delegatedstaking1.WithEnableSetCommission(in.EnableSetCommission, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDelegatedStaking(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AdminCreateAppDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.AdminCreateAppDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AdminCreateAppDelegatedStakingResponse{
		Info: info,
	}, nil
}
